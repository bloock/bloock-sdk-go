package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	"strconv"
	"time"
)

type Service interface {
	Write(hash []byte) error
	Search(hash [][]byte) (*Receipts, error)
	Wait(hashes [][]byte) (*Receipts, error)
}

var (
	messagesStack []Message
	channel       chan SendResponse
	done          chan bool
)

type service struct {
	apiKey string
	http   http.Client
	params cloud.SdkParams
}

func NewService(apiKey string, http http.Client, params cloud.SdkParams) Service {
	channel = make(chan SendResponse)
	done = make(chan bool)
	s := &service{apiKey, http, params}
	go s.scheduler(done)
	return s
}

type Receipts struct {
	Messages []Receipt `json:"messages"`
}

type Receipt struct {
	Message string `json:"message"`
	Anchor  int    `json:"anchor"`
	Client  string `json:"client"`
	Status  string `json:"status"`
}

type ReceiptOld struct {
	Root    string
	Message string
	TxHash  string
	Status  string
	Error   string
}

type SendResponse struct {
	Body  WriteResponse
	Error error
}

func (s *service) Write(hash []byte) error {
	message, err := New(hash)
	if err != nil {
		return err
	}
	messagesStack = append(messagesStack, *message)

	return nil
}

// TODO should receive client ID parameter
// TODO messages should be: the message string in bytes or the message hash in bytes??
func (s *service) Search(messages [][]byte) (*Receipts, error) {
	var hashes []string
	for _, message := range messages {
		m, err := New(message)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, m.hash)
	}

	body := FetchRequest{
		Messages: hashes,
		// TODO client id
		Client: "e7f97c33-ab8c-48b8-b349-fd3ef9ce974e",
	}

	// TODO remove hardcoded
	resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s%s", "http://localhost:3000", "/v1/messages/fetch"), body)
	//resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s/v1%s", s.params.Host, s.params.MessageFetch), nil, body)
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}
	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, err
	}

	if res["status"] == "error" {
		return nil, errors.New(fmt.Sprintf("%v", res["message"]))
	}

	var receipts Receipts
	bytes, _ := json.Marshal(res)
	if err := json.Unmarshal(bytes, &receipts); err != nil {
		return nil, err
	}

	return &receipts, nil
}

func (s *service) Wait(messages [][]byte) (*Receipts, error) {
	var complete bool
	var attempts int

	var receipts *Receipts
	for !complete {
		var err error
		receipts, err = s.Search(messages)
		if err != nil {
			return nil, err
		}

		if len(receipts.Messages) >= len(messages) {
			for _, r := range receipts.Messages {
				if r.Status == "success" || r.Status == "error" {
					complete = true
				} else {
					complete = false
					break
				}
			}
		}

		if complete {
			break
		}

		waitIntervalDefault, _ := strconv.Atoi(s.params.WaitIntervalDefault)
		waitIntervalFactor, _ := strconv.Atoi(s.params.WaitIntervalFactor)
		time.Sleep(time.Duration(waitIntervalDefault+(attempts*waitIntervalFactor)) * time.Millisecond)
	}

	return receipts, nil
}

// send does a POST request to Enchainté's API to write all hashes stored in the message stack to the blockchain.
func (s *service) send() (*WriteResponse, error) {
	if len(messagesStack) == 0 {
		return nil, nil
	}

	var hashes []string
	for _, message := range messagesStack {
		hashes = append(hashes, message.Hash())
	}

	body := WriteRequest{
		Messages: hashes,
		// TODO add client id param
		Client: "",
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

	if respMap["status"] == "error" {
		return nil, errors.New(fmt.Sprintf("%v", respMap["message"]))
	}

	var anchor WriteResponse
	bytes, _ := json.Marshal(respMap)
	if err := json.Unmarshal(bytes, &anchor); err != nil {
		return nil, err
	}

	return &anchor, nil
}

// scheduler executes periodically the checkStack method
func (s *service) scheduler(done chan bool) {
	interval, _ := strconv.Atoi(s.params.WriteInterval)
	ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case <-ticker.C:
			s.checkStack()
		}
	}
}

// checkStack executes the send method when the message stack is not empty and sends the result to the channel
func (s *service) checkStack() {
	if messagesStack == nil {
		//fmt.Println("Empty stack... leaving")
		//fmt.Println()
		return
	}

	//fmt.Println("Sending message stack...")
	//fmt.Println(messagesStack)
	//fmt.Println()
	resp, err := s.send()

	messagesStack = nil

	channel <- SendResponse{
		Body:  *resp,
		Error: err,
	}
	return
}

// Receive returns the value held by the channel
func Receive() SendResponse {
	return <-channel
}

// Done closes the Ticker channel
func Done() {
	done <- true
}
