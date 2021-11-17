package service

import "github.com/enchainte/enchainte-sdk-go/internal/proof/repository"

type ProofService struct {
	proofRepository repository.ProoferRepository
}

func NewProofService(pr repository.ProoferRepository) ProofService {
	return ProofService{
		proofRepository: pr,
	}
}

func(p ProofService) RetrieveProof() {

}

func(p ProofService) VerifyRecords() {

}

func(p ProofService) VerifyProof() {

}
