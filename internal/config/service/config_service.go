package service

import (
	"github.com/enchainte/enchainte-sdk-go/internal/config/entity"
	"github.com/enchainte/enchainte-sdk-go/internal/config/repository"
)

type ConfigService struct {
	configRepository repository.ConfigurerRepository
}

func NewConfigService(configRepo repository.ConfigurerRepository) ConfigService {
	return ConfigService{
		configRepository: configRepo,
	}
}

func (c ConfigService) GetNetworkConfiguration(network string) entity.NetworkConfiguration {
	return c.configRepository.GetNetworkConfiguration(network)
}

func (c ConfigService) GetApiBaseUrl() string {
	return c.configRepository.GetConfiguration().Host
}

func (c ConfigService) GetConfiguration() entity.Configuration {
	return c.configRepository.GetConfiguration()
}

func (c ConfigService) SetNetworkConfiguration(network string, config entity.NetworkConfiguration) {
	c.configRepository.SetNetworkConfiguration(network, config)
}

func (c ConfigService) SetApiHost(host string) {
	c.configRepository.SetApiHost(host)
}
