package http

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	apiKey := "123456789abcdef"

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		log.Println(w.Header())
	}))
	defer func() { testServer.Close() }()

	req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	req.Header.Set("X-API-KEY", apiKey)
	assert.NoError(t, err)

	resp, err := testServer.Client().Do(req)
	//resp, err := http.DefaultClient.Do(req)
	log.Println(resp)
	assert.NoError(t, err)
	assert.Equal(t, apiKey, resp.Header.Get("X-API-KEY"), "")

}
