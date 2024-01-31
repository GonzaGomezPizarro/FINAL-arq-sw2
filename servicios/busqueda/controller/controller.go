package controller

import (
	"net/http"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/dto"
	elasticsearch "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/elasticsearch"
	"github.com/gin-gonic/gin"
)

func GetQuery(c *gin.Context) {
	var itemsDto dto.Items
	query := c.Param("searchQuery")

	itemsDto, err := elasticsearch.GetQuery(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, itemsDto)
		return
	}

	c.JSON(http.StatusOK, itemsDto)

}

func GetAll(c *gin.Context) {
	var itemsDto dto.Items

	itemsDto, err := elasticsearch.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, itemsDto)
		return
	}

	c.JSON(http.StatusOK, itemsDto)

}

func GetItemById(c *gin.Context) {
	id := c.Param("id")
	item, err := elasticsearch.GetItemById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, item)
}
