package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Request(verb, endpoint string, body interface{}) (map[string]interface{}, error)
	PostRequest(endpoint string, body interface{}) ([]byte, error)
}

type httpClient struct {
	host string
	port string
}

func NewHttpRequest(host, port string) Client {
	return &httpClient{host, port}
}

func (c *httpClient) Request(verb, endpoint string, body interface{}) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(verb, fmt.Sprintf("%s:%s/%s", c.host, c.port, endpoint), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res map[string]interface{}
	if err := json.Unmarshal(respBytes, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *httpClient) PostRequest(endpoint string, body interface{}) ([]byte, error) {
	bodyRequest, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s:%s/%s", c.host, c.port, endpoint), "application/json", bytes.NewBuffer(bodyRequest))
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
