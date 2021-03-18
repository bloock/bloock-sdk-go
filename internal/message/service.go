package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
)

type Service interface {
	Write(datasets [][]byte) (*WriteResponse, error)
	Search(hash [][]byte) (*[]Receipt, error)
}

type service struct {
	apiKey string
	http   http.Client
	params cloud.SdkParams
}

func NewService(apiKey string, http http.Client, params cloud.SdkParams) Service {
	return &service{apiKey, http, params}
}

type SendResponse struct {
	Body  WriteResponse
	Error error
}

func (s *service) Write(dataset [][]byte) (*WriteResponse, error) {
	var messages []Message
	for _, data := range dataset {
		message, _ := New(data)
		messages = append(messages, *message)
	}

	resp, err := s.send(messages)
	if err != nil {
		return  nil, err
	}
	return resp, nil
}

// TODO should receive client ID parameter
// Search retrieves the information related to the provided messages. It takes an array bytes as parameter and returns
// a Receipts struct or an error if something goes wrong.
func (s *service) Search(messages [][]byte) (*[]Receipt, error) {
	var hashes []string
	for _, message := range messages {
		m, err := New(message)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, m.Hash())
	}

	body := FetchRequest{
		Messages: hashes,
		// TODO client id
		//Client: "e7f97c33-ab8c-48b8-b349-fd3ef9ce974e",
	}

	// TODO remove hardcoded
	resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s%s", "http://localhost:3000", "/v1/messages/fetch"), body)
	//resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s/v1%s", s.params.Host, s.params.MessageFetch), nil, body)
	if err != nil {
		return nil, err
	}

	var respMap map[string]interface{}
	if err := json.Unmarshal(resp, &respMap); err != nil {
		return nil, err
	}

	if respMap["success"] == false {
		return nil, errors.New(fmt.Sprintf("%v", respMap["error"].(map[string]interface{})["message"]))
	}

	var searchMessage SearchMessageResponse
	bytes, _ := json.Marshal(respMap)
	if err := json.Unmarshal(bytes, &searchMessage); err != nil {
			return nil, err
		}

	return searchMessage.Data, nil
}

// send does a POST request to Enchainté's API to write all hashes stored in the message stack to the blockchain.
func (s *service) send(messages []Message) (*WriteResponse, error) {
	var hashes []string
	for _, message := range messages {
		hashes = append(hashes, message.Hash())
	}

	body := WriteRequest{
		Messages: hashes,
		// TODO add client id param
		//Client: "",
	}
	// TODO remove hardcoded
	resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s%s", "http://localhost:3000", "/v1/messages"), body)
	//resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s/v1%s", s.params.Host, s.params.MessageWrite), nil, body)
	if err != nil {
		return nil, err
	}

	var respMap map[string]interface{}
	if err := json.Unmarshal(resp, &respMap); err != nil {
		return nil, err
	}

	if respMap["success"] == false {
		return nil, errors.New(fmt.Sprintf("%v", respMap["error"].(map[string]interface{})["message"]))
	}

	var anchor WriteResponse
	bytes, _ := json.Marshal(respMap)
	if err := json.Unmarshal(bytes, &anchor); err != nil {
		return nil, err
	}

	return &anchor, nil
}
