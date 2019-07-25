// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package btcResonance

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

// BtcResonanceABI is the input ABI used to generate the binding from.
const BtcResonanceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"registResonanceNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"resonanceTrades\",\"outputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"resonanced\",\"type\":\"bool\"},{\"name\":\"signcount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"invitationRest\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_resonanceId\",\"type\":\"bytes32\"}],\"name\":\"approveResonance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"resonanceTrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PoolAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogResonance\",\"type\":\"event\"}]"

// BtcResonanceBin is the compiled bytecode used for deploying new contracts.
var BtcResonanceBin = "0x60806040526040516020806108768339810180604052602081101561002357600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600181905550506107eb8061008b6000396000f3fe6080604052600436106100555760003560e01c806307583f051461005a5780632851ff9f1461006457806381c3650f146100f8578063c47f19e514610102578063d4e533581461017b578063fd67c324146101d6575b600080fd5b61006261022d565b005b34801561007057600080fd5b5061009d6004803603602081101561008757600080fd5b81019080803590602001909291905050506102ea565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018315151515815260200182815260200194505050505060405180910390f35b610100610347565b005b34801561010e57600080fd5b506101656004803603606081101561012557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610349565b6040518082815260200191505060405180910390f35b34801561018757600080fd5b506101d46004803603604081101561019e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506106da565b005b3480156101e257600080fd5b506101eb610773565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61271034116102a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f746865206465706f736974206973206e6f7420656e6f7567680000000000000081525060200191505060405180910390fd5b34600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900460ff16908060030154905084565b565b6000336000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156103e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806107996027913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166003600085815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561053a5760405180608001604052808673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200160001515815260200160018152506003600085815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160030155905050600360008481526020019081526020016000206003015491506106d2565b60016003600085815260200190815260200160002060030154016003600085815260200190815260200160002060030181905550600380600085815260200190815260200160002060030154101580156105bb5750600015156003600085815260200190815260200160002060020160009054906101000a900460ff161515145b156106d15760016003600085815260200190815260200160002060020160006101000a81548160ff0219169083151502179055506003600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc60036000868152602001908152602001600020600101549081150290604051600060405180830381858888f19350505050158015610681573d6000803e3d6000fd5b508473ffffffffffffffffffffffffffffffffffffffff167fc59ee75c48b4d11cb884f7c91453ef4147c4d1d1a4bc801c6aaec50db398f2bb856040518082815260200191505060405180910390a25b5b509392505050565b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610720573d6000803e3d6000fd5b508173ffffffffffffffffffffffffffffffffffffffff167fc59ee75c48b4d11cb884f7c91453ef4147c4d1d1a4bc801c6aaec50db398f2bb826040518082815260200191505060405180910390a25050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fe4e6f646561646472657373206973206e6f74207265676573745265736f6e616e63654e6f646521a165627a7a7230582006669e6dd8d1d48f59fbdfb5096634986d974a1d7fd7f305543ef92c45e938280029"

