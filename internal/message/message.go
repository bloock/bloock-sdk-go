package message

import "encoding/hex"

type message struct {
	hash string
}

func New(bytes []byte) (*message, error) {
	hash := hex.EncodeToString(bytes)
	m := message{hash: hash}

	return &m, nil
}

func (m *message) Hash() string {
	return m.hash
}
