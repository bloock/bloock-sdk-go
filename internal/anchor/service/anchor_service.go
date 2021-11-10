package service

import (
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/repository"
)

type AnchorService struct {
	anchorRepository repository.AnchorRepository
	configService service.ConfigService
}

func NewAnchorService(ar repository.AnchorRepository, conf service.ConfigService) AnchorService {
	return AnchorService{
		anchorRepository: ar,
		configService: conf,
	}
}

func(a AnchorService) GetAnchor(anchorId int) entity.Anchor {
	return entity.Anchor{}
}

func(a AnchorService) WaitAnchor(anchorId int, timeout int) entity.Anchor {
	return entity.Anchor{}
}
