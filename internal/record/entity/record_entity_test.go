package entity

import (
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordEntity(t *testing.T) {
	var record RecordEntity

	t.Run("Given a valid hash should return that hash", func(t *testing.T) {
		r := record.fromHash("test_hash")

		assert.Equal(t, "test_hash", r.GetHash())
	})

	t.Run("Given a valid hex should return a hashed Keccak256", func(t *testing.T) {
		s := "10101010101010101010101010101010101010101010101010101010101010101111111111111111111111111111111111111111111111111111111111111111"
		p := "e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994a5"
		r, err := record.fromHex(s)

		assert.Nil(t, err)
		assert.Equal(t, p, r.GetHash())
	})

	t.Run("Given a valid string should return a hashed Keccak256", func(t *testing.T) {
		s := "testing keccak"
		p := "7e5e383e8e70e55cdccfccf40dfc5d4bed935613dffc806b16b4675b555be139"
		r := record.fromString(s)

		assert.Equal(t, p, r.GetHash())
	})

	t.Run("Given a valid []byte should return a hashed Keccak256", func(t *testing.T) {
		s := "testing keccak"
		b := shared.StringToBytes(s)
		p := "7e5e383e8e70e55cdccfccf40dfc5d4bed935613dffc806b16b4675b555be139"
		r := record.fromUint8Array(b)

		assert.Equal(t, p, r.GetHash())
	})

	t.Run("Given a valid hash, should be valid", func(t *testing.T) {
		r := record.fromHash("1010101010101010101010101010101010101010101010101010101010101010")

		assert.True(t, r.IsValid(r))
	})

	t.Run("Given a hash with missing char, should be invalid", func(t *testing.T) {
		r := record.fromHash("010101010101010101010101010101010101010101010101010101010101010")

		assert.False(t, r.IsValid(r))
	})

	t.Run("Given a hash with wrong char, should be invalid", func(t *testing.T) {
		r := record.fromHash("G010101010101010101010101010101010101010101010101010101010101010")

		assert.False(t, r.IsValid(r))
	})
}
