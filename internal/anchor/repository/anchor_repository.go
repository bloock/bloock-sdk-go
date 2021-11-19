package repository

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor/entity/dto"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure"
)

type AnchorRepository struct {
	httpClient infrastructure.HttpClient
	configService service.ConfigurerService
}

func NewAnchorRepository(httpClient infrastructure.HttpClient, configService service.ConfigurerService) AnchorRepository{
	return AnchorRepository{
		httpClient: httpClient,
		configService: configService,
	}
}

func(a AnchorRepository) GetAnchor(anchor int) (entity.Anchor, error) {
	url := fmt.Sprintf("%s/core/anchor/%d", a.configService.GetApiBaseUrl(), anchor)
	response, err := a.httpClient.Get(url, nil)
	if err != nil {
		return entity.Anchor{}, err
	}

	var r = dto.AnchorRetrieveResponse{}
	if err = json.Unmarshal(response, &r); err != nil {
		return entity.Anchor{}, fmt.Errorf("getAnchor.unmarshall2: %s", err)
	}

	anchorResponse := entity.NewAnchor(r.AnchorId, r.BlockRoots, r.Networks, r.Root, r.Status)

	return anchorResponse, nil
}
