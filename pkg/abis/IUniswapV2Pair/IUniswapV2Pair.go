// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iuniswapv2pair

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
	_ = abi.ConvertType
)

// Iuniswapv2pairMetaData contains all meta data concerning the Iuniswapv2pair contract.
var Iuniswapv2pairMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Iuniswapv2pairABI is the input ABI used to generate the binding from.
// Deprecated: Use Iuniswapv2pairMetaData.ABI instead.
var Iuniswapv2pairABI = Iuniswapv2pairMetaData.ABI

// Iuniswapv2pair is an auto generated Go binding around an Ethereum contract.
type Iuniswapv2pair struct {
	Iuniswapv2pairCaller     // Read-only binding to the contract
	Iuniswapv2pairTransactor // Write-only binding to the contract
	Iuniswapv2pairFilterer   // Log filterer for contract events
}

// Iuniswapv2pairCaller is an auto generated read-only Go binding around an Ethereum contract.
type Iuniswapv2pairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iuniswapv2pairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Iuniswapv2pairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iuniswapv2pairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iuniswapv2pairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iuniswapv2pairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iuniswapv2pairSession struct {
	Contract     *Iuniswapv2pair   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Iuniswapv2pairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iuniswapv2pairCallerSession struct {
	Contract *Iuniswapv2pairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Iuniswapv2pairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iuniswapv2pairTransactorSession struct {
	Contract     *Iuniswapv2pairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Iuniswapv2pairRaw is an auto generated low-level Go binding around an Ethereum contract.
type Iuniswapv2pairRaw struct {
	Contract *Iuniswapv2pair // Generic contract binding to access the raw methods on
}

// Iuniswapv2pairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iuniswapv2pairCallerRaw struct {
	Contract *Iuniswapv2pairCaller // Generic read-only contract binding to access the raw methods on
}

// Iuniswapv2pairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iuniswapv2pairTransactorRaw struct {
	Contract *Iuniswapv2pairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIuniswapv2pair creates a new instance of Iuniswapv2pair, bound to a specific deployed contract.
func NewIuniswapv2pair(address common.Address, backend bind.ContractBackend) (*Iuniswapv2pair, error) {
	contract, err := bindIuniswapv2pair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pair{Iuniswapv2pairCaller: Iuniswapv2pairCaller{contract: contract}, Iuniswapv2pairTransactor: Iuniswapv2pairTransactor{contract: contract}, Iuniswapv2pairFilterer: Iuniswapv2pairFilterer{contract: contract}}, nil
}

// NewIuniswapv2pairCaller creates a new read-only instance of Iuniswapv2pair, bound to a specific deployed contract.
func NewIuniswapv2pairCaller(address common.Address, caller bind.ContractCaller) (*Iuniswapv2pairCaller, error) {
	contract, err := bindIuniswapv2pair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pairCaller{contract: contract}, nil
}

// NewIuniswapv2pairTransactor creates a new write-only instance of Iuniswapv2pair, bound to a specific deployed contract.
func NewIuniswapv2pairTransactor(address common.Address, transactor bind.ContractTransactor) (*Iuniswapv2pairTransactor, error) {
	contract, err := bindIuniswapv2pair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pairTransactor{contract: contract}, nil
}

// NewIuniswapv2pairFilterer creates a new log filterer instance of Iuniswapv2pair, bound to a specific deployed contract.
func NewIuniswapv2pairFilterer(address common.Address, filterer bind.ContractFilterer) (*Iuniswapv2pairFilterer, error) {
	contract, err := bindIuniswapv2pair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pairFilterer{contract: contract}, nil
}

// bindIuniswapv2pair binds a generic wrapper to an already deployed contract.
func bindIuniswapv2pair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Iuniswapv2pairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iuniswapv2pair *Iuniswapv2pairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iuniswapv2pair.Contract.Iuniswapv2pairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iuniswapv2pair *Iuniswapv2pairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Iuniswapv2pairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iuniswapv2pair *Iuniswapv2pairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Iuniswapv2pairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iuniswapv2pair *Iuniswapv2pairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iuniswapv2pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iuniswapv2pair *Iuniswapv2pairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iuniswapv2pair *Iuniswapv2pairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Iuniswapv2pair *Iuniswapv2pairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Iuniswapv2pair.Contract.DOMAINSEPARATOR(&_Iuniswapv2pair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Iuniswapv2pair.Contract.DOMAINSEPARATOR(&_Iuniswapv2pair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.MINIMUMLIQUIDITY(&_Iuniswapv2pair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.MINIMUMLIQUIDITY(&_Iuniswapv2pair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_Iuniswapv2pair *Iuniswapv2pairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Iuniswapv2pair.Contract.PERMITTYPEHASH(&_Iuniswapv2pair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Iuniswapv2pair.Contract.PERMITTYPEHASH(&_Iuniswapv2pair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Iuniswapv2pair.Contract.Allowance(&_Iuniswapv2pair.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Iuniswapv2pair.Contract.Allowance(&_Iuniswapv2pair.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Iuniswapv2pair.Contract.BalanceOf(&_Iuniswapv2pair.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Iuniswapv2pair.Contract.BalanceOf(&_Iuniswapv2pair.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Decimals() (uint8, error) {
	return _Iuniswapv2pair.Contract.Decimals(&_Iuniswapv2pair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Decimals() (uint8, error) {
	return _Iuniswapv2pair.Contract.Decimals(&_Iuniswapv2pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Factory() (common.Address, error) {
	return _Iuniswapv2pair.Contract.Factory(&_Iuniswapv2pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Factory() (common.Address, error) {
	return _Iuniswapv2pair.Contract.Factory(&_Iuniswapv2pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Iuniswapv2pair *Iuniswapv2pairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Iuniswapv2pair.Contract.GetReserves(&_Iuniswapv2pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Iuniswapv2pair.Contract.GetReserves(&_Iuniswapv2pair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairSession) KLast() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.KLast(&_Iuniswapv2pair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) KLast() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.KLast(&_Iuniswapv2pair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Name() (string, error) {
	return _Iuniswapv2pair.Contract.Name(&_Iuniswapv2pair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Name() (string, error) {
	return _Iuniswapv2pair.Contract.Name(&_Iuniswapv2pair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Iuniswapv2pair.Contract.Nonces(&_Iuniswapv2pair.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Iuniswapv2pair.Contract.Nonces(&_Iuniswapv2pair.CallOpts, owner)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Price0CumulativeLast() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.Price0CumulativeLast(&_Iuniswapv2pair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.Price0CumulativeLast(&_Iuniswapv2pair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Price1CumulativeLast() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.Price1CumulativeLast(&_Iuniswapv2pair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.Price1CumulativeLast(&_Iuniswapv2pair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Symbol() (string, error) {
	return _Iuniswapv2pair.Contract.Symbol(&_Iuniswapv2pair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Symbol() (string, error) {
	return _Iuniswapv2pair.Contract.Symbol(&_Iuniswapv2pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Token0() (common.Address, error) {
	return _Iuniswapv2pair.Contract.Token0(&_Iuniswapv2pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Token0() (common.Address, error) {
	return _Iuniswapv2pair.Contract.Token0(&_Iuniswapv2pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Token1() (common.Address, error) {
	return _Iuniswapv2pair.Contract.Token1(&_Iuniswapv2pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) Token1() (common.Address, error) {
	return _Iuniswapv2pair.Contract.Token1(&_Iuniswapv2pair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv2pair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairSession) TotalSupply() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.TotalSupply(&_Iuniswapv2pair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Iuniswapv2pair *Iuniswapv2pairCallerSession) TotalSupply() (*big.Int, error) {
	return _Iuniswapv2pair.Contract.TotalSupply(&_Iuniswapv2pair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Approve(&_Iuniswapv2pair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Approve(&_Iuniswapv2pair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Burn(&_Iuniswapv2pair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Burn(&_Iuniswapv2pair.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) Initialize(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "initialize", arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_Iuniswapv2pair *Iuniswapv2pairSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Initialize(&_Iuniswapv2pair.TransactOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Initialize(&_Iuniswapv2pair.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Mint(&_Iuniswapv2pair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Mint(&_Iuniswapv2pair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Iuniswapv2pair *Iuniswapv2pairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Permit(&_Iuniswapv2pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Permit(&_Iuniswapv2pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Iuniswapv2pair *Iuniswapv2pairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Skim(&_Iuniswapv2pair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Skim(&_Iuniswapv2pair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Iuniswapv2pair *Iuniswapv2pairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Swap(&_Iuniswapv2pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Swap(&_Iuniswapv2pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Iuniswapv2pair *Iuniswapv2pairSession) Sync() (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Sync(&_Iuniswapv2pair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) Sync() (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Sync(&_Iuniswapv2pair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Iuniswapv2pair *Iuniswapv2pairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Transfer(&_Iuniswapv2pair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.Transfer(&_Iuniswapv2pair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Iuniswapv2pair *Iuniswapv2pairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2pair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Iuniswapv2pair *Iuniswapv2pairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.TransferFrom(&_Iuniswapv2pair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Iuniswapv2pair *Iuniswapv2pairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2pair.Contract.TransferFrom(&_Iuniswapv2pair.TransactOpts, from, to, value)
}

// Iuniswapv2pairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Iuniswapv2pair contract.
type Iuniswapv2pairApprovalIterator struct {
	Event *Iuniswapv2pairApproval // Event containing the contract specifics and raw log

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
func (it *Iuniswapv2pairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iuniswapv2pairApproval)
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
		it.Event = new(Iuniswapv2pairApproval)
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
func (it *Iuniswapv2pairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iuniswapv2pairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iuniswapv2pairApproval represents a Approval event raised by the Iuniswapv2pair contract.
type Iuniswapv2pairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Iuniswapv2pairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pairApprovalIterator{contract: _Iuniswapv2pair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Iuniswapv2pairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iuniswapv2pairApproval)
				if err := _Iuniswapv2pair.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) ParseApproval(log types.Log) (*Iuniswapv2pairApproval, error) {
	event := new(Iuniswapv2pairApproval)
	if err := _Iuniswapv2pair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Iuniswapv2pairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Iuniswapv2pair contract.
type Iuniswapv2pairBurnIterator struct {
	Event *Iuniswapv2pairBurn // Event containing the contract specifics and raw log

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
func (it *Iuniswapv2pairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iuniswapv2pairBurn)
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
		it.Event = new(Iuniswapv2pairBurn)
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
func (it *Iuniswapv2pairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iuniswapv2pairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iuniswapv2pairBurn represents a Burn event raised by the Iuniswapv2pair contract.
type Iuniswapv2pairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*Iuniswapv2pairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pairBurnIterator{contract: _Iuniswapv2pair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *Iuniswapv2pairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iuniswapv2pairBurn)
				if err := _Iuniswapv2pair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) ParseBurn(log types.Log) (*Iuniswapv2pairBurn, error) {
	event := new(Iuniswapv2pairBurn)
	if err := _Iuniswapv2pair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Iuniswapv2pairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Iuniswapv2pair contract.
type Iuniswapv2pairMintIterator struct {
	Event *Iuniswapv2pairMint // Event containing the contract specifics and raw log

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
func (it *Iuniswapv2pairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iuniswapv2pairMint)
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
		it.Event = new(Iuniswapv2pairMint)
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
func (it *Iuniswapv2pairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iuniswapv2pairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iuniswapv2pairMint represents a Mint event raised by the Iuniswapv2pair contract.
type Iuniswapv2pairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*Iuniswapv2pairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pairMintIterator{contract: _Iuniswapv2pair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Iuniswapv2pairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iuniswapv2pairMint)
				if err := _Iuniswapv2pair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) ParseMint(log types.Log) (*Iuniswapv2pairMint, error) {
	event := new(Iuniswapv2pairMint)
	if err := _Iuniswapv2pair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Iuniswapv2pairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Iuniswapv2pair contract.
type Iuniswapv2pairSwapIterator struct {
	Event *Iuniswapv2pairSwap // Event containing the contract specifics and raw log

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
func (it *Iuniswapv2pairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iuniswapv2pairSwap)
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
		it.Event = new(Iuniswapv2pairSwap)
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
func (it *Iuniswapv2pairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iuniswapv2pairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iuniswapv2pairSwap represents a Swap event raised by the Iuniswapv2pair contract.
type Iuniswapv2pairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*Iuniswapv2pairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pairSwapIterator{contract: _Iuniswapv2pair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *Iuniswapv2pairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iuniswapv2pairSwap)
				if err := _Iuniswapv2pair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) ParseSwap(log types.Log) (*Iuniswapv2pairSwap, error) {
	event := new(Iuniswapv2pairSwap)
	if err := _Iuniswapv2pair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Iuniswapv2pairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Iuniswapv2pair contract.
type Iuniswapv2pairSyncIterator struct {
	Event *Iuniswapv2pairSync // Event containing the contract specifics and raw log

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
func (it *Iuniswapv2pairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iuniswapv2pairSync)
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
		it.Event = new(Iuniswapv2pairSync)
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
func (it *Iuniswapv2pairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iuniswapv2pairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iuniswapv2pairSync represents a Sync event raised by the Iuniswapv2pair contract.
type Iuniswapv2pairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) FilterSync(opts *bind.FilterOpts) (*Iuniswapv2pairSyncIterator, error) {

	logs, sub, err := _Iuniswapv2pair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pairSyncIterator{contract: _Iuniswapv2pair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *Iuniswapv2pairSync) (event.Subscription, error) {

	logs, sub, err := _Iuniswapv2pair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iuniswapv2pairSync)
				if err := _Iuniswapv2pair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) ParseSync(log types.Log) (*Iuniswapv2pairSync, error) {
	event := new(Iuniswapv2pairSync)
	if err := _Iuniswapv2pair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Iuniswapv2pairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Iuniswapv2pair contract.
type Iuniswapv2pairTransferIterator struct {
	Event *Iuniswapv2pairTransfer // Event containing the contract specifics and raw log

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
func (it *Iuniswapv2pairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iuniswapv2pairTransfer)
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
		it.Event = new(Iuniswapv2pairTransfer)
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
func (it *Iuniswapv2pairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iuniswapv2pairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iuniswapv2pairTransfer represents a Transfer event raised by the Iuniswapv2pair contract.
type Iuniswapv2pairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Iuniswapv2pairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2pairTransferIterator{contract: _Iuniswapv2pair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Iuniswapv2pairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iuniswapv2pair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iuniswapv2pairTransfer)
				if err := _Iuniswapv2pair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Iuniswapv2pair *Iuniswapv2pairFilterer) ParseTransfer(log types.Log) (*Iuniswapv2pairTransfer, error) {
	event := new(Iuniswapv2pairTransfer)
	if err := _Iuniswapv2pair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
