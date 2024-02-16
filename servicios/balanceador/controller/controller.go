package controller

import (
	"log"
	"net/http"
	"strings"

	cacheLocal "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/balanceador/cachelocal"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/balanceador/cola"
	"github.com/gin-gonic/gin"
)

func GetToItems(c *gin.Context) {
	url := c.Request.URL.String() // Obtener la URL solicitada
	log.Println("url to get: ", url)

	jsonn := "" // vacio por que no se usa

	if url == "/items" {
		items, httpStatusCode, err := cola.SendToItems("GET", url, jsonn)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			log.Println(err.Error())
			return
		}
		c.JSON(httpStatusCode, items)
		return
	}

	// Si la URL es para obtener un ítem específico
	// Obtener el ID del ítem de la URL
	parts := strings.Split(url, "/")
	if len(parts) != 3 {
		// La URL no tiene el formato esperado, devolver un error
		c.JSON(http.StatusBadRequest, "URL no válida")
		log.Println("URL no válida")
		return
	}
	itemID := parts[2]

	// verificar la caché local primero
	itemm, bul := cacheLocal.CacheInstance.Get(itemID)
	if bul {
		c.JSON(http.StatusOK, itemm)
		return
	}

	//si no esta en cache
	item, httpStatusCode, err := cola.SendToItems("GET", url, jsonn)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	//guardo en cachelocal
	cacheLocal.CacheInstance.Set(item[0])

	c.JSON(httpStatusCode, item)
}

func PostToItems(c *gin.Context) {
	url := c.Request.URL.String() // Obtener la URL solicitada
	log.Println("url to post: ", url)

	jsonBody, err := c.GetRawData() // Obtener el JSON del cuerpo como bytes
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		log.Println("Error reading request body:", err.Error())
		return
	}

	jsonString := string(jsonBody) // Convertir los bytes en una cadena JSON

	items, httpStatusCode, err := cola.SendToItems("POST", url, jsonString)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	//guardo en cachelocal
	cacheLocal.CacheInstance.Set(items[0])

	c.JSON(httpStatusCode, items)
}

func DeleteToItems(c *gin.Context) {
	url := c.Request.URL.String() // Obtener la URL solicitada
	log.Println("url to delete: ", url)

	jsonn := "" //vacio por q no se usa

	items, httpStatusCode, err := cola.SendToItems("DELETE", url, jsonn)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	//guardo en cachelocal
	cacheLocal.CacheInstance.Delete(items[0].Id)

	c.JSON(httpStatusCode, items)
}

// func GetToUsuarios(c *gin.Context) {

// }

// func PostToUsuarios(c *gin.Context) {

// }

// func DeleteToUsuarios(c *gin.Context) {

// }

// func GetToMensajes(c *gin.Context) {

// }

// func PostToMensajes(c *gin.Context) {

// }

// func DeleteToMensajes(c *gin.Context) {

// }

// func GetToBusqueda(c *gin.Context) {

// }
