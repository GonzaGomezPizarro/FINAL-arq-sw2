package app

import (
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/search/:searchQuery", controller.GetQuery)
	router.GET("/searchAll", controller.GetAll)
	router.GET("/item/:id", controller.GetItemById)
	router.GET("/items/:UserId", controller.GetItemsByUserId)

	log.Info("Finishing mappings configurations")
}
