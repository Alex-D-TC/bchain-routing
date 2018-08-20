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
const RelayHandlerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_sentBytes\",\"type\":\"uint128\"},{\"name\":\"_sentBytesHash\",\"type\":\"uint256\"},{\"name\":\"_sentBytesSignature\",\"type\":\"bytes\"},{\"name\":\"_senderPublicKey\",\"type\":\"bytes\"},{\"name\":\"_ids\",\"type\":\"uint128[][3]\"},{\"name\":\"_keysN\",\"type\":\"bytes[][2]\"},{\"name\":\"_keysE\",\"type\":\"uint256[][2]\"},{\"name\":\"_signatures\",\"type\":\"bytes[][2]\"},{\"name\":\"_porRaw\",\"type\":\"bytes[]\"}],\"name\":\"submitRelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getRelay\",\"outputs\":[{\"name\":\"sentBytes\",\"type\":\"uint128\"},{\"name\":\"sentBytesSignature\",\"type\":\"bytes\"},{\"name\":\"senderPublicKey\",\"type\":\"bytes\"},{\"name\":\"ids\",\"type\":\"uint128[][3]\"},{\"name\":\"keysN\",\"type\":\"bytes[][2]\"},{\"name\":\"keysE\",\"type\":\"uint256[][2]\"},{\"name\":\"signatures\",\"type\":\"bytes[][2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"RelayHonored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"RelayPaymentReqeusted\",\"type\":\"event\"}]"

