package entity

import (
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
)

type MessageEntity struct {
	hashAlgorithm infrastructure.HashingClient
	hash string
}

func NewMessageEntity(hash string) MessageEntity{
	return MessageEntity{
		hash: hash,
	}
}

func(m MessageEntity) fromObject(data interface{}) MessageEntity {
	return m.fromString(shared.Stringify(data))
}

func(m MessageEntity) fromHash(hash string) MessageEntity {
	return NewMessageEntity(hash)
}

func(m MessageEntity) fromHex(hex string) (MessageEntity, error) {
	dataArray, err := shared.HexToBytes(hex)
	if err != nil {
		return MessageEntity{}, err
	}
	return NewMessageEntity(m.hashAlgorithm.GenerateHash(dataArray)), nil
}

func(m MessageEntity) fromString(string string) MessageEntity {
	dataArray := shared.StringToBytes(string)
	return NewMessageEntity(m.hashAlgorithm.GenerateHash(dataArray))
}

func(m MessageEntity) fromUint8Array(array []byte) MessageEntity {
	return NewMessageEntity(m.hashAlgorithm.GenerateHash(array))
}

func(m MessageEntity) isValid(message MessageEntity) bool {
	if isType(message) {
		record := m.GetHash()
		if len(record) == 64 && shared.IsHex(record) {
			return true
		}
	}
	return false
}

func(m MessageEntity) GetHash() string {
	return m.hash
}

func isType(t interface{}) bool {
	switch t.(type) {
	case MessageEntity:
		return true
	default:
		return false
	}
}