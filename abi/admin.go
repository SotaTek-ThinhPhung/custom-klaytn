// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AdminABI is the input ABI used to generate the binding from.
const AdminABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"BlockUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"UnblockUser\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addOperatorAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"blockUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"changeSuperAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"checkOperatorAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSuperAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"isBlocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeOperatorAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"unblockUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdminBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const AdminBinRuntime = `608060405234801561001057600080fd5b50600436106100935760003560e01c806341858c4b1161006657806341858c4b1461010f5780638204c3261461012257806386c58d3e14610142578063ca56107014610155578063ca9e7d281461016857600080fd5b8063138557e71461009857806339d319e5146100ad5780633da01325146100c0578063416ae768146100d3575b600080fd5b6100ab6100a63660046105b8565b6101a3565b005b6100ab6100bb3660046105b8565b6101fe565b6100ab6100ce3660046105b8565b610285565b6100fc6100e13660046105b8565b6001600160a01b031660009081526001602052604090205490565b6040519081526020015b60405180910390f35b6100ab61011d3660046105b8565b610433565b61012a61048d565b6040516001600160a01b039091168152602001610106565b6100fc6101503660046105da565b6104c7565b6100ab6101633660046105b8565b610547565b6101936101763660046105b8565b6001600160a01b0316600090815260026020526040902054151590565b6040519015158152602001610106565b6101ab61048d565b6001600160a01b0316336001600160a01b0316146101e45760405162461bcd60e51b81526004016101db9061060d565b60405180910390fd5b6001600160a01b0316600090815260026020526040812055565b33600090815260026020526040812054900361022c5760405162461bcd60e51b81526004016101db9061060d565b6001600160a01b0381166000818152600160209081526040808320929092558151338152908101929092527fdbefd7bcfd1cd7894a6b203734a0c53c5400672e7ac3c9ae6e5affa20462b865910160405180910390a150565b3360009081526002602052604081205490036102b35760405162461bcd60e51b81526004016101db9061060d565b806001600160a01b03811661030a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f7420736574207a65726f206164647265737300000000000000000060448201526064016101db565b61031261048d565b6001600160a01b0316816001600160a01b0316036103725760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74207365742073757065722061646d696e2061646472657373000060448201526064016101db565b6001600160a01b038116610600036103d85760405162461bcd60e51b815260206004820152602360248201527f43616e6e6f74207365742067656e6573697320636f6e7472616374206164647260448201526265737360e81b60648201526084016101db565b6001600160a01b038216600081815260016020818152604092839020919091558151338152908101929092527f2c3259f69c8a35d86e1e8a3d898e65d7b5b884c3b4a39ec7d84ac20c62b46e49910160405180910390a15050565b61043b61048d565b6001600160a01b0316336001600160a01b03161461046b5760405162461bcd60e51b81526004016101db9061060d565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600080546001600160a01b03166104b7575073348ec4f32236e24265cd497e3636a5362c3cba4490565b506000546001600160a01b031690565b60006104d161048d565b6001600160a01b0316826001600160a01b0316036104f157506000610541565b6001600160a01b0383166000908152600160205260409020541561051757506001610541565b6001600160a01b0382166000908152600160205260409020541561053d57506002610541565b5060005b92915050565b61054f61048d565b6001600160a01b0316336001600160a01b03161461057f5760405162461bcd60e51b81526004016101db9061060d565b6001600160a01b0316600090815260026020526040902060019055565b80356001600160a01b03811681146105b357600080fd5b919050565b6000602082840312156105ca57600080fd5b6105d38261059c565b9392505050565b600080604083850312156105ed57600080fd5b6105f68361059c565b91506106046020840161059c565b90509250929050565b6020808252601690820152752237903737ba103430bb32903832b936b4b9b9b4b7b760511b60408201526060019056fea26469706673582212206bc6732af7c0f524377227cc383032fd27f0b1694f4522c49f96435a596d181164736f6c634300080d0033`

