// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resonance

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

// ResonanceABI is the input ABI used to generate the binding from.
const ResonanceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"invitationRest\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poolBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"Ismanger\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_resonanceId\",\"type\":\"bytes32\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint256[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"resonanceTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"resonanceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"resonancer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogResonance\",\"type\":\"event\"}]"

// ResonanceBin is the compiled bytecode used for deploying new contracts.
var ResonanceBin = "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034600281905550610ab18061005a6000396000f3fe6080604052600436106100705760003560e01c806396365d441161004e57806396365d4414610121578063ac18de431461014c578063e6db233e1461019d578063faf892ff1461020657610070565b80632d06177a1461007557806351cff8d9146100c657806381c3650f14610117575b600080fd5b34801561008157600080fd5b506100c46004803603602081101561009857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061043f565b005b3480156100d257600080fd5b50610115600480360360208110156100e957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104f2565b005b61011f6105ac565b005b34801561012d57600080fd5b506101366105ae565b6040518082815260200191505060405180910390f35b34801561015857600080fd5b5061019b6004803603602081101561016f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105cd565b005b3480156101a957600080fd5b506101ec600480360360208110156101c057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610681565b604051808215151515815260200191505060405180910390f35b34801561021257600080fd5b50610425600480360360c081101561022957600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561027a57600080fd5b82018360208201111561028c57600080fd5b803590602001918460208302840111640100000000831117156102ae57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561030e57600080fd5b82018360208201111561032057600080fd5b8035906020019184602083028401116401000000008311171561034257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156103a257600080fd5b8201836020820111156103b457600080fd5b803590602001918460208302840111640100000000831117156103d657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506106d7565b604051808215151515815260200191505060405180910390f35b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461049857600080fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461054b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f193505050501580156105a8573d6000803e3d6000fd5b5050565b565b60003073ffffffffffffffffffffffffffffffffffffffff1631905090565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461062657600080fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b600060011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461073657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561077057600080fd5b843073ffffffffffffffffffffffffffffffffffffffff1631101561079457600080fd5b82518451146107a257600080fd5b81518451146107b057600080fd5b6003610823888888604051602001808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140182815260200193505050506040516020818303038152906040528051906020012086868661095e565b60ff16101561089a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f746865207369676e436f756e74206c657373207468656e20330000000000000081525060200191505060405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff166108fc869081150290604051600060405180830381858888f193505050501580156108e0573d6000803e3d6000fd5b507f7241e5669bce2b1644b188c5dba4211619e53a4da45e4680c420d35b64cf9ebe878787604051808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a19695505050505050565b600080600090505b8451811015610a7c57600060018787848151811061098057fe5b602002602001015187858151811061099457fe5b60200260200101518786815181106109a857fe5b602002602001015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610a07573d6000803e3d6000fd5b505050602060405103519050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610a6e5782806001019350505b508080600101915050610966565b5094935050505056fea165627a7a723058203aaf767fd0a46013a7e0eb2ce76aaf19cb8cae47e66819d11e88fa868d27338b0029"

