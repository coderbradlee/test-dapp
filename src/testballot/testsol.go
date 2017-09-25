// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BallotABI is the input ABI used to generate the binding from.
const BallotABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"proposalIndex\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"voteCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testtxorigin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winningProposal\",\"outputs\":[{\"name\":\"winningProposalIndex\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testvalue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\"},{\"name\":\"voted\",\"type\":\"bool\"},{\"name\":\"delegate\",\"type\":\"address\"},{\"name\":\"vote\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testdata\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testmsg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testgas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winnerName\",\"outputs\":[{\"name\":\"winnerNames\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testGasPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BallotBin is the compiled bytecode used for deploying new contracts.
const BallotBin = `0x6060604052341561000f57600080fd5b6040516108103803806108108339810160405280805160008054600160a060020a03191633600160a060020a039081169190911780835516815260016020819052604082205592019190505b81518110156100d057600280546001810161007683826100d7565b91600052602060002090600202016000604080519081016040528086868151811061009d57fe5b906020019060200201518152600060209091015291905081518155602082015160019182015592909201915061005b9050565b505061012f565b815481835581811511610103576002028160020283600052602060002091820191016101039190610108565b505050565b61012c91905b80821115610128576000808255600182015560020161010e565b5090565b90565b6106d28061013e6000396000f300606060405236156100cd5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630121b93f81146100d2578063013cf08b146100ea5780632e4176cf146101185780634d66a3ab146101475780635c19a95c1461015a578063609ff1bd1461017957806399244d1a1461019e5780639e7b8d61146101b1578063a3ec138d146101d0578063b449c68914610224578063be8f316a146102ae578063c8e9fc8b146102c1578063e2ba53f0146102d4578063e9de432f146102e7575b600080fd5b34156100dd57600080fd5b6100e86004356102fa565b005b34156100f557600080fd5b61010060043561036d565b60405191825260208201526040908101905180910390f35b341561012357600080fd5b61012b610399565b604051600160a060020a03909116815260200160405180910390f35b341561015257600080fd5b61012b6103a8565b341561016557600080fd5b6100e8600160a060020a03600435166103ac565b341561018457600080fd5b61018c6104f2565b60405190815260200160405180910390f35b34156101a957600080fd5b61018c61055c565b34156101bc57600080fd5b6100e8600160a060020a0360043516610560565b34156101db57600080fd5b6101ef600160a060020a03600435166105e5565b6040519384529115156020840152600160a060020a031660408084019190915260608301919091526080909101905180910390f35b341561022f57600080fd5b610237610619565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561027357808201518382015260200161025b565b50505050905090810190601f1680156102a05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102b957600080fd5b61012b610657565b34156102cc57600080fd5b61018c61065b565b34156102df57600080fd5b61018c610663565b34156102f257600080fd5b61018c610690565b600160a060020a03331660009081526001602081905260409091209081015460ff161561032657600080fd5b6001818101805460ff191690911790556002808201839055815481549091908490811061034f57fe5b60009182526020909120600160029092020101805490910190555050565b600280548290811061037b57fe5b60009182526020909120600290910201805460019091015490915082565b600054600160a060020a031681565b3290565b600160a060020a033316600090815260016020819052604082209081015490919060ff16156103da57600080fd5b33600160a060020a031683600160a060020a0316141515156103fb57600080fd5b600160a060020a03838116600090815260016020819052604090912001546101009004161561045d57600160a060020a03928316600090815260016020819052604090912001546101009004831692331683141561045857600080fd5b6103fb565b506001818101805460ff1916821774ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a0386169081029190911790915560009081526020829052604090209081015460ff16156104e5578154600282810154815481106104c657fe5b60009182526020909120600160029092020101805490910190556104ed565b815481540181555b505050565b600080805b600254811015610557578160028281548110151561051157fe5b906000526020600020906002020160010154111561054f57600280548290811061053757fe5b90600052602060002090600202016001015491508092505b6001016104f7565b505090565b3490565b60005433600160a060020a03908116911614801561059b5750600160a060020a0381166000908152600160208190526040909120015460ff16155b80156105bd5750600160a060020a038116600090815260016020526040902054155b15156105c857600080fd5b600160a060020a0316600090815260016020819052604090912055565b600160208190526000918252604090912080549181015460029091015460ff8216916101009004600160a060020a03169084565b610621610694565b6000368080601f016020809104026020016040519081016040528181529291906020840183838082843750949550505050505090565b3390565b60005a905090565b6000600261066f6104f2565b8154811061067957fe5b906000526020600020906002020160000154905090565b3a90565b602060405190810160405260008152905600a165627a7a723058204e088e82b02d6ac9d44984cf3f4b63f35d92c306b6bdbe9a392001ceb56d9ecc0029`

