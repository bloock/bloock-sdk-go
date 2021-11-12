package repository

import "github.com/enchainte/enchainte-sdk-go/internal/message/entity"

type MessagerRepository interface {
	SendRecords(messages []entity.MessageEntity)
	GetRecords(messages []entity.MessageEntity)
}
