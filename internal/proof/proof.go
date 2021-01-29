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
	leaves []string
	nodes  []string
	depth  string
	bitmap string
}

func New(leaves, nodes []string, depth, bitmap string) (*Proof, error) {
	p := Proof{
		leaves: leaves,
		nodes:  nodes,
		depth:  depth,
		bitmap: bitmap,
	}
	if err := p.validate(); err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Proof) validate() error {
	for _, l := range p.leaves {
		if !isHex(l) || len(l) != length {
			return ErrInvalidLeave
		}
	}
	for _, n := range p.nodes {
		if !isHex(n) || len(n) != length {
			return ErrInvalidNode
		}
	}
	if !isHex(p.depth) {
		return ErrInvalidDepth
	}
	if !isHex(p.bitmap) {
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

func (p *Proof) Leaves() []string {
	return p.leaves
}

func (p *Proof) Nodes() []string {
	return p.nodes
}

func (p *Proof) Depth() string {
	return p.depth
}

func (p *Proof) Bitmap() string {
	return p.bitmap
}
