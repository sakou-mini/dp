// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ScryProtocolABI is the input ABI used to generate the binding from.
const ScryProtocolABI = "[{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"despDataId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"supportVerify\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"DataPublish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"proofIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"needVerify\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"TransactionCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"TransactionClose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"proofIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"VerifiersChosen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"metaDataIdEncBuyer\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"ReadyForDownload\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"publishId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"proofIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"metaDataIdEncArbitrator\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"ArbitrationBegin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"judge\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"identify\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"ArbitrationResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"RegisterVerifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"judge\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"comments\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"seqNo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"VerifierDisable\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"}],\"name\":\"registerAsVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"judge\",\"type\":\"bool\"},{\"name\":\"comments\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"verifierIndex\",\"type\":\"uint8\"},{\"name\":\"credit\",\"type\":\"uint8\"}],\"name\":\"creditsToVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"judge\",\"type\":\"bool\"}],\"name\":\"arbitrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"publishId\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"metaDataIdEncSeller\",\"type\":\"bytes\"},{\"name\":\"proofDataIds\",\"type\":\"bytes32[]\"},{\"name\":\"descDataId\",\"type\":\"string\"},{\"name\":\"supportVerify\",\"type\":\"bool\"}],\"name\":\"publishDataInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"publishId\",\"type\":\"string\"},{\"name\":\"startVerify\",\"type\":\"bool\"}],\"name\":\"createTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"buyData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"cancelTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"encryptedMetaDataId\",\"type\":\"bytes\"},{\"name\":\"encryptedMetaDataIds\",\"type\":\"bytes\"}],\"name\":\"reEncryptMetaDataIdBySeller\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seqNo\",\"type\":\"string\"},{\"name\":\"txId\",\"type\":\"uint256\"},{\"name\":\"truth\",\"type\":\"bool\"}],\"name\":\"confirmDataTruth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"getBuyer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"txId\",\"type\":\"uint256\"}],\"name\":\"getArbitrators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ScryProtocol is an auto generated Go binding around an Ethereum contract.
type ScryProtocol struct {
	ScryProtocolCaller     // Read-only binding to the contract
	ScryProtocolTransactor // Write-only binding to the contract
	ScryProtocolFilterer   // Log filterer for contract events
}

// ScryProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScryProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScryProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScryProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScryProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScryProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScryProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScryProtocolSession struct {
	Contract     *ScryProtocol     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScryProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScryProtocolCallerSession struct {
	Contract *ScryProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ScryProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScryProtocolTransactorSession struct {
	Contract     *ScryProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ScryProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScryProtocolRaw struct {
	Contract *ScryProtocol // Generic contract binding to access the raw methods on
}

// ScryProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScryProtocolCallerRaw struct {
	Contract *ScryProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ScryProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScryProtocolTransactorRaw struct {
	Contract *ScryProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScryProtocol creates a new instance of ScryProtocol, bound to a specific deployed contract.
func NewScryProtocol(address common.Address, backend bind.ContractBackend) (*ScryProtocol, error) {
	contract, err := bindScryProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScryProtocol{ScryProtocolCaller: ScryProtocolCaller{contract: contract}, ScryProtocolTransactor: ScryProtocolTransactor{contract: contract}, ScryProtocolFilterer: ScryProtocolFilterer{contract: contract}}, nil
}

// NewScryProtocolCaller creates a new read-only instance of ScryProtocol, bound to a specific deployed contract.
func NewScryProtocolCaller(address common.Address, caller bind.ContractCaller) (*ScryProtocolCaller, error) {
	contract, err := bindScryProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScryProtocolCaller{contract: contract}, nil
}

// NewScryProtocolTransactor creates a new write-only instance of ScryProtocol, bound to a specific deployed contract.
func NewScryProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ScryProtocolTransactor, error) {
	contract, err := bindScryProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScryProtocolTransactor{contract: contract}, nil
}

// NewScryProtocolFilterer creates a new log filterer instance of ScryProtocol, bound to a specific deployed contract.
func NewScryProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ScryProtocolFilterer, error) {
	contract, err := bindScryProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScryProtocolFilterer{contract: contract}, nil
}

// bindScryProtocol binds a generic wrapper to an already deployed contract.
func bindScryProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScryProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScryProtocol *ScryProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ScryProtocol.Contract.ScryProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScryProtocol *ScryProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ScryProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScryProtocol *ScryProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ScryProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScryProtocol *ScryProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ScryProtocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScryProtocol *ScryProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScryProtocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScryProtocol *ScryProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScryProtocol.Contract.contract.Transact(opts, method, params...)
}