// AdminFuncSigs maps the 4-byte function signature to its string representation.
var AdminFuncSigs = map[string]string{
	"ca561070": "addOperatorAdmin(address)",
	"3da01325": "blockUser(address)",
	"41858c4b": "changeSuperAdmin(address)",
	"ca9e7d28": "checkOperatorAdmin(address)",
	"8204c326": "getSuperAdmin()",
	"416ae768": "getUserState(address)",
	"86c58d3e": "isBlocked(address,address)",
	"138557e7": "revokeOperatorAdmin(address)",
	"39d319e5": "unblockUser(address)",
}

// AdminBin is the compiled bytecode used for deploying new contracts.
var AdminBin = "0x608060405234801561001057600080fd5b50610673806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806341858c4b1161006657806341858c4b1461010f5780638204c3261461012257806386c58d3e14610142578063ca56107014610155578063ca9e7d281461016857600080fd5b8063138557e71461009857806339d319e5146100ad5780633da01325146100c0578063416ae768146100d3575b600080fd5b6100ab6100a63660046105b8565b6101a3565b005b6100ab6100bb3660046105b8565b6101fe565b6100ab6100ce3660046105b8565b610285565b6100fc6100e13660046105b8565b6001600160a01b031660009081526001602052604090205490565b6040519081526020015b60405180910390f35b6100ab61011d3660046105b8565b610433565b61012a61048d565b6040516001600160a01b039091168152602001610106565b6100fc6101503660046105da565b6104c7565b6100ab6101633660046105b8565b610547565b6101936101763660046105b8565b6001600160a01b0316600090815260026020526040902054151590565b6040519015158152602001610106565b6101ab61048d565b6001600160a01b0316336001600160a01b0316146101e45760405162461bcd60e51b81526004016101db9061060d565b60405180910390fd5b6001600160a01b0316600090815260026020526040812055565b33600090815260026020526040812054900361022c5760405162461bcd60e51b81526004016101db9061060d565b6001600160a01b0381166000818152600160209081526040808320929092558151338152908101929092527fdbefd7bcfd1cd7894a6b203734a0c53c5400672e7ac3c9ae6e5affa20462b865910160405180910390a150565b3360009081526002602052604081205490036102b35760405162461bcd60e51b81526004016101db9061060d565b806001600160a01b03811661030a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f7420736574207a65726f206164647265737300000000000000000060448201526064016101db565b61031261048d565b6001600160a01b0316816001600160a01b0316036103725760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74207365742073757065722061646d696e2061646472657373000060448201526064016101db565b6001600160a01b038116610600036103d85760405162461bcd60e51b815260206004820152602360248201527f43616e6e6f74207365742067656e6573697320636f6e7472616374206164647260448201526265737360e81b60648201526084016101db565b6001600160a01b038216600081815260016020818152604092839020919091558151338152908101929092527f2c3259f69c8a35d86e1e8a3d898e65d7b5b884c3b4a39ec7d84ac20c62b46e49910160405180910390a15050565b61043b61048d565b6001600160a01b0316336001600160a01b03161461046b5760405162461bcd60e51b81526004016101db9061060d565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600080546001600160a01b03166104b7575073348ec4f32236e24265cd497e3636a5362c3cba4490565b506000546001600160a01b031690565b60006104d161048d565b6001600160a01b0316826001600160a01b0316036104f157506000610541565b6001600160a01b0383166000908152600160205260409020541561051757506001610541565b6001600160a01b0382166000908152600160205260409020541561053d57506002610541565b5060005b92915050565b61054f61048d565b6001600160a01b0316336001600160a01b03161461057f5760405162461bcd60e51b81526004016101db9061060d565b6001600160a01b0316600090815260026020526040902060019055565b80356001600160a01b03811681146105b357600080fd5b919050565b6000602082840312156105ca57600080fd5b6105d38261059c565b9392505050565b600080604083850312156105ed57600080fd5b6105f68361059c565b91506106046020840161059c565b90509250929050565b6020808252601690820152752237903737ba103430bb32903832b936b4b9b9b4b7b760511b60408201526060019056fea26469706673582212206bc6732af7c0f524377227cc383032fd27f0b1694f4522c49f96435a596d181164736f6c634300080d0033"

