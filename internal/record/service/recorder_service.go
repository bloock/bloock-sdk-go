package service

import "github.com/enchainte/enchainte-sdk-go/internal/record/entity"

type RecorderService interface {
	SendRecords(records []entity.RecordEntity)
	GetRecords(records []entity.RecordEntity)
}
