package app

import (
	controller "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/users", controller.GetUsers)
	router.POST("/user", controller.NewUser)
	router.GET("/user/:id", controller.GetUserById)
	router.DELETE("/user/:id", controller.DeleteUser)

	log.Info("Finishing mappings configurations")
}