// GetArbitrators is a free data retrieval call binding the contract method 0x3297e591.
//
// Solidity: function getArbitrators(txId uint256) constant returns(address[])
func (_ScryProtocol *ScryProtocolCaller) GetArbitrators(opts *bind.CallOpts, txId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ScryProtocol.contract.Call(opts, out, "getArbitrators", txId)
	return *ret0, err
}

// GetArbitrators is a free data retrieval call binding the contract method 0x3297e591.
//
// Solidity: function getArbitrators(txId uint256) constant returns(address[])
func (_ScryProtocol *ScryProtocolSession) GetArbitrators(txId *big.Int) ([]common.Address, error) {
	return _ScryProtocol.Contract.GetArbitrators(&_ScryProtocol.CallOpts, txId)
}

// GetArbitrators is a free data retrieval call binding the contract method 0x3297e591.
//
// Solidity: function getArbitrators(txId uint256) constant returns(address[])
func (_ScryProtocol *ScryProtocolCallerSession) GetArbitrators(txId *big.Int) ([]common.Address, error) {
	return _ScryProtocol.Contract.GetArbitrators(&_ScryProtocol.CallOpts, txId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(txId uint256) constant returns(address)
func (_ScryProtocol *ScryProtocolCaller) GetBuyer(opts *bind.CallOpts, txId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ScryProtocol.contract.Call(opts, out, "getBuyer", txId)
	return *ret0, err
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(txId uint256) constant returns(address)
func (_ScryProtocol *ScryProtocolSession) GetBuyer(txId *big.Int) (common.Address, error) {
	return _ScryProtocol.Contract.GetBuyer(&_ScryProtocol.CallOpts, txId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(txId uint256) constant returns(address)
func (_ScryProtocol *ScryProtocolCallerSession) GetBuyer(txId *big.Int) (common.Address, error) {
	return _ScryProtocol.Contract.GetBuyer(&_ScryProtocol.CallOpts, txId)
}

// Arbitrate is a paid mutator transaction binding the contract method 0xdaff1168.
//
// Solidity: function arbitrate(seqNo string, txId uint256, judge bool) returns()
func (_ScryProtocol *ScryProtocolTransactor) Arbitrate(opts *bind.TransactOpts, seqNo string, txId *big.Int, judge bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "arbitrate", seqNo, txId, judge)
}

// Arbitrate is a paid mutator transaction binding the contract method 0xdaff1168.
//
// Solidity: function arbitrate(seqNo string, txId uint256, judge bool) returns()
func (_ScryProtocol *ScryProtocolSession) Arbitrate(seqNo string, txId *big.Int, judge bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Arbitrate(&_ScryProtocol.TransactOpts, seqNo, txId, judge)
}

// Arbitrate is a paid mutator transaction binding the contract method 0xdaff1168.
//
// Solidity: function arbitrate(seqNo string, txId uint256, judge bool) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) Arbitrate(seqNo string, txId *big.Int, judge bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Arbitrate(&_ScryProtocol.TransactOpts, seqNo, txId, judge)
}

// BuyData is a paid mutator transaction binding the contract method 0x9a756a99.
//
// Solidity: function buyData(seqNo string, txId uint256) returns()
func (_ScryProtocol *ScryProtocolTransactor) BuyData(opts *bind.TransactOpts, seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "buyData", seqNo, txId)
}

// BuyData is a paid mutator transaction binding the contract method 0x9a756a99.
//
// Solidity: function buyData(seqNo string, txId uint256) returns()
func (_ScryProtocol *ScryProtocolSession) BuyData(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.BuyData(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// BuyData is a paid mutator transaction binding the contract method 0x9a756a99.
//
// Solidity: function buyData(seqNo string, txId uint256) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) BuyData(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.BuyData(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0xcca8f8c3.
//
// Solidity: function cancelTransaction(seqNo string, txId uint256) returns()
func (_ScryProtocol *ScryProtocolTransactor) CancelTransaction(opts *bind.TransactOpts, seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "cancelTransaction", seqNo, txId)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0xcca8f8c3.
//
// Solidity: function cancelTransaction(seqNo string, txId uint256) returns()
func (_ScryProtocol *ScryProtocolSession) CancelTransaction(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CancelTransaction(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// CancelTransaction is a paid mutator transaction binding the contract method 0xcca8f8c3.
//
// Solidity: function cancelTransaction(seqNo string, txId uint256) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) CancelTransaction(seqNo string, txId *big.Int) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CancelTransaction(&_ScryProtocol.TransactOpts, seqNo, txId)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0xdd16673b.
//
// Solidity: function confirmDataTruth(seqNo string, txId uint256, truth bool) returns()
func (_ScryProtocol *ScryProtocolTransactor) ConfirmDataTruth(opts *bind.TransactOpts, seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "confirmDataTruth", seqNo, txId, truth)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0xdd16673b.
//
// Solidity: function confirmDataTruth(seqNo string, txId uint256, truth bool) returns()
func (_ScryProtocol *ScryProtocolSession) ConfirmDataTruth(seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmDataTruth(&_ScryProtocol.TransactOpts, seqNo, txId, truth)
}

// ConfirmDataTruth is a paid mutator transaction binding the contract method 0xdd16673b.
//
// Solidity: function confirmDataTruth(seqNo string, txId uint256, truth bool) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) ConfirmDataTruth(seqNo string, txId *big.Int, truth bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ConfirmDataTruth(&_ScryProtocol.TransactOpts, seqNo, txId, truth)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xccc2ba76.
//
// Solidity: function createTransaction(seqNo string, publishId string, startVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactor) CreateTransaction(opts *bind.TransactOpts, seqNo string, publishId string, startVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "createTransaction", seqNo, publishId, startVerify)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xccc2ba76.
//
// Solidity: function createTransaction(seqNo string, publishId string, startVerify bool) returns()
func (_ScryProtocol *ScryProtocolSession) CreateTransaction(seqNo string, publishId string, startVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreateTransaction(&_ScryProtocol.TransactOpts, seqNo, publishId, startVerify)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xccc2ba76.
//
// Solidity: function createTransaction(seqNo string, publishId string, startVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) CreateTransaction(seqNo string, publishId string, startVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreateTransaction(&_ScryProtocol.TransactOpts, seqNo, publishId, startVerify)
}

// CreditsToVerifier is a paid mutator transaction binding the contract method 0xbd7ff8e5.
//
// Solidity: function creditsToVerifier(seqNo string, txId uint256, verifierIndex uint8, credit uint8) returns()
func (_ScryProtocol *ScryProtocolTransactor) CreditsToVerifier(opts *bind.TransactOpts, seqNo string, txId *big.Int, verifierIndex uint8, credit uint8) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "creditsToVerifier", seqNo, txId, verifierIndex, credit)
}

// CreditsToVerifier is a paid mutator transaction binding the contract method 0xbd7ff8e5.
//
// Solidity: function creditsToVerifier(seqNo string, txId uint256, verifierIndex uint8, credit uint8) returns()
func (_ScryProtocol *ScryProtocolSession) CreditsToVerifier(seqNo string, txId *big.Int, verifierIndex uint8, credit uint8) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreditsToVerifier(&_ScryProtocol.TransactOpts, seqNo, txId, verifierIndex, credit)
}

// CreditsToVerifier is a paid mutator transaction binding the contract method 0xbd7ff8e5.
//
// Solidity: function creditsToVerifier(seqNo string, txId uint256, verifierIndex uint8, credit uint8) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) CreditsToVerifier(seqNo string, txId *big.Int, verifierIndex uint8, credit uint8) (*types.Transaction, error) {
	return _ScryProtocol.Contract.CreditsToVerifier(&_ScryProtocol.TransactOpts, seqNo, txId, verifierIndex, credit)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0x377caed1.
//
// Solidity: function publishDataInfo(seqNo string, publishId string, price uint256, metaDataIdEncSeller bytes, proofDataIds bytes32[], descDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactor) PublishDataInfo(opts *bind.TransactOpts, seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, descDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "publishDataInfo", seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, descDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0x377caed1.
//
// Solidity: function publishDataInfo(seqNo string, publishId string, price uint256, metaDataIdEncSeller bytes, proofDataIds bytes32[], descDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolSession) PublishDataInfo(seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, descDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, descDataId, supportVerify)
}

// PublishDataInfo is a paid mutator transaction binding the contract method 0x377caed1.
//
// Solidity: function publishDataInfo(seqNo string, publishId string, price uint256, metaDataIdEncSeller bytes, proofDataIds bytes32[], descDataId string, supportVerify bool) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) PublishDataInfo(seqNo string, publishId string, price *big.Int, metaDataIdEncSeller []byte, proofDataIds [][32]byte, descDataId string, supportVerify bool) (*types.Transaction, error) {
	return _ScryProtocol.Contract.PublishDataInfo(&_ScryProtocol.TransactOpts, seqNo, publishId, price, metaDataIdEncSeller, proofDataIds, descDataId, supportVerify)
}

// ReEncryptMetaDataIdBySeller is a paid mutator transaction binding the contract method 0x275e7bc8.
//
// Solidity: function reEncryptMetaDataIdBySeller(seqNo string, txId uint256, encryptedMetaDataId bytes, encryptedMetaDataIds bytes) returns()
func (_ScryProtocol *ScryProtocolTransactor) ReEncryptMetaDataIdBySeller(opts *bind.TransactOpts, seqNo string, txId *big.Int, encryptedMetaDataId []byte, encryptedMetaDataIds []byte) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "reEncryptMetaDataIdBySeller", seqNo, txId, encryptedMetaDataId, encryptedMetaDataIds)
}