// DeployAdmin deploys a new Klaytn contract, binding an instance of Admin to it.
func DeployAdmin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Admin, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Admin{AdminCaller: AdminCaller{contract: contract}, AdminTransactor: AdminTransactor{contract: contract}, AdminFilterer: AdminFilterer{contract: contract}}, nil
}

// Admin is an auto generated Go binding around a Klaytn contract.
type Admin struct {
	AdminCaller     // Read-only binding to the contract
	AdminTransactor // Write-only binding to the contract
	AdminFilterer   // Log filterer for contract events
}

// AdminCaller is an auto generated read-only Go binding around a Klaytn contract.
type AdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminTransactor is an auto generated write-only Go binding around a Klaytn contract.
type AdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type AdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type AdminSession struct {
	Contract     *Admin            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type AdminCallerSession struct {
	Contract *AdminCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AdminTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type AdminTransactorSession struct {
	Contract     *AdminTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminRaw is an auto generated low-level Go binding around a Klaytn contract.
type AdminRaw struct {
	Contract *Admin // Generic contract binding to access the raw methods on
}

// AdminCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type AdminCallerRaw struct {
	Contract *AdminCaller // Generic read-only contract binding to access the raw methods on
}

// AdminTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type AdminTransactorRaw struct {
	Contract *AdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdmin creates a new instance of Admin, bound to a specific deployed contract.
func NewAdmin(address common.Address, backend bind.ContractBackend) (*Admin, error) {
	contract, err := bindAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Admin{AdminCaller: AdminCaller{contract: contract}, AdminTransactor: AdminTransactor{contract: contract}, AdminFilterer: AdminFilterer{contract: contract}}, nil
}

// NewAdminCaller creates a new read-only instance of Admin, bound to a specific deployed contract.
func NewAdminCaller(address common.Address, caller bind.ContractCaller) (*AdminCaller, error) {
	contract, err := bindAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminCaller{contract: contract}, nil
}

// NewAdminTransactor creates a new write-only instance of Admin, bound to a specific deployed contract.
func NewAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminTransactor, error) {
	contract, err := bindAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminTransactor{contract: contract}, nil
}

// NewAdminFilterer creates a new log filterer instance of Admin, bound to a specific deployed contract.
func NewAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminFilterer, error) {
	contract, err := bindAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminFilterer{contract: contract}, nil
}

// bindAdmin binds a generic wrapper to an already deployed contract.
func bindAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.AdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transact(opts, method, params...)
}

// CheckOperatorAdmin is a free data retrieval call binding the contract method 0xca9e7d28.
//
// Solidity: function checkOperatorAdmin(address _address) view returns(bool)
func (_Admin *AdminCaller) CheckOperatorAdmin(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Admin.contract.Call(opts, out, "checkOperatorAdmin", _address)
	return *ret0, err
}

// CheckOperatorAdmin is a free data retrieval call binding the contract method 0xca9e7d28.
//
// Solidity: function checkOperatorAdmin(address _address) view returns(bool)
func (_Admin *AdminSession) CheckOperatorAdmin(_address common.Address) (bool, error) {
	return _Admin.Contract.CheckOperatorAdmin(&_Admin.CallOpts, _address)
}

// CheckOperatorAdmin is a free data retrieval call binding the contract method 0xca9e7d28.
//
// Solidity: function checkOperatorAdmin(address _address) view returns(bool)
func (_Admin *AdminCallerSession) CheckOperatorAdmin(_address common.Address) (bool, error) {
	return _Admin.Contract.CheckOperatorAdmin(&_Admin.CallOpts, _address)
}