// DeployBallot deploys a new Ethereum contract, binding an instance of Ballot to it.
func DeployBallot(auth *bind.TransactOpts, backend bind.ContractBackend, proposalNames [][32]byte) (common.Address, *types.Transaction, *Ballot, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BallotBin), backend, proposalNames)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ballot{BallotCaller: BallotCaller{contract: contract}, BallotTransactor: BallotTransactor{contract: contract}}, nil
}

// Ballot is an auto generated Go binding around an Ethereum contract.
type Ballot struct {
	BallotCaller     // Read-only binding to the contract
	BallotTransactor // Write-only binding to the contract
}

// BallotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotSession struct {
	Contract     *Ballot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotCallerSession struct {
	Contract *BallotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BallotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotTransactorSession struct {
	Contract     *BallotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotRaw struct {
	Contract *Ballot // Generic contract binding to access the raw methods on
}

// BallotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotCallerRaw struct {
	Contract *BallotCaller // Generic read-only contract binding to access the raw methods on
}

// BallotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotTransactorRaw struct {
	Contract *BallotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallot creates a new instance of Ballot, bound to a specific deployed contract.
func NewBallot(address common.Address, backend bind.ContractBackend) (*Ballot, error) {
	contract, err := bindBallot(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ballot{BallotCaller: BallotCaller{contract: contract}, BallotTransactor: BallotTransactor{contract: contract}}, nil
}

// NewBallotCaller creates a new read-only instance of Ballot, bound to a specific deployed contract.
func NewBallotCaller(address common.Address, caller bind.ContractCaller) (*BallotCaller, error) {
	contract, err := bindBallot(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BallotCaller{contract: contract}, nil
}

// NewBallotTransactor creates a new write-only instance of Ballot, bound to a specific deployed contract.
func NewBallotTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotTransactor, error) {
	contract, err := bindBallot(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BallotTransactor{contract: contract}, nil
}

// bindBallot binds a generic wrapper to an already deployed contract.
func bindBallot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.BallotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Ballot *BallotCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ballot.contract.Call(opts, out, "chairperson")
	return *ret0, err
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Ballot *BallotSession) Chairperson() (common.Address, error) {
	return _Ballot.Contract.Chairperson(&_Ballot.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() constant returns(address)
func (_Ballot *BallotCallerSession) Chairperson() (common.Address, error) {
	return _Ballot.Contract.Chairperson(&_Ballot.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals( uint256) constant returns(name bytes32, voteCount uint256)
func (_Ballot *BallotCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	ret := new(struct {
		Name      [32]byte
		VoteCount *big.Int
	})
	out := ret
	err := _Ballot.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals( uint256) constant returns(name bytes32, voteCount uint256)
func (_Ballot *BallotSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Ballot.Contract.Proposals(&_Ballot.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals( uint256) constant returns(name bytes32, voteCount uint256)
func (_Ballot *BallotCallerSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Ballot.Contract.Proposals(&_Ballot.CallOpts, arg0)
}

// TestGasPrice is a free data retrieval call binding the contract method 0xe9de432f.
//
// Solidity: function testGasPrice() constant returns(uint256)
func (_Ballot *BallotCaller) TestGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ballot.contract.Call(opts, out, "testGasPrice")
	return *ret0, err
}

// TestGasPrice is a free data retrieval call binding the contract method 0xe9de432f.
//
// Solidity: function testGasPrice() constant returns(uint256)
func (_Ballot *BallotSession) TestGasPrice() (*big.Int, error) {
	return _Ballot.Contract.TestGasPrice(&_Ballot.CallOpts)
}

// TestGasPrice is a free data retrieval call binding the contract method 0xe9de432f.
//
// Solidity: function testGasPrice() constant returns(uint256)
func (_Ballot *BallotCallerSession) TestGasPrice() (*big.Int, error) {
	return _Ballot.Contract.TestGasPrice(&_Ballot.CallOpts)
}

// Testdata is a free data retrieval call binding the contract method 0xb449c689.
//
// Solidity: function testdata() constant returns(bytes)
func (_Ballot *BallotCaller) Testdata(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Ballot.contract.Call(opts, out, "testdata")
	return *ret0, err
}

// Testdata is a free data retrieval call binding the contract method 0xb449c689.
//
// Solidity: function testdata() constant returns(bytes)
func (_Ballot *BallotSession) Testdata() ([]byte, error) {
	return _Ballot.Contract.Testdata(&_Ballot.CallOpts)
}

// Testdata is a free data retrieval call binding the contract method 0xb449c689.
//
// Solidity: function testdata() constant returns(bytes)
func (_Ballot *BallotCallerSession) Testdata() ([]byte, error) {
	return _Ballot.Contract.Testdata(&_Ballot.CallOpts)
}

// Testgas is a free data retrieval call binding the contract method 0xc8e9fc8b.
//
// Solidity: function testgas() constant returns(uint256)
func (_Ballot *BallotCaller) Testgas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ballot.contract.Call(opts, out, "testgas")
	return *ret0, err
}

// Testgas is a free data retrieval call binding the contract method 0xc8e9fc8b.
//
// Solidity: function testgas() constant returns(uint256)
func (_Ballot *BallotSession) Testgas() (*big.Int, error) {
	return _Ballot.Contract.Testgas(&_Ballot.CallOpts)
}

// Testgas is a free data retrieval call binding the contract method 0xc8e9fc8b.
//
// Solidity: function testgas() constant returns(uint256)
func (_Ballot *BallotCallerSession) Testgas() (*big.Int, error) {
	return _Ballot.Contract.Testgas(&_Ballot.CallOpts)
}

// Testmsg is a free data retrieval call binding the contract method 0xbe8f316a.
//
// Solidity: function testmsg() constant returns(address)
func (_Ballot *BallotCaller) Testmsg(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ballot.contract.Call(opts, out, "testmsg")
	return *ret0, err
}

// Testmsg is a free data retrieval call binding the contract method 0xbe8f316a.
//
// Solidity: function testmsg() constant returns(address)
func (_Ballot *BallotSession) Testmsg() (common.Address, error) {
	return _Ballot.Contract.Testmsg(&_Ballot.CallOpts)
}

// Testmsg is a free data retrieval call binding the contract method 0xbe8f316a.
//
// Solidity: function testmsg() constant returns(address)
func (_Ballot *BallotCallerSession) Testmsg() (common.Address, error) {
	return _Ballot.Contract.Testmsg(&_Ballot.CallOpts)
}

// Testtxorigin is a free data retrieval call binding the contract method 0x4d66a3ab.
//
// Solidity: function testtxorigin() constant returns(address)
func (_Ballot *BallotCaller) Testtxorigin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ballot.contract.Call(opts, out, "testtxorigin")
	return *ret0, err
}

// Testtxorigin is a free data retrieval call binding the contract method 0x4d66a3ab.
//
// Solidity: function testtxorigin() constant returns(address)
func (_Ballot *BallotSession) Testtxorigin() (common.Address, error) {
	return _Ballot.Contract.Testtxorigin(&_Ballot.CallOpts)
}

// Testtxorigin is a free data retrieval call binding the contract method 0x4d66a3ab.
//
// Solidity: function testtxorigin() constant returns(address)
func (_Ballot *BallotCallerSession) Testtxorigin() (common.Address, error) {
	return _Ballot.Contract.Testtxorigin(&_Ballot.CallOpts)
}

// Testvalue is a free data retrieval call binding the contract method 0x99244d1a.
//
// Solidity: function testvalue() constant returns(uint256)
func (_Ballot *BallotCaller) Testvalue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ballot.contract.Call(opts, out, "testvalue")
	return *ret0, err
}

// Testvalue is a free data retrieval call binding the contract method 0x99244d1a.
//
// Solidity: function testvalue() constant returns(uint256)
func (_Ballot *BallotSession) Testvalue() (*big.Int, error) {
	return _Ballot.Contract.Testvalue(&_Ballot.CallOpts)
}

// Testvalue is a free data retrieval call binding the contract method 0x99244d1a.
//
// Solidity: function testvalue() constant returns(uint256)
func (_Ballot *BallotCallerSession) Testvalue() (*big.Int, error) {
	return _Ballot.Contract.Testvalue(&_Ballot.CallOpts)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters( address) constant returns(weight uint256, voted bool, delegate address, vote uint256)
func (_Ballot *BallotCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	ret := new(struct {
		Weight   *big.Int
		Voted    bool
		Delegate common.Address
		Vote     *big.Int
	})
	out := ret
	err := _Ballot.contract.Call(opts, out, "voters", arg0)
	return *ret, err
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters( address) constant returns(weight uint256, voted bool, delegate address, vote uint256)
func (_Ballot *BallotSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Ballot.Contract.Voters(&_Ballot.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters( address) constant returns(weight uint256, voted bool, delegate address, vote uint256)
func (_Ballot *BallotCallerSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Ballot.Contract.Voters(&_Ballot.CallOpts, arg0)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() constant returns(winnerNames bytes32)
func (_Ballot *BallotCaller) WinnerName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Ballot.contract.Call(opts, out, "winnerName")
	return *ret0, err
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() constant returns(winnerNames bytes32)
func (_Ballot *BallotSession) WinnerName() ([32]byte, error) {
	return _Ballot.Contract.WinnerName(&_Ballot.CallOpts)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() constant returns(winnerNames bytes32)
func (_Ballot *BallotCallerSession) WinnerName() ([32]byte, error) {
	return _Ballot.Contract.WinnerName(&_Ballot.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() constant returns(winningProposalIndex uint256)
func (_Ballot *BallotCaller) WinningProposal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ballot.contract.Call(opts, out, "winningProposal")
	return *ret0, err
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() constant returns(winningProposalIndex uint256)
func (_Ballot *BallotSession) WinningProposal() (*big.Int, error) {
	return _Ballot.Contract.WinningProposal(&_Ballot.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() constant returns(winningProposalIndex uint256)
func (_Ballot *BallotCallerSession) WinningProposal() (*big.Int, error) {
	return _Ballot.Contract.WinningProposal(&_Ballot.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(to address) returns()
func (_Ballot *BallotTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(to address) returns()
func (_Ballot *BallotSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.Delegate(&_Ballot.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(to address) returns()
func (_Ballot *BallotTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.Delegate(&_Ballot.TransactOpts, to)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(voter address) returns()
func (_Ballot *BallotTransactor) GiveRightToVote(opts *bind.TransactOpts, voter common.Address) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "giveRightToVote", voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(voter address) returns()
func (_Ballot *BallotSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.GiveRightToVote(&_Ballot.TransactOpts, voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(voter address) returns()
func (_Ballot *BallotTransactorSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.GiveRightToVote(&_Ballot.TransactOpts, voter)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(proposalIndex uint256) returns()
func (_Ballot *BallotTransactor) Vote(opts *bind.TransactOpts, proposalIndex *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "vote", proposalIndex)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(proposalIndex uint256) returns()
func (_Ballot *BallotSession) Vote(proposalIndex *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, proposalIndex)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(proposalIndex uint256) returns()
func (_Ballot *BallotTransactorSession) Vote(proposalIndex *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, proposalIndex)
}
