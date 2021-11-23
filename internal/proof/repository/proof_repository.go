package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
	entity2 "github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/proof/entity/dto"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"log"
	"math"
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
	url := fmt.Sprintf("%s/core/proof", p.configService.GetApiBaseUrl())
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
	type Stack struct {
		Depth int
		Hash  []byte
	}

	leaves, hashes, bitmap, depth, err := initializeVariables(proof)
	if err != nil {
		return entity.RecordEntity{}, err
	}

	itHashes := 0
	itLeaves := 0
	stack := make([]Stack, 0)
	log.Printf("Actual hashes: %+v leaves %+v\n Depth %+v bitmap %+v", len(hashes), len(leaves), len(depth), bitmap)

	for len(hashes) > itHashes || len(leaves) > itLeaves {
		actDepth := int(depth[itHashes + itLeaves])

		log.Printf("actDepth %+v Depth position %+v", actDepth, itHashes+itLeaves)
		var actHash []byte

		log.Printf("evaluate bitmap %+v", bitmap[int(math.Floor(float64((itHashes+itLeaves)/8)))] & (1 << (7 - ((itHashes + itLeaves) % 8))))
		if (bitmap[(itHashes+itLeaves)/8] & (1 << (7 - ((itHashes + itLeaves) % 8)))) > 0 {
			actHash = hashes[itHashes]
			log.Printf("enter if actualHash %+v", actHash)
			itHashes += 1
		} else {
			actHash = leaves[itLeaves]
			log.Printf("enter else actualHash %+v", actHash)
			itLeaves += 1
		}
		log.Println(len(stack))
		for len(stack) > 0 && stack[len(stack)-1].Depth == actDepth {
			log.Printf("stck Depth %+v Hash %+v",stack[len(stack)-1].Depth,stack[len(stack)-1].Hash)
			lastHash := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if lastHash.Hash == nil {
				return entity.RecordEntity{}, errors.New("verify: Stack got empty before capturing its value")
			}
			actHash, err = entity.Merge(lastHash.Hash, actHash)
			if err != nil {
				return entity.RecordEntity{}, err
			}
			actDepth -= 1
		}
		log.Printf("push to stack actdepth: %+v actHash: %+v",actDepth, actHash)
		stack = append(stack, Stack{actDepth, actHash})
		fmt.Printf("stack %+v\n", stack)
	}

	result := entity.FromHash(shared.BytesToHex(stack[0].Hash))
	return result, nil
}

func(p ProofRepository) ValidateRoot(network string, record entity.RecordEntity) (int, error) {
	log.Printf("Root: %+v", record.GetHash())
	r, err := p.blockchainClient.ValidateRoot(network, record.GetHash())
	return int(r), err
}

func initializeVariables(proof entity2.Proof) (l, h [][]byte, b []byte, d []uint16, err error) {
	leaves := make([][]byte, 0)
	for _, p := range proof.Leaves {
		b := entity.FromHash(p)
		r, err := b.GetByteArray()
		if err != nil {
			return nil, nil, nil, nil, err
		}
		leaves = append(leaves, r)
	}
	hashes := make([][]byte, 0)
	for _, n := range proof.Nodes {
		h, err := shared.HexToBytes(n)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		hashes = append(hashes, h)
	}
	depth, err := shared.HexToBytes16(proof.Depth)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	bitmap, err := shared.HexToBytes(proof.Bitmap)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return leaves, hashes, bitmap, depth, nil
}
