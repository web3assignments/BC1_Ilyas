// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"ChatRoomClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChatRoomOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"}],\"name\":\"MessageBroadcasted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"ParticipantJoined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"ParticipantLeft\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"}],\"name\":\"broadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParticipantCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOpened\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x60806040526000600360006101000a81548160ff0219169083600181111561002357fe5b021790555034801561003457600080fd5b50610f24806100446000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063c822d7f01161005b578063c822d7f014610192578063f2c298be14610265578063f8ca7b2814610320578063fcfff16f1461047257610088565b806343d726d61461008d57806347535d7b146100975780636598a1ae146100b9578063ad60572914610174575b600080fd5b61009561047c565b005b61009f6104ce565b604051808215151515815260200191505060405180910390f35b610172600480360360208110156100cf57600080fd5b81019080803590602001906401000000008111156100ec57600080fd5b8201836020820111156100fe57600080fd5b8035906020019184600183028401116401000000008311171561012057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506104fd565b005b61017c6107a4565b6040518082815260200191505060405180910390f35b61024b600480360360208110156101a857600080fd5b81019080803590602001906401000000008111156101c557600080fd5b8201836020820111156101d757600080fd5b803590602001918460018302840111640100000000831117156101f957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107b1565b604051808215151515815260200191505060405180910390f35b61031e6004803603602081101561027b57600080fd5b810190808035906020019064010000000081111561029857600080fd5b8201836020820111156102aa57600080fd5b803590602001918460018302840111640100000000831117156102cc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061082a565b005b6104706004803603604081101561033657600080fd5b810190808035906020019064010000000081111561035357600080fd5b82018360208201111561036557600080fd5b8035906020019184600183028401116401000000008311171561038757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156103ea57600080fd5b8201836020820111156103fc57600080fd5b8035906020019184600183028401116401000000008311171561041e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610afa565b005b61047a610d6f565b005b6001600360006101000a81548160ff0219169083600181111561049b57fe5b02179055507f228b0adf58c434f90f0f19d95b5d5d4fc677b4d6b1b20c8d1f4ade69f9058de660405160405180910390a1565b60008060018111156104dc57fe5b600360009054906101000a900460ff1660018111156104f757fe5b14905090565b600080826040518082805190602001908083835b602083106105345780518252602082019150602081019050602083039250610511565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000015414156105df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f5573657220646f6573206e6f742065786973740000000000000000000000000081525060200191505060405180910390fd5b60016000826040518082805190602001908083835b6020831061061757805182526020820191506020810190506020830392506105f4565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600001548154811061065857fe5b906000526020600020906002020160008082016000905560018201600061067f9190610dc1565b50506000816040518082805190602001908083835b602083106106b75780518252602082019150602081019050602083039250610694565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000808201600090556001820160006107039190610dc1565b50507fd386af6b577d8c71d173442a8cab0a172e24a7a9047ee284da4aa74081a05053816040518080602001828103825283818151815260200191508051906020019080838360005b8381101561076757808201518184015260208101905061074c565b50505050905090810190601f1680156107945780820380516001836020036101000a031916815260200191505b509250505060405180910390a150565b6000600180549050905090565b6000806000836040518082805190602001908083835b602083106107ea57805182526020820191506020810190506020830392506107c7565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000015414159050919050565b610832610e09565b6000826040518082805190602001908083835b602083106108685780518252602082019150602081019050602083039250610845565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060405180604001604052908160008201548152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109495780601f1061091e57610100808354040283529160200191610949565b820191906000526020600020905b81548152906001019060200180831161092c57829003601f168201915b505050505081525050905081816020018190525060016109676107a4565b01816000018181525050806000836040518082805190602001908083835b602083106109a85780518252602082019150602081019050602083039250610985565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600082015181600001556020820151816001019080519060200190610a01929190610e23565b509050506001819080600181540180825580915050600190039060005260206000209060020201600090919091909150600082015181600001556020820151816001019080519060200190610a57929190610e23565b5050507fe7ad261bb7ea20d1848a57c9d6084b77f0190124e122bcb41323a554f2009092826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610abc578082015181840152602081019050610aa1565b50505050905090810190601f168015610ae95780820380516001836020036101000a031916815260200191505b509250505060405180910390a15050565b600180811115610b0657fe5b600360009054906101000a900460ff166001811115610b2157fe5b1415610b95576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f4368617420726f6f6d20697320636c6f7365640000000000000000000000000081525060200191505060405180910390fd5b600081511415610c0d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f43616e6e6f74207075626c69736820656d70747920746578742076616c75657381525060200191505060405180910390fd5b6002604051806020016040528083815250908060018154018082558091505060019003906000526020600020016000909190919091506000820151816000019080519060200190610c5f929190610e23565b5050507f0977aadd0ff28d010547fbd2fd5260eaaef85a93d238b230c08377267742db7b8282604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610cc9578082015181840152602081019050610cae565b50505050905090810190601f168015610cf65780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610d2f578082015181840152602081019050610d14565b50505050905090810190601f168015610d5c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15050565b6000600360006101000a81548160ff02191690836001811115610d8e57fe5b02179055507f8879568228d56ae9c174f11f1089ce290d8853217f592250dabf329be26bbd3760405160405180910390a1565b50805460018160011615610100020316600290046000825580601f10610de75750610e06565b601f016020900490600052602060002090810190610e059190610ea3565b5b50565b604051806040016040528060008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e6457805160ff1916838001178555610e92565b82800160010185558215610e92579182015b82811115610e91578251825591602001919060010190610e76565b5b509050610e9f9190610ea3565b5090565b610ec591905b80821115610ec1576000816000905550600101610ea9565b5090565b9056fea2646970667358221220a555071cd0c497c1648642a3c723702f9218e0a47ac7b7dda19b2eed366d9cb664736f6c637828302e362e302d646576656c6f702e323031392e31322e31342b636f6d6d69742e31633031633639650059"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Broadcast is a paid mutator transaction binding the contract method 0xf8ca7b28.
//
// Solidity: function broadcast(string displayName, string text) returns()
func (_Store *StoreTransactor) Broadcast(opts *bind.TransactOpts, displayName string, text string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "broadcast", displayName, text)
}

