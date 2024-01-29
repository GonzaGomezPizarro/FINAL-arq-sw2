package app

import (
	controller "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/messages", controller.GetMessages)
	router.GET("/messagesByUser/:id", controller.GetMessagesByUserId)
	router.GET("/message/:id", controller.GetMessageById)
	router.GET("/messagesByItem/:id", controller.GetMessagesByItemId)

	router.POST("/message/", controller.PostMessage)
	router.POST("/messages", controller.PostMessages)

	router.DELETE("/message/:id", controller.DeleteMessageById)

	log.Info("Finishing mappings configurations")
}
