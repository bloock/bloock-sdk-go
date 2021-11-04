// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"role_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"state_manager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATE_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"state_root\",\"type\":\"bytes32\"}],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"state_root\",\"type\":\"bytes32\"}],\"name\":\"isStatePresent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"state_root\",\"type\":\"bytes32\"}],\"name\":\"updateState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"content\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"bitmap\",\"type\":\"bytes\"},{\"internalType\":\"uint32[]\",\"name\":\"depths\",\"type\":\"uint32[]\"}],\"name\":\"verifyInclusionProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000efc38038062000efc833981016040819052620000349162000159565b6200004160008362000075565b6200006d7f7c0c08811839d3a8bfad3f26fd05feea7daf5d75bf9d6f3fe140cd8f62b7af388262000075565b505062000190565b62000081828262000085565b5050565b6200009182826200010f565b62000081576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620000cb62000138565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b3390565b80516001600160a01b03811681146200015457600080fd5b919050565b600080604083850312156200016c578182fd5b62000177836200013c565b915062000187602084016200013c565b90509250929050565b610d5c80620001a06000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80632f2ff15d116100715780632f2ff15d1461014357806336568abe146101585780638ea59e1d1461016b57806391d148541461017e578063a217fddf14610191578063d547741f14610199576100b4565b806301ffc9a7146100b957806309648a9d146100e25780630aeb5fa21461010257806313bf2c97146101155780632450a4a214610128578063248a9ca314610130575b600080fd5b6100cc6100c736600461098f565b6101ac565b6040516100d991906109e2565b60405180910390f35b6100f56100f036600461093d565b6101d7565b6040516100d991906109ed565b6100f5610110366004610853565b6101e9565b6100cc61012336600461093d565b610549565b6100f561055d565b6100f561013e36600461093d565b610581565b610156610151366004610955565b610596565b005b610156610166366004610955565b6105dd565b61015661017936600461093d565b61061f565b6100cc61018c366004610955565b610678565b6100f56106a1565b6101566101a7366004610955565b6106a6565b60006001600160e01b03198216637965db0b60e01b14806101d157506101d1826106ce565b92915050565b60009081526001602052604090205490565b60008080806101f8898c610b3d565b67ffffffffffffffff81111561021e57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561025757816020015b6102446107f3565b81526020019060019003908161023c5790505b50905060005b8983108061026a57508b84105b156104f7576000878761027d8787610b3d565b81811061029a57634e487b7160e01b600052603260045260246000fd5b90506020020160208101906102af91906109be565b9050600060086102bf8787610b3d565b6102c99190610ce6565b6102d4906007610caa565b6102df906002610baf565b6102ea906001610c81565b60f81b8b8b60086102fb8a8a610b3d565b6103059190610b55565b81811061032257634e487b7160e01b600052603260045260246000fd5b6001600160f81b0319920135929092161615905061037f57610345856001610b3d565b94508c8c610354600188610caa565b81811061037157634e487b7160e01b600052603260045260246000fd5b9050602002013590506103c0565b61038a866001610b3d565b95508e8e610399600189610caa565b8181106103b657634e487b7160e01b600052603260045260246000fd5b9050602002013590505b821580159061040f575063ffffffff8216846103dd600186610caa565b815181106103fb57634e487b7160e01b600052603260045260246000fd5b60200260200101516020015163ffffffff16145b1561049d57600084610422600186610caa565b8151811061044057634e487b7160e01b600052603260045260246000fd5b602002602001015160000151905060018461045b9190610caa565b935080826040516020016104709291906109f6565b6040516020818303038152906040528051906020012091506001836104959190610cc1565b9250506103c0565b60405180604001604052808281526020018363ffffffff168152508484815181106104d857634e487b7160e01b600052603260045260246000fd5b60209081029190910101526104ee836001610b3d565b9250505061025d565b600160008360008151811061051c57634e487b7160e01b600052603260045260246000fd5b60200260200101516000015181526020019081526020016000205494505050505098975050505050505050565b600090815260016020526040902054151590565b7f7c0c08811839d3a8bfad3f26fd05feea7daf5d75bf9d6f3fe140cd8f62b7af3881565b60009081526020819052604090206001015490565b6105aa6105a283610581565b61018c6106e7565b6105cf5760405162461bcd60e51b81526004016105c690610a04565b60405180910390fd5b6105d982826106eb565b5050565b6105e56106e7565b6001600160a01b0316816001600160a01b0316146106155760405162461bcd60e51b81526004016105c690610aee565b6105d98282610770565b6106497f7c0c08811839d3a8bfad3f26fd05feea7daf5d75bf9d6f3fe140cd8f62b7af3833610678565b6106655760405162461bcd60e51b81526004016105c690610a53565b6000908152600160205260409020429055565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b600081565b6106b26105a283610581565b6106155760405162461bcd60e51b81526004016105c690610a9e565b6001600160e01b031981166301ffc9a760e01b14919050565b3390565b6106f58282610678565b6105d9576000828152602081815260408083206001600160a01b03851684529091529020805460ff1916600117905561072c6106e7565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61077a8282610678565b156105d9576000828152602081815260408083206001600160a01b03851684529091529020805460ff191690556107af6106e7565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b604080518082019091526000808252602082015290565b60008083601f84011261081b578182fd5b50813567ffffffffffffffff811115610832578182fd5b602083019150836020808302850101111561084c57600080fd5b9250929050565b6000806000806000806000806080898b03121561086e578384fd5b883567ffffffffffffffff80821115610885578586fd5b6108918c838d0161080a565b909a50985060208b01359150808211156108a9578586fd5b6108b58c838d0161080a565b909850965060408b01359150808211156108cd578586fd5b818b0191508b601f8301126108e0578586fd5b8135818111156108ee578687fd5b8c60208285010111156108ff578687fd5b6020830196508095505060608b013591508082111561091c578384fd5b506109298b828c0161080a565b999c989b5096995094979396929594505050565b60006020828403121561094e578081fd5b5035919050565b60008060408385031215610967578182fd5b8235915060208301356001600160a01b0381168114610984578182fd5b809150509250929050565b6000602082840312156109a0578081fd5b81356001600160e01b0319811681146109b7578182fd5b9392505050565b6000602082840312156109cf578081fd5b813563ffffffff811681146109b7578182fd5b901515815260200190565b90815260200190565b918252602082015260400190565b6020808252602f908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201526e0818591b5a5b881d1bc819dc985b9d608a1b606082015260800190565b6020808252602b908201527f426c6f6f636b53746174653a3a75706461746553746174653a204f4e4c595f4160408201526a4c4c4f5745445f524f4c4560a81b606082015260800190565b60208082526030908201527f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60408201526f2061646d696e20746f207265766f6b6560801b606082015260800190565b6020808252602f908201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560408201526e103937b632b9903337b91039b2b63360891b606082015260800190565b60008219821115610b5057610b50610cfa565b500190565b600082610b6457610b64610d10565b500490565b80825b6001808611610b7b5750610ba6565b818704821115610b8d57610b8d610cfa565b80861615610b9a57918102915b9490941c938002610b6c565b94509492505050565b60006109b760ff80851660ff8516600082610bcc575060016109b7565b81610bd9575060006109b7565b8160018114610bef5760028114610bf957610c26565b60019150506109b7565b60ff841115610c0a57610c0a610cfa565b6001841b915084821115610c2057610c20610cfa565b506109b7565b5060208310610133831016604e8410600b8410161715610c59575081810a83811115610c5457610c54610cfa565b6109b7565b610c668484846001610b69565b808604821115610c7857610c78610cfa565b02949350505050565b600060ff821660ff84168160ff0481118215151615610ca257610ca2610cfa565b029392505050565b600082821015610cbc57610cbc610cfa565b500390565b600063ffffffff83811690831681811015610cde57610cde610cfa565b039392505050565b600082610cf557610cf5610d10565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fdfea2646970667358221220a8e519fbcd9d08b4e0b96c98f59a1579a67f4fe73f44dd381ec3768dd121e2f464736f6c63430008010033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, role_manager common.Address, state_manager common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, role_manager, state_manager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contracts.Contract.DEFAULTADMINROLE(&_Contracts.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contracts.Contract.DEFAULTADMINROLE(&_Contracts.CallOpts)
}

