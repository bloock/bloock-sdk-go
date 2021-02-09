package proof

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
	"github.com/enchainte/enchainte-sdk-go/internal/message"
	"github.com/enchainte/enchainte-sdk-go/pkg/blockchain"
	"github.com/enchainte/enchainte-sdk-go/pkg/crypto"
	"github.com/enchainte/enchainte-sdk-go/pkg/http"
)

type service struct {
	apiKey string
	http   http.Client
	constants config.Constants
	hasher crypto.Hasher
	bc blockchain.Client
}

type Service interface {
	Proof(hashes [][]byte) (*Proof, error)
	Verify(leaves, nodes []string, depth, bitmap string) (bool, error)
	CalculateRoot(proof *Proof) (string, error)
}

func NewService(apiKey string, http http.Client, constants config.Constants, hasher crypto.Hasher, bc blockchain.Client) Service {
	return &service{apiKey,http,constants,hasher, bc}
}

func (s *service) Proof(hashesBytes [][]byte) (*Proof, error) {
	var hashes []string
	for _, h := range hashesBytes {
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
	resp, err := s.http.Request(s.apiKey, "POST", fmt.Sprintf("%s%s",s.constants.Api.Host, s.constants.Api.Endpoints.MessageProof),nil,  body)
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

	var proof Proof
	bytes, _ := json.Marshal(res)
	if err := json.Unmarshal(bytes, &proof); err != nil {
		return nil, err
	}

	return &proof, nil
}

func (s *service) Verify(leaves, nodes []string, depth, bitmap string) (bool, error) {
	proof, err := New(leaves, nodes, depth, bitmap)
	if err != nil {
		return false, err
	}

	var root string
	root, err = s.CalculateRoot(proof)
	if err != nil {
		return false, err
	}

	return s.bc.ValidateRoot(root)
}

func (s *service) CalculateRoot(proof *Proof) (string, error) {

	lsb := proof.Leaves.Bytes()
	nsb := proof.Nodes.Bytes()
	db := proof.Depth.Bytes()
	bb := proof.Bitmap.Bytes()

	var leavesIt, nodesIt int
	var actNode []byte

	type pair struct {
		Node  []byte
		Depth byte
	}
	var stack []pair

	for nodesIt < len(nsb) || leavesIt < len(lsb) {
		actDepth := db[leavesIt+nodesIt]

		if (bb[(nodesIt+leavesIt)/8] & (1 << (7 - ((nodesIt + leavesIt) % 8)))) > 0 {
			actNode = nsb[nodesIt]
			nodesIt++
		} else {
			actNode = lsb[leavesIt]
			leavesIt++
		}

		for len(stack) > 0 && stack[len(stack)-1].Depth == actDepth {
			if len(stack) == 0 {
				return "", errors.New("error: stack got empty before capturing its value")
			}
			// array pop
			lastNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var err error
			actNode, err = s.merge(lastNode.Node, actNode)
			if err != nil {
				return "", err
			}

			actDepth--
		}
		stack = append(stack, pair{actNode, actDepth})
	}
	return hex.EncodeToString(stack[0].Node), nil
}

func (s *service) merge(hashes ...[]byte) ([]byte, error) {
	var concat []byte
	for _, h := range hashes {
		concat = append(concat, h...)
	}

	newHash, err := s.hasher.Hash(concat)
	if err != nil {
		return nil, err
	}

	return newHash, nil
}