// DeployBtcResonance deploys a new Ethereum contract, binding an instance of BtcResonance to it.
func DeployBtcResonance(auth *bind.TransactOpts, backend bind.ContractBackend, amount *big.Int) (common.Address, *types.Transaction, *BtcResonance, error) {
	parsed, err := abi.JSON(strings.NewReader(BtcResonanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BtcResonanceBin), backend, amount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BtcResonance{BtcResonanceCaller: BtcResonanceCaller{contract: contract}, BtcResonanceTransactor: BtcResonanceTransactor{contract: contract}, BtcResonanceFilterer: BtcResonanceFilterer{contract: contract}}, nil
}

// BtcResonance is an auto generated Go binding around an Ethereum contract.
type BtcResonance struct {
	BtcResonanceCaller     // Read-only binding to the contract
	BtcResonanceTransactor // Write-only binding to the contract
	BtcResonanceFilterer   // Log filterer for contract events
}

// BtcResonanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BtcResonanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcResonanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BtcResonanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcResonanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BtcResonanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcResonanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BtcResonanceSession struct {
	Contract     *BtcResonance     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtcResonanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BtcResonanceCallerSession struct {
	Contract *BtcResonanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BtcResonanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BtcResonanceTransactorSession struct {
	Contract     *BtcResonanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BtcResonanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BtcResonanceRaw struct {
	Contract *BtcResonance // Generic contract binding to access the raw methods on
}

// BtcResonanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BtcResonanceCallerRaw struct {
	Contract *BtcResonanceCaller // Generic read-only contract binding to access the raw methods on
}

// BtcResonanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BtcResonanceTransactorRaw struct {
	Contract *BtcResonanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBtcResonance creates a new instance of BtcResonance, bound to a specific deployed contract.
func NewBtcResonance(address common.Address, backend bind.ContractBackend) (*BtcResonance, error) {
	contract, err := bindBtcResonance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BtcResonance{BtcResonanceCaller: BtcResonanceCaller{contract: contract}, BtcResonanceTransactor: BtcResonanceTransactor{contract: contract}, BtcResonanceFilterer: BtcResonanceFilterer{contract: contract}}, nil
}

// NewBtcResonanceCaller creates a new read-only instance of BtcResonance, bound to a specific deployed contract.
func NewBtcResonanceCaller(address common.Address, caller bind.ContractCaller) (*BtcResonanceCaller, error) {
	contract, err := bindBtcResonance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BtcResonanceCaller{contract: contract}, nil
}

// NewBtcResonanceTransactor creates a new write-only instance of BtcResonance, bound to a specific deployed contract.
func NewBtcResonanceTransactor(address common.Address, transactor bind.ContractTransactor) (*BtcResonanceTransactor, error) {
	contract, err := bindBtcResonance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BtcResonanceTransactor{contract: contract}, nil
}

// NewBtcResonanceFilterer creates a new log filterer instance of BtcResonance, bound to a specific deployed contract.
func NewBtcResonanceFilterer(address common.Address, filterer bind.ContractFilterer) (*BtcResonanceFilterer, error) {
	contract, err := bindBtcResonance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BtcResonanceFilterer{contract: contract}, nil
}

// bindBtcResonance binds a generic wrapper to an already deployed contract.
func bindBtcResonance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BtcResonanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BtcResonance *BtcResonanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BtcResonance.Contract.BtcResonanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BtcResonance *BtcResonanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtcResonance.Contract.BtcResonanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BtcResonance *BtcResonanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BtcResonance.Contract.BtcResonanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BtcResonance *BtcResonanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BtcResonance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BtcResonance *BtcResonanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtcResonance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BtcResonance *BtcResonanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BtcResonance.Contract.contract.Transact(opts, method, params...)
}

// PoolAddr is a free data retrieval call binding the contract method 0xfd67c324.
//
// Solidity: function PoolAddr() constant returns(address)
func (_BtcResonance *BtcResonanceCaller) PoolAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BtcResonance.contract.Call(opts, out, "PoolAddr")
	return *ret0, err
}

// PoolAddr is a free data retrieval call binding the contract method 0xfd67c324.
//
// Solidity: function PoolAddr() constant returns(address)
func (_BtcResonance *BtcResonanceSession) PoolAddr() (common.Address, error) {
	return _BtcResonance.Contract.PoolAddr(&_BtcResonance.CallOpts)
}

// PoolAddr is a free data retrieval call binding the contract method 0xfd67c324.
//
// Solidity: function PoolAddr() constant returns(address)
func (_BtcResonance *BtcResonanceCallerSession) PoolAddr() (common.Address, error) {
	return _BtcResonance.Contract.PoolAddr(&_BtcResonance.CallOpts)
}

// ResonanceTrades is a free data retrieval call binding the contract method 0x2851ff9f.
//
// Solidity: function resonanceTrades(bytes32 ) constant returns(address receiver, uint256 amount, bool resonanced, uint256 signcount)
func (_BtcResonance *BtcResonanceCaller) ResonanceTrades(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver   common.Address
	Amount     *big.Int
	Resonanced bool
	Signcount  *big.Int
}, error) {
	ret := new(struct {
		Receiver   common.Address
		Amount     *big.Int
		Resonanced bool
		Signcount  *big.Int
	})
	out := ret
	err := _BtcResonance.contract.Call(opts, out, "resonanceTrades", arg0)
	return *ret, err
}

