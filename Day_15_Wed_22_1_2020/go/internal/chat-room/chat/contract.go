// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chat

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

// ChannelABI is the input ABI used to generate the binding from.
const ChannelABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doesExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChannelFuncSigs maps the 4-byte function signature to its string representation.
var ChannelFuncSigs = map[string]string{
	"83197ef0": "destroy()",
	"5ad70d12": "doesExist()",
	"267c4ae4": "exists()",
}

// ChannelBin is the compiled bytecode used for deploying new contracts.
var ChannelBin = "0x608060405234801561001057600080fd5b5060008054600160ff1990911617610100600160a81b031916610100330217815560f190819061004090396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c8063267c4ae41460415780635ad70d1214605b57806383197ef0146061575b600080fd5b60476069565b604080519115158252519081900360200190f35b60476072565b6067607b565b005b60005460ff1681565b60005460ff1690565b6000805460ff19169081905561010090046001600160a01b0316fffea2646970667358221220bb72c6c85976b3ce16afbab76752ba06c5c59e304e626172fe0a40b20bfbe04264736f6c637826302e362e332d646576656c6f702e323032302e322e342b636f6d6d69742e38613765316436350057"

// DeployChannel deploys a new Ethereum contract, binding an instance of Channel to it.
func DeployChannel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Channel, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChannelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Channel{ChannelCaller: ChannelCaller{contract: contract}, ChannelTransactor: ChannelTransactor{contract: contract}, ChannelFilterer: ChannelFilterer{contract: contract}}, nil
}

// Channel is an auto generated Go binding around an Ethereum contract.
type Channel struct {
	ChannelCaller     // Read-only binding to the contract
	ChannelTransactor // Write-only binding to the contract
	ChannelFilterer   // Log filterer for contract events
}

// ChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelSession struct {
	Contract     *Channel          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelCallerSession struct {
	Contract *ChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelTransactorSession struct {
	Contract     *ChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelRaw struct {
	Contract *Channel // Generic contract binding to access the raw methods on
}

// ChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelCallerRaw struct {
	Contract *ChannelCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelTransactorRaw struct {
	Contract *ChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannel creates a new instance of Channel, bound to a specific deployed contract.
func NewChannel(address common.Address, backend bind.ContractBackend) (*Channel, error) {
	contract, err := bindChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Channel{ChannelCaller: ChannelCaller{contract: contract}, ChannelTransactor: ChannelTransactor{contract: contract}, ChannelFilterer: ChannelFilterer{contract: contract}}, nil
}

// NewChannelCaller creates a new read-only instance of Channel, bound to a specific deployed contract.
func NewChannelCaller(address common.Address, caller bind.ContractCaller) (*ChannelCaller, error) {
	contract, err := bindChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelCaller{contract: contract}, nil
}

// NewChannelTransactor creates a new write-only instance of Channel, bound to a specific deployed contract.
func NewChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelTransactor, error) {
	contract, err := bindChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelTransactor{contract: contract}, nil
}

// NewChannelFilterer creates a new log filterer instance of Channel, bound to a specific deployed contract.
func NewChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelFilterer, error) {
	contract, err := bindChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelFilterer{contract: contract}, nil
}

// bindChannel binds a generic wrapper to an already deployed contract.
func bindChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.ChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transact(opts, method, params...)
}

// DoesExist is a free data retrieval call binding the contract method 0x5ad70d12.
//
// Solidity: function doesExist() constant returns(bool)
func (_Channel *ChannelCaller) DoesExist(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Channel.contract.Call(opts, out, "doesExist")
	return *ret0, err
}

// DoesExist is a free data retrieval call binding the contract method 0x5ad70d12.
//
// Solidity: function doesExist() constant returns(bool)
func (_Channel *ChannelSession) DoesExist() (bool, error) {
	return _Channel.Contract.DoesExist(&_Channel.CallOpts)
}

