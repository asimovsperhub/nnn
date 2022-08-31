// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnscontract

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

// ColdBootMetaData contains all meta data concerning the ColdBoot contract.
var ColdBootMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"whiteList_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"CardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumColdBoot.PassCardColor\",\"name\":\"color\",\"type\":\"uint8\"}],\"name\":\"EVMintCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CardID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count_\",\"type\":\"uint256\"}],\"name\":\"MintByOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MintPassCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MintPassCardOne\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"PassCardType\",\"outputs\":[{\"internalType\":\"enumColdBoot.PassCardColor\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PassCardUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cardDefault\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx_\",\"type\":\"uint256\"}],\"name\":\"getTokenList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openToMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"passCard\",\"outputs\":[{\"internalType\":\"enumColdBoot.PassCardColor\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri_\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"count_\",\"type\":\"uint8\"}],\"name\":\"setCardDefaultNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance_\",\"type\":\"uint256\"}],\"name\":\"setCheckEthBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"open_\",\"type\":\"bool\"}],\"name\":\"setOpenToMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"white_\",\"type\":\"address\"}],\"name\":\"setWhiteListAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o_\",\"type\":\"address\"}],\"name\":\"setWhiteOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whiteListC\",\"outputs\":[{\"internalType\":\"contractIPassCardWhiteList\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whiteOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ColdBootABI is the input ABI used to generate the binding from.
// Deprecated: Use ColdBootMetaData.ABI instead.
var ColdBootABI = ColdBootMetaData.ABI

// ColdBoot is an auto generated Go binding around an Ethereum contract.
type ColdBoot struct {
	ColdBootCaller     // Read-only binding to the contract
	ColdBootTransactor // Write-only binding to the contract
	ColdBootFilterer   // Log filterer for contract events
}

