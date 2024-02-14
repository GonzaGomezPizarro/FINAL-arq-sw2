package main

import (
	"log"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/balanceador/app"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/balanceador/cola"
)

func main() {
	err := cola.CheckConection()
	if err != nil {
		log.Println(err)
		return
	}

	app.StartRoute()
}
