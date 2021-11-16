package repository

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity/dto"
)

type RecordRepository struct {
	httpClient infrastructure.HttpClient
	configService service.ConfigurerService
}

func NewRecordRepository(httpClient infrastructure.HttpClient, configService service.ConfigurerService) RecordRepository {
	return RecordRepository{
		httpClient: httpClient,
		configService: configService,
	}
}

func(m RecordRepository) SendRecords(records []entity.RecordEntity) (dto.RecordWriteResponse, error) {
	url := fmt.Sprintf("%s/core/messages", "https://api.bloock.dev")
	recordArray := entity.MapHashToStringArray(records)
	body := dto.NewRecordWriteRequest(recordArray)
	resp, err := m.httpClient.Post(url, body, nil)
	if err != nil {
		return dto.RecordWriteResponse{}, err
	}

	var recWriteResp dto.RecordWriteResponse
	respToBytes, err := json.Marshal(resp)
	if err != nil {
		return dto.RecordWriteResponse{}, err
	}

	if err := json.Unmarshal(respToBytes, &recWriteResp); err != nil {
		return dto.RecordWriteResponse{}, err
	}

	return recWriteResp, nil
}

func(m RecordRepository) FetchRecords(records []entity.RecordEntity) ([]dto.RecordRetrieveResponse, error) {
	url := fmt.Sprintf("%s/core/messages/fetch", "https://api.bloock.dev")
	recordArray := entity.MapHashToStringArray(records)
	body := dto.NewRecordRetrieveRequest(recordArray)
	resp, err := m.httpClient.Post(url, body, nil)
	if err != nil {
		return []dto.RecordRetrieveResponse{}, err
	}

	var recRetrieveResp []dto.RecordRetrieveResponse
	respToBytes, err := json.Marshal(resp)
	if err != nil {
		return []dto.RecordRetrieveResponse{}, err
	}

	if err := json.Unmarshal(respToBytes, &recRetrieveResp); err != nil {
		return []dto.RecordRetrieveResponse{}, err
	}

	return recRetrieveResp, err
}


