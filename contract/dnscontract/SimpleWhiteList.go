// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnscontract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SimpleWhiteListABI is the input ABI used to generate the binding from.
const SimpleWhiteListABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"EvAddWhiteAddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"EvUpdateWhiteAddr\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"AddMultiWhiteAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"op\",\"type\":\"address\"}],\"name\":\"AddOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"level_\",\"type\":\"uint8\"}],\"name\":\"AddWhiteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"LevelAddr\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"level_\",\"type\":\"uint8\"}],\"name\":\"UpdateWhiteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"cnt_\",\"type\":\"uint8\"}],\"name\":\"setDefaultLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SimpleWhiteList is an auto generated Go binding around an Ethereum contract.
type SimpleWhiteList struct {
	SimpleWhiteListCaller     // Read-only binding to the contract
	SimpleWhiteListTransactor // Write-only binding to the contract
	SimpleWhiteListFilterer   // Log filterer for contract events
}

// SimpleWhiteListCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleWhiteListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleWhiteListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleWhiteListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleWhiteListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleWhiteListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleWhiteListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleWhiteListSession struct {
	Contract     *SimpleWhiteList  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleWhiteListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleWhiteListCallerSession struct {
	Contract *SimpleWhiteListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SimpleWhiteListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleWhiteListTransactorSession struct {
	Contract     *SimpleWhiteListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SimpleWhiteListRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleWhiteListRaw struct {
	Contract *SimpleWhiteList // Generic contract binding to access the raw methods on
}

// SimpleWhiteListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleWhiteListCallerRaw struct {
	Contract *SimpleWhiteListCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleWhiteListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleWhiteListTransactorRaw struct {
	Contract *SimpleWhiteListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleWhiteList creates a new instance of SimpleWhiteList, bound to a specific deployed contract.
func NewSimpleWhiteList(address common.Address, backend bind.ContractBackend) (*SimpleWhiteList, error) {
	contract, err := bindSimpleWhiteList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleWhiteList{SimpleWhiteListCaller: SimpleWhiteListCaller{contract: contract}, SimpleWhiteListTransactor: SimpleWhiteListTransactor{contract: contract}, SimpleWhiteListFilterer: SimpleWhiteListFilterer{contract: contract}}, nil
}

// NewSimpleWhiteListCaller creates a new read-only instance of SimpleWhiteList, bound to a specific deployed contract.
func NewSimpleWhiteListCaller(address common.Address, caller bind.ContractCaller) (*SimpleWhiteListCaller, error) {
	contract, err := bindSimpleWhiteList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleWhiteListCaller{contract: contract}, nil
}

// NewSimpleWhiteListTransactor creates a new write-only instance of SimpleWhiteList, bound to a specific deployed contract.
func NewSimpleWhiteListTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleWhiteListTransactor, error) {
	contract, err := bindSimpleWhiteList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleWhiteListTransactor{contract: contract}, nil
}

// NewSimpleWhiteListFilterer creates a new log filterer instance of SimpleWhiteList, bound to a specific deployed contract.
func NewSimpleWhiteListFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleWhiteListFilterer, error) {
	contract, err := bindSimpleWhiteList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleWhiteListFilterer{contract: contract}, nil
}

// bindSimpleWhiteList binds a generic wrapper to an already deployed contract.
func bindSimpleWhiteList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleWhiteListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleWhiteList *SimpleWhiteListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleWhiteList.Contract.SimpleWhiteListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleWhiteList *SimpleWhiteListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.SimpleWhiteListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleWhiteList *SimpleWhiteListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.SimpleWhiteListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleWhiteList *SimpleWhiteListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleWhiteList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleWhiteList *SimpleWhiteListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleWhiteList *SimpleWhiteListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.contract.Transact(opts, method, params...)
}

