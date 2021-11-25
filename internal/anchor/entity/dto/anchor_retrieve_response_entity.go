package dto

import "github.com/bloock/bloock-sdk-go/internal/anchor/entity"

type AnchorRetrieveResponse struct {
	AnchorId int `json:"anchor_id"`
	BlockRoots []string `json:"block_roots"`
	Networks []entity.Network `json:"networks"`
	Root string `json:"root"`
	Status string `json:"status"`
}
