package infrastructure

//go:generate mockgen -source=internal/infrastructure/blockchain_client.go -destination internal/infrastructure/blockchain/mockblockchain/mocks_blockchain.go -package=mockblockchain
type BlockchainClient interface {
	ValidateRoot(network string, root string) (int64, error)
}
