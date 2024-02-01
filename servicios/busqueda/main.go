package main

import (
	"log"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/app"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/motordebusqueda"
	notificacion "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/notificaciones"
)

func main() {
	errr := motordebusqueda.StartSearchEngine() // iniciamos la conexcion con elasticsearch
	if errr != nil {
		panic(errr)
	}
	log.Println("-> Connectado a elasticsearch")

	// indexamos la base de datos
	err := motordebusqueda.IndexAll()
	if err != nil {
		panic(err)
	}
	log.Println("-> Items indexados")

	// Iniciar la escucha de mensajes en una goroutine
	messages := make(chan string)
	go notificacion.Receive(messages)

	app.StartRoute()
}