// Broadcast is a paid mutator transaction binding the contract method 0xf8ca7b28.
//
// Solidity: function broadcast(string displayName, string text) returns()
func (_Store *StoreSession) Broadcast(displayName string, text string) (*types.Transaction, error) {
	return _Store.Contract.Broadcast(&_Store.TransactOpts, displayName, text)
}

// Broadcast is a paid mutator transaction binding the contract method 0xf8ca7b28.
//
// Solidity: function broadcast(string displayName, string text) returns()
func (_Store *StoreTransactorSession) Broadcast(displayName string, text string) (*types.Transaction, error) {
	return _Store.Contract.Broadcast(&_Store.TransactOpts, displayName, text)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Store *StoreTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Store *StoreSession) Close() (*types.Transaction, error) {
	return _Store.Contract.Close(&_Store.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Store *StoreTransactorSession) Close() (*types.Transaction, error) {
	return _Store.Contract.Close(&_Store.TransactOpts)
}

// GetParticipantCount is a paid mutator transaction binding the contract method 0xad605729.
//
// Solidity: function getParticipantCount() returns(uint256 count)
func (_Store *StoreTransactor) GetParticipantCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "getParticipantCount")
}

// GetParticipantCount is a paid mutator transaction binding the contract method 0xad605729.
//
// Solidity: function getParticipantCount() returns(uint256 count)
func (_Store *StoreSession) GetParticipantCount() (*types.Transaction, error) {
	return _Store.Contract.GetParticipantCount(&_Store.TransactOpts)
}

// GetParticipantCount is a paid mutator transaction binding the contract method 0xad605729.
//
// Solidity: function getParticipantCount() returns(uint256 count)
func (_Store *StoreTransactorSession) GetParticipantCount() (*types.Transaction, error) {
	return _Store.Contract.GetParticipantCount(&_Store.TransactOpts)
}

// IsOpen is a paid mutator transaction binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() returns(bool isOpened)
func (_Store *StoreTransactor) IsOpen(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "isOpen")
}

// IsOpen is a paid mutator transaction binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() returns(bool isOpened)
func (_Store *StoreSession) IsOpen() (*types.Transaction, error) {
	return _Store.Contract.IsOpen(&_Store.TransactOpts)
}

// IsOpen is a paid mutator transaction binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() returns(bool isOpened)
func (_Store *StoreTransactorSession) IsOpen() (*types.Transaction, error) {
	return _Store.Contract.IsOpen(&_Store.TransactOpts)
}

// IsRegistered is a paid mutator transaction binding the contract method 0xc822d7f0.
//
// Solidity: function isRegistered(string displayName) returns(bool exists)
func (_Store *StoreTransactor) IsRegistered(opts *bind.TransactOpts, displayName string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "isRegistered", displayName)
}

