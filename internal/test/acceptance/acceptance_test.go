package acceptance

import (
	"github.com/bloock/bloock-sdk-go/internal"
	exceptionEntity "github.com/bloock/bloock-sdk-go/internal/anchor/entity/exception"
	configEntity "github.com/bloock/bloock-sdk-go/internal/config/entity"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
	"github.com/bloock/bloock-sdk-go/internal/record/entity/exception"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func GetSdk() internal.BloockClient {
	apiKey := os.Getenv("API_KEY")
	apiHost := os.Getenv("API_HOST")
	client := internal.NewBloockClient(apiKey)
	client.SetApiHost(apiHost)
	return client
}

func TestAcceptance(t *testing.T) {
	sdk := GetSdk()

	t.Run("Basic test E2E", func(t *testing.T) {
		record := entity.FromString(randHex(64))
		records := []entity.RecordEntity{record}

		rr, err := sdk.SendRecords(records)
		assert.Nil(t, err)
		assert.NotEqual(t, entity.RecordReceipt{}, rr[0])

		sdk.WaitAnchor(rr[0].Anchor, 120000)

		// Retrieving record proof
		proof, err := sdk.GetProof(records)
		assert.Nil(t, err)
		timestamp, err := sdk.VerifyProof(proof, configEntity.BloockChain)
		assert.Greater(t, timestamp, 5000)
	})

	t.Run("Test send records invalid record input wrong char", func(t *testing.T) {
		record := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aG")
		records := []entity.RecordEntity{record}

		_, err := sdk.SendRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test send records invalid record input missing chars", func(t *testing.T) {
		record1 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.SendRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test send records invalid record input wrong start", func(t *testing.T) {
		record1 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994bb")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.SendRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test send records empty record input", func(t *testing.T) {
		res, err := sdk.SendRecords([]entity.RecordEntity{})
		assert.Nil(t, err)
		assert.Equal(t, []entity.RecordReceipt{}, res)
	})

	t.Run("Test get records invalid record input wrong char", func(t *testing.T) {
		record := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aG")

		records := []entity.RecordEntity{record}

		_, err := sdk.GetRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get records invalid record input missing chars", func(t *testing.T) {
		record1 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.GetRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get records invalid record input wrong start", func(t *testing.T) {
		record1 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994bb")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.GetRecords(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get anchor non existing anchor", func(t *testing.T) {
		_, err := sdk.GetAnchor(666666666666666666)
		assert.NotNil(t, err)
		assert.IsType(t, exceptionEntity.NewAnchorNotFoundException(), err)
		assert.Equal(t, "Anchor not found", err.Error())
	})

	t.Run("Test wait anchor non existing anchor", func(t *testing.T) {
		_, err := sdk.WaitAnchor(666666666666666666, 3000)
		assert.NotNil(t, err)
		assert.IsType(t, exceptionEntity.WaitAnchorTimeoutException{}, err)
		assert.Equal(t, "Timeout exceeded while waiting for anchor", err.Error())
	})

	t.Run("Test get proof invalid record input wrong char", func(t *testing.T) {
		record := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aG")
		records := []entity.RecordEntity{record}

		_, err := sdk.GetProof(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get proof invalid record input missing chars", func(t *testing.T) {
		record1 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.GetProof(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test get proof invalid record input wrong start", func(t *testing.T) {
		record1 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994bb")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.GetProof(records)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	//Review because should salt Http error, or should control that the proof exists
	t.Run("Test get proof none existing leaf", func(t *testing.T) {
		record := entity.FromHash("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdee")

		records := []entity.RecordEntity{record}

		_, err := sdk.GetProof(records)
		assert.NotNil(t, err)
		assert.Equal(t, "couldn't get proof for specified records", err.Error())
	})

	t.Run("Test verify records invalid record input wrong char", func(t *testing.T) {
		record := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aG")
		records := []entity.RecordEntity{record}

		_, err := sdk.VerifyRecords(records, configEntity.BloockChain)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test verify records invalid record input missing chars", func(t *testing.T) {
		record1 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("e016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.VerifyRecords(records, configEntity.BloockChain)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	t.Run("Test verify records invalid record input wrong start", func(t *testing.T) {
		record1 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994aa")
		record2 := entity.FromHash("0xe016214a5c4abb88b8b614a916b1a6f075dfcf6fbc16c1e9d6e8ebcec81994bb")

		records := []entity.RecordEntity{record1, record2}

		_, err := sdk.VerifyRecords(records, configEntity.BloockChain)
		assert.NotNil(t, err)
		assert.IsType(t, exception.InvalidRecordException{}, err)
		assert.Equal(t, "Record not valid", err.Error())
	})

	//Review because should be HttpError
	t.Run("Test verify records none existing leaf", func(t *testing.T) {
		record := entity.FromHash("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdee")

		records := []entity.RecordEntity{record}

		_, err := sdk.VerifyRecords(records, configEntity.BloockChain)
		assert.NotNil(t, err)
		assert.Equal(t, "couldn't get proof for specified records", err.Error())
	})
}

func randHex(length int) string {
	maxlength := 8
	min := math.Pow(16, math.Min(float64(length), float64(maxlength))-1)
	max := math.Pow(16, math.Min(float64(length), float64(maxlength))) - 1
	n := int((rand.Float64() * (max - min + 1)) + min)
	r := strconv.Itoa(n)
	for len(r) < length {
		r += randHex(length - maxlength)
	}
	return r
}
