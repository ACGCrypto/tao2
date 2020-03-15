// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"github.com/taoblockchain/tao2"
	"github.com/taoblockchain/tao2/accounts/abi"
	"github.com/taoblockchain/tao2/accounts/abi/bind"
	"github.com/taoblockchain/tao2/common"
	"github.com/taoblockchain/tao2/core/types"
	"github.com/taoblockchain/tao2/event"
	"math/big"
	"strings"
)

// AbstractTokenTRC2ABI is the input ABI used to generate the binding from.
const AbstractTokenTRC2ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AbstractTokenTRC2Bin is the compiled bytecode used for deploying new contracts.
const AbstractTokenTRC2Bin = `0x`

// DeployAbstractTokenTRC2 deploys a new Ethereum contract, binding an instance of AbstractTokenTRC2 to it.
func DeployAbstractTokenTRC2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AbstractTokenTRC2, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractTokenTRC2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbstractTokenTRC2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AbstractTokenTRC2{AbstractTokenTRC2Caller: AbstractTokenTRC2Caller{contract: contract}, AbstractTokenTRC2Transactor: AbstractTokenTRC2Transactor{contract: contract}, AbstractTokenTRC2Filterer: AbstractTokenTRC2Filterer{contract: contract}}, nil
}

// AbstractTokenTRC2 is an auto generated Go binding around an Ethereum contract.
type AbstractTokenTRC2 struct {
	AbstractTokenTRC2Caller     // Read-only binding to the contract
	AbstractTokenTRC2Transactor // Write-only binding to the contract
	AbstractTokenTRC2Filterer   // Log filterer for contract events
}