// IsRegistered is a paid mutator transaction binding the contract method 0xc822d7f0.
//
// Solidity: function isRegistered(string displayName) returns(bool exists)
func (_Store *StoreSession) IsRegistered(displayName string) (*types.Transaction, error) {
	return _Store.Contract.IsRegistered(&_Store.TransactOpts, displayName)
}

// IsRegistered is a paid mutator transaction binding the contract method 0xc822d7f0.
//
// Solidity: function isRegistered(string displayName) returns(bool exists)
func (_Store *StoreTransactorSession) IsRegistered(displayName string) (*types.Transaction, error) {
	return _Store.Contract.IsRegistered(&_Store.TransactOpts, displayName)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Store *StoreTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Store *StoreSession) Open() (*types.Transaction, error) {
	return _Store.Contract.Open(&_Store.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Store *StoreTransactorSession) Open() (*types.Transaction, error) {
	return _Store.Contract.Open(&_Store.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string displayName) returns()
func (_Store *StoreTransactor) Register(opts *bind.TransactOpts, displayName string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "register", displayName)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string displayName) returns()
func (_Store *StoreSession) Register(displayName string) (*types.Transaction, error) {
	return _Store.Contract.Register(&_Store.TransactOpts, displayName)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string displayName) returns()
func (_Store *StoreTransactorSession) Register(displayName string) (*types.Transaction, error) {
	return _Store.Contract.Register(&_Store.TransactOpts, displayName)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string displayName) returns()
func (_Store *StoreTransactor) Unregister(opts *bind.TransactOpts, displayName string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "unregister", displayName)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string displayName) returns()
func (_Store *StoreSession) Unregister(displayName string) (*types.Transaction, error) {
	return _Store.Contract.Unregister(&_Store.TransactOpts, displayName)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string displayName) returns()
func (_Store *StoreTransactorSession) Unregister(displayName string) (*types.Transaction, error) {
	return _Store.Contract.Unregister(&_Store.TransactOpts, displayName)
}

