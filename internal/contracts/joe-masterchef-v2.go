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

// MasterChefMetaData contains all meta data concerning the MasterChef contract.
var MasterChefMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractJoeToken\",\"name\":\"_joe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_investorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_joePerSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_devPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_treasuryPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_investorPercent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetDevAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_joePerSec\",\"type\":\"uint256\"}],\"name\":\"UpdateEmissionRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accJoePerShare\",\"type\":\"uint256\"}],\"name\":\"UpdatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devAddr\",\"type\":\"address\"}],\"name\":\"dev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"investorAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"investorPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joe\",\"outputs\":[{\"internalType\":\"contractJoeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joePerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingJoe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accJoePerShare\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"rewarderBonusTokenInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDevPercent\",\"type\":\"uint256\"}],\"name\":\"setDevPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_investorAddr\",\"type\":\"address\"}],\"name\":\"setInvestorAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newInvestorPercent\",\"type\":\"uint256\"}],\"name\":\"setInvestorPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"}],\"name\":\"setTreasuryAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"setTreasuryPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_joePerSec\",\"type\":\"uint256\"}],\"name\":\"updateEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MasterChefABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterChefMetaData.ABI instead.
var MasterChefABI = MasterChefMetaData.ABI

// MasterChef is an auto generated Go binding around an Ethereum contract.
type MasterChef struct {
	MasterChefCaller     // Read-only binding to the contract
	MasterChefTransactor // Write-only binding to the contract
	MasterChefFilterer   // Log filterer for contract events
}

// MasterChefCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterChefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterChefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterChefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterChefSession struct {
	Contract     *MasterChef       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterChefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterChefCallerSession struct {
	Contract *MasterChefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MasterChefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterChefTransactorSession struct {
	Contract     *MasterChefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MasterChefRaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterChefRaw struct {
	Contract *MasterChef // Generic contract binding to access the raw methods on
}

// MasterChefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterChefCallerRaw struct {
	Contract *MasterChefCaller // Generic read-only contract binding to access the raw methods on
}

// MasterChefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterChefTransactorRaw struct {
	Contract *MasterChefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterChef creates a new instance of MasterChef, bound to a specific deployed contract.
func NewMasterChef(address common.Address, backend bind.ContractBackend) (*MasterChef, error) {
	contract, err := bindMasterChef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterChef{MasterChefCaller: MasterChefCaller{contract: contract}, MasterChefTransactor: MasterChefTransactor{contract: contract}, MasterChefFilterer: MasterChefFilterer{contract: contract}}, nil
}

// NewMasterChefCaller creates a new read-only instance of MasterChef, bound to a specific deployed contract.
func NewMasterChefCaller(address common.Address, caller bind.ContractCaller) (*MasterChefCaller, error) {
	contract, err := bindMasterChef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterChefCaller{contract: contract}, nil
}

// NewMasterChefTransactor creates a new write-only instance of MasterChef, bound to a specific deployed contract.
func NewMasterChefTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterChefTransactor, error) {
	contract, err := bindMasterChef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterChefTransactor{contract: contract}, nil
}

// NewMasterChefFilterer creates a new log filterer instance of MasterChef, bound to a specific deployed contract.
func NewMasterChefFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterChefFilterer, error) {
	contract, err := bindMasterChef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterChefFilterer{contract: contract}, nil
}

// bindMasterChef binds a generic wrapper to an already deployed contract.
func bindMasterChef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasterChefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterChef *MasterChefRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterChef.Contract.MasterChefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterChef *MasterChefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChef.Contract.MasterChefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterChef *MasterChefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterChef.Contract.MasterChefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterChef *MasterChefCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterChef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterChef *MasterChefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterChef *MasterChefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterChef.Contract.contract.Transact(opts, method, params...)
}

// DevAddr is a free data retrieval call binding the contract method 0xda09c72c.
//
// Solidity: function devAddr() view returns(address)
func (_MasterChef *MasterChefCaller) DevAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "devAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DevAddr is a free data retrieval call binding the contract method 0xda09c72c.
//
// Solidity: function devAddr() view returns(address)
func (_MasterChef *MasterChefSession) DevAddr() (common.Address, error) {
	return _MasterChef.Contract.DevAddr(&_MasterChef.CallOpts)
}

