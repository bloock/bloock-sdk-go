package crypto

import (
	b2b "golang.org/x/crypto/blake2b"
)

type Hasher interface {
	Hash(key []byte) []byte
}

type blake2b struct {
}

func Blake2b() Hasher {
	return &blake2b{}
}

func (b *blake2b) Hash(key []byte) []byte {
	hash := b2b.Sum256(key)

	return hash[:]

	//hash, err := b2b.New256(key)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return hash.Sum(nil), nil
}