// ReEncryptMetaDataIdBySeller is a paid mutator transaction binding the contract method 0x275e7bc8.
//
// Solidity: function reEncryptMetaDataIdBySeller(seqNo string, txId uint256, encryptedMetaDataId bytes, encryptedMetaDataIds bytes) returns()
func (_ScryProtocol *ScryProtocolSession) ReEncryptMetaDataIdBySeller(seqNo string, txId *big.Int, encryptedMetaDataId []byte, encryptedMetaDataIds []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ReEncryptMetaDataIdBySeller(&_ScryProtocol.TransactOpts, seqNo, txId, encryptedMetaDataId, encryptedMetaDataIds)
}

// ReEncryptMetaDataIdBySeller is a paid mutator transaction binding the contract method 0x275e7bc8.
//
// Solidity: function reEncryptMetaDataIdBySeller(seqNo string, txId uint256, encryptedMetaDataId bytes, encryptedMetaDataIds bytes) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) ReEncryptMetaDataIdBySeller(seqNo string, txId *big.Int, encryptedMetaDataId []byte, encryptedMetaDataIds []byte) (*types.Transaction, error) {
	return _ScryProtocol.Contract.ReEncryptMetaDataIdBySeller(&_ScryProtocol.TransactOpts, seqNo, txId, encryptedMetaDataId, encryptedMetaDataIds)
}

