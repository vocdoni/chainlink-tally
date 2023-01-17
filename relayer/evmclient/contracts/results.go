// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractResults

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

// ContractResultsMetaData contains all meta data concerning the ContractResults contract.
var ContractResultsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chainLinkToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_chainlinkOracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"electionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"rawResults\",\"type\":\"bytes\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chainlinkOracle\",\"type\":\"address\"}],\"name\":\"_setChainlinkOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chainlinkToken\",\"type\":\"address\"}],\"name\":\"_setChainlinkToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_result\",\"type\":\"bytes\"}],\"name\":\"decodeResult\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"electionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceNetworkAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[][]\",\"name\":\"scrutiny\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_electionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_organizationId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_sourceNetworkAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[][]\",\"name\":\"_scrutiny\",\"type\":\"uint256[][]\"}],\"name\":\"encodeResult\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"endpoints\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_rawResult\",\"type\":\"bytes\"}],\"name\":\"fulfillRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainLinkOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainLinkToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_electionId\",\"type\":\"bytes32\"}],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceNetworkAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[][]\",\"name\":\"scrutiny\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"results\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceNetworkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_endpoint\",\"type\":\"string\"}],\"name\":\"setEndpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_electionId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_chainName\",\"type\":\"string\"}],\"name\":\"setResult\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractResultsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractResultsMetaData.ABI instead.
var ContractResultsABI = ContractResultsMetaData.ABI

// ContractResults is an auto generated Go binding around an Ethereum contract.
type ContractResults struct {
	ContractResultsCaller     // Read-only binding to the contract
	ContractResultsTransactor // Write-only binding to the contract
	ContractResultsFilterer   // Log filterer for contract events
}

// ContractResultsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractResultsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractResultsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractResultsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractResultsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractResultsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractResultsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractResultsSession struct {
	Contract     *ContractResults  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractResultsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractResultsCallerSession struct {
	Contract *ContractResultsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContractResultsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractResultsTransactorSession struct {
	Contract     *ContractResultsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractResultsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractResultsRaw struct {
	Contract *ContractResults // Generic contract binding to access the raw methods on
}

// ContractResultsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractResultsCallerRaw struct {
	Contract *ContractResultsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractResultsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractResultsTransactorRaw struct {
	Contract *ContractResultsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractResults creates a new instance of ContractResults, bound to a specific deployed contract.
func NewContractResults(address common.Address, backend bind.ContractBackend) (*ContractResults, error) {
	contract, err := bindContractResults(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractResults{ContractResultsCaller: ContractResultsCaller{contract: contract}, ContractResultsTransactor: ContractResultsTransactor{contract: contract}, ContractResultsFilterer: ContractResultsFilterer{contract: contract}}, nil
}

// NewContractResultsCaller creates a new read-only instance of ContractResults, bound to a specific deployed contract.
func NewContractResultsCaller(address common.Address, caller bind.ContractCaller) (*ContractResultsCaller, error) {
	contract, err := bindContractResults(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractResultsCaller{contract: contract}, nil
}

// NewContractResultsTransactor creates a new write-only instance of ContractResults, bound to a specific deployed contract.
func NewContractResultsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractResultsTransactor, error) {
	contract, err := bindContractResults(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractResultsTransactor{contract: contract}, nil
}

// NewContractResultsFilterer creates a new log filterer instance of ContractResults, bound to a specific deployed contract.
func NewContractResultsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractResultsFilterer, error) {
	contract, err := bindContractResults(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractResultsFilterer{contract: contract}, nil
}

// bindContractResults binds a generic wrapper to an already deployed contract.
func bindContractResults(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractResultsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractResults *ContractResultsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractResults.Contract.ContractResultsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractResults *ContractResultsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractResults.Contract.ContractResultsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractResults *ContractResultsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractResults.Contract.ContractResultsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractResults *ContractResultsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractResults.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractResults *ContractResultsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractResults.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractResults *ContractResultsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractResults.Contract.contract.Transact(opts, method, params...)
}

// DecodeResult is a free data retrieval call binding the contract method 0x7ecc184d.
//
// Solidity: function decodeResult(bytes _result) pure returns(bytes32 electionId, address organizationId, bytes32 censusRoot, address sourceNetworkAddress, uint256[][] scrutiny)
func (_ContractResults *ContractResultsCaller) DecodeResult(opts *bind.CallOpts, _result []byte) (struct {
	ElectionId           [32]byte
	OrganizationId       common.Address
	CensusRoot           [32]byte
	SourceNetworkAddress common.Address
	Scrutiny             [][]*big.Int
}, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "decodeResult", _result)

	outstruct := new(struct {
		ElectionId           [32]byte
		OrganizationId       common.Address
		CensusRoot           [32]byte
		SourceNetworkAddress common.Address
		Scrutiny             [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ElectionId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.OrganizationId = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CensusRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.SourceNetworkAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Scrutiny = *abi.ConvertType(out[4], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// DecodeResult is a free data retrieval call binding the contract method 0x7ecc184d.
//
// Solidity: function decodeResult(bytes _result) pure returns(bytes32 electionId, address organizationId, bytes32 censusRoot, address sourceNetworkAddress, uint256[][] scrutiny)
func (_ContractResults *ContractResultsSession) DecodeResult(_result []byte) (struct {
	ElectionId           [32]byte
	OrganizationId       common.Address
	CensusRoot           [32]byte
	SourceNetworkAddress common.Address
	Scrutiny             [][]*big.Int
}, error) {
	return _ContractResults.Contract.DecodeResult(&_ContractResults.CallOpts, _result)
}

// DecodeResult is a free data retrieval call binding the contract method 0x7ecc184d.
//
// Solidity: function decodeResult(bytes _result) pure returns(bytes32 electionId, address organizationId, bytes32 censusRoot, address sourceNetworkAddress, uint256[][] scrutiny)
func (_ContractResults *ContractResultsCallerSession) DecodeResult(_result []byte) (struct {
	ElectionId           [32]byte
	OrganizationId       common.Address
	CensusRoot           [32]byte
	SourceNetworkAddress common.Address
	Scrutiny             [][]*big.Int
}, error) {
	return _ContractResults.Contract.DecodeResult(&_ContractResults.CallOpts, _result)
}

// EncodeResult is a free data retrieval call binding the contract method 0x980c72bc.
//
// Solidity: function encodeResult(bytes32 _electionId, address _organizationId, bytes32 _censusRoot, address _sourceNetworkAddress, uint256[][] _scrutiny) pure returns(bytes)
func (_ContractResults *ContractResultsCaller) EncodeResult(opts *bind.CallOpts, _electionId [32]byte, _organizationId common.Address, _censusRoot [32]byte, _sourceNetworkAddress common.Address, _scrutiny [][]*big.Int) ([]byte, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "encodeResult", _electionId, _organizationId, _censusRoot, _sourceNetworkAddress, _scrutiny)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeResult is a free data retrieval call binding the contract method 0x980c72bc.
//
// Solidity: function encodeResult(bytes32 _electionId, address _organizationId, bytes32 _censusRoot, address _sourceNetworkAddress, uint256[][] _scrutiny) pure returns(bytes)
func (_ContractResults *ContractResultsSession) EncodeResult(_electionId [32]byte, _organizationId common.Address, _censusRoot [32]byte, _sourceNetworkAddress common.Address, _scrutiny [][]*big.Int) ([]byte, error) {
	return _ContractResults.Contract.EncodeResult(&_ContractResults.CallOpts, _electionId, _organizationId, _censusRoot, _sourceNetworkAddress, _scrutiny)
}

// EncodeResult is a free data retrieval call binding the contract method 0x980c72bc.
//
// Solidity: function encodeResult(bytes32 _electionId, address _organizationId, bytes32 _censusRoot, address _sourceNetworkAddress, uint256[][] _scrutiny) pure returns(bytes)
func (_ContractResults *ContractResultsCallerSession) EncodeResult(_electionId [32]byte, _organizationId common.Address, _censusRoot [32]byte, _sourceNetworkAddress common.Address, _scrutiny [][]*big.Int) ([]byte, error) {
	return _ContractResults.Contract.EncodeResult(&_ContractResults.CallOpts, _electionId, _organizationId, _censusRoot, _sourceNetworkAddress, _scrutiny)
}

// Endpoints is a free data retrieval call binding the contract method 0x84777429.
//
// Solidity: function endpoints(string ) view returns(string)
func (_ContractResults *ContractResultsCaller) Endpoints(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "endpoints", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Endpoints is a free data retrieval call binding the contract method 0x84777429.
//
// Solidity: function endpoints(string ) view returns(string)
func (_ContractResults *ContractResultsSession) Endpoints(arg0 string) (string, error) {
	return _ContractResults.Contract.Endpoints(&_ContractResults.CallOpts, arg0)
}

// Endpoints is a free data retrieval call binding the contract method 0x84777429.
//
// Solidity: function endpoints(string ) view returns(string)
func (_ContractResults *ContractResultsCallerSession) Endpoints(arg0 string) (string, error) {
	return _ContractResults.Contract.Endpoints(&_ContractResults.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ContractResults *ContractResultsCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ContractResults *ContractResultsSession) Fee() (*big.Int, error) {
	return _ContractResults.Contract.Fee(&_ContractResults.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ContractResults *ContractResultsCallerSession) Fee() (*big.Int, error) {
	return _ContractResults.Contract.Fee(&_ContractResults.CallOpts)
}

// GetChainLinkOracle is a free data retrieval call binding the contract method 0xbdf21ca3.
//
// Solidity: function getChainLinkOracle() view returns(address)
func (_ContractResults *ContractResultsCaller) GetChainLinkOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "getChainLinkOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChainLinkOracle is a free data retrieval call binding the contract method 0xbdf21ca3.
//
// Solidity: function getChainLinkOracle() view returns(address)
func (_ContractResults *ContractResultsSession) GetChainLinkOracle() (common.Address, error) {
	return _ContractResults.Contract.GetChainLinkOracle(&_ContractResults.CallOpts)
}

// GetChainLinkOracle is a free data retrieval call binding the contract method 0xbdf21ca3.
//
// Solidity: function getChainLinkOracle() view returns(address)
func (_ContractResults *ContractResultsCallerSession) GetChainLinkOracle() (common.Address, error) {
	return _ContractResults.Contract.GetChainLinkOracle(&_ContractResults.CallOpts)
}

// GetChainLinkToken is a free data retrieval call binding the contract method 0x17e13d08.
//
// Solidity: function getChainLinkToken() view returns(address)
func (_ContractResults *ContractResultsCaller) GetChainLinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "getChainLinkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChainLinkToken is a free data retrieval call binding the contract method 0x17e13d08.
//
// Solidity: function getChainLinkToken() view returns(address)
func (_ContractResults *ContractResultsSession) GetChainLinkToken() (common.Address, error) {
	return _ContractResults.Contract.GetChainLinkToken(&_ContractResults.CallOpts)
}

// GetChainLinkToken is a free data retrieval call binding the contract method 0x17e13d08.
//
// Solidity: function getChainLinkToken() view returns(address)
func (_ContractResults *ContractResultsCallerSession) GetChainLinkToken() (common.Address, error) {
	return _ContractResults.Contract.GetChainLinkToken(&_ContractResults.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _electionId) view returns(address organizationId, bytes32 censusRoot, address sourceNetworkAddress, uint256[][] scrutiny)
func (_ContractResults *ContractResultsCaller) GetResult(opts *bind.CallOpts, _electionId [32]byte) (struct {
	OrganizationId       common.Address
	CensusRoot           [32]byte
	SourceNetworkAddress common.Address
	Scrutiny             [][]*big.Int
}, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "getResult", _electionId)

	outstruct := new(struct {
		OrganizationId       common.Address
		CensusRoot           [32]byte
		SourceNetworkAddress common.Address
		Scrutiny             [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrganizationId = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CensusRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SourceNetworkAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Scrutiny = *abi.ConvertType(out[3], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _electionId) view returns(address organizationId, bytes32 censusRoot, address sourceNetworkAddress, uint256[][] scrutiny)
func (_ContractResults *ContractResultsSession) GetResult(_electionId [32]byte) (struct {
	OrganizationId       common.Address
	CensusRoot           [32]byte
	SourceNetworkAddress common.Address
	Scrutiny             [][]*big.Int
}, error) {
	return _ContractResults.Contract.GetResult(&_ContractResults.CallOpts, _electionId)
}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _electionId) view returns(address organizationId, bytes32 censusRoot, address sourceNetworkAddress, uint256[][] scrutiny)
func (_ContractResults *ContractResultsCallerSession) GetResult(_electionId [32]byte) (struct {
	OrganizationId       common.Address
	CensusRoot           [32]byte
	SourceNetworkAddress common.Address
	Scrutiny             [][]*big.Int
}, error) {
	return _ContractResults.Contract.GetResult(&_ContractResults.CallOpts, _electionId)
}

// JobId is a free data retrieval call binding the contract method 0xc2939d97.
//
// Solidity: function jobId() view returns(bytes32)
func (_ContractResults *ContractResultsCaller) JobId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "jobId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// JobId is a free data retrieval call binding the contract method 0xc2939d97.
//
// Solidity: function jobId() view returns(bytes32)
func (_ContractResults *ContractResultsSession) JobId() ([32]byte, error) {
	return _ContractResults.Contract.JobId(&_ContractResults.CallOpts)
}

// JobId is a free data retrieval call binding the contract method 0xc2939d97.
//
// Solidity: function jobId() view returns(bytes32)
func (_ContractResults *ContractResultsCallerSession) JobId() ([32]byte, error) {
	return _ContractResults.Contract.JobId(&_ContractResults.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractResults *ContractResultsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractResults *ContractResultsSession) Owner() (common.Address, error) {
	return _ContractResults.Contract.Owner(&_ContractResults.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractResults *ContractResultsCallerSession) Owner() (common.Address, error) {
	return _ContractResults.Contract.Owner(&_ContractResults.CallOpts)
}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(address organizationId, bytes32 censusRoot, address sourceNetworkAddress)
func (_ContractResults *ContractResultsCaller) Results(opts *bind.CallOpts, arg0 [32]byte) (struct {
	OrganizationId       common.Address
	CensusRoot           [32]byte
	SourceNetworkAddress common.Address
}, error) {
	var out []interface{}
	err := _ContractResults.contract.Call(opts, &out, "results", arg0)

	outstruct := new(struct {
		OrganizationId       common.Address
		CensusRoot           [32]byte
		SourceNetworkAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrganizationId = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CensusRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SourceNetworkAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(address organizationId, bytes32 censusRoot, address sourceNetworkAddress)
func (_ContractResults *ContractResultsSession) Results(arg0 [32]byte) (struct {
	OrganizationId       common.Address
	CensusRoot           [32]byte
	SourceNetworkAddress common.Address
}, error) {
	return _ContractResults.Contract.Results(&_ContractResults.CallOpts, arg0)
}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(address organizationId, bytes32 censusRoot, address sourceNetworkAddress)
func (_ContractResults *ContractResultsCallerSession) Results(arg0 [32]byte) (struct {
	OrganizationId       common.Address
	CensusRoot           [32]byte
	SourceNetworkAddress common.Address
}, error) {
	return _ContractResults.Contract.Results(&_ContractResults.CallOpts, arg0)
}

// SetChainlinkOracle is a paid mutator transaction binding the contract method 0x2ccbc4a8.
//
// Solidity: function _setChainlinkOracle(address _chainlinkOracle) returns()
func (_ContractResults *ContractResultsTransactor) SetChainlinkOracle(opts *bind.TransactOpts, _chainlinkOracle common.Address) (*types.Transaction, error) {
	return _ContractResults.contract.Transact(opts, "_setChainlinkOracle", _chainlinkOracle)
}

// SetChainlinkOracle is a paid mutator transaction binding the contract method 0x2ccbc4a8.
//
// Solidity: function _setChainlinkOracle(address _chainlinkOracle) returns()
func (_ContractResults *ContractResultsSession) SetChainlinkOracle(_chainlinkOracle common.Address) (*types.Transaction, error) {
	return _ContractResults.Contract.SetChainlinkOracle(&_ContractResults.TransactOpts, _chainlinkOracle)
}

// SetChainlinkOracle is a paid mutator transaction binding the contract method 0x2ccbc4a8.
//
// Solidity: function _setChainlinkOracle(address _chainlinkOracle) returns()
func (_ContractResults *ContractResultsTransactorSession) SetChainlinkOracle(_chainlinkOracle common.Address) (*types.Transaction, error) {
	return _ContractResults.Contract.SetChainlinkOracle(&_ContractResults.TransactOpts, _chainlinkOracle)
}

// SetChainlinkToken is a paid mutator transaction binding the contract method 0x5c86aae7.
//
// Solidity: function _setChainlinkToken(address _chainlinkToken) returns()
func (_ContractResults *ContractResultsTransactor) SetChainlinkToken(opts *bind.TransactOpts, _chainlinkToken common.Address) (*types.Transaction, error) {
	return _ContractResults.contract.Transact(opts, "_setChainlinkToken", _chainlinkToken)
}

// SetChainlinkToken is a paid mutator transaction binding the contract method 0x5c86aae7.
//
// Solidity: function _setChainlinkToken(address _chainlinkToken) returns()
func (_ContractResults *ContractResultsSession) SetChainlinkToken(_chainlinkToken common.Address) (*types.Transaction, error) {
	return _ContractResults.Contract.SetChainlinkToken(&_ContractResults.TransactOpts, _chainlinkToken)
}

// SetChainlinkToken is a paid mutator transaction binding the contract method 0x5c86aae7.
//
// Solidity: function _setChainlinkToken(address _chainlinkToken) returns()
func (_ContractResults *ContractResultsTransactorSession) SetChainlinkToken(_chainlinkToken common.Address) (*types.Transaction, error) {
	return _ContractResults.Contract.SetChainlinkToken(&_ContractResults.TransactOpts, _chainlinkToken)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0x4d48f790.
//
// Solidity: function fulfillRequest(bytes32 _requestId, bytes _rawResult) returns()
func (_ContractResults *ContractResultsTransactor) FulfillRequest(opts *bind.TransactOpts, _requestId [32]byte, _rawResult []byte) (*types.Transaction, error) {
	return _ContractResults.contract.Transact(opts, "fulfillRequest", _requestId, _rawResult)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0x4d48f790.
//
// Solidity: function fulfillRequest(bytes32 _requestId, bytes _rawResult) returns()
func (_ContractResults *ContractResultsSession) FulfillRequest(_requestId [32]byte, _rawResult []byte) (*types.Transaction, error) {
	return _ContractResults.Contract.FulfillRequest(&_ContractResults.TransactOpts, _requestId, _rawResult)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0x4d48f790.
//
// Solidity: function fulfillRequest(bytes32 _requestId, bytes _rawResult) returns()
func (_ContractResults *ContractResultsTransactorSession) FulfillRequest(_requestId [32]byte, _rawResult []byte) (*types.Transaction, error) {
	return _ContractResults.Contract.FulfillRequest(&_ContractResults.TransactOpts, _requestId, _rawResult)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractResults *ContractResultsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractResults.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractResults *ContractResultsSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractResults.Contract.RenounceOwnership(&_ContractResults.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractResults *ContractResultsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractResults.Contract.RenounceOwnership(&_ContractResults.TransactOpts)
}

// SetEndpoint is a paid mutator transaction binding the contract method 0xc7571d73.
//
// Solidity: function setEndpoint(string _name, string _endpoint) returns()
func (_ContractResults *ContractResultsTransactor) SetEndpoint(opts *bind.TransactOpts, _name string, _endpoint string) (*types.Transaction, error) {
	return _ContractResults.contract.Transact(opts, "setEndpoint", _name, _endpoint)
}

// SetEndpoint is a paid mutator transaction binding the contract method 0xc7571d73.
//
// Solidity: function setEndpoint(string _name, string _endpoint) returns()
func (_ContractResults *ContractResultsSession) SetEndpoint(_name string, _endpoint string) (*types.Transaction, error) {
	return _ContractResults.Contract.SetEndpoint(&_ContractResults.TransactOpts, _name, _endpoint)
}

// SetEndpoint is a paid mutator transaction binding the contract method 0xc7571d73.
//
// Solidity: function setEndpoint(string _name, string _endpoint) returns()
func (_ContractResults *ContractResultsTransactorSession) SetEndpoint(_name string, _endpoint string) (*types.Transaction, error) {
	return _ContractResults.Contract.SetEndpoint(&_ContractResults.TransactOpts, _name, _endpoint)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_ContractResults *ContractResultsTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _ContractResults.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_ContractResults *ContractResultsSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _ContractResults.Contract.SetFee(&_ContractResults.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_ContractResults *ContractResultsTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _ContractResults.Contract.SetFee(&_ContractResults.TransactOpts, _fee)
}

// SetResult is a paid mutator transaction binding the contract method 0xef0caaa9.
//
// Solidity: function setResult(bytes32 _electionId, string _chainName) returns(bytes32 requestId)
func (_ContractResults *ContractResultsTransactor) SetResult(opts *bind.TransactOpts, _electionId [32]byte, _chainName string) (*types.Transaction, error) {
	return _ContractResults.contract.Transact(opts, "setResult", _electionId, _chainName)
}

// SetResult is a paid mutator transaction binding the contract method 0xef0caaa9.
//
// Solidity: function setResult(bytes32 _electionId, string _chainName) returns(bytes32 requestId)
func (_ContractResults *ContractResultsSession) SetResult(_electionId [32]byte, _chainName string) (*types.Transaction, error) {
	return _ContractResults.Contract.SetResult(&_ContractResults.TransactOpts, _electionId, _chainName)
}

// SetResult is a paid mutator transaction binding the contract method 0xef0caaa9.
//
// Solidity: function setResult(bytes32 _electionId, string _chainName) returns(bytes32 requestId)
func (_ContractResults *ContractResultsTransactorSession) SetResult(_electionId [32]byte, _chainName string) (*types.Transaction, error) {
	return _ContractResults.Contract.SetResult(&_ContractResults.TransactOpts, _electionId, _chainName)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractResults *ContractResultsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractResults.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractResults *ContractResultsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractResults.Contract.TransferOwnership(&_ContractResults.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractResults *ContractResultsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractResults.Contract.TransferOwnership(&_ContractResults.TransactOpts, newOwner)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_ContractResults *ContractResultsTransactor) WithdrawLink(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractResults.contract.Transact(opts, "withdrawLink")
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_ContractResults *ContractResultsSession) WithdrawLink() (*types.Transaction, error) {
	return _ContractResults.Contract.WithdrawLink(&_ContractResults.TransactOpts)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_ContractResults *ContractResultsTransactorSession) WithdrawLink() (*types.Transaction, error) {
	return _ContractResults.Contract.WithdrawLink(&_ContractResults.TransactOpts)
}

// ContractResultsChainlinkCancelledIterator is returned from FilterChainlinkCancelled and is used to iterate over the raw logs and unpacked data for ChainlinkCancelled events raised by the ContractResults contract.
type ContractResultsChainlinkCancelledIterator struct {
	Event *ContractResultsChainlinkCancelled // Event containing the contract specifics and raw log

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
func (it *ContractResultsChainlinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractResultsChainlinkCancelled)
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
		it.Event = new(ContractResultsChainlinkCancelled)
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
func (it *ContractResultsChainlinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractResultsChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractResultsChainlinkCancelled represents a ChainlinkCancelled event raised by the ContractResults contract.
type ContractResultsChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCancelled is a free log retrieval operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_ContractResults *ContractResultsFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*ContractResultsChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractResults.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractResultsChainlinkCancelledIterator{contract: _ContractResults.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

// WatchChainlinkCancelled is a free log subscription operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_ContractResults *ContractResultsFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *ContractResultsChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractResults.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractResultsChainlinkCancelled)
				if err := _ContractResults.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
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

// ParseChainlinkCancelled is a log parse operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_ContractResults *ContractResultsFilterer) ParseChainlinkCancelled(log types.Log) (*ContractResultsChainlinkCancelled, error) {
	event := new(ContractResultsChainlinkCancelled)
	if err := _ContractResults.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractResultsChainlinkFulfilledIterator is returned from FilterChainlinkFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkFulfilled events raised by the ContractResults contract.
type ContractResultsChainlinkFulfilledIterator struct {
	Event *ContractResultsChainlinkFulfilled // Event containing the contract specifics and raw log

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
func (it *ContractResultsChainlinkFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractResultsChainlinkFulfilled)
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
		it.Event = new(ContractResultsChainlinkFulfilled)
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
func (it *ContractResultsChainlinkFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractResultsChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractResultsChainlinkFulfilled represents a ChainlinkFulfilled event raised by the ContractResults contract.
type ContractResultsChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkFulfilled is a free log retrieval operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_ContractResults *ContractResultsFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*ContractResultsChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractResults.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractResultsChainlinkFulfilledIterator{contract: _ContractResults.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkFulfilled is a free log subscription operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_ContractResults *ContractResultsFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *ContractResultsChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractResults.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractResultsChainlinkFulfilled)
				if err := _ContractResults.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
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

// ParseChainlinkFulfilled is a log parse operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_ContractResults *ContractResultsFilterer) ParseChainlinkFulfilled(log types.Log) (*ContractResultsChainlinkFulfilled, error) {
	event := new(ContractResultsChainlinkFulfilled)
	if err := _ContractResults.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractResultsChainlinkRequestedIterator is returned from FilterChainlinkRequested and is used to iterate over the raw logs and unpacked data for ChainlinkRequested events raised by the ContractResults contract.
type ContractResultsChainlinkRequestedIterator struct {
	Event *ContractResultsChainlinkRequested // Event containing the contract specifics and raw log

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
func (it *ContractResultsChainlinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractResultsChainlinkRequested)
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
		it.Event = new(ContractResultsChainlinkRequested)
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
func (it *ContractResultsChainlinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractResultsChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractResultsChainlinkRequested represents a ChainlinkRequested event raised by the ContractResults contract.
type ContractResultsChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkRequested is a free log retrieval operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_ContractResults *ContractResultsFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*ContractResultsChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractResults.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractResultsChainlinkRequestedIterator{contract: _ContractResults.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

// WatchChainlinkRequested is a free log subscription operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_ContractResults *ContractResultsFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *ContractResultsChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractResults.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractResultsChainlinkRequested)
				if err := _ContractResults.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
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

// ParseChainlinkRequested is a log parse operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_ContractResults *ContractResultsFilterer) ParseChainlinkRequested(log types.Log) (*ContractResultsChainlinkRequested, error) {
	event := new(ContractResultsChainlinkRequested)
	if err := _ContractResults.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractResultsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractResults contract.
type ContractResultsOwnershipTransferredIterator struct {
	Event *ContractResultsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractResultsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractResultsOwnershipTransferred)
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
		it.Event = new(ContractResultsOwnershipTransferred)
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
func (it *ContractResultsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractResultsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractResultsOwnershipTransferred represents a OwnershipTransferred event raised by the ContractResults contract.
type ContractResultsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractResults *ContractResultsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractResultsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractResults.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractResultsOwnershipTransferredIterator{contract: _ContractResults.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractResults *ContractResultsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractResultsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractResults.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractResultsOwnershipTransferred)
				if err := _ContractResults.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractResults *ContractResultsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractResultsOwnershipTransferred, error) {
	event := new(ContractResultsOwnershipTransferred)
	if err := _ContractResults.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractResultsRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the ContractResults contract.
type ContractResultsRequestFulfilledIterator struct {
	Event *ContractResultsRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *ContractResultsRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractResultsRequestFulfilled)
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
		it.Event = new(ContractResultsRequestFulfilled)
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
func (it *ContractResultsRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractResultsRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractResultsRequestFulfilled represents a RequestFulfilled event raised by the ContractResults contract.
type ContractResultsRequestFulfilled struct {
	RequestId  [32]byte
	ElectionId [32]byte
	RawResults common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0x1a98d10aaebe922faa52b48e7716f10b15f116411a84e22f0d9fddeb48509fa9.
//
// Solidity: event RequestFulfilled(bytes32 indexed requestId, bytes32 indexed electionId, bytes indexed rawResults)
func (_ContractResults *ContractResultsFilterer) FilterRequestFulfilled(opts *bind.FilterOpts, requestId [][32]byte, electionId [][32]byte, rawResults [][]byte) (*ContractResultsRequestFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}
	var rawResultsRule []interface{}
	for _, rawResultsItem := range rawResults {
		rawResultsRule = append(rawResultsRule, rawResultsItem)
	}

	logs, sub, err := _ContractResults.contract.FilterLogs(opts, "RequestFulfilled", requestIdRule, electionIdRule, rawResultsRule)
	if err != nil {
		return nil, err
	}
	return &ContractResultsRequestFulfilledIterator{contract: _ContractResults.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0x1a98d10aaebe922faa52b48e7716f10b15f116411a84e22f0d9fddeb48509fa9.
//
// Solidity: event RequestFulfilled(bytes32 indexed requestId, bytes32 indexed electionId, bytes indexed rawResults)
func (_ContractResults *ContractResultsFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *ContractResultsRequestFulfilled, requestId [][32]byte, electionId [][32]byte, rawResults [][]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}
	var rawResultsRule []interface{}
	for _, rawResultsItem := range rawResults {
		rawResultsRule = append(rawResultsRule, rawResultsItem)
	}

	logs, sub, err := _ContractResults.contract.WatchLogs(opts, "RequestFulfilled", requestIdRule, electionIdRule, rawResultsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractResultsRequestFulfilled)
				if err := _ContractResults.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

// ParseRequestFulfilled is a log parse operation binding the contract event 0x1a98d10aaebe922faa52b48e7716f10b15f116411a84e22f0d9fddeb48509fa9.
//
// Solidity: event RequestFulfilled(bytes32 indexed requestId, bytes32 indexed electionId, bytes indexed rawResults)
func (_ContractResults *ContractResultsFilterer) ParseRequestFulfilled(log types.Log) (*ContractResultsRequestFulfilled, error) {
	event := new(ContractResultsRequestFulfilled)
	if err := _ContractResults.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
