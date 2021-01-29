package message

import "encoding/hex"

type Message struct {
	hash string
}

func New(bytes []byte) (*Message, error) {
	hash := hex.EncodeToString(bytes)
	m := Message{hash: hash}

	return &m, nil
}

func (m *Message) Hash() string {
	return m.hash
}
