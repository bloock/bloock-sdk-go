package repository

import (
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/message/entity"
)

type MessageRepository struct {
	httpClient HttpClient
	configService service.ConfigService
}

func NewMessageRepository(httpClient HttpClient, configService service.ConfigService) MessageRepository {
	return MessageRepository{
		httpClient: httpClient,
		configService: configService,
	}
}

func(m MessageRepository) SendRecords(messages []entity.MessageEntity) {
	url := fmt.Sprintf("%s/core/messages", m.configService.GetApiBaseUrl())

}

func(m MessageRepository) FetchRecords(messages []entity.MessageEntity) {

}
