package service

import (
	"errors"
	proofEntity "github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/proof/repository"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	entityException "github.com/enchainte/enchainte-sdk-go/internal/record/entity/exception"
)

type ProofService struct {
	proofRepository repository.ProoferRepository
}

func NewProofService(pr repository.ProoferRepository) ProofService {
	return ProofService{
		proofRepository: pr,
	}
}

func (p ProofService) RetrieveProof(records []entity.RecordEntity) (proofEntity.Proof, error) {
	for _, r := range records {
		if !r.IsValid(r) {
			return proofEntity.Proof{}, entityException.NewInvalidRecordException()
		}
	}

	sorted := entity.Sort(records)

	return p.proofRepository.RetrieveProof(sorted)
}

func (p ProofService) VerifyRecords(records []entity.RecordEntity, network string) (int, error) {
	for _, r := range records {
		if !r.IsValid(r) {
			return -1, entityException.NewInvalidRecordException()
		}
	}

	proof, err := p.RetrieveProof(records)
	if err != nil {
		return -1, err
	}
	if proof.Depth == "" && proof.Nodes == nil {
		return -1, errors.New("couldn't get proof for specified records")
	}

	return p.VerifyProof(proof, network)
}

func (p ProofService) VerifyProof(proof proofEntity.Proof, network string) (int, error) {
	root, err := p.proofRepository.VerifyProof(proof)
	if err != nil {
		return -1, err
	}
	if root == (entity.RecordEntity{}) {
		return -1, errors.New("the provided proof is invalid")
	}

	return p.proofRepository.ValidateRoot(network, root)
}
