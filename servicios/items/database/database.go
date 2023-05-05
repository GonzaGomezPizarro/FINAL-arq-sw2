package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

func init() {
	// Configura la conexión a la base de datos
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conecta con la base de datos
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Comprueba la conexión
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Establece la base de datos por defecto
	database = client.Database("nombre_de_la_base_de_datos")

	fmt.Println("Conexión a la base de datos establecida correctamente.")
}

// StartDBEngine inicia la conexión con la base de datos MongoDB
func StartDBEngine() *mongo.Database {
	return database
}
