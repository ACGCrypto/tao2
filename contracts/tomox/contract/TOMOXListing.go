// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/tao2-core/tao2-core/accounts/abi"
	"github.com/tao2-core/tao2-core/accounts/abi/bind"
	"github.com/tao2-core/tao2-core/common"
	"github.com/tao2-core/tao2-core/core/types"
)

// WAIHUIListingABI is the input ABI used to generate the binding from.
const WAIHUIListingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// WAIHUIListingBin is the compiled bytecode used for deploying new contracts.
const WAIHUIListingBin = `0x608060405234801561001057600080fd5b506102bf806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639d63848a811461005b578063a3ff31b5146100c0578063c6b32f34146100f5575b600080fd5b34801561006757600080fd5b5061007061010b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ac578181015183820152602001610094565b505050509050019250505060405180910390f35b3480156100cc57600080fd5b506100e1600160a060020a036004351661016d565b604080519115158252519081900360200190f35b610109600160a060020a036004351661018b565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561016357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610145575b5050505050905090565b600160a060020a031660009081526001602052604090205460ff1690565b80600160a060020a03811615156101a157600080fd5b600160a060020a03811660009081526001602081905260409091205460ff16151514156101cd57600080fd5b683635c9adc5dea000003410156101e357600080fd5b6040516068903480156108fc02916000818181858888f19350505050158015610210573d6000803e3d6000fd5b505060008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556040805160208082018352838252948452919093529190209051815460ff19169015151790555600a165627a7a72305820a258c7f15c7c6507a28499e1a95c7e7ca19f22f78bcf25bf0b842006720fd85d0029`

// DeployWAIHUIListing deploys a new Ethereum contract, binding an instance of WAIHUIListing to it.
func DeployWAIHUIListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WAIHUIListing, error) {
	parsed, err := abi.JSON(strings.NewReader(WAIHUIListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WAIHUIListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WAIHUIListing{WAIHUIListingCaller: WAIHUIListingCaller{contract: contract}, WAIHUIListingTransactor: WAIHUIListingTransactor{contract: contract}, WAIHUIListingFilterer: WAIHUIListingFilterer{contract: contract}}, nil
}

// WAIHUIListing is an auto generated Go binding around an Ethereum contract.
type WAIHUIListing struct {
	WAIHUIListingCaller     // Read-only binding to the contract
	WAIHUIListingTransactor // Write-only binding to the contract
	WAIHUIListingFilterer   // Log filterer for contract events
}

// WAIHUIListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type WAIHUIListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WAIHUIListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WAIHUIListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WAIHUIListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WAIHUIListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WAIHUIListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WAIHUIListingSession struct {
	Contract     *WAIHUIListing     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WAIHUIListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WAIHUIListingCallerSession struct {
	Contract *WAIHUIListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WAIHUIListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WAIHUIListingTransactorSession struct {
	Contract     *WAIHUIListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WAIHUIListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type WAIHUIListingRaw struct {
	Contract *WAIHUIListing // Generic contract binding to access the raw methods on
}

// WAIHUIListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WAIHUIListingCallerRaw struct {
	Contract *WAIHUIListingCaller // Generic read-only contract binding to access the raw methods on
}

// WAIHUIListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WAIHUIListingTransactorRaw struct {
	Contract *WAIHUIListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWAIHUIListing creates a new instance of WAIHUIListing, bound to a specific deployed contract.
func NewWAIHUIListing(address common.Address, backend bind.ContractBackend) (*WAIHUIListing, error) {
	contract, err := bindWAIHUIListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WAIHUIListing{WAIHUIListingCaller: WAIHUIListingCaller{contract: contract}, WAIHUIListingTransactor: WAIHUIListingTransactor{contract: contract}, WAIHUIListingFilterer: WAIHUIListingFilterer{contract: contract}}, nil
}

// NewWAIHUIListingCaller creates a new read-only instance of WAIHUIListing, bound to a specific deployed contract.
func NewWAIHUIListingCaller(address common.Address, caller bind.ContractCaller) (*WAIHUIListingCaller, error) {
	contract, err := bindWAIHUIListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WAIHUIListingCaller{contract: contract}, nil
}

// NewWAIHUIListingTransactor creates a new write-only instance of WAIHUIListing, bound to a specific deployed contract.
func NewWAIHUIListingTransactor(address common.Address, transactor bind.ContractTransactor) (*WAIHUIListingTransactor, error) {
	contract, err := bindWAIHUIListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WAIHUIListingTransactor{contract: contract}, nil
}

// NewWAIHUIListingFilterer creates a new log filterer instance of WAIHUIListing, bound to a specific deployed contract.
func NewWAIHUIListingFilterer(address common.Address, filterer bind.ContractFilterer) (*WAIHUIListingFilterer, error) {
	contract, err := bindWAIHUIListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WAIHUIListingFilterer{contract: contract}, nil
}

// bindWAIHUIListing binds a generic wrapper to an already deployed contract.
func bindWAIHUIListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WAIHUIListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WAIHUIListing *WAIHUIListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WAIHUIListing.Contract.WAIHUIListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WAIHUIListing *WAIHUIListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WAIHUIListing.Contract.WAIHUIListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WAIHUIListing *WAIHUIListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WAIHUIListing.Contract.WAIHUIListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WAIHUIListing *WAIHUIListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WAIHUIListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WAIHUIListing *WAIHUIListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WAIHUIListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WAIHUIListing *WAIHUIListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WAIHUIListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_WAIHUIListing *WAIHUIListingCaller) GetTokenStatus(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WAIHUIListing.contract.Call(opts, out, "getTokenStatus", token)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_WAIHUIListing *WAIHUIListingSession) GetTokenStatus(token common.Address) (bool, error) {
	return _WAIHUIListing.Contract.GetTokenStatus(&_WAIHUIListing.CallOpts, token)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_WAIHUIListing *WAIHUIListingCallerSession) GetTokenStatus(token common.Address) (bool, error) {
	return _WAIHUIListing.Contract.GetTokenStatus(&_WAIHUIListing.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_WAIHUIListing *WAIHUIListingCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _WAIHUIListing.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_WAIHUIListing *WAIHUIListingSession) Tokens() ([]common.Address, error) {
	return _WAIHUIListing.Contract.Tokens(&_WAIHUIListing.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_WAIHUIListing *WAIHUIListingCallerSession) Tokens() ([]common.Address, error) {
	return _WAIHUIListing.Contract.Tokens(&_WAIHUIListing.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_WAIHUIListing *WAIHUIListingTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _WAIHUIListing.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_WAIHUIListing *WAIHUIListingSession) Apply(token common.Address) (*types.Transaction, error) {
	return _WAIHUIListing.Contract.Apply(&_WAIHUIListing.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_WAIHUIListing *WAIHUIListingTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _WAIHUIListing.Contract.Apply(&_WAIHUIListing.TransactOpts, token)
}
