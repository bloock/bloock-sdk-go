package entity

import (
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
)

type RecordEntity struct {
	hashAlgorithm infrastructure.HashingClient
	hash string
}

func NewRecordEntity(hash string) RecordEntity {
	return RecordEntity{
		hash: hash,
	}
}

func(m RecordEntity) fromObject(data interface{}) RecordEntity {
	return m.fromString(shared.Stringify(data))
}

func(m RecordEntity) fromHash(hash string) RecordEntity {
	return NewRecordEntity(hash)
}

func(m RecordEntity) fromHex(hex string) (RecordEntity, error) {
	dataArray, err := shared.HexToBytes(hex)
	if err != nil {
		return RecordEntity{}, err
	}
	return NewRecordEntity(m.hashAlgorithm.GenerateHash(dataArray)), nil
}

func(m RecordEntity) fromString(string string) RecordEntity {
	dataArray := shared.StringToBytes(string)
	return NewRecordEntity(m.hashAlgorithm.GenerateHash(dataArray))
}

func(m RecordEntity) fromUint8Array(array []byte) RecordEntity {
	return NewRecordEntity(m.hashAlgorithm.GenerateHash(array))
}

func(m RecordEntity) IsValid(record RecordEntity) bool {
	if isType(record) {
		record := m.GetHash()
		if len(record) == 64 && shared.IsHex(record) {
			return true
		}
	}
	return false
}

func(m RecordEntity) GetHash() string {
	return m.hash
}

func isType(t interface{}) bool {
	switch t.(type) {
	case RecordEntity:
		return true
	default:
		return false
	}
}

func MapHashToStringArray(records []RecordEntity) []string {
	recordArray := make([]string, 0)
	for _, m := range records {
		recordArray = append(recordArray, m.GetHash())
	}
	return recordArray
}