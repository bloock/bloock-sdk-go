package service

import (
	"github.com/enchainte/enchainte-sdk-go/config/entity"
	"github.com/enchainte/enchainte-sdk-go/config/repository"
)

type ConfigServiceImpl struct {
	configRepository repository.ConfigRepository
}

func NewConfigServiceImpl(configRepo repository.ConfigRepository) ConfigServiceImpl {
	return ConfigServiceImpl{
		configRepository: configRepo,
	}
}

func(c ConfigServiceImpl) GetNetworkConfiguration(network entity.Network) entity.NetworkConfiguration {
	return c.configRepository.GetNetworkConfiguration(network)
}
