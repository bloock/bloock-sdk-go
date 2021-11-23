package entity

import (
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/hashing"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"sort"
	"strings"
)

var hashAlgorithm hashing.Keccak

type RecordEntity struct {
	hash          string
}

func NewRecordEntity(hash string) RecordEntity {
	return RecordEntity{
		hash: hash,
	}
}

func FromObject(data interface{}) RecordEntity {
	return FromString(shared.Stringify(data))
}

func FromHash(hash string) RecordEntity {
	return NewRecordEntity(hash)
}

func FromHex(hex string) (RecordEntity, error) {
	dataArray, err := shared.HexToBytes(hex)
	if err != nil {
		return RecordEntity{}, err
	}
	return NewRecordEntity(hashAlgorithm.GenerateHash(dataArray)), nil
}

func FromString(string string) RecordEntity {
	dataArray := shared.StringToBytes(string)
	return NewRecordEntity(hashAlgorithm.GenerateHash(dataArray))
}

func FromUint8Array(array []byte) RecordEntity {
	return NewRecordEntity(hashAlgorithm.GenerateHash(array))
}

func Sort(records []RecordEntity) []RecordEntity {
	sort.SliceStable(records, func(i, j int) bool {
		return strings.ToUpper(records[i].GetHash()) < strings.ToUpper(records[j].GetHash())
	})
	return records
}

func (m RecordEntity) IsValid(record RecordEntity) bool {
	if isType(record) {
		record := m.GetHash()
		if len(record) == 64 && shared.IsHex(record) {
			return true
		}
	}

	return false
}

func (m RecordEntity) GetHash() string {
	return m.hash
}

func (m RecordEntity) GetByteArray() ([]byte, error) {
	ret, err := shared.HexToBytes(m.hash)
	if err != nil {
		return nil, err
	}
	return ret, nil
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

func Merge(left, right []byte) ([]byte, error) {
	concat := make([]byte, len(left)+len(right))
	concat = append(left, right...)

	return FromUint8Array(concat).GetByteArray()
}
