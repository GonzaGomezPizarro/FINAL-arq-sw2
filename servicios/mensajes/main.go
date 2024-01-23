package main

import (
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/app"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/db"
)

func main() {

	db.StartDbEngine()
	app.StartRoute()
}
