package crypto

import (
	b2b "golang.org/x/crypto/blake2b"
)

type Hasher interface {
	Hash(key []byte) ([]byte, error)
	Validate(hash string) bool
}

type blake2b struct {
	size int
}

func Blake2b(size int) Hasher {
	return &blake2b{256}
}

func (b *blake2b) Hash(key []byte) ([]byte, error) {
	hash, err := b2b.New(b.size, key)
	if err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

func (b *blake2b) Validate(hash string) bool {
	panic("implement me")
}
