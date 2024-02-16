package app

import (
	controller "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/contenedores/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/imagenes", controller.GetImages)
	router.GET("/contenedores", controller.GetContenedores)

	router.POST("/contenedor", controller.NewContenedor)

	router.DELETE("/contenedor/:id", controller.DeleteItem)

	log.Info("Finishing mappings configurations")
}
