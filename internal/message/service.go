package message

import (
	"encoding/json"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
	"time"
)

type Service interface {
	Send(hash []byte) bool
	Search(hash [][]byte) (*Receipts, error)
	Verify(hashes [][]byte) bool
	Wait(hashes [][]byte) (*Receipts, error)
}

type service struct {
	http http.Client
}

func NewService(http http.Client) Service {
	return &service{http}
}

type Receipts struct {
	Messages []Receipt `json:"messages"`
}
//type Receipts []Receipt `json:"messages"`

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

func (s *service) Send(message []byte) bool {
	panic("implement me")
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

	body := ApiFetchRequestBody{
		Messages: hashes,
		// TODO
		Client: "",
	}

	// TODO add endpoint to config
	resp, err := s.http.PostRequest("/v1/messages/fetch", body)
	if err != nil {
		return nil, err
	}

	var rs Receipts
	if err := json.Unmarshal(resp, &rs); err != nil {
		return nil, err
	}

	return &rs, nil
}

func (s *service) Verify(hashes [][]byte) bool {
	panic("implement me")
}

func (s *service) Wait(hashes [][]byte) (*Receipts, error) {
	var complete bool
	var attempts int

	var receipts Receipts
	for !complete {
		receipts, err := s.Search(hashes)
		if err != nil {
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
