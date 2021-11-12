package credential

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
)

type Service interface {
	Delete(apiKey, id string) error
}

type service struct {
	apiKey string
	http   http.Client
	params cloud.SdkParams
}

func NewService(apiKey string, http http.Client, params cloud.SdkParams) Service {
	return &service{apiKey, http, params}
}

func (s *service) Delete(apiKey, id string) error {
	body := DeleteApiKeyRequest{ApiKey: apiKey}

	// TODO add azure endpoint
	resp, err := s.http.Request(s.apiKey, "DELETE", fmt.Sprintf("%s%s%s%s", "http://localhost:3000", "/v1/clients/", id, "/credentials"), body)
	if err != nil {
		return err
	}

	var res map[string]interface{}
	if err := json.Unmarshal(resp, &res); err != nil {
		return err
	}

	if !res["success"].(bool) {
		return errors.New(fmt.Sprintf("%v", res["error"].(map[string]interface{})["record"]))
	}

	return nil
}
