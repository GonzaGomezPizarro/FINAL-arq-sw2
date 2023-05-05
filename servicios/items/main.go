package main

import (
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/app"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/database"
)

func main() {
	database.StartDbEngine()
	app.StartRoute()
}
