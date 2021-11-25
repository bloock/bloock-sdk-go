package entity

import (
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/hashing"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"sort"
	"strings"
)

var hashAlgorithm hashing.Keccak

/*
RecordEntity is the struct in charge of computing and storing the
value of the data sent to Bloock.
This class is intended to be used by calling "from" methods to create instances of RecordEntity.
*/
type RecordEntity struct {
	hash string
}

func NewRecordEntity(hash string) RecordEntity {
	return RecordEntity{
		hash: hash,
	}
}

/*
FromObject
Given an JSON object, returns a Record with its value hashed.
Parameters:
	{interface{}} data
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
*/
func FromObject(data interface{}) RecordEntity {
	return FromString(shared.Stringify(data))
}

/*
FromHash
Given a value already hashed creates a Record containing it.
Parameters:
	{string} Hexadecimal string without prefix and length 64.
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
*/
func FromHash(hash string) RecordEntity {
	return NewRecordEntity(hash)
}

/*
FromHex
Given a hexadecimal string (with no 0x prefix) returns a Record with its value hashed.
Parameters:
	{string} Hexadecimal string without prefix.
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
	{error} any type of error when hashing or converting
*/
func FromHex(hex string) (RecordEntity, error) {
	dataArray, err := shared.HexToBytes(hex)
	if err != nil {
		return RecordEntity{}, err
	}
	return NewRecordEntity(hashAlgorithm.GenerateHash(dataArray)), nil
}

/*
FromString
Given a string returns a Record with its value hashed.
Parameters:
	{string} String object.
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
*/
func FromString(string string) RecordEntity {
	dataArray := shared.StringToBytes(string)
	return NewRecordEntity(hashAlgorithm.GenerateHash(dataArray))
}

/*
FromUint8Array
Given a bytes object returns a Record with its value hashed.
Parameters:
	{[]byte} Bytes object.
Returns:
	{RecordEntity} RecordEntity object of the hashed input.
*/
func FromUint8Array(array []byte) RecordEntity {
	return NewRecordEntity(hashAlgorithm.GenerateHash(array))
}

func Sort(records []RecordEntity) []RecordEntity {
	sort.SliceStable(records, func(i, j int) bool {
		return strings.ToUpper(records[i].GetHash()) < strings.ToUpper(records[j].GetHash())
	})
	return records
}

/*
IsValid
Given a RecordEntity returns True if its contents are valid to be sent to Bloock's API or False otherwise.
Parameters:
	{RecordEntity} RecordEntity object.
Returns:
	{boolean} Boolean indicating if the RecordEntity is susceptible to be sent (True) or not (False).
*/
func (m RecordEntity) IsValid(record RecordEntity) bool {
	if isType(record) {
		record := m.GetHash()
		if len(record) == 64 && shared.IsHex(record) {
			return true
		}
	}

	return false
}

/*
GetHash
Returns the hashed representation of the Record string.
Returns:
	{string} String containing the RecordEntity hash as a hexadecimal (with no "0x" prefix).
*/
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
