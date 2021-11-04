package infrastructure

import "github.com/enchainte/enchainte-sdk-go/config/entity"

type BlockchainClient interface {
	ValidateRoot(network entity.Network, root string) int
}


