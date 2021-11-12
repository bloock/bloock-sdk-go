package service

import "github.com/enchainte/enchainte-sdk-go/internal/message/entity"

type MessageService struct {
	messageRepository MessageRepository
}

func NewMessageService(messageRepo MessageRepository) MessageService {
	return MessageService{
		messageRepository: messageRepo,
	}
}

func(m MessageService) SendRecords(messages []entity.MessageEntity) {

}

func(m MessageService) GetRecords(messages []entity.MessageEntity) {

}
