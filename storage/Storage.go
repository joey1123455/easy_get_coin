// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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

// GameHistoryGameSession is an auto generated low-level Go binding around an user-defined struct.
type GameHistoryGameSession struct {
	Gid  *big.Int
	Gtid string
	Uid  string
	Data string
	Time *big.Int
}

// GameHistoryPayment is an auto generated low-level Go binding around an user-defined struct.
type GameHistoryPayment struct {
	Sender common.Address
	Amount *big.Int
	Time   *big.Int
}

// GameHistoryMetaData contains all meta data concerning the GameHistory contract.
var GameHistoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedLessThanTarget\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gid\",\"type\":\"uint256\"}],\"name\":\"getGameHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gtid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"internalType\":\"structGameHistory.GameSession[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uid\",\"type\":\"string\"}],\"name\":\"getUserHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gtid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"internalType\":\"structGameHistory.GameSession[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_gtid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"storeGameData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddressEGC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userStakeHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"internalType\":\"structGameHistory.Payment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// GameHistoryABI is the input ABI used to generate the binding from.
// Deprecated: Use GameHistoryMetaData.ABI instead.
var GameHistoryABI = GameHistoryMetaData.ABI

// GameHistory is an auto generated Go binding around an Ethereum contract.
type GameHistory struct {
	GameHistoryCaller     // Read-only binding to the contract
	GameHistoryTransactor // Write-only binding to the contract
	GameHistoryFilterer   // Log filterer for contract events
}

// GameHistoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameHistoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameHistoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameHistoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameHistoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameHistoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameHistorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameHistorySession struct {
	Contract     *GameHistory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameHistoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameHistoryCallerSession struct {
	Contract *GameHistoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GameHistoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameHistoryTransactorSession struct {
	Contract     *GameHistoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GameHistoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameHistoryRaw struct {
	Contract *GameHistory // Generic contract binding to access the raw methods on
}

// GameHistoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameHistoryCallerRaw struct {
	Contract *GameHistoryCaller // Generic read-only contract binding to access the raw methods on
}

// GameHistoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameHistoryTransactorRaw struct {
	Contract *GameHistoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameHistory creates a new instance of GameHistory, bound to a specific deployed contract.
func NewGameHistory(address common.Address, backend bind.ContractBackend) (*GameHistory, error) {
	contract, err := bindGameHistory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameHistory{GameHistoryCaller: GameHistoryCaller{contract: contract}, GameHistoryTransactor: GameHistoryTransactor{contract: contract}, GameHistoryFilterer: GameHistoryFilterer{contract: contract}}, nil
}

// NewGameHistoryCaller creates a new read-only instance of GameHistory, bound to a specific deployed contract.
func NewGameHistoryCaller(address common.Address, caller bind.ContractCaller) (*GameHistoryCaller, error) {
	contract, err := bindGameHistory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameHistoryCaller{contract: contract}, nil
}

// NewGameHistoryTransactor creates a new write-only instance of GameHistory, bound to a specific deployed contract.
func NewGameHistoryTransactor(address common.Address, transactor bind.ContractTransactor) (*GameHistoryTransactor, error) {
	contract, err := bindGameHistory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameHistoryTransactor{contract: contract}, nil
}

// NewGameHistoryFilterer creates a new log filterer instance of GameHistory, bound to a specific deployed contract.
func NewGameHistoryFilterer(address common.Address, filterer bind.ContractFilterer) (*GameHistoryFilterer, error) {
	contract, err := bindGameHistory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameHistoryFilterer{contract: contract}, nil
}

// bindGameHistory binds a generic wrapper to an already deployed contract.
func bindGameHistory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GameHistoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameHistory *GameHistoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameHistory.Contract.GameHistoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameHistory *GameHistoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameHistory.Contract.GameHistoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameHistory *GameHistoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameHistory.Contract.GameHistoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameHistory *GameHistoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameHistory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameHistory *GameHistoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameHistory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameHistory *GameHistoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameHistory.Contract.contract.Transact(opts, method, params...)
}

// GetGameHistory is a free data retrieval call binding the contract method 0xbcc412da.
//
// Solidity: function getGameHistory(uint256 _gid) view returns((uint256,string,string,string,uint256)[])
func (_GameHistory *GameHistoryCaller) GetGameHistory(opts *bind.CallOpts, _gid *big.Int) ([]GameHistoryGameSession, error) {
	var out []interface{}
	err := _GameHistory.contract.Call(opts, &out, "getGameHistory", _gid)

	if err != nil {
		return *new([]GameHistoryGameSession), err
	}

	out0 := *abi.ConvertType(out[0], new([]GameHistoryGameSession)).(*[]GameHistoryGameSession)

	return out0, err

}

