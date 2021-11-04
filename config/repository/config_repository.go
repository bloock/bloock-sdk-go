package repository

import "github.com/enchainte/enchainte-sdk-go/config/entity"

//go:generate mockgen -source=config/repository/config_repository.go -destination config/mockconfig/mock_config_repository.go -package=mockconfig
type ConfigRepository interface {
	GetNetworkConfiguration(network entity.Network) entity.NetworkConfiguration
}
