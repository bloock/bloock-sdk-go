package service

import "github.com/enchainte/enchainte-sdk-go/internal/record/entity"

type RecorderService interface {
	SendRecords(records []entity.RecordEntity) ([]entity.RecordReceipt, error)
	GetRecords(records []entity.RecordEntity) ([]entity.RecordReceipt, error)
}
