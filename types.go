package bloock

import (
	anchorEntity "github.com/bloock/bloock-sdk-go/internal/anchor/entity"
	configNetwork "github.com/bloock/bloock-sdk-go/internal/config/entity"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
)

func NewNetworkParams() configNetwork.NetworkParams {
	return configNetwork.NetworkParams{}
}

func NewAnchorParams() anchorEntity.AnchorParams {
	return anchorEntity.AnchorParams{}
}

type Record = entity.RecordEntity

var BloockChain = configNetwork.BloockChain
var EthereumRinkeby = configNetwork.EthereumRinkeby
var EthereumMainnet = configNetwork.EthereumMainnet