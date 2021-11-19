package service

import "github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"

//go:generate mockgen -source=internal/anchor/service/anchorer_service.go -destination internal/anchor/mockanchor/mock_anchor_service.go -package=mockanchor
type AnchorerService interface {
	GetAnchor(anchorId int) (entity.Anchor, error)
	WaitAnchor(anchorId int, timeout int) (entity.Anchor, error)
}
