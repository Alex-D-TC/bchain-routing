// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RelayHandlerABI is the input ABI used to generate the binding from.
const RelayHandlerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// RelayHandlerBin is the compiled bytecode used for deploying new contracts.
const RelayHandlerBin = `0x6080604052348015600f57600080fd5b50604380601d6000396000f3fe6080604052600080fdfea265627a7a723058207d9b706421931a10cc894f386b06a8e1cc70847a27b67b4610b820e31b6d941d6c6578706572696d656e74616cf50037`

// DeployRelayHandler deploys a new Ethereum contract, binding an instance of RelayHandler to it.
func DeployRelayHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RelayHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RelayHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelayHandler{RelayHandlerCaller: RelayHandlerCaller{contract: contract}, RelayHandlerTransactor: RelayHandlerTransactor{contract: contract}, RelayHandlerFilterer: RelayHandlerFilterer{contract: contract}}, nil
}

// RelayHandler is an auto generated Go binding around an Ethereum contract.
type RelayHandler struct {
	RelayHandlerCaller     // Read-only binding to the contract
	RelayHandlerTransactor // Write-only binding to the contract
	RelayHandlerFilterer   // Log filterer for contract events
}

// RelayHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayHandlerSession struct {
	Contract     *RelayHandler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayHandlerCallerSession struct {
	Contract *RelayHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RelayHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayHandlerTransactorSession struct {
	Contract     *RelayHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RelayHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayHandlerRaw struct {
	Contract *RelayHandler // Generic contract binding to access the raw methods on
}

// RelayHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayHandlerCallerRaw struct {
	Contract *RelayHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// RelayHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayHandlerTransactorRaw struct {
	Contract *RelayHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayHandler creates a new instance of RelayHandler, bound to a specific deployed contract.
func NewRelayHandler(address common.Address, backend bind.ContractBackend) (*RelayHandler, error) {
	contract, err := bindRelayHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayHandler{RelayHandlerCaller: RelayHandlerCaller{contract: contract}, RelayHandlerTransactor: RelayHandlerTransactor{contract: contract}, RelayHandlerFilterer: RelayHandlerFilterer{contract: contract}}, nil
}

// NewRelayHandlerCaller creates a new read-only instance of RelayHandler, bound to a specific deployed contract.
func NewRelayHandlerCaller(address common.Address, caller bind.ContractCaller) (*RelayHandlerCaller, error) {
	contract, err := bindRelayHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHandlerCaller{contract: contract}, nil
}

// NewRelayHandlerTransactor creates a new write-only instance of RelayHandler, bound to a specific deployed contract.
func NewRelayHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayHandlerTransactor, error) {
	contract, err := bindRelayHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHandlerTransactor{contract: contract}, nil
}

// NewRelayHandlerFilterer creates a new log filterer instance of RelayHandler, bound to a specific deployed contract.
func NewRelayHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayHandlerFilterer, error) {
	contract, err := bindRelayHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayHandlerFilterer{contract: contract}, nil
}

// bindRelayHandler binds a generic wrapper to an already deployed contract.
func bindRelayHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHandler *RelayHandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RelayHandler.Contract.RelayHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHandler *RelayHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHandler.Contract.RelayHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHandler *RelayHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHandler.Contract.RelayHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHandler *RelayHandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RelayHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHandler *RelayHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHandler *RelayHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHandler.Contract.contract.Transact(opts, method, params...)
}
