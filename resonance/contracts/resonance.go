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
const ResonanceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"PoolBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_resonanceId\",\"type\":\"bytes32\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createResonance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"invitationRest\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_resonanceId\",\"type\":\"bytes32\"}],\"name\":\"haveResonance\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_resonanceId\",\"type\":\"bytes32\"}],\"name\":\"sigResonance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"creater\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"resonanceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogResonanceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"resonanceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogResonanced\",\"type\":\"event\"}]"

// ResonanceBin is the compiled bytecode used for deploying new contracts.
var ResonanceBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610b50806100606000396000f3fe6080604052600436106100705760003560e01c806381c3650f1161004e57806381c3650f14610153578063a3a29fa21461015d578063a49c3b01146101b0578063ac18de43146101eb57610070565b80631be0819c146100725780632d06177a1461009d5780637019f029146100ee575b005b34801561007e57600080fd5b5061008761023c565b6040518082815260200191505060405180910390f35b3480156100a957600080fd5b506100ec600480360360208110156100c057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102b8565b005b3480156100fa57600080fd5b506101516004803603606081101561011157600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061036b565b005b61015b61057f565b005b34801561016957600080fd5b506101966004803603602081101561018057600080fd5b8101908080359060200190929190505050610581565b604051808215151515815260200191505060405180910390f35b3480156101bc57600080fd5b506101e9600480360360208110156101d357600080fd5b81019080803590602001909291905050506105f0565b005b3480156101f757600080fd5b5061023a6004803603602081101561020e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a36565b005b600060011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461029b57600080fd5b3073ffffffffffffffffffffffffffffffffffffffff1631905090565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461031157600080fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515146103c857600080fd5b803073ffffffffffffffffffffffffffffffffffffffff163110156103ec57600080fd5b6103f4610aea565b82816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050818160200181815250506000816040019060ff16908160ff1681525050806004600086815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548160ff021916908360ff1602179055509050507f684b666fb285edbb362bc05c378ae46cc7943988d1e3241615281b33ec00f8df33858585604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a150505050565b565b60008073ffffffffffffffffffffffffffffffffffffffff166004600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b60011515600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461064d57600080fd5b8061065781610581565b6106c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f7265736f6e616e6365496420646f6573206e6f7420657869737400000000000081525060200191505060405180910390fd5b6000600460008481526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561073e57600080fd5b8060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561079b57600080fd5b60018160030160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1660ff1614156107fa57600080fd5b60018160030160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff16021790555080600201600081819054906101000a900460ff168092919060010191906101000a81548160ff021916908360ff1602179055505060028160020160009054906101000a900460ff1660ff1610610a315780600101543073ffffffffffffffffffffffffffffffffffffffff163110156108cd57600080fd5b8060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc82600101549081150290604051600060405180830381858888f1935050505015801561093b573d6000803e3d6000fd5b507f2a17f0bf846a9f4ce8fb9dbe21e12a5d2c79bc07d0adf5d41146ac459084692e838260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360010154604051808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a160046000848152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560018201600090556002820160006101000a81549060ff021916905550505b505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a8f57600080fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600060ff168152509056fea165627a7a723058206d64c24d505b7b8cfbc812b39a4e027b2210294bf71acd8612ee4f0985b2c7220029"

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

// PoolBalance is a free data retrieval call binding the contract method 0x1be0819c.
//
// Solidity: function PoolBalance() constant returns(uint256)
func (_Resonance *ResonanceCaller) PoolBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Resonance.contract.Call(opts, out, "PoolBalance")
	return *ret0, err
}