// GetGameHistory is a free data retrieval call binding the contract method 0xbcc412da.
//
// Solidity: function getGameHistory(uint256 _gid) view returns((uint256,string,string,string,uint256)[])
func (_GameHistory *GameHistorySession) GetGameHistory(_gid *big.Int) ([]GameHistoryGameSession, error) {
	return _GameHistory.Contract.GetGameHistory(&_GameHistory.CallOpts, _gid)
}

// GetGameHistory is a free data retrieval call binding the contract method 0xbcc412da.
//
// Solidity: function getGameHistory(uint256 _gid) view returns((uint256,string,string,string,uint256)[])
func (_GameHistory *GameHistoryCallerSession) GetGameHistory(_gid *big.Int) ([]GameHistoryGameSession, error) {
	return _GameHistory.Contract.GetGameHistory(&_GameHistory.CallOpts, _gid)
}

// GetUserHistory is a free data retrieval call binding the contract method 0xc6c17135.
//
// Solidity: function getUserHistory(string _uid) view returns((uint256,string,string,string,uint256)[])
func (_GameHistory *GameHistoryCaller) GetUserHistory(opts *bind.CallOpts, _uid string) ([]GameHistoryGameSession, error) {
	var out []interface{}
	err := _GameHistory.contract.Call(opts, &out, "getUserHistory", _uid)

	if err != nil {
		return *new([]GameHistoryGameSession), err
	}

	out0 := *abi.ConvertType(out[0], new([]GameHistoryGameSession)).(*[]GameHistoryGameSession)

	return out0, err

}

// GetUserHistory is a free data retrieval call binding the contract method 0xc6c17135.
//
// Solidity: function getUserHistory(string _uid) view returns((uint256,string,string,string,uint256)[])
func (_GameHistory *GameHistorySession) GetUserHistory(_uid string) ([]GameHistoryGameSession, error) {
	return _GameHistory.Contract.GetUserHistory(&_GameHistory.CallOpts, _uid)
}

// GetUserHistory is a free data retrieval call binding the contract method 0xc6c17135.
//
// Solidity: function getUserHistory(string _uid) view returns((uint256,string,string,string,uint256)[])
func (_GameHistory *GameHistoryCallerSession) GetUserHistory(_uid string) ([]GameHistoryGameSession, error) {
	return _GameHistory.Contract.GetUserHistory(&_GameHistory.CallOpts, _uid)
}

