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

// RelayHandlerABI is the input ABI used to generate the binding from.
const RelayHandlerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"sentBytes\",\"type\":\"uint128\"},{\"name\":\"sentBytesSignature\",\"type\":\"bytes\"},{\"name\":\"senderPublicKey\",\"type\":\"bytes\"},{\"name\":\"ids\",\"type\":\"uint128[][3]\"},{\"name\":\"keysN\",\"type\":\"bytes[][2]\"},{\"name\":\"keysE\",\"type\":\"uint256[][2]\"},{\"name\":\"signatures\",\"type\":\"bytes[][2]\"}],\"name\":\"submitRelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"RelayHonored\",\"type\":\"event\"}]"

// RelayHandlerBin is the compiled bytecode used for deploying new contracts.
const RelayHandlerBin = `0x608060405234801561001057600080fd5b50610c4b806100206000396000f3006080604052600436106100405763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166332760d5f8114610045575b600080fd5b34801561005157600080fd5b506100656100603660046109ca565b610067565b005b61006f6103f8565b610077610445565b5050604080516101208101825285518152602080870151818301528683015182840152855160608084019190915285516080808501919091528783015160a08501528683015160c0850152855160e085015285830151610100850152845190810185526fffffffffffffffffffffffffffffffff8c1681529182018a905292810188905291820181905290600061010d886102c2565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260208181526040822080546001808201808455928552938390208751600c9092020180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff90921691909117815586830151805195965091948794919361019893928501920190610486565b50604082015180516101b4916002840191602090910190610486565b5060608201518051805160038401916101d291839160200190610504565b5060208281015180516101eb9260018501920190610504565b5060408201518051610207916002840191602090910190610504565b50606082015180516102239160038401916020909101906105d3565b506080820151805161023f91600484019160209091019061062c565b5060a0820151805161025b9160058401916020909101906105d3565b5060c0820151805161027791600684019160209091019061062c565b5060e082015180516102939160078401916020909101906105d3565b5061010082015180516102b09160088401916020909101906105d3565b50505050505050505050505050505050565b60008060008060006014865110151515610311576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030890610b4c565b60405180910390fd5b6000935060018651039250600091505b60148210156103ee57858381518110151561033857fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169050806c0100000000000000000000000090048401935060088473ffffffffffffffffffffffffffffffffffffffff169060020a02935082600190039250816001019150610321565b5091949350505050565b610120604051908101604052806060815260200160608152602001606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b6101806040519081016040528060006fffffffffffffffffffffffffffffffff16815260200160608152602001606081526020016104816103f8565b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106104c757805160ff19168380011785556104f4565b828001600101855582156104f4579182015b828111156104f45782518255916020019190600101906104d9565b50610500929150610666565b5090565b828054828255906000526020600020906001016002900481019282156105c75791602002820160005b8382111561058957835183826101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff1602179055509260200192601001602081600f0104928301926001030261052d565b80156105c55782816101000a8154906fffffffffffffffffffffffffffffffff0219169055601001602081600f01049283019260010302610589565b505b50610500929150610683565b828054828255906000526020600020908101928215610620579160200282015b828111156106205782518051610610918491602090910190610486565b50916020019190600101906105f3565b506105009291506106b0565b8280548282559060005260206000209081019282156104f457916020028201828111156104f45782518255916020019190600101906104d9565b61068091905b80821115610500576000815560010161066c565b90565b61068091905b808211156105005780546fffffffffffffffffffffffffffffffff19168155600101610689565b61068091905b808211156105005760006106ca82826106d3565b506001016106b6565b50805460018160011615610100020316600290046000825580601f106106f95750610717565b601f0160209004906000526020600020908101906107179190610666565b50565b6000601f8201831361072b57600080fd5b600261073e61073982610b89565b610b62565b9150818360005b83811015610771578135860161075b888261081f565b8452506020928301929190910190600101610745565b5050505092915050565b6000601f8201831361078c57600080fd5b600361079a61073982610b89565b9150818360005b8381101561077157813586016107b7888261087c565b84525060209283019291909101906001016107a1565b6000601f820183136107de57600080fd5b60026107ec61073982610b89565b9150818360005b83811015610771578135860161080988826108ec565b84525060209283019291909101906001016107f3565b6000601f8201831361083057600080fd5b813561083e61073982610ba7565b81815260209384019390925082018360005b838110156107715781358601610866888261095c565b8452506020928301929190910190600101610850565b6000601f8201831361088d57600080fd5b813561089b61073982610ba7565b915081818352602084019350602081019050838560208402820111156108c057600080fd5b60005b8381101561077157816108d688826109ab565b84525060209283019291909101906001016108c3565b6000601f820183136108fd57600080fd5b813561090b61073982610ba7565b9150818183526020840193506020810190508385602084028201111561093057600080fd5b60005b83811015610771578161094688826109be565b8452506020928301929190910190600101610933565b6000601f8201831361096d57600080fd5b813561097b61073982610bc8565b9150808252602083016020830185838301111561099757600080fd5b6109a2838284610c05565b50505092915050565b60006109b78235610bf0565b9392505050565b60006109b78235610680565b600080600080600080600060e0888a0312156109e557600080fd5b60006109f18a8a6109ab565b975050602088013567ffffffffffffffff811115610a0e57600080fd5b610a1a8a828b0161095c565b965050604088013567ffffffffffffffff811115610a3757600080fd5b610a438a828b0161095c565b955050606088013567ffffffffffffffff811115610a6057600080fd5b610a6c8a828b0161077b565b945050608088013567ffffffffffffffff811115610a8957600080fd5b610a958a828b0161071a565b93505060a088013567ffffffffffffffff811115610ab257600080fd5b610abe8a828b016107cd565b92505060c088013567ffffffffffffffff811115610adb57600080fd5b610ae78a828b0161071a565b91505092959891949750929550565b602481527f546865206b6579206d757374206265206f66206174206c65617374203230206260208201527f7974657300000000000000000000000000000000000000000000000000000000604082015260600190565b60208082528101610b5c81610af6565b92915050565b60405181810167ffffffffffffffff81118282101715610b8157600080fd5b604052919050565b600067ffffffffffffffff821115610ba057600080fd5b5060200290565b600067ffffffffffffffff821115610bbe57600080fd5b5060209081020190565b600067ffffffffffffffff821115610bdf57600080fd5b506020601f91909101601f19160190565b6fffffffffffffffffffffffffffffffff1690565b828183375060009101525600a265627a7a72305820ce59d99faf25d4dad3dc8ccd8389e34930e9f2523a1ea38f1435509a89cb53f26c6578706572696d656e74616cf50037`

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

