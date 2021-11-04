package blockchain

import (
	"errors"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config/entity"
	"github.com/enchainte/enchainte-sdk-go/config/service"
	"github.com/enchainte/enchainte-sdk-go/internal/infrastructure/blockchain/contracts"
	"github.com/enchainte/enchainte-sdk-go/internal/shared"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Web3Client struct {
	configService service.ConfigService
}

func NewWeb3(config service.ConfigService) Web3Client {
	return Web3Client{
		configService: config,
	}
}

func(w Web3Client) ValidateRoot(network entity.Network, root string) (*big.Int, error) {
	config := w.configService.GetNetworkConfiguration(network)

	client, err := ethclient.Dial(config.HttpProvider)
	if err != nil {
		return big.NewInt(-1), fmt.Errorf("validateRoot.Dial: %s", err)
	}
	address := common.HexToAddress(config.ContractAddress)
	instance, err := contracts.NewContracts(address, client)
	if err != nil {
		return big.NewInt(-1), fmt.Errorf("validateRoot.NewContract: %s", err)
	}

	rootByte, err := shared.HexToBytes32(root)
	if err != nil {
		return big.NewInt(-1), err
	}
	timestamp, err := instance.GetState(nil, rootByte)
	if err != nil {
		return big.NewInt(-1), fmt.Errorf("validateRoot.GetState: %s", err)
	}

	if timestamp == nil {
		return big.NewInt(-1), errors.New("returned timestamp is not valid")
	}

	return timestamp, nil
}


