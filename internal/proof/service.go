package proof

import (
	"encoding/json"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
)

type service struct {
	http http.Client
}

type Service interface {
	Proof(hashes [][]byte) (*Proof, error)
	Verify(leaves, nodes []string, depth, bitmap string) bool
}

func NewService(http http.Client) Service {
	return &service{http}
}

func (s *service) Proof(hashesBytes [][]byte) (*Proof, error) {
	var hashes []string
	for _, h := range  hashesBytes {
		m, err := message.New(h)
		if err != nil {
			return nil, err
		}
		hashes = append(hashes, m.Hash())
	}

	body := ApiProofRequestBody{
		Messages: hashes,
		Client:   "",
	}

	// TODO sort necessary?
	bytes, err := s.http.PostRequest("/v1/messages/proof", body)
	if err != nil {
		return nil, err
	}

	var proof Proof
	if err := json.Unmarshal(bytes, &proof); err != nil {
		return nil, err
	}
	return &proof, nil
}

func (s *service) Verify(leaves, nodes []string, depth, bitmap string) bool {
	panic("implement me")
}
