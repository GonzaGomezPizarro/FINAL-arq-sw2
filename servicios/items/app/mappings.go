package app

import (
	controller "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/items", controller.GetItems)
	router.GET("/item/:id", controller.GetItemById)

	router.POST("/item", controller.NewItem)
	router.POST("/items", controller.NewItems)

	router.DELETE("/item/:id", controller.DeleteItem)

	log.Info("Finishing mappings configurations")
}