// AbstractTokenTRC2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractTokenTRC2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractTokenTRC2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractTokenTRC2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractTokenTRC2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractTokenTRC2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractTokenTRC2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractTokenTRC2Session struct {
	Contract     *AbstractTokenTRC2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AbstractTokenTRC2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractTokenTRC2CallerSession struct {
	Contract *AbstractTokenTRC2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AbstractTokenTRC2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractTokenTRC2TransactorSession struct {
	Contract     *AbstractTokenTRC2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AbstractTokenTRC2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractTokenTRC2Raw struct {
	Contract *AbstractTokenTRC2 // Generic contract binding to access the raw methods on
}

// AbstractTokenTRC2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractTokenTRC2CallerRaw struct {
	Contract *AbstractTokenTRC2Caller // Generic read-only contract binding to access the raw methods on
}

// AbstractTokenTRC2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractTokenTRC2TransactorRaw struct {
	Contract *AbstractTokenTRC2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractTokenTRC2 creates a new instance of AbstractTokenTRC2, bound to a specific deployed contract.
func NewAbstractTokenTRC2(address common.Address, backend bind.ContractBackend) (*AbstractTokenTRC2, error) {
	contract, err := bindAbstractTokenTRC2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenTRC2{AbstractTokenTRC2Caller: AbstractTokenTRC2Caller{contract: contract}, AbstractTokenTRC2Transactor: AbstractTokenTRC2Transactor{contract: contract}, AbstractTokenTRC2Filterer: AbstractTokenTRC2Filterer{contract: contract}}, nil
}

// NewAbstractTokenTRC2Caller creates a new read-only instance of AbstractTokenTRC2, bound to a specific deployed contract.
func NewAbstractTokenTRC2Caller(address common.Address, caller bind.ContractCaller) (*AbstractTokenTRC2Caller, error) {
	contract, err := bindAbstractTokenTRC2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenTRC2Caller{contract: contract}, nil
}

// NewAbstractTokenTRC2Transactor creates a new write-only instance of AbstractTokenTRC2, bound to a specific deployed contract.
func NewAbstractTokenTRC2Transactor(address common.Address, transactor bind.ContractTransactor) (*AbstractTokenTRC2Transactor, error) {
	contract, err := bindAbstractTokenTRC2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenTRC2Transactor{contract: contract}, nil
}

// NewAbstractTokenTRC2Filterer creates a new log filterer instance of AbstractTokenTRC2, bound to a specific deployed contract.
func NewAbstractTokenTRC2Filterer(address common.Address, filterer bind.ContractFilterer) (*AbstractTokenTRC2Filterer, error) {
	contract, err := bindAbstractTokenTRC2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractTokenTRC2Filterer{contract: contract}, nil
}

// bindAbstractTokenTRC2 binds a generic wrapper to an already deployed contract.
func bindAbstractTokenTRC2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractTokenTRC2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractTokenTRC2 *AbstractTokenTRC2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AbstractTokenTRC2.Contract.AbstractTokenTRC2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractTokenTRC2 *AbstractTokenTRC2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractTokenTRC2.Contract.AbstractTokenTRC2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractTokenTRC2 *AbstractTokenTRC2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractTokenTRC2.Contract.AbstractTokenTRC2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractTokenTRC2 *AbstractTokenTRC2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AbstractTokenTRC2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractTokenTRC2 *AbstractTokenTRC2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractTokenTRC2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractTokenTRC2 *AbstractTokenTRC2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractTokenTRC2.Contract.contract.Transact(opts, method, params...)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_AbstractTokenTRC2 *AbstractTokenTRC2Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AbstractTokenTRC2.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_AbstractTokenTRC2 *AbstractTokenTRC2Session) Issuer() (common.Address, error) {
	return _AbstractTokenTRC2.Contract.Issuer(&_AbstractTokenTRC2.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_AbstractTokenTRC2 *AbstractTokenTRC2CallerSession) Issuer() (common.Address, error) {
	return _AbstractTokenTRC2.Contract.Issuer(&_AbstractTokenTRC2.CallOpts)
}

// TRC2IssuerABI is the input ABI used to generate the binding from.
const TRC2IssuerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"minCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenCapacity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"charge\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Apply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"supporter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Charge\",\"type\":\"event\"}]"

// TRC2IssuerBin is the compiled bytecode used for deploying new contracts.
const TRC2IssuerBin = `0x608060405234801561001057600080fd5b506040516020806104578339810160405251600055610423806100346000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633fa615b081146100715780638f3a981c146100985780639d63848a146100b9578063c6b32f341461011e578063fc6bd76a14610134575b600080fd5b34801561007d57600080fd5b50610086610148565b60408051918252519081900360200190f35b3480156100a457600080fd5b50610086600160a060020a036004351661014e565b3480156100c557600080fd5b506100ce610169565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561010a5781810151838201526020016100f2565b505050509050019250505060405180910390f35b610132600160a060020a03600435166101cb565b005b610132600160a060020a036004351661035d565b60005490565b600160a060020a031660009081526002602052604090205490565b606060018054806020026020016040519081016040528092919081815260200182805480156101c157602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116101a3575b5050505050905090565b600081600160a060020a03811615156101e357600080fd5b6000543410156101f257600080fd5b82915033600160a060020a031682600160a060020a0316631d1438486040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561025657600080fd5b505af115801561026a573d6000803e3d6000fd5b505050506040513d602081101561028057600080fd5b5051600160a060020a03161461029557600080fd5b600180548082019091557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03851690811790915560009081526002602052604090205461030390346103de565b600160a060020a0384166000818152600260209081526040918290209390935580513481529051919233927f2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d7169281900390910190a3505050565b600160a060020a038116600090815260026020526040902054610386903463ffffffff6103de16565b600160a060020a0382166000818152600260209081526040918290209390935580513481529051919233927f5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e069281900390910190a350565b6000828201838110156103f057600080fd5b93925050505600a165627a7a7230582005dc9504c7a156980fbaadfe03ffb20a475e65b947f9a8ef3e6d6beee52325a80029`

// DeployTRC2Issuer deploys a new Ethereum contract, binding an instance of TRC2Issuer to it.
func DeployTRC2Issuer(auth *bind.TransactOpts, backend bind.ContractBackend, value *big.Int) (common.Address, *types.Transaction, *TRC2Issuer, error) {
	parsed, err := abi.JSON(strings.NewReader(TRC2IssuerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TRC2IssuerBin), backend, value)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TRC2Issuer{TRC2IssuerCaller: TRC2IssuerCaller{contract: contract}, TRC2IssuerTransactor: TRC2IssuerTransactor{contract: contract}, TRC2IssuerFilterer: TRC2IssuerFilterer{contract: contract}}, nil
}

// TRC2Issuer is an auto generated Go binding around an Ethereum contract.
type TRC2Issuer struct {
	TRC2IssuerCaller     // Read-only binding to the contract
	TRC2IssuerTransactor // Write-only binding to the contract
	TRC2IssuerFilterer   // Log filterer for contract events
}

// TRC2IssuerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TRC2IssuerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TRC2IssuerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TRC2IssuerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TRC2IssuerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TRC2IssuerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TRC2IssuerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TRC2IssuerSession struct {
	Contract     *TRC2Issuer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TRC2IssuerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TRC2IssuerCallerSession struct {
	Contract *TRC2IssuerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TRC2IssuerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TRC2IssuerTransactorSession struct {
	Contract     *TRC2IssuerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TRC2IssuerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TRC2IssuerRaw struct {
	Contract *TRC2Issuer // Generic contract binding to access the raw methods on
}

// TRC2IssuerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TRC2IssuerCallerRaw struct {
	Contract *TRC2IssuerCaller // Generic read-only contract binding to access the raw methods on
}

// TRC2IssuerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TRC2IssuerTransactorRaw struct {
	Contract *TRC2IssuerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTRC2Issuer creates a new instance of TRC2Issuer, bound to a specific deployed contract.
func NewTRC2Issuer(address common.Address, backend bind.ContractBackend) (*TRC2Issuer, error) {
	contract, err := bindTRC2Issuer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TRC2Issuer{TRC2IssuerCaller: TRC2IssuerCaller{contract: contract}, TRC2IssuerTransactor: TRC2IssuerTransactor{contract: contract}, TRC2IssuerFilterer: TRC2IssuerFilterer{contract: contract}}, nil
}

// NewTRC2IssuerCaller creates a new read-only instance of TRC2Issuer, bound to a specific deployed contract.
func NewTRC2IssuerCaller(address common.Address, caller bind.ContractCaller) (*TRC2IssuerCaller, error) {
	contract, err := bindTRC2Issuer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TRC2IssuerCaller{contract: contract}, nil
}

// NewTRC2IssuerTransactor creates a new write-only instance of TRC2Issuer, bound to a specific deployed contract.
func NewTRC2IssuerTransactor(address common.Address, transactor bind.ContractTransactor) (*TRC2IssuerTransactor, error) {
	contract, err := bindTRC2Issuer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TRC2IssuerTransactor{contract: contract}, nil
}

// NewTRC2IssuerFilterer creates a new log filterer instance of TRC2Issuer, bound to a specific deployed contract.
func NewTRC2IssuerFilterer(address common.Address, filterer bind.ContractFilterer) (*TRC2IssuerFilterer, error) {
	contract, err := bindTRC2Issuer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TRC2IssuerFilterer{contract: contract}, nil
}

// bindTRC2Issuer binds a generic wrapper to an already deployed contract.
func bindTRC2Issuer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TRC2IssuerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TRC2Issuer *TRC2IssuerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TRC2Issuer.Contract.TRC2IssuerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TRC2Issuer *TRC2IssuerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TRC2Issuer.Contract.TRC2IssuerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TRC2Issuer *TRC2IssuerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TRC2Issuer.Contract.TRC2IssuerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TRC2Issuer *TRC2IssuerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TRC2Issuer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TRC2Issuer *TRC2IssuerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TRC2Issuer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TRC2Issuer *TRC2IssuerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TRC2Issuer.Contract.contract.Transact(opts, method, params...)
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_TRC2Issuer *TRC2IssuerCaller) GetTokenCapacity(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TRC2Issuer.contract.Call(opts, out, "getTokenCapacity", token)
	return *ret0, err
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_TRC2Issuer *TRC2IssuerSession) GetTokenCapacity(token common.Address) (*big.Int, error) {
	return _TRC2Issuer.Contract.GetTokenCapacity(&_TRC2Issuer.CallOpts, token)
}

// GetTokenCapacity is a free data retrieval call binding the contract method 0x8f3a981c.
//
// Solidity: function getTokenCapacity(token address) constant returns(uint256)
func (_TRC2Issuer *TRC2IssuerCallerSession) GetTokenCapacity(token common.Address) (*big.Int, error) {
	return _TRC2Issuer.Contract.GetTokenCapacity(&_TRC2Issuer.CallOpts, token)
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_TRC2Issuer *TRC2IssuerCaller) MinCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TRC2Issuer.contract.Call(opts, out, "minCap")
	return *ret0, err
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_TRC2Issuer *TRC2IssuerSession) MinCap() (*big.Int, error) {
	return _TRC2Issuer.Contract.MinCap(&_TRC2Issuer.CallOpts)
}