// ResonanceTrades is a free data retrieval call binding the contract method 0x2851ff9f.
//
// Solidity: function resonanceTrades(bytes32 ) constant returns(address receiver, uint256 amount, bool resonanced, uint256 signcount)
func (_BtcResonance *BtcResonanceSession) ResonanceTrades(arg0 [32]byte) (struct {
	Receiver   common.Address
	Amount     *big.Int
	Resonanced bool
	Signcount  *big.Int
}, error) {
	return _BtcResonance.Contract.ResonanceTrades(&_BtcResonance.CallOpts, arg0)
}

// ResonanceTrades is a free data retrieval call binding the contract method 0x2851ff9f.
//
// Solidity: function resonanceTrades(bytes32 ) constant returns(address receiver, uint256 amount, bool resonanced, uint256 signcount)
func (_BtcResonance *BtcResonanceCallerSession) ResonanceTrades(arg0 [32]byte) (struct {
	Receiver   common.Address
	Amount     *big.Int
	Resonanced bool
	Signcount  *big.Int
}, error) {
	return _BtcResonance.Contract.ResonanceTrades(&_BtcResonance.CallOpts, arg0)
}

// ApproveResonance is a paid mutator transaction binding the contract method 0xc47f19e5.
//
// Solidity: function approveResonance(address _receiver, uint256 _amount, bytes32 _resonanceId) returns(uint256)
func (_BtcResonance *BtcResonanceTransactor) ApproveResonance(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int, _resonanceId [32]byte) (*types.Transaction, error) {
	return _BtcResonance.contract.Transact(opts, "approveResonance", _receiver, _amount, _resonanceId)
}

// ApproveResonance is a paid mutator transaction binding the contract method 0xc47f19e5.
//
// Solidity: function approveResonance(address _receiver, uint256 _amount, bytes32 _resonanceId) returns(uint256)
func (_BtcResonance *BtcResonanceSession) ApproveResonance(_receiver common.Address, _amount *big.Int, _resonanceId [32]byte) (*types.Transaction, error) {
	return _BtcResonance.Contract.ApproveResonance(&_BtcResonance.TransactOpts, _receiver, _amount, _resonanceId)
}

// ApproveResonance is a paid mutator transaction binding the contract method 0xc47f19e5.
//
// Solidity: function approveResonance(address _receiver, uint256 _amount, bytes32 _resonanceId) returns(uint256)
func (_BtcResonance *BtcResonanceTransactorSession) ApproveResonance(_receiver common.Address, _amount *big.Int, _resonanceId [32]byte) (*types.Transaction, error) {
	return _BtcResonance.Contract.ApproveResonance(&_BtcResonance.TransactOpts, _receiver, _amount, _resonanceId)
}

// InvitationRest is a paid mutator transaction binding the contract method 0x81c3650f.
//
// Solidity: function invitationRest() returns()
func (_BtcResonance *BtcResonanceTransactor) InvitationRest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtcResonance.contract.Transact(opts, "invitationRest")
}

// InvitationRest is a paid mutator transaction binding the contract method 0x81c3650f.
//
// Solidity: function invitationRest() returns()
func (_BtcResonance *BtcResonanceSession) InvitationRest() (*types.Transaction, error) {
	return _BtcResonance.Contract.InvitationRest(&_BtcResonance.TransactOpts)
}

// InvitationRest is a paid mutator transaction binding the contract method 0x81c3650f.
//
// Solidity: function invitationRest() returns()
func (_BtcResonance *BtcResonanceTransactorSession) InvitationRest() (*types.Transaction, error) {
	return _BtcResonance.Contract.InvitationRest(&_BtcResonance.TransactOpts)
}

// RegistResonanceNode is a paid mutator transaction binding the contract method 0x07583f05.
//
// Solidity: function registResonanceNode() returns()
func (_BtcResonance *BtcResonanceTransactor) RegistResonanceNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtcResonance.contract.Transact(opts, "registResonanceNode")
}

// RegistResonanceNode is a paid mutator transaction binding the contract method 0x07583f05.
//
// Solidity: function registResonanceNode() returns()
func (_BtcResonance *BtcResonanceSession) RegistResonanceNode() (*types.Transaction, error) {
	return _BtcResonance.Contract.RegistResonanceNode(&_BtcResonance.TransactOpts)
}

