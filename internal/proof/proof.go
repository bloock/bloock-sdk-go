package proof

import (
	"encoding/hex"
	"errors"
)

const (
	length = 64
)

var (
	ErrInvalidLeave  = errors.New("invalid leave provided")
	ErrInvalidNode   = errors.New("invalid node provided")
	ErrInvalidDepth  = errors.New("invalid depth provided")
	ErrInvalidBitmap = errors.New("invalid bitmap provided")
)

type Proof struct {
	Leaves []string
	Nodes  []string
	Depth  string
	Bitmap string
}

func New(leaves, nodes []string, depth, bitmap string) (*Proof, error) {
	p := Proof{
		Leaves: leaves,
		Nodes:  nodes,
		Depth:  depth,
		Bitmap: bitmap,
	}
	if err := p.validate(); err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Proof) validate() error {
	for _, l := range p.Leaves {
		if !isHex(l) || len(l) != length {
			return ErrInvalidLeave
		}
	}
	for _, n := range p.Nodes {
		if !isHex(n) || len(n) != length {
			return ErrInvalidNode
		}
	}
	if !isHex(p.Depth) {
		return ErrInvalidDepth
	}
	if !isHex(p.Bitmap) {
		return ErrInvalidBitmap
	}

	return nil
}

func isHex(s string) bool {
	if _, err := hex.DecodeString(s); err != nil {
		return false
	}
	return true
}

