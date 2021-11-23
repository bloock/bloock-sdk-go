mocks:
	mockgen -source=internal/config/service/configurer_service.go -destination internal/config/mockconfig/mock_config_service.go -package=mockconfig
	mockgen -source=internal/config/repository/configurer_repository.go -destination internal/config/mockconfig/mock_config_repository.go -package=mockconfig
	mockgen -source=internal/anchor/repository/anchorer_repository.go -destination internal/anchor/mockanchor/mock_anchor_repository.go -package=mockanchor
	mockgen -source=internal/infrastructure/http_client.go -destination internal/infrastructure/http/mockhttp/mocks_http.go -package=mockhttp
	mockgen -source=internal/record/repository/recorder_repository.go -destination internal/record/mockrecord/mock_record_repository.go -package=mockrecord
	mockgen -source=internal/infrastructure/blockchain_client.go -destination internal/infrastructure/blockchain/mockblockchain/mocks_blockchain.go -package=mockblockchain
