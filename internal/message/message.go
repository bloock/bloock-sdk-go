package message

import (
	"encoding/hex"
	"github.com/enchainte/enchainte-sdk-go/pkg/crypto"
)

type Message struct {
	Hasher crypto.Hasher
	hash string
}

// New creates a new message instance. It takes a content data in bytes format and returns the message instance
// containing the hash of the content data represented in hexadecimal format.
func New(content []byte) (*Message, error) {
	m := Message{Hasher: crypto.Blake2b(), hash: string(content)}

	return &m, nil
}

// HashData hashes the provided key and returns its hexadecimal representation
func (m *Message) HashData() {
	m.hash = hex.EncodeToString(m.Hasher.Hash([]byte(m.hash)))
}

// Hash is a getter of the message hash property
func (m *Message) Hash() string {
	if !isHex(m.hash) || len(m.hash) != 64 {
		m.HashData()
	}
	return m.hash
}

func isHex(s string) bool {
	if _, err := hex.DecodeString(s); err != nil {
		return false
	}
	return true
}