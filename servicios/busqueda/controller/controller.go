package controller

import (
	"log"
	"net/http"
	"strconv"

	cliente "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/client"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/dto"
	"github.com/gin-gonic/gin"
)

func GetQuery(c *gin.Context) {
	var itemsDto dto.Items
	query := c.Param("searchQuery")

	itemsDto, err := cliente.GetQuery(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, itemsDto)
		log.Println(err.Error())
		return
	}

	itemsDto.ConvertirImagenes()

	c.JSON(http.StatusOK, itemsDto)
}

func GetAll(c *gin.Context) {
	var itemsDto dto.Items

	itemsDto, err := cliente.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, itemsDto)
		log.Println(err.Error())
		return
	}

	itemsDto.ConvertirImagenes()

	c.JSON(http.StatusOK, itemsDto)

}

func GetItemById(c *gin.Context) {
	id := c.Param("id")
	item, err := cliente.GetItemById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Println(err.Error())
		return
	}

	item.ConvertirImagenes()

	c.JSON(http.StatusOK, item)
}

func GetItemsByUserId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("UserId"))
	items, err := cliente.GetItemByUserId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Println(err.Error())
		return
	}

	items.ConvertirImagenes()

	c.JSON(http.StatusOK, items)
}
