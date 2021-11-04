package service

import "github.com/enchainte/enchainte-sdk-go/config/entity"

//go:generate mockgen -source=config/service/config_service.go -destination config/mockconfig/mocks_config_service.go -package=mockconfig
type ConfigService interface {
	GetNetworkConfiguration(network entity.Network) entity.NetworkConfiguration
}