// MinCap is a free data retrieval call binding the contract method 0x3fa615b0.
//
// Solidity: function minCap() constant returns(uint256)
func (_TRC2Issuer *TRC2IssuerCallerSession) MinCap() (*big.Int, error) {
	return _TRC2Issuer.Contract.MinCap(&_TRC2Issuer.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_TRC2Issuer *TRC2IssuerCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TRC2Issuer.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_TRC2Issuer *TRC2IssuerSession) Tokens() ([]common.Address, error) {
	return _TRC2Issuer.Contract.Tokens(&_TRC2Issuer.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_TRC2Issuer *TRC2IssuerCallerSession) Tokens() ([]common.Address, error) {
	return _TRC2Issuer.Contract.Tokens(&_TRC2Issuer.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_TRC2Issuer *TRC2IssuerTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TRC2Issuer.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_TRC2Issuer *TRC2IssuerSession) Apply(token common.Address) (*types.Transaction, error) {
	return _TRC2Issuer.Contract.Apply(&_TRC2Issuer.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_TRC2Issuer *TRC2IssuerTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _TRC2Issuer.Contract.Apply(&_TRC2Issuer.TransactOpts, token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_TRC2Issuer *TRC2IssuerTransactor) Charge(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TRC2Issuer.contract.Transact(opts, "charge", token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_TRC2Issuer *TRC2IssuerSession) Charge(token common.Address) (*types.Transaction, error) {
	return _TRC2Issuer.Contract.Charge(&_TRC2Issuer.TransactOpts, token)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(token address) returns()
func (_TRC2Issuer *TRC2IssuerTransactorSession) Charge(token common.Address) (*types.Transaction, error) {
	return _TRC2Issuer.Contract.Charge(&_TRC2Issuer.TransactOpts, token)
}

// TRC2IssuerApplyIterator is returned from FilterApply and is used to iterate over the raw logs and unpacked data for Apply events raised by the TRC2Issuer contract.
type TRC2IssuerApplyIterator struct {
	Event *TRC2IssuerApply // Event containing the contract specifics and raw log

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
func (it *TRC2IssuerApplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TRC2IssuerApply)
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
		it.Event = new(TRC2IssuerApply)
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
func (it *TRC2IssuerApplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TRC2IssuerApplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TRC2IssuerApply represents a Apply event raised by the TRC2Issuer contract.
type TRC2IssuerApply struct {
	Issuer common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApply is a free log retrieval operation binding the contract event 0x2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d716.
//
// Solidity: event Apply(issuer indexed address, token indexed address, value uint256)
func (_TRC2Issuer *TRC2IssuerFilterer) FilterApply(opts *bind.FilterOpts, issuer []common.Address, token []common.Address) (*TRC2IssuerApplyIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TRC2Issuer.contract.FilterLogs(opts, "Apply", issuerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TRC2IssuerApplyIterator{contract: _TRC2Issuer.contract, event: "Apply", logs: logs, sub: sub}, nil
}

// WatchApply is a free log subscription operation binding the contract event 0x2d485624158277d5113a56388c3abf5c20e3511dd112123ba376d16b21e4d716.
//
// Solidity: event Apply(issuer indexed address, token indexed address, value uint256)
func (_TRC2Issuer *TRC2IssuerFilterer) WatchApply(opts *bind.WatchOpts, sink chan<- *TRC2IssuerApply, issuer []common.Address, token []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TRC2Issuer.contract.WatchLogs(opts, "Apply", issuerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TRC2IssuerApply)
				if err := _TRC2Issuer.contract.UnpackLog(event, "Apply", log); err != nil {
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

// TRC2IssuerChargeIterator is returned from FilterCharge and is used to iterate over the raw logs and unpacked data for Charge events raised by the TRC2Issuer contract.
type TRC2IssuerChargeIterator struct {
	Event *TRC2IssuerCharge // Event containing the contract specifics and raw log

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
func (it *TRC2IssuerChargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TRC2IssuerCharge)
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
		it.Event = new(TRC2IssuerCharge)
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
func (it *TRC2IssuerChargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TRC2IssuerChargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TRC2IssuerCharge represents a Charge event raised by the TRC2Issuer contract.
type TRC2IssuerCharge struct {
	Supporter common.Address
	Token     common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCharge is a free log retrieval operation binding the contract event 0x5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e06.
//
// Solidity: event Charge(supporter indexed address, token indexed address, value uint256)
func (_TRC2Issuer *TRC2IssuerFilterer) FilterCharge(opts *bind.FilterOpts, supporter []common.Address, token []common.Address) (*TRC2IssuerChargeIterator, error) {

	var supporterRule []interface{}
	for _, supporterItem := range supporter {
		supporterRule = append(supporterRule, supporterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TRC2Issuer.contract.FilterLogs(opts, "Charge", supporterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TRC2IssuerChargeIterator{contract: _TRC2Issuer.contract, event: "Charge", logs: logs, sub: sub}, nil
}

// WatchCharge is a free log subscription operation binding the contract event 0x5cffac866325fd9b2a8ea8df2f110a0058313b279402d15ae28dd324a2282e06.
//
// Solidity: event Charge(supporter indexed address, token indexed address, value uint256)
func (_TRC2Issuer *TRC2IssuerFilterer) WatchCharge(opts *bind.WatchOpts, sink chan<- *TRC2IssuerCharge, supporter []common.Address, token []common.Address) (event.Subscription, error) {

	var supporterRule []interface{}
	for _, supporterItem := range supporter {
		supporterRule = append(supporterRule, supporterItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TRC2Issuer.contract.WatchLogs(opts, "Charge", supporterRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TRC2IssuerCharge)
				if err := _TRC2Issuer.contract.UnpackLog(event, "Charge", log); err != nil {
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
