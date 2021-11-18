package repository

import (
	entity2 "github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
)

type ProoferRepository interface {
	RetrieveProof(records []entity.RecordEntity) (entity2.Proof, error)
	VerifyProof(proof entity2.Proof) (entity.RecordEntity, error)
	ValidateRoot(network string, record entity.RecordEntity) (int, error)
}
