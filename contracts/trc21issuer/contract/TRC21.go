// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/tao2-core/taoblockchain"
	"github.com/tao2-core/tao2-core/accounts/abi"
	"github.com/tao2-core/tao2-core/accounts/abi/bind"
	"github.com/tao2-core/tao2-core/common"
	"github.com/tao2-core/tao2-core/core/types"
	"github.com/tao2-core/tao2-core/event"
)

// ITRC2ABI is the input ABI used to generate the binding from.
const ITRC2ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// ITRC2Bin is the compiled bytecode used for deploying new contracts.
const ITRC2Bin = `0x`

// DeployITRC2 deploys a new Ethereum contract, binding an instance of ITRC2 to it.
func DeployITRC2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ITRC2, error) {
	parsed, err := abi.JSON(strings.NewReader(ITRC2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ITRC2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ITRC2{ITRC2Caller: ITRC2Caller{contract: contract}, ITRC2Transactor: ITRC2Transactor{contract: contract}, ITRC2Filterer: ITRC2Filterer{contract: contract}}, nil
}

// ITRC2 is an auto generated Go binding around an Ethereum contract.
type ITRC2 struct {
	ITRC2Caller     // Read-only binding to the contract
	ITRC2Transactor // Write-only binding to the contract
	ITRC2Filterer   // Log filterer for contract events
}

// ITRC2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ITRC2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITRC2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ITRC2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITRC2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITRC2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITRC2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITRC2Session struct {
	Contract     *ITRC2           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITRC2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITRC2CallerSession struct {
	Contract *ITRC2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITRC2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITRC2TransactorSession struct {
	Contract     *ITRC2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITRC2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ITRC2Raw struct {
	Contract *ITRC2 // Generic contract binding to access the raw methods on
}

// ITRC2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITRC2CallerRaw struct {
	Contract *ITRC2Caller // Generic read-only contract binding to access the raw methods on
}

// ITRC2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITRC2TransactorRaw struct {
	Contract *ITRC2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewITRC2 creates a new instance of ITRC2, bound to a specific deployed contract.
func NewITRC2(address common.Address, backend bind.ContractBackend) (*ITRC2, error) {
	contract, err := bindITRC2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITRC2{ITRC2Caller: ITRC2Caller{contract: contract}, ITRC2Transactor: ITRC2Transactor{contract: contract}, ITRC2Filterer: ITRC2Filterer{contract: contract}}, nil
}

// NewITRC2Caller creates a new read-only instance of ITRC2, bound to a specific deployed contract.
func NewITRC2Caller(address common.Address, caller bind.ContractCaller) (*ITRC2Caller, error) {
	contract, err := bindITRC2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITRC2Caller{contract: contract}, nil
}

// NewITRC2Transactor creates a new write-only instance of ITRC2, bound to a specific deployed contract.
func NewITRC2Transactor(address common.Address, transactor bind.ContractTransactor) (*ITRC2Transactor, error) {
	contract, err := bindITRC2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITRC2Transactor{contract: contract}, nil
}

// NewITRC2Filterer creates a new log filterer instance of ITRC2, bound to a specific deployed contract.
func NewITRC2Filterer(address common.Address, filterer bind.ContractFilterer) (*ITRC2Filterer, error) {
	contract, err := bindITRC2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITRC2Filterer{contract: contract}, nil
}

// bindITRC2 binds a generic wrapper to an already deployed contract.
func bindITRC2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITRC2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITRC2 *ITRC2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITRC2.Contract.ITRC2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITRC2 *ITRC2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITRC2.Contract.ITRC2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITRC2 *ITRC2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITRC2.Contract.ITRC2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITRC2 *ITRC2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITRC2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITRC2 *ITRC2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITRC2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITRC2 *ITRC2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITRC2.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ITRC2 *ITRC2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ITRC2.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ITRC2 *ITRC2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ITRC2.Contract.Allowance(&_ITRC2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ITRC2 *ITRC2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ITRC2.Contract.Allowance(&_ITRC2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ITRC2 *ITRC2Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ITRC2.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ITRC2 *ITRC2Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ITRC2.Contract.BalanceOf(&_ITRC2.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ITRC2 *ITRC2CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ITRC2.Contract.BalanceOf(&_ITRC2.CallOpts, who)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_ITRC2 *ITRC2Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ITRC2.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_ITRC2 *ITRC2Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _ITRC2.Contract.EstimateFee(&_ITRC2.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_ITRC2 *ITRC2CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _ITRC2.Contract.EstimateFee(&_ITRC2.CallOpts, value)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ITRC2 *ITRC2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ITRC2.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ITRC2 *ITRC2Session) TotalSupply() (*big.Int, error) {
	return _ITRC2.Contract.TotalSupply(&_ITRC2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ITRC2 *ITRC2CallerSession) TotalSupply() (*big.Int, error) {
	return _ITRC2.Contract.TotalSupply(&_ITRC2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ITRC2 *ITRC2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ITRC2.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ITRC2 *ITRC2Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ITRC2.Contract.Approve(&_ITRC2.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ITRC2 *ITRC2TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ITRC2.Contract.Approve(&_ITRC2.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ITRC2 *ITRC2Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ITRC2.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ITRC2 *ITRC2Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ITRC2.Contract.Transfer(&_ITRC2.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ITRC2 *ITRC2TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ITRC2.Contract.Transfer(&_ITRC2.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ITRC2 *ITRC2Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ITRC2.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ITRC2 *ITRC2Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ITRC2.Contract.TransferFrom(&_ITRC2.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ITRC2 *ITRC2TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ITRC2.Contract.TransferFrom(&_ITRC2.TransactOpts, from, to, value)
}

// ITRC2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ITRC2 contract.
type ITRC2ApprovalIterator struct {
	Event *ITRC2Approval // Event containing the contract specifics and raw log

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
func (it *ITRC2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITRC2Approval)
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
		it.Event = new(ITRC2Approval)
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
func (it *ITRC2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITRC2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITRC2Approval represents a Approval event raised by the ITRC2 contract.
type ITRC2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ITRC2 *ITRC2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ITRC2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ITRC2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ITRC2ApprovalIterator{contract: _ITRC2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ITRC2 *ITRC2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ITRC2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ITRC2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITRC2Approval)
				if err := _ITRC2.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ITRC2FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the ITRC2 contract.
type ITRC2FeeIterator struct {
	Event *ITRC2Fee // Event containing the contract specifics and raw log

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
func (it *ITRC2FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITRC2Fee)
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
		it.Event = new(ITRC2Fee)
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
func (it *ITRC2FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITRC2FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITRC2Fee represents a Fee event raised by the ITRC2 contract.
type ITRC2Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_ITRC2 *ITRC2Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*ITRC2FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ITRC2.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ITRC2FeeIterator{contract: _ITRC2.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_ITRC2 *ITRC2Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *ITRC2Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ITRC2.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITRC2Fee)
				if err := _ITRC2.contract.UnpackLog(event, "Fee", log); err != nil {
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

// ITRC2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ITRC2 contract.
type ITRC2TransferIterator struct {
	Event *ITRC2Transfer // Event containing the contract specifics and raw log

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
func (it *ITRC2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITRC2Transfer)
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
		it.Event = new(ITRC2Transfer)
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
func (it *ITRC2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITRC2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITRC2Transfer represents a Transfer event raised by the ITRC2 contract.
type ITRC2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ITRC2 *ITRC2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ITRC2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ITRC2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ITRC2TransferIterator{contract: _ITRC2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ITRC2 *ITRC2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ITRC2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ITRC2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITRC2Transfer)
				if err := _ITRC2.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// MyTRC2ABI is the input ABI used to generate the binding from.
const MyTRC2ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"cap\",\"type\":\"uint256\"},{\"name\":\"minFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// MyTRC2Bin is the compiled bytecode used for deploying new contracts.
const MyTRC2Bin = `0x60806040523480156200001157600080fd5b5060405162000b3138038062000b3183398101604090815281516020808401519284015160608501516080860151938601805190969590950194919390929091620000639160059190880190620001e7565b50835162000079906006906020870190620001e7565b506007805460ff191660ff85161790556200009e3383640100000000620000d1810204565b620000b23364010000000062000190810204565b620000c681640100000000620001c8810204565b50505050506200028c565b600160a060020a0382161515620000e757600080fd5b60045462000104908264010000000062000842620001cd82021704565b600455600160a060020a0382166000908152602081905260409020546200013a908264010000000062000842620001cd82021704565b600160a060020a0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600160a060020a0381161515620001a657600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600155565b600082820183811015620001e057600080fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022a57805160ff19168380011785556200025a565b828001600101855582156200025a579182015b828111156200025a5782518255916020019190600101906200023d565b50620002689291506200026c565b5090565b6200028991905b8082111562000268576000815560010162000273565b90565b610895806200029c6000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100be578063095ea7b314610148578063127e8e4d1461018057806318160ddd146101aa5780631d143848146101bf57806323b872dd146101f057806324ec75901461021a578063313ce5671461022f57806370a082311461025a57806395d89b411461027b578063a9059cbb14610290578063dd62ed3e146102b4575b600080fd5b3480156100ca57600080fd5b506100d36102db565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010d5781810151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015457600080fd5b5061016c600160a060020a0360043516602435610371565b604080519115158252519081900360200190f35b34801561018c57600080fd5b5061019860043561042b565b60408051918252519081900360200190f35b3480156101b657600080fd5b50610198610457565b3480156101cb57600080fd5b506101d461045d565b60408051600160a060020a039092168252519081900360200190f35b3480156101fc57600080fd5b5061016c600160a060020a036004358116906024351660443561046c565b34801561022657600080fd5b506101986105ac565b34801561023b57600080fd5b506102446105b2565b6040805160ff9092168252519081900360200190f35b34801561026657600080fd5b50610198600160a060020a03600435166105bb565b34801561028757600080fd5b506100d36105d6565b34801561029c57600080fd5b5061016c600160a060020a0360043516602435610637565b3480156102c057600080fd5b50610198600160a060020a03600435811690602435166106f0565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103675780601f1061033c57610100808354040283529160200191610367565b820191906000526020600020905b81548152906001019060200180831161034a57829003601f168201915b5050505050905090565b6000600160a060020a038316151561038857600080fd5b6001543360009081526020819052604090205410156103a657600080fd5b336000818152600360209081526040808320600160a060020a03888116855292529091208490556002546001546103e29392919091169061071b565b604080518381529051600160a060020a0385169133917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a350600192915050565b60015460009061045190610445848463ffffffff61080d16565b9063ffffffff61084216565b92915050565b60045490565b600254600160a060020a031690565b6000806104846001548461084290919063ffffffff16565b9050600160a060020a038416151561049b57600080fd5b808311156104a857600080fd5b600160a060020a03851660009081526003602090815260408083203384529091529020548111156104d857600080fd5b600160a060020a038516600090815260036020908152604080832033845290915290205461050c908263ffffffff61085416565b600160a060020a038616600090815260036020908152604080832033845290915290205561053b85858561071b565b600254600154610558918791600160a060020a039091169061071b565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4506001949350505050565b60015490565b60075460ff1690565b600160a060020a031660009081526020819052604090205490565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103675780601f1061033c57610100808354040283529160200191610367565b60008061064f6001548461084290919063ffffffff16565b9050600160a060020a038416151561066657600080fd5b8083111561067357600080fd5b61067e33858561071b565b60025460015461069b913391600160a060020a039091169061071b565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4600191505b5092915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600160a060020a03831660009081526020819052604090205481111561074057600080fd5b600160a060020a038216151561075557600080fd5b600160a060020a03831660009081526020819052604090205461077e908263ffffffff61085416565b600160a060020a0380851660009081526020819052604080822093909355908416815220546107b3908263ffffffff61084216565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008083151561082057600091506106e9565b5082820282848281151561083057fe5b041461083b57600080fd5b9392505050565b60008282018381101561083b57600080fd5b60008282111561086357600080fd5b509003905600a165627a7a72305820fafa860f05b0db36e3b8067a75d9068ddc0c82bea68b125e228ef1a02a3a62140029`

// DeployMyTRC2 deploys a new Ethereum contract, binding an instance of MyTRC2 to it.
func DeployMyTRC2(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8, cap *big.Int, minFee *big.Int) (common.Address, *types.Transaction, *MyTRC2, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTRC2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyTRC2Bin), backend, name, symbol, decimals, cap, minFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyTRC2{MyTRC2Caller: MyTRC2Caller{contract: contract}, MyTRC2Transactor: MyTRC2Transactor{contract: contract}, MyTRC2Filterer: MyTRC2Filterer{contract: contract}}, nil
}

// MyTRC2 is an auto generated Go binding around an Ethereum contract.
type MyTRC2 struct {
	MyTRC2Caller     // Read-only binding to the contract
	MyTRC2Transactor // Write-only binding to the contract
	MyTRC2Filterer   // Log filterer for contract events
}

// MyTRC2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyTRC2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTRC2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyTRC2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTRC2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyTRC2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTRC2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyTRC2Session struct {
	Contract     *MyTRC2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyTRC2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyTRC2CallerSession struct {
	Contract *MyTRC2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyTRC2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyTRC2TransactorSession struct {
	Contract     *MyTRC2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyTRC2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyTRC2Raw struct {
	Contract *MyTRC2 // Generic contract binding to access the raw methods on
}

// MyTRC2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyTRC2CallerRaw struct {
	Contract *MyTRC2Caller // Generic read-only contract binding to access the raw methods on
}

// MyTRC2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyTRC2TransactorRaw struct {
	Contract *MyTRC2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyTRC2 creates a new instance of MyTRC2, bound to a specific deployed contract.
func NewMyTRC2(address common.Address, backend bind.ContractBackend) (*MyTRC2, error) {
	contract, err := bindMyTRC2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyTRC2{MyTRC2Caller: MyTRC2Caller{contract: contract}, MyTRC2Transactor: MyTRC2Transactor{contract: contract}, MyTRC2Filterer: MyTRC2Filterer{contract: contract}}, nil
}

// NewMyTRC2Caller creates a new read-only instance of MyTRC2, bound to a specific deployed contract.
func NewMyTRC2Caller(address common.Address, caller bind.ContractCaller) (*MyTRC2Caller, error) {
	contract, err := bindMyTRC2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyTRC2Caller{contract: contract}, nil
}

// NewMyTRC2Transactor creates a new write-only instance of MyTRC2, bound to a specific deployed contract.
func NewMyTRC2Transactor(address common.Address, transactor bind.ContractTransactor) (*MyTRC2Transactor, error) {
	contract, err := bindMyTRC2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyTRC2Transactor{contract: contract}, nil
}

// NewMyTRC2Filterer creates a new log filterer instance of MyTRC2, bound to a specific deployed contract.
func NewMyTRC2Filterer(address common.Address, filterer bind.ContractFilterer) (*MyTRC2Filterer, error) {
	contract, err := bindMyTRC2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyTRC2Filterer{contract: contract}, nil
}

// bindMyTRC2 binds a generic wrapper to an already deployed contract.
func bindMyTRC2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTRC2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyTRC2 *MyTRC2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyTRC2.Contract.MyTRC2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyTRC2 *MyTRC2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyTRC2.Contract.MyTRC2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyTRC2 *MyTRC2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyTRC2.Contract.MyTRC2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyTRC2 *MyTRC2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyTRC2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyTRC2 *MyTRC2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyTRC2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyTRC2 *MyTRC2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyTRC2.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyTRC2 *MyTRC2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyTRC2.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyTRC2 *MyTRC2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyTRC2.Contract.Allowance(&_MyTRC2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyTRC2 *MyTRC2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyTRC2.Contract.Allowance(&_MyTRC2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyTRC2 *MyTRC2Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyTRC2.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyTRC2 *MyTRC2Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyTRC2.Contract.BalanceOf(&_MyTRC2.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyTRC2 *MyTRC2CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyTRC2.Contract.BalanceOf(&_MyTRC2.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyTRC2 *MyTRC2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MyTRC2.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyTRC2 *MyTRC2Session) Decimals() (uint8, error) {
	return _MyTRC2.Contract.Decimals(&_MyTRC2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyTRC2 *MyTRC2CallerSession) Decimals() (uint8, error) {
	return _MyTRC2.Contract.Decimals(&_MyTRC2.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyTRC2 *MyTRC2Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyTRC2.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyTRC2 *MyTRC2Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _MyTRC2.Contract.EstimateFee(&_MyTRC2.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyTRC2 *MyTRC2CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _MyTRC2.Contract.EstimateFee(&_MyTRC2.CallOpts, value)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyTRC2 *MyTRC2Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyTRC2.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyTRC2 *MyTRC2Session) Issuer() (common.Address, error) {
	return _MyTRC2.Contract.Issuer(&_MyTRC2.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyTRC2 *MyTRC2CallerSession) Issuer() (common.Address, error) {
	return _MyTRC2.Contract.Issuer(&_MyTRC2.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyTRC2 *MyTRC2Caller) MinFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyTRC2.contract.Call(opts, out, "minFee")
	return *ret0, err
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyTRC2 *MyTRC2Session) MinFee() (*big.Int, error) {
	return _MyTRC2.Contract.MinFee(&_MyTRC2.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyTRC2 *MyTRC2CallerSession) MinFee() (*big.Int, error) {
	return _MyTRC2.Contract.MinFee(&_MyTRC2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyTRC2 *MyTRC2Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyTRC2.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyTRC2 *MyTRC2Session) Name() (string, error) {
	return _MyTRC2.Contract.Name(&_MyTRC2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyTRC2 *MyTRC2CallerSession) Name() (string, error) {
	return _MyTRC2.Contract.Name(&_MyTRC2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyTRC2 *MyTRC2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyTRC2.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyTRC2 *MyTRC2Session) Symbol() (string, error) {
	return _MyTRC2.Contract.Symbol(&_MyTRC2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyTRC2 *MyTRC2CallerSession) Symbol() (string, error) {
	return _MyTRC2.Contract.Symbol(&_MyTRC2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyTRC2 *MyTRC2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyTRC2.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyTRC2 *MyTRC2Session) TotalSupply() (*big.Int, error) {
	return _MyTRC2.Contract.TotalSupply(&_MyTRC2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyTRC2 *MyTRC2CallerSession) TotalSupply() (*big.Int, error) {
	return _MyTRC2.Contract.TotalSupply(&_MyTRC2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyTRC2 *MyTRC2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyTRC2.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyTRC2 *MyTRC2Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyTRC2.Contract.Approve(&_MyTRC2.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyTRC2 *MyTRC2TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyTRC2.Contract.Approve(&_MyTRC2.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyTRC2 *MyTRC2Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyTRC2.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyTRC2 *MyTRC2Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyTRC2.Contract.Transfer(&_MyTRC2.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyTRC2 *MyTRC2TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyTRC2.Contract.Transfer(&_MyTRC2.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyTRC2 *MyTRC2Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyTRC2.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyTRC2 *MyTRC2Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyTRC2.Contract.TransferFrom(&_MyTRC2.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyTRC2 *MyTRC2TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyTRC2.Contract.TransferFrom(&_MyTRC2.TransactOpts, from, to, value)
}

// MyTRC2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyTRC2 contract.
type MyTRC2ApprovalIterator struct {
	Event *MyTRC2Approval // Event containing the contract specifics and raw log

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
func (it *MyTRC2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTRC2Approval)
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
		it.Event = new(MyTRC2Approval)
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
func (it *MyTRC2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTRC2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTRC2Approval represents a Approval event raised by the MyTRC2 contract.
type MyTRC2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_MyTRC2 *MyTRC2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MyTRC2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyTRC2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MyTRC2ApprovalIterator{contract: _MyTRC2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_MyTRC2 *MyTRC2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyTRC2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyTRC2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTRC2Approval)
				if err := _MyTRC2.contract.UnpackLog(event, "Approval", log); err != nil {
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

// MyTRC2FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the MyTRC2 contract.
type MyTRC2FeeIterator struct {
	Event *MyTRC2Fee // Event containing the contract specifics and raw log

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
func (it *MyTRC2FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTRC2Fee)
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
		it.Event = new(MyTRC2Fee)
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
func (it *MyTRC2FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTRC2FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTRC2Fee represents a Fee event raised by the MyTRC2 contract.
type MyTRC2Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_MyTRC2 *MyTRC2Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*MyTRC2FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _MyTRC2.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &MyTRC2FeeIterator{contract: _MyTRC2.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_MyTRC2 *MyTRC2Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *MyTRC2Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _MyTRC2.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTRC2Fee)
				if err := _MyTRC2.contract.UnpackLog(event, "Fee", log); err != nil {
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

// MyTRC2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyTRC2 contract.
type MyTRC2TransferIterator struct {
	Event *MyTRC2Transfer // Event containing the contract specifics and raw log

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
func (it *MyTRC2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTRC2Transfer)
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
		it.Event = new(MyTRC2Transfer)
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
func (it *MyTRC2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTRC2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTRC2Transfer represents a Transfer event raised by the MyTRC2 contract.
type MyTRC2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyTRC2 *MyTRC2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyTRC2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyTRC2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyTRC2TransferIterator{contract: _MyTRC2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyTRC2 *MyTRC2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyTRC2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyTRC2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTRC2Transfer)
				if err := _MyTRC2.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// TRC2ABI is the input ABI used to generate the binding from.
const TRC2ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// TRC2Bin is the compiled bytecode used for deploying new contracts.
const TRC2Bin = `0x608060405234801561001057600080fd5b506106aa806100206000396000f3006080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009d578063127e8e4d146100d557806318160ddd146100ff5780631d1438481461011457806323b872dd1461014557806324ec75901461016f57806370a0823114610184578063a9059cbb146101a5578063dd62ed3e146101c9575b600080fd5b3480156100a957600080fd5b506100c1600160a060020a03600435166024356101f0565b604080519115158252519081900360200190f35b3480156100e157600080fd5b506100ed6004356102aa565b60408051918252519081900360200190f35b34801561010b57600080fd5b506100ed6102d6565b34801561012057600080fd5b506101296102dc565b60408051600160a060020a039092168252519081900360200190f35b34801561015157600080fd5b506100c1600160a060020a03600435811690602435166044356102eb565b34801561017b57600080fd5b506100ed61042b565b34801561019057600080fd5b506100ed600160a060020a0360043516610431565b3480156101b157600080fd5b506100c1600160a060020a036004351660243561044c565b3480156101d557600080fd5b506100ed600160a060020a0360043581169060243516610505565b6000600160a060020a038316151561020757600080fd5b60015433600090815260208190526040902054101561022557600080fd5b336000818152600360209081526040808320600160a060020a038881168552925290912084905560025460015461026193929190911690610530565b604080518381529051600160a060020a0385169133917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a350600192915050565b6001546000906102d0906102c4848463ffffffff61062216565b9063ffffffff61065716565b92915050565b60045490565b600254600160a060020a031690565b6000806103036001548461065790919063ffffffff16565b9050600160a060020a038416151561031a57600080fd5b8083111561032757600080fd5b600160a060020a038516600090815260036020908152604080832033845290915290205481111561035757600080fd5b600160a060020a038516600090815260036020908152604080832033845290915290205461038b908263ffffffff61066916565b600160a060020a03861660009081526003602090815260408083203384529091529020556103ba858585610530565b6002546001546103d7918791600160a060020a0390911690610530565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4506001949350505050565b60015490565b600160a060020a031660009081526020819052604090205490565b6000806104646001548461065790919063ffffffff16565b9050600160a060020a038416151561047b57600080fd5b8083111561048857600080fd5b610493338585610530565b6002546001546104b0913391600160a060020a0390911690610530565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4600191505b5092915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600160a060020a03831660009081526020819052604090205481111561055557600080fd5b600160a060020a038216151561056a57600080fd5b600160a060020a038316600090815260208190526040902054610593908263ffffffff61066916565b600160a060020a0380851660009081526020819052604080822093909355908416815220546105c8908263ffffffff61065716565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008083151561063557600091506104fe565b5082820282848281151561064557fe5b041461065057600080fd5b9392505050565b60008282018381101561065057600080fd5b60008282111561067857600080fd5b509003905600a165627a7a72305820a40ecfbb2241cf25b0b2e3e120bec997e2d61dc709b4b92c70939c9b604b417f0029`

// DeployTRC2 deploys a new Ethereum contract, binding an instance of TRC2 to it.
func DeployTRC2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TRC2, error) {
	parsed, err := abi.JSON(strings.NewReader(TRC2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TRC2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TRC2{TRC2Caller: TRC2Caller{contract: contract}, TRC2Transactor: TRC2Transactor{contract: contract}, TRC2Filterer: TRC2Filterer{contract: contract}}, nil
}

// TRC2 is an auto generated Go binding around an Ethereum contract.
type TRC2 struct {
	TRC2Caller     // Read-only binding to the contract
	TRC2Transactor // Write-only binding to the contract
	TRC2Filterer   // Log filterer for contract events
}

// TRC2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TRC2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TRC2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TRC2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TRC2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TRC2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TRC2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TRC2Session struct {
	Contract     *TRC2            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TRC2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TRC2CallerSession struct {
	Contract *TRC2Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TRC2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TRC2TransactorSession struct {
	Contract     *TRC2Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TRC2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TRC2Raw struct {
	Contract *TRC2 // Generic contract binding to access the raw methods on
}

// TRC2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TRC2CallerRaw struct {
	Contract *TRC2Caller // Generic read-only contract binding to access the raw methods on
}

// TRC2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TRC2TransactorRaw struct {
	Contract *TRC2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTRC2 creates a new instance of TRC2, bound to a specific deployed contract.
func NewTRC2(address common.Address, backend bind.ContractBackend) (*TRC2, error) {
	contract, err := bindTRC2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TRC2{TRC2Caller: TRC2Caller{contract: contract}, TRC2Transactor: TRC2Transactor{contract: contract}, TRC2Filterer: TRC2Filterer{contract: contract}}, nil
}

// NewTRC2Caller creates a new read-only instance of TRC2, bound to a specific deployed contract.
func NewTRC2Caller(address common.Address, caller bind.ContractCaller) (*TRC2Caller, error) {
	contract, err := bindTRC2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TRC2Caller{contract: contract}, nil
}

// NewTRC2Transactor creates a new write-only instance of TRC2, bound to a specific deployed contract.
func NewTRC2Transactor(address common.Address, transactor bind.ContractTransactor) (*TRC2Transactor, error) {
	contract, err := bindTRC2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TRC2Transactor{contract: contract}, nil
}

// NewTRC2Filterer creates a new log filterer instance of TRC2, bound to a specific deployed contract.
func NewTRC2Filterer(address common.Address, filterer bind.ContractFilterer) (*TRC2Filterer, error) {
	contract, err := bindTRC2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TRC2Filterer{contract: contract}, nil
}

// bindTRC2 binds a generic wrapper to an already deployed contract.
func bindTRC2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TRC2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TRC2 *TRC2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TRC2.Contract.TRC2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TRC2 *TRC2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TRC2.Contract.TRC2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TRC2 *TRC2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TRC2.Contract.TRC2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TRC2 *TRC2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TRC2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TRC2 *TRC2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TRC2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TRC2 *TRC2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TRC2.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_TRC2 *TRC2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TRC2.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_TRC2 *TRC2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TRC2.Contract.Allowance(&_TRC2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_TRC2 *TRC2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TRC2.Contract.Allowance(&_TRC2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_TRC2 *TRC2Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TRC2.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_TRC2 *TRC2Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TRC2.Contract.BalanceOf(&_TRC2.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_TRC2 *TRC2CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TRC2.Contract.BalanceOf(&_TRC2.CallOpts, owner)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_TRC2 *TRC2Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TRC2.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_TRC2 *TRC2Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _TRC2.Contract.EstimateFee(&_TRC2.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_TRC2 *TRC2CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _TRC2.Contract.EstimateFee(&_TRC2.CallOpts, value)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_TRC2 *TRC2Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TRC2.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_TRC2 *TRC2Session) Issuer() (common.Address, error) {
	return _TRC2.Contract.Issuer(&_TRC2.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_TRC2 *TRC2CallerSession) Issuer() (common.Address, error) {
	return _TRC2.Contract.Issuer(&_TRC2.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_TRC2 *TRC2Caller) MinFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TRC2.contract.Call(opts, out, "minFee")
	return *ret0, err
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_TRC2 *TRC2Session) MinFee() (*big.Int, error) {
	return _TRC2.Contract.MinFee(&_TRC2.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_TRC2 *TRC2CallerSession) MinFee() (*big.Int, error) {
	return _TRC2.Contract.MinFee(&_TRC2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TRC2 *TRC2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TRC2.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TRC2 *TRC2Session) TotalSupply() (*big.Int, error) {
	return _TRC2.Contract.TotalSupply(&_TRC2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TRC2 *TRC2CallerSession) TotalSupply() (*big.Int, error) {
	return _TRC2.Contract.TotalSupply(&_TRC2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_TRC2 *TRC2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TRC2.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_TRC2 *TRC2Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TRC2.Contract.Approve(&_TRC2.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_TRC2 *TRC2TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TRC2.Contract.Approve(&_TRC2.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_TRC2 *TRC2Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TRC2.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_TRC2 *TRC2Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TRC2.Contract.Transfer(&_TRC2.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_TRC2 *TRC2TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TRC2.Contract.Transfer(&_TRC2.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_TRC2 *TRC2Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TRC2.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_TRC2 *TRC2Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TRC2.Contract.TransferFrom(&_TRC2.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_TRC2 *TRC2TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TRC2.Contract.TransferFrom(&_TRC2.TransactOpts, from, to, value)
}

// TRC2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TRC2 contract.
type TRC2ApprovalIterator struct {
	Event *TRC2Approval // Event containing the contract specifics and raw log

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
func (it *TRC2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TRC2Approval)
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
		it.Event = new(TRC2Approval)
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
func (it *TRC2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TRC2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TRC2Approval represents a Approval event raised by the TRC2 contract.
type TRC2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_TRC2 *TRC2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TRC2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TRC2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TRC2ApprovalIterator{contract: _TRC2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_TRC2 *TRC2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TRC2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TRC2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TRC2Approval)
				if err := _TRC2.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TRC2FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the TRC2 contract.
type TRC2FeeIterator struct {
	Event *TRC2Fee // Event containing the contract specifics and raw log

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
func (it *TRC2FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TRC2Fee)
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
		it.Event = new(TRC2Fee)
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
func (it *TRC2FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TRC2FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TRC2Fee represents a Fee event raised by the TRC2 contract.
type TRC2Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_TRC2 *TRC2Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*TRC2FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _TRC2.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &TRC2FeeIterator{contract: _TRC2.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_TRC2 *TRC2Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *TRC2Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _TRC2.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TRC2Fee)
				if err := _TRC2.contract.UnpackLog(event, "Fee", log); err != nil {
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

// TRC2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TRC2 contract.
type TRC2TransferIterator struct {
	Event *TRC2Transfer // Event containing the contract specifics and raw log

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
func (it *TRC2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TRC2Transfer)
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
		it.Event = new(TRC2Transfer)
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
func (it *TRC2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TRC2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TRC2Transfer represents a Transfer event raised by the TRC2 contract.
type TRC2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_TRC2 *TRC2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TRC2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TRC2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TRC2TransferIterator{contract: _TRC2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_TRC2 *TRC2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TRC2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TRC2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TRC2Transfer)
				if err := _TRC2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
