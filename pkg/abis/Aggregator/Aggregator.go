// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggregator

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

// AggregatorMetaData contains all meta data concerning the Aggregator contract.
var AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"getPairKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"keyToBestBuy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVolume\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dexAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"keyToBestSell\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVolume\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dexAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tradeContract\",\"type\":\"address\"}],\"name\":\"updateBestAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tradeContract\",\"type\":\"address\"}],\"name\":\"updateBestBuy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteVolume\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tradeContract\",\"type\":\"address\"}],\"name\":\"updateBestSell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610bc98061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061007b575f3560e01c8063c638987911610059578063c6389879146100e7578063cade791f14610103578063df791e5014610137578063eac06175146101535761007b565b806358618e7d1461007f57806363a830601461009b578063b9e16722146100cb575b5f80fd5b61009960048036038101906100949190610929565b610187565b005b6100b560048036038101906100b091906109a0565b6102f8565b6040516100c291906109f6565b60405180910390f35b6100e560048036038101906100e09190610929565b61038f565b005b61010160048036038101906100fc9190610929565b610606565b005b61011d60048036038101906101189190610a39565b610776565b60405161012e959493929190610a82565b60405180910390f35b610151600480360381019061014c9190610ad3565b610805565b005b61016d60048036038101906101689190610a39565b61080a565b60405161017e959493929190610a82565b60405180910390f35b6040518060a001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018481526020018273ffffffffffffffffffffffffffffffffffffffff1681525060015f6101fd88886102f8565b81526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505050505050565b5f8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16101561035d578282604051602001610340929190610b68565b604051602081830303815290604052805190602001209050610389565b8183604051602001610370929190610b68565b6040516020818303038152906040528051906020012090505b92915050565b5f6040518060a001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018581526020018373ffffffffffffffffffffffffffffffffffffffff168152509050805f8061040889896102f8565b81526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050508060015f61050a89896102f8565b81526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050505050505050565b6040518060a001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018481526020018273ffffffffffffffffffffffffffffffffffffffff168152505f8061067b88886102f8565b81526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505050505050565b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015490806003015490806004015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905085565b505050565b5f602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015490806003015490806004015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905085565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108c58261089c565b9050919050565b6108d5816108bb565b81146108df575f80fd5b50565b5f813590506108f0816108cc565b92915050565b5f819050919050565b610908816108f6565b8114610912575f80fd5b50565b5f81359050610923816108ff565b92915050565b5f805f805f60a0868803121561094257610941610898565b5b5f61094f888289016108e2565b9550506020610960888289016108e2565b945050604061097188828901610915565b935050606061098288828901610915565b9250506080610993888289016108e2565b9150509295509295909350565b5f80604083850312156109b6576109b5610898565b5b5f6109c3858286016108e2565b92505060206109d4858286016108e2565b9150509250929050565b5f819050919050565b6109f0816109de565b82525050565b5f602082019050610a095f8301846109e7565b92915050565b610a18816109de565b8114610a22575f80fd5b50565b5f81359050610a3381610a0f565b92915050565b5f60208284031215610a4e57610a4d610898565b5b5f610a5b84828501610a25565b91505092915050565b610a6d816108bb565b82525050565b610a7c816108f6565b82525050565b5f60a082019050610a955f830188610a64565b610aa26020830187610a64565b610aaf6040830186610a73565b610abc6060830185610a73565b610ac96080830184610a64565b9695505050505050565b5f805f60608486031215610aea57610ae9610898565b5b5f610af7868287016108e2565b9350506020610b08868287016108e2565b9250506040610b1986828701610915565b9150509250925092565b5f8160601b9050919050565b5f610b3982610b23565b9050919050565b5f610b4a82610b2f565b9050919050565b610b62610b5d826108bb565b610b40565b82525050565b5f610b738285610b51565b601482019150610b838284610b51565b601482019150819050939250505056fea26469706673582212209b22d536881d6375ed31577f1f5a3ec3d80b0a9c1808bb6f550cd5bb451ae3ca64736f6c63430008170033",
}

// AggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorMetaData.ABI instead.
var AggregatorABI = AggregatorMetaData.ABI

// AggregatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AggregatorMetaData.Bin instead.
var AggregatorBin = AggregatorMetaData.Bin

