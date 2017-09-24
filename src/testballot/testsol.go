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
const BallotABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"proposalIndex\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"voteCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winningProposal\",\"outputs\":[{\"name\":\"winningProposalIndex\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testvalue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\"},{\"name\":\"voted\",\"type\":\"bool\"},{\"name\":\"delegate\",\"type\":\"address\"},{\"name\":\"vote\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testdata\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testmsg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testgas\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winnerName\",\"outputs\":[{\"name\":\"winnerNames\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testGasPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BallotBin is the compiled bytecode used for deploying new contracts.
const BallotBin = `0x6060604052341561000f57600080fd5b6040516107ee3803806107ee8339810160405280805160008054600160a060020a03191633600160a060020a039081169190911780835516815260016020819052604082205592019190505b81518110156100d057600280546001810161007683826100d7565b91600052602060002090600202016000604080519081016040528086868151811061009d57fe5b906020019060200201518152600060209091015291905081518155602082015160019182015592909201915061005b9050565b505061012f565b815481835581811511610103576002028160020283600052602060002091820191016101039190610108565b505050565b61012c91905b80821115610128576000808255600182015560020161010e565b5090565b90565b6106b08061013e6000396000f300606060405236156100c25763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630121b93f81146100c7578063013cf08b146100df5780632e4176cf1461010d5780635c19a95c1461013c578063609ff1bd1461015b57806399244d1a146101805780639e7b8d6114610193578063a3ec138d146101b2578063b449c68914610206578063be8f316a14610290578063c8e9fc8b146102a3578063e2ba53f0146102b6578063e9de432f146102c9575b600080fd5b34156100d257600080fd5b6100dd6004356102dc565b005b34156100ea57600080fd5b6100f560043561034f565b60405191825260208201526040908101905180910390f35b341561011857600080fd5b61012061037b565b604051600160a060020a03909116815260200160405180910390f35b341561014757600080fd5b6100dd600160a060020a036004351661038a565b341561016657600080fd5b61016e6104d0565b60405190815260200160405180910390f35b341561018b57600080fd5b61016e61053a565b341561019e57600080fd5b6100dd600160a060020a036004351661053e565b34156101bd57600080fd5b6101d1600160a060020a03600435166105c3565b6040519384529115156020840152600160a060020a031660408084019190915260608301919091526080909101905180910390f35b341561021157600080fd5b6102196105f7565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561025557808201518382015260200161023d565b50505050905090810190601f1680156102825780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561029b57600080fd5b610120610635565b34156102ae57600080fd5b61016e610639565b34156102c157600080fd5b61016e610641565b34156102d457600080fd5b61016e61066e565b600160a060020a03331660009081526001602081905260409091209081015460ff161561030857600080fd5b6001818101805460ff191690911790556002808201839055815481549091908490811061033157fe5b60009182526020909120600160029092020101805490910190555050565b600280548290811061035d57fe5b60009182526020909120600290910201805460019091015490915082565b600054600160a060020a031681565b600160a060020a033316600090815260016020819052604082209081015490919060ff16156103b857600080fd5b33600160a060020a031683600160a060020a0316141515156103d957600080fd5b600160a060020a03838116600090815260016020819052604090912001546101009004161561043b57600160a060020a03928316600090815260016020819052604090912001546101009004831692331683141561043657600080fd5b6103d9565b506001818101805460ff1916821774ffffffffffffffffffffffffffffffffffffffff001916610100600160a060020a0386169081029190911790915560009081526020829052604090209081015460ff16156104c3578154600282810154815481106104a457fe5b60009182526020909120600160029092020101805490910190556104cb565b815481540181555b505050565b600080805b60025481101561053557816002828154811015156104ef57fe5b906000526020600020906002020160010154111561052d57600280548290811061051557fe5b90600052602060002090600202016001015491508092505b6001016104d5565b505090565b3490565b60005433600160a060020a0390811691161480156105795750600160a060020a0381166000908152600160208190526040909120015460ff16155b801561059b5750600160a060020a038116600090815260016020526040902054155b15156105a657600080fd5b600160a060020a0316600090815260016020819052604090912055565b600160208190526000918252604090912080549181015460029091015460ff8216916101009004600160a060020a03169084565b6105ff610672565b6000368080601f016020809104026020016040519081016040528181529291906020840183838082843750949550505050505090565b3390565b60005a905090565b6000600261064d6104d0565b8154811061065757fe5b906000526020600020906002020160000154905090565b3a90565b602060405190810160405260008152905600a165627a7a7230582032f3081a466e20f3b4ccfa6b7da0f3396ccece428a543b320454e2551b7b9b570029`

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
