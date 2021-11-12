package service

import (
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/repository"
)

type RecordService struct {
	recordRepository repository.RecordRepository
}

func NewRecordService(recordRepo repository.RecordRepository) RecordService {
	return RecordService{
		recordRepository: recordRepo,
	}
}

func(m RecordService) SendRecords(records []entity.RecordEntity) {

}

func(m RecordService) GetRecords(records []entity.RecordEntity) {

}