// DevAddr is a free data retrieval call binding the contract method 0xda09c72c.
//
// Solidity: function devAddr() view returns(address)
func (_MasterChef *MasterChefCallerSession) DevAddr() (common.Address, error) {
	return _MasterChef.Contract.DevAddr(&_MasterChef.CallOpts)
}

// DevPercent is a free data retrieval call binding the contract method 0xfc3c28af.
//
// Solidity: function devPercent() view returns(uint256)
func (_MasterChef *MasterChefCaller) DevPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "devPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DevPercent is a free data retrieval call binding the contract method 0xfc3c28af.
//
// Solidity: function devPercent() view returns(uint256)
func (_MasterChef *MasterChefSession) DevPercent() (*big.Int, error) {
	return _MasterChef.Contract.DevPercent(&_MasterChef.CallOpts)
}

// DevPercent is a free data retrieval call binding the contract method 0xfc3c28af.
//
// Solidity: function devPercent() view returns(uint256)
func (_MasterChef *MasterChefCallerSession) DevPercent() (*big.Int, error) {
	return _MasterChef.Contract.DevPercent(&_MasterChef.CallOpts)
}

// InvestorAddr is a free data retrieval call binding the contract method 0xacc4cc50.
//
// Solidity: function investorAddr() view returns(address)
func (_MasterChef *MasterChefCaller) InvestorAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "investorAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InvestorAddr is a free data retrieval call binding the contract method 0xacc4cc50.
//
// Solidity: function investorAddr() view returns(address)
func (_MasterChef *MasterChefSession) InvestorAddr() (common.Address, error) {
	return _MasterChef.Contract.InvestorAddr(&_MasterChef.CallOpts)
}

// InvestorAddr is a free data retrieval call binding the contract method 0xacc4cc50.
//
// Solidity: function investorAddr() view returns(address)
func (_MasterChef *MasterChefCallerSession) InvestorAddr() (common.Address, error) {
	return _MasterChef.Contract.InvestorAddr(&_MasterChef.CallOpts)
}

// InvestorPercent is a free data retrieval call binding the contract method 0x0735b208.
//
// Solidity: function investorPercent() view returns(uint256)
func (_MasterChef *MasterChefCaller) InvestorPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "investorPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvestorPercent is a free data retrieval call binding the contract method 0x0735b208.
//
// Solidity: function investorPercent() view returns(uint256)
func (_MasterChef *MasterChefSession) InvestorPercent() (*big.Int, error) {
	return _MasterChef.Contract.InvestorPercent(&_MasterChef.CallOpts)
}

// InvestorPercent is a free data retrieval call binding the contract method 0x0735b208.
//
// Solidity: function investorPercent() view returns(uint256)
func (_MasterChef *MasterChefCallerSession) InvestorPercent() (*big.Int, error) {
	return _MasterChef.Contract.InvestorPercent(&_MasterChef.CallOpts)
}

// Joe is a free data retrieval call binding the contract method 0xb985a3a0.
//
// Solidity: function joe() view returns(address)
func (_MasterChef *MasterChefCaller) Joe(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "joe")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Joe is a free data retrieval call binding the contract method 0xb985a3a0.
//
// Solidity: function joe() view returns(address)
func (_MasterChef *MasterChefSession) Joe() (common.Address, error) {
	return _MasterChef.Contract.Joe(&_MasterChef.CallOpts)
}

// Joe is a free data retrieval call binding the contract method 0xb985a3a0.
//
// Solidity: function joe() view returns(address)
func (_MasterChef *MasterChefCallerSession) Joe() (common.Address, error) {
	return _MasterChef.Contract.Joe(&_MasterChef.CallOpts)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256)
