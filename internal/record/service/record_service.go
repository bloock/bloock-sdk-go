package service

import (
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity/exception"
	"github.com/enchainte/enchainte-sdk-go/internal/record/repository"
	exception2 "github.com/enchainte/enchainte-sdk-go/internal/shared/entity/exception"
	"reflect"
)

type RecordService struct {
	recordRepository repository.RecordRepository
}

func NewRecordService(recordRepo repository.RecordRepository) RecordService {
	return RecordService{
		recordRepository: recordRepo,
	}
}

func(m RecordService) SendRecords(records []entity.RecordEntity) ([]entity.RecordReceipt, error) {
	if len(records) == 0 {
		return []entity.RecordReceipt{}, nil
	}

	typeOfRecords := reflect.TypeOf(records).Kind()
	if typeOfRecords != reflect.Slice {
		return []entity.RecordReceipt{}, exception2.NewInvalidArgumentException()
	}

	for _, r := range records {
		if !r.IsValid(r) {
			return []entity.RecordReceipt{}, exception.NewInvalidRecordException()
		}
	}

	recds, err := m.recordRepository.SendRecords(records)
	if err != nil {
		return []entity.RecordReceipt{}, err
	}

	var result []entity.RecordReceipt
	for _, r := range records {
		rr := entity.NewRecordReceipt(recds.Anchor, recds.Client, r.GetHash(), recds.Status)
		result = append(result, rr)
	}
	return result, nil
}

func(m RecordService) GetRecords(records []entity.RecordEntity) ([]entity.RecordReceipt, error) {
	if len(records) == 0 {
		return []entity.RecordReceipt{}, nil
	}

	typeOfRecords := reflect.TypeOf(records).Kind()
	if typeOfRecords != reflect.Slice {
		return []entity.RecordReceipt{}, nil //Retornar error tipus InvalidArgumentException
	}

	for _, r := range records {
		if !r.IsValid(r) {
			return []entity.RecordReceipt{}, nil //Retornar error tipus InvalidArgumentException
		}
	}

	recds, err := m.recordRepository.FetchRecords(records)
	if err != nil {
		return []entity.RecordReceipt{}, err
	}
	if len(recds) == 0 {
		return []entity.RecordReceipt{}, nil
	}

	var result []entity.RecordReceipt
	for _, r := range recds {
		rr := entity.NewRecordReceipt(r.Anchor, r.Client, r.Message, r.Status)
		result = append(result, rr)
	}

	return result, nil
}
