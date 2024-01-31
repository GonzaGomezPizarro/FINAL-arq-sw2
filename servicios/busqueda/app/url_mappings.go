package app

import (
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/search/:searchQuery", controller.GetQuery)
	router.GET("/searchAll", controller.GetAll)
	router.GET("/items/:id", controller.GetItemById)

	log.Info("Finishing mappings configurations")
}
