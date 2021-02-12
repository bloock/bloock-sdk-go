package message_test

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
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
	cts := config.Constants{}

	s := message.NewService("", http, cts)

	defer message.Done()


	// should successfully return the receipts
	var messages [][]byte
	messages = append(messages, []byte("first Message"))
	messages = append(messages, []byte("second Message"))

	var hashes []string
	for _, mb := range messages {
		m, err := message.New(mb)
		if err != nil {
			assert.Fail(t, fmt.Sprintf("error not expected: %s", err.Error()))
		}
		hashes = append(hashes, m.Hash())
	}

	body := message.FetchRequest{
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

	jsonRes, _ := json.Marshal(res)
	http.EXPECT().Request("", "POST", gomock.Any(), nil, body).Return(jsonRes, nil)

	receipts, err := s.Search(messages)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error not expected searching messages: %s", err.Error()))
	}
	assert.Equal(t, &res, receipts)
}

func TestMessageServiceWait(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	http := mocks.NewHttpClient(mockCtrl)
	cts := config.Constants{}

	s := message.NewService("", http, cts)

	defer message.Done()


	var messages [][]byte
	messages = append(messages, []byte("first Message"))
	messages = append(messages, []byte("second Message"))

	var search1ItResp = message.Receipts{
		Messages: []message.Receipt{
			{
				Message: "13aae7e862151f61a64183a322475d0d60cbd52719dbf28f0942b1b97f50aee0",
				Anchor:  1,
				Client:  "9281b537-b78b-4d2b-8f7d-9421a9f0ffbd",
				Status:  "success",
			},
		},
	}
	bytes1It, err := json.Marshal(search1ItResp)
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("error marashaling: %s", err.Error()))
	}

	var searchResp = message.Receipts{
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
		},
	}
	bytesResp, err := json.Marshal(searchResp)
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("error marashaling: %s", err.Error()))
	}

	// should successfully return the receipts doing 1 iteration
	http.EXPECT().Request("", "POST", gomock.Any(), nil, gomock.Any()).Return(bytesResp, nil).Times(1)

	receipts, err := s.Wait(messages)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error not expected searching messages: %s", err.Error()))
	}
	assert.Equal(t, &searchResp, receipts)


	// should successfully return the receipts doing 2 iteration
	http.EXPECT().Request("", "POST", gomock.Any(), nil, gomock.Any()).Return(bytes1It, nil).Times(1)
	http.EXPECT().Request("", "POST", gomock.Any(), nil, gomock.Any()).Return(bytesResp, nil).Times(1)

	receipts, err = s.Wait(messages)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error not expected searching messages: %s", err.Error()))
	}
	assert.Equal(t, &searchResp, receipts)
}

func TestMessageServiceWrite(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	http := mocks.NewHttpClient(mockCtrl)
	cts := config.Constants{}

	s := message.NewService("", http, cts)

	defer message.Done()

	var messages [][]byte
	messages = append(messages, []byte("first Message"))
	messages = append(messages, []byte("second Message"))

	writeResponse := message.WriteResponse{Anchor: 2}
	bytesResp, err := json.Marshal(writeResponse)
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("error marashaling: %s", err.Error()))
	}
	http.EXPECT().Request("", "POST", gomock.Any(), nil, gomock.Any()).Return(bytesResp, nil).Times(1)

	for _, m := range messages {
		if err := s.Write(m); err != nil {
			assert.FailNow(t, fmt.Sprintf("unexpected error: %s", err.Error()))
		}
	}
	sendResponse := message.SendResponse{Body: writeResponse}

	res := message.Receive()
	assert.Equal(t, sendResponse, res)
}
