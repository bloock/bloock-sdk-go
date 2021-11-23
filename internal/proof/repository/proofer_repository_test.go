package repository

import (
	"encoding/json"
	"github.com/enchainte/enchainte-sdk-go/internal/config/mockconfig"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/blockchain/mockblockchain"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http/mockhttp"
	"github.com/enchainte/enchainte-sdk-go/internal/proof/entity"
	entity2 "github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRetrieveProofRepository(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	hc := mockhttp.NewMockHttpClient(crtl)
	b := mockblockchain.NewMockBlockchainClient(crtl)
	cs := mockconfig.NewMockConfigurerService(crtl)
	pr := NewProofRepository(hc, b, cs)

	t.Run("Given valid records, should return the proof", func(t *testing.T) {
		nodes := []string{
			"bb6986853646d083929d1d92638f3d4741a3b7149bd2b63c6bfedd32e3c684d3",
			"0616067c793ac533815ae2d48d785d339e0330ce5bb5345b5e6217dd9d1dbeab",
			"68b8f6b25cc700e64ed3e3d33f2f246e24801f93d29786589fbbab3b11f5bcee"}

		resp := entity.Proof{
			Leaves: []string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"},
			Nodes:  nodes,
			Bitmap: "bfdf7000",
			Depth:  "000400060006000500030002000400060007000800090009",
		}

		respBytes, err := json.Marshal(resp)
		require.Nil(t, err)

		cs.EXPECT().GetApiBaseUrl()
		hc.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(respBytes, nil).Times(1)

		r := entity2.FromHash("02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5")
		actual, err := pr.RetrieveProof([]entity2.RecordEntity{r})
		assert.Nil(t, err)
		assert.IsType(t, entity.Proof{}, actual)
		assert.Equal(t, "bfdf7000", actual.Bitmap)
		assert.Equal(t, "000400060006000500030002000400060007000800090009", actual.Depth)
		assert.Equal(t, []string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"}, actual.Leaves)
		assert.Equal(t, nodes, actual.Nodes)
	})
}

func TestVerifyProofRepository(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	hc := mockhttp.NewMockHttpClient(crtl)
	b := mockblockchain.NewMockBlockchainClient(crtl)
	cs := mockconfig.NewMockConfigurerService(crtl)
	pr := NewProofRepository(hc, b, cs)

	t.Run("Given a proof, should verify it", func(t *testing.T) {
		leaves := []string{
			"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5",
			"5e1712aca5f3925fc0ce628e7da2e1e407e2cc7b358e83a7152b1958f7982dab",
		}
		nodes := []string{
			"1ca0e9d9a206f08d38a4e2cf485351674ffc9b0f3175e0cb6dbd8e0e19829b97",
			"1ca0e9d9a206f08d38a4e2cf485351674ffc9b0f3175e0cb6dbd8e0e19829b97",
			"54944fcea707a57048c17ca7453fa5078a031143b44629776750e7f0ff7940f0",
			"d6f9bcd042be70b39b65dc2a8168858606b0a2fcf6d02c0a1812b1804efc0c37",
			"e663ec001b81b96eceabd1b766d49ec5d99adedc3e5f03d245b0d90f603f66d3",
		}
		depth := "0004000400030004000400030001"
		bitmap := "7600"
		root := "a1fd8b878cee593a7debf12b5bcbf081a972bbec40e103c6d82197db2751ced7"

		re, err := pr.VerifyProof(entity.NewProof(leaves, nodes, depth, bitmap))
		assert.Nil(t, err)
		assert.Equal(t, root, re.GetHash())
	})

	t.Run("Given a proof, should verify it 2", func(t *testing.T) {
		leaves := []string{
			"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5",
			"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5",
			"5e1712aca5f3925fc0ce628e7da2e1e407e2cc7b358e83a7152b1958f7982dab",
		}
		nodes := []string{
			"1ca0e9d9a206f08d38a4e2cf485351674ffc9b0f3175e0cb6dbd8e0e19829b97",
			"1ca0e9d9a206f08d38a4e2cf485351674ffc9b0f3175e0cb6dbd8e0e19829b97",
			"1509877db1aa81c699a144d1a240c5d463c9ff08b2df489b40a35802844baeb6",
			"54944fcea707a57048c17ca7453fa5078a031143b44629776750e7f0ff7940f0",
			"d6f9bcd042be70b39b65dc2a8168858606b0a2fcf6d02c0a1812b1804efc0c37",
			"e663ec001b81b96eceabd1b766d49ec5d99adedc3e5f03d245b0d90f603f66d3",
		}
		depth := "000500050004000400040004000400030001"
		bitmap := "6d80"
		root := "7e1f3c7e6d3515389b6117cc8c1ef5512d51c59743dc097c70de405a91861d2b"

		re, err := pr.VerifyProof(entity.NewProof(leaves, nodes, depth, bitmap))
		assert.Nil(t, err)
		assert.Equal(t, root, re.GetHash())
	})

	t.Run("Given a proof, should verify it 3", func(t *testing.T) {
		leaves := []string{"0000000000000000000000000000000000000000000000000000000000000000"}
		nodes := []string{"f49d70da1c2c8989766908e06b8d2277a6954ec8533696b9a404b631b0b7735a"}
		depth := "00010001"
		bitmap := "4000"
		root := "5c67902dc31624d9278c286ef4ce469451d8f1d04c1edb29a5941ca0e03ddc8d"

		re, err := pr.VerifyProof(entity.NewProof(leaves, nodes, depth, bitmap))
		assert.Nil(t, err)
		assert.Equal(t, root, re.GetHash())
	})
}
