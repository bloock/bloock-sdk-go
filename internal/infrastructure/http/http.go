package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/http/exception"
	"io"
	http "net/http"
	"time"
)

type Http struct {
	httpData DataHttp
}

func NewHttp(data DataHttp) Http {
	return Http{
		httpData: data,
	}
}

func (h Http) Get(url string, headers map[string]string) (interface{}, error) {
	return h.request("GET", url, headers, nil)
}

func (h Http) Post(url string, body interface{}, headers map[string]string) (interface{}, error) {
	return h.request("POST", url, headers, body)
}

func (h Http) request(method, url string, headers map[string]string, body interface{}) (interface{}, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("http.marshal: %s", err)
	}

	client := http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, exception.NewHttpRequestException(err.Error())
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("X-API-KEY", h.httpData.GetApiKey())
	response, err := client.Do(req)
	if err != nil {
		return nil, exception.NewHttpRequestException(err.Error())
	}
	defer response.Body.Close()

	respByte, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, exception.NewHttpRequestException(err.Error())
	}

	var resp interface{}
	if err := json.Unmarshal(respByte, &resp); err != nil {
		return nil, exception.NewHttpRequestException(err.Error())
	}
	return resp, nil
}