// RegistResonanceNode is a paid mutator transaction binding the contract method 0x07583f05.
//
// Solidity: function registResonanceNode() returns()
func (_BtcResonance *BtcResonanceTransactorSession) RegistResonanceNode() (*types.Transaction, error) {
	return _BtcResonance.Contract.RegistResonanceNode(&_BtcResonance.TransactOpts)
}

// ResonanceTrade is a paid mutator transaction binding the contract method 0xd4e53358.
//
// Solidity: function resonanceTrade(address _receiver, uint256 _amount) returns()
func (_BtcResonance *BtcResonanceTransactor) ResonanceTrade(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BtcResonance.contract.Transact(opts, "resonanceTrade", _receiver, _amount)
}

// ResonanceTrade is a paid mutator transaction binding the contract method 0xd4e53358.
//
// Solidity: function resonanceTrade(address _receiver, uint256 _amount) returns()
func (_BtcResonance *BtcResonanceSession) ResonanceTrade(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BtcResonance.Contract.ResonanceTrade(&_BtcResonance.TransactOpts, _receiver, _amount)
}

// ResonanceTrade is a paid mutator transaction binding the contract method 0xd4e53358.
//
// Solidity: function resonanceTrade(address _receiver, uint256 _amount) returns()
func (_BtcResonance *BtcResonanceTransactorSession) ResonanceTrade(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BtcResonance.Contract.ResonanceTrade(&_BtcResonance.TransactOpts, _receiver, _amount)
}

// BtcResonanceLogResonanceIterator is returned from FilterLogResonance and is used to iterate over the raw logs and unpacked data for LogResonance events raised by the BtcResonance contract.
type BtcResonanceLogResonanceIterator struct {
	Event *BtcResonanceLogResonance // Event containing the contract specifics and raw log

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
func (it *BtcResonanceLogResonanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcResonanceLogResonance)
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
		it.Event = new(BtcResonanceLogResonance)
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
func (it *BtcResonanceLogResonanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcResonanceLogResonanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcResonanceLogResonance represents a LogResonance event raised by the BtcResonance contract.
type BtcResonanceLogResonance struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogResonance is a free log retrieval operation binding the contract event 0xc59ee75c48b4d11cb884f7c91453ef4147c4d1d1a4bc801c6aaec50db398f2bb.
//
// Solidity: event LogResonance(address indexed receiver, uint256 amount)
func (_BtcResonance *BtcResonanceFilterer) FilterLogResonance(opts *bind.FilterOpts, receiver []common.Address) (*BtcResonanceLogResonanceIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BtcResonance.contract.FilterLogs(opts, "LogResonance", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BtcResonanceLogResonanceIterator{contract: _BtcResonance.contract, event: "LogResonance", logs: logs, sub: sub}, nil
}

// WatchLogResonance is a free log subscription operation binding the contract event 0xc59ee75c48b4d11cb884f7c91453ef4147c4d1d1a4bc801c6aaec50db398f2bb.
//
// Solidity: event LogResonance(address indexed receiver, uint256 amount)
func (_BtcResonance *BtcResonanceFilterer) WatchLogResonance(opts *bind.WatchOpts, sink chan<- *BtcResonanceLogResonance, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BtcResonance.contract.WatchLogs(opts, "LogResonance", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcResonanceLogResonance)
				if err := _BtcResonance.contract.UnpackLog(event, "LogResonance", log); err != nil {
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

// ParseLogResonance is a log parse operation binding the contract event 0xc59ee75c48b4d11cb884f7c91453ef4147c4d1d1a4bc801c6aaec50db398f2bb.
//
// Solidity: event LogResonance(address indexed receiver, uint256 amount)
func (_BtcResonance *BtcResonanceFilterer) ParseLogResonance(log types.Log) (*BtcResonanceLogResonance, error) {
	event := new(BtcResonanceLogResonance)
	if err := _BtcResonance.contract.UnpackLog(event, "LogResonance", log); err != nil {
		return nil, err
	}
	return event, nil
}
