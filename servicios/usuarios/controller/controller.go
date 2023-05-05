package controller

import (
	"net/http"
	"strconv"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/dto"

	service "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/service"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

var jwtKey = []byte("secret_key")

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto.User

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		log.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	var usersDto dto.Users
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, usersDto) // esto es lo que no esta andando !!!!!!!!!!!
}

func NewUser(c *gin.Context) {
	var userDto dto.User
	err := c.BindJSON(&userDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, errr := service.UserService.NewUser(userDto)

	if errr != nil {
		c.JSON(errr.Status(), errr)
		log.Error(errr.Error())
		return
	}

	c.JSON(http.StatusOK, userDto)
}

func DeleteUser(c *gin.Context) {
	// Obtener el ID del usuario a eliminar desde los parámetros de la URL
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Llamar al servicio para eliminar el usuario
	errr := service.UserService.DeleteUser(userID)
	if errr != nil {
		c.JSON(500, errr)
		log.Error(errr.Error())
		return
	}

	// Si se eliminó el usuario correctamente, devolver una respuesta 204 (sin contenido)
	c.Status(http.StatusNoContent)
}