// RegisterAsVerifier is a paid mutator transaction binding the contract method 0x93151dd5.
//
// Solidity: function registerAsVerifier(seqNo string) returns()
func (_ScryProtocol *ScryProtocolTransactor) RegisterAsVerifier(opts *bind.TransactOpts, seqNo string) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "registerAsVerifier", seqNo)
}

// RegisterAsVerifier is a paid mutator transaction binding the contract method 0x93151dd5.
//
// Solidity: function registerAsVerifier(seqNo string) returns()
func (_ScryProtocol *ScryProtocolSession) RegisterAsVerifier(seqNo string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.RegisterAsVerifier(&_ScryProtocol.TransactOpts, seqNo)
}

// RegisterAsVerifier is a paid mutator transaction binding the contract method 0x93151dd5.
//
// Solidity: function registerAsVerifier(seqNo string) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) RegisterAsVerifier(seqNo string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.RegisterAsVerifier(&_ScryProtocol.TransactOpts, seqNo)
}

// Vote is a paid mutator transaction binding the contract method 0x980da40d.
//
// Solidity: function vote(seqNo string, txId uint256, judge bool, comments string) returns()
func (_ScryProtocol *ScryProtocolTransactor) Vote(opts *bind.TransactOpts, seqNo string, txId *big.Int, judge bool, comments string) (*types.Transaction, error) {
	return _ScryProtocol.contract.Transact(opts, "vote", seqNo, txId, judge, comments)
}

// Vote is a paid mutator transaction binding the contract method 0x980da40d.
//
// Solidity: function vote(seqNo string, txId uint256, judge bool, comments string) returns()
func (_ScryProtocol *ScryProtocolSession) Vote(seqNo string, txId *big.Int, judge bool, comments string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Vote(&_ScryProtocol.TransactOpts, seqNo, txId, judge, comments)
}

// Vote is a paid mutator transaction binding the contract method 0x980da40d.
//
// Solidity: function vote(seqNo string, txId uint256, judge bool, comments string) returns()
func (_ScryProtocol *ScryProtocolTransactorSession) Vote(seqNo string, txId *big.Int, judge bool, comments string) (*types.Transaction, error) {
	return _ScryProtocol.Contract.Vote(&_ScryProtocol.TransactOpts, seqNo, txId, judge, comments)
}