// TokenAddressEGC is a free data retrieval call binding the contract method 0x2578b84f.
//
// Solidity: function tokenAddressEGC() view returns(address)
func (_GameHistory *GameHistoryCaller) TokenAddressEGC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GameHistory.contract.Call(opts, &out, "tokenAddressEGC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddressEGC is a free data retrieval call binding the contract method 0x2578b84f.
//
// Solidity: function tokenAddressEGC() view returns(address)
func (_GameHistory *GameHistorySession) TokenAddressEGC() (common.Address, error) {
	return _GameHistory.Contract.TokenAddressEGC(&_GameHistory.CallOpts)
}

// TokenAddressEGC is a free data retrieval call binding the contract method 0x2578b84f.
//
// Solidity: function tokenAddressEGC() view returns(address)
func (_GameHistory *GameHistoryCallerSession) TokenAddressEGC() (common.Address, error) {
	return _GameHistory.Contract.TokenAddressEGC(&_GameHistory.CallOpts)
}

// UserStakeHistory is a free data retrieval call binding the contract method 0x5a48a1d9.
//
// Solidity: function userStakeHistory(address _user) view returns((address,uint256,uint256)[])
func (_GameHistory *GameHistoryCaller) UserStakeHistory(opts *bind.CallOpts, _user common.Address) ([]GameHistoryPayment, error) {
	var out []interface{}
	err := _GameHistory.contract.Call(opts, &out, "userStakeHistory", _user)

	if err != nil {
		return *new([]GameHistoryPayment), err
	}

	out0 := *abi.ConvertType(out[0], new([]GameHistoryPayment)).(*[]GameHistoryPayment)

	return out0, err

}

// UserStakeHistory is a free data retrieval call binding the contract method 0x5a48a1d9.
//
// Solidity: function userStakeHistory(address _user) view returns((address,uint256,uint256)[])
func (_GameHistory *GameHistorySession) UserStakeHistory(_user common.Address) ([]GameHistoryPayment, error) {
	return _GameHistory.Contract.UserStakeHistory(&_GameHistory.CallOpts, _user)
}

// UserStakeHistory is a free data retrieval call binding the contract method 0x5a48a1d9.
//
// Solidity: function userStakeHistory(address _user) view returns((address,uint256,uint256)[])
func (_GameHistory *GameHistoryCallerSession) UserStakeHistory(_user common.Address) ([]GameHistoryPayment, error) {
	return _GameHistory.Contract.UserStakeHistory(&_GameHistory.CallOpts, _user)
}

// UserTotal is a free data retrieval call binding the contract method 0xe601cf44.
//
// Solidity: function userTotal(address _user) view returns(uint256)
func (_GameHistory *GameHistoryCaller) UserTotal(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GameHistory.contract.Call(opts, &out, "userTotal", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTotal is a free data retrieval call binding the contract method 0xe601cf44.
//
// Solidity: function userTotal(address _user) view returns(uint256)
func (_GameHistory *GameHistorySession) UserTotal(_user common.Address) (*big.Int, error) {
	return _GameHistory.Contract.UserTotal(&_GameHistory.CallOpts, _user)
}

// UserTotal is a free data retrieval call binding the contract method 0xe601cf44.
//
// Solidity: function userTotal(address _user) view returns(uint256)
func (_GameHistory *GameHistoryCallerSession) UserTotal(_user common.Address) (*big.Int, error) {
	return _GameHistory.Contract.UserTotal(&_GameHistory.CallOpts, _user)
}

// StoreGameData is a paid mutator transaction binding the contract method 0x92f4019a.
//
// Solidity: function storeGameData(uint256 _gid, string _gtid, string _uid, string _data, uint256 _time) returns()
func (_GameHistory *GameHistoryTransactor) StoreGameData(opts *bind.TransactOpts, _gid *big.Int, _gtid string, _uid string, _data string, _time *big.Int) (*types.Transaction, error) {
	return _GameHistory.contract.Transact(opts, "storeGameData", _gid, _gtid, _uid, _data, _time)
}

// StoreGameData is a paid mutator transaction binding the contract method 0x92f4019a.
//
// Solidity: function storeGameData(uint256 _gid, string _gtid, string _uid, string _data, uint256 _time) returns()
func (_GameHistory *GameHistorySession) StoreGameData(_gid *big.Int, _gtid string, _uid string, _data string, _time *big.Int) (*types.Transaction, error) {
	return _GameHistory.Contract.StoreGameData(&_GameHistory.TransactOpts, _gid, _gtid, _uid, _data, _time)
}

// StoreGameData is a paid mutator transaction binding the contract method 0x92f4019a.
//
// Solidity: function storeGameData(uint256 _gid, string _gtid, string _uid, string _data, uint256 _time) returns()
func (_GameHistory *GameHistoryTransactorSession) StoreGameData(_gid *big.Int, _gtid string, _uid string, _data string, _time *big.Int) (*types.Transaction, error) {
	return _GameHistory.Contract.StoreGameData(&_GameHistory.TransactOpts, _gid, _gtid, _uid, _data, _time)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GameHistory *GameHistoryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameHistory.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GameHistory *GameHistorySession) Receive() (*types.Transaction, error) {
	return _GameHistory.Contract.Receive(&_GameHistory.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GameHistory *GameHistoryTransactorSession) Receive() (*types.Transaction, error) {
	return _GameHistory.Contract.Receive(&_GameHistory.TransactOpts)
}

// GameHistoryOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the GameHistory contract.
type GameHistoryOwnerSetIterator struct {
	Event *GameHistoryOwnerSet // Event containing the contract specifics and raw log

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
func (it *GameHistoryOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameHistoryOwnerSet)
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
		it.Event = new(GameHistoryOwnerSet)
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
func (it *GameHistoryOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameHistoryOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameHistoryOwnerSet represents a OwnerSet event raised by the GameHistory contract.
type GameHistoryOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_GameHistory *GameHistoryFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*GameHistoryOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameHistory.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameHistoryOwnerSetIterator{contract: _GameHistory.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_GameHistory *GameHistoryFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *GameHistoryOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameHistory.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameHistoryOwnerSet)
				if err := _GameHistory.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// ParseOwnerSet is a log parse operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_GameHistory *GameHistoryFilterer) ParseOwnerSet(log types.Log) (*GameHistoryOwnerSet, error) {
	event := new(GameHistoryOwnerSet)
	if err := _GameHistory.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameHistoryReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the GameHistory contract.
type GameHistoryReceivedIterator struct {
	Event *GameHistoryReceived // Event containing the contract specifics and raw log

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
func (it *GameHistoryReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameHistoryReceived)
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
		it.Event = new(GameHistoryReceived)
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
func (it *GameHistoryReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameHistoryReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameHistoryReceived represents a Received event raised by the GameHistory contract.
type GameHistoryReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 amount)
func (_GameHistory *GameHistoryFilterer) FilterReceived(opts *bind.FilterOpts, sender []common.Address) (*GameHistoryReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameHistory.contract.FilterLogs(opts, "Received", senderRule)
	if err != nil {
		return nil, err
	}
	return &GameHistoryReceivedIterator{contract: _GameHistory.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 amount)
func (_GameHistory *GameHistoryFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *GameHistoryReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameHistory.contract.WatchLogs(opts, "Received", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameHistoryReceived)
				if err := _GameHistory.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 amount)
func (_GameHistory *GameHistoryFilterer) ParseReceived(log types.Log) (*GameHistoryReceived, error) {
	event := new(GameHistoryReceived)
	if err := _GameHistory.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameHistoryReceivedLessThanTargetIterator is returned from FilterReceivedLessThanTarget and is used to iterate over the raw logs and unpacked data for ReceivedLessThanTarget events raised by the GameHistory contract.
type GameHistoryReceivedLessThanTargetIterator struct {
	Event *GameHistoryReceivedLessThanTarget // Event containing the contract specifics and raw log

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
func (it *GameHistoryReceivedLessThanTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameHistoryReceivedLessThanTarget)
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
		it.Event = new(GameHistoryReceivedLessThanTarget)
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
func (it *GameHistoryReceivedLessThanTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameHistoryReceivedLessThanTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameHistoryReceivedLessThanTarget represents a ReceivedLessThanTarget event raised by the GameHistory contract.
type GameHistoryReceivedLessThanTarget struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedLessThanTarget is a free log retrieval operation binding the contract event 0xa5ebd4198835aee14a329cca5bd059673b15ae2a4a66e8c800648b3b4e9fbff8.
//
// Solidity: event ReceivedLessThanTarget(address indexed sender, uint256 amount)
func (_GameHistory *GameHistoryFilterer) FilterReceivedLessThanTarget(opts *bind.FilterOpts, sender []common.Address) (*GameHistoryReceivedLessThanTargetIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameHistory.contract.FilterLogs(opts, "ReceivedLessThanTarget", senderRule)
	if err != nil {
		return nil, err
	}
	return &GameHistoryReceivedLessThanTargetIterator{contract: _GameHistory.contract, event: "ReceivedLessThanTarget", logs: logs, sub: sub}, nil
}

// WatchReceivedLessThanTarget is a free log subscription operation binding the contract event 0xa5ebd4198835aee14a329cca5bd059673b15ae2a4a66e8c800648b3b4e9fbff8.
//
// Solidity: event ReceivedLessThanTarget(address indexed sender, uint256 amount)
func (_GameHistory *GameHistoryFilterer) WatchReceivedLessThanTarget(opts *bind.WatchOpts, sink chan<- *GameHistoryReceivedLessThanTarget, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GameHistory.contract.WatchLogs(opts, "ReceivedLessThanTarget", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameHistoryReceivedLessThanTarget)
				if err := _GameHistory.contract.UnpackLog(event, "ReceivedLessThanTarget", log); err != nil {
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

// ParseReceivedLessThanTarget is a log parse operation binding the contract event 0xa5ebd4198835aee14a329cca5bd059673b15ae2a4a66e8c800648b3b4e9fbff8.
//
// Solidity: event ReceivedLessThanTarget(address indexed sender, uint256 amount)
func (_GameHistory *GameHistoryFilterer) ParseReceivedLessThanTarget(log types.Log) (*GameHistoryReceivedLessThanTarget, error) {
	event := new(GameHistoryReceivedLessThanTarget)
	if err := _GameHistory.contract.UnpackLog(event, "ReceivedLessThanTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
