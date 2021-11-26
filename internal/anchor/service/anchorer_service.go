package service

import "github.com/bloock/bloock-sdk-go/internal/anchor/entity"

type AnchorerService interface {
	GetAnchor(anchorId int) (entity.Anchor, error)
	WaitAnchor(anchorId int, params entity.AnchorParams) (entity.Anchor, error)
}
