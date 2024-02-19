package controller

import (
	"net/http"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/contenedores/docker"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/contenedores/dto"
	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	imagenes, err := docker.GetImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, imagenes)
}

func GetContenedores(c *gin.Context) {
	contenedores, err := docker.GetContenedores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, contenedores)
}

func NewContenedor(c *gin.Context) {
	var contenedor dto.Contenedor
	err := c.BindJSON(&contenedor)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	contenedor, errr := docker.PostContenedor(contenedor)
	if errr != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, contenedor)
}

func PlayContenedor(c *gin.Context) {
	id := c.Param("id")

	err := docker.StartContenedor(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func StopContenedor(c *gin.Context) {
	id := c.Param("id")

	err := docker.StopContenedor(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func DeleteContenedor(c *gin.Context) {
	id := c.Param("id")

	err := docker.DeleteContenedor(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)

}
