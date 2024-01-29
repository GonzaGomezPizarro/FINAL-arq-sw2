package client

import (
	"errors"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/model"

	e "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetMessageById(id int) model.Message {
	var message model.Message
	Db.Where("id = ?", id).First(&message)
	log.Debug("Message: ", message)
	return message
}

func GetMessagesByUserId(userId int) model.Messages {
	var messages model.Messages
	Db.Where("user_id = ?", userId).Find(&messages)
	log.Debug("Messages: ", messages)
	return messages
}

func GetMessagesByItemId(itemId string) model.Messages {
	var messages model.Messages
	Db.Where("item_id = ?", itemId).Find(&messages)
	log.Debug("Messages: ", messages)
	return messages
}

func GetMessages() model.Messages {
	var messages model.Messages
	Db.Find(&messages)
	log.Debug("Messages: ", messages)
	return messages
}

func NewMessage(message model.Message) (model.Message, e.ApiError) {
	// Asumiendo que Db está configurado y es accesible
	result := Db.Create(&message)

	if result.Error != nil {
		log.Error("Error al crear un nuevo mensaje:", result.Error)
		return model.Message{}, e.NewInternalServerApiError("Error al crear un nuevo mensaje", result.Error)
	}

	log.Debug("Nuevo mensaje creado:", message)
	return message, nil
}

func DeleteMessage(id int) (model.Message, e.ApiError) {
	var message model.Message

	// Buscar el mensaje por ID
	result := Db.First(&message, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Message{}, e.NewNotFoundApiError("Message does not exist")
		} else {
			return model.Message{}, e.NewInternalServerApiError("Error while searching for message", result.Error)
		}
	}

	// Almacenar copia del mensaje para devolverlo después de la eliminación
	deletedMessage := message

	// Eliminar el mensaje
	result = Db.Delete(&message)
	if result.Error != nil {
		return model.Message{}, e.NewInternalServerApiError("Error while deleting message", result.Error)
	}

	// Devolver el mensaje eliminado
	return deletedMessage, nil
}
