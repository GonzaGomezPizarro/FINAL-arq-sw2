package main

import (
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/app"
	cacheLocal "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/cachelocal"
)

func main() {
	cacheLocal.InitCache()
	app.StartRoute()
}
