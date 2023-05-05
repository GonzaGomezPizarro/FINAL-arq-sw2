package controller

import (
	"net/http"
	"strconv"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/dto"
	service "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/service"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

var jwtKey = []byte("secret_key")

func GetItemById(c *gin.Context) {
	log.Debug("Item id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var itemDto dto.Item

	itemDto, err := service.ItemService.GetItemById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		log.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, itemDto)
}

func GetItems(c *gin.Context) {
	var usersDto dto.Items
	usersDto, err := service.ItemService.GetItems()

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		log.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func NewItem(c *gin.Context) {
	var itemDto dto.Item
	err := c.BindJSON(&itemDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemDto, errr := service.ItemService.NewItem(itemDto)

	if errr != nil {
		c.JSON(errr.Status(), errr)
		log.Error(errr.Error())
		return
	}

	c.JSON(http.StatusOK, itemDto)
}

func NewItems(c *gin.Context) {
	var itemsDto dto.Items
	err := c.BindJSON(&itemsDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemsDto, errr := service.ItemService.NewItems(itemsDto)

	if errr != nil {
		c.JSON(errr.Status(), errr)
		log.Error(errr.Error())
		return
	}

	c.JSON(http.StatusOK, itemsDto)
}

func DeleteItem(c *gin.Context) {
	// Obtener el ID del item a eliminar desde los parámetros de la URL
	itemIDStr := c.Param("id")
	itemID, err := strconv.Atoi(itemIDStr)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Llamar al servicio para eliminar el usuario
	errr := service.ItemService.DeleteItem(itemID)
	if errr != nil {
		c.JSON(500, errr)
		log.Error(errr.Error())
		return
	}

	// Si se eliminó el usuario correctamente, devolver una respuesta 204 (sin contenido)
	c.Status(http.StatusNoContent)
}
