package infrastructure

type BlockchainClient interface {
	ValidateRoot(network string, root string) (int64, error)
}


