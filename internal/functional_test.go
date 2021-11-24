package internal

import (
	anchorEntity "github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	configEntity "github.com/enchainte/enchainte-sdk-go/internal/config/entity"
	proofEntity "github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetSdk() BloockClient {
	apiKey := "test_xculO0olb1Itp-tFMNCjpsLgx4Bik3E7Wd-iUfdL1c2lsgyKvhAZQnd7U8vlPnJX" //clau de l'entorn de test de pro
	apiHost := "https://api.bloock.com" //endpoint de pro
	client := NewBloockClient(apiKey)
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

	a, err := sdk.WaitAnchor(r[0].Anchor, 5000)
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

	sdk.WaitAnchor(r[0].Anchor, 5000)

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

	timestamp, err := sdk.VerifyProof(p, configEntity.BloockChain)
	assert.Nil(t, err)
	assert.Greater(t, timestamp, 0)
}

