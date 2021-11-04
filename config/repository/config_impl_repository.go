package repository

import "github.com/enchainte/enchainte-sdk-go/config/entity"

type ConfigRepositoryImpl struct {
	configData ConfigData
}

func NewConfigRepositoryImpl(configData ConfigData) ConfigRepositoryImpl {
	return ConfigRepositoryImpl{
		configData: configData,
	}
}

func(c ConfigRepositoryImpl) GetNetworkConfiguration(network entity.Network) entity.NetworkConfiguration {
	return c.configData.getNetworkConfiguration(network)
}
