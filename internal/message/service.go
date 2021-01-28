package message

import (
	"encoding/json"
	"github.com/enchainte/enchainte-sdk-go/internal/pkg/http"
	"time"
)

type Service interface {
	Send(hash []byte) bool
	Search(hash [][]byte) ([]Receipt, error)
	Verify(hashes [][]byte) bool
	Wait(hashes [][]byte) ([]Receipt, error)
}

type service struct {
	http http.Client
}

func NewService(http http.Client) Service {
	return &service{http}
}

type Receipt struct {
	Root string
	Message string
	TxHash string
	Status string
	Error string
}

func (s *service) Send(message []byte) bool {
	panic("implement me")
}

func (s *service) Search(hashes [][]byte) ([]Receipt, error) {
	var ms []message
	for _, h := range hashes {
		m, err := New(h)
		if err != nil {
			return nil, err
		}
		ms = append(ms, *m)
	}

	body := struct {
		Messages []string `json:"messages"`
		Client string `json:"client"`
	}{}

	// TODO add endpoint to config
	resp, err := s.http.PostRequest("/v1/messages/fetch", body)
	if err != nil {
		return nil, err
	}

	var rcpt []Receipt
	if err := json.Unmarshal(resp, &rcpt); err != nil {
		return nil, err
	}

	return rcpt, nil
}

func (s *service) Verify(hashes [][]byte) bool {
	panic("implement me")
}

func (s *service) Wait(hashes [][]byte) ([]Receipt, error) {
	var complete bool
	var attempts int

	var receipts []Receipt
	for !complete {
		receipts, err := s.Search(hashes)
		if err != nil {
			return nil, err
		}

		if len(receipts) <= len(hashes) {
			continue
		}

		for _, r := range receipts {
			if r.Status == "success" || r.Status == "error"{
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
		time.Sleep(time.Duration(3 + (attempts * 1)) * time.Second)

		attempts++
	}

	return receipts, nil
}