// STATEMANAGER is a free data retrieval call binding the contract method 0x2450a4a2.
//
// Solidity: function STATE_MANAGER() view returns(bytes32)
func (_Contracts *ContractsCaller) STATEMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "STATE_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STATEMANAGER is a free data retrieval call binding the contract method 0x2450a4a2.
//
// Solidity: function STATE_MANAGER() view returns(bytes32)
func (_Contracts *ContractsSession) STATEMANAGER() ([32]byte, error) {
	return _Contracts.Contract.STATEMANAGER(&_Contracts.CallOpts)
}

// STATEMANAGER is a free data retrieval call binding the contract method 0x2450a4a2.
//
// Solidity: function STATE_MANAGER() view returns(bytes32)
func (_Contracts *ContractsCallerSession) STATEMANAGER() ([32]byte, error) {
	return _Contracts.Contract.STATEMANAGER(&_Contracts.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contracts *ContractsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contracts *ContractsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetRoleAdmin(&_Contracts.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contracts *ContractsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetRoleAdmin(&_Contracts.CallOpts, role)
}

// GetState is a free data retrieval call binding the contract method 0x09648a9d.
//
// Solidity: function getState(bytes32 state_root) view returns(uint256)
func (_Contracts *ContractsCaller) GetState(opts *bind.CallOpts, state_root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getState", state_root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x09648a9d.
//
// Solidity: function getState(bytes32 state_root) view returns(uint256)
func (_Contracts *ContractsSession) GetState(state_root [32]byte) (*big.Int, error) {
	return _Contracts.Contract.GetState(&_Contracts.CallOpts, state_root)
}

// GetState is a free data retrieval call binding the contract method 0x09648a9d.
//
// Solidity: function getState(bytes32 state_root) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetState(state_root [32]byte) (*big.Int, error) {
	return _Contracts.Contract.GetState(&_Contracts.CallOpts, state_root)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contracts *ContractsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contracts *ContractsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contracts.Contract.HasRole(&_Contracts.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contracts *ContractsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contracts.Contract.HasRole(&_Contracts.CallOpts, role, account)
}

// IsStatePresent is a free data retrieval call binding the contract method 0x13bf2c97.
//
// Solidity: function isStatePresent(bytes32 state_root) view returns(bool)
func (_Contracts *ContractsCaller) IsStatePresent(opts *bind.CallOpts, state_root [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isStatePresent", state_root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStatePresent is a free data retrieval call binding the contract method 0x13bf2c97.
//
// Solidity: function isStatePresent(bytes32 state_root) view returns(bool)
func (_Contracts *ContractsSession) IsStatePresent(state_root [32]byte) (bool, error) {
	return _Contracts.Contract.IsStatePresent(&_Contracts.CallOpts, state_root)
}

// IsStatePresent is a free data retrieval call binding the contract method 0x13bf2c97.
//
// Solidity: function isStatePresent(bytes32 state_root) view returns(bool)
func (_Contracts *ContractsCallerSession) IsStatePresent(state_root [32]byte) (bool, error) {
	return _Contracts.Contract.IsStatePresent(&_Contracts.CallOpts, state_root)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x0aeb5fa2.
//
// Solidity: function verifyInclusionProof(bytes32[] content, bytes32[] hashes, bytes bitmap, uint32[] depths) view returns(uint256)
func (_Contracts *ContractsCaller) VerifyInclusionProof(opts *bind.CallOpts, content [][32]byte, hashes [][32]byte, bitmap []byte, depths []uint32) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "verifyInclusionProof", content, hashes, bitmap, depths)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x0aeb5fa2.
//
// Solidity: function verifyInclusionProof(bytes32[] content, bytes32[] hashes, bytes bitmap, uint32[] depths) view returns(uint256)
func (_Contracts *ContractsSession) VerifyInclusionProof(content [][32]byte, hashes [][32]byte, bitmap []byte, depths []uint32) (*big.Int, error) {
	return _Contracts.Contract.VerifyInclusionProof(&_Contracts.CallOpts, content, hashes, bitmap, depths)
}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x0aeb5fa2.
//
// Solidity: function verifyInclusionProof(bytes32[] content, bytes32[] hashes, bytes bitmap, uint32[] depths) view returns(uint256)
func (_Contracts *ContractsCallerSession) VerifyInclusionProof(content [][32]byte, hashes [][32]byte, bitmap []byte, depths []uint32) (*big.Int, error) {
	return _Contracts.Contract.VerifyInclusionProof(&_Contracts.CallOpts, content, hashes, bitmap, depths)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contracts *ContractsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantRole(&_Contracts.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantRole(&_Contracts.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Contracts *ContractsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RenounceRole(&_Contracts.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RenounceRole(&_Contracts.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contracts *ContractsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeRole(&_Contracts.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeRole(&_Contracts.TransactOpts, role, account)
}

// UpdateState is a paid mutator transaction binding the contract method 0x8ea59e1d.
//
// Solidity: function updateState(bytes32 state_root) returns()
func (_Contracts *ContractsTransactor) UpdateState(opts *bind.TransactOpts, state_root [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateState", state_root)
}

// UpdateState is a paid mutator transaction binding the contract method 0x8ea59e1d.
//
// Solidity: function updateState(bytes32 state_root) returns()
func (_Contracts *ContractsSession) UpdateState(state_root [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateState(&_Contracts.TransactOpts, state_root)
}

// UpdateState is a paid mutator transaction binding the contract method 0x8ea59e1d.
//
// Solidity: function updateState(bytes32 state_root) returns()
func (_Contracts *ContractsTransactorSession) UpdateState(state_root [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateState(&_Contracts.TransactOpts, state_root)
}

// ContractsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Contracts contract.
type ContractsRoleAdminChangedIterator struct {
	Event *ContractsRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoleAdminChanged represents a RoleAdminChanged event raised by the Contracts contract.
type ContractsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contracts *ContractsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractsRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoleAdminChangedIterator{contract: _Contracts.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contracts *ContractsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoleAdminChanged)
				if err := _Contracts.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contracts *ContractsFilterer) ParseRoleAdminChanged(log types.Log) (*ContractsRoleAdminChanged, error) {
	event := new(ContractsRoleAdminChanged)
	if err := _Contracts.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Contracts contract.
type ContractsRoleGrantedIterator struct {
	Event *ContractsRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoleGranted represents a RoleGranted event raised by the Contracts contract.
type ContractsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractsRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoleGrantedIterator{contract: _Contracts.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoleGranted)
				if err := _Contracts.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) ParseRoleGranted(log types.Log) (*ContractsRoleGranted, error) {
	event := new(ContractsRoleGranted)
	if err := _Contracts.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Contracts contract.
type ContractsRoleRevokedIterator struct {
	Event *ContractsRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoleRevoked represents a RoleRevoked event raised by the Contracts contract.
type ContractsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractsRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoleRevokedIterator{contract: _Contracts.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoleRevoked)
				if err := _Contracts.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) ParseRoleRevoked(log types.Log) (*ContractsRoleRevoked, error) {
	event := new(ContractsRoleRevoked)
	if err := _Contracts.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