// LevelAddr is a free data retrieval call binding the contract method 0x5767c0e1.
//
// Solidity: function LevelAddr(address user_) view returns(uint8)
func (_SimpleWhiteList *SimpleWhiteListCaller) LevelAddr(opts *bind.CallOpts, user_ common.Address) (uint8, error) {
	var out []interface{}
	err := _SimpleWhiteList.contract.Call(opts, &out, "LevelAddr", user_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LevelAddr is a free data retrieval call binding the contract method 0x5767c0e1.
//
// Solidity: function LevelAddr(address user_) view returns(uint8)
func (_SimpleWhiteList *SimpleWhiteListSession) LevelAddr(user_ common.Address) (uint8, error) {
	return _SimpleWhiteList.Contract.LevelAddr(&_SimpleWhiteList.CallOpts, user_)
}

// LevelAddr is a free data retrieval call binding the contract method 0x5767c0e1.
//
// Solidity: function LevelAddr(address user_) view returns(uint8)
func (_SimpleWhiteList *SimpleWhiteListCallerSession) LevelAddr(user_ common.Address) (uint8, error) {
	return _SimpleWhiteList.Contract.LevelAddr(&_SimpleWhiteList.CallOpts, user_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SimpleWhiteList *SimpleWhiteListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleWhiteList.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SimpleWhiteList *SimpleWhiteListSession) Owner() (common.Address, error) {
	return _SimpleWhiteList.Contract.Owner(&_SimpleWhiteList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SimpleWhiteList *SimpleWhiteListCallerSession) Owner() (common.Address, error) {
	return _SimpleWhiteList.Contract.Owner(&_SimpleWhiteList.CallOpts)
}

// AddMultiWhiteAddr is a paid mutator transaction binding the contract method 0xbbfa057b.
//
// Solidity: function AddMultiWhiteAddr(address[] addrs) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactor) AddMultiWhiteAddr(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _SimpleWhiteList.contract.Transact(opts, "AddMultiWhiteAddr", addrs)
}

// AddMultiWhiteAddr is a paid mutator transaction binding the contract method 0xbbfa057b.
//
// Solidity: function AddMultiWhiteAddr(address[] addrs) returns()
func (_SimpleWhiteList *SimpleWhiteListSession) AddMultiWhiteAddr(addrs []common.Address) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.AddMultiWhiteAddr(&_SimpleWhiteList.TransactOpts, addrs)
}

// AddMultiWhiteAddr is a paid mutator transaction binding the contract method 0xbbfa057b.
//
// Solidity: function AddMultiWhiteAddr(address[] addrs) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactorSession) AddMultiWhiteAddr(addrs []common.Address) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.AddMultiWhiteAddr(&_SimpleWhiteList.TransactOpts, addrs)
}

// AddOperator is a paid mutator transaction binding the contract method 0x4c141abc.
//
// Solidity: function AddOperator(address op) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactor) AddOperator(opts *bind.TransactOpts, op common.Address) (*types.Transaction, error) {
	return _SimpleWhiteList.contract.Transact(opts, "AddOperator", op)
}

// AddOperator is a paid mutator transaction binding the contract method 0x4c141abc.
//
// Solidity: function AddOperator(address op) returns()
func (_SimpleWhiteList *SimpleWhiteListSession) AddOperator(op common.Address) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.AddOperator(&_SimpleWhiteList.TransactOpts, op)
}

// AddOperator is a paid mutator transaction binding the contract method 0x4c141abc.
//
// Solidity: function AddOperator(address op) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactorSession) AddOperator(op common.Address) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.AddOperator(&_SimpleWhiteList.TransactOpts, op)
}

// AddWhiteAddress is a paid mutator transaction binding the contract method 0xbf078664.
//
// Solidity: function AddWhiteAddress(address user_, uint8 level_) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactor) AddWhiteAddress(opts *bind.TransactOpts, user_ common.Address, level_ uint8) (*types.Transaction, error) {
	return _SimpleWhiteList.contract.Transact(opts, "AddWhiteAddress", user_, level_)
}

// AddWhiteAddress is a paid mutator transaction binding the contract method 0xbf078664.
//
// Solidity: function AddWhiteAddress(address user_, uint8 level_) returns()
func (_SimpleWhiteList *SimpleWhiteListSession) AddWhiteAddress(user_ common.Address, level_ uint8) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.AddWhiteAddress(&_SimpleWhiteList.TransactOpts, user_, level_)
}

// AddWhiteAddress is a paid mutator transaction binding the contract method 0xbf078664.
//
// Solidity: function AddWhiteAddress(address user_, uint8 level_) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactorSession) AddWhiteAddress(user_ common.Address, level_ uint8) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.AddWhiteAddress(&_SimpleWhiteList.TransactOpts, user_, level_)
}

// UpdateWhiteAddress is a paid mutator transaction binding the contract method 0xf3998ace.
//
// Solidity: function UpdateWhiteAddress(address user_, uint8 level_) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactor) UpdateWhiteAddress(opts *bind.TransactOpts, user_ common.Address, level_ uint8) (*types.Transaction, error) {
	return _SimpleWhiteList.contract.Transact(opts, "UpdateWhiteAddress", user_, level_)
}

