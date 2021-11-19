package repository

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
	entity2 "github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/proof/entity/dto"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
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
	/*leaves, hashes, depth, bitmap, err := initializeVariables(proof)
	if err != nil {
		return entity.RecordEntity{}, err
	}

	itHashes := 0
	itLeaves := 0

	for len(hashes) > itHashes || len(leaves) > itLeaves {
		actDepth := depth[itHashes + itLeaves]
		var actHash []byte

		if (bitmap[int(math.Floor(float64((itHashes+itLeaves)/8)))] & 1 << (7 - ((itHashes + itLeaves) % 8))) > 0 {
			actHash = hashes[itHashes]
			itHashes += 1
		} else {
			actHash = leaves[itLeaves]
			itLeaves += 1
		}
	}*/





	return entity.RecordEntity{}, nil
}

func initializeVariables(proof entity2.Proof) (l, h, d, b []byte, err error) {
	var leaves []byte
	for _, p := range proof.Leaves {
		b := entity.FromHash(p)
		r, err := b.GetByteArray()
		if err != nil {
			return nil, nil, nil, nil, err
		}
		leaves = r
	}
	var hashes []byte
	for _, n := range proof.Nodes {
		h, err := shared.HexToBytes(n)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		hashes = h
	}
	depth, err := shared.HexToBytes(proof.Depth)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	bitmap, err := shared.HexToBytes(proof.Bitmap)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return leaves, hashes, depth, bitmap, nil
}

func(p ProofRepository) ValidateRoot(network string, record entity.RecordEntity) (int, error) {
	r, err := p.blockchainClient.ValidateRoot(network, record.GetHash())
	return int(r), err
}