// RelayHandlerBin is the compiled bytecode used for deploying new contracts.
const RelayHandlerBin = `0x608060405234801561001057600080fd5b506119d9806100206000396000f30060806040526004361061004b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663112ddfb78114610050578063aea7a62d14610086575b600080fd5b34801561005c57600080fd5b5061007061006b366004611365565b6100b9565b60405161007d9190611882565b60405180910390f35b34801561009257600080fd5b506100a66100a136600461132b565b610351565b60405161007d97969594939291906117f1565b60006100c3610cc5565b6100cb610d19565b60408051610140810182528951815289820151602080830191909152808b01519282019290925288516060820152875160808201528882015160a08201529087015160c0820152855160e0820152610100810185905260009061012081018760016020020151815250925060a0604051908101604052808e6001608060020a031681526020018d81526020018c81526020018b81526020018481525091506101768260600151610b98565b73ffffffffffffffffffffffffffffffffffffffff811660009081526020818152604080832080546001808201808455928652948490208851600e9092020180546fffffffffffffffffffffffffffffffff19166001608060020a0390921691909117815587840151948101949094559086015180519495509093869392610205926002850192910190610d58565b5060608201518051610221916003840191602090910190610d58565b50608082015180518051600484019161023f91839160200190610dd6565b5060208281015180516102589260018501920190610dd6565b5060408201518051610274916002840191602090910190610dd6565b5060608201518051610290916003840191602090910190610e8a565b50608082015180516102ac916004840191602090910190610ee3565b5060a082015180516102c8916005840191602090910190610e8a565b5060c082015180516102e4916006840191602090910190610ee3565b5060e08201518051610300916007840191602090910190610e8a565b50610100820151805161031d916008840191602090910190610e8a565b50610120820151805161033a916009840191602090910190610e8a565b505050505093505050509998505050505050505050565b600060608061035e610f1d565b610366610f45565b61036e610f45565b610376610f45565b73ffffffffffffffffffffffffffffffffffffffff891660009081526020819052604081205489106103dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d4906117e1565b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a16600090815260208190526040902080548a90811061040e57fe5b6000918252602091829020600e9091020180546002808301805460408051601f600019600185161561010002019093169490940491820187900487028401870190528083526001608060020a039093169c50929450928301828280156104b55780601f1061048a576101008083540402835291602001916104b5565b820191906000526020600020905b81548152906001019060200180831161049857829003601f168201915b5050505060038301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949b509192508301828280156105455780601f1061051a57610100808354040283529160200191610545565b820191906000526020600020905b81548152906001019060200180831161052857829003601f168201915b50505060048401805460408051602080840282018101909252828152959b5091935091508301828280156105ca57602002820191906000526020600020906000905b82829054906101000a90046001608060020a03166001608060020a031681526020019060100190602082600f010492830192600103820291508084116105875790505b50899350600092506105da915050565b60200201819052508060040160020180548060200260200160405190810160405280929190818152602001828054801561066557602002820191906000526020600020906000905b82829054906101000a90046001608060020a03166001608060020a031681526020019060100190602082600f010492830192600103820291508084116106225790505b5089935060019250610675915050565b60200201819052508060040160010180548060200260200160405190810160405280929190818152602001828054801561070057602002820191906000526020600020906000905b82829054906101000a90046001608060020a03166001608060020a031681526020019060100190602082600f010492830192600103820291508084116106bd5790505b5089935060029250610710915050565b602002018190525080600401600301805480602002602001604051908101604052809291908181526020016000905b828210156107ea5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156107d65780601f106107ab576101008083540402835291602001916107d6565b820191906000526020600020905b8154815290600101906020018083116107b957829003601f168201915b50505050508152602001906001019061073f565b50879250600091506107f99050565b60200201819052508060040160040180548060200260200160405190810160405280929190818152602001828054801561085257602002820191906000526020600020905b81548152602001906001019080831161083e575b5087935060009250610862915050565b602002018190525080600401600501805480602002602001604051908101604052809291908181526020016000905b8282101561093c5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156109285780601f106108fd57610100808354040283529160200191610928565b820191906000526020600020905b81548152906001019060200180831161090b57829003601f168201915b505050505081526020019060010190610891565b508792506001915061094b9050565b6020020181905250806004016006018054806020026020016040519081016040528092919081815260200182805480156109a457602002820191906000526020600020905b815481526020019060010190808311610990575b50879350600192506109b4915050565b602002018190525080600401600701805480602002602001604051908101604052809291908181526020016000905b82821015610a8e5760008481526020908190208301805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015610a7a5780601f10610a4f57610100808354040283529160200191610a7a565b820191906000526020600020905b815481529060010190602001808311610a5d57829003601f168201915b5050505050815260200190600101906109e3565b5085925060009150610a9d9050565b602002018190525080600401600901805480602002602001604051908101604052809291908181526020016000905b82821015610b775760008481526020908190208301805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015610b635780601f10610b3857610100808354040283529160200191610b63565b820191906000526020600020905b815481529060010190602001808311610b4657829003601f168201915b505050505081526020019060010190610acc565b5085925060019150610b869050565b60200201525092959891949750929550565b60008060008060006014865110151515610bde576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d4906117cb565b6000935060018651039250600091505b6014821015610cbb578583815181101515610c0557fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169050806c0100000000000000000000000090048401935060088473ffffffffffffffffffffffffffffffffffffffff169060020a02935082600190039250816001019150610bee565b5091949350505050565b61014060405190810160405280606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b6101c06040519081016040528060006001608060020a03168152602001600081526020016060815260200160608152602001610d53610cc5565b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610d9957805160ff1916838001178555610dc6565b82800160010185558215610dc6579182015b82811115610dc6578251825591602001919060010190610dab565b50610dd2929150610f5e565b5090565b82805482825590600052602060002090600101600290048101928215610e7e5791602002820160005b83821115610e4957835183826101000a8154816001608060020a0302191690836001608060020a031602179055509260200192601001602081600f01049283019260010302610dff565b8015610e7c5782816101000a8154906001608060020a030219169055601001602081600f01049283019260010302610e49565b505b50610dd2929150610f7b565b828054828255906000526020600020908101928215610ed7579160200282015b82811115610ed75782518051610ec7918491602090910190610d58565b5091602001919060010190610eaa565b50610dd2929150610fa8565b828054828255906000526020600020908101928215610dc65791602002820182811115610dc6578251825591602001919060010190610dab565b6060604051908101604052806003905b6060815260200190600190039081610f2d5790505090565b6040805180820190915260608152600160208201610f2d565b610f7891905b80821115610dd25760008155600101610f64565b90565b610f7891905b80821115610dd25780546fffffffffffffffffffffffffffffffff19168155600101610f81565b610f7891905b80821115610dd2576000610fc28282610fcb565b50600101610fae565b50805460018160011615610100020316600290046000825580601f10610ff1575061100f565b601f01602090049060005260206000209081019061100f9190610f5e565b50565b600061101e8235611940565b9392505050565b6000601f8201831361103657600080fd5b6002611049611044826118b7565b611890565b9150818360005b8381101561107c5781358601611066888261112a565b8452506020928301929190910190600101611050565b5050505092915050565b6000601f8201831361109757600080fd5b60036110a5611044826118b7565b9150818360005b8381101561107c57813586016110c288826111e4565b84525060209283019291909101906001016110ac565b6000601f820183136110e957600080fd5b60026110f7611044826118b7565b9150818360005b8381101561107c57813586016111148882611254565b84525060209283019291909101906001016110fe565b6000601f8201831361113b57600080fd5b8135611149611044826118d5565b81815260209384019390925082018360005b8381101561107c578135860161117188826112c4565b845250602092830192919091019060010161115b565b6000601f8201831361119857600080fd5b81356111a6611044826118d5565b81815260209384019390925082018360005b8381101561107c57813586016111ce88826112c4565b84525060209283019291909101906001016111b8565b6000601f820183136111f557600080fd5b8135611203611044826118d5565b9150818183526020840193506020810190508385602084028201111561122857600080fd5b60005b8381101561107c578161123e8882611313565b845250602092830192919091019060010161122b565b6000601f8201831361126557600080fd5b8135611273611044826118d5565b9150818183526020840193506020810190508385602084028201111561129857600080fd5b60005b8381101561107c57816112ae888261131f565b845250602092830192919091019060010161129b565b6000601f820183136112d557600080fd5b81356112e3611044826118f6565b915080825260208301602083018583830111156112ff57600080fd5b61130a838284611959565b50505092915050565b600061101e8235611934565b600061101e8235610f78565b6000806040838503121561133e57600080fd5b600061134a8585611012565b925050602061135b8582860161131f565b9150509250929050565b60008060008060008060008060006101208a8c03121561138457600080fd5b60006113908c8c611313565b99505060206113a18c828d0161131f565b98505060408a013567ffffffffffffffff8111156113be57600080fd5b6113ca8c828d016112c4565b97505060608a013567ffffffffffffffff8111156113e757600080fd5b6113f38c828d016112c4565b96505060808a013567ffffffffffffffff81111561141057600080fd5b61141c8c828d01611086565b95505060a08a013567ffffffffffffffff81111561143957600080fd5b6114458c828d01611025565b94505060c08a013567ffffffffffffffff81111561146257600080fd5b61146e8c828d016110d8565b93505060e08a013567ffffffffffffffff81111561148b57600080fd5b6114978c828d01611025565b9250506101008a013567ffffffffffffffff8111156114b557600080fd5b6114c18c828d01611187565b9150509295985092959850929598565b60006114dc82611924565b836020820285016114ec85610f78565b60005b848110156115235783830388526115078383516115d3565b92506115128261191e565b6020989098019791506001016114ef565b50909695505050505050565b600061153a8261192a565b8360208202850161154a85610f78565b60005b8481101561152357838303885261156583835161162e565b92506115708261191e565b60209890980197915060010161154d565b600061158c82611924565b8360208202850161159c85610f78565b60005b848110156115235783830388526115b7838351611685565b92506115c28261191e565b60209890980197915060010161159f565b60006115de82611930565b808452602084019350836020820285016115f78561191e565b60005b848110156115235783830388526116128383516116d2565b925061161d8261191e565b6020989098019791506001016115fa565b600061163982611930565b80845260208401935061164b8361191e565b60005b8281101561167b576116618683516117b3565b61166a8261191e565b60209690960195915060010161164e565b5093949350505050565b600061169082611930565b8084526020840193506116a28361191e565b60005b8281101561167b576116b88683516117c2565b6116c18261191e565b6020969096019591506001016116a5565b60006116dd82611930565b8084526116f1816020860160208601611965565b6116fa81611995565b9093016020019392505050565b602481527f546865206b6579206d757374206265206f66206174206c65617374203230206260208201527f7974657300000000000000000000000000000000000000000000000000000000604082015260600190565b602681527f52656c617920776974682074686520676976656e20696420646f6573206e6f7460208201527f2065786973740000000000000000000000000000000000000000000000000000604082015260600190565b6117bc81611934565b82525050565b6117bc81610f78565b602080825281016117db81611707565b92915050565b602080825281016117db8161175d565b60e081016117ff828a6117b3565b818103602083015261181181896116d2565b9050818103604083015261182581886116d2565b90508181036060830152611839818761152f565b9050818103608083015261184d81866114d1565b905081810360a08301526118618185611581565b905081810360c083015261187581846114d1565b9998505050505050505050565b602081016117db82846117c2565b60405181810167ffffffffffffffff811182821017156118af57600080fd5b604052919050565b600067ffffffffffffffff8211156118ce57600080fd5b5060200290565b600067ffffffffffffffff8211156118ec57600080fd5b5060209081020190565b600067ffffffffffffffff82111561190d57600080fd5b506020601f91909101601f19160190565b60200190565b50600290565b50600390565b5190565b6001608060020a031690565b73ffffffffffffffffffffffffffffffffffffffff1690565b82818337506000910152565b60005b83811015611980578181015183820152602001611968565b8381111561198f576000848401525b50505050565b601f01601f1916905600a265627a7a7230582083e6a375f2285b2392796622c4f6f4b4167fd0126e10e4895044c6933ba44aeb6c6578706572696d656e74616cf50037`

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