// DeployAggregator deploys a new Ethereum contract, binding an instance of Aggregator to it.
func DeployAggregator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Aggregator, error) {
	parsed, err := AggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AggregatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aggregator{AggregatorCaller: AggregatorCaller{contract: contract}, AggregatorTransactor: AggregatorTransactor{contract: contract}, AggregatorFilterer: AggregatorFilterer{contract: contract}}, nil
}

// Aggregator is an auto generated Go binding around an Ethereum contract.
type Aggregator struct {
	AggregatorCaller     // Read-only binding to the contract
	AggregatorTransactor // Write-only binding to the contract
	AggregatorFilterer   // Log filterer for contract events
}

// AggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorSession struct {
	Contract     *Aggregator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorCallerSession struct {
	Contract *AggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorTransactorSession struct {
	Contract     *AggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorRaw struct {
	Contract *Aggregator // Generic contract binding to access the raw methods on
}

// AggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorCallerRaw struct {
	Contract *AggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorTransactorRaw struct {
	Contract *AggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregator creates a new instance of Aggregator, bound to a specific deployed contract.
func NewAggregator(address common.Address, backend bind.ContractBackend) (*Aggregator, error) {
	contract, err := bindAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggregator{AggregatorCaller: AggregatorCaller{contract: contract}, AggregatorTransactor: AggregatorTransactor{contract: contract}, AggregatorFilterer: AggregatorFilterer{contract: contract}}, nil
}

// NewAggregatorCaller creates a new read-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorCaller(address common.Address, caller bind.ContractCaller) (*AggregatorCaller, error) {
	contract, err := bindAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorCaller{contract: contract}, nil
}

// NewAggregatorTransactor creates a new write-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorTransactor, error) {
	contract, err := bindAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorTransactor{contract: contract}, nil
}

// NewAggregatorFilterer creates a new log filterer instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorFilterer, error) {
	contract, err := bindAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorFilterer{contract: contract}, nil
}

// bindAggregator binds a generic wrapper to an already deployed contract.
func bindAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.AggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transact(opts, method, params...)
}

// GetPairKey is a free data retrieval call binding the contract method 0x63a83060.
//
// Solidity: function getPairKey(address base, address quote) pure returns(bytes32)
func (_Aggregator *AggregatorCaller) GetPairKey(opts *bind.CallOpts, base common.Address, quote common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "getPairKey", base, quote)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPairKey is a free data retrieval call binding the contract method 0x63a83060.
//
// Solidity: function getPairKey(address base, address quote) pure returns(bytes32)
func (_Aggregator *AggregatorSession) GetPairKey(base common.Address, quote common.Address) ([32]byte, error) {
	return _Aggregator.Contract.GetPairKey(&_Aggregator.CallOpts, base, quote)
}

// GetPairKey is a free data retrieval call binding the contract method 0x63a83060.
//
// Solidity: function getPairKey(address base, address quote) pure returns(bytes32)
func (_Aggregator *AggregatorCallerSession) GetPairKey(base common.Address, quote common.Address) ([32]byte, error) {
	return _Aggregator.Contract.GetPairKey(&_Aggregator.CallOpts, base, quote)
}

