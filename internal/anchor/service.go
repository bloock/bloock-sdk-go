package anchor

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	"strconv"
	"strings"
	"time"
)

type Service interface {
	Get(id int) (*GetAnchorResponse, error)
	Wait(anchorId int) (bool, error)
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

func (s *service) Wait(anchorId int) (bool, error) {
	var complete bool
	var attempts int

	var anchorResp *GetAnchorResponse
	for !complete {
		var err error
		anchorResp, err = s.Get(anchorId)
		if err != nil {
			return false, err
		}

		if strings.ToLower(anchorResp.Data.Status) == "success" {
			complete = true
		}

		if complete {
			break
		}

		waitIntervalDefault, _ := strconv.Atoi(s.params.WaitIntervalDefault)
		waitIntervalFactor, _ := strconv.Atoi(s.params.WaitIntervalFactor)
		time.Sleep(time.Duration(waitIntervalDefault+(attempts*waitIntervalFactor)) * time.Millisecond)
	}

	return true, nil
}
