package http

import (
	"encoding/json"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http/exception"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	apiKey := "123456789abcdef"

	t.Run("Given an api key in a get request should return that api key", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			h := r.Header.Get("X-API-KEY")
			w.Write([]byte(h))
		}))
		defer func() { testServer.Close() }()
		req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
		req.Header.Set("X-API-KEY", apiKey)
		assert.NoError(t, err)

		resp, err := testServer.Client().Do(req)
		respBytes, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, apiKey, string(respBytes), "Api Keys should be equal")
	})

	t.Run("Given an api key in a post request should return that api key", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			h := r.Header.Get("X-API-KEY")
			w.Write([]byte(h))
		}))
		defer func() { testServer.Close() }()
		req, err := http.NewRequest(http.MethodPost, testServer.URL, nil)
		req.Header.Set("X-API-KEY", apiKey)
		assert.NoError(t, err)

		resp, err := testServer.Client().Do(req)
		respBytes, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, apiKey, string(respBytes), "Api Keys should be equal")
	})
	t.Run("Given an invalid JSON response should raise an HttpRequestException", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			err := exception.NewHttpRequestException("")
			w.Write([]byte(err.Error()))
		}))
		defer func() { testServer.Close() }()

		data := NewDataHttp(apiKey)
		httpClient := NewHttp(data)
		resp, _ := httpClient.Get(testServer.URL, nil)
		var err exception.HttpRequestException
		json.Unmarshal(resp, &err)
		assert.NotNil(t, err)
		assert.IsType(t, exception.HttpRequestException{}, err)
		assert.Equal(t, "HttpClient was not successful: ", err.Error())
	})
	t.Run("Given an invalid JSON response should raise an HttpRequestException", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			err := exception.NewHttpRequestException("")
			w.Write([]byte(err.Error()))
		}))
		defer func() { testServer.Close() }()

		data := NewDataHttp(apiKey)
		httpClient := NewHttp(data)
		resp, _ := httpClient.Post(testServer.URL, "", nil)
		var err exception.HttpRequestException
		json.Unmarshal(resp, &err)
		assert.NotNil(t, err)
		assert.IsType(t, exception.HttpRequestException{}, err)
		assert.Equal(t, "HttpClient was not successful: ", err.Error())
	})

}
