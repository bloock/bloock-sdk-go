package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
	"github.com/enchainte/enchainte-sdk-go/internal/proof"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	"time"
)

type Service interface {
	Write(hash []byte) error
	Search(hash [][]byte) (*Receipts, error)
	Verify(hashes [][]byte) (bool, error)
	Wait(hashes [][]byte) (*Receipts, error)
}

const period = 2 // seconds

var (
	messagesStack []Message
)

type service struct {
	channel chan SendResponse
	apiKey string
	http http.Client
	constants config.Constants
	proof proof.Service
}

func NewService(ch chan SendResponse, apiKey string, http http.Client, constants config.Constants, proof proof.Service) Service {
	return &service{ch,apiKey,http, constants, proof}
}

func (s *service) init() {
	go s.scheduler(time.Second * period, s.channel)
}

type Receipts struct {
	Messages []Receipt `json:"messages"`
}
//type Receipts []Receipt `json:"messages"`

type Receipt struct {
	Message string `json:"Message"`
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
	body WriteResponse
	err  error
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

	resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s%s",s.constants.Api.Host, s.constants.Api.Endpoints.MessageFetch), nil, body)
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

func (s *service) Verify(hashes [][]byte) (bool, error) {

	p, err := s.proof.Proof(hashes)
	if err != nil {
		return false, err
	}

	return s.proof.Verify(p.Leaves, p.Nodes, string(p.Depth), string(p.Bitmap))
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

		if len(receipts.Messages) <= len(hashes) {
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

		// TODO change by time and coef define in config
		time.Sleep(time.Duration(3+(attempts*1)) * time.Second)

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
	resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s%s",s.constants.Api.Host, s.constants.Api.Endpoints.MessageWrite), nil, body)
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
func (s *service) scheduler(period time.Duration, ch chan SendResponse) {
	for {
		c := time.Tick(period)
		for range c {
			s.checkStack(ch)
		}
	}
}

// checkStack executes the send method when the message stack is not empty and sends the result to the channel
func (s *service) checkStack(ch chan SendResponse) {
	if messagesStack == nil {
		fmt.Println("Empty stack... leaving")
		fmt.Println()
		return
	}

	fmt.Println("Sending message stack...")
	fmt.Println()
	resp, err := s.send()

	messagesStack = nil
	ch <- SendResponse{
		body: *resp,
		err:  err,
	}
	return
}
