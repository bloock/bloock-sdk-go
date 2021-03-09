package anchor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
)

type Service interface {
	Get(id int) (*GetAnchorResponse, error)
}

type service struct {	
	apiKey string
	http   http.Client
	params cloud.SdkParams
}

func NewService(apiKey string, http http.Client, params cloud.SdkParams) Service {
	return &service{apiKey, http, params}
}

func (s *service) Get(id int) (*GetAnchorResponse, error) {
	// TODO add azure endpoint
	resp, err := s.http.Request(s.apiKey, "GET", fmt.Sprintf("%s%s%d", "http://localhost:3000", "/v1/anchors/", id), nil)
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}

	if !res["success"].(bool) {
		return nil, errors.New(fmt.Sprintf("%v", res["error"].(map[string]interface{})["message"]))
	}

	var data GetAnchorResponse
	bytes, _ := json.Marshal(res)
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
