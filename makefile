mocks:
	mockgen -source=config/service/configurer_service.go -destination config/mockconfig/mock_config_service.go -package=mockconfig
	mockgen -source=config/repository/configurer_repository.go -destination config/mockconfig/mock_config_repository.go -package=mockconfig
	mockgen -source=internal/infrastructure/http_client.go -destination internal/infrastructure/http/mockhttp/mocks_http.go -package=mockhttp
	mockgen -source=internal/record/repository/recorder_repository.go -destination internal/record/mockrecord/mock_record_repository.go -package=mockrecord