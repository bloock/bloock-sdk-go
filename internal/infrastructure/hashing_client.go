package infrastructure

type HashingClient interface {
	GenerateHash(data []byte) string
}
