package service

import (
	entity2 "github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
)

type ProoferService interface {
	RetrieveProof(records []entity.RecordEntity) (entity2.Proof, error)
	VerifyRecords(records []entity.RecordEntity, network string) (int, error)
	VerifyProof(proof entity2.Proof, network string) (int, error)
}