// GetRelay is a free data retrieval call binding the contract method 0xaea7a62d.
//
// Solidity: function getRelay(_addr address, _id uint256) constant returns(sentBytes uint128, sentBytesSignature bytes, senderPublicKey bytes, ids uint128[][3], keysN bytes[][2], keysE uint256[][2], signatures bytes[][2])
func (_RelayHandler *RelayHandlerCaller) GetRelay(opts *bind.CallOpts, _addr common.Address, _id *big.Int) (struct {
	SentBytes          *big.Int
	SentBytesSignature []byte
	SenderPublicKey    []byte
	Ids                [3][]*big.Int
	KeysN              [2][][]byte
	KeysE              [2][]*big.Int
	Signatures         [2][][]byte
}, error) {
	ret := new(struct {
		SentBytes          *big.Int
		SentBytesSignature []byte
		SenderPublicKey    []byte
		Ids                [3][]*big.Int
		KeysN              [2][][]byte
		KeysE              [2][]*big.Int
		Signatures         [2][][]byte
	})
	out := ret
	err := _RelayHandler.contract.Call(opts, out, "getRelay", _addr, _id)
	return *ret, err
}

// GetRelay is a free data retrieval call binding the contract method 0xaea7a62d.
//
// Solidity: function getRelay(_addr address, _id uint256) constant returns(sentBytes uint128, sentBytesSignature bytes, senderPublicKey bytes, ids uint128[][3], keysN bytes[][2], keysE uint256[][2], signatures bytes[][2])
func (_RelayHandler *RelayHandlerSession) GetRelay(_addr common.Address, _id *big.Int) (struct {
	SentBytes          *big.Int
	SentBytesSignature []byte
	SenderPublicKey    []byte
	Ids                [3][]*big.Int
	KeysN              [2][][]byte
	KeysE              [2][]*big.Int
	Signatures         [2][][]byte
}, error) {
	return _RelayHandler.Contract.GetRelay(&_RelayHandler.CallOpts, _addr, _id)
}

