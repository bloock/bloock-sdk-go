package blockchain

import (
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/blockchain/contracts"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type Web3Client struct {
	configService service.ConfigurerService
}

func NewWeb3(config service.ConfigurerService) Web3Client {
	return Web3Client{
		configService: config,
	}
}

func(w Web3Client) ValidateRoot(network string, root string) (int64, error) {
	log.Println(root)
	config := w.configService.GetNetworkConfiguration(network)

	client, err := ethclient.Dial(config.HttpProvider)
	if err != nil {
		return -1, fmt.Errorf("validateRoot.Dial: %s", err)
	}
	address := common.HexToAddress(config.ContractAddress)
	instance, err := contracts.NewContracts(address, client)
	if err != nil {
		return -1, fmt.Errorf("validateRoot.NewContract: %s", err)
	}

	rootByte, err := shared.HexToBytes32(root)
	if err != nil {
		return -1, err
	}
	timestamp, err := instance.GetState(nil, rootByte)
	if err != nil {
		return -1, fmt.Errorf("validateRoot.GetState: %s", err)
	}

	if timestamp == nil {
		return -1, errors.New("returned timestamp is not valid")
	}

	return timestamp.Int64(), nil
}


