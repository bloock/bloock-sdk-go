package repository

import (
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity/dto"
)

//go:generate mockgen -source=internal/record/repository/recorder_repository.go -destination internal/record/mockrecord/mock_record_repository.go -package=mockrecord
type RecorderRepository interface {
	SendRecords(records []entity.RecordEntity) (dto.RecordWriteResponse, error)
	FetchRecords(records []entity.RecordEntity) ([]dto.RecordRetrieveResponse, error)
}