// GetRelay is a free data retrieval call binding the contract method 0xaea7a62d.
//
// Solidity: function getRelay(_addr address, _id uint256) constant returns(sentBytes uint128, sentBytesSignature bytes, senderPublicKey bytes, ids uint128[][3], keysN bytes[][2], keysE uint256[][2], signatures bytes[][2])
func (_RelayHandler *RelayHandlerCallerSession) GetRelay(_addr common.Address, _id *big.Int) (struct {
	SentBytes          *big.Int
	SentBytesSignature []byte
	SenderPublicKey    []byte
	Ids                [3][]*big.Int
	KeysN              [2][][]byte
	KeysE              [2][]*big.Int
	Signatures         [2][][]byte
}, error) {
	return _RelayHandler.Contract.GetRelay(&_RelayHandler.CallOpts, _addr, _id)
}

// SubmitRelay is a paid mutator transaction binding the contract method 0x112ddfb7.
//
// Solidity: function submitRelay(_sentBytes uint128, _sentBytesHash uint256, _sentBytesSignature bytes, _senderPublicKey bytes, _ids uint128[][3], _keysN bytes[][2], _keysE uint256[][2], _signatures bytes[][2], _porRaw bytes[]) returns(uint256)
func (_RelayHandler *RelayHandlerTransactor) SubmitRelay(opts *bind.TransactOpts, _sentBytes *big.Int, _sentBytesHash *big.Int, _sentBytesSignature []byte, _senderPublicKey []byte, _ids [3][]*big.Int, _keysN [2][][]byte, _keysE [2][]*big.Int, _signatures [2][][]byte, _porRaw [][]byte) (*types.Transaction, error) {
	return _RelayHandler.contract.Transact(opts, "submitRelay", _sentBytes, _sentBytesHash, _sentBytesSignature, _senderPublicKey, _ids, _keysN, _keysE, _signatures, _porRaw)
}

