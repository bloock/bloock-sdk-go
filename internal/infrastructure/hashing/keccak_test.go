package hashing

import (
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	var data1, data2, data3 = "Test Data", "Bloock", "testing keccak"
	var dataArray = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	var client infrastructure.HashingClient
	var hashingAlgorithm = NewKeccak(client)

	t.Run("Given data1 should return the correct hexadecimal hash", func(t *testing.T) {
		result := hashingAlgorithm.GenerateHash(shared.StringToBytes(data1))
		assert.Equal(t, "e287462a142cd6237de5a89891ad82065f6aca6644c161b1a61c933c5d26117a", result)
	})

	t.Run("Given data2 should return the correct hexadecimal hash", func(t *testing.T) {
		result := hashingAlgorithm.GenerateHash(shared.StringToBytes(data2))
		assert.Equal(t, "3a7ae5d1ca472a7459e484babf13adf1aa7fe78326755969e3e2f5fc7766f6ee", result)
	})

	t.Run("Given data3 should return the correct hexadecimal hash", func(t *testing.T) {
		result := hashingAlgorithm.GenerateHash(shared.StringToBytes(data3))
		assert.Equal(t, "7e5e383e8e70e55cdccfccf40dfc5d4bed935613dffc806b16b4675b555be139", result)
	})

	t.Run("Given data3 should return the correct hexadecimal hash", func(t *testing.T) {
		result := hashingAlgorithm.GenerateHash(dataArray)
		assert.Equal(t, "d5f4f7e1d989848480236fb0a5f808d5877abf778364ae50845234dd6c1e80fc", result)
	})


}