// GetSuperAdmin is a free data retrieval call binding the contract method 0x8204c326.
//
// Solidity: function getSuperAdmin() view returns(address)
func (_Admin *AdminCaller) GetSuperAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Admin.contract.Call(opts, out, "getSuperAdmin")
	return *ret0, err
}

// GetSuperAdmin is a free data retrieval call binding the contract method 0x8204c326.
//
// Solidity: function getSuperAdmin() view returns(address)
func (_Admin *AdminSession) GetSuperAdmin() (common.Address, error) {
	return _Admin.Contract.GetSuperAdmin(&_Admin.CallOpts)
}

// GetSuperAdmin is a free data retrieval call binding the contract method 0x8204c326.
//
// Solidity: function getSuperAdmin() view returns(address)
func (_Admin *AdminCallerSession) GetSuperAdmin() (common.Address, error) {
	return _Admin.Contract.GetSuperAdmin(&_Admin.CallOpts)
}

// GetUserState is a free data retrieval call binding the contract method 0x416ae768.
//
// Solidity: function getUserState(address _address) view returns(uint256)
func (_Admin *AdminCaller) GetUserState(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Admin.contract.Call(opts, out, "getUserState", _address)
	return *ret0, err
}

// GetUserState is a free data retrieval call binding the contract method 0x416ae768.
//
// Solidity: function getUserState(address _address) view returns(uint256)
func (_Admin *AdminSession) GetUserState(_address common.Address) (*big.Int, error) {
	return _Admin.Contract.GetUserState(&_Admin.CallOpts, _address)
}

// GetUserState is a free data retrieval call binding the contract method 0x416ae768.
//
// Solidity: function getUserState(address _address) view returns(uint256)
func (_Admin *AdminCallerSession) GetUserState(_address common.Address) (*big.Int, error) {
	return _Admin.Contract.GetUserState(&_Admin.CallOpts, _address)
}

// IsBlocked is a free data retrieval call binding the contract method 0x86c58d3e.
//
// Solidity: function isBlocked(address from, address to) view returns(uint256)
func (_Admin *AdminCaller) IsBlocked(opts *bind.CallOpts, from common.Address, to common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Admin.contract.Call(opts, out, "isBlocked", from, to)
	return *ret0, err
}

// IsBlocked is a free data retrieval call binding the contract method 0x86c58d3e.
//
// Solidity: function isBlocked(address from, address to) view returns(uint256)
func (_Admin *AdminSession) IsBlocked(from common.Address, to common.Address) (*big.Int, error) {
	return _Admin.Contract.IsBlocked(&_Admin.CallOpts, from, to)
}

// IsBlocked is a free data retrieval call binding the contract method 0x86c58d3e.
//
// Solidity: function isBlocked(address from, address to) view returns(uint256)
func (_Admin *AdminCallerSession) IsBlocked(from common.Address, to common.Address) (*big.Int, error) {
	return _Admin.Contract.IsBlocked(&_Admin.CallOpts, from, to)
}

// AddOperatorAdmin is a paid mutator transaction binding the contract method 0xca561070.
//
// Solidity: function addOperatorAdmin(address _address) returns()
func (_Admin *AdminTransactor) AddOperatorAdmin(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "addOperatorAdmin", _address)
}

// AddOperatorAdmin is a paid mutator transaction binding the contract method 0xca561070.
//
// Solidity: function addOperatorAdmin(address _address) returns()
func (_Admin *AdminSession) AddOperatorAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddOperatorAdmin(&_Admin.TransactOpts, _address)
}

// AddOperatorAdmin is a paid mutator transaction binding the contract method 0xca561070.
//
// Solidity: function addOperatorAdmin(address _address) returns()
func (_Admin *AdminTransactorSession) AddOperatorAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.AddOperatorAdmin(&_Admin.TransactOpts, _address)
}

