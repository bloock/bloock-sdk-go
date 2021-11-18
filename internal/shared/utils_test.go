package shared

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytesToHex(t *testing.T) {
	array := []byte("Bloock")

	t.Run("Given an array should return an hexadecimal", func(t *testing.T) {
		hex := BytesToHex(array)
		assert.Equal(t, "426c6f6f636b", hex, "The two results should be the same")
	})
}

func TestHexToBytes(t *testing.T) {
	data := "c7afe76d6dabae68c10c32e5673ed20535ebb00436e615eccc208f14c0993744"

	t.Run("Given an hexadecimal string should return an array of byte", func(t *testing.T){
		bytes, err := HexToBytes(data)
		assert.Nil(t, err)
		assert.IsType(t, []byte{}, bytes, "Type should be []byte")
	})
}

func TestStringToBytes(t *testing.T) {
	data := "Bloock"

	t.Run("Given a string should return an array of byte", func(t *testing.T) {
		arr := StringToBytes(data)
		assert.NotNil(t, arr)
		assert.IsType(t, []byte{}, arr, "The two types should be []byte")
	})
}

func TestHexToBytes32(t *testing.T) {
	data := "c7afe76d6dabae68c10c32e5673ed20535ebb00436e615eccc208f14c0993744"

	t.Run("Given an hexadecimal string should return an array of 32 byte", func(t *testing.T) {
		bytes32, err := HexToBytes32(data)
		assert.Nil(t, err)
		assert.IsType(t, [32]byte{}, bytes32, "The two types should be [32]byte")
	})
}

func TestIsHex(t *testing.T) {
	hex := "abcdefg"

	t.Run("Given an invalid regexp should return false", func(t *testing.T) {
		exp := IsHex(hex)
		assert.False(t, exp)
	})
}

func TestStringify(t *testing.T) {
	i := 123456789
	s := "test"
	arr := []int{1,2,3,4}
	b := true

	t.Run("Given an int value, should convert to string", func(t *testing.T) {
		res := Stringify(i)
		assert.Equal(t, "123456789", res)
	})

	t.Run("Given a string value, should convert to string", func(t *testing.T) {
		res := Stringify(s)
		assert.Equal(t, "test", res)
	})

	t.Run("Given array value, should convert to string", func(t *testing.T) {
		res := Stringify(arr)
		assert.Equal(t, "[1 2 3 4]", res)
	})

	t.Run("Given boolean value, should convert to string", func(t *testing.T) {
		res := Stringify(b)
		assert.Equal(t, "true", res)
	})
}
