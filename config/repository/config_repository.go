package repository

import "github.com/enchainte/enchainte-sdk-go/config/entity"

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