// BlockUser is a paid mutator transaction binding the contract method 0x3da01325.
//
// Solidity: function blockUser(address target) returns()
func (_Admin *AdminTransactor) BlockUser(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "blockUser", target)
}

// BlockUser is a paid mutator transaction binding the contract method 0x3da01325.
//
// Solidity: function blockUser(address target) returns()
func (_Admin *AdminSession) BlockUser(target common.Address) (*types.Transaction, error) {
	return _Admin.Contract.BlockUser(&_Admin.TransactOpts, target)
}

// BlockUser is a paid mutator transaction binding the contract method 0x3da01325.
//
// Solidity: function blockUser(address target) returns()
func (_Admin *AdminTransactorSession) BlockUser(target common.Address) (*types.Transaction, error) {
	return _Admin.Contract.BlockUser(&_Admin.TransactOpts, target)
}

// ChangeSuperAdmin is a paid mutator transaction binding the contract method 0x41858c4b.
//
// Solidity: function changeSuperAdmin(address _address) returns()
func (_Admin *AdminTransactor) ChangeSuperAdmin(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "changeSuperAdmin", _address)
}

// ChangeSuperAdmin is a paid mutator transaction binding the contract method 0x41858c4b.
//
// Solidity: function changeSuperAdmin(address _address) returns()
func (_Admin *AdminSession) ChangeSuperAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.ChangeSuperAdmin(&_Admin.TransactOpts, _address)
}

// ChangeSuperAdmin is a paid mutator transaction binding the contract method 0x41858c4b.
//
// Solidity: function changeSuperAdmin(address _address) returns()
func (_Admin *AdminTransactorSession) ChangeSuperAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.ChangeSuperAdmin(&_Admin.TransactOpts, _address)
}

// RevokeOperatorAdmin is a paid mutator transaction binding the contract method 0x138557e7.
//
// Solidity: function revokeOperatorAdmin(address _address) returns()
func (_Admin *AdminTransactor) RevokeOperatorAdmin(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "revokeOperatorAdmin", _address)
}

// RevokeOperatorAdmin is a paid mutator transaction binding the contract method 0x138557e7.
//
// Solidity: function revokeOperatorAdmin(address _address) returns()
func (_Admin *AdminSession) RevokeOperatorAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RevokeOperatorAdmin(&_Admin.TransactOpts, _address)
}

// RevokeOperatorAdmin is a paid mutator transaction binding the contract method 0x138557e7.
//
// Solidity: function revokeOperatorAdmin(address _address) returns()
func (_Admin *AdminTransactorSession) RevokeOperatorAdmin(_address common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RevokeOperatorAdmin(&_Admin.TransactOpts, _address)
}

// UnblockUser is a paid mutator transaction binding the contract method 0x39d319e5.
//
// Solidity: function unblockUser(address target) returns()
func (_Admin *AdminTransactor) UnblockUser(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "unblockUser", target)
}

// UnblockUser is a paid mutator transaction binding the contract method 0x39d319e5.
//
// Solidity: function unblockUser(address target) returns()
func (_Admin *AdminSession) UnblockUser(target common.Address) (*types.Transaction, error) {
	return _Admin.Contract.UnblockUser(&_Admin.TransactOpts, target)
}

// UnblockUser is a paid mutator transaction binding the contract method 0x39d319e5.
//
// Solidity: function unblockUser(address target) returns()
func (_Admin *AdminTransactorSession) UnblockUser(target common.Address) (*types.Transaction, error) {
	return _Admin.Contract.UnblockUser(&_Admin.TransactOpts, target)
}

