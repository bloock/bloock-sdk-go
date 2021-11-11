mocks:
	mockgen -source=config/service/configurer_service.go -destination config/mockconfig/mock_config_service.go -package=mockconfig
	mockgen -source=config/repository/configurer_repository.go -destination config/mockconfig/mock_config_repository.go -package=mockconfig
	mockgen -source=internal/anchor/service/anchorer_service.go -destination internal/anchor/mockanchor/mock_anchor_service.go -package=mockanchor
	mockgen -source=internal/anchor/repository/anchorer_repository.go -destination internal/anchor/mockanchor/mock_anchor_repository.go -package=mockanchor