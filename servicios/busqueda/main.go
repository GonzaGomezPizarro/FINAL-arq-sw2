package main

import (
	"log"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/app"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/motordebusqueda"
	notificacion "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/notificaciones"
)

func main() {
	errr := motordebusqueda.Check() // iniciamos la conexcion con elasticsearch
	if errr != nil {
		panic(errr)
	}
	log.Println("-> Connectado a elasticsearch")

	// indexamos la base de datos
	err := motordebusqueda.IndexAll()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err := notificacion.Receive()
			if err != nil {
				log.Println("Error recibiendo notificación:", err)
				log.Println("Intentando volver a iniciar la recepción de notificaciones...")
			}
		}
	}()

	app.StartRoute()

}