// ColdBootCaller is an auto generated read-only Go binding around an Ethereum contract.
type ColdBootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ColdBootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ColdBootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ColdBootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ColdBootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ColdBootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ColdBootSession struct {
	Contract     *ColdBoot         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ColdBootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ColdBootCallerSession struct {
	Contract *ColdBootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ColdBootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ColdBootTransactorSession struct {
	Contract     *ColdBootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ColdBootRaw is an auto generated low-level Go binding around an Ethereum contract.
type ColdBootRaw struct {
	Contract *ColdBoot // Generic contract binding to access the raw methods on
}

// ColdBootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ColdBootCallerRaw struct {
	Contract *ColdBootCaller // Generic read-only contract binding to access the raw methods on
}

// ColdBootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ColdBootTransactorRaw struct {
	Contract *ColdBootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewColdBoot creates a new instance of ColdBoot, bound to a specific deployed contract.
func NewColdBoot(address common.Address, backend bind.ContractBackend) (*ColdBoot, error) {
	contract, err := bindColdBoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ColdBoot{ColdBootCaller: ColdBootCaller{contract: contract}, ColdBootTransactor: ColdBootTransactor{contract: contract}, ColdBootFilterer: ColdBootFilterer{contract: contract}}, nil
}

// NewColdBootCaller creates a new read-only instance of ColdBoot, bound to a specific deployed contract.
func NewColdBootCaller(address common.Address, caller bind.ContractCaller) (*ColdBootCaller, error) {
	contract, err := bindColdBoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ColdBootCaller{contract: contract}, nil
}

// NewColdBootTransactor creates a new write-only instance of ColdBoot, bound to a specific deployed contract.
func NewColdBootTransactor(address common.Address, transactor bind.ContractTransactor) (*ColdBootTransactor, error) {
	contract, err := bindColdBoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ColdBootTransactor{contract: contract}, nil
}

// NewColdBootFilterer creates a new log filterer instance of ColdBoot, bound to a specific deployed contract.
func NewColdBootFilterer(address common.Address, filterer bind.ContractFilterer) (*ColdBootFilterer, error) {
	contract, err := bindColdBoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ColdBootFilterer{contract: contract}, nil
}

// bindColdBoot binds a generic wrapper to an already deployed contract.
func bindColdBoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ColdBootABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ColdBoot *ColdBootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ColdBoot.Contract.ColdBootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ColdBoot *ColdBootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ColdBoot.Contract.ColdBootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ColdBoot *ColdBootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ColdBoot.Contract.ColdBootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ColdBoot *ColdBootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ColdBoot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ColdBoot *ColdBootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ColdBoot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ColdBoot *ColdBootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ColdBoot.Contract.contract.Transact(opts, method, params...)
}

// CardID is a free data retrieval call binding the contract method 0x3c2666d4.
//
// Solidity: function CardID() view returns(uint256)
func (_ColdBoot *ColdBootCaller) CardID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "CardID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CardID is a free data retrieval call binding the contract method 0x3c2666d4.
//
// Solidity: function CardID() view returns(uint256)
func (_ColdBoot *ColdBootSession) CardID() (*big.Int, error) {
	return _ColdBoot.Contract.CardID(&_ColdBoot.CallOpts)
}

// CardID is a free data retrieval call binding the contract method 0x3c2666d4.
//
// Solidity: function CardID() view returns(uint256)
func (_ColdBoot *ColdBootCallerSession) CardID() (*big.Int, error) {
	return _ColdBoot.Contract.CardID(&_ColdBoot.CallOpts)
}

// PassCardType is a free data retrieval call binding the contract method 0xcb45547f.
//
// Solidity: function PassCardType(uint256 tokenId_) view returns(uint8)
func (_ColdBoot *ColdBootCaller) PassCardType(opts *bind.CallOpts, tokenId_ *big.Int) (uint8, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "PassCardType", tokenId_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PassCardType is a free data retrieval call binding the contract method 0xcb45547f.
//
// Solidity: function PassCardType(uint256 tokenId_) view returns(uint8)
func (_ColdBoot *ColdBootSession) PassCardType(tokenId_ *big.Int) (uint8, error) {
	return _ColdBoot.Contract.PassCardType(&_ColdBoot.CallOpts, tokenId_)
}

// PassCardType is a free data retrieval call binding the contract method 0xcb45547f.
//
// Solidity: function PassCardType(uint256 tokenId_) view returns(uint8)
func (_ColdBoot *ColdBootCallerSession) PassCardType(tokenId_ *big.Int) (uint8, error) {
	return _ColdBoot.Contract.PassCardType(&_ColdBoot.CallOpts, tokenId_)
}

// PassCardUsed is a free data retrieval call binding the contract method 0x121a0a28.
//
// Solidity: function PassCardUsed() view returns(uint256, uint256, uint256)
func (_ColdBoot *ColdBootCaller) PassCardUsed(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "PassCardUsed")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// PassCardUsed is a free data retrieval call binding the contract method 0x121a0a28.
//
// Solidity: function PassCardUsed() view returns(uint256, uint256, uint256)
func (_ColdBoot *ColdBootSession) PassCardUsed() (*big.Int, *big.Int, *big.Int, error) {
	return _ColdBoot.Contract.PassCardUsed(&_ColdBoot.CallOpts)
}

// PassCardUsed is a free data retrieval call binding the contract method 0x121a0a28.
//
// Solidity: function PassCardUsed() view returns(uint256, uint256, uint256)
func (_ColdBoot *ColdBootCallerSession) PassCardUsed() (*big.Int, *big.Int, *big.Int, error) {
	return _ColdBoot.Contract.PassCardUsed(&_ColdBoot.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ColdBoot *ColdBootCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ColdBoot *ColdBootSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ColdBoot.Contract.BalanceOf(&_ColdBoot.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ColdBoot *ColdBootCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ColdBoot.Contract.BalanceOf(&_ColdBoot.CallOpts, owner)
}

// CardDefault is a free data retrieval call binding the contract method 0xebeb697a.
//
// Solidity: function cardDefault() view returns(uint8)
func (_ColdBoot *ColdBootCaller) CardDefault(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "cardDefault")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CardDefault is a free data retrieval call binding the contract method 0xebeb697a.
//
// Solidity: function cardDefault() view returns(uint8)
func (_ColdBoot *ColdBootSession) CardDefault() (uint8, error) {
	return _ColdBoot.Contract.CardDefault(&_ColdBoot.CallOpts)
}

// CardDefault is a free data retrieval call binding the contract method 0xebeb697a.
//
// Solidity: function cardDefault() view returns(uint8)
func (_ColdBoot *ColdBootCallerSession) CardDefault() (uint8, error) {
	return _ColdBoot.Contract.CardDefault(&_ColdBoot.CallOpts)
}

// EthBalance is a free data retrieval call binding the contract method 0x4e6630b0.
//
// Solidity: function ethBalance() view returns(uint256)
func (_ColdBoot *ColdBootCaller) EthBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "ethBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthBalance is a free data retrieval call binding the contract method 0x4e6630b0.
//
// Solidity: function ethBalance() view returns(uint256)
func (_ColdBoot *ColdBootSession) EthBalance() (*big.Int, error) {
	return _ColdBoot.Contract.EthBalance(&_ColdBoot.CallOpts)
}

// EthBalance is a free data retrieval call binding the contract method 0x4e6630b0.
//
// Solidity: function ethBalance() view returns(uint256)
func (_ColdBoot *ColdBootCallerSession) EthBalance() (*big.Int, error) {
	return _ColdBoot.Contract.EthBalance(&_ColdBoot.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ColdBoot *ColdBootCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ColdBoot *ColdBootSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ColdBoot.Contract.GetApproved(&_ColdBoot.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ColdBoot *ColdBootCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ColdBoot.Contract.GetApproved(&_ColdBoot.CallOpts, tokenId)
}

// GetTokenList is a free data retrieval call binding the contract method 0x41b63bd8.
//
// Solidity: function getTokenList(uint256 idx_) view returns(bool, uint256)
func (_ColdBoot *ColdBootCaller) GetTokenList(opts *bind.CallOpts, idx_ *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "getTokenList", idx_)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTokenList is a free data retrieval call binding the contract method 0x41b63bd8.
//
// Solidity: function getTokenList(uint256 idx_) view returns(bool, uint256)
func (_ColdBoot *ColdBootSession) GetTokenList(idx_ *big.Int) (bool, *big.Int, error) {
	return _ColdBoot.Contract.GetTokenList(&_ColdBoot.CallOpts, idx_)
}

// GetTokenList is a free data retrieval call binding the contract method 0x41b63bd8.
//
// Solidity: function getTokenList(uint256 idx_) view returns(bool, uint256)
func (_ColdBoot *ColdBootCallerSession) GetTokenList(idx_ *big.Int) (bool, *big.Int, error) {
	return _ColdBoot.Contract.GetTokenList(&_ColdBoot.CallOpts, idx_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ColdBoot *ColdBootCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ColdBoot *ColdBootSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ColdBoot.Contract.IsApprovedForAll(&_ColdBoot.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ColdBoot *ColdBootCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ColdBoot.Contract.IsApprovedForAll(&_ColdBoot.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ColdBoot *ColdBootCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ColdBoot *ColdBootSession) Name() (string, error) {
	return _ColdBoot.Contract.Name(&_ColdBoot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ColdBoot *ColdBootCallerSession) Name() (string, error) {
	return _ColdBoot.Contract.Name(&_ColdBoot.CallOpts)
}

// OpenToMint is a free data retrieval call binding the contract method 0xb266a686.
//
// Solidity: function openToMint() view returns(bool)
func (_ColdBoot *ColdBootCaller) OpenToMint(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "openToMint")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OpenToMint is a free data retrieval call binding the contract method 0xb266a686.
//
// Solidity: function openToMint() view returns(bool)
func (_ColdBoot *ColdBootSession) OpenToMint() (bool, error) {
	return _ColdBoot.Contract.OpenToMint(&_ColdBoot.CallOpts)
}

// OpenToMint is a free data retrieval call binding the contract method 0xb266a686.
//
// Solidity: function openToMint() view returns(bool)
func (_ColdBoot *ColdBootCallerSession) OpenToMint() (bool, error) {
	return _ColdBoot.Contract.OpenToMint(&_ColdBoot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ColdBoot *ColdBootCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ColdBoot *ColdBootSession) Owner() (common.Address, error) {
	return _ColdBoot.Contract.Owner(&_ColdBoot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ColdBoot *ColdBootCallerSession) Owner() (common.Address, error) {
	return _ColdBoot.Contract.Owner(&_ColdBoot.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ColdBoot *ColdBootCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ColdBoot *ColdBootSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ColdBoot.Contract.OwnerOf(&_ColdBoot.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ColdBoot *ColdBootCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ColdBoot.Contract.OwnerOf(&_ColdBoot.CallOpts, tokenId)
}

// PassCard is a free data retrieval call binding the contract method 0x7a04d3b4.
//
// Solidity: function passCard(uint256 ) view returns(uint8)
func (_ColdBoot *ColdBootCaller) PassCard(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "passCard", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PassCard is a free data retrieval call binding the contract method 0x7a04d3b4.
//
// Solidity: function passCard(uint256 ) view returns(uint8)
func (_ColdBoot *ColdBootSession) PassCard(arg0 *big.Int) (uint8, error) {
	return _ColdBoot.Contract.PassCard(&_ColdBoot.CallOpts, arg0)
}

// PassCard is a free data retrieval call binding the contract method 0x7a04d3b4.
//
// Solidity: function passCard(uint256 ) view returns(uint8)
func (_ColdBoot *ColdBootCallerSession) PassCard(arg0 *big.Int) (uint8, error) {
	return _ColdBoot.Contract.PassCard(&_ColdBoot.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ColdBoot *ColdBootCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ColdBoot *ColdBootSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ColdBoot.Contract.SupportsInterface(&_ColdBoot.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ColdBoot *ColdBootCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ColdBoot.Contract.SupportsInterface(&_ColdBoot.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ColdBoot *ColdBootCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ColdBoot *ColdBootSession) Symbol() (string, error) {
	return _ColdBoot.Contract.Symbol(&_ColdBoot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ColdBoot *ColdBootCallerSession) Symbol() (string, error) {
	return _ColdBoot.Contract.Symbol(&_ColdBoot.CallOpts)
}

// TokenList is a free data retrieval call binding the contract method 0xbebfd690.
//
// Solidity: function tokenList(address , uint256 ) view returns(uint256)
func (_ColdBoot *ColdBootCaller) TokenList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "tokenList", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenList is a free data retrieval call binding the contract method 0xbebfd690.
//
// Solidity: function tokenList(address , uint256 ) view returns(uint256)
func (_ColdBoot *ColdBootSession) TokenList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ColdBoot.Contract.TokenList(&_ColdBoot.CallOpts, arg0, arg1)
}

// TokenList is a free data retrieval call binding the contract method 0xbebfd690.
//
// Solidity: function tokenList(address , uint256 ) view returns(uint256)
func (_ColdBoot *ColdBootCallerSession) TokenList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ColdBoot.Contract.TokenList(&_ColdBoot.CallOpts, arg0, arg1)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ColdBoot *ColdBootCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ColdBoot *ColdBootSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ColdBoot.Contract.TokenURI(&_ColdBoot.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ColdBoot *ColdBootCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ColdBoot.Contract.TokenURI(&_ColdBoot.CallOpts, tokenId)
}

// WhiteListC is a free data retrieval call binding the contract method 0xde0f8282.
//
// Solidity: function whiteListC() view returns(address)
func (_ColdBoot *ColdBootCaller) WhiteListC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "whiteListC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhiteListC is a free data retrieval call binding the contract method 0xde0f8282.
//
// Solidity: function whiteListC() view returns(address)
func (_ColdBoot *ColdBootSession) WhiteListC() (common.Address, error) {
	return _ColdBoot.Contract.WhiteListC(&_ColdBoot.CallOpts)
}

// WhiteListC is a free data retrieval call binding the contract method 0xde0f8282.
//
// Solidity: function whiteListC() view returns(address)
func (_ColdBoot *ColdBootCallerSession) WhiteListC() (common.Address, error) {
	return _ColdBoot.Contract.WhiteListC(&_ColdBoot.CallOpts)
}

// WhiteOperator is a free data retrieval call binding the contract method 0x07c5ae77.
//
// Solidity: function whiteOperator() view returns(address)
func (_ColdBoot *ColdBootCaller) WhiteOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ColdBoot.contract.Call(opts, &out, "whiteOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhiteOperator is a free data retrieval call binding the contract method 0x07c5ae77.
//
// Solidity: function whiteOperator() view returns(address)
func (_ColdBoot *ColdBootSession) WhiteOperator() (common.Address, error) {
	return _ColdBoot.Contract.WhiteOperator(&_ColdBoot.CallOpts)
}

// WhiteOperator is a free data retrieval call binding the contract method 0x07c5ae77.
//
// Solidity: function whiteOperator() view returns(address)
func (_ColdBoot *ColdBootCallerSession) WhiteOperator() (common.Address, error) {
	return _ColdBoot.Contract.WhiteOperator(&_ColdBoot.CallOpts)
}

// MintByOperator is a paid mutator transaction binding the contract method 0x81e1378a.
//
// Solidity: function MintByOperator(address user_, uint256 count_) returns()
func (_ColdBoot *ColdBootTransactor) MintByOperator(opts *bind.TransactOpts, user_ common.Address, count_ *big.Int) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "MintByOperator", user_, count_)
}

// MintByOperator is a paid mutator transaction binding the contract method 0x81e1378a.
//
// Solidity: function MintByOperator(address user_, uint256 count_) returns()
func (_ColdBoot *ColdBootSession) MintByOperator(user_ common.Address, count_ *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.MintByOperator(&_ColdBoot.TransactOpts, user_, count_)
}

// MintByOperator is a paid mutator transaction binding the contract method 0x81e1378a.
//
// Solidity: function MintByOperator(address user_, uint256 count_) returns()
func (_ColdBoot *ColdBootTransactorSession) MintByOperator(user_ common.Address, count_ *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.MintByOperator(&_ColdBoot.TransactOpts, user_, count_)
}

// MintPassCard is a paid mutator transaction binding the contract method 0xc346f624.
//
// Solidity: function MintPassCard() returns()
func (_ColdBoot *ColdBootTransactor) MintPassCard(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "MintPassCard")
}

// MintPassCard is a paid mutator transaction binding the contract method 0xc346f624.
//
// Solidity: function MintPassCard() returns()
func (_ColdBoot *ColdBootSession) MintPassCard() (*types.Transaction, error) {
	return _ColdBoot.Contract.MintPassCard(&_ColdBoot.TransactOpts)
}

// MintPassCard is a paid mutator transaction binding the contract method 0xc346f624.
//
// Solidity: function MintPassCard() returns()
func (_ColdBoot *ColdBootTransactorSession) MintPassCard() (*types.Transaction, error) {
	return _ColdBoot.Contract.MintPassCard(&_ColdBoot.TransactOpts)
}

// MintPassCardOne is a paid mutator transaction binding the contract method 0x17aef551.
//
// Solidity: function MintPassCardOne() returns()
func (_ColdBoot *ColdBootTransactor) MintPassCardOne(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "MintPassCardOne")
}

// MintPassCardOne is a paid mutator transaction binding the contract method 0x17aef551.
//
// Solidity: function MintPassCardOne() returns()
func (_ColdBoot *ColdBootSession) MintPassCardOne() (*types.Transaction, error) {
	return _ColdBoot.Contract.MintPassCardOne(&_ColdBoot.TransactOpts)
}

// MintPassCardOne is a paid mutator transaction binding the contract method 0x17aef551.
//
// Solidity: function MintPassCardOne() returns()
func (_ColdBoot *ColdBootTransactorSession) MintPassCardOne() (*types.Transaction, error) {
	return _ColdBoot.Contract.MintPassCardOne(&_ColdBoot.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ColdBoot *ColdBootTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ColdBoot *ColdBootSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.Approve(&_ColdBoot.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ColdBoot *ColdBootTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.Approve(&_ColdBoot.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ColdBoot *ColdBootTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ColdBoot *ColdBootSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.SafeTransferFrom(&_ColdBoot.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ColdBoot *ColdBootTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.SafeTransferFrom(&_ColdBoot.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ColdBoot *ColdBootTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ColdBoot *ColdBootSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ColdBoot.Contract.SafeTransferFrom0(&_ColdBoot.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ColdBoot *ColdBootTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ColdBoot.Contract.SafeTransferFrom0(&_ColdBoot.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ColdBoot *ColdBootTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ColdBoot *ColdBootSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetApprovalForAll(&_ColdBoot.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ColdBoot *ColdBootTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetApprovalForAll(&_ColdBoot.TransactOpts, operator, approved)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_ColdBoot *ColdBootTransactor) SetBaseUri(opts *bind.TransactOpts, baseUri_ string) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "setBaseUri", baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_ColdBoot *ColdBootSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetBaseUri(&_ColdBoot.TransactOpts, baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_ColdBoot *ColdBootTransactorSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetBaseUri(&_ColdBoot.TransactOpts, baseUri_)
}

// SetCardDefaultNum is a paid mutator transaction binding the contract method 0xcf7dbcfb.
//
// Solidity: function setCardDefaultNum(uint8 count_) returns()
func (_ColdBoot *ColdBootTransactor) SetCardDefaultNum(opts *bind.TransactOpts, count_ uint8) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "setCardDefaultNum", count_)
}

// SetCardDefaultNum is a paid mutator transaction binding the contract method 0xcf7dbcfb.
//
// Solidity: function setCardDefaultNum(uint8 count_) returns()
func (_ColdBoot *ColdBootSession) SetCardDefaultNum(count_ uint8) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetCardDefaultNum(&_ColdBoot.TransactOpts, count_)
}

// SetCardDefaultNum is a paid mutator transaction binding the contract method 0xcf7dbcfb.
//
// Solidity: function setCardDefaultNum(uint8 count_) returns()
func (_ColdBoot *ColdBootTransactorSession) SetCardDefaultNum(count_ uint8) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetCardDefaultNum(&_ColdBoot.TransactOpts, count_)
}

// SetCheckEthBalance is a paid mutator transaction binding the contract method 0x254e73d0.
//
// Solidity: function setCheckEthBalance(uint256 balance_) returns()
func (_ColdBoot *ColdBootTransactor) SetCheckEthBalance(opts *bind.TransactOpts, balance_ *big.Int) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "setCheckEthBalance", balance_)
}

// SetCheckEthBalance is a paid mutator transaction binding the contract method 0x254e73d0.
//
// Solidity: function setCheckEthBalance(uint256 balance_) returns()
func (_ColdBoot *ColdBootSession) SetCheckEthBalance(balance_ *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetCheckEthBalance(&_ColdBoot.TransactOpts, balance_)
}

// SetCheckEthBalance is a paid mutator transaction binding the contract method 0x254e73d0.
//
// Solidity: function setCheckEthBalance(uint256 balance_) returns()
func (_ColdBoot *ColdBootTransactorSession) SetCheckEthBalance(balance_ *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetCheckEthBalance(&_ColdBoot.TransactOpts, balance_)
}

// SetOpenToMint is a paid mutator transaction binding the contract method 0x20e79b71.
//
// Solidity: function setOpenToMint(bool open_) returns()
func (_ColdBoot *ColdBootTransactor) SetOpenToMint(opts *bind.TransactOpts, open_ bool) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "setOpenToMint", open_)
}

// SetOpenToMint is a paid mutator transaction binding the contract method 0x20e79b71.
//
// Solidity: function setOpenToMint(bool open_) returns()
func (_ColdBoot *ColdBootSession) SetOpenToMint(open_ bool) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetOpenToMint(&_ColdBoot.TransactOpts, open_)
}

// SetOpenToMint is a paid mutator transaction binding the contract method 0x20e79b71.
//
// Solidity: function setOpenToMint(bool open_) returns()
func (_ColdBoot *ColdBootTransactorSession) SetOpenToMint(open_ bool) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetOpenToMint(&_ColdBoot.TransactOpts, open_)
}

// SetWhiteListAddr is a paid mutator transaction binding the contract method 0x4f65a164.
//
// Solidity: function setWhiteListAddr(address white_) returns()
func (_ColdBoot *ColdBootTransactor) SetWhiteListAddr(opts *bind.TransactOpts, white_ common.Address) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "setWhiteListAddr", white_)
}

// SetWhiteListAddr is a paid mutator transaction binding the contract method 0x4f65a164.
//
// Solidity: function setWhiteListAddr(address white_) returns()
func (_ColdBoot *ColdBootSession) SetWhiteListAddr(white_ common.Address) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetWhiteListAddr(&_ColdBoot.TransactOpts, white_)
}

// SetWhiteListAddr is a paid mutator transaction binding the contract method 0x4f65a164.
//
// Solidity: function setWhiteListAddr(address white_) returns()
func (_ColdBoot *ColdBootTransactorSession) SetWhiteListAddr(white_ common.Address) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetWhiteListAddr(&_ColdBoot.TransactOpts, white_)
}

// SetWhiteOperator is a paid mutator transaction binding the contract method 0x13bf02d8.
//
// Solidity: function setWhiteOperator(address o_) returns()
func (_ColdBoot *ColdBootTransactor) SetWhiteOperator(opts *bind.TransactOpts, o_ common.Address) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "setWhiteOperator", o_)
}

// SetWhiteOperator is a paid mutator transaction binding the contract method 0x13bf02d8.
//
// Solidity: function setWhiteOperator(address o_) returns()
func (_ColdBoot *ColdBootSession) SetWhiteOperator(o_ common.Address) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetWhiteOperator(&_ColdBoot.TransactOpts, o_)
}

// SetWhiteOperator is a paid mutator transaction binding the contract method 0x13bf02d8.
//
// Solidity: function setWhiteOperator(address o_) returns()
func (_ColdBoot *ColdBootTransactorSession) SetWhiteOperator(o_ common.Address) (*types.Transaction, error) {
	return _ColdBoot.Contract.SetWhiteOperator(&_ColdBoot.TransactOpts, o_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ColdBoot *ColdBootTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ColdBoot *ColdBootSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.TransferFrom(&_ColdBoot.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ColdBoot *ColdBootTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ColdBoot.Contract.TransferFrom(&_ColdBoot.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ColdBoot *ColdBootTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ColdBoot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ColdBoot *ColdBootSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ColdBoot.Contract.TransferOwnership(&_ColdBoot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ColdBoot *ColdBootTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ColdBoot.Contract.TransferOwnership(&_ColdBoot.TransactOpts, newOwner)
}

// ColdBootApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ColdBoot contract.
type ColdBootApprovalIterator struct {
	Event *ColdBootApproval // Event containing the contract specifics and raw log

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
func (it *ColdBootApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ColdBootApproval)
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
		it.Event = new(ColdBootApproval)
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
func (it *ColdBootApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ColdBootApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ColdBootApproval represents a Approval event raised by the ColdBoot contract.
type ColdBootApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ColdBoot *ColdBootFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ColdBootApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ColdBoot.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ColdBootApprovalIterator{contract: _ColdBoot.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ColdBoot *ColdBootFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ColdBootApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ColdBoot.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ColdBootApproval)
				if err := _ColdBoot.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ColdBoot *ColdBootFilterer) ParseApproval(log types.Log) (*ColdBootApproval, error) {
	event := new(ColdBootApproval)
	if err := _ColdBoot.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ColdBootApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ColdBoot contract.
type ColdBootApprovalForAllIterator struct {
	Event *ColdBootApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ColdBootApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ColdBootApprovalForAll)
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
		it.Event = new(ColdBootApprovalForAll)
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
func (it *ColdBootApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ColdBootApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ColdBootApprovalForAll represents a ApprovalForAll event raised by the ColdBoot contract.
type ColdBootApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ColdBoot *ColdBootFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ColdBootApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ColdBoot.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ColdBootApprovalForAllIterator{contract: _ColdBoot.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ColdBoot *ColdBootFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ColdBootApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ColdBoot.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ColdBootApprovalForAll)
				if err := _ColdBoot.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ColdBoot *ColdBootFilterer) ParseApprovalForAll(log types.Log) (*ColdBootApprovalForAll, error) {
	event := new(ColdBootApprovalForAll)
	if err := _ColdBoot.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ColdBootEVMintCardIterator is returned from FilterEVMintCard and is used to iterate over the raw logs and unpacked data for EVMintCard events raised by the ColdBoot contract.
type ColdBootEVMintCardIterator struct {
	Event *ColdBootEVMintCard // Event containing the contract specifics and raw log

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
func (it *ColdBootEVMintCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ColdBootEVMintCard)
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
		it.Event = new(ColdBootEVMintCard)
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
func (it *ColdBootEVMintCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ColdBootEVMintCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ColdBootEVMintCard represents a EVMintCard event raised by the ColdBoot contract.
type ColdBootEVMintCard struct {
	User   common.Address
	CardId *big.Int
	Color  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEVMintCard is a free log retrieval operation binding the contract event 0x6bfc38f760b8c165ed9fe7935e605974e1f7c793e621cb1e76ee21a8ff9bae4f.
//
// Solidity: event EVMintCard(address user, uint256 CardId, uint8 color)
func (_ColdBoot *ColdBootFilterer) FilterEVMintCard(opts *bind.FilterOpts) (*ColdBootEVMintCardIterator, error) {

	logs, sub, err := _ColdBoot.contract.FilterLogs(opts, "EVMintCard")
	if err != nil {
		return nil, err
	}
	return &ColdBootEVMintCardIterator{contract: _ColdBoot.contract, event: "EVMintCard", logs: logs, sub: sub}, nil
}

// WatchEVMintCard is a free log subscription operation binding the contract event 0x6bfc38f760b8c165ed9fe7935e605974e1f7c793e621cb1e76ee21a8ff9bae4f.
//
// Solidity: event EVMintCard(address user, uint256 CardId, uint8 color)
func (_ColdBoot *ColdBootFilterer) WatchEVMintCard(opts *bind.WatchOpts, sink chan<- *ColdBootEVMintCard) (event.Subscription, error) {

	logs, sub, err := _ColdBoot.contract.WatchLogs(opts, "EVMintCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ColdBootEVMintCard)
				if err := _ColdBoot.contract.UnpackLog(event, "EVMintCard", log); err != nil {
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

// ParseEVMintCard is a log parse operation binding the contract event 0x6bfc38f760b8c165ed9fe7935e605974e1f7c793e621cb1e76ee21a8ff9bae4f.
//
// Solidity: event EVMintCard(address user, uint256 CardId, uint8 color)
func (_ColdBoot *ColdBootFilterer) ParseEVMintCard(log types.Log) (*ColdBootEVMintCard, error) {
	event := new(ColdBootEVMintCard)
	if err := _ColdBoot.contract.UnpackLog(event, "EVMintCard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ColdBootTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ColdBoot contract.
type ColdBootTransferIterator struct {
	Event *ColdBootTransfer // Event containing the contract specifics and raw log

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
func (it *ColdBootTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ColdBootTransfer)
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
		it.Event = new(ColdBootTransfer)
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
func (it *ColdBootTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ColdBootTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ColdBootTransfer represents a Transfer event raised by the ColdBoot contract.
type ColdBootTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ColdBoot *ColdBootFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ColdBootTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ColdBoot.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ColdBootTransferIterator{contract: _ColdBoot.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ColdBoot *ColdBootFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ColdBootTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ColdBoot.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ColdBootTransfer)
				if err := _ColdBoot.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ColdBoot *ColdBootFilterer) ParseTransfer(log types.Log) (*ColdBootTransfer, error) {
	event := new(ColdBootTransfer)
	if err := _ColdBoot.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
