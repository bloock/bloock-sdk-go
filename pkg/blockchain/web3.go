package blockchain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/enchainte/enchainte-sdk-go/config"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	ValidateRoot(root string) (bool, error)
}

type web3Client struct {
	config config.Constants
	params cloud.SdkParams
}

var conn *EthereumConnection

func Web3(constants config.Constants, params cloud.SdkParams) Client {
	conn = SetupEthereumBlockchain(params)
	return &web3Client{constants, params}
}

func (w *web3Client) ValidateRoot(root string) (bool, error) {

	rootBytes, err := hexToBytes32(root)
	if err != nil {
		return false, err
	}

	return conn.Contracts.GetCheckpoint( nil, rootBytes)
}

type EthereumConnection struct {
	Client     *ethclient.Client
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
	Contracts  *Contract
}

// SetupBC - Sets up a Blockchain Connection
func SetupEthereumBlockchain(params cloud.SdkParams) *EthereumConnection {

	client, err := ethclient.Dial(params.Web3Provider)
	if err != nil {
		panic(err.Error())
	}

	address := common.HexToAddress(params.SmartContractAddress)
	instance, err := NewContract(address, client)
	if err != nil {
		panic(err.Error())
	}

	return &EthereumConnection{
		Client:     client,
		Contracts:  instance,
	}
}

func hexToBytes32(hexStr string) ([32]byte, error) {
	var bytes32 [32]byte
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return bytes32, err
	}
	copy(bytes32[:], bytes)
	return bytes32, nil
}

