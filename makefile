mocks:
	mockgen -source=config/service/config_service.go -destination config/mockconfig/mock_config_service.go -package=mockconfig
	mockgen -source=config/repository/config_repository.go -destination config/mockconfig/mock_config_repository.go -package=mockconfig
