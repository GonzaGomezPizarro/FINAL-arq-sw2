package app

import (
	// userController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/user" -- Linea original
	controller "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/controller" // Puede ser que haga falta agregar / user
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/users", controller.GetUsers)
	router.POST("/user", controller.NewUser)
	router.GET("/user/:id", controller.GetUserById)
	router.DELETE("/user/:id", controller.DeleteUser)

	log.Info("Finishing mappings configurations")
}
