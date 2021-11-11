package repository

import "github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"

//go:generate mockgen -source=internal/anchor/repository/anchorer_repository.go -destination internal/anchor/mockanchor/mock_anchor_repository.go -package=mockanchor
type AnchorerRepository interface {
	GetAnchor(anchor int) (entity.Anchor, error)
}