// KeyToBestBuy is a free data retrieval call binding the contract method 0xeac06175.
//
// Solidity: function keyToBestBuy(bytes32 ) view returns(address quote, address base, uint256 quoteVolume, uint256 baseVolume, address dexAddress)
func (_Aggregator *AggregatorCaller) KeyToBestBuy(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Quote       common.Address
	Base        common.Address
	QuoteVolume *big.Int
	BaseVolume  *big.Int
	DexAddress  common.Address
}, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "keyToBestBuy", arg0)

	outstruct := new(struct {
		Quote       common.Address
		Base        common.Address
		QuoteVolume *big.Int
		BaseVolume  *big.Int
		DexAddress  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quote = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Base = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.QuoteVolume = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BaseVolume = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DexAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// KeyToBestBuy is a free data retrieval call binding the contract method 0xeac06175.
//
// Solidity: function keyToBestBuy(bytes32 ) view returns(address quote, address base, uint256 quoteVolume, uint256 baseVolume, address dexAddress)
func (_Aggregator *AggregatorSession) KeyToBestBuy(arg0 [32]byte) (struct {
	Quote       common.Address
	Base        common.Address
	QuoteVolume *big.Int
	BaseVolume  *big.Int
	DexAddress  common.Address
}, error) {
	return _Aggregator.Contract.KeyToBestBuy(&_Aggregator.CallOpts, arg0)
}

// KeyToBestBuy is a free data retrieval call binding the contract method 0xeac06175.
//
// Solidity: function keyToBestBuy(bytes32 ) view returns(address quote, address base, uint256 quoteVolume, uint256 baseVolume, address dexAddress)
func (_Aggregator *AggregatorCallerSession) KeyToBestBuy(arg0 [32]byte) (struct {
	Quote       common.Address
	Base        common.Address
	QuoteVolume *big.Int
	BaseVolume  *big.Int
	DexAddress  common.Address
}, error) {
	return _Aggregator.Contract.KeyToBestBuy(&_Aggregator.CallOpts, arg0)
}

// KeyToBestSell is a free data retrieval call binding the contract method 0xcade791f.
//
// Solidity: function keyToBestSell(bytes32 ) view returns(address quote, address base, uint256 quoteVolume, uint256 baseVolume, address dexAddress)
func (_Aggregator *AggregatorCaller) KeyToBestSell(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Quote       common.Address
	Base        common.Address
	QuoteVolume *big.Int
	BaseVolume  *big.Int
	DexAddress  common.Address
}, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "keyToBestSell", arg0)

	outstruct := new(struct {
		Quote       common.Address
		Base        common.Address
		QuoteVolume *big.Int
		BaseVolume  *big.Int
		DexAddress  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quote = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Base = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.QuoteVolume = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BaseVolume = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DexAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// KeyToBestSell is a free data retrieval call binding the contract method 0xcade791f.
//
// Solidity: function keyToBestSell(bytes32 ) view returns(address quote, address base, uint256 quoteVolume, uint256 baseVolume, address dexAddress)
func (_Aggregator *AggregatorSession) KeyToBestSell(arg0 [32]byte) (struct {
	Quote       common.Address
	Base        common.Address
	QuoteVolume *big.Int
	BaseVolume  *big.Int
	DexAddress  common.Address
}, error) {
	return _Aggregator.Contract.KeyToBestSell(&_Aggregator.CallOpts, arg0)
}

// KeyToBestSell is a free data retrieval call binding the contract method 0xcade791f.
//
// Solidity: function keyToBestSell(bytes32 ) view returns(address quote, address base, uint256 quoteVolume, uint256 baseVolume, address dexAddress)
func (_Aggregator *AggregatorCallerSession) KeyToBestSell(arg0 [32]byte) (struct {
	Quote       common.Address
	Base        common.Address
	QuoteVolume *big.Int
	BaseVolume  *big.Int
	DexAddress  common.Address
}, error) {
	return _Aggregator.Contract.KeyToBestSell(&_Aggregator.CallOpts, arg0)
}

// Swap is a paid mutator transaction binding the contract method 0xdf791e50.
//
// Solidity: function swap(address token0, address token1, uint256 amt) returns()
func (_Aggregator *AggregatorTransactor) Swap(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "swap", token0, token1, amt)
}

// Swap is a paid mutator transaction binding the contract method 0xdf791e50.
//
// Solidity: function swap(address token0, address token1, uint256 amt) returns()
func (_Aggregator *AggregatorSession) Swap(token0 common.Address, token1 common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.Swap(&_Aggregator.TransactOpts, token0, token1, amt)
}

// Swap is a paid mutator transaction binding the contract method 0xdf791e50.
//
// Solidity: function swap(address token0, address token1, uint256 amt) returns()
func (_Aggregator *AggregatorTransactorSession) Swap(token0 common.Address, token1 common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.Swap(&_Aggregator.TransactOpts, token0, token1, amt)
}

// UpdateBestAll is a paid mutator transaction binding the contract method 0xb9e16722.
//
// Solidity: function updateBestAll(address base, address quote, uint256 baseVolume, uint256 quoteVolume, address tradeContract) returns()
func (_Aggregator *AggregatorTransactor) UpdateBestAll(opts *bind.TransactOpts, base common.Address, quote common.Address, baseVolume *big.Int, quoteVolume *big.Int, tradeContract common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "updateBestAll", base, quote, baseVolume, quoteVolume, tradeContract)
}

// UpdateBestAll is a paid mutator transaction binding the contract method 0xb9e16722.
//
// Solidity: function updateBestAll(address base, address quote, uint256 baseVolume, uint256 quoteVolume, address tradeContract) returns()
func (_Aggregator *AggregatorSession) UpdateBestAll(base common.Address, quote common.Address, baseVolume *big.Int, quoteVolume *big.Int, tradeContract common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateBestAll(&_Aggregator.TransactOpts, base, quote, baseVolume, quoteVolume, tradeContract)
}

// UpdateBestAll is a paid mutator transaction binding the contract method 0xb9e16722.
//
// Solidity: function updateBestAll(address base, address quote, uint256 baseVolume, uint256 quoteVolume, address tradeContract) returns()
func (_Aggregator *AggregatorTransactorSession) UpdateBestAll(base common.Address, quote common.Address, baseVolume *big.Int, quoteVolume *big.Int, tradeContract common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateBestAll(&_Aggregator.TransactOpts, base, quote, baseVolume, quoteVolume, tradeContract)
}

// UpdateBestBuy is a paid mutator transaction binding the contract method 0xc6389879.
//
// Solidity: function updateBestBuy(address base, address quote, uint256 baseVolume, uint256 quoteVolume, address tradeContract) returns()
func (_Aggregator *AggregatorTransactor) UpdateBestBuy(opts *bind.TransactOpts, base common.Address, quote common.Address, baseVolume *big.Int, quoteVolume *big.Int, tradeContract common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "updateBestBuy", base, quote, baseVolume, quoteVolume, tradeContract)
}