// SubmitRelay is a paid mutator transaction binding the contract method 0x32760d5f.
//
// Solidity: function submitRelay(sentBytes uint128, sentBytesSignature bytes, senderPublicKey bytes, ids uint128[][3], keysN bytes[][2], keysE uint256[][2], signatures bytes[][2]) returns()
func (_RelayHandler *RelayHandlerTransactor) SubmitRelay(opts *bind.TransactOpts, sentBytes *big.Int, sentBytesSignature []byte, senderPublicKey []byte, ids [3][]*big.Int, keysN [2][][]byte, keysE [2][]*big.Int, signatures [2][][]byte) (*types.Transaction, error) {
	return _RelayHandler.contract.Transact(opts, "submitRelay", sentBytes, sentBytesSignature, senderPublicKey, ids, keysN, keysE, signatures)
}

// SubmitRelay is a paid mutator transaction binding the contract method 0x32760d5f.
//
// Solidity: function submitRelay(sentBytes uint128, sentBytesSignature bytes, senderPublicKey bytes, ids uint128[][3], keysN bytes[][2], keysE uint256[][2], signatures bytes[][2]) returns()
func (_RelayHandler *RelayHandlerSession) SubmitRelay(sentBytes *big.Int, sentBytesSignature []byte, senderPublicKey []byte, ids [3][]*big.Int, keysN [2][][]byte, keysE [2][]*big.Int, signatures [2][][]byte) (*types.Transaction, error) {
	return _RelayHandler.Contract.SubmitRelay(&_RelayHandler.TransactOpts, sentBytes, sentBytesSignature, senderPublicKey, ids, keysN, keysE, signatures)
}

// SubmitRelay is a paid mutator transaction binding the contract method 0x32760d5f.
//
// Solidity: function submitRelay(sentBytes uint128, sentBytesSignature bytes, senderPublicKey bytes, ids uint128[][3], keysN bytes[][2], keysE uint256[][2], signatures bytes[][2]) returns()
func (_RelayHandler *RelayHandlerTransactorSession) SubmitRelay(sentBytes *big.Int, sentBytesSignature []byte, senderPublicKey []byte, ids [3][]*big.Int, keysN [2][][]byte, keysE [2][]*big.Int, signatures [2][][]byte) (*types.Transaction, error) {
	return _RelayHandler.Contract.SubmitRelay(&_RelayHandler.TransactOpts, sentBytes, sentBytesSignature, senderPublicKey, ids, keysN, keysE, signatures)
}

// RelayHandlerRelayHonoredIterator is returned from FilterRelayHonored and is used to iterate over the raw logs and unpacked data for RelayHonored events raised by the RelayHandler contract.
type RelayHandlerRelayHonoredIterator struct {
	Event *RelayHandlerRelayHonored // Event containing the contract specifics and raw log

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
func (it *RelayHandlerRelayHonoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHandlerRelayHonored)
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
		it.Event = new(RelayHandlerRelayHonored)
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
func (it *RelayHandlerRelayHonoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHandlerRelayHonoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHandlerRelayHonored represents a RelayHonored event raised by the RelayHandler contract.
type RelayHandlerRelayHonored struct {
	common.Address
	*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRelayHonored is a free log retrieval operation binding the contract event 0x5e2962388ace4e86a9836399fe6b8ca76e519b2cc912a26b8a8f581010bd3c79.
//
// Solidity: e RelayHonored( address,  uint256)
func (_RelayHandler *RelayHandlerFilterer) FilterRelayHonored(opts *bind.FilterOpts) (*RelayHandlerRelayHonoredIterator, error) {

	logs, sub, err := _RelayHandler.contract.FilterLogs(opts, "RelayHonored")
	if err != nil {
		return nil, err
	}
	return &RelayHandlerRelayHonoredIterator{contract: _RelayHandler.contract, event: "RelayHonored", logs: logs, sub: sub}, nil
}

// WatchRelayHonored is a free log subscription operation binding the contract event 0x5e2962388ace4e86a9836399fe6b8ca76e519b2cc912a26b8a8f581010bd3c79.
//
// Solidity: e RelayHonored( address,  uint256)
func (_RelayHandler *RelayHandlerFilterer) WatchRelayHonored(opts *bind.WatchOpts, sink chan<- *RelayHandlerRelayHonored) (event.Subscription, error) {

	logs, sub, err := _RelayHandler.contract.WatchLogs(opts, "RelayHonored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHandlerRelayHonored)
				if err := _RelayHandler.contract.UnpackLog(event, "RelayHonored", log); err != nil {
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
