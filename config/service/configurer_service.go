package service

import "github.com/enchainte/enchainte-sdk-go/config/entity"

//go:generate mockgen -source=config/service/configurer_service.go -destination config/mockconfig/mocks_config_service.go -package=mockconfig
type ConfigurerService interface {
	GetNetworkConfiguration(network string) entity.NetworkConfiguration
	GetApiBaseUrl() string
}
