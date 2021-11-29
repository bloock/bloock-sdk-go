package service

import (
	"github.com/bloock/bloock-sdk-go/internal/config/entity"
)

//go:generate mockgen -source=internal/config/service/configurer_service.go -destination internal/config/mockconfig/mocks_config_service.go -package=mockconfig
type ConfigurerService interface {
	GetNetworkConfiguration(network string) entity.NetworkConfiguration
	GetApiBaseUrl() string
	GetConfiguration() entity.Configuration
	SetNetworkConfiguration(network string, config entity.NetworkConfiguration)
	SetApiHost(host string)
}