// ScryProtocolArbitrationBeginIterator is returned from FilterArbitrationBegin and is used to iterate over the raw logs and unpacked data for ArbitrationBegin events raised by the ScryProtocol contract.
type ScryProtocolArbitrationBeginIterator struct {
	Event *ScryProtocolArbitrationBegin // Event containing the contract specifics and raw log

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
func (it *ScryProtocolArbitrationBeginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolArbitrationBegin)
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
		it.Event = new(ScryProtocolArbitrationBegin)
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
func (it *ScryProtocolArbitrationBeginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolArbitrationBeginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolArbitrationBegin represents a ArbitrationBegin event raised by the ScryProtocol contract.
type ScryProtocolArbitrationBegin struct {
	SeqNo                   string
	TransactionId           *big.Int
	PublishId               string
	ProofIds                [][32]byte
	MetaDataIdEncArbitrator []byte
	Users                   []common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterArbitrationBegin is a free log retrieval operation binding the contract event 0x541c5ec36a883a22048ce07131e93f1c295bd6ccfed09aaab4e233e9b07c87b0.
//
// Solidity: e ArbitrationBegin(seqNo string, transactionId uint256, publishId string, proofIds bytes32[], metaDataIdEncArbitrator bytes, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterArbitrationBegin(opts *bind.FilterOpts) (*ScryProtocolArbitrationBeginIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "ArbitrationBegin")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolArbitrationBeginIterator{contract: _ScryProtocol.contract, event: "ArbitrationBegin", logs: logs, sub: sub}, nil
}

// WatchArbitrationBegin is a free log subscription operation binding the contract event 0x541c5ec36a883a22048ce07131e93f1c295bd6ccfed09aaab4e233e9b07c87b0.
//
// Solidity: e ArbitrationBegin(seqNo string, transactionId uint256, publishId string, proofIds bytes32[], metaDataIdEncArbitrator bytes, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchArbitrationBegin(opts *bind.WatchOpts, sink chan<- *ScryProtocolArbitrationBegin) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "ArbitrationBegin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolArbitrationBegin)
				if err := _ScryProtocol.contract.UnpackLog(event, "ArbitrationBegin", log); err != nil {
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

// ScryProtocolArbitrationResultIterator is returned from FilterArbitrationResult and is used to iterate over the raw logs and unpacked data for ArbitrationResult events raised by the ScryProtocol contract.
type ScryProtocolArbitrationResultIterator struct {
	Event *ScryProtocolArbitrationResult // Event containing the contract specifics and raw log

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
func (it *ScryProtocolArbitrationResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolArbitrationResult)
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
		it.Event = new(ScryProtocolArbitrationResult)
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
func (it *ScryProtocolArbitrationResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolArbitrationResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolArbitrationResult represents a ArbitrationResult event raised by the ScryProtocol contract.
type ScryProtocolArbitrationResult struct {
	SeqNo         string
	TransactionId *big.Int
	Judge         bool
	Identify      uint8
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterArbitrationResult is a free log retrieval operation binding the contract event 0x8b3b8af680cfe70f3a16a2750d9dee312f8002fd35bd049d61befe3cb7c5a3c7.
//
// Solidity: e ArbitrationResult(seqNo string, transactionId uint256, judge bool, identify uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterArbitrationResult(opts *bind.FilterOpts) (*ScryProtocolArbitrationResultIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "ArbitrationResult")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolArbitrationResultIterator{contract: _ScryProtocol.contract, event: "ArbitrationResult", logs: logs, sub: sub}, nil
}

// WatchArbitrationResult is a free log subscription operation binding the contract event 0x8b3b8af680cfe70f3a16a2750d9dee312f8002fd35bd049d61befe3cb7c5a3c7.
//
// Solidity: e ArbitrationResult(seqNo string, transactionId uint256, judge bool, identify uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchArbitrationResult(opts *bind.WatchOpts, sink chan<- *ScryProtocolArbitrationResult) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "ArbitrationResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolArbitrationResult)
				if err := _ScryProtocol.contract.UnpackLog(event, "ArbitrationResult", log); err != nil {
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

// ScryProtocolBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the ScryProtocol contract.
type ScryProtocolBuyIterator struct {
	Event *ScryProtocolBuy // Event containing the contract specifics and raw log

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
func (it *ScryProtocolBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolBuy)
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
		it.Event = new(ScryProtocolBuy)
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
func (it *ScryProtocolBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolBuy represents a Buy event raised by the ScryProtocol contract.
type ScryProtocolBuy struct {
	SeqNo               string
	TransactionId       *big.Int
	PublishId           string
	MetaDataIdEncSeller []byte
	State               uint8
	Index               uint8
	Users               []common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0x8fc09fab957b20f7258d7947b08e0b4ef8cca78ab1b4111fc8af6104d024613b.
//
// Solidity: e Buy(seqNo string, transactionId uint256, publishId string, metaDataIdEncSeller bytes, state uint8, index uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterBuy(opts *bind.FilterOpts) (*ScryProtocolBuyIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "Buy")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolBuyIterator{contract: _ScryProtocol.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0x8fc09fab957b20f7258d7947b08e0b4ef8cca78ab1b4111fc8af6104d024613b.
//
// Solidity: e Buy(seqNo string, transactionId uint256, publishId string, metaDataIdEncSeller bytes, state uint8, index uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *ScryProtocolBuy) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "Buy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolBuy)
				if err := _ScryProtocol.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ScryProtocolDataPublishIterator is returned from FilterDataPublish and is used to iterate over the raw logs and unpacked data for DataPublish events raised by the ScryProtocol contract.
type ScryProtocolDataPublishIterator struct {
	Event *ScryProtocolDataPublish // Event containing the contract specifics and raw log

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
func (it *ScryProtocolDataPublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolDataPublish)
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
		it.Event = new(ScryProtocolDataPublish)
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
func (it *ScryProtocolDataPublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolDataPublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolDataPublish represents a DataPublish event raised by the ScryProtocol contract.
type ScryProtocolDataPublish struct {
	SeqNo         string
	PublishId     string
	Price         *big.Int
	DespDataId    string
	SupportVerify bool
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDataPublish is a free log retrieval operation binding the contract event 0x382c732523b607b116dd2681b4e54c31b1ee534cf9eed7bcb98e1ec8d43ecbea.
//
// Solidity: e DataPublish(seqNo string, publishId string, price uint256, despDataId string, supportVerify bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterDataPublish(opts *bind.FilterOpts) (*ScryProtocolDataPublishIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "DataPublish")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolDataPublishIterator{contract: _ScryProtocol.contract, event: "DataPublish", logs: logs, sub: sub}, nil
}

// WatchDataPublish is a free log subscription operation binding the contract event 0x382c732523b607b116dd2681b4e54c31b1ee534cf9eed7bcb98e1ec8d43ecbea.
//
// Solidity: e DataPublish(seqNo string, publishId string, price uint256, despDataId string, supportVerify bool, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchDataPublish(opts *bind.WatchOpts, sink chan<- *ScryProtocolDataPublish) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "DataPublish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolDataPublish)
				if err := _ScryProtocol.contract.UnpackLog(event, "DataPublish", log); err != nil {
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

// ScryProtocolReadyForDownloadIterator is returned from FilterReadyForDownload and is used to iterate over the raw logs and unpacked data for ReadyForDownload events raised by the ScryProtocol contract.
type ScryProtocolReadyForDownloadIterator struct {
	Event *ScryProtocolReadyForDownload // Event containing the contract specifics and raw log

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
func (it *ScryProtocolReadyForDownloadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolReadyForDownload)
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
		it.Event = new(ScryProtocolReadyForDownload)
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
func (it *ScryProtocolReadyForDownloadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolReadyForDownloadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolReadyForDownload represents a ReadyForDownload event raised by the ScryProtocol contract.
type ScryProtocolReadyForDownload struct {
	SeqNo              string
	TransactionId      *big.Int
	MetaDataIdEncBuyer []byte
	State              uint8
	Index              uint8
	Users              []common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReadyForDownload is a free log retrieval operation binding the contract event 0xef3b9511479459f28053c72c9c2deba123aa2ddf021496d5cce3b788f7494e5c.
//
// Solidity: e ReadyForDownload(seqNo string, transactionId uint256, metaDataIdEncBuyer bytes, state uint8, index uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterReadyForDownload(opts *bind.FilterOpts) (*ScryProtocolReadyForDownloadIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "ReadyForDownload")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolReadyForDownloadIterator{contract: _ScryProtocol.contract, event: "ReadyForDownload", logs: logs, sub: sub}, nil
}

// WatchReadyForDownload is a free log subscription operation binding the contract event 0xef3b9511479459f28053c72c9c2deba123aa2ddf021496d5cce3b788f7494e5c.
//
// Solidity: e ReadyForDownload(seqNo string, transactionId uint256, metaDataIdEncBuyer bytes, state uint8, index uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchReadyForDownload(opts *bind.WatchOpts, sink chan<- *ScryProtocolReadyForDownload) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "ReadyForDownload")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolReadyForDownload)
				if err := _ScryProtocol.contract.UnpackLog(event, "ReadyForDownload", log); err != nil {
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

// ScryProtocolRegisterVerifierIterator is returned from FilterRegisterVerifier and is used to iterate over the raw logs and unpacked data for RegisterVerifier events raised by the ScryProtocol contract.
type ScryProtocolRegisterVerifierIterator struct {
	Event *ScryProtocolRegisterVerifier // Event containing the contract specifics and raw log

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
func (it *ScryProtocolRegisterVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolRegisterVerifier)
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
		it.Event = new(ScryProtocolRegisterVerifier)
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
func (it *ScryProtocolRegisterVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolRegisterVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolRegisterVerifier represents a RegisterVerifier event raised by the ScryProtocol contract.
type ScryProtocolRegisterVerifier struct {
	SeqNo string
	Users []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRegisterVerifier is a free log retrieval operation binding the contract event 0x476785064b6fb8cce78cd4377a03177c7bac7803ef345a1eaf34d1dbdbf0e864.
//
// Solidity: e RegisterVerifier(seqNo string, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterRegisterVerifier(opts *bind.FilterOpts) (*ScryProtocolRegisterVerifierIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "RegisterVerifier")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolRegisterVerifierIterator{contract: _ScryProtocol.contract, event: "RegisterVerifier", logs: logs, sub: sub}, nil
}

// WatchRegisterVerifier is a free log subscription operation binding the contract event 0x476785064b6fb8cce78cd4377a03177c7bac7803ef345a1eaf34d1dbdbf0e864.
//
// Solidity: e RegisterVerifier(seqNo string, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchRegisterVerifier(opts *bind.WatchOpts, sink chan<- *ScryProtocolRegisterVerifier) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "RegisterVerifier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolRegisterVerifier)
				if err := _ScryProtocol.contract.UnpackLog(event, "RegisterVerifier", log); err != nil {
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

// ScryProtocolTransactionCloseIterator is returned from FilterTransactionClose and is used to iterate over the raw logs and unpacked data for TransactionClose events raised by the ScryProtocol contract.
type ScryProtocolTransactionCloseIterator struct {
	Event *ScryProtocolTransactionClose // Event containing the contract specifics and raw log

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
func (it *ScryProtocolTransactionCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolTransactionClose)
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
		it.Event = new(ScryProtocolTransactionClose)
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
func (it *ScryProtocolTransactionCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolTransactionCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolTransactionClose represents a TransactionClose event raised by the ScryProtocol contract.
type ScryProtocolTransactionClose struct {
	SeqNo         string
	TransactionId *big.Int
	State         uint8
	Index         uint8
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionClose is a free log retrieval operation binding the contract event 0xef210c4e1a93373cd651f7dea507cd35c6eb5b14f64ce0f9dd941e0d4be1cbb7.
//
// Solidity: e TransactionClose(seqNo string, transactionId uint256, state uint8, index uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterTransactionClose(opts *bind.FilterOpts) (*ScryProtocolTransactionCloseIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "TransactionClose")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolTransactionCloseIterator{contract: _ScryProtocol.contract, event: "TransactionClose", logs: logs, sub: sub}, nil
}

// WatchTransactionClose is a free log subscription operation binding the contract event 0xef210c4e1a93373cd651f7dea507cd35c6eb5b14f64ce0f9dd941e0d4be1cbb7.
//
// Solidity: e TransactionClose(seqNo string, transactionId uint256, state uint8, index uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchTransactionClose(opts *bind.WatchOpts, sink chan<- *ScryProtocolTransactionClose) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "TransactionClose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolTransactionClose)
				if err := _ScryProtocol.contract.UnpackLog(event, "TransactionClose", log); err != nil {
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

// ScryProtocolTransactionCreateIterator is returned from FilterTransactionCreate and is used to iterate over the raw logs and unpacked data for TransactionCreate events raised by the ScryProtocol contract.
type ScryProtocolTransactionCreateIterator struct {
	Event *ScryProtocolTransactionCreate // Event containing the contract specifics and raw log

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
func (it *ScryProtocolTransactionCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolTransactionCreate)
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
		it.Event = new(ScryProtocolTransactionCreate)
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
func (it *ScryProtocolTransactionCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolTransactionCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolTransactionCreate represents a TransactionCreate event raised by the ScryProtocol contract.
type ScryProtocolTransactionCreate struct {
	SeqNo         string
	TransactionId *big.Int
	PublishId     string
	ProofIds      [][32]byte
	NeedVerify    bool
	State         uint8
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionCreate is a free log retrieval operation binding the contract event 0x87712a8083757ec54dddd564e3755ff6db99131f0a2f8e8db57795d71926a781.
//
// Solidity: e TransactionCreate(seqNo string, transactionId uint256, publishId string, proofIds bytes32[], needVerify bool, state uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterTransactionCreate(opts *bind.FilterOpts) (*ScryProtocolTransactionCreateIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "TransactionCreate")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolTransactionCreateIterator{contract: _ScryProtocol.contract, event: "TransactionCreate", logs: logs, sub: sub}, nil
}

// WatchTransactionCreate is a free log subscription operation binding the contract event 0x87712a8083757ec54dddd564e3755ff6db99131f0a2f8e8db57795d71926a781.
//
// Solidity: e TransactionCreate(seqNo string, transactionId uint256, publishId string, proofIds bytes32[], needVerify bool, state uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchTransactionCreate(opts *bind.WatchOpts, sink chan<- *ScryProtocolTransactionCreate) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "TransactionCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolTransactionCreate)
				if err := _ScryProtocol.contract.UnpackLog(event, "TransactionCreate", log); err != nil {
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

// ScryProtocolVerifierDisableIterator is returned from FilterVerifierDisable and is used to iterate over the raw logs and unpacked data for VerifierDisable events raised by the ScryProtocol contract.
type ScryProtocolVerifierDisableIterator struct {
	Event *ScryProtocolVerifierDisable // Event containing the contract specifics and raw log

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
func (it *ScryProtocolVerifierDisableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolVerifierDisable)
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
		it.Event = new(ScryProtocolVerifierDisable)
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
func (it *ScryProtocolVerifierDisableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolVerifierDisableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolVerifierDisable represents a VerifierDisable event raised by the ScryProtocol contract.
type ScryProtocolVerifierDisable struct {
	SeqNo    string
	Verifier common.Address
	Users    []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierDisable is a free log retrieval operation binding the contract event 0xd0e2127bc672e5762b6852e277e4d77594e1eecad47e50f41f56d9a87c1f7505.
//
// Solidity: e VerifierDisable(seqNo string, verifier address, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterVerifierDisable(opts *bind.FilterOpts) (*ScryProtocolVerifierDisableIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "VerifierDisable")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolVerifierDisableIterator{contract: _ScryProtocol.contract, event: "VerifierDisable", logs: logs, sub: sub}, nil
}

// WatchVerifierDisable is a free log subscription operation binding the contract event 0xd0e2127bc672e5762b6852e277e4d77594e1eecad47e50f41f56d9a87c1f7505.
//
// Solidity: e VerifierDisable(seqNo string, verifier address, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchVerifierDisable(opts *bind.WatchOpts, sink chan<- *ScryProtocolVerifierDisable) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "VerifierDisable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolVerifierDisable)
				if err := _ScryProtocol.contract.UnpackLog(event, "VerifierDisable", log); err != nil {
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

// ScryProtocolVerifiersChosenIterator is returned from FilterVerifiersChosen and is used to iterate over the raw logs and unpacked data for VerifiersChosen events raised by the ScryProtocol contract.
type ScryProtocolVerifiersChosenIterator struct {
	Event *ScryProtocolVerifiersChosen // Event containing the contract specifics and raw log

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
func (it *ScryProtocolVerifiersChosenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolVerifiersChosen)
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
		it.Event = new(ScryProtocolVerifiersChosen)
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
func (it *ScryProtocolVerifiersChosenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolVerifiersChosenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolVerifiersChosen represents a VerifiersChosen event raised by the ScryProtocol contract.
type ScryProtocolVerifiersChosen struct {
	SeqNo         string
	TransactionId *big.Int
	PublishId     string
	ProofIds      [][32]byte
	State         uint8
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVerifiersChosen is a free log retrieval operation binding the contract event 0x8c48707022c17239432366b846b96252d3a8c7f241cbba874edb30ec41c1acb3.
//
// Solidity: e VerifiersChosen(seqNo string, transactionId uint256, publishId string, proofIds bytes32[], state uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterVerifiersChosen(opts *bind.FilterOpts) (*ScryProtocolVerifiersChosenIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "VerifiersChosen")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolVerifiersChosenIterator{contract: _ScryProtocol.contract, event: "VerifiersChosen", logs: logs, sub: sub}, nil
}

// WatchVerifiersChosen is a free log subscription operation binding the contract event 0x8c48707022c17239432366b846b96252d3a8c7f241cbba874edb30ec41c1acb3.
//
// Solidity: e VerifiersChosen(seqNo string, transactionId uint256, publishId string, proofIds bytes32[], state uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchVerifiersChosen(opts *bind.WatchOpts, sink chan<- *ScryProtocolVerifiersChosen) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "VerifiersChosen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolVerifiersChosen)
				if err := _ScryProtocol.contract.UnpackLog(event, "VerifiersChosen", log); err != nil {
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

// ScryProtocolVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the ScryProtocol contract.
type ScryProtocolVoteIterator struct {
	Event *ScryProtocolVote // Event containing the contract specifics and raw log

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
func (it *ScryProtocolVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScryProtocolVote)
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
		it.Event = new(ScryProtocolVote)
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
func (it *ScryProtocolVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScryProtocolVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScryProtocolVote represents a Vote event raised by the ScryProtocol contract.
type ScryProtocolVote struct {
	SeqNo         string
	TransactionId *big.Int
	Judge         bool
	Comments      string
	State         uint8
	Index         uint8
	Users         []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x0ffb7cbdd6a39a576bf0eedc7cf48b5d1da05a4d9e868397dafb21ffc1df9c6b.
//
// Solidity: e Vote(seqNo string, transactionId uint256, judge bool, comments string, state uint8, index uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) FilterVote(opts *bind.FilterOpts) (*ScryProtocolVoteIterator, error) {

	logs, sub, err := _ScryProtocol.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &ScryProtocolVoteIterator{contract: _ScryProtocol.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x0ffb7cbdd6a39a576bf0eedc7cf48b5d1da05a4d9e868397dafb21ffc1df9c6b.
//
// Solidity: e Vote(seqNo string, transactionId uint256, judge bool, comments string, state uint8, index uint8, users address[])
func (_ScryProtocol *ScryProtocolFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *ScryProtocolVote) (event.Subscription, error) {

	logs, sub, err := _ScryProtocol.contract.WatchLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScryProtocolVote)
				if err := _ScryProtocol.contract.UnpackLog(event, "Vote", log); err != nil {
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