// UpdateBestBuy is a paid mutator transaction binding the contract method 0xc6389879.
//
// Solidity: function updateBestBuy(address base, address quote, uint256 baseVolume, uint256 quoteVolume, address tradeContract) returns()
func (_Aggregator *AggregatorSession) UpdateBestBuy(base common.Address, quote common.Address, baseVolume *big.Int, quoteVolume *big.Int, tradeContract common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateBestBuy(&_Aggregator.TransactOpts, base, quote, baseVolume, quoteVolume, tradeContract)
}

// UpdateBestBuy is a paid mutator transaction binding the contract method 0xc6389879.
//
// Solidity: function updateBestBuy(address base, address quote, uint256 baseVolume, uint256 quoteVolume, address tradeContract) returns()
func (_Aggregator *AggregatorTransactorSession) UpdateBestBuy(base common.Address, quote common.Address, baseVolume *big.Int, quoteVolume *big.Int, tradeContract common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateBestBuy(&_Aggregator.TransactOpts, base, quote, baseVolume, quoteVolume, tradeContract)
}

// UpdateBestSell is a paid mutator transaction binding the contract method 0x58618e7d.
//
// Solidity: function updateBestSell(address base, address quote, uint256 baseVolume, uint256 quoteVolume, address tradeContract) returns()
func (_Aggregator *AggregatorTransactor) UpdateBestSell(opts *bind.TransactOpts, base common.Address, quote common.Address, baseVolume *big.Int, quoteVolume *big.Int, tradeContract common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "updateBestSell", base, quote, baseVolume, quoteVolume, tradeContract)
}

// UpdateBestSell is a paid mutator transaction binding the contract method 0x58618e7d.
//
// Solidity: function updateBestSell(address base, address quote, uint256 baseVolume, uint256 quoteVolume, address tradeContract) returns()
func (_Aggregator *AggregatorSession) UpdateBestSell(base common.Address, quote common.Address, baseVolume *big.Int, quoteVolume *big.Int, tradeContract common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateBestSell(&_Aggregator.TransactOpts, base, quote, baseVolume, quoteVolume, tradeContract)
}

// UpdateBestSell is a paid mutator transaction binding the contract method 0x58618e7d.
//
// Solidity: function updateBestSell(address base, address quote, uint256 baseVolume, uint256 quoteVolume, address tradeContract) returns()
func (_Aggregator *AggregatorTransactorSession) UpdateBestSell(base common.Address, quote common.Address, baseVolume *big.Int, quoteVolume *big.Int, tradeContract common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateBestSell(&_Aggregator.TransactOpts, base, quote, baseVolume, quoteVolume, tradeContract)
}
