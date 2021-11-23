package entity

import (
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"math"
)

type Proof struct {
	Leaves []string `json:"leaves"`
	Nodes  []string `json:"nodes"`
	Depth  string   `json:"depth"`
	Bitmap string   `json:"bitmap"`
}

func NewProof(leaves, nodes []string, depth, bitmap string) Proof {
	return Proof{
		Leaves: leaves,
		Nodes:  nodes,
		Depth:  depth,
		Bitmap: bitmap,
	}
}

func IsValid(proof Proof) bool {
	if isType(proof) {
		for _, l := range proof.Leaves {
			if !shared.IsHex(l) || len(l) != 64 {
				return false
			}
		}
		for _, n := range proof.Nodes {
			if !shared.IsHex(n) || len(n) != 64 {
				return false
			}
		}
		nElements := len(proof.Leaves) + len(proof.Nodes)
		if len(proof.Depth) != nElements*4 && shared.IsHex(proof.Depth) {
			return false
		}
		if len(proof.Depth) != nElements*4 {
			return false
		}

		if math.Floor(float64(len(proof.Bitmap)/2)) < math.Floor(float64((nElements+8-(nElements%8))/8)) {
			return false
		}
		return true
	}
	return false
}

func isType(t interface{}) bool {
	switch t.(type) {
	case Proof:
		return true
	default:
		return false
	}
}
