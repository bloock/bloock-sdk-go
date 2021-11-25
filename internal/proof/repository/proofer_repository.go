package repository

import (
	proofEntity "github.com/bloock/bloock-sdk-go/internal/proof/entity"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
)

type ProoferRepository interface {
	RetrieveProof(records []entity.RecordEntity) (proofEntity.Proof, error)
	VerifyProof(proof proofEntity.Proof) (entity.RecordEntity, error)
	ValidateRoot(network string, record entity.RecordEntity) (int, error)
}
