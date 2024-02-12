package client

import (
	"context"
	"errors"
	"log"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/cache"
	cacheLocal "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/cachelocal"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/database"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/model"
	notificacion "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/notificaciones"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	e "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/errors"
)

func GetItemById(id string) (model.Item, error) {
	//intentar obtener item from cchelocal
	i, bul := cacheLocal.CacheInstance.Get(id)
	if bul == true {
		log.Println(" -> Encontrado en cacheLocal")
		return i, nil
	}

	// Intenta obtener el objeto de Memcached
	cachedItem, err := GetItemFromCache(id)
	if err == nil {
		bul := cacheLocal.CacheInstance.Set(cachedItem)
		if bul == true {
			log.Println(" -> Almacenado en cacheLocal")
		} else {
			log.Println(" -> Error al almacenar en cacheLocal")
		}
		return cachedItem, nil
	}

	// Si no se encuentra en Memcached, busca en la base de datos
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
	log.Println(" -> Encontrdo en BD")
	// Almacena el objeto en Memcached solo si se encontró en la base de datos
	cachedItem, errr := InsertItemToCache(item, item.Id.Hex())
	if errr != nil {
		log.Println(errr)
	}

	// insertar en cacheLocal
	ii := cacheLocal.CacheInstance.Set(item)
	if ii == true {
		log.Println(" -> Almacenado en cacheLocal")
	} else {
		log.Println(" -> Error al almacenar en cacheLocal")
	}

	//

	return item, nil
}

func GetItems() (model.Items, error) { // voy directo a base de datos por que trae todos los items de la base de datos
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

	// Guarda los primeros 50 items en la caché
	for i := 0; i < 50 && i < len(items); i++ {
		_, err := InsertItemToCache(items[i], items[i].Id.Hex())
		if err != nil {
			log.Println(" -> Error al almacenar en la caché:", err)
		}
	}

	// guarda los primeros 10 items en cachelocal
	for i := 0; i < 10 && i < len(items); i++ {
		bul := cacheLocal.CacheInstance.Set(items[i])
		if bul == false {
			log.Println(" -> Error al almacecenar en cacheLocal: ", items[i])
		} else {
			log.Println(" -> Almacenado en cacheLocal", items[i])
		}
	}

	return items, nil
}

func NewItem(item model.Item) (model.Item, e.ApiError) {
	// Insertar el ítem en la base de datos
	db := database.StartDBEngine()
	collection := db.Collection("items")

	res, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return model.Item{}, e.NewInternalServerApiError("Error al crear el item", err)
	}
	log.Println("insertado en BD")

	// Obtener el ID insertado y asignarlo al ítem
	objectID := res.InsertedID.(primitive.ObjectID)
	item.Id = objectID

	// Guardar el ítem en la caché
	_, err = InsertItemToCache(item, item.Id.Hex())
	if err != nil {
		log.Println("Error al almacenar en caché:", err)
	}

	// Guardar item en cacheLocal
	bul := cacheLocal.CacheInstance.Set(item)
	if bul == true {
		log.Println(" -> Almacenado en cacheLocal")
	} else {
		log.Println(" -> Error al almacenar en cacheLocal")
	}

	//notificamos que se modifico un nuevo item
	id := item.Id.Hex()
	notificacion.Send(id)

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

	log.Println("items almacenados en BD")

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

	// Guarda los primeros 50 items en la caché
	for i := 0; i < 50 && i < len(items); i++ {
		_, err := InsertItemToCache(items[i], items[i].Id.Hex())
		if err != nil {
			log.Println("Error al almacenar en la caché:", err)
		}
	}

	// Guarda los primeros 10 items en la cachéLocal
	for i := 0; i < 10 && i < len(items); i++ {
		bul := cacheLocal.CacheInstance.Set(items[i])
		if bul == true {
			log.Println(" -> Almacenado en cacheLocal")
		} else {
			log.Println(" -> Error al almacenar en cacheLocal")
		}

	}

	for i := 0; i < len(items); i++ {
		//notificamos que se modifico cada item
		id := items[i].Id.Hex()
		notificacion.Send(id)
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

	log.Println("item deleted from DB", objectId)

	//notificamos que se modifico un nuevo item
	notificacion.Send(itemId)

	// Borrar el item de la caché
	err = cache.DeleteFromCache("item:" + itemId)
	if err != nil {
		log.Println("Error al borrar de la caché:", err)
	} else {
		log.Println("item deleted from cache", objectId)
	}

	//borrar el item de la cacheLocal
	cacheLocal.CacheInstance.Delete(itemId)

	return nil
}

func InsertItemToCache(item model.Item, id string) (model.Item, error) {
	// Almacena el objeto en Memcached
	bsonItem, err := bson.Marshal(item)
	if err != nil {
		log.Println("Error marshaling item to BSON:", err)
		return model.Item{}, errors.New("Error marshaling item to BSON")
	}

	err = cache.SetToCache("item:"+id, string(bsonItem))
	if err != nil {
		log.Println("Error almacenando en Memcached:", err)
		return model.Item{}, errors.New("Error almacenando en Memcached")
	}

	log.Println(" -> Item almacenado en cache")

	// Devuelve el item
	return item, nil
}

func GetItemFromCache(id string) (model.Item, error) {
	// Intenta obtener el objeto de Memcached
	cachedItem, err := cache.GetFromCache("item:" + id)
	if err != nil {
		return model.Item{}, errors.New("Item not found in cache")
	}

	// Decodifica el objeto BSON
	var item model.Item
	err = bson.Unmarshal([]byte(cachedItem), &item)
	if err != nil {
		log.Println("Error decoding cached item:", err)
		return model.Item{}, errors.New("Error decoding cached item")
	}

	log.Println(" -> Encontrado en Memcached")
	return item, nil
}
