package service

import "github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"

type AnchorerService interface {
	GetAnchor(anchorId int) entity.Anchor
	WaitAnchor(anchorId int, timeout int) entity.Anchor
}