// StoreChatRoomClosedIterator is returned from FilterChatRoomClosed and is used to iterate over the raw logs and unpacked data for ChatRoomClosed events raised by the Store contract.
type StoreChatRoomClosedIterator struct {
	Event *StoreChatRoomClosed // Event containing the contract specifics and raw log

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
func (it *StoreChatRoomClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreChatRoomClosed)
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
		it.Event = new(StoreChatRoomClosed)
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
func (it *StoreChatRoomClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreChatRoomClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreChatRoomClosed represents a ChatRoomClosed event raised by the Store contract.
type StoreChatRoomClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChatRoomClosed is a free log retrieval operation binding the contract event 0x228b0adf58c434f90f0f19d95b5d5d4fc677b4d6b1b20c8d1f4ade69f9058de6.
//
// Solidity: event ChatRoomClosed()
func (_Store *StoreFilterer) FilterChatRoomClosed(opts *bind.FilterOpts) (*StoreChatRoomClosedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ChatRoomClosed")
	if err != nil {
		return nil, err
	}
	return &StoreChatRoomClosedIterator{contract: _Store.contract, event: "ChatRoomClosed", logs: logs, sub: sub}, nil
}

// WatchChatRoomClosed is a free log subscription operation binding the contract event 0x228b0adf58c434f90f0f19d95b5d5d4fc677b4d6b1b20c8d1f4ade69f9058de6.
//
// Solidity: event ChatRoomClosed()
func (_Store *StoreFilterer) WatchChatRoomClosed(opts *bind.WatchOpts, sink chan<- *StoreChatRoomClosed) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ChatRoomClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreChatRoomClosed)
				if err := _Store.contract.UnpackLog(event, "ChatRoomClosed", log); err != nil {
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

// ParseChatRoomClosed is a log parse operation binding the contract event 0x228b0adf58c434f90f0f19d95b5d5d4fc677b4d6b1b20c8d1f4ade69f9058de6.
//
// Solidity: event ChatRoomClosed()
func (_Store *StoreFilterer) ParseChatRoomClosed(log types.Log) (*StoreChatRoomClosed, error) {
	event := new(StoreChatRoomClosed)
	if err := _Store.contract.UnpackLog(event, "ChatRoomClosed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreChatRoomOpenedIterator is returned from FilterChatRoomOpened and is used to iterate over the raw logs and unpacked data for ChatRoomOpened events raised by the Store contract.
type StoreChatRoomOpenedIterator struct {
	Event *StoreChatRoomOpened // Event containing the contract specifics and raw log

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
func (it *StoreChatRoomOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreChatRoomOpened)
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
		it.Event = new(StoreChatRoomOpened)
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
func (it *StoreChatRoomOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreChatRoomOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreChatRoomOpened represents a ChatRoomOpened event raised by the Store contract.
type StoreChatRoomOpened struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChatRoomOpened is a free log retrieval operation binding the contract event 0x8879568228d56ae9c174f11f1089ce290d8853217f592250dabf329be26bbd37.
//
// Solidity: event ChatRoomOpened()
func (_Store *StoreFilterer) FilterChatRoomOpened(opts *bind.FilterOpts) (*StoreChatRoomOpenedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ChatRoomOpened")
	if err != nil {
		return nil, err
	}
	return &StoreChatRoomOpenedIterator{contract: _Store.contract, event: "ChatRoomOpened", logs: logs, sub: sub}, nil
}

// WatchChatRoomOpened is a free log subscription operation binding the contract event 0x8879568228d56ae9c174f11f1089ce290d8853217f592250dabf329be26bbd37.
//
// Solidity: event ChatRoomOpened()
func (_Store *StoreFilterer) WatchChatRoomOpened(opts *bind.WatchOpts, sink chan<- *StoreChatRoomOpened) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ChatRoomOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreChatRoomOpened)
				if err := _Store.contract.UnpackLog(event, "ChatRoomOpened", log); err != nil {
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

// ParseChatRoomOpened is a log parse operation binding the contract event 0x8879568228d56ae9c174f11f1089ce290d8853217f592250dabf329be26bbd37.
//
// Solidity: event ChatRoomOpened()
func (_Store *StoreFilterer) ParseChatRoomOpened(log types.Log) (*StoreChatRoomOpened, error) {
	event := new(StoreChatRoomOpened)
	if err := _Store.contract.UnpackLog(event, "ChatRoomOpened", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreMessageBroadcastedIterator is returned from FilterMessageBroadcasted and is used to iterate over the raw logs and unpacked data for MessageBroadcasted events raised by the Store contract.
type StoreMessageBroadcastedIterator struct {
	Event *StoreMessageBroadcasted // Event containing the contract specifics and raw log

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
func (it *StoreMessageBroadcastedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreMessageBroadcasted)
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
		it.Event = new(StoreMessageBroadcasted)
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
func (it *StoreMessageBroadcastedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreMessageBroadcastedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreMessageBroadcasted represents a MessageBroadcasted event raised by the Store contract.
type StoreMessageBroadcasted struct {
	DisplayName string
	Text        string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageBroadcasted is a free log retrieval operation binding the contract event 0x0977aadd0ff28d010547fbd2fd5260eaaef85a93d238b230c08377267742db7b.
//
// Solidity: event MessageBroadcasted(string displayName, string text)
func (_Store *StoreFilterer) FilterMessageBroadcasted(opts *bind.FilterOpts) (*StoreMessageBroadcastedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "MessageBroadcasted")
	if err != nil {
		return nil, err
	}
	return &StoreMessageBroadcastedIterator{contract: _Store.contract, event: "MessageBroadcasted", logs: logs, sub: sub}, nil
}

// WatchMessageBroadcasted is a free log subscription operation binding the contract event 0x0977aadd0ff28d010547fbd2fd5260eaaef85a93d238b230c08377267742db7b.
//
// Solidity: event MessageBroadcasted(string displayName, string text)
func (_Store *StoreFilterer) WatchMessageBroadcasted(opts *bind.WatchOpts, sink chan<- *StoreMessageBroadcasted) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "MessageBroadcasted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreMessageBroadcasted)
				if err := _Store.contract.UnpackLog(event, "MessageBroadcasted", log); err != nil {
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

// ParseMessageBroadcasted is a log parse operation binding the contract event 0x0977aadd0ff28d010547fbd2fd5260eaaef85a93d238b230c08377267742db7b.
//
// Solidity: event MessageBroadcasted(string displayName, string text)
func (_Store *StoreFilterer) ParseMessageBroadcasted(log types.Log) (*StoreMessageBroadcasted, error) {
	event := new(StoreMessageBroadcasted)
	if err := _Store.contract.UnpackLog(event, "MessageBroadcasted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreParticipantJoinedIterator is returned from FilterParticipantJoined and is used to iterate over the raw logs and unpacked data for ParticipantJoined events raised by the Store contract.
type StoreParticipantJoinedIterator struct {
	Event *StoreParticipantJoined // Event containing the contract specifics and raw log

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
func (it *StoreParticipantJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreParticipantJoined)
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
		it.Event = new(StoreParticipantJoined)
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
func (it *StoreParticipantJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreParticipantJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreParticipantJoined represents a ParticipantJoined event raised by the Store contract.
type StoreParticipantJoined struct {
	DisplayName string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantJoined is a free log retrieval operation binding the contract event 0xe7ad261bb7ea20d1848a57c9d6084b77f0190124e122bcb41323a554f2009092.
//
// Solidity: event ParticipantJoined(string displayName)
func (_Store *StoreFilterer) FilterParticipantJoined(opts *bind.FilterOpts) (*StoreParticipantJoinedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ParticipantJoined")
	if err != nil {
		return nil, err
	}
	return &StoreParticipantJoinedIterator{contract: _Store.contract, event: "ParticipantJoined", logs: logs, sub: sub}, nil
}

// WatchParticipantJoined is a free log subscription operation binding the contract event 0xe7ad261bb7ea20d1848a57c9d6084b77f0190124e122bcb41323a554f2009092.
//
// Solidity: event ParticipantJoined(string displayName)
func (_Store *StoreFilterer) WatchParticipantJoined(opts *bind.WatchOpts, sink chan<- *StoreParticipantJoined) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ParticipantJoined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreParticipantJoined)
				if err := _Store.contract.UnpackLog(event, "ParticipantJoined", log); err != nil {
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

// ParseParticipantJoined is a log parse operation binding the contract event 0xe7ad261bb7ea20d1848a57c9d6084b77f0190124e122bcb41323a554f2009092.
//
// Solidity: event ParticipantJoined(string displayName)
func (_Store *StoreFilterer) ParseParticipantJoined(log types.Log) (*StoreParticipantJoined, error) {
	event := new(StoreParticipantJoined)
	if err := _Store.contract.UnpackLog(event, "ParticipantJoined", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StoreParticipantLeftIterator is returned from FilterParticipantLeft and is used to iterate over the raw logs and unpacked data for ParticipantLeft events raised by the Store contract.
type StoreParticipantLeftIterator struct {
	Event *StoreParticipantLeft // Event containing the contract specifics and raw log

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
func (it *StoreParticipantLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreParticipantLeft)
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
		it.Event = new(StoreParticipantLeft)
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
func (it *StoreParticipantLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreParticipantLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreParticipantLeft represents a ParticipantLeft event raised by the Store contract.
type StoreParticipantLeft struct {
	DisplayName string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantLeft is a free log retrieval operation binding the contract event 0xd386af6b577d8c71d173442a8cab0a172e24a7a9047ee284da4aa74081a05053.
//
// Solidity: event ParticipantLeft(string displayName)
func (_Store *StoreFilterer) FilterParticipantLeft(opts *bind.FilterOpts) (*StoreParticipantLeftIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ParticipantLeft")
	if err != nil {
		return nil, err
	}
	return &StoreParticipantLeftIterator{contract: _Store.contract, event: "ParticipantLeft", logs: logs, sub: sub}, nil
}

// WatchParticipantLeft is a free log subscription operation binding the contract event 0xd386af6b577d8c71d173442a8cab0a172e24a7a9047ee284da4aa74081a05053.
//
// Solidity: event ParticipantLeft(string displayName)
func (_Store *StoreFilterer) WatchParticipantLeft(opts *bind.WatchOpts, sink chan<- *StoreParticipantLeft) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ParticipantLeft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreParticipantLeft)
				if err := _Store.contract.UnpackLog(event, "ParticipantLeft", log); err != nil {
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

// ParseParticipantLeft is a log parse operation binding the contract event 0xd386af6b577d8c71d173442a8cab0a172e24a7a9047ee284da4aa74081a05053.
//
// Solidity: event ParticipantLeft(string displayName)
func (_Store *StoreFilterer) ParseParticipantLeft(log types.Log) (*StoreParticipantLeft, error) {
	event := new(StoreParticipantLeft)
	if err := _Store.contract.UnpackLog(event, "ParticipantLeft", log); err != nil {
		return nil, err
	}
	return event, nil
}
