package app

import (
	controller "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/balanceador/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	//items

	router.GET("/items", controller.GetToItems)
	router.GET("/item/:id", controller.GetToItems)

	router.POST("/item", controller.PostToItems)
	router.POST("/items", controller.PostToItems)

	router.DELETE("/item/:id", controller.DeleteToItems)

	//mensajes

	// router.GET("/messages", controller.GetToMensajes)
	// router.GET("/messagesByUser/:id", controller.GetToMensajes) //no esta devolviendo mensajes
	// router.GET("/message/:id", controller.GetToMensajes)
	// router.GET("/messagesByItem/:id", controller.GetToMensajes)

	// router.POST("/message", controller.PostToMensajes)
	// router.POST("/messages", controller.PostToMensajes)

	// router.DELETE("/message/:id", controller.DeleteToMensajes)

	// //usuarios

	// router.GET("/users", controller.GetToUsuarios)
	// router.GET("/user/:id", controller.GetToUsuarios)

	// router.DELETE("/user/:id", controller.DeleteToUsuarios)

	// router.POST("/user", controller.PostToUsuarios)
	// router.POST("/login", controller.PostToUsuarios)

	// //busqueda

	// router.GET("/search/:searchQuery", controller.GetToBusqueda)
	// router.GET("/searchAll", controller.GetToBusqueda)
	// router.GET("/item/:id", controller.GetToBusqueda)
	// router.GET("/items/:UserId", controller.GetToBusqueda)

	log.Info("Finishing mappings configurations")
}
