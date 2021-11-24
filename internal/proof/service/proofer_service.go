package service

import (
	proofEntity "github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
)

type ProoferService interface {
	RetrieveProof(records []entity.RecordEntity) (proofEntity.Proof, error)
	VerifyRecords(records []entity.RecordEntity, network string) (int, error)
	VerifyProof(proof proofEntity.Proof, network string) (int, error)
}
