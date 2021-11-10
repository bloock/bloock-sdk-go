package repository

import (
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
)

type AnchorRepository struct {
	httpClient HttpClient
	configService service.ConfigService
}

func NewAnchorRepository(httpClient HttpClient, configService service.ConfigService) AnchorRepository{
	return AnchorRepository{
		httpClient: httpClient,
		configService: configService,
	}
}

func(a AnchorRepository) GetAnchor(anchor int) entity.Anchor {
	url := fmt.Sprintf("%s/core/anchor/%d", a.configService.GetApiBaseUrl(), anchor)
	response := "get anchor from http"
	return entity.Anchor{}
}
