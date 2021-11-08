mocks:
	mockgen -source=config/service/configurer_service.go -destination config/mockconfig/mock_config_service.go -package=mockconfig
	mockgen -source=config/repository/configurer_repository.go -destination config/mockconfig/mock_config_repository.go -package=mockconfig
