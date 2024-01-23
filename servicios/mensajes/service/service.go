package service

import (
	cliente "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/client"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/dto"
	e "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/errors"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/model"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/funciones"
)

type messageService struct{}

type messageServiceInterface interface {
	GetMessageById(id int) (dto.Message, e.ApiError)
	GetMessages() (dto.Messages, e.ApiError)
	GetMessagesByUserId(id int) (dto.Messages, e.ApiError)
	GetMessagesByItemId(id string) (dto.Messages, e.ApiError)

	PostMessage(messageDto dto.Message) (dto.Message, e.ApiError)

	DeleteMessageById(id int) e.ApiError
}

var jwtKey = []byte("secret_key")

var (
	MessageService messageServiceInterface
)

func init() {
	MessageService = &umessageService{}
}

func (s *messageService) GetMessageById(id int) (dto.Message, e.apiError) {
	var message model.Message = cliente.GetMessageById(id)
	var messageDto dto.Message

	if message.Id == 0 {
		return messageDto, e.NewBadRequestApiError("message not found")
	}

	messageDto.Id = message.Id
	messageDto.Content = message.Content
	messageDto.ItemId = message.ItemId
	messageDto.ReceiverId = message.ReceiverId

	return messageDto, nil
}

func (s *messageService) GetMessages() (dto.Messages, error) {
	var messages model.Message = cliente.GetMessages()
	var messagesDto dto.Messages

	for _, message := range messages {
		var messageDto dto.Message
		messageDto.Id = message.Id
		messageDto.Content = message.Content
		messageDto.ReceiverId = message.ReceiverId
		messagesDto = append(messagesDto, messageDto)
	}

	return messagesDto, nil
}

func (s *messageService) GetMessagesByUserId(id int) (dto.Messages, e.ApiError) {
	var messages model.Messages = cliente.GetMessagesByUserId(id)
	var messagesDto dto.Messages

	for _, message := range messages {
		var messageDto dto.Message
		messageDto.Id = message.Id
		messageDto.Content = message.Content
		messageDto.ReceiverId = message.ReceiverId
		messagesDto = append(messagesDto, messageDto)
	}
	return messagesDto, nil
}

func (s *messageService) GetMessagesByItemId(id string) (dto.Messages, e.ApiError) {
	var messages model.Messages = cliente.GetMessagesByItemId(id)
	var messagesDto dto.Messages

	for _, message := range messages {
		var messageDto dto.Message
		messageDto.Id = message.Id
		messageDto.Content = message.Content
		messageDto.ReceiverId = message.ReceiverId
		messagesDto = append(messagesDto, messageDto)
	}
	return messagesDto, nil
}

func (s *userService) NewUser(userDto dto.User) (dto.User, e.ApiError) {
	var user model.User

	user.Username = userDto.Username
	user.Email = userDto.Email
	user.Password = funciones.SSHA256(userDto.Password)
	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName

	var err e.ApiError
	user, err = cliente.NewUser(user)

	userDto.Password = user.Password
	userDto.Id = user.Id

	return userDto, err

}

func (s *messageService) PostMessage(messageDto dto.Message) (dto.Message, e.ApiError) {
	var message model.Message

	message.Content = messageDto.Content
	message.ItemId = messageDto.ItemId
	message.ReceiverId = messageDto.ReceiverId

	var err e.ApiError
	message, err = cliente.NewMessage(message)

	messageDto.id = message.Id

	return messageDto, err
}

func (s *userService) DeleteUser(id int) e.ApiError {
	err := cliente.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *messageService) DeleteMessageById(id int) e.ApiError {
	err := cliente.DeleteMessageById(id)

	if err != nil {
		return err
	}

	return nil
}
