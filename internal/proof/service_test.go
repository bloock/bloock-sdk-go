package proof_test

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/internal/mocks"
	"github.com/enchainte/enchainte-sdk-go/internal/proof"
	"github.com/enchainte/enchainte-sdk-go/pkg/crypto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProofServiceProof(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	http := mocks.NewHttpClient(mockCtrl)
	hasher := mocks.NewHasher(mockCtrl)
	sdkParams := cloud.SdkParams{}
	bc := mocks.NewBlockchainClient(mockCtrl)
	s := proof.NewService("", http, sdkParams, hasher, bc)

	var hashesBytes [][]byte
	hashesBytes = append(hashesBytes, []byte("first hash"))
	hashesBytes = append(hashesBytes, []byte("second hash"))

	var hashes []string
	for _, hb := range hashesBytes {
		m, err := message.New(hb)
		if err != nil {
			assert.Fail(t, fmt.Sprintf("error not expected: %s", err.Error()))
		}
		hashes = append(hashes, m.Hash())
	}

	body := proof.ApiProofRequestBody{
		Messages: hashes,
		//Client:   "",
	}

	res := proof.ProofResponse{
		Success: true,
		Data: &proof.Proof{
			Leaves: []string{"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9"},
			Nodes:  []string{"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99"},
			Depth:  "020304050501",
			Bitmap: "f4",
		},
	}
	jsonRes, _ := json.Marshal(res)
	http.EXPECT().Request("", "POST", gomock.Any(), body).Return(jsonRes, nil).Times(1)

	p, err := s.Proof(hashesBytes)
	if err != nil {
		assert.Fail(t, "no error expected executing Proof service")
	}
	assert.Equal(t, res.Data, p)
}

func TestProofServiceCalculateRoot(t *testing.T) {
	var leaves = []string{"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9"}
	var nodes = []string{
		"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99",
		"707cb86e449cd3990c85fb3ae9ec967ee12b82f21eae9e6ea35180e6c331c3e8",
		"23950edeb3ca719e814d8b04d63d90d39327b49b7df5baf2f72305c1f2b260b7",
		"72aae7e86eb50f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
		"517e320992fb35553575750153992d6360268d04a1e4d9e2cae7e5c3736ac627",
	}
	var depth = "020304050501"
	var bitmap = "f4"
	var root = "6608fd2c5d9c28124b41d6e441d552ad811a51fc6fdae0f33aa64bf3f43ca699"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	http := mocks.NewHttpClient(mockCtrl)
	hasher := crypto.Blake2b()
	sdkParams := cloud.SdkParams{}
	bc := mocks.NewBlockchainClient(mockCtrl)

	s := proof.NewService("", http, sdkParams, hasher, bc)

	proof, err := proof.New(leaves, nodes, depth, bitmap)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("unexpected error: %s", err.Error()))
	}

	var res string
	res, err = s.CalculateRoot(proof)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("unexpected error: %s", err.Error()))
	}

	assert.Equal(t, root, res)
}

func TestProofServiceVerify(t *testing.T) {

	var proofResponse = proof.ProofResponse{
		Success: true,
		Data: &proof.Proof{
			Leaves: []string{"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9"},
			Nodes: []string{
				"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99",
				"707cb86e449cd3990c85fb3ae9ec967ee12b82f21eae9e6ea35180e6c331c3e8",
				"23950edeb3ca719e814d8b04d63d90d39327b49b7df5baf2f72305c1f2b260b7",
				"72aae7e86eb50f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9",
				"517e320992fb35553575750153992d6360268d04a1e4d9e2cae7e5c3736ac627",
			},
			Depth:  "020304050501",
			Bitmap: "f4",
		},
	}

	var root = "6608fd2c5d9c28124b41d6e441d552ad811a51fc6fdae0f33aa64bf3f43ca699"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	http := mocks.NewHttpClient(mockCtrl)
	hasher := crypto.Blake2b()
	sdkParams := cloud.SdkParams{}
	bc := mocks.NewBlockchainClient(mockCtrl)

	s := proof.NewService("", http, sdkParams, hasher, bc)

	jsonRes, _ := json.Marshal(proofResponse)
	http.EXPECT().Request("", "POST", gomock.Any(), gomock.Any()).Return(jsonRes, nil).Times(1)

	bc.EXPECT().ValidateRoot(root).Return(true, nil).Times(1)

	var hashesBytes [][]byte
	hashesBytes = append(hashesBytes, []byte("first hash"))
	hashesBytes = append(hashesBytes, []byte("second hash"))

	// the result will be mock so no interest in the result
	_, err := s.Verify(hashesBytes)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("unexpected error: %s", err.Error()))
	}
}
