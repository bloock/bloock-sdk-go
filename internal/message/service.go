package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
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
func (s *service) Search(messageBytes [][]byte) (*Receipts, error) {
	var hashes []string
	for _, mb := range messageBytes {
		m, err := New(mb)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, m.hash)
	}

	body := FetchRequest{
		Messages: hashes,
		// TODO client id
		Client: "",
	}

	resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s%s", s.params.Host, s.params.MessageFetch), nil, body)
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

func (s *service) Wait(hashes [][]byte) (*Receipts, error) {
	var complete bool
	var attempts int

	var receipts Receipts
	for !complete {
		res, err := s.Search(hashes)
		if err != nil {
			return nil, err
		}

		bytes, _ := json.Marshal(res)
		if err := json.Unmarshal(bytes, &receipts); err != nil {
			return nil, err
		}

		if len(receipts.Messages) < len(hashes) {
			continue
		}

		for _, r := range receipts.Messages {
			if r.Status == "success" || r.Status == "error" {
				complete = true
			} else {
				complete = false
				break
			}
		}

		if complete {
			break
		}

		time.Sleep(time.Duration(s.params.WaitIntervalDefault+(attempts*s.params.WaitIntervalFactor)) * time.Second)

		attempts++
	}

	return &receipts, nil
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
		// TODO
		Client: "",
	}
	resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s%s", s.params.Host, s.params.MessageWrite), nil, body)
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

	var anchor WriteResponse
	bytes, _ := json.Marshal(res)
	if err := json.Unmarshal(bytes, &anchor); err != nil {
		return nil, err
	}

	return &anchor, nil
}

// scheduler executes periodically the checkStack method
func (s *service) scheduler(done chan bool) {

	ticker := time.NewTicker(time.Duration(s.params.WriteInterval) * time.Second)
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