// DoesExist is a free data retrieval call binding the contract method 0x5ad70d12.
//
// Solidity: function doesExist() constant returns(bool)
func (_Channel *ChannelCallerSession) DoesExist() (bool, error) {
	return _Channel.Contract.DoesExist(&_Channel.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x267c4ae4.
//
// Solidity: function exists() constant returns(bool)
func (_Channel *ChannelCaller) Exists(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Channel.contract.Call(opts, out, "exists")
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x267c4ae4.
//
// Solidity: function exists() constant returns(bool)
func (_Channel *ChannelSession) Exists() (bool, error) {
	return _Channel.Contract.Exists(&_Channel.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x267c4ae4.
//
// Solidity: function exists() constant returns(bool)
func (_Channel *ChannelCallerSession) Exists() (bool, error) {
	return _Channel.Contract.Exists(&_Channel.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Channel *ChannelTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Channel *ChannelSession) Destroy() (*types.Transaction, error) {
	return _Channel.Contract.Destroy(&_Channel.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Channel *ChannelTransactorSession) Destroy() (*types.Transaction, error) {
	return _Channel.Contract.Destroy(&_Channel.TransactOpts)
}

// ChatRoomABI is the input ABI used to generate the binding from.
const ChatRoomABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ChannelCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ChannelDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChatRoomClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ChatRoomOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"channelName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"}],\"name\":\"MessageBroadcasted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"ParticipantJoined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"ParticipantLeft\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOpened\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"}],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChatRoomFuncSigs maps the 4-byte function signature to its string representation.
var ChatRoomFuncSigs = map[string]string{
	"43d726d6": "close()",
	"47535d7b": "isOpen()",
	"c822d7f0": "isRegistered(string)",
	"fcfff16f": "open()",
	"f2c298be": "register(string)",
	"5615264d": "send(string,string,string)",
	"6598a1ae": "unregister(string)",
}

// ChatRoomBin is the compiled bytecode used for deploying new contracts.
var ChatRoomBin = "0x60806040526002805460ff1916905534801561001a57600080fd5b5060408051808201909152600781526611d95b995c985b60ca1b602082015261004b906001600160e01b0361008216565b6040805180820190915260098152684f66662d546f70696360b81b602082015261007d906001600160e01b0361008216565b6101d0565b6000604051610090906101c3565b604051809103906000f0801580156100ac573d6000803e3d6000fd5b509050806001836040518082805190602001908083835b602083106100e25780518252601f1990920191602091820191016100c3565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201852080546001600160a01b0319166001600160a01b03979097169690961790955580845286518482015286517fcf3bc676593dbd38dcb2832540b70f009982838d486233f971ebd306a79b27af958895945084935083019185019080838360005b8381101561018557818101518382015260200161016d565b50505050905090810190601f1680156101b25780820380516001836020036101000a031916815260200191505b509250505060405180910390a15050565b61013180610c9b83390190565b610abc806101df6000396000f3fe6080604052600436106100705760003560e01c80636598a1ae1161004e5780636598a1ae146101d0578063c822d7f014610281578063f2c298be14610332578063fcfff16f146103d657610070565b806343d726d61461007557806347535d7b1461008c5780635615264d146100b5575b600080fd5b34801561008157600080fd5b5061008a6103eb565b005b34801561009857600080fd5b506100a1610423565b604080519115158252519081900360200190f35b3480156100c157600080fd5b5061008a600480360360608110156100d857600080fd5b810190602081018135600160201b8111156100f257600080fd5b82018360208201111561010457600080fd5b803590602001918460018302840111600160201b8311171561012557600080fd5b919390929091602081019035600160201b81111561014257600080fd5b82018360208201111561015457600080fd5b803590602001918460018302840111600160201b8311171561017557600080fd5b919390929091602081019035600160201b81111561019257600080fd5b8201836020820111156101a457600080fd5b803590602001918460018302840111600160201b831117156101c557600080fd5b50909250905061043e565b3480156101dc57600080fd5b5061008a600480360360208110156101f357600080fd5b810190602081018135600160201b81111561020d57600080fd5b82018360208201111561021f57600080fd5b803590602001918460018302840111600160201b8311171561024057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610599945050505050565b34801561028d57600080fd5b506100a1600480360360208110156102a457600080fd5b810190602081018135600160201b8111156102be57600080fd5b8201836020820111156102d057600080fd5b803590602001918460018302840111600160201b831117156102f157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106a0945050505050565b61008a6004803603602081101561034857600080fd5b810190602081018135600160201b81111561036257600080fd5b82018360208201111561037457600080fd5b803590602001918460018302840111600160201b8311171561039557600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061070e945050505050565b3480156103e257600080fd5b5061008a610972565b6002805460ff191660011790556040517f228b0adf58c434f90f0f19d95b5d5d4fc677b4d6b1b20c8d1f4ade69f9058de690600090a1565b60008060025460ff16600181111561043757fe5b1490505b90565b60006001878760405180838380828437919091019485525050604080516020948190038501812054632d6b868960e11b825291516001600160a01b0390921695508594635ad70d12945060048083019450909291829003018186803b1580156104a657600080fd5b505afa1580156104ba573d6000803e3d6000fd5b505050506040513d60208110156104d057600080fd5b505115610590577fef5b6616d0ef7a7345c7d7d30de504c95f981d0f19c9a1559d5e1ee093a1bb7c8787878787876040518080602001806020018060200184810384528a8a82818152602001925080828437600083820152601f01601f191690910185810384528881526020019050888880828437600083820152601f01601f191690910185810383528681526020019050868680828437600083820152604051601f909101601f19169092018290039b50909950505050505050505050a15b50505050505050565b600080826040518082805190602001908083835b602083106105cc5780518252601f1990920191602091820191016105ad565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520600201805460ff19169615159690961790955580845285518482015285517fd386af6b577d8c71d173442a8cab0a172e24a7a9047ee284da4aa74081a05053958795945084935083019185019080838360005b8381101561066357818101518382015260200161064b565b50505050905090810190601f1680156106905780820380516001836020036101000a031916815260200191505b509250505060405180910390a150565b600080826040518082805190602001908083835b602083106106d35780518252601f1990920191602091820191016106b4565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206002015460ff1615949350505050565b6107166109a7565b6000826040518082805190602001908083835b602083106107485780518252601f199092019160209182019101610729565b518151602093840361010090810a60001990810180199094169390921692909217909252929094019687526040805197889003820188206060890182528054895260018082018054845160029382161590970290980190971604601f810184900484028501840190925281845296508782019550919392508301828280156108115780601f106107e657610100808354040283529160200191610811565b820191906000526020600020905b8154815290600101906020018083116107f457829003601f168201915b50505091835250506002919091015460ff1615156020918201526001604080840191909152828201859052518451929350839260009286929182918401908083835b602083106108725780518252601f199092019160209182019101610853565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451815584840151805191946108bc945060018601935001906109ca565b50604091820151600291909101805460ff19169115159190911790558051602080825284518183015284517fe7ad261bb7ea20d1848a57c9d6084b77f0190124e122bcb41323a554f2009092938693928392918301919085019080838360005b8381101561093457818101518382015260200161091c565b50505050905090810190601f1680156109615780820380516001836020036101000a031916815260200191505b509250505060405180910390a15050565b6002805460ff191690556040517f8879568228d56ae9c174f11f1089ce290d8853217f592250dabf329be26bbd3790600090a1565b604051806060016040528060008152602001606081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a0b57805160ff1916838001178555610a38565b82800160010185558215610a38579182015b82811115610a38578251825591602001919060010190610a1d565b50610a44929150610a48565b5090565b61043b91905b80821115610a445760008155600101610a4e56fea26469706673582212204f2bd2ef9cac0bd808b916fe70a22de4b0ef9c50df10959c6f9cac1b0a18612764736f6c637826302e362e332d646576656c6f702e323032302e322e342b636f6d6d69742e38613765316436350057608060405234801561001057600080fd5b5060008054600160ff1990911617610100600160a81b031916610100330217815560f190819061004090396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c8063267c4ae41460415780635ad70d1214605b57806383197ef0146061575b600080fd5b60476069565b604080519115158252519081900360200190f35b60476072565b6067607b565b005b60005460ff1681565b60005460ff1690565b6000805460ff19169081905561010090046001600160a01b0316fffea2646970667358221220bb72c6c85976b3ce16afbab76752ba06c5c59e304e626172fe0a40b20bfbe04264736f6c637826302e362e332d646576656c6f702e323032302e322e342b636f6d6d69742e38613765316436350057"

// DeployChatRoom deploys a new Ethereum contract, binding an instance of ChatRoom to it.
func DeployChatRoom(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChatRoom, error) {
	parsed, err := abi.JSON(strings.NewReader(ChatRoomABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChatRoomBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChatRoom{ChatRoomCaller: ChatRoomCaller{contract: contract}, ChatRoomTransactor: ChatRoomTransactor{contract: contract}, ChatRoomFilterer: ChatRoomFilterer{contract: contract}}, nil
}

// ChatRoom is an auto generated Go binding around an Ethereum contract.
type ChatRoom struct {
	ChatRoomCaller     // Read-only binding to the contract
	ChatRoomTransactor // Write-only binding to the contract
	ChatRoomFilterer   // Log filterer for contract events
}

// ChatRoomCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChatRoomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatRoomTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChatRoomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatRoomFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChatRoomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatRoomSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChatRoomSession struct {
	Contract     *ChatRoom         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChatRoomCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChatRoomCallerSession struct {
	Contract *ChatRoomCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ChatRoomTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChatRoomTransactorSession struct {
	Contract     *ChatRoomTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ChatRoomRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChatRoomRaw struct {
	Contract *ChatRoom // Generic contract binding to access the raw methods on
}

// ChatRoomCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChatRoomCallerRaw struct {
	Contract *ChatRoomCaller // Generic read-only contract binding to access the raw methods on
}

// ChatRoomTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChatRoomTransactorRaw struct {
	Contract *ChatRoomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChatRoom creates a new instance of ChatRoom, bound to a specific deployed contract.
func NewChatRoom(address common.Address, backend bind.ContractBackend) (*ChatRoom, error) {
	contract, err := bindChatRoom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChatRoom{ChatRoomCaller: ChatRoomCaller{contract: contract}, ChatRoomTransactor: ChatRoomTransactor{contract: contract}, ChatRoomFilterer: ChatRoomFilterer{contract: contract}}, nil
}

// NewChatRoomCaller creates a new read-only instance of ChatRoom, bound to a specific deployed contract.
func NewChatRoomCaller(address common.Address, caller bind.ContractCaller) (*ChatRoomCaller, error) {
	contract, err := bindChatRoom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChatRoomCaller{contract: contract}, nil
}

// NewChatRoomTransactor creates a new write-only instance of ChatRoom, bound to a specific deployed contract.
func NewChatRoomTransactor(address common.Address, transactor bind.ContractTransactor) (*ChatRoomTransactor, error) {
	contract, err := bindChatRoom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChatRoomTransactor{contract: contract}, nil
}

// NewChatRoomFilterer creates a new log filterer instance of ChatRoom, bound to a specific deployed contract.
func NewChatRoomFilterer(address common.Address, filterer bind.ContractFilterer) (*ChatRoomFilterer, error) {
	contract, err := bindChatRoom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChatRoomFilterer{contract: contract}, nil
}

// bindChatRoom binds a generic wrapper to an already deployed contract.
func bindChatRoom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChatRoomABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChatRoom *ChatRoomRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChatRoom.Contract.ChatRoomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChatRoom *ChatRoomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatRoom.Contract.ChatRoomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChatRoom *ChatRoomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChatRoom.Contract.ChatRoomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChatRoom *ChatRoomCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChatRoom.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChatRoom *ChatRoomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatRoom.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChatRoom *ChatRoomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChatRoom.Contract.contract.Transact(opts, method, params...)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() constant returns(bool isOpened)
func (_ChatRoom *ChatRoomCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChatRoom.contract.Call(opts, out, "isOpen")
	return *ret0, err
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() constant returns(bool isOpened)
func (_ChatRoom *ChatRoomSession) IsOpen() (bool, error) {
	return _ChatRoom.Contract.IsOpen(&_ChatRoom.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() constant returns(bool isOpened)
func (_ChatRoom *ChatRoomCallerSession) IsOpen() (bool, error) {
	return _ChatRoom.Contract.IsOpen(&_ChatRoom.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc822d7f0.
//
// Solidity: function isRegistered(string displayName) constant returns(bool exists)
func (_ChatRoom *ChatRoomCaller) IsRegistered(opts *bind.CallOpts, displayName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChatRoom.contract.Call(opts, out, "isRegistered", displayName)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc822d7f0.
//
// Solidity: function isRegistered(string displayName) constant returns(bool exists)
func (_ChatRoom *ChatRoomSession) IsRegistered(displayName string) (bool, error) {
	return _ChatRoom.Contract.IsRegistered(&_ChatRoom.CallOpts, displayName)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc822d7f0.
//
// Solidity: function isRegistered(string displayName) constant returns(bool exists)
func (_ChatRoom *ChatRoomCallerSession) IsRegistered(displayName string) (bool, error) {
	return _ChatRoom.Contract.IsRegistered(&_ChatRoom.CallOpts, displayName)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_ChatRoom *ChatRoomTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatRoom.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_ChatRoom *ChatRoomSession) Close() (*types.Transaction, error) {
	return _ChatRoom.Contract.Close(&_ChatRoom.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_ChatRoom *ChatRoomTransactorSession) Close() (*types.Transaction, error) {
	return _ChatRoom.Contract.Close(&_ChatRoom.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_ChatRoom *ChatRoomTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChatRoom.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_ChatRoom *ChatRoomSession) Open() (*types.Transaction, error) {
	return _ChatRoom.Contract.Open(&_ChatRoom.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_ChatRoom *ChatRoomTransactorSession) Open() (*types.Transaction, error) {
	return _ChatRoom.Contract.Open(&_ChatRoom.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string displayName) returns()
func (_ChatRoom *ChatRoomTransactor) Register(opts *bind.TransactOpts, displayName string) (*types.Transaction, error) {
	return _ChatRoom.contract.Transact(opts, "register", displayName)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string displayName) returns()
func (_ChatRoom *ChatRoomSession) Register(displayName string) (*types.Transaction, error) {
	return _ChatRoom.Contract.Register(&_ChatRoom.TransactOpts, displayName)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(string displayName) returns()
func (_ChatRoom *ChatRoomTransactorSession) Register(displayName string) (*types.Transaction, error) {
	return _ChatRoom.Contract.Register(&_ChatRoom.TransactOpts, displayName)
}

// Send is a paid mutator transaction binding the contract method 0x5615264d.
//
// Solidity: function send(string channelName, string displayName, string text) returns()
func (_ChatRoom *ChatRoomTransactor) Send(opts *bind.TransactOpts, channelName string, displayName string, text string) (*types.Transaction, error) {
	return _ChatRoom.contract.Transact(opts, "send", channelName, displayName, text)
}

// Send is a paid mutator transaction binding the contract method 0x5615264d.
//
// Solidity: function send(string channelName, string displayName, string text) returns()
func (_ChatRoom *ChatRoomSession) Send(channelName string, displayName string, text string) (*types.Transaction, error) {
	return _ChatRoom.Contract.Send(&_ChatRoom.TransactOpts, channelName, displayName, text)
}

// Send is a paid mutator transaction binding the contract method 0x5615264d.
//
// Solidity: function send(string channelName, string displayName, string text) returns()
func (_ChatRoom *ChatRoomTransactorSession) Send(channelName string, displayName string, text string) (*types.Transaction, error) {
	return _ChatRoom.Contract.Send(&_ChatRoom.TransactOpts, channelName, displayName, text)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string displayName) returns()
func (_ChatRoom *ChatRoomTransactor) Unregister(opts *bind.TransactOpts, displayName string) (*types.Transaction, error) {
	return _ChatRoom.contract.Transact(opts, "unregister", displayName)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string displayName) returns()
func (_ChatRoom *ChatRoomSession) Unregister(displayName string) (*types.Transaction, error) {
	return _ChatRoom.Contract.Unregister(&_ChatRoom.TransactOpts, displayName)
}

// Unregister is a paid mutator transaction binding the contract method 0x6598a1ae.
//
// Solidity: function unregister(string displayName) returns()
func (_ChatRoom *ChatRoomTransactorSession) Unregister(displayName string) (*types.Transaction, error) {
	return _ChatRoom.Contract.Unregister(&_ChatRoom.TransactOpts, displayName)
}

// ChatRoomChannelCreatedIterator is returned from FilterChannelCreated and is used to iterate over the raw logs and unpacked data for ChannelCreated events raised by the ChatRoom contract.
type ChatRoomChannelCreatedIterator struct {
	Event *ChatRoomChannelCreated // Event containing the contract specifics and raw log

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
func (it *ChatRoomChannelCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatRoomChannelCreated)
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
		it.Event = new(ChatRoomChannelCreated)
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
func (it *ChatRoomChannelCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatRoomChannelCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatRoomChannelCreated represents a ChannelCreated event raised by the ChatRoom contract.
type ChatRoomChannelCreated struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChannelCreated is a free log retrieval operation binding the contract event 0xcf3bc676593dbd38dcb2832540b70f009982838d486233f971ebd306a79b27af.
//
// Solidity: event ChannelCreated(string name)
func (_ChatRoom *ChatRoomFilterer) FilterChannelCreated(opts *bind.FilterOpts) (*ChatRoomChannelCreatedIterator, error) {

	logs, sub, err := _ChatRoom.contract.FilterLogs(opts, "ChannelCreated")
	if err != nil {
		return nil, err
	}
	return &ChatRoomChannelCreatedIterator{contract: _ChatRoom.contract, event: "ChannelCreated", logs: logs, sub: sub}, nil
}

// WatchChannelCreated is a free log subscription operation binding the contract event 0xcf3bc676593dbd38dcb2832540b70f009982838d486233f971ebd306a79b27af.
//
// Solidity: event ChannelCreated(string name)
func (_ChatRoom *ChatRoomFilterer) WatchChannelCreated(opts *bind.WatchOpts, sink chan<- *ChatRoomChannelCreated) (event.Subscription, error) {

	logs, sub, err := _ChatRoom.contract.WatchLogs(opts, "ChannelCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatRoomChannelCreated)
				if err := _ChatRoom.contract.UnpackLog(event, "ChannelCreated", log); err != nil {
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

// ParseChannelCreated is a log parse operation binding the contract event 0xcf3bc676593dbd38dcb2832540b70f009982838d486233f971ebd306a79b27af.
//
// Solidity: event ChannelCreated(string name)
func (_ChatRoom *ChatRoomFilterer) ParseChannelCreated(log types.Log) (*ChatRoomChannelCreated, error) {
	event := new(ChatRoomChannelCreated)
	if err := _ChatRoom.contract.UnpackLog(event, "ChannelCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChatRoomChannelDestroyedIterator is returned from FilterChannelDestroyed and is used to iterate over the raw logs and unpacked data for ChannelDestroyed events raised by the ChatRoom contract.
type ChatRoomChannelDestroyedIterator struct {
	Event *ChatRoomChannelDestroyed // Event containing the contract specifics and raw log

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
func (it *ChatRoomChannelDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatRoomChannelDestroyed)
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
		it.Event = new(ChatRoomChannelDestroyed)
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
func (it *ChatRoomChannelDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatRoomChannelDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatRoomChannelDestroyed represents a ChannelDestroyed event raised by the ChatRoom contract.
type ChatRoomChannelDestroyed struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChannelDestroyed is a free log retrieval operation binding the contract event 0x0835516533e30f96908f88d4b5dfd5da084e97276659ca3448c9c8c475fbbd3b.
//
// Solidity: event ChannelDestroyed(string name)
func (_ChatRoom *ChatRoomFilterer) FilterChannelDestroyed(opts *bind.FilterOpts) (*ChatRoomChannelDestroyedIterator, error) {

	logs, sub, err := _ChatRoom.contract.FilterLogs(opts, "ChannelDestroyed")
	if err != nil {
		return nil, err
	}
	return &ChatRoomChannelDestroyedIterator{contract: _ChatRoom.contract, event: "ChannelDestroyed", logs: logs, sub: sub}, nil
}

// WatchChannelDestroyed is a free log subscription operation binding the contract event 0x0835516533e30f96908f88d4b5dfd5da084e97276659ca3448c9c8c475fbbd3b.
//
// Solidity: event ChannelDestroyed(string name)
func (_ChatRoom *ChatRoomFilterer) WatchChannelDestroyed(opts *bind.WatchOpts, sink chan<- *ChatRoomChannelDestroyed) (event.Subscription, error) {

	logs, sub, err := _ChatRoom.contract.WatchLogs(opts, "ChannelDestroyed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatRoomChannelDestroyed)
				if err := _ChatRoom.contract.UnpackLog(event, "ChannelDestroyed", log); err != nil {
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

// ParseChannelDestroyed is a log parse operation binding the contract event 0x0835516533e30f96908f88d4b5dfd5da084e97276659ca3448c9c8c475fbbd3b.
//
// Solidity: event ChannelDestroyed(string name)
func (_ChatRoom *ChatRoomFilterer) ParseChannelDestroyed(log types.Log) (*ChatRoomChannelDestroyed, error) {
	event := new(ChatRoomChannelDestroyed)
	if err := _ChatRoom.contract.UnpackLog(event, "ChannelDestroyed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChatRoomChatRoomClosedIterator is returned from FilterChatRoomClosed and is used to iterate over the raw logs and unpacked data for ChatRoomClosed events raised by the ChatRoom contract.
type ChatRoomChatRoomClosedIterator struct {
	Event *ChatRoomChatRoomClosed // Event containing the contract specifics and raw log

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
func (it *ChatRoomChatRoomClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatRoomChatRoomClosed)
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
		it.Event = new(ChatRoomChatRoomClosed)
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
func (it *ChatRoomChatRoomClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatRoomChatRoomClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatRoomChatRoomClosed represents a ChatRoomClosed event raised by the ChatRoom contract.
type ChatRoomChatRoomClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChatRoomClosed is a free log retrieval operation binding the contract event 0x228b0adf58c434f90f0f19d95b5d5d4fc677b4d6b1b20c8d1f4ade69f9058de6.
//
// Solidity: event ChatRoomClosed()
func (_ChatRoom *ChatRoomFilterer) FilterChatRoomClosed(opts *bind.FilterOpts) (*ChatRoomChatRoomClosedIterator, error) {

	logs, sub, err := _ChatRoom.contract.FilterLogs(opts, "ChatRoomClosed")
	if err != nil {
		return nil, err
	}
	return &ChatRoomChatRoomClosedIterator{contract: _ChatRoom.contract, event: "ChatRoomClosed", logs: logs, sub: sub}, nil
}

// WatchChatRoomClosed is a free log subscription operation binding the contract event 0x228b0adf58c434f90f0f19d95b5d5d4fc677b4d6b1b20c8d1f4ade69f9058de6.
//
// Solidity: event ChatRoomClosed()
func (_ChatRoom *ChatRoomFilterer) WatchChatRoomClosed(opts *bind.WatchOpts, sink chan<- *ChatRoomChatRoomClosed) (event.Subscription, error) {

	logs, sub, err := _ChatRoom.contract.WatchLogs(opts, "ChatRoomClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatRoomChatRoomClosed)
				if err := _ChatRoom.contract.UnpackLog(event, "ChatRoomClosed", log); err != nil {
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
func (_ChatRoom *ChatRoomFilterer) ParseChatRoomClosed(log types.Log) (*ChatRoomChatRoomClosed, error) {
	event := new(ChatRoomChatRoomClosed)
	if err := _ChatRoom.contract.UnpackLog(event, "ChatRoomClosed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChatRoomChatRoomOpenedIterator is returned from FilterChatRoomOpened and is used to iterate over the raw logs and unpacked data for ChatRoomOpened events raised by the ChatRoom contract.
type ChatRoomChatRoomOpenedIterator struct {
	Event *ChatRoomChatRoomOpened // Event containing the contract specifics and raw log

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
func (it *ChatRoomChatRoomOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatRoomChatRoomOpened)
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
		it.Event = new(ChatRoomChatRoomOpened)
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
func (it *ChatRoomChatRoomOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatRoomChatRoomOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatRoomChatRoomOpened represents a ChatRoomOpened event raised by the ChatRoom contract.
type ChatRoomChatRoomOpened struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChatRoomOpened is a free log retrieval operation binding the contract event 0x8879568228d56ae9c174f11f1089ce290d8853217f592250dabf329be26bbd37.
//
// Solidity: event ChatRoomOpened()
func (_ChatRoom *ChatRoomFilterer) FilterChatRoomOpened(opts *bind.FilterOpts) (*ChatRoomChatRoomOpenedIterator, error) {

	logs, sub, err := _ChatRoom.contract.FilterLogs(opts, "ChatRoomOpened")
	if err != nil {
		return nil, err
	}
	return &ChatRoomChatRoomOpenedIterator{contract: _ChatRoom.contract, event: "ChatRoomOpened", logs: logs, sub: sub}, nil
}

// WatchChatRoomOpened is a free log subscription operation binding the contract event 0x8879568228d56ae9c174f11f1089ce290d8853217f592250dabf329be26bbd37.
//
// Solidity: event ChatRoomOpened()
func (_ChatRoom *ChatRoomFilterer) WatchChatRoomOpened(opts *bind.WatchOpts, sink chan<- *ChatRoomChatRoomOpened) (event.Subscription, error) {

	logs, sub, err := _ChatRoom.contract.WatchLogs(opts, "ChatRoomOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatRoomChatRoomOpened)
				if err := _ChatRoom.contract.UnpackLog(event, "ChatRoomOpened", log); err != nil {
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
func (_ChatRoom *ChatRoomFilterer) ParseChatRoomOpened(log types.Log) (*ChatRoomChatRoomOpened, error) {
	event := new(ChatRoomChatRoomOpened)
	if err := _ChatRoom.contract.UnpackLog(event, "ChatRoomOpened", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChatRoomMessageBroadcastedIterator is returned from FilterMessageBroadcasted and is used to iterate over the raw logs and unpacked data for MessageBroadcasted events raised by the ChatRoom contract.
type ChatRoomMessageBroadcastedIterator struct {
	Event *ChatRoomMessageBroadcasted // Event containing the contract specifics and raw log

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
func (it *ChatRoomMessageBroadcastedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatRoomMessageBroadcasted)
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
		it.Event = new(ChatRoomMessageBroadcasted)
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
func (it *ChatRoomMessageBroadcastedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatRoomMessageBroadcastedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatRoomMessageBroadcasted represents a MessageBroadcasted event raised by the ChatRoom contract.
type ChatRoomMessageBroadcasted struct {
	ChannelName string
	DisplayName string
	Text        string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageBroadcasted is a free log retrieval operation binding the contract event 0xef5b6616d0ef7a7345c7d7d30de504c95f981d0f19c9a1559d5e1ee093a1bb7c.
//
// Solidity: event MessageBroadcasted(string channelName, string displayName, string text)
func (_ChatRoom *ChatRoomFilterer) FilterMessageBroadcasted(opts *bind.FilterOpts) (*ChatRoomMessageBroadcastedIterator, error) {

	logs, sub, err := _ChatRoom.contract.FilterLogs(opts, "MessageBroadcasted")
	if err != nil {
		return nil, err
	}
	return &ChatRoomMessageBroadcastedIterator{contract: _ChatRoom.contract, event: "MessageBroadcasted", logs: logs, sub: sub}, nil
}

// WatchMessageBroadcasted is a free log subscription operation binding the contract event 0xef5b6616d0ef7a7345c7d7d30de504c95f981d0f19c9a1559d5e1ee093a1bb7c.
//
// Solidity: event MessageBroadcasted(string channelName, string displayName, string text)
func (_ChatRoom *ChatRoomFilterer) WatchMessageBroadcasted(opts *bind.WatchOpts, sink chan<- *ChatRoomMessageBroadcasted) (event.Subscription, error) {

	logs, sub, err := _ChatRoom.contract.WatchLogs(opts, "MessageBroadcasted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatRoomMessageBroadcasted)
				if err := _ChatRoom.contract.UnpackLog(event, "MessageBroadcasted", log); err != nil {
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

// ParseMessageBroadcasted is a log parse operation binding the contract event 0xef5b6616d0ef7a7345c7d7d30de504c95f981d0f19c9a1559d5e1ee093a1bb7c.
//
// Solidity: event MessageBroadcasted(string channelName, string displayName, string text)
func (_ChatRoom *ChatRoomFilterer) ParseMessageBroadcasted(log types.Log) (*ChatRoomMessageBroadcasted, error) {
	event := new(ChatRoomMessageBroadcasted)
	if err := _ChatRoom.contract.UnpackLog(event, "MessageBroadcasted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChatRoomParticipantJoinedIterator is returned from FilterParticipantJoined and is used to iterate over the raw logs and unpacked data for ParticipantJoined events raised by the ChatRoom contract.
type ChatRoomParticipantJoinedIterator struct {
	Event *ChatRoomParticipantJoined // Event containing the contract specifics and raw log

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
func (it *ChatRoomParticipantJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatRoomParticipantJoined)
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
		it.Event = new(ChatRoomParticipantJoined)
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
func (it *ChatRoomParticipantJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatRoomParticipantJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatRoomParticipantJoined represents a ParticipantJoined event raised by the ChatRoom contract.
type ChatRoomParticipantJoined struct {
	DisplayName string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantJoined is a free log retrieval operation binding the contract event 0xe7ad261bb7ea20d1848a57c9d6084b77f0190124e122bcb41323a554f2009092.
//
// Solidity: event ParticipantJoined(string displayName)
func (_ChatRoom *ChatRoomFilterer) FilterParticipantJoined(opts *bind.FilterOpts) (*ChatRoomParticipantJoinedIterator, error) {

	logs, sub, err := _ChatRoom.contract.FilterLogs(opts, "ParticipantJoined")
	if err != nil {
		return nil, err
	}
	return &ChatRoomParticipantJoinedIterator{contract: _ChatRoom.contract, event: "ParticipantJoined", logs: logs, sub: sub}, nil
}

// WatchParticipantJoined is a free log subscription operation binding the contract event 0xe7ad261bb7ea20d1848a57c9d6084b77f0190124e122bcb41323a554f2009092.
//
// Solidity: event ParticipantJoined(string displayName)
func (_ChatRoom *ChatRoomFilterer) WatchParticipantJoined(opts *bind.WatchOpts, sink chan<- *ChatRoomParticipantJoined) (event.Subscription, error) {

	logs, sub, err := _ChatRoom.contract.WatchLogs(opts, "ParticipantJoined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatRoomParticipantJoined)
				if err := _ChatRoom.contract.UnpackLog(event, "ParticipantJoined", log); err != nil {
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
func (_ChatRoom *ChatRoomFilterer) ParseParticipantJoined(log types.Log) (*ChatRoomParticipantJoined, error) {
	event := new(ChatRoomParticipantJoined)
	if err := _ChatRoom.contract.UnpackLog(event, "ParticipantJoined", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChatRoomParticipantLeftIterator is returned from FilterParticipantLeft and is used to iterate over the raw logs and unpacked data for ParticipantLeft events raised by the ChatRoom contract.
type ChatRoomParticipantLeftIterator struct {
	Event *ChatRoomParticipantLeft // Event containing the contract specifics and raw log

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
func (it *ChatRoomParticipantLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatRoomParticipantLeft)
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
		it.Event = new(ChatRoomParticipantLeft)
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
func (it *ChatRoomParticipantLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatRoomParticipantLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatRoomParticipantLeft represents a ParticipantLeft event raised by the ChatRoom contract.
type ChatRoomParticipantLeft struct {
	DisplayName string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantLeft is a free log retrieval operation binding the contract event 0xd386af6b577d8c71d173442a8cab0a172e24a7a9047ee284da4aa74081a05053.
//
// Solidity: event ParticipantLeft(string displayName)
func (_ChatRoom *ChatRoomFilterer) FilterParticipantLeft(opts *bind.FilterOpts) (*ChatRoomParticipantLeftIterator, error) {

	logs, sub, err := _ChatRoom.contract.FilterLogs(opts, "ParticipantLeft")
	if err != nil {
		return nil, err
	}
	return &ChatRoomParticipantLeftIterator{contract: _ChatRoom.contract, event: "ParticipantLeft", logs: logs, sub: sub}, nil
}

// WatchParticipantLeft is a free log subscription operation binding the contract event 0xd386af6b577d8c71d173442a8cab0a172e24a7a9047ee284da4aa74081a05053.
//
// Solidity: event ParticipantLeft(string displayName)
func (_ChatRoom *ChatRoomFilterer) WatchParticipantLeft(opts *bind.WatchOpts, sink chan<- *ChatRoomParticipantLeft) (event.Subscription, error) {

	logs, sub, err := _ChatRoom.contract.WatchLogs(opts, "ParticipantLeft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatRoomParticipantLeft)
				if err := _ChatRoom.contract.UnpackLog(event, "ParticipantLeft", log); err != nil {
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
func (_ChatRoom *ChatRoomFilterer) ParseParticipantLeft(log types.Log) (*ChatRoomParticipantLeft, error) {
	event := new(ChatRoomParticipantLeft)
	if err := _ChatRoom.contract.UnpackLog(event, "ParticipantLeft", log); err != nil {
		return nil, err
	}
	return event, nil
}