// UpdateWhiteAddress is a paid mutator transaction binding the contract method 0xf3998ace.
//
// Solidity: function UpdateWhiteAddress(address user_, uint8 level_) returns()
func (_SimpleWhiteList *SimpleWhiteListSession) UpdateWhiteAddress(user_ common.Address, level_ uint8) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.UpdateWhiteAddress(&_SimpleWhiteList.TransactOpts, user_, level_)
}

// UpdateWhiteAddress is a paid mutator transaction binding the contract method 0xf3998ace.
//
// Solidity: function UpdateWhiteAddress(address user_, uint8 level_) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactorSession) UpdateWhiteAddress(user_ common.Address, level_ uint8) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.UpdateWhiteAddress(&_SimpleWhiteList.TransactOpts, user_, level_)
}

// SetDefaultLevel is a paid mutator transaction binding the contract method 0x66278950.
//
// Solidity: function setDefaultLevel(uint8 cnt_) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactor) SetDefaultLevel(opts *bind.TransactOpts, cnt_ uint8) (*types.Transaction, error) {
	return _SimpleWhiteList.contract.Transact(opts, "setDefaultLevel", cnt_)
}

// SetDefaultLevel is a paid mutator transaction binding the contract method 0x66278950.
//
// Solidity: function setDefaultLevel(uint8 cnt_) returns()
func (_SimpleWhiteList *SimpleWhiteListSession) SetDefaultLevel(cnt_ uint8) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.SetDefaultLevel(&_SimpleWhiteList.TransactOpts, cnt_)
}

// SetDefaultLevel is a paid mutator transaction binding the contract method 0x66278950.
//
// Solidity: function setDefaultLevel(uint8 cnt_) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactorSession) SetDefaultLevel(cnt_ uint8) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.SetDefaultLevel(&_SimpleWhiteList.TransactOpts, cnt_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SimpleWhiteList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SimpleWhiteList *SimpleWhiteListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.TransferOwnership(&_SimpleWhiteList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SimpleWhiteList *SimpleWhiteListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleWhiteList.Contract.TransferOwnership(&_SimpleWhiteList.TransactOpts, newOwner)
}

