package test

import (
	"github.com/bloock/bloock-sdk-go"
	anchorEntity "github.com/bloock/bloock-sdk-go/internal/anchor/entity"
	configEntity "github.com/bloock/bloock-sdk-go/internal/config/entity"
	proofEntity "github.com/bloock/bloock-sdk-go/internal/proof/entity"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func GetSdk() bloock.BloockClient {
	apiKey := os.Getenv("API_KEY")
	apiHost := os.Getenv("API_HOST")
	client := bloock.NewBloockClient(apiKey)
	client.SetApiHost(apiHost)
	return client
}

func TestFunctionalSendRecord(t *testing.T) {
	sdk := GetSdk()

	records := make([]entity.RecordEntity, 0)
	records = append(records, entity.FromString("Example Data 1"))
	records = append(records, entity.FromString("Example Data 2"))
	records = append(records, entity.FromString("Example Data 3"))

	r, err := sdk.SendRecords(records)
	assert.Nil(t, err)
	assert.IsType(t, entity.RecordReceipt{}, r[0])
	assert.Greater(t, r[0].Anchor, 0)
	assert.Greater(t, len(r[0].Client), 0)
	assert.Equal(t, r[0].Record, records[0].GetHash())
	assert.Equal(t, "Pending", r[0].Status)
}

func TestFunctionalWaitAnchor(t *testing.T) {
	sdk := GetSdk()

	records := make([]entity.RecordEntity, 0)
	records = append(records, entity.FromString("Example Data 1"))
	records = append(records, entity.FromString("Example Data 2"))
	records = append(records, entity.FromString("Example Data 3"))

	r, err := sdk.SendRecords(records)
	assert.Nil(t, err)
	assert.IsType(t, entity.RecordReceipt{}, r[0])
	assert.NotNil(t, r)
	assert.NotEqual(t, entity.RecordReceipt{}, r[0])

	a, err := sdk.WaitAnchor(r[0].Anchor, anchorEntity.AnchorParams{})
	assert.Nil(t, err)
	assert.IsType(t, anchorEntity.Anchor{}, a)
	assert.Greater(t, a.ID(), 0)
	assert.Greater(t, len(a.BlockRoots()), 0)
	assert.Greater(t, len(a.Networks()), 0)
	assert.Greater(t, len(a.Root()), 0)
	assert.Greater(t, len(a.Status()), 0)
}

func TestFunctionalFetchRecords(t *testing.T) {
	sdk := GetSdk()

	records := make([]entity.RecordEntity, 0)
	records = append(records, entity.FromString("Example Data 1"))
	records = append(records, entity.FromString("Example Data 2"))
	records = append(records, entity.FromString("Example Data 3"))

	r, err := sdk.SendRecords(records)
	assert.Nil(t, err)
	assert.IsType(t, entity.RecordReceipt{}, r[0])
	assert.NotNil(t, r)
	assert.NotEqual(t, entity.RecordReceipt{}, r[0])

	sdk.WaitAnchor(r[0].Anchor, anchorEntity.AnchorParams{})

	rr, err := sdk.GetRecords(records)
	assert.Nil(t, err)
	for _, r := range rr {
		assert.Equal(t, "Success", r.Status)
	}
}

func TestFunctionalGetProof(t *testing.T) {
	sdk := GetSdk()

	records := make([]entity.RecordEntity, 0)
	records = append(records, entity.FromString("Example Data 1"))
	records = append(records, entity.FromString("Example Data 2"))
	records = append(records, entity.FromString("Example Data 3"))

	p, err := sdk.GetProof(records)
	assert.Nil(t, err)
	assert.IsType(t, proofEntity.Proof{}, p)
	assert.NotNil(t, p)
	assert.NotEqual(t, proofEntity.Proof{}, p)
}

func TestFunctionalVerifyProof(t *testing.T) {
	sdk := GetSdk()

	records := make([]entity.RecordEntity, 0)
	records = append(records, entity.FromString("Example Data 1"))
	records = append(records, entity.FromString("Example Data 2"))
	records = append(records, entity.FromString("Example Data 3"))

	p, err := sdk.GetProof(records)
	assert.Nil(t, err)
	assert.IsType(t, proofEntity.Proof{}, p)
	assert.NotNil(t, p)
	assert.NotEqual(t, proofEntity.Proof{}, p)

	timestamp, err := sdk.VerifyProof(p, configEntity.NetworkParams{Network: configEntity.BloockChain})
	assert.Nil(t, err)
	assert.Greater(t, timestamp, 0)
}
