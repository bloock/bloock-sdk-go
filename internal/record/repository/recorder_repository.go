package repository

import "github.com/enchainte/enchainte-sdk-go/internal/record/entity"

type RecorderRepository interface {
	SendRecords(records []entity.RecordEntity)
	GetRecords(records []entity.RecordEntity)
}
