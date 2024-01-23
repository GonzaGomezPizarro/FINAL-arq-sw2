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
	GetMessages() (dto.Messages, e.ApiError)
	GetMessagesByUserId(id int) (dto.Messages, e.ApiError)
	GetMessageById(id int) (dto.Message, e.ApiError)
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

func (s *userService) GetUserById(id int) (dto.User, e.ApiError) {

	var user model.User = cliente.GetUserById(id)
	var userDto dto.User

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}

	userDto.Id = user.Id
	userDto.Username = user.Username
	userDto.Password = user.Password
	userDto.Email = user.Email
	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName

	return userDto, nil
}

func (s *messageService) GetMessageById(id int) (dto.Message, e.apiError) {
	var message model.Message = cliente.GetMessageById(id)
}

func (s *messageService) GetMessages() ([]dto.Message, error) {
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

func (s *userService) DeleteUser(id int) e.ApiError {
	err := cliente.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}
