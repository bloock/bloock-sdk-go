package repository

import (
	"github.com/enchainte/enchainte-sdk-go/internal/config/entity"
)

type ConfigRepository struct {
	configData ConfigData
}

func NewConfigRepository(configData ConfigData) ConfigRepository {
	return ConfigRepository{
		configData: configData,
	}
}

func(c ConfigRepository) GetNetworkConfiguration(network string) entity.NetworkConfiguration {
	return c.configData.getNetworkConfiguration(network)
}

func(c ConfigRepository) GetConfiguration() entity.Configuration {
	return c.configData.getConfiguration()
}

func(c *ConfigRepository) SetNetworkConfiguration(network string, config entity.NetworkConfiguration) {
	c.configData.setNetworkConfiguration(network, config)
}

func(c *ConfigRepository) SetApiHost(host string) {
	c.configData.setApiHost(host)
}