func (_MasterChef *MasterChefCaller) JoePerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "joePerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256)
func (_MasterChef *MasterChefSession) JoePerSec() (*big.Int, error) {
	return _MasterChef.Contract.JoePerSec(&_MasterChef.CallOpts)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256)
func (_MasterChef *MasterChefCallerSession) JoePerSec() (*big.Int, error) {
	return _MasterChef.Contract.JoePerSec(&_MasterChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChef *MasterChefCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChef *MasterChefSession) Owner() (common.Address, error) {
	return _MasterChef.Contract.Owner(&_MasterChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChef *MasterChefCallerSession) Owner() (common.Address, error) {
	return _MasterChef.Contract.Owner(&_MasterChef.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_MasterChef *MasterChefCaller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "pendingTokens", _pid, _user)

	outstruct := new(struct {
		PendingJoe        *big.Int
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
		PendingBonusToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingJoe = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PendingBonusToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_MasterChef *MasterChefSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _MasterChef.Contract.PendingTokens(&_MasterChef.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_MasterChef *MasterChefCallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _MasterChef.Contract.PendingTokens(&_MasterChef.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardTimestamp, uint256 accJoePerShare, address rewarder)
func (_MasterChef *MasterChefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardTimestamp *big.Int
	AccJoePerShare      *big.Int
	Rewarder            common.Address
}, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken             common.Address
		AllocPoint          *big.Int
		LastRewardTimestamp *big.Int
		AccJoePerShare      *big.Int
		Rewarder            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccJoePerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rewarder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardTimestamp, uint256 accJoePerShare, address rewarder)
func (_MasterChef *MasterChefSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardTimestamp *big.Int
	AccJoePerShare      *big.Int
	Rewarder            common.Address
}, error) {
	return _MasterChef.Contract.PoolInfo(&_MasterChef.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardTimestamp, uint256 accJoePerShare, address rewarder)
func (_MasterChef *MasterChefCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardTimestamp *big.Int
	AccJoePerShare      *big.Int
	Rewarder            common.Address
}, error) {
	return _MasterChef.Contract.PoolInfo(&_MasterChef.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterChef *MasterChefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterChef *MasterChefSession) PoolLength() (*big.Int, error) {
	return _MasterChef.Contract.PoolLength(&_MasterChef.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterChef *MasterChefCallerSession) PoolLength() (*big.Int, error) {
	return _MasterChef.Contract.PoolLength(&_MasterChef.CallOpts)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_MasterChef *MasterChefCaller) RewarderBonusTokenInfo(opts *bind.CallOpts, _pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "rewarderBonusTokenInfo", _pid)

	outstruct := new(struct {
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BonusTokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_MasterChef *MasterChefSession) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _MasterChef.Contract.RewarderBonusTokenInfo(&_MasterChef.CallOpts, _pid)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_MasterChef *MasterChefCallerSession) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _MasterChef.Contract.RewarderBonusTokenInfo(&_MasterChef.CallOpts, _pid)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_MasterChef *MasterChefCaller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "startTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_MasterChef *MasterChefSession) StartTimestamp() (*big.Int, error) {
	return _MasterChef.Contract.StartTimestamp(&_MasterChef.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_MasterChef *MasterChefCallerSession) StartTimestamp() (*big.Int, error) {
	return _MasterChef.Contract.StartTimestamp(&_MasterChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MasterChef *MasterChefCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MasterChef *MasterChefSession) TotalAllocPoint() (*big.Int, error) {
	return _MasterChef.Contract.TotalAllocPoint(&_MasterChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MasterChef *MasterChefCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _MasterChef.Contract.TotalAllocPoint(&_MasterChef.CallOpts)
}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_MasterChef *MasterChefCaller) TreasuryAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "treasuryAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_MasterChef *MasterChefSession) TreasuryAddr() (common.Address, error) {
	return _MasterChef.Contract.TreasuryAddr(&_MasterChef.CallOpts)
}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_MasterChef *MasterChefCallerSession) TreasuryAddr() (common.Address, error) {
	return _MasterChef.Contract.TreasuryAddr(&_MasterChef.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_MasterChef *MasterChefCaller) TreasuryPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "treasuryPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_MasterChef *MasterChefSession) TreasuryPercent() (*big.Int, error) {
	return _MasterChef.Contract.TreasuryPercent(&_MasterChef.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_MasterChef *MasterChefCallerSession) TreasuryPercent() (*big.Int, error) {
	return _MasterChef.Contract.TreasuryPercent(&_MasterChef.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_MasterChef *MasterChefCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _MasterChef.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_MasterChef *MasterChefSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _MasterChef.Contract.UserInfo(&_MasterChef.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_MasterChef *MasterChefCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _MasterChef.Contract.UserInfo(&_MasterChef.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, address _rewarder) returns()
func (_MasterChef *MasterChefTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "add", _allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, address _rewarder) returns()
func (_MasterChef *MasterChefSession) Add(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.Add(&_MasterChef.TransactOpts, _allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, address _rewarder) returns()
func (_MasterChef *MasterChefTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.Add(&_MasterChef.TransactOpts, _allocPoint, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MasterChef *MasterChefTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MasterChef *MasterChefSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.Deposit(&_MasterChef.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MasterChef *MasterChefTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.Deposit(&_MasterChef.TransactOpts, _pid, _amount)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devAddr) returns()
func (_MasterChef *MasterChefTransactor) Dev(opts *bind.TransactOpts, _devAddr common.Address) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "dev", _devAddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devAddr) returns()
func (_MasterChef *MasterChefSession) Dev(_devAddr common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.Dev(&_MasterChef.TransactOpts, _devAddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devAddr) returns()
func (_MasterChef *MasterChefTransactorSession) Dev(_devAddr common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.Dev(&_MasterChef.TransactOpts, _devAddr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChef *MasterChefTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChef *MasterChefSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.EmergencyWithdraw(&_MasterChef.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChef *MasterChefTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.EmergencyWithdraw(&_MasterChef.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChef *MasterChefTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChef *MasterChefSession) MassUpdatePools() (*types.Transaction, error) {
	return _MasterChef.Contract.MassUpdatePools(&_MasterChef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChef *MasterChefTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _MasterChef.Contract.MassUpdatePools(&_MasterChef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChef *MasterChefTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChef *MasterChefSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterChef.Contract.RenounceOwnership(&_MasterChef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChef *MasterChefTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterChef.Contract.RenounceOwnership(&_MasterChef.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_MasterChef *MasterChefTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "set", _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_MasterChef *MasterChefSession) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _MasterChef.Contract.Set(&_MasterChef.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_MasterChef *MasterChefTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _MasterChef.Contract.Set(&_MasterChef.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// SetDevPercent is a paid mutator transaction binding the contract method 0x6eaddad2.
//
// Solidity: function setDevPercent(uint256 _newDevPercent) returns()
func (_MasterChef *MasterChefTransactor) SetDevPercent(opts *bind.TransactOpts, _newDevPercent *big.Int) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "setDevPercent", _newDevPercent)
}

// SetDevPercent is a paid mutator transaction binding the contract method 0x6eaddad2.
//
// Solidity: function setDevPercent(uint256 _newDevPercent) returns()
func (_MasterChef *MasterChefSession) SetDevPercent(_newDevPercent *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.SetDevPercent(&_MasterChef.TransactOpts, _newDevPercent)
}

// SetDevPercent is a paid mutator transaction binding the contract method 0x6eaddad2.
//
// Solidity: function setDevPercent(uint256 _newDevPercent) returns()
func (_MasterChef *MasterChefTransactorSession) SetDevPercent(_newDevPercent *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.SetDevPercent(&_MasterChef.TransactOpts, _newDevPercent)
}

// SetInvestorAddr is a paid mutator transaction binding the contract method 0x0f51f8ff.
//
// Solidity: function setInvestorAddr(address _investorAddr) returns()
func (_MasterChef *MasterChefTransactor) SetInvestorAddr(opts *bind.TransactOpts, _investorAddr common.Address) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "setInvestorAddr", _investorAddr)
}

// SetInvestorAddr is a paid mutator transaction binding the contract method 0x0f51f8ff.
//
// Solidity: function setInvestorAddr(address _investorAddr) returns()
func (_MasterChef *MasterChefSession) SetInvestorAddr(_investorAddr common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.SetInvestorAddr(&_MasterChef.TransactOpts, _investorAddr)
}

// SetInvestorAddr is a paid mutator transaction binding the contract method 0x0f51f8ff.
//
// Solidity: function setInvestorAddr(address _investorAddr) returns()
func (_MasterChef *MasterChefTransactorSession) SetInvestorAddr(_investorAddr common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.SetInvestorAddr(&_MasterChef.TransactOpts, _investorAddr)
}

// SetInvestorPercent is a paid mutator transaction binding the contract method 0x876d3c9c.
//
// Solidity: function setInvestorPercent(uint256 _newInvestorPercent) returns()
func (_MasterChef *MasterChefTransactor) SetInvestorPercent(opts *bind.TransactOpts, _newInvestorPercent *big.Int) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "setInvestorPercent", _newInvestorPercent)
}

// SetInvestorPercent is a paid mutator transaction binding the contract method 0x876d3c9c.
//
// Solidity: function setInvestorPercent(uint256 _newInvestorPercent) returns()
func (_MasterChef *MasterChefSession) SetInvestorPercent(_newInvestorPercent *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.SetInvestorPercent(&_MasterChef.TransactOpts, _newInvestorPercent)
}

// SetInvestorPercent is a paid mutator transaction binding the contract method 0x876d3c9c.
//
// Solidity: function setInvestorPercent(uint256 _newInvestorPercent) returns()
func (_MasterChef *MasterChefTransactorSession) SetInvestorPercent(_newInvestorPercent *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.SetInvestorPercent(&_MasterChef.TransactOpts, _newInvestorPercent)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_MasterChef *MasterChefTransactor) SetTreasuryAddr(opts *bind.TransactOpts, _treasuryAddr common.Address) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "setTreasuryAddr", _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_MasterChef *MasterChefSession) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.SetTreasuryAddr(&_MasterChef.TransactOpts, _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_MasterChef *MasterChefTransactorSession) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.SetTreasuryAddr(&_MasterChef.TransactOpts, _treasuryAddr)
}

// SetTreasuryPercent is a paid mutator transaction binding the contract method 0x89a2bc25.
//
// Solidity: function setTreasuryPercent(uint256 _newTreasuryPercent) returns()
func (_MasterChef *MasterChefTransactor) SetTreasuryPercent(opts *bind.TransactOpts, _newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "setTreasuryPercent", _newTreasuryPercent)
}

// SetTreasuryPercent is a paid mutator transaction binding the contract method 0x89a2bc25.
//
// Solidity: function setTreasuryPercent(uint256 _newTreasuryPercent) returns()
func (_MasterChef *MasterChefSession) SetTreasuryPercent(_newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.SetTreasuryPercent(&_MasterChef.TransactOpts, _newTreasuryPercent)
}

// SetTreasuryPercent is a paid mutator transaction binding the contract method 0x89a2bc25.
//
// Solidity: function setTreasuryPercent(uint256 _newTreasuryPercent) returns()
func (_MasterChef *MasterChefTransactorSession) SetTreasuryPercent(_newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.SetTreasuryPercent(&_MasterChef.TransactOpts, _newTreasuryPercent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChef *MasterChefTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChef *MasterChefSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.TransferOwnership(&_MasterChef.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChef *MasterChefTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterChef.Contract.TransferOwnership(&_MasterChef.TransactOpts, newOwner)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _joePerSec) returns()
func (_MasterChef *MasterChefTransactor) UpdateEmissionRate(opts *bind.TransactOpts, _joePerSec *big.Int) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "updateEmissionRate", _joePerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _joePerSec) returns()
func (_MasterChef *MasterChefSession) UpdateEmissionRate(_joePerSec *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.UpdateEmissionRate(&_MasterChef.TransactOpts, _joePerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _joePerSec) returns()
func (_MasterChef *MasterChefTransactorSession) UpdateEmissionRate(_joePerSec *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.UpdateEmissionRate(&_MasterChef.TransactOpts, _joePerSec)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterChef *MasterChefTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterChef *MasterChefSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.UpdatePool(&_MasterChef.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterChef *MasterChefTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.UpdatePool(&_MasterChef.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChef *MasterChefTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChef.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChef *MasterChefSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.Withdraw(&_MasterChef.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChef *MasterChefTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChef.Contract.Withdraw(&_MasterChef.TransactOpts, _pid, _amount)
}

// MasterChefAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the MasterChef contract.
type MasterChefAddIterator struct {
	Event *MasterChefAdd // Event containing the contract specifics and raw log

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
func (it *MasterChefAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefAdd)
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
		it.Event = new(MasterChefAdd)
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
func (it *MasterChefAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefAdd represents a Add event raised by the MasterChef contract.
type MasterChefAdd struct {
	Pid        *big.Int
	AllocPoint *big.Int
	LpToken    common.Address
	Rewarder   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_MasterChef *MasterChefFilterer) FilterAdd(opts *bind.FilterOpts, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (*MasterChefAddIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefAddIterator{contract: _MasterChef.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_MasterChef *MasterChefFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *MasterChefAdd, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefAdd)
				if err := _MasterChef.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_MasterChef *MasterChefFilterer) ParseAdd(log types.Log) (*MasterChefAdd, error) {
	event := new(MasterChefAdd)
	if err := _MasterChef.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MasterChef contract.
type MasterChefDepositIterator struct {
	Event *MasterChefDeposit // Event containing the contract specifics and raw log

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
func (it *MasterChefDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDeposit)
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
		it.Event = new(MasterChefDeposit)
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
func (it *MasterChefDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDeposit represents a Deposit event raised by the MasterChef contract.
type MasterChefDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefDepositIterator{contract: _MasterChef.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MasterChefDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDeposit)
				if err := _MasterChef.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) ParseDeposit(log types.Log) (*MasterChefDeposit, error) {
	event := new(MasterChefDeposit)
	if err := _MasterChef.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the MasterChef contract.
type MasterChefEmergencyWithdrawIterator struct {
	Event *MasterChefEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *MasterChefEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefEmergencyWithdraw)
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
		it.Event = new(MasterChefEmergencyWithdraw)
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
func (it *MasterChefEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefEmergencyWithdraw represents a EmergencyWithdraw event raised by the MasterChef contract.
type MasterChefEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefEmergencyWithdrawIterator{contract: _MasterChef.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *MasterChefEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefEmergencyWithdraw)
				if err := _MasterChef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) ParseEmergencyWithdraw(log types.Log) (*MasterChefEmergencyWithdraw, error) {
	event := new(MasterChefEmergencyWithdraw)
	if err := _MasterChef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the MasterChef contract.
type MasterChefHarvestIterator struct {
	Event *MasterChefHarvest // Event containing the contract specifics and raw log

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
func (it *MasterChefHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefHarvest)
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
		it.Event = new(MasterChefHarvest)
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
func (it *MasterChefHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefHarvest represents a Harvest event raised by the MasterChef contract.
type MasterChefHarvest struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefHarvestIterator{contract: _MasterChef.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *MasterChefHarvest, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefHarvest)
				if err := _MasterChef.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) ParseHarvest(log types.Log) (*MasterChefHarvest, error) {
	event := new(MasterChefHarvest)
	if err := _MasterChef.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MasterChef contract.
type MasterChefOwnershipTransferredIterator struct {
	Event *MasterChefOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MasterChefOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefOwnershipTransferred)
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
		it.Event = new(MasterChefOwnershipTransferred)
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
func (it *MasterChefOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefOwnershipTransferred represents a OwnershipTransferred event raised by the MasterChef contract.
type MasterChefOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterChef *MasterChefFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MasterChefOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefOwnershipTransferredIterator{contract: _MasterChef.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterChef *MasterChefFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MasterChefOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefOwnershipTransferred)
				if err := _MasterChef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MasterChef *MasterChefFilterer) ParseOwnershipTransferred(log types.Log) (*MasterChefOwnershipTransferred, error) {
	event := new(MasterChefOwnershipTransferred)
	if err := _MasterChef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the MasterChef contract.
type MasterChefSetIterator struct {
	Event *MasterChefSet // Event containing the contract specifics and raw log

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
func (it *MasterChefSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefSet)
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
		it.Event = new(MasterChefSet)
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
func (it *MasterChefSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefSet represents a Set event raised by the MasterChef contract.
type MasterChefSet struct {
	Pid        *big.Int
	AllocPoint *big.Int
	Rewarder   common.Address
	Overwrite  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_MasterChef *MasterChefFilterer) FilterSet(opts *bind.FilterOpts, pid []*big.Int, rewarder []common.Address) (*MasterChefSetIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefSetIterator{contract: _MasterChef.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_MasterChef *MasterChefFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *MasterChefSet, pid []*big.Int, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefSet)
				if err := _MasterChef.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_MasterChef *MasterChefFilterer) ParseSet(log types.Log) (*MasterChefSet, error) {
	event := new(MasterChefSet)
	if err := _MasterChef.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefSetDevAddressIterator is returned from FilterSetDevAddress and is used to iterate over the raw logs and unpacked data for SetDevAddress events raised by the MasterChef contract.
type MasterChefSetDevAddressIterator struct {
	Event *MasterChefSetDevAddress // Event containing the contract specifics and raw log

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
func (it *MasterChefSetDevAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefSetDevAddress)
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
		it.Event = new(MasterChefSetDevAddress)
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
func (it *MasterChefSetDevAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefSetDevAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefSetDevAddress represents a SetDevAddress event raised by the MasterChef contract.
type MasterChefSetDevAddress struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetDevAddress is a free log retrieval operation binding the contract event 0x618c54559e94f1499a808aad71ee8729f8e74e8c48e979616328ce493a1a52e7.
//
// Solidity: event SetDevAddress(address indexed oldAddress, address indexed newAddress)
func (_MasterChef *MasterChefFilterer) FilterSetDevAddress(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*MasterChefSetDevAddressIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "SetDevAddress", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefSetDevAddressIterator{contract: _MasterChef.contract, event: "SetDevAddress", logs: logs, sub: sub}, nil
}

// WatchSetDevAddress is a free log subscription operation binding the contract event 0x618c54559e94f1499a808aad71ee8729f8e74e8c48e979616328ce493a1a52e7.
//
// Solidity: event SetDevAddress(address indexed oldAddress, address indexed newAddress)
func (_MasterChef *MasterChefFilterer) WatchSetDevAddress(opts *bind.WatchOpts, sink chan<- *MasterChefSetDevAddress, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "SetDevAddress", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefSetDevAddress)
				if err := _MasterChef.contract.UnpackLog(event, "SetDevAddress", log); err != nil {
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

// ParseSetDevAddress is a log parse operation binding the contract event 0x618c54559e94f1499a808aad71ee8729f8e74e8c48e979616328ce493a1a52e7.
//
// Solidity: event SetDevAddress(address indexed oldAddress, address indexed newAddress)
func (_MasterChef *MasterChefFilterer) ParseSetDevAddress(log types.Log) (*MasterChefSetDevAddress, error) {
	event := new(MasterChefSetDevAddress)
	if err := _MasterChef.contract.UnpackLog(event, "SetDevAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefUpdateEmissionRateIterator is returned from FilterUpdateEmissionRate and is used to iterate over the raw logs and unpacked data for UpdateEmissionRate events raised by the MasterChef contract.
type MasterChefUpdateEmissionRateIterator struct {
	Event *MasterChefUpdateEmissionRate // Event containing the contract specifics and raw log

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
func (it *MasterChefUpdateEmissionRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefUpdateEmissionRate)
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
		it.Event = new(MasterChefUpdateEmissionRate)
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
func (it *MasterChefUpdateEmissionRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefUpdateEmissionRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefUpdateEmissionRate represents a UpdateEmissionRate event raised by the MasterChef contract.
type MasterChefUpdateEmissionRate struct {
	User      common.Address
	JoePerSec *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateEmissionRate is a free log retrieval operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 _joePerSec)
func (_MasterChef *MasterChefFilterer) FilterUpdateEmissionRate(opts *bind.FilterOpts, user []common.Address) (*MasterChefUpdateEmissionRateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "UpdateEmissionRate", userRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefUpdateEmissionRateIterator{contract: _MasterChef.contract, event: "UpdateEmissionRate", logs: logs, sub: sub}, nil
}

// WatchUpdateEmissionRate is a free log subscription operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 _joePerSec)
func (_MasterChef *MasterChefFilterer) WatchUpdateEmissionRate(opts *bind.WatchOpts, sink chan<- *MasterChefUpdateEmissionRate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "UpdateEmissionRate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefUpdateEmissionRate)
				if err := _MasterChef.contract.UnpackLog(event, "UpdateEmissionRate", log); err != nil {
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

// ParseUpdateEmissionRate is a log parse operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 _joePerSec)
func (_MasterChef *MasterChefFilterer) ParseUpdateEmissionRate(log types.Log) (*MasterChefUpdateEmissionRate, error) {
	event := new(MasterChefUpdateEmissionRate)
	if err := _MasterChef.contract.UnpackLog(event, "UpdateEmissionRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefUpdatePoolIterator is returned from FilterUpdatePool and is used to iterate over the raw logs and unpacked data for UpdatePool events raised by the MasterChef contract.
type MasterChefUpdatePoolIterator struct {
	Event *MasterChefUpdatePool // Event containing the contract specifics and raw log

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
func (it *MasterChefUpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefUpdatePool)
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
		it.Event = new(MasterChefUpdatePool)
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
func (it *MasterChefUpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefUpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefUpdatePool represents a UpdatePool event raised by the MasterChef contract.
type MasterChefUpdatePool struct {
	Pid                 *big.Int
	LastRewardTimestamp *big.Int
	LpSupply            *big.Int
	AccJoePerShare      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdatePool is a free log retrieval operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accJoePerShare)
func (_MasterChef *MasterChefFilterer) FilterUpdatePool(opts *bind.FilterOpts, pid []*big.Int) (*MasterChefUpdatePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefUpdatePoolIterator{contract: _MasterChef.contract, event: "UpdatePool", logs: logs, sub: sub}, nil
}

// WatchUpdatePool is a free log subscription operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accJoePerShare)
func (_MasterChef *MasterChefFilterer) WatchUpdatePool(opts *bind.WatchOpts, sink chan<- *MasterChefUpdatePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefUpdatePool)
				if err := _MasterChef.contract.UnpackLog(event, "UpdatePool", log); err != nil {
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

// ParseUpdatePool is a log parse operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accJoePerShare)
func (_MasterChef *MasterChefFilterer) ParseUpdatePool(log types.Log) (*MasterChefUpdatePool, error) {
	event := new(MasterChefUpdatePool)
	if err := _MasterChef.contract.UnpackLog(event, "UpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MasterChef contract.
type MasterChefWithdrawIterator struct {
	Event *MasterChefWithdraw // Event containing the contract specifics and raw log

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
func (it *MasterChefWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefWithdraw)
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
		it.Event = new(MasterChefWithdraw)
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
func (it *MasterChefWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefWithdraw represents a Withdraw event raised by the MasterChef contract.
type MasterChefWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefWithdrawIterator{contract: _MasterChef.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MasterChefWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChef.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefWithdraw)
				if err := _MasterChef.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChef *MasterChefFilterer) ParseWithdraw(log types.Log) (*MasterChefWithdraw, error) {
	event := new(MasterChefWithdraw)
	if err := _MasterChef.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
