package proof

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProofServiceProof(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	http := mocks.NewHttpClient(mockCtrl)
	s := NewService(http)

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
	
	body := ApiProofRequestBody{
		Messages: hashes,
		Client:   "",
	}

	res := Proof{
		Leaves: []string{"72aae7e86eb51f61a620831320475d9d61cbd52749dbf18fa942b1b97f50aee9"},
		Nodes:  []string{"359b5206452a4ca5058129727fb48f0860a36c0afee0ec62baa874927e9d4b99"},
		Depth:  "020304050501",
		Bitmap: "f4",
	}
	
	jsonRes, _ := json.Marshal(res)
	http.EXPECT().PostRequest("/v1/messages/proof", body).Return(jsonRes, nil).Times(1)

	p, err := s.Proof(hashesBytes)
	if err != nil {
		assert.Fail(t, "no error expected executing Proof service")
	}
	assert.Equal(t, &res, p)
}