// DeployResonance deploys a new Ethereum contract, binding an instance of Resonance to it.
func DeployResonance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Resonance, error) {
	parsed, err := abi.JSON(strings.NewReader(ResonanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ResonanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Resonance{ResonanceCaller: ResonanceCaller{contract: contract}, ResonanceTransactor: ResonanceTransactor{contract: contract}, ResonanceFilterer: ResonanceFilterer{contract: contract}}, nil
}

// Resonance is an auto generated Go binding around an Ethereum contract.
type Resonance struct {
	ResonanceCaller     // Read-only binding to the contract
	ResonanceTransactor // Write-only binding to the contract
	ResonanceFilterer   // Log filterer for contract events
}

// ResonanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResonanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResonanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResonanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResonanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResonanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResonanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResonanceSession struct {
	Contract     *Resonance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResonanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResonanceCallerSession struct {
	Contract *ResonanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ResonanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResonanceTransactorSession struct {
	Contract     *ResonanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ResonanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResonanceRaw struct {
	Contract *Resonance // Generic contract binding to access the raw methods on
}

// ResonanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResonanceCallerRaw struct {
	Contract *ResonanceCaller // Generic read-only contract binding to access the raw methods on
}

// ResonanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResonanceTransactorRaw struct {
	Contract *ResonanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResonance creates a new instance of Resonance, bound to a specific deployed contract.
func NewResonance(address common.Address, backend bind.ContractBackend) (*Resonance, error) {
	contract, err := bindResonance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Resonance{ResonanceCaller: ResonanceCaller{contract: contract}, ResonanceTransactor: ResonanceTransactor{contract: contract}, ResonanceFilterer: ResonanceFilterer{contract: contract}}, nil
}

// NewResonanceCaller creates a new read-only instance of Resonance, bound to a specific deployed contract.
func NewResonanceCaller(address common.Address, caller bind.ContractCaller) (*ResonanceCaller, error) {
	contract, err := bindResonance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResonanceCaller{contract: contract}, nil
}

// NewResonanceTransactor creates a new write-only instance of Resonance, bound to a specific deployed contract.
func NewResonanceTransactor(address common.Address, transactor bind.ContractTransactor) (*ResonanceTransactor, error) {
	contract, err := bindResonance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResonanceTransactor{contract: contract}, nil
}

// NewResonanceFilterer creates a new log filterer instance of Resonance, bound to a specific deployed contract.
func NewResonanceFilterer(address common.Address, filterer bind.ContractFilterer) (*ResonanceFilterer, error) {
	contract, err := bindResonance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResonanceFilterer{contract: contract}, nil
}

// bindResonance binds a generic wrapper to an already deployed contract.
func bindResonance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResonanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resonance *ResonanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Resonance.Contract.ResonanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resonance *ResonanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resonance.Contract.ResonanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resonance *ResonanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resonance.Contract.ResonanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resonance *ResonanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Resonance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resonance *ResonanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resonance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resonance *ResonanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resonance.Contract.contract.Transact(opts, method, params...)
}

// Ismanger is a free data retrieval call binding the contract method 0xe6db233e.
//
// Solidity: function Ismanger(address _manager) constant returns(bool)
func (_Resonance *ResonanceCaller) Ismanger(opts *bind.CallOpts, _manager common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Resonance.contract.Call(opts, out, "Ismanger", _manager)
	return *ret0, err
}

// Ismanger is a free data retrieval call binding the contract method 0xe6db233e.
//
// Solidity: function Ismanger(address _manager) constant returns(bool)
func (_Resonance *ResonanceSession) Ismanger(_manager common.Address) (bool, error) {
	return _Resonance.Contract.Ismanger(&_Resonance.CallOpts, _manager)
}

// Ismanger is a free data retrieval call binding the contract method 0xe6db233e.
//
// Solidity: function Ismanger(address _manager) constant returns(bool)
func (_Resonance *ResonanceCallerSession) Ismanger(_manager common.Address) (bool, error) {
	return _Resonance.Contract.Ismanger(&_Resonance.CallOpts, _manager)
}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() constant returns(uint256)
func (_Resonance *ResonanceCaller) PoolBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Resonance.contract.Call(opts, out, "poolBalance")
	return *ret0, err
}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() constant returns(uint256)
func (_Resonance *ResonanceSession) PoolBalance() (*big.Int, error) {
	return _Resonance.Contract.PoolBalance(&_Resonance.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() constant returns(uint256)
func (_Resonance *ResonanceCallerSession) PoolBalance() (*big.Int, error) {
	return _Resonance.Contract.PoolBalance(&_Resonance.CallOpts)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _manager) returns()
func (_Resonance *ResonanceTransactor) AddManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Resonance.contract.Transact(opts, "addManager", _manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _manager) returns()
func (_Resonance *ResonanceSession) AddManager(_manager common.Address) (*types.Transaction, error) {
	return _Resonance.Contract.AddManager(&_Resonance.TransactOpts, _manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _manager) returns()
func (_Resonance *ResonanceTransactorSession) AddManager(_manager common.Address) (*types.Transaction, error) {
	return _Resonance.Contract.AddManager(&_Resonance.TransactOpts, _manager)
}

// InvitationRest is a paid mutator transaction binding the contract method 0x81c3650f.
//
// Solidity: function invitationRest() returns()
func (_Resonance *ResonanceTransactor) InvitationRest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resonance.contract.Transact(opts, "invitationRest")
}

// InvitationRest is a paid mutator transaction binding the contract method 0x81c3650f.
//
// Solidity: function invitationRest() returns()
func (_Resonance *ResonanceSession) InvitationRest() (*types.Transaction, error) {
	return _Resonance.Contract.InvitationRest(&_Resonance.TransactOpts)
}

// InvitationRest is a paid mutator transaction binding the contract method 0x81c3650f.
//
// Solidity: function invitationRest() returns()
func (_Resonance *ResonanceTransactorSession) InvitationRest() (*types.Transaction, error) {
	return _Resonance.Contract.InvitationRest(&_Resonance.TransactOpts)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address _manager) returns()
func (_Resonance *ResonanceTransactor) RemoveManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Resonance.contract.Transact(opts, "removeManager", _manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address _manager) returns()
func (_Resonance *ResonanceSession) RemoveManager(_manager common.Address) (*types.Transaction, error) {
	return _Resonance.Contract.RemoveManager(&_Resonance.TransactOpts, _manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address _manager) returns()
func (_Resonance *ResonanceTransactorSession) RemoveManager(_manager common.Address) (*types.Transaction, error) {
	return _Resonance.Contract.RemoveManager(&_Resonance.TransactOpts, _manager)
}

// ResonanceTrade is a paid mutator transaction binding the contract method 0xfaf892ff.
//
// Solidity: function resonanceTrade(bytes32 _resonanceId, address _receiver, uint256 _amount, uint256[] v, bytes32[] r, bytes32[] s) returns(bool)
func (_Resonance *ResonanceTransactor) ResonanceTrade(opts *bind.TransactOpts, _resonanceId [32]byte, _receiver common.Address, _amount *big.Int, v []*big.Int, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Resonance.contract.Transact(opts, "resonanceTrade", _resonanceId, _receiver, _amount, v, r, s)
}

// ResonanceTrade is a paid mutator transaction binding the contract method 0xfaf892ff.
//
// Solidity: function resonanceTrade(bytes32 _resonanceId, address _receiver, uint256 _amount, uint256[] v, bytes32[] r, bytes32[] s) returns(bool)
func (_Resonance *ResonanceSession) ResonanceTrade(_resonanceId [32]byte, _receiver common.Address, _amount *big.Int, v []*big.Int, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Resonance.Contract.ResonanceTrade(&_Resonance.TransactOpts, _resonanceId, _receiver, _amount, v, r, s)
}

// ResonanceTrade is a paid mutator transaction binding the contract method 0xfaf892ff.
//
// Solidity: function resonanceTrade(bytes32 _resonanceId, address _receiver, uint256 _amount, uint256[] v, bytes32[] r, bytes32[] s) returns(bool)
func (_Resonance *ResonanceTransactorSession) ResonanceTrade(_resonanceId [32]byte, _receiver common.Address, _amount *big.Int, v []*big.Int, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Resonance.Contract.ResonanceTrade(&_Resonance.TransactOpts, _resonanceId, _receiver, _amount, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _addr) returns()
func (_Resonance *ResonanceTransactor) Withdraw(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Resonance.contract.Transact(opts, "withdraw", _addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _addr) returns()
func (_Resonance *ResonanceSession) Withdraw(_addr common.Address) (*types.Transaction, error) {
	return _Resonance.Contract.Withdraw(&_Resonance.TransactOpts, _addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _addr) returns()
func (_Resonance *ResonanceTransactorSession) Withdraw(_addr common.Address) (*types.Transaction, error) {
	return _Resonance.Contract.Withdraw(&_Resonance.TransactOpts, _addr)
}

// ResonanceLogResonanceIterator is returned from FilterLogResonance and is used to iterate over the raw logs and unpacked data for LogResonance events raised by the Resonance contract.
type ResonanceLogResonanceIterator struct {
	Event *ResonanceLogResonance // Event containing the contract specifics and raw log

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
func (it *ResonanceLogResonanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResonanceLogResonance)
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
		it.Event = new(ResonanceLogResonance)
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
func (it *ResonanceLogResonanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResonanceLogResonanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResonanceLogResonance represents a LogResonance event raised by the Resonance contract.
type ResonanceLogResonance struct {
	ResonanceId [32]byte
	Resonancer  common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogResonance is a free log retrieval operation binding the contract event 0x7241e5669bce2b1644b188c5dba4211619e53a4da45e4680c420d35b64cf9ebe.
//
// Solidity: event LogResonance(bytes32 resonanceId, address resonancer, uint256 amount)
func (_Resonance *ResonanceFilterer) FilterLogResonance(opts *bind.FilterOpts) (*ResonanceLogResonanceIterator, error) {

	logs, sub, err := _Resonance.contract.FilterLogs(opts, "LogResonance")
	if err != nil {
		return nil, err
	}
	return &ResonanceLogResonanceIterator{contract: _Resonance.contract, event: "LogResonance", logs: logs, sub: sub}, nil
}

// WatchLogResonance is a free log subscription operation binding the contract event 0x7241e5669bce2b1644b188c5dba4211619e53a4da45e4680c420d35b64cf9ebe.
//
// Solidity: event LogResonance(bytes32 resonanceId, address resonancer, uint256 amount)
func (_Resonance *ResonanceFilterer) WatchLogResonance(opts *bind.WatchOpts, sink chan<- *ResonanceLogResonance) (event.Subscription, error) {

	logs, sub, err := _Resonance.contract.WatchLogs(opts, "LogResonance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResonanceLogResonance)
				if err := _Resonance.contract.UnpackLog(event, "LogResonance", log); err != nil {
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

// ParseLogResonance is a log parse operation binding the contract event 0x7241e5669bce2b1644b188c5dba4211619e53a4da45e4680c420d35b64cf9ebe.
//
// Solidity: event LogResonance(bytes32 resonanceId, address resonancer, uint256 amount)
func (_Resonance *ResonanceFilterer) ParseLogResonance(log types.Log) (*ResonanceLogResonance, error) {
	event := new(ResonanceLogResonance)
	if err := _Resonance.contract.UnpackLog(event, "LogResonance", log); err != nil {
		return nil, err
	}
	return event, nil
}