// PoolBalance is a free data retrieval call binding the contract method 0x1be0819c.
//
// Solidity: function PoolBalance() constant returns(uint256)
func (_Resonance *ResonanceSession) PoolBalance() (*big.Int, error) {
	return _Resonance.Contract.PoolBalance(&_Resonance.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0x1be0819c.
//
// Solidity: function PoolBalance() constant returns(uint256)
func (_Resonance *ResonanceCallerSession) PoolBalance() (*big.Int, error) {
	return _Resonance.Contract.PoolBalance(&_Resonance.CallOpts)
}

// HaveResonance is a free data retrieval call binding the contract method 0xa3a29fa2.
//
// Solidity: function haveResonance(bytes32 _resonanceId) constant returns(bool exists)
func (_Resonance *ResonanceCaller) HaveResonance(opts *bind.CallOpts, _resonanceId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Resonance.contract.Call(opts, out, "haveResonance", _resonanceId)
	return *ret0, err
}

// HaveResonance is a free data retrieval call binding the contract method 0xa3a29fa2.
//
// Solidity: function haveResonance(bytes32 _resonanceId) constant returns(bool exists)
func (_Resonance *ResonanceSession) HaveResonance(_resonanceId [32]byte) (bool, error) {
	return _Resonance.Contract.HaveResonance(&_Resonance.CallOpts, _resonanceId)
}

// HaveResonance is a free data retrieval call binding the contract method 0xa3a29fa2.
//
// Solidity: function haveResonance(bytes32 _resonanceId) constant returns(bool exists)
func (_Resonance *ResonanceCallerSession) HaveResonance(_resonanceId [32]byte) (bool, error) {
	return _Resonance.Contract.HaveResonance(&_Resonance.CallOpts, _resonanceId)
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

// CreateResonance is a paid mutator transaction binding the contract method 0x7019f029.
//
// Solidity: function createResonance(bytes32 _resonanceId, address _receiver, uint256 _amount) returns()
func (_Resonance *ResonanceTransactor) CreateResonance(opts *bind.TransactOpts, _resonanceId [32]byte, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Resonance.contract.Transact(opts, "createResonance", _resonanceId, _receiver, _amount)
}

// CreateResonance is a paid mutator transaction binding the contract method 0x7019f029.
//
// Solidity: function createResonance(bytes32 _resonanceId, address _receiver, uint256 _amount) returns()
func (_Resonance *ResonanceSession) CreateResonance(_resonanceId [32]byte, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Resonance.Contract.CreateResonance(&_Resonance.TransactOpts, _resonanceId, _receiver, _amount)
}

// CreateResonance is a paid mutator transaction binding the contract method 0x7019f029.
//
// Solidity: function createResonance(bytes32 _resonanceId, address _receiver, uint256 _amount) returns()
func (_Resonance *ResonanceTransactorSession) CreateResonance(_resonanceId [32]byte, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Resonance.Contract.CreateResonance(&_Resonance.TransactOpts, _resonanceId, _receiver, _amount)
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

// SigResonance is a paid mutator transaction binding the contract method 0xa49c3b01.
//
// Solidity: function sigResonance(bytes32 _resonanceId) returns()
func (_Resonance *ResonanceTransactor) SigResonance(opts *bind.TransactOpts, _resonanceId [32]byte) (*types.Transaction, error) {
	return _Resonance.contract.Transact(opts, "sigResonance", _resonanceId)
}

// SigResonance is a paid mutator transaction binding the contract method 0xa49c3b01.
//
// Solidity: function sigResonance(bytes32 _resonanceId) returns()
func (_Resonance *ResonanceSession) SigResonance(_resonanceId [32]byte) (*types.Transaction, error) {
	return _Resonance.Contract.SigResonance(&_Resonance.TransactOpts, _resonanceId)
}

// SigResonance is a paid mutator transaction binding the contract method 0xa49c3b01.
//
// Solidity: function sigResonance(bytes32 _resonanceId) returns()
func (_Resonance *ResonanceTransactorSession) SigResonance(_resonanceId [32]byte) (*types.Transaction, error) {
	return _Resonance.Contract.SigResonance(&_Resonance.TransactOpts, _resonanceId)
}

// ResonanceLogResonanceCreatedIterator is returned from FilterLogResonanceCreated and is used to iterate over the raw logs and unpacked data for LogResonanceCreated events raised by the Resonance contract.
type ResonanceLogResonanceCreatedIterator struct {
	Event *ResonanceLogResonanceCreated // Event containing the contract specifics and raw log

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
func (it *ResonanceLogResonanceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResonanceLogResonanceCreated)
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
		it.Event = new(ResonanceLogResonanceCreated)
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
func (it *ResonanceLogResonanceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResonanceLogResonanceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResonanceLogResonanceCreated represents a LogResonanceCreated event raised by the Resonance contract.
type ResonanceLogResonanceCreated struct {
	Creater     common.Address
	ResonanceId [32]byte
	Receiver    common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogResonanceCreated is a free log retrieval operation binding the contract event 0x684b666fb285edbb362bc05c378ae46cc7943988d1e3241615281b33ec00f8df.
//
// Solidity: event LogResonanceCreated(address creater, bytes32 resonanceId, address receiver, uint256 amount)
func (_Resonance *ResonanceFilterer) FilterLogResonanceCreated(opts *bind.FilterOpts) (*ResonanceLogResonanceCreatedIterator, error) {

	logs, sub, err := _Resonance.contract.FilterLogs(opts, "LogResonanceCreated")
	if err != nil {
		return nil, err
	}
	return &ResonanceLogResonanceCreatedIterator{contract: _Resonance.contract, event: "LogResonanceCreated", logs: logs, sub: sub}, nil
}

// WatchLogResonanceCreated is a free log subscription operation binding the contract event 0x684b666fb285edbb362bc05c378ae46cc7943988d1e3241615281b33ec00f8df.
//
// Solidity: event LogResonanceCreated(address creater, bytes32 resonanceId, address receiver, uint256 amount)
func (_Resonance *ResonanceFilterer) WatchLogResonanceCreated(opts *bind.WatchOpts, sink chan<- *ResonanceLogResonanceCreated) (event.Subscription, error) {

	logs, sub, err := _Resonance.contract.WatchLogs(opts, "LogResonanceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResonanceLogResonanceCreated)
				if err := _Resonance.contract.UnpackLog(event, "LogResonanceCreated", log); err != nil {
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

// ParseLogResonanceCreated is a log parse operation binding the contract event 0x684b666fb285edbb362bc05c378ae46cc7943988d1e3241615281b33ec00f8df.
//
// Solidity: event LogResonanceCreated(address creater, bytes32 resonanceId, address receiver, uint256 amount)
func (_Resonance *ResonanceFilterer) ParseLogResonanceCreated(log types.Log) (*ResonanceLogResonanceCreated, error) {
	event := new(ResonanceLogResonanceCreated)
	if err := _Resonance.contract.UnpackLog(event, "LogResonanceCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ResonanceLogResonancedIterator is returned from FilterLogResonanced and is used to iterate over the raw logs and unpacked data for LogResonanced events raised by the Resonance contract.
type ResonanceLogResonancedIterator struct {
	Event *ResonanceLogResonanced // Event containing the contract specifics and raw log

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
func (it *ResonanceLogResonancedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResonanceLogResonanced)
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
		it.Event = new(ResonanceLogResonanced)
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
func (it *ResonanceLogResonancedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResonanceLogResonancedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResonanceLogResonanced represents a LogResonanced event raised by the Resonance contract.
type ResonanceLogResonanced struct {
	ResonanceId [32]byte
	Receiver    common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogResonanced is a free log retrieval operation binding the contract event 0x2a17f0bf846a9f4ce8fb9dbe21e12a5d2c79bc07d0adf5d41146ac459084692e.
//
// Solidity: event LogResonanced(bytes32 resonanceId, address receiver, uint256 amount)
func (_Resonance *ResonanceFilterer) FilterLogResonanced(opts *bind.FilterOpts) (*ResonanceLogResonancedIterator, error) {

	logs, sub, err := _Resonance.contract.FilterLogs(opts, "LogResonanced")
	if err != nil {
		return nil, err
	}
	return &ResonanceLogResonancedIterator{contract: _Resonance.contract, event: "LogResonanced", logs: logs, sub: sub}, nil
}

// WatchLogResonanced is a free log subscription operation binding the contract event 0x2a17f0bf846a9f4ce8fb9dbe21e12a5d2c79bc07d0adf5d41146ac459084692e.
//
// Solidity: event LogResonanced(bytes32 resonanceId, address receiver, uint256 amount)
func (_Resonance *ResonanceFilterer) WatchLogResonanced(opts *bind.WatchOpts, sink chan<- *ResonanceLogResonanced) (event.Subscription, error) {

	logs, sub, err := _Resonance.contract.WatchLogs(opts, "LogResonanced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResonanceLogResonanced)
				if err := _Resonance.contract.UnpackLog(event, "LogResonanced", log); err != nil {
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

// ParseLogResonanced is a log parse operation binding the contract event 0x2a17f0bf846a9f4ce8fb9dbe21e12a5d2c79bc07d0adf5d41146ac459084692e.
//
// Solidity: event LogResonanced(bytes32 resonanceId, address receiver, uint256 amount)
func (_Resonance *ResonanceFilterer) ParseLogResonanced(log types.Log) (*ResonanceLogResonanced, error) {
	event := new(ResonanceLogResonanced)
	if err := _Resonance.contract.UnpackLog(event, "LogResonanced", log); err != nil {
		return nil, err
	}
	return event, nil
}