// AdminBlockUserIterator is returned from FilterBlockUser and is used to iterate over the raw logs and unpacked data for BlockUser events raised by the Admin contract.
type AdminBlockUserIterator struct {
	Event *AdminBlockUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminBlockUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminBlockUser)
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
		it.Event = new(AdminBlockUser)
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
func (it *AdminBlockUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminBlockUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminBlockUser represents a BlockUser event raised by the Admin contract.
type AdminBlockUser struct {
	Caller common.Address
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBlockUser is a free log retrieval operation binding the contract event 0x2c3259f69c8a35d86e1e8a3d898e65d7b5b884c3b4a39ec7d84ac20c62b46e49.
//
// Solidity: event BlockUser(address caller, address target)
func (_Admin *AdminFilterer) FilterBlockUser(opts *bind.FilterOpts) (*AdminBlockUserIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "BlockUser")
	if err != nil {
		return nil, err
	}
	return &AdminBlockUserIterator{contract: _Admin.contract, event: "BlockUser", logs: logs, sub: sub}, nil
}

// WatchBlockUser is a free log subscription operation binding the contract event 0x2c3259f69c8a35d86e1e8a3d898e65d7b5b884c3b4a39ec7d84ac20c62b46e49.
//
// Solidity: event BlockUser(address caller, address target)
func (_Admin *AdminFilterer) WatchBlockUser(opts *bind.WatchOpts, sink chan<- *AdminBlockUser) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "BlockUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminBlockUser)
				if err := _Admin.contract.UnpackLog(event, "BlockUser", log); err != nil {
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

// ParseBlockUser is a log parse operation binding the contract event 0x2c3259f69c8a35d86e1e8a3d898e65d7b5b884c3b4a39ec7d84ac20c62b46e49.
//
// Solidity: event BlockUser(address caller, address target)
func (_Admin *AdminFilterer) ParseBlockUser(log types.Log) (*AdminBlockUser, error) {
	event := new(AdminBlockUser)
	if err := _Admin.contract.UnpackLog(event, "BlockUser", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminUnblockUserIterator is returned from FilterUnblockUser and is used to iterate over the raw logs and unpacked data for UnblockUser events raised by the Admin contract.
type AdminUnblockUserIterator struct {
	Event *AdminUnblockUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminUnblockUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminUnblockUser)
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
		it.Event = new(AdminUnblockUser)
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
func (it *AdminUnblockUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminUnblockUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminUnblockUser represents a UnblockUser event raised by the Admin contract.
type AdminUnblockUser struct {
	Caller common.Address
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnblockUser is a free log retrieval operation binding the contract event 0xdbefd7bcfd1cd7894a6b203734a0c53c5400672e7ac3c9ae6e5affa20462b865.
//
// Solidity: event UnblockUser(address caller, address target)
func (_Admin *AdminFilterer) FilterUnblockUser(opts *bind.FilterOpts) (*AdminUnblockUserIterator, error) {

	logs, sub, err := _Admin.contract.FilterLogs(opts, "UnblockUser")
	if err != nil {
		return nil, err
	}
	return &AdminUnblockUserIterator{contract: _Admin.contract, event: "UnblockUser", logs: logs, sub: sub}, nil
}

// WatchUnblockUser is a free log subscription operation binding the contract event 0xdbefd7bcfd1cd7894a6b203734a0c53c5400672e7ac3c9ae6e5affa20462b865.
//
// Solidity: event UnblockUser(address caller, address target)
func (_Admin *AdminFilterer) WatchUnblockUser(opts *bind.WatchOpts, sink chan<- *AdminUnblockUser) (event.Subscription, error) {

	logs, sub, err := _Admin.contract.WatchLogs(opts, "UnblockUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminUnblockUser)
				if err := _Admin.contract.UnpackLog(event, "UnblockUser", log); err != nil {
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

// ParseUnblockUser is a log parse operation binding the contract event 0xdbefd7bcfd1cd7894a6b203734a0c53c5400672e7ac3c9ae6e5affa20462b865.
//
// Solidity: event UnblockUser(address caller, address target)
func (_Admin *AdminFilterer) ParseUnblockUser(log types.Log) (*AdminUnblockUser, error) {
	event := new(AdminUnblockUser)
	if err := _Admin.contract.UnpackLog(event, "UnblockUser", log); err != nil {
		return nil, err
	}
	return event, nil
}
