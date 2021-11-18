package repository

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
	entity2 "github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/proof/entity/dto"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
)

type ProofRepository struct {
	httpClient infrastructure.HttpClient
	blockchainClient infrastructure.BlockchainClient
	configService service.ConfigurerService
}

func NewProofRepository(h infrastructure.HttpClient, b infrastructure.BlockchainClient, cs service.ConfigurerService) ProofRepository {
	return ProofRepository{
		httpClient: h,
		blockchainClient: b,
		configService: cs,
	}
}

func(p ProofRepository) RetrieveProof(records []entity.RecordEntity) (entity2.Proof, error) {
	url := fmt.Sprintf("%s/core/proof", "https://api.bloock.dev")
	recordArray := entity.MapHashToStringArray(records)
	body := dto.NewProofRetrieveRequest(recordArray)
	resp, err := p.httpClient.Post(url, body, nil)
	if err != nil {
		return entity2.Proof{}, err
	}

	var proof entity2.Proof
	if err := json.Unmarshal(resp, &proof); err != nil {
		return entity2.Proof{}, err
	}

	return proof, nil
}

func(p ProofRepository) VerifyProof(proof entity2.Proof) (entity.RecordEntity, error) {
	return entity.RecordEntity{}, nil
}

func(p ProofRepository) ValidateRoot(network string, record entity.RecordEntity) (int, error) {
	r, err := p.blockchainClient.ValidateRoot(network, record.GetHash())
	return int(r), err
}
