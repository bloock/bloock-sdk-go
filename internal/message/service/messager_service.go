package service

import "github.com/enchainte/enchainte-sdk-go/internal/message/entity"

type MessagerService interface {
	SendRecords(messages []entity.MessageEntity)
	GetRecords(messages []entity.MessageEntity)
}
