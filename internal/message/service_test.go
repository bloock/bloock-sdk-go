package message_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMessageServiceWrite(t *testing.T) {
	var datasets = [][]byte{[]byte("message 1"), []byte("message 2")}
	var smtErr = errors.New("smt error")
	var httpErr = errors.New("something went wrong")

	testCases := []struct {
		name        string
		inDatasets  [][]byte
		buildStubs  func(mocks.HttpClient)
		outExpected func(*message.WriteResponse, error)
	}{
		{
			name:       "success",
			inDatasets: datasets,
			buildStubs: func(http mocks.HttpClient) {
				var hashes []string
				for _, data := range datasets {
					m, err := message.New(data)
					if err != nil {
						require.Nil(t, err)
					}
					hashes = append(hashes, m.Hash())
				}
				body := message.WriteRequest{
					Messages: hashes,
				}
				wr := message.WriteResponse{
					Success: true,
					Data:    &message.WriteResponseData{Anchor: 1},
				}
				apiResp, err := json.Marshal(wr)
				require.Nil(t, err)
				http.EXPECT().Request(gomock.Any(), "POST", "http://localhost:3000/v1/messages", body).Return(apiResp, nil).Times(1)
			},
			outExpected: func(resp *message.WriteResponse, err error) {
				require.Nil(t, err)

				require.True(t, resp.Success)
				require.Equal(t, "int", fmt.Sprintf("%T", resp.Data.Anchor))
				require.Equal(t, 1, resp.Data.Anchor)
			},
		},
		{
			name:       "anchor_failure",
			inDatasets: datasets,
			buildStubs: func(http mocks.HttpClient) {
				var hashes []string
				for _, data := range datasets {
					m, err := message.New(data)
					if err != nil {
						require.Nil(t, err)
					}
					hashes = append(hashes, m.Hash())
				}
				body := message.WriteRequest{
					Messages: hashes,
				}
				wr := map[string]interface{}{
					"success": false,
					"error": map[string]interface{}{
						"message": "smt error",
					},
				}
				apiResp, err := json.Marshal(wr)
				require.Nil(t, err)
				http.EXPECT().Request(gomock.Any(), "POST", "http://localhost:3000/v1/messages", body).Return(apiResp, nil).Times(1)
			},
			outExpected: func(resp *message.WriteResponse, err error) {
				require.Nil(t, resp)
				require.Equal(t, smtErr, err)
			},
		},
		{
			name:       "error_http_request",
			inDatasets: datasets,
			buildStubs: func(http mocks.HttpClient) {
				var hashes []string
				for _, data := range datasets {
					m, err := message.New(data)
					if err != nil {
						require.Nil(t, err)
					}
					hashes = append(hashes, m.Hash())
				}
				body := message.WriteRequest{
					Messages: hashes,
				}
				http.EXPECT().Request(gomock.Any(), "POST", "http://localhost:3000/v1/messages", body).Return(nil, httpErr).Times(1)
			},
			outExpected: func(resp *message.WriteResponse, err error) {
				require.Nil(t, resp)

				require.NotNil(t, err)
				require.Equal(t, httpErr, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			http := mocks.NewHttpClient(mockCtrl)
			sdkParams := cloud.SdkParams{}
			s := message.NewService("", http, sdkParams)

			tc.buildStubs(*http)
			resp, err := s.Write(tc.inDatasets)
			tc.outExpected(resp, err)
		})
	}
}

func TestMessageServiceFetch(t *testing.T) {
	var datasets = [][]byte{[]byte("message 1"), []byte("message 2")}
	var receipts = []message.Receipt{
		{
			Message: "13aae7e862151f61a64183a322475d0d60cbd52719dbf28f0942b1b97f50aee0",
			Anchor:  1,
			Client:  "9281b537-b78b-4d2b-8f7d-9421a9f0ffbd",
			Status:  "success",
		},
		{
			Message: "987ae7e862151f61a64183a322475d0d60cbd52719dbf28f0942b1b97f50a123",
			Anchor:  2,
			Client:  "9281b537-b78b-4d2b-8f7d-9421a9f0ffbd",
			Status:  "success",
		},
	}
	var smtErr = errors.New("smt error")
	var httpErr = errors.New("something went wrong")

	testCases := []struct {
		name        string
		inDatasets  [][]byte
		buildStubs  func(mocks.HttpClient)
		outExpected func(*[]message.Receipt, error)
	}{
		{
			name:       "success",
			inDatasets: datasets,
			buildStubs: func(http mocks.HttpClient) {
				var hashes []string
				for _, data := range datasets {
					m, err := message.New(data)
					if err != nil {
						require.Nil(t, err)
					}
					hashes = append(hashes, m.Hash())
				}
				body := message.FetchRequest{
					Messages: hashes,
				}
				sr := message.SearchMessageResponse{
					Success: true,
					Data:    &receipts,
				}
				apiResp, err := json.Marshal(sr)
				require.Nil(t, err)
				http.EXPECT().Request(gomock.Any(), "POST", "http://localhost:3000/v1/messages/fetch", body).Return(apiResp, nil).Times(1)
			},
			outExpected: func(resp *[]message.Receipt, err error) {
				require.Nil(t, err)
				require.Equal(t, &receipts, resp)
			},
		},
		{
			name:       "fetch_message_failure",
			inDatasets: datasets,
			buildStubs: func(http mocks.HttpClient) {
				var hashes []string
				for _, data := range datasets {
					m, err := message.New(data)
					if err != nil {
						require.Nil(t, err)
					}
					hashes = append(hashes, m.Hash())
				}
				body := message.FetchRequest{
					Messages: hashes,
				}
				sr := map[string]interface{}{
					"success": false,
					"error": map[string]interface{}{
						"message": "smt error",
					},
				}
				apiResp, err := json.Marshal(sr)
				require.Nil(t, err)
				http.EXPECT().Request(gomock.Any(), "POST", "http://localhost:3000/v1/messages/fetch", body).Return(apiResp, nil).Times(1)
			},
			outExpected: func(resp *[]message.Receipt, err error) {
				require.Nil(t, resp)
				require.NotNil(t, err)
				require.Equal(t, smtErr, err)
			},
		},
		{
			name:       "http_request_error",
			inDatasets: datasets,
			buildStubs: func(http mocks.HttpClient) {
				var hashes []string
				for _, data := range datasets {
					m, err := message.New(data)
					if err != nil {
						require.Nil(t, err)
					}
					hashes = append(hashes, m.Hash())
				}
				body := message.FetchRequest{
					Messages: hashes,
				}
				http.EXPECT().Request(gomock.Any(), "POST", "http://localhost:3000/v1/messages/fetch", body).Return(nil, httpErr).Times(1)
			},
			outExpected: func(resp *[]message.Receipt, err error) {
				require.Nil(t, resp)
				require.NotNil(t, err)
				require.Equal(t, httpErr, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			http := mocks.NewHttpClient(mockCtrl)
			sdkParams := cloud.SdkParams{}
			s := message.NewService("", http, sdkParams)

			tc.buildStubs(*http)
			resp, err := s.Search(tc.inDatasets)
			tc.outExpected(resp, err)
		})
	}
}
