package bloock

import (
	anchorEntity "github.com/bloock/bloock-sdk-go/internal/anchor/entity"
	configNetwork "github.com/bloock/bloock-sdk-go/internal/config/entity"
	"github.com/bloock/bloock-sdk-go/internal/record/entity"
)

/*
NewNetworkParams
function that returns a NetworkParams object. You can set another network
Parameters:
	{void}
Returns:
	{NetworkParams} NetworkParams object return, where you can set the Network field.
 */
func NewNetworkParams() configNetwork.NetworkParams {
	return configNetwork.NetworkParams{}
}

/*
NewAnchorParams
function that returns an AnchorParam object. You can set another timeout
Parameters:
	{void}
Returns:
	{AnchorParams} AnchorParams object return, where you can set the Timeout field
 */
func NewAnchorParams() anchorEntity.AnchorParams {
	return anchorEntity.AnchorParams{}
}

/*
Record
Call Record to get a RecordEntity type
 */
type Record = entity.RecordEntity

type networks struct {
	BloockChain     string
	EthereumRinkeby string
	EthereumMainnet string
}

/*
ListOfNetworks
Gets a networks object with all different networks available to use.
Parameters:
	{void}
Returns:
	{networks} returns all networks available to use.
 */
func ListOfNetworks() networks {
	return networks{
		BloockChain:     configNetwork.BloockChain,
		EthereumRinkeby: configNetwork.EthereumRinkeby,
		EthereumMainnet: configNetwork.EthereumMainnet,
	}
}
