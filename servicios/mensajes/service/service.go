package service

import (
	cliente "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/client"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/dto"
	e "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/errors"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/model"
)

type messageService struct{}

type messageServiceInterface interface {
	GetMessageById(id int) (dto.Message, e.ApiError)
	GetMessages() (dto.Messages, e.ApiError)
	GetMessagesByUserId(id int) (dto.Messages, e.ApiError)
	GetMessagesByItemId(id string) (dto.Messages, e.ApiError)

	PostMessage(messageDto dto.Message) (dto.Message, e.ApiError)
	PostMessages(messages dto.Messages) (dto.Messages, e.ApiError)

	DeleteMessageById(id int) (dto.Message, e.ApiError)
}

var jwtKey = []byte("secret_key")

var (
	MessageService messageServiceInterface
)

func init() {
	MessageService = &messageService{}
}

func (s *messageService) GetMessageById(id int) (dto.Message, e.ApiError) {
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

func (s *messageService) GetMessages() (dto.Messages, e.ApiError) {
	var messages model.Messages = cliente.GetMessages()
	var messagesDto dto.Messages

	for _, message := range messages {
		var messageDto dto.Message
		messageDto.Id = message.Id
		messageDto.Content = message.Content
		messageDto.ReceiverId = message.ReceiverId
		messageDto.ItemId = message.ItemId
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
		messageDto.ItemId = message.ItemId
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
		messageDto.ItemId = message.ItemId
		messagesDto = append(messagesDto, messageDto)
	}
	return messagesDto, nil
}

func (s *messageService) PostMessage(messageDto dto.Message) (dto.Message, e.ApiError) {
	var message model.Message

	message.Content = messageDto.Content
	message.ItemId = messageDto.ItemId
	message.ReceiverId = messageDto.ReceiverId

	var err e.ApiError
	message, err = cliente.NewMessage(message)

	messageDto.Id = message.Id

	return messageDto, err
}

func (s *messageService) PostMessages(messages dto.Messages) (dto.Messages, e.ApiError) {
	var MessagesDto dto.Messages

	for _, messageDto := range messages {
		var message model.Message

		message.Content = messageDto.Content
		message.ItemId = messageDto.ItemId
		message.ReceiverId = messageDto.ReceiverId

		message, err := cliente.NewMessage(message)

		if err != nil {
			return nil, err
		}

		messageDto.Id = message.Id

		MessagesDto = append(MessagesDto, messageDto)
	}

	return MessagesDto, nil
}

func (s *messageService) DeleteMessageById(id int) (dto.Message, e.ApiError) {
	deletedMessage, err := cliente.DeleteMessage(id)

	if err != nil {
		return dto.Message{}, err
	}

	var message dto.Message
	message.Content = deletedMessage.Content
	message.Id = deletedMessage.Id
	message.ItemId = deletedMessage.ItemId
	message.ReceiverId = deletedMessage.ReceiverId

	return message, nil
}
