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

// TokenPairMetaData contains all meta data concerning the TokenPair contract.
var TokenPairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TokenPairABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenPairMetaData.ABI instead.
var TokenPairABI = TokenPairMetaData.ABI

// TokenPair is an auto generated Go binding around an Ethereum contract.
type TokenPair struct {
	TokenPairCaller     // Read-only binding to the contract
	TokenPairTransactor // Write-only binding to the contract
	TokenPairFilterer   // Log filterer for contract events
}

// TokenPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenPairSession struct {
	Contract     *TokenPair        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenPairCallerSession struct {
	Contract *TokenPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenPairTransactorSession struct {
	Contract     *TokenPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenPairRaw struct {
	Contract *TokenPair // Generic contract binding to access the raw methods on
}

// TokenPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenPairCallerRaw struct {
	Contract *TokenPairCaller // Generic read-only contract binding to access the raw methods on
}

// TokenPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenPairTransactorRaw struct {
	Contract *TokenPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenPair creates a new instance of TokenPair, bound to a specific deployed contract.
func NewTokenPair(address common.Address, backend bind.ContractBackend) (*TokenPair, error) {
	contract, err := bindTokenPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPair{TokenPairCaller: TokenPairCaller{contract: contract}, TokenPairTransactor: TokenPairTransactor{contract: contract}, TokenPairFilterer: TokenPairFilterer{contract: contract}}, nil
}

// NewTokenPairCaller creates a new read-only instance of TokenPair, bound to a specific deployed contract.
func NewTokenPairCaller(address common.Address, caller bind.ContractCaller) (*TokenPairCaller, error) {
	contract, err := bindTokenPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPairCaller{contract: contract}, nil
}

// NewTokenPairTransactor creates a new write-only instance of TokenPair, bound to a specific deployed contract.
func NewTokenPairTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPairTransactor, error) {
	contract, err := bindTokenPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPairTransactor{contract: contract}, nil
}

// NewTokenPairFilterer creates a new log filterer instance of TokenPair, bound to a specific deployed contract.
func NewTokenPairFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPairFilterer, error) {
	contract, err := bindTokenPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPairFilterer{contract: contract}, nil
}

// bindTokenPair binds a generic wrapper to an already deployed contract.
func bindTokenPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenPairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPair *TokenPairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPair.Contract.TokenPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPair *TokenPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPair.Contract.TokenPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPair *TokenPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPair.Contract.TokenPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPair *TokenPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPair *TokenPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPair *TokenPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPair.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TokenPair *TokenPairCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TokenPair *TokenPairSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TokenPair.Contract.BalanceOf(&_TokenPair.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TokenPair *TokenPairCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TokenPair.Contract.BalanceOf(&_TokenPair.CallOpts, owner)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_TokenPair *TokenPairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_TokenPair *TokenPairSession) Price0CumulativeLast() (*big.Int, error) {
	return _TokenPair.Contract.Price0CumulativeLast(&_TokenPair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_TokenPair *TokenPairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _TokenPair.Contract.Price0CumulativeLast(&_TokenPair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_TokenPair *TokenPairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_TokenPair *TokenPairSession) Price1CumulativeLast() (*big.Int, error) {
	return _TokenPair.Contract.Price1CumulativeLast(&_TokenPair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_TokenPair *TokenPairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _TokenPair.Contract.Price1CumulativeLast(&_TokenPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_TokenPair *TokenPairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_TokenPair *TokenPairSession) Symbol() (string, error) {
	return _TokenPair.Contract.Symbol(&_TokenPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_TokenPair *TokenPairCallerSession) Symbol() (string, error) {
	return _TokenPair.Contract.Symbol(&_TokenPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_TokenPair *TokenPairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_TokenPair *TokenPairSession) Token0() (common.Address, error) {
	return _TokenPair.Contract.Token0(&_TokenPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_TokenPair *TokenPairCallerSession) Token0() (common.Address, error) {
	return _TokenPair.Contract.Token0(&_TokenPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_TokenPair *TokenPairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_TokenPair *TokenPairSession) Token1() (common.Address, error) {
	return _TokenPair.Contract.Token1(&_TokenPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_TokenPair *TokenPairCallerSession) Token1() (common.Address, error) {
	return _TokenPair.Contract.Token1(&_TokenPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenPair *TokenPairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenPair *TokenPairSession) TotalSupply() (*big.Int, error) {
	return _TokenPair.Contract.TotalSupply(&_TokenPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenPair *TokenPairCallerSession) TotalSupply() (*big.Int, error) {
	return _TokenPair.Contract.TotalSupply(&_TokenPair.CallOpts)
}
