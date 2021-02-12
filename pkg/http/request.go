package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Request(apiKey, verb, url string, headers map[string]string, body interface{}) ([]byte, error)
	PostRequest(url string, body interface{}) ([]byte, error)
}

type httpClient struct {
}

func NewClient() Client {
	return &httpClient{}
}

func (c *httpClient) Request(apiKey, verb, url string, headers map[string]string, body interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(verb, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s" , apiKey))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *httpClient) PostRequest(url string, body interface{}) ([]byte, error) {
	bodyRequest, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyRequest))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
