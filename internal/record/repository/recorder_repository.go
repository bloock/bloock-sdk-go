package repository

import (
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity/dto"
)

type RecorderRepository interface {
	SendRecords(records []entity.RecordEntity) (dto.RecordWriteResponse, error)
	FetchRecords(records []entity.RecordEntity) ([]dto.RecordRetrieveResponse, error)
}