// SimpleWhiteListEvAddWhiteAddrIterator is returned from FilterEvAddWhiteAddr and is used to iterate over the raw logs and unpacked data for EvAddWhiteAddr events raised by the SimpleWhiteList contract.
type SimpleWhiteListEvAddWhiteAddrIterator struct {
	Event *SimpleWhiteListEvAddWhiteAddr // Event containing the contract specifics and raw log

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
func (it *SimpleWhiteListEvAddWhiteAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleWhiteListEvAddWhiteAddr)
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
		it.Event = new(SimpleWhiteListEvAddWhiteAddr)
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
func (it *SimpleWhiteListEvAddWhiteAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleWhiteListEvAddWhiteAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleWhiteListEvAddWhiteAddr represents a EvAddWhiteAddr event raised by the SimpleWhiteList contract.
type SimpleWhiteListEvAddWhiteAddr struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvAddWhiteAddr is a free log retrieval operation binding the contract event 0xf7c8eeb0fa392d7af983982e140349d250e35e75d849a18b41899ed00893b936.
//
// Solidity: event EvAddWhiteAddr(address arg0, address arg1, uint8 arg2)
func (_SimpleWhiteList *SimpleWhiteListFilterer) FilterEvAddWhiteAddr(opts *bind.FilterOpts) (*SimpleWhiteListEvAddWhiteAddrIterator, error) {

	logs, sub, err := _SimpleWhiteList.contract.FilterLogs(opts, "EvAddWhiteAddr")
	if err != nil {
		return nil, err
	}
	return &SimpleWhiteListEvAddWhiteAddrIterator{contract: _SimpleWhiteList.contract, event: "EvAddWhiteAddr", logs: logs, sub: sub}, nil
}

// WatchEvAddWhiteAddr is a free log subscription operation binding the contract event 0xf7c8eeb0fa392d7af983982e140349d250e35e75d849a18b41899ed00893b936.
//
// Solidity: event EvAddWhiteAddr(address arg0, address arg1, uint8 arg2)
func (_SimpleWhiteList *SimpleWhiteListFilterer) WatchEvAddWhiteAddr(opts *bind.WatchOpts, sink chan<- *SimpleWhiteListEvAddWhiteAddr) (event.Subscription, error) {

	logs, sub, err := _SimpleWhiteList.contract.WatchLogs(opts, "EvAddWhiteAddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleWhiteListEvAddWhiteAddr)
				if err := _SimpleWhiteList.contract.UnpackLog(event, "EvAddWhiteAddr", log); err != nil {
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

// ParseEvAddWhiteAddr is a log parse operation binding the contract event 0xf7c8eeb0fa392d7af983982e140349d250e35e75d849a18b41899ed00893b936.
//
// Solidity: event EvAddWhiteAddr(address arg0, address arg1, uint8 arg2)
func (_SimpleWhiteList *SimpleWhiteListFilterer) ParseEvAddWhiteAddr(log types.Log) (*SimpleWhiteListEvAddWhiteAddr, error) {
	event := new(SimpleWhiteListEvAddWhiteAddr)
	if err := _SimpleWhiteList.contract.UnpackLog(event, "EvAddWhiteAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleWhiteListEvUpdateWhiteAddrIterator is returned from FilterEvUpdateWhiteAddr and is used to iterate over the raw logs and unpacked data for EvUpdateWhiteAddr events raised by the SimpleWhiteList contract.
type SimpleWhiteListEvUpdateWhiteAddrIterator struct {
	Event *SimpleWhiteListEvUpdateWhiteAddr // Event containing the contract specifics and raw log

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
func (it *SimpleWhiteListEvUpdateWhiteAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleWhiteListEvUpdateWhiteAddr)
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
		it.Event = new(SimpleWhiteListEvUpdateWhiteAddr)
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
func (it *SimpleWhiteListEvUpdateWhiteAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleWhiteListEvUpdateWhiteAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleWhiteListEvUpdateWhiteAddr represents a EvUpdateWhiteAddr event raised by the SimpleWhiteList contract.
type SimpleWhiteListEvUpdateWhiteAddr struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvUpdateWhiteAddr is a free log retrieval operation binding the contract event 0x76a1864ced3d80617af5372d04dde2283c288043e0305228aeb2bee922e2ff98.
//
// Solidity: event EvUpdateWhiteAddr(address arg0, address arg1, uint8 arg2)
func (_SimpleWhiteList *SimpleWhiteListFilterer) FilterEvUpdateWhiteAddr(opts *bind.FilterOpts) (*SimpleWhiteListEvUpdateWhiteAddrIterator, error) {

	logs, sub, err := _SimpleWhiteList.contract.FilterLogs(opts, "EvUpdateWhiteAddr")
	if err != nil {
		return nil, err
	}
	return &SimpleWhiteListEvUpdateWhiteAddrIterator{contract: _SimpleWhiteList.contract, event: "EvUpdateWhiteAddr", logs: logs, sub: sub}, nil
}

// WatchEvUpdateWhiteAddr is a free log subscription operation binding the contract event 0x76a1864ced3d80617af5372d04dde2283c288043e0305228aeb2bee922e2ff98.
//
// Solidity: event EvUpdateWhiteAddr(address arg0, address arg1, uint8 arg2)
func (_SimpleWhiteList *SimpleWhiteListFilterer) WatchEvUpdateWhiteAddr(opts *bind.WatchOpts, sink chan<- *SimpleWhiteListEvUpdateWhiteAddr) (event.Subscription, error) {

	logs, sub, err := _SimpleWhiteList.contract.WatchLogs(opts, "EvUpdateWhiteAddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleWhiteListEvUpdateWhiteAddr)
				if err := _SimpleWhiteList.contract.UnpackLog(event, "EvUpdateWhiteAddr", log); err != nil {
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

// ParseEvUpdateWhiteAddr is a log parse operation binding the contract event 0x76a1864ced3d80617af5372d04dde2283c288043e0305228aeb2bee922e2ff98.
//
// Solidity: event EvUpdateWhiteAddr(address arg0, address arg1, uint8 arg2)
func (_SimpleWhiteList *SimpleWhiteListFilterer) ParseEvUpdateWhiteAddr(log types.Log) (*SimpleWhiteListEvUpdateWhiteAddr, error) {
	event := new(SimpleWhiteListEvUpdateWhiteAddr)
	if err := _SimpleWhiteList.contract.UnpackLog(event, "EvUpdateWhiteAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
