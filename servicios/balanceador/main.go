package main

import (
	"log"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/balanceador/app"
	cacheLocal "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/balanceador/cachelocal"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/balanceador/cola"
)

func main() {
	cacheLocal.InitCache()

	err := cola.CheckConection()
	if err != nil {
		log.Println(err)
		return
	}

	app.StartRoute()
}
