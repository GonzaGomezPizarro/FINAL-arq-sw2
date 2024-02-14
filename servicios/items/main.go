package main

import (
	"log"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/app"
	cacheLocal "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/cachelocal"
	notificacion "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/notificaciones"
)

func main() {
	cacheLocal.InitCache()

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
