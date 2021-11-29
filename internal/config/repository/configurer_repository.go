package repository

import (
	"github.com/bloock/bloock-sdk-go/internal/config/entity"
)

//go:generate mockgen -source=internal/config/repository/configurer_repository.go -destination internal/config/mockconfig/mock_config_repository.go -package=mockconfig
type ConfigurerRepository interface {
	GetNetworkConfiguration(network string) entity.NetworkConfiguration
	GetConfiguration() entity.Configuration
	SetNetworkConfiguration(network string, config entity.NetworkConfiguration)
	SetApiHost(host string)
}
