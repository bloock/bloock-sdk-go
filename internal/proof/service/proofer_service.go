package service

import (
	configEntity "github.com/bloock/bloock-sdk-go/internal/config/entity"
	proofEntity "github.com/bloock/bloock-sdk-go/internal/proof/entity"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
)

type ProoferService interface {
	RetrieveProof(records []entity.RecordEntity) (proofEntity.Proof, error)
	VerifyRecords(records []entity.RecordEntity, params configEntity.NetworkParams) (int, error)
	VerifyProof(proof proofEntity.Proof, params configEntity.NetworkParams) (int, error)
}
