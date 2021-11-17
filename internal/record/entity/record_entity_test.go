package entity

import (
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordEntity(t *testing.T) {
	var record RecordEntity

	t.Run("Given a valid hash should return that hash", func(t *testing.T) {
		r := record.FromHash("test_hash")

		assert.Equal(t, "test_hash", r.GetHash())
	})

	t.Run("Given a valid hex should return a hashed Keccak256", func(t *testing.T) {
		s := "10101010101010101010101010101010101010101010101010101010101010101111111111111111111111111111111111111111111111111111111111111111"
		p := "e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994a5"
		r, err := record.FromHex(s)

		assert.Nil(t, err)
		assert.Equal(t, p, r.GetHash())
	})

	t.Run("Given a valid string should return a hashed Keccak256", func(t *testing.T) {
		s := "testing keccak"
		p := "7e5e383e8e70e55cdccfccf40dfc5d4bed935613dffc806b16b4675b555be139"
		r := record.FromString(s)

		assert.Equal(t, p, r.GetHash())
	})

	t.Run("Given a valid []byte should return a hashed Keccak256", func(t *testing.T) {
		s := "testing keccak"
		b := shared.StringToBytes(s)
		p := "7e5e383e8e70e55cdccfccf40dfc5d4bed935613dffc806b16b4675b555be139"
		r := record.FromUint8Array(b)

		assert.Equal(t, p, r.GetHash())
	})

	t.Run("Given a valid hash, should be valid", func(t *testing.T) {
		r := record.FromHash("1010101010101010101010101010101010101010101010101010101010101010")

		assert.True(t, r.IsValid(r))
	})

	t.Run("Given a hash with missing char, should be invalid", func(t *testing.T) {
		r := record.FromHash("010101010101010101010101010101010101010101010101010101010101010")

		assert.False(t, r.IsValid(r))
	})

	t.Run("Given a hash with wrong char, should be invalid", func(t *testing.T) {
		r := record.FromHash("G010101010101010101010101010101010101010101010101010101010101010")

		assert.False(t, r.IsValid(r))
	})

	t.Run("Given four records, should sort with success", func(t *testing.T) {
		rec := make([]RecordEntity, 0)
		rec = append(rec, NewRecordEntity("6a83f545cb5693a32b5d56fb4a0530f7054df0c7e2e6b0a9fef36e26a2a96b04"))
		rec = append(rec, NewRecordEntity("2d9130eb0900a08f22dee5e0330672861e6035eb858e1d1ac0d0d5d98a676800"))
		rec = append(rec, NewRecordEntity("cadc5a160b48bde5727b08e1f8d1b8fe08704ff3cc730bf4919a2ef10ae6e291"))
		rec = append(rec, NewRecordEntity("db6d0af6e743ca02954f1feb7dec3033fe4f86d429b8dd5b7dd654b794d71dee"))

		result := record.Sort(rec)
		assert.Equal(t, "2d9130eb0900a08f22dee5e0330672861e6035eb858e1d1ac0d0d5d98a676800", result[0].GetHash())
		assert.Equal(t, "6a83f545cb5693a32b5d56fb4a0530f7054df0c7e2e6b0a9fef36e26a2a96b04", result[1].GetHash())
		assert.Equal(t, "cadc5a160b48bde5727b08e1f8d1b8fe08704ff3cc730bf4919a2ef10ae6e291", result[2].GetHash())
		assert.Equal(t, "db6d0af6e743ca02954f1feb7dec3033fe4f86d429b8dd5b7dd654b794d71dee", result[3].GetHash())
	})
}
