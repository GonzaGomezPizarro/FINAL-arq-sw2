package client

import (
	"context"
	"errors"
	"log"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/database"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	e "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/errors"
)

func GetItemById(id string) (model.Item, error) {
	db := database.StartDBEngine()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Item{}, errors.New("Invalid ID format")
	}

	var item model.Item
	err = db.Collection("items").FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&item)
	if err != nil {
		log.Println(err)
		return model.Item{}, errors.New("Item not found")
	}

	return item, nil
}

func GetItems() (model.Items, error) {
	db := database.StartDBEngine()

	var items model.Items
	cursor, err := db.Collection("items").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
		return items, errors.New("Error fetching items")
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var item model.Item
		if err := cursor.Decode(&item); err != nil {
			log.Println(err)
			continue
		}
		items = append(items, item)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return items, errors.New("Error fetching items")
	}

	return items, nil
}

func NewItem(item model.Item) (model.Item, e.ApiError) {
	db := database.StartDBEngine()
	collection := db.Collection("items")

	res, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return model.Item{}, e.NewInternalServerApiError("Error al crear el item", err)
	}

	objectID := res.InsertedID.(primitive.ObjectID)
	item.Id = objectID

	return item, nil
}

func NewItems(items model.Items) (model.Items, e.ApiError) {
	db := database.StartDBEngine()
	collection := db.Collection("items")

	// Inserta los items en la base de datos
	documents := make([]interface{}, len(items))
	for i, item := range items {
		documents[i] = item
	}
	result, err := collection.InsertMany(context.Background(), documents)
	if err != nil {
		return nil, e.NewApiError("Error al insertar los items", err.Error(), 500, e.CauseList{})
	}

	// Obtiene los IDs asignados por MongoDB a los nuevos items
	objectIds := result.InsertedIDs
	if len(objectIds) != len(items) {
		return nil, e.NewInternalServerApiError("Error al obtener los IDs de los items insertados", nil)
	}
	for i := 0; i < len(items); i++ {
		id, ok := objectIds[i].(primitive.ObjectID)
		if !ok {
			return nil, e.NewInternalServerApiError("Error al obtener los IDs de los items insertados", nil)
		}
		items[i].Id = id
	}

	return items, nil
}

func DeleteItem(itemId string) e.ApiError {
	db := database.StartDBEngine()
	collection := db.Collection("items")

	objectId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return e.NewNotFoundApiError("No se encontró el item")
	}

	filter := bson.M{"_id": objectId}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return e.NewInternalServerApiError("Error al eliminar el item", err)
	}

	return nil
}
