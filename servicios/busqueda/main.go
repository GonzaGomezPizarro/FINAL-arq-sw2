package main

import (
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/app"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/motordebusqueda"
	notificacion "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/notificaciones"
)

func main() {
	motordebusqueda.Indexall()

	// Iniciar la escucha de mensajes en una goroutine
	messages := make(chan string)
	go notificacion.Receive(messages)

	app.StartRoute()
}
