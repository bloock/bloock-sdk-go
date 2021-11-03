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

func TestStringToBytes(t *testing.T) {
	data := "Bloock"

	t.Run("Given a string should return an array of byte", func(t *testing.T) {
		arr := StringToBytes(data)
		assert.NotNil(t, arr)
		assert.IsType(t, []byte{}, arr, "The two types should be []byte")
	})
}
