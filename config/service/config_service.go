package service

import (
	"github.com/enchainte/enchainte-sdk-go/config/entity"
	"github.com/enchainte/enchainte-sdk-go/config/repository"
)

type ConfigService struct {
	configRepository repository.ConfigurerRepository
}

func NewConfigService(configRepo repository.ConfigurerRepository) ConfigService {
	return ConfigService{
		configRepository: configRepo,
	}
}

func(c ConfigService) GetNetworkConfiguration(network string) entity.NetworkConfiguration {
	return c.configRepository.GetNetworkConfiguration(network)
}

func(c ConfigService) GetApiBaseUrl() string {
	return c.configRepository.GetConfiguration().Host
}
