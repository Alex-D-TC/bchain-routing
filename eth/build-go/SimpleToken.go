// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// SimpleTokenABI is the input ABI used to generate the binding from.
const SimpleTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"supply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"sendTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// SimpleTokenBin is the compiled bytecode used for deploying new contracts.
const SimpleTokenBin = `0x608060405234801561001057600080fd5b5060405161070038038061070083398101604090815281516020808401519284015191840180519094939093019261004e916001919086019061007f565b50815161006290600290602085019061007f565b50600355505060008054600160a060020a0319163317905561011a565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c057805160ff19168380011785556100ed565b828001600101855582156100ed579182015b828111156100ed5782518255916020019190600101906100d2565b506100f99291506100fd565b5090565b61011791905b808211156100f95760008155600101610103565b90565b6105d7806101296000396000f3006080604052600436106100a35763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663047fc9aa81146100a857806306fdde03146100cf57806318160ddd1461015957806340c10f191461016e57806370a08231146101945780638da5cb5b146101b557806395d89b41146101e65780639e1a00aa146101fb578063a6f9dae11461021f578063d5abeb0114610240575b600080fd5b3480156100b457600080fd5b506100bd610255565b60408051918252519081900360200190f35b3480156100db57600080fd5b506100e461025b565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561011e578181015183820152602001610106565b50505050905090810190601f16801561014b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016557600080fd5b506100bd6102e8565b34801561017a57600080fd5b50610192600160a060020a03600435166024356102ee565b005b3480156101a057600080fd5b506100bd600160a060020a03600435166103cd565b3480156101c157600080fd5b506101ca6103e8565b60408051600160a060020a039092168252519081900360200190f35b3480156101f257600080fd5b506100e46103f7565b34801561020757600080fd5b50610192600160a060020a036004351660243561044f565b34801561022b57600080fd5b50610192600160a060020a036004351661055c565b34801561024c57600080fd5b506100bd6105a5565b60045481565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102e05780601f106102b5576101008083540402835291602001916102e0565b820191906000526020600020905b8154815290600101906020018083116102c357829003601f168201915b505050505081565b60045490565b6000543390600160a060020a0316811461030757600080fd5b600354600454830111156103a257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f43616e6e6f74206d696e74206d6f7265207468616e20746865206d6178696d7560448201527f6d20616c6c6f7765640000000000000000000000000000000000000000000000606482015290519081900360840190fd5b50600160a060020a039091166000908152600560205260409020805482019055600480549091019055565b600160a060020a031660009081526005602052604090205490565b600054600160a060020a031681565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156102e05780601f106102b5576101008083540402835291602001916102e0565b336000908152600560205260409020548111156104f357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f43616e6e6f742073656e64206d6f726520636f696e73207468616e2049206f7760448201527f6e00000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b33600081815260056020908152604080832080548690039055600160a060020a03861680845292819020805486019055805185815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35050565b6000543390600160a060020a0316811461057557600080fd5b506000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354815600a165627a7a7230582004b2ace3f93a553f7cc7be10b4c6a73ff91bbeea3f9b764721d12ccd49337e070029`

// DeploySimpleToken deploys a new Ethereum contract, binding an instance of SimpleToken to it.
func DeploySimpleToken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _maxSupply *big.Int) (common.Address, *types.Transaction, *SimpleToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleTokenBin), backend, _name, _symbol, _maxSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleToken{SimpleTokenCaller: SimpleTokenCaller{contract: contract}, SimpleTokenTransactor: SimpleTokenTransactor{contract: contract}, SimpleTokenFilterer: SimpleTokenFilterer{contract: contract}}, nil
}

// SimpleToken is an auto generated Go binding around an Ethereum contract.
type SimpleToken struct {
	SimpleTokenCaller     // Read-only binding to the contract
	SimpleTokenTransactor // Write-only binding to the contract
	SimpleTokenFilterer   // Log filterer for contract events
}

// SimpleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleTokenSession struct {
	Contract     *SimpleToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleTokenCallerSession struct {
	Contract *SimpleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SimpleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleTokenTransactorSession struct {
	Contract     *SimpleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SimpleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleTokenRaw struct {
	Contract *SimpleToken // Generic contract binding to access the raw methods on
}

// SimpleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleTokenCallerRaw struct {
	Contract *SimpleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleTokenTransactorRaw struct {
	Contract *SimpleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleToken creates a new instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleToken(address common.Address, backend bind.ContractBackend) (*SimpleToken, error) {
	contract, err := bindSimpleToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleToken{SimpleTokenCaller: SimpleTokenCaller{contract: contract}, SimpleTokenTransactor: SimpleTokenTransactor{contract: contract}, SimpleTokenFilterer: SimpleTokenFilterer{contract: contract}}, nil
}

// NewSimpleTokenCaller creates a new read-only instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleTokenCaller(address common.Address, caller bind.ContractCaller) (*SimpleTokenCaller, error) {
	contract, err := bindSimpleToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenCaller{contract: contract}, nil
}

// NewSimpleTokenTransactor creates a new write-only instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleTokenTransactor, error) {
	contract, err := bindSimpleToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenTransactor{contract: contract}, nil
}

// NewSimpleTokenFilterer creates a new log filterer instance of SimpleToken, bound to a specific deployed contract.
func NewSimpleTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleTokenFilterer, error) {
	contract, err := bindSimpleToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenFilterer{contract: contract}, nil
}

// bindSimpleToken binds a generic wrapper to an already deployed contract.
func bindSimpleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleToken *SimpleTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleToken.Contract.SimpleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleToken *SimpleTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleToken.Contract.SimpleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleToken *SimpleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleToken.Contract.SimpleTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleToken *SimpleTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleToken *SimpleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleToken *SimpleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.BalanceOf(&_SimpleToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _SimpleToken.Contract.BalanceOf(&_SimpleToken.CallOpts, who)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "maxSupply")
	return *ret0, err
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) MaxSupply() (*big.Int, error) {
	return _SimpleToken.Contract.MaxSupply(&_SimpleToken.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) MaxSupply() (*big.Int, error) {
	return _SimpleToken.Contract.MaxSupply(&_SimpleToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SimpleToken *SimpleTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SimpleToken *SimpleTokenSession) Name() (string, error) {
	return _SimpleToken.Contract.Name(&_SimpleToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SimpleToken *SimpleTokenCallerSession) Name() (string, error) {
	return _SimpleToken.Contract.Name(&_SimpleToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SimpleToken *SimpleTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SimpleToken *SimpleTokenSession) Owner() (common.Address, error) {
	return _SimpleToken.Contract.Owner(&_SimpleToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SimpleToken *SimpleTokenCallerSession) Owner() (common.Address, error) {
	return _SimpleToken.Contract.Owner(&_SimpleToken.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "supply")
	return *ret0, err
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) Supply() (*big.Int, error) {
	return _SimpleToken.Contract.Supply(&_SimpleToken.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) Supply() (*big.Int, error) {
	return _SimpleToken.Contract.Supply(&_SimpleToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SimpleToken *SimpleTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SimpleToken *SimpleTokenSession) Symbol() (string, error) {
	return _SimpleToken.Contract.Symbol(&_SimpleToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SimpleToken *SimpleTokenCallerSession) Symbol() (string, error) {
	return _SimpleToken.Contract.Symbol(&_SimpleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenSession) TotalSupply() (*big.Int, error) {
	return _SimpleToken.Contract.TotalSupply(&_SimpleToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SimpleToken *SimpleTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SimpleToken.Contract.TotalSupply(&_SimpleToken.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_SimpleToken *SimpleTokenTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_SimpleToken *SimpleTokenSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleToken.Contract.ChangeOwner(&_SimpleToken.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_SimpleToken *SimpleTokenTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleToken.Contract.ChangeOwner(&_SimpleToken.TransactOpts, newOwner)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, value uint256) returns()
func (_SimpleToken *SimpleTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, value uint256) returns()
func (_SimpleToken *SimpleTokenSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.Mint(&_SimpleToken.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, value uint256) returns()
func (_SimpleToken *SimpleTokenTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.Mint(&_SimpleToken.TransactOpts, to, value)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(to address, value uint256) returns()
func (_SimpleToken *SimpleTokenTransactor) SendTo(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.contract.Transact(opts, "sendTo", to, value)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(to address, value uint256) returns()
func (_SimpleToken *SimpleTokenSession) SendTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.SendTo(&_SimpleToken.TransactOpts, to, value)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(to address, value uint256) returns()
func (_SimpleToken *SimpleTokenTransactorSession) SendTo(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SimpleToken.Contract.SendTo(&_SimpleToken.TransactOpts, to, value)
}

// SimpleTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SimpleToken contract.
type SimpleTokenTransferIterator struct {
	Event *SimpleTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SimpleTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleTokenTransfer)
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
		it.Event = new(SimpleTokenTransfer)
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
func (it *SimpleTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleTokenTransfer represents a Transfer event raised by the SimpleToken contract.
type SimpleTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_SimpleToken *SimpleTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimpleTokenTransferIterator{contract: _SimpleToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_SimpleToken *SimpleTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SimpleTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleTokenTransfer)
				if err := _SimpleToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
