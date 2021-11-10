package repository

import "github.com/enchainte/enchainte-sdk-go/config/entity"

//go:generate mockgen -source=config/repository/configurer_repository.go -destination config/mockconfig/mock_config_repository.go -package=mockconfig
type ConfigurerRepository interface {
	GetNetworkConfiguration(network string) entity.NetworkConfiguration
	GetConfiguration() entity.Configuration
}
