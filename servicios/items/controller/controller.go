package controller

import (
	"net/http"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/dto"
	service "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/service"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func GetItemById(c *gin.Context) {
	log.Debug("Item id to load: " + c.Param("id"))
	id := c.Param("id")
	var itemDto dto.Item

	itemDto, err := service.ItemService.GetItemById(id)

	if err != nil {
		c.JSON(err.Status(), dto.Item{})
		return
	}
	c.JSON(http.StatusOK, itemDto)
}

func GetItems(c *gin.Context) {
	var itemsDto dto.Items
	itemsDto, err := service.ItemService.GetItems()

	if err != nil {
		c.JSON(err.Status(), dto.Items{})
		return
	}

	c.JSON(http.StatusOK, itemsDto)
}

func NewItem(c *gin.Context) {
	var itemDto dto.Item
	err := c.BindJSON(&itemDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Item{})
		return
	}

	itemDto, errr := service.ItemService.NewItem(itemDto)

	if errr != nil {
		c.JSON(errr.Status(), dto.Item{})
		return
	}

	c.JSON(http.StatusCreated, itemDto)
}

func NewItems(c *gin.Context) {
	var itemsDto dto.Items
	err := c.BindJSON(&itemsDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Items{})
		return
	}

	itemsDto, errr := service.ItemService.NewItems(itemsDto)

	if errr != nil {
		c.JSON(errr.Status(), dto.Items{})
		return
	}

	c.JSON(http.StatusCreated, itemsDto)
}

func DeleteItem(c *gin.Context) {
	// Obtener el ID del item a eliminar desde los parámetros de la URL
	itemID := c.Param("id")

	// Llamar al servicio para eliminar el usuario
	errr := service.ItemService.DeleteItem(itemID)
	if errr != nil {
		c.Status(errr.Status())
		return
	}

	// Si se eliminó el usuario correctamente, devolver una respuesta 204 (sin contenido)
	c.Status(http.StatusNoContent)
}
