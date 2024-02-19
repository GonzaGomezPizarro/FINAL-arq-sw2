package app

import (
	controller "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/contenedores/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/imagenes", controller.GetImages)
	router.GET("/contenedores", controller.GetContenedores)

	router.POST("/contenedor", controller.NewContenedor)

	router.PUT("/Pcontenedor/:id", controller.PlayContenedor)
	router.PUT("/Scontenedor/:id", controller.StopContenedor)

	router.DELETE("/contenedor/:id", controller.DeleteContenedor)

	log.Info("Finishing mappings configurations")
}
