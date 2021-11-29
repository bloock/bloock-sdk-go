package service

import (
	"errors"
	configEntity "github.com/bloock/bloock-sdk-go/internal/config/entity"
	proofEntity "github.com/bloock/bloock-sdk-go/internal/proof/entity"
	"github.com/bloock/bloock-sdk-go/internal/proof/repository"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
	entityException "github.com/bloock/bloock-sdk-go/internal/record/entity/exception"
	"log"
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

	proof, err := p.proofRepository.RetrieveProof(sorted)
	if err != nil {
		return proofEntity.Proof{}, err
	}

	if proof.Depth == "" && proof.Nodes == nil {
		return proofEntity.Proof{}, errors.New("couldn't get proof for specified records")
	}

	return proof, nil
}

func (p ProofService) VerifyRecords(records []entity.RecordEntity, params configEntity.NetworkParams) (int, error) {
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

	return p.VerifyProof(proof, params)
}

func (p ProofService) VerifyProof(proof proofEntity.Proof, params configEntity.NetworkParams) (int, error) {
	if params.Network == "" {
		params.Network = configEntity.EthereumMainnet
	}

	root, err := p.proofRepository.VerifyProof(proof)
	if err != nil {
		return -1, err
	}
	if root == (entity.RecordEntity{}) {
		return -1, errors.New("the provided proof is invalid")
	}
	log.Println(params.Network)
	return p.proofRepository.ValidateRoot(params.Network, root)
}
