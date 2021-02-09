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

type Leaves []string
type Nodes []string
type Depth string
type Bitmap string

type Proof struct {
	Leaves Leaves
	Nodes  Nodes
	Depth  Depth
	Bitmap Bitmap
}

func New(leaves, nodes []string, depth, bitmap string) (*Proof, error) {
	p := Proof{
		Leaves: leaves,
		Nodes:  nodes,
		Depth:  Depth(depth),
		Bitmap: Bitmap(bitmap),
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
	if !isHex(string(p.Depth)) {
		return ErrInvalidDepth
	}
	if !isHex(string(p.Bitmap)) {
		return ErrInvalidBitmap
	}

	return nil
}

func (ls *Leaves) Bytes() [][]byte  {
	var leaves [][]byte
	for _, l := range *ls {
		bytes, _ := hexToBytes(string(l))
		leaves = append(leaves, bytes)
	}
	return leaves
}

func (ns *Nodes) Bytes() [][]byte  {
	var nodes [][]byte
	for _, n := range *ns {
		bytes, _ := hexToBytes(n)
		nodes = append(nodes, bytes)
	}
	return nodes
}

func (d *Depth) Bytes() []byte  {
	bytes, _ := hexToBytes(string(*d))

	return bytes
}

func (b *Bitmap) Bytes() []byte {
	bytes, _ := hexToBytes(string(*b))

	return bytes
}

func isHex(s string) bool {
	if _, err := hex.DecodeString(s); err != nil {
		return false
	}
	return true
}

func hexToBytes(hexStr string) ([]byte, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
