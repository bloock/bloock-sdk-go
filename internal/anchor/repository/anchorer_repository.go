package repository

import "github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"

type AnchorerRepository interface {
	GetAnchor(anchor int) entity.Anchor
}
