package repository

import (
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity/dto"
)

type RecordRepository struct {
	httpClient HttpClient
	configService service.ConfigService
}

func NewRecordRepository(httpClient HttpClient, configService service.ConfigService) RecordRepository {
	return RecordRepository{
		httpClient: httpClient,
		configService: configService,
	}
}

func(m RecordRepository) SendRecords(records []entity.RecordEntity) {
	url := fmt.Sprintf("%s/core/messages", m.configService.GetApiBaseUrl())
	recordArray := mapHashToStringArray(records)
	body := dto.NewRecordWriteRequest(recordArray)
	m.httpClient.Post(url, body, nil)

}



func(m RecordRepository) FetchRecords(records []entity.RecordEntity) {

}

func mapHashToStringArray(records []entity.RecordEntity) []string {
	recordArray := make([]string, 0)
	for _, m := range records {
		recordArray = append(recordArray, m.GetHash())
	}
	return recordArray
}
