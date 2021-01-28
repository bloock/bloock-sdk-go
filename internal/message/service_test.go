package message_test

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageServiceSearch(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	http := mocks.NewHttpClient(mockCtrl)
	s := message.NewService(http)


	// should successfully return the receipts
	var messages [][]byte
	messages = append(messages, []byte("first message"))
	messages = append(messages, []byte("second message"))

	var hashes []string
	for _, mb := range messages {
		m, err := message.New(mb)
		if err != nil {
			assert.Fail(t, fmt.Sprintf("error not expected: %s", err.Error()))
		}
		hashes = append(hashes, m.Hash())
	}

	body := message.ApiFetchRequestBody{
		Messages: hashes,
		Client:   "",
	}

	res := message.Receipts{
		Messages: []message.Receipt{
			{
				Message: "13aae7e862151f61a64183a322475d0d60cbd52719dbf28f0942b1b97f50aee0",
				Anchor:  1,
				Client:  "9281b537-b78b-4d2b-8f7d-9421a9f0ffbd",
				Status:  "success",
			},
			{
				Message: "13aae7e862151f61a64183a322475d0d60cbd52719dbf28f0942b1b97f50aee0",
				Anchor:  2,
				Client:  "9281b537-b78b-4d2b-8f7d-9421a9f0ffbd",
				Status:  "success",
			},
		}}

	jsonREs, _ := json.Marshal(res)
	http.EXPECT().PostRequest("/v1/messages/fetch", body).Return(jsonREs, nil).Times(1)

	receipts, err := s.Search(messages)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error not expected searching messages: %s", err.Error()))
	}
	assert.Equal(t, &res, receipts)

}