// SubmitRelay is a paid mutator transaction binding the contract method 0x112ddfb7.
//
// Solidity: function submitRelay(_sentBytes uint128, _sentBytesHash uint256, _sentBytesSignature bytes, _senderPublicKey bytes, _ids uint128[][3], _keysN bytes[][2], _keysE uint256[][2], _signatures bytes[][2], _porRaw bytes[]) returns(uint256)
func (_RelayHandler *RelayHandlerSession) SubmitRelay(_sentBytes *big.Int, _sentBytesHash *big.Int, _sentBytesSignature []byte, _senderPublicKey []byte, _ids [3][]*big.Int, _keysN [2][][]byte, _keysE [2][]*big.Int, _signatures [2][][]byte, _porRaw [][]byte) (*types.Transaction, error) {
	return _RelayHandler.Contract.SubmitRelay(&_RelayHandler.TransactOpts, _sentBytes, _sentBytesHash, _sentBytesSignature, _senderPublicKey, _ids, _keysN, _keysE, _signatures, _porRaw)
}

// SubmitRelay is a paid mutator transaction binding the contract method 0x112ddfb7.
//
// Solidity: function submitRelay(_sentBytes uint128, _sentBytesHash uint256, _sentBytesSignature bytes, _senderPublicKey bytes, _ids uint128[][3], _keysN bytes[][2], _keysE uint256[][2], _signatures bytes[][2], _porRaw bytes[]) returns(uint256)
func (_RelayHandler *RelayHandlerTransactorSession) SubmitRelay(_sentBytes *big.Int, _sentBytesHash *big.Int, _sentBytesSignature []byte, _senderPublicKey []byte, _ids [3][]*big.Int, _keysN [2][][]byte, _keysE [2][]*big.Int, _signatures [2][][]byte, _porRaw [][]byte) (*types.Transaction, error) {
	return _RelayHandler.Contract.SubmitRelay(&_RelayHandler.TransactOpts, _sentBytes, _sentBytesHash, _sentBytesSignature, _senderPublicKey, _ids, _keysN, _keysE, _signatures, _porRaw)
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

// RelayHandlerRelayPaymentReqeustedIterator is returned from FilterRelayPaymentReqeusted and is used to iterate over the raw logs and unpacked data for RelayPaymentReqeusted events raised by the RelayHandler contract.
type RelayHandlerRelayPaymentReqeustedIterator struct {
	Event *RelayHandlerRelayPaymentReqeusted // Event containing the contract specifics and raw log

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
func (it *RelayHandlerRelayPaymentReqeustedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHandlerRelayPaymentReqeusted)
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
		it.Event = new(RelayHandlerRelayPaymentReqeusted)
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
func (it *RelayHandlerRelayPaymentReqeustedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHandlerRelayPaymentReqeustedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHandlerRelayPaymentReqeusted represents a RelayPaymentReqeusted event raised by the RelayHandler contract.
type RelayHandlerRelayPaymentReqeusted struct {
	common.Address
	*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRelayPaymentReqeusted is a free log retrieval operation binding the contract event 0x1a9bb4ab2c55e44b41b759f4c5a04cef4bdb8340b39a12ff37d9d0f62f3b9bfe.
//
// Solidity: e RelayPaymentReqeusted( address,  uint256)
func (_RelayHandler *RelayHandlerFilterer) FilterRelayPaymentReqeusted(opts *bind.FilterOpts) (*RelayHandlerRelayPaymentReqeustedIterator, error) {

	logs, sub, err := _RelayHandler.contract.FilterLogs(opts, "RelayPaymentReqeusted")
	if err != nil {
		return nil, err
	}
	return &RelayHandlerRelayPaymentReqeustedIterator{contract: _RelayHandler.contract, event: "RelayPaymentReqeusted", logs: logs, sub: sub}, nil
}

// WatchRelayPaymentReqeusted is a free log subscription operation binding the contract event 0x1a9bb4ab2c55e44b41b759f4c5a04cef4bdb8340b39a12ff37d9d0f62f3b9bfe.
//
// Solidity: e RelayPaymentReqeusted( address,  uint256)
func (_RelayHandler *RelayHandlerFilterer) WatchRelayPaymentReqeusted(opts *bind.WatchOpts, sink chan<- *RelayHandlerRelayPaymentReqeusted) (event.Subscription, error) {

	logs, sub, err := _RelayHandler.contract.WatchLogs(opts, "RelayPaymentReqeusted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHandlerRelayPaymentReqeusted)
				if err := _RelayHandler.contract.UnpackLog(event, "RelayPaymentReqeusted", log); err != nil {
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
