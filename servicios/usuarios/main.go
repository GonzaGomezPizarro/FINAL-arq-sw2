package main

import (
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/app"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/db"
)

func main() {

	db.StartDbEngine()
	app.StartRoute()
}
