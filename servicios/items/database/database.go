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
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// Agrega la credencial al objeto clientOptions
	clientOptions.Auth = &options.Credential{
		Username: "root",
		Password: "CONTRASENA",
	}

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
	database = client.Database("itemss")

	fmt.Println("Conexión a la base de datos establecida correctamente.")
}

// StartDBEngine inicia la conexión con la base de datos MongoDB
func StartDBEngine() *mongo.Database {
	return database
}
