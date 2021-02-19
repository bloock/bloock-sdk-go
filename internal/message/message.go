package message

import (
	"encoding/hex"
	"github.com/enchainte/enchainte-sdk-go/pkg/crypto"
)

type Message struct {
	Hasher crypto.Hasher
	hash string
}

func New(bytes []byte) (*Message, error) {
	//hash := hex.EncodeToString(bytes)
	m := Message{Hasher: crypto.Blake2b()}
	m.hash = m.HashData(bytes)

	return &m, nil
}

func (m *Message) HashData(key []byte) string {
	return hex.EncodeToString(m.Hasher.Hash(key))
}

func (m *Message) Hash() string {
	return m.hash
}
