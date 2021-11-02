package hashing

import (
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"github.com/ethereum/go-ethereum/crypto"
)

type Keccak struct {
	HashingClient infrastructure.HashingClient
}

func NewKeccak(client infrastructure.HashingClient) Keccak {
	return Keccak{
		HashingClient: client,
	}
}

func (c Keccak) GenerateHash(data []byte) string {
	hash := crypto.Keccak256(data)
	hex := shared.BytesToHex(hash)
	return hex
}