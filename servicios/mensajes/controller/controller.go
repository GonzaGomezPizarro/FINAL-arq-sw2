package controller

import (
	"net/http"
	"strconv"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/dto"

	service "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/service"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

var jwtKey = []byte("secret_key")

func GetMessages(c *gin.Context) {
	MessagesDto, err := service.MessageService.GetMessages()

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, MessagesDto)
}
func GetMessagesByUserId(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := c.Param("id")
	var messagesDto dto.Messages

	messagesDto, err := service.MessageService.GetMessagesByUserId(id)

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		log.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, messagesDto)
}
func GetMessageById(c *gin.Context) {
	log.Debug("Message id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var messageDto dto.Message

	messageDto, err := service.MessageService.GetMessageById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		log.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, messageDto)
}
func GetMessageByItemId(c *gin.Context {
	log.Debug("Item id to load: " + c.Param("id"))

	id, _ := strconv
}
func PostMessage(c *gin.Context) {
	var messageDto dto.Message
	err := c.BindJSON(&messageDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	messageDto, errr := service.MessageService.PostMessage(messageDto)

	if errr != nil {
		c.JSON(errr.Status(), errr)
		log.Error(errr.Error())
		return
	}

	c.JSON(http.StatusOK, messageDto)
}
func DeleteMessageById(c *gin.Context) {
	// Obtener el ID del mensaje a eliminar desde los parámetros de la URL
	messageIDStr := c.Param("id")
	messageID, err := strconv.Atoi(messageIDStr)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Llamar al servicio para eliminar el usuario
	errr := service.MessageService.DeleteMessageById(messageID)
	if errr != nil {
		c.JSON(500, errr)
		log.Error(errr.Error())
		return
	}

	// Si se eliminó el usuario correctamente, devolver una respuesta 204 (sin contenido)
	c.Status(http.StatusNoContent)
}
