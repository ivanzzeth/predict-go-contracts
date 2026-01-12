// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yb_negriskadapter

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

// YBNegRiskAdapterMetaData contains all meta data concerning the YBNegRiskAdapter contract.
var YBNegRiskAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ctf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yieldManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DeterminedFlagAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeBipsOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIndexSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketAlreadyDetermined\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketAlreadyPrepared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketNotPrepared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoConvertiblePositions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedForAll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedCollateralToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"IsConvertPositionsGatedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"IsMergePositionsGatedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"IsSplitPositionGatedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MarketPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"outcome\",\"type\":\"bool\"}],\"name\":\"OutcomeReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"name\":\"PayoutRedemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"canSplitPosition\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"canMergePositions\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"canConvertPositions\",\"type\":\"bool\"}],\"name\":\"PermissionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionSplit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"indexSet\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionsConverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionsMerge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"marketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"QuestionPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedOperator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NO_TOKEN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"col\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_indexSet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertPositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctf\",\"outputs\":[{\"internalType\":\"contractIConditionalTokens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"}],\"name\":\"getConditionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getDetermined\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getFeeBips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getMarketData\",\"outputs\":[{\"internalType\":\"MarketData\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_outcome\",\"type\":\"bool\"}],\"name\":\"getPositionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getQuestionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"}],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isConvertPositionsGated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMergePositionsGated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSplitPositionGated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mergePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mergePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeBips\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"}],\"name\":\"prepareMarket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_marketId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"}],\"name\":\"prepareQuestion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"redeemPositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_questionId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_outcome\",\"type\":\"bool\"}],\"name\":\"reportOutcome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"setIsConvertPositionsGated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"setIsMergePositionsGated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"setIsSplitPositionGated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_canSplitPosition\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canMergePositions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canConvertPositions\",\"type\":\"bool\"}],\"name\":\"setPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"splitPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"splitPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userPermissions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"canSplitPosition\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canMergePositions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canConvertPositions\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wcol\",\"outputs\":[{\"internalType\":\"contractYieldBearingWrappedCollateral\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YBNegRiskAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use YBNegRiskAdapterMetaData.ABI instead.
var YBNegRiskAdapterABI = YBNegRiskAdapterMetaData.ABI

// YBNegRiskAdapter is an auto generated Go binding around an Ethereum contract.
type YBNegRiskAdapter struct {
	YBNegRiskAdapterCaller     // Read-only binding to the contract
	YBNegRiskAdapterTransactor // Write-only binding to the contract
	YBNegRiskAdapterFilterer   // Log filterer for contract events
}

// YBNegRiskAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type YBNegRiskAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBNegRiskAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YBNegRiskAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBNegRiskAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YBNegRiskAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBNegRiskAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YBNegRiskAdapterSession struct {
	Contract     *YBNegRiskAdapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YBNegRiskAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YBNegRiskAdapterCallerSession struct {
	Contract *YBNegRiskAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// YBNegRiskAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YBNegRiskAdapterTransactorSession struct {
	Contract     *YBNegRiskAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// YBNegRiskAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type YBNegRiskAdapterRaw struct {
	Contract *YBNegRiskAdapter // Generic contract binding to access the raw methods on
}

// YBNegRiskAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YBNegRiskAdapterCallerRaw struct {
	Contract *YBNegRiskAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// YBNegRiskAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YBNegRiskAdapterTransactorRaw struct {
	Contract *YBNegRiskAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYBNegRiskAdapter creates a new instance of YBNegRiskAdapter, bound to a specific deployed contract.
func NewYBNegRiskAdapter(address common.Address, backend bind.ContractBackend) (*YBNegRiskAdapter, error) {
	contract, err := bindYBNegRiskAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapter{YBNegRiskAdapterCaller: YBNegRiskAdapterCaller{contract: contract}, YBNegRiskAdapterTransactor: YBNegRiskAdapterTransactor{contract: contract}, YBNegRiskAdapterFilterer: YBNegRiskAdapterFilterer{contract: contract}}, nil
}

// NewYBNegRiskAdapterCaller creates a new read-only instance of YBNegRiskAdapter, bound to a specific deployed contract.
func NewYBNegRiskAdapterCaller(address common.Address, caller bind.ContractCaller) (*YBNegRiskAdapterCaller, error) {
	contract, err := bindYBNegRiskAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterCaller{contract: contract}, nil
}

// NewYBNegRiskAdapterTransactor creates a new write-only instance of YBNegRiskAdapter, bound to a specific deployed contract.
func NewYBNegRiskAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*YBNegRiskAdapterTransactor, error) {
	contract, err := bindYBNegRiskAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterTransactor{contract: contract}, nil
}

// NewYBNegRiskAdapterFilterer creates a new log filterer instance of YBNegRiskAdapter, bound to a specific deployed contract.
func NewYBNegRiskAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*YBNegRiskAdapterFilterer, error) {
	contract, err := bindYBNegRiskAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterFilterer{contract: contract}, nil
}

// bindYBNegRiskAdapter binds a generic wrapper to an already deployed contract.
func bindYBNegRiskAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YBNegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YBNegRiskAdapter *YBNegRiskAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YBNegRiskAdapter.Contract.YBNegRiskAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YBNegRiskAdapter *YBNegRiskAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.YBNegRiskAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YBNegRiskAdapter *YBNegRiskAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.YBNegRiskAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YBNegRiskAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.contract.Transact(opts, method, params...)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) FEEDENOMINATOR() (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.FEEDENOMINATOR(&_YBNegRiskAdapter.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.FEEDENOMINATOR(&_YBNegRiskAdapter.CallOpts)
}

// NOTOKENBURNADDRESS is a free data retrieval call binding the contract method 0x7ad7fe36.
//
// Solidity: function NO_TOKEN_BURN_ADDRESS() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) NOTOKENBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "NO_TOKEN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NOTOKENBURNADDRESS is a free data retrieval call binding the contract method 0x7ad7fe36.
//
// Solidity: function NO_TOKEN_BURN_ADDRESS() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) NOTOKENBURNADDRESS() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.NOTOKENBURNADDRESS(&_YBNegRiskAdapter.CallOpts)
}

// NOTOKENBURNADDRESS is a free data retrieval call binding the contract method 0x7ad7fe36.
//
// Solidity: function NO_TOKEN_BURN_ADDRESS() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) NOTOKENBURNADDRESS() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.NOTOKENBURNADDRESS(&_YBNegRiskAdapter.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.Admins(&_YBNegRiskAdapter.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.Admins(&_YBNegRiskAdapter.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "balanceOf", _owner, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.BalanceOf(&_YBNegRiskAdapter.CallOpts, _owner, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.BalanceOf(&_YBNegRiskAdapter.CallOpts, _owner, _id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) BalanceOfBatch(opts *bind.CallOpts, _owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "balanceOfBatch", _owners, _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _YBNegRiskAdapter.Contract.BalanceOfBatch(&_YBNegRiskAdapter.CallOpts, _owners, _ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _YBNegRiskAdapter.Contract.BalanceOfBatch(&_YBNegRiskAdapter.CallOpts, _owners, _ids)
}

// Col is a free data retrieval call binding the contract method 0xa78695b0.
//
// Solidity: function col() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) Col(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "col")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Col is a free data retrieval call binding the contract method 0xa78695b0.
//
// Solidity: function col() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) Col() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.Col(&_YBNegRiskAdapter.CallOpts)
}

// Col is a free data retrieval call binding the contract method 0xa78695b0.
//
// Solidity: function col() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) Col() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.Col(&_YBNegRiskAdapter.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) Ctf() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.Ctf(&_YBNegRiskAdapter.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) Ctf() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.Ctf(&_YBNegRiskAdapter.CallOpts)
}

// GetConditionId is a free data retrieval call binding the contract method 0x04329c03.
//
// Solidity: function getConditionId(bytes32 _questionId) view returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) GetConditionId(opts *bind.CallOpts, _questionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "getConditionId", _questionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConditionId is a free data retrieval call binding the contract method 0x04329c03.
//
// Solidity: function getConditionId(bytes32 _questionId) view returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) GetConditionId(_questionId [32]byte) ([32]byte, error) {
	return _YBNegRiskAdapter.Contract.GetConditionId(&_YBNegRiskAdapter.CallOpts, _questionId)
}

// GetConditionId is a free data retrieval call binding the contract method 0x04329c03.
//
// Solidity: function getConditionId(bytes32 _questionId) view returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) GetConditionId(_questionId [32]byte) ([32]byte, error) {
	return _YBNegRiskAdapter.Contract.GetConditionId(&_YBNegRiskAdapter.CallOpts, _questionId)
}

// GetDetermined is a free data retrieval call binding the contract method 0x7ae2e67b.
//
// Solidity: function getDetermined(bytes32 _marketId) view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) GetDetermined(opts *bind.CallOpts, _marketId [32]byte) (bool, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "getDetermined", _marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetDetermined is a free data retrieval call binding the contract method 0x7ae2e67b.
//
// Solidity: function getDetermined(bytes32 _marketId) view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) GetDetermined(_marketId [32]byte) (bool, error) {
	return _YBNegRiskAdapter.Contract.GetDetermined(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetDetermined is a free data retrieval call binding the contract method 0x7ae2e67b.
//
// Solidity: function getDetermined(bytes32 _marketId) view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) GetDetermined(_marketId [32]byte) (bool, error) {
	return _YBNegRiskAdapter.Contract.GetDetermined(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetFeeBips is a free data retrieval call binding the contract method 0x2582cb5e.
//
// Solidity: function getFeeBips(bytes32 _marketId) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) GetFeeBips(opts *bind.CallOpts, _marketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "getFeeBips", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBips is a free data retrieval call binding the contract method 0x2582cb5e.
//
// Solidity: function getFeeBips(bytes32 _marketId) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) GetFeeBips(_marketId [32]byte) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.GetFeeBips(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetFeeBips is a free data retrieval call binding the contract method 0x2582cb5e.
//
// Solidity: function getFeeBips(bytes32 _marketId) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) GetFeeBips(_marketId [32]byte) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.GetFeeBips(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetMarketData is a free data retrieval call binding the contract method 0x30f4f4bb.
//
// Solidity: function getMarketData(bytes32 _marketId) view returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) GetMarketData(opts *bind.CallOpts, _marketId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "getMarketData", _marketId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMarketData is a free data retrieval call binding the contract method 0x30f4f4bb.
//
// Solidity: function getMarketData(bytes32 _marketId) view returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) GetMarketData(_marketId [32]byte) ([32]byte, error) {
	return _YBNegRiskAdapter.Contract.GetMarketData(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetMarketData is a free data retrieval call binding the contract method 0x30f4f4bb.
//
// Solidity: function getMarketData(bytes32 _marketId) view returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) GetMarketData(_marketId [32]byte) ([32]byte, error) {
	return _YBNegRiskAdapter.Contract.GetMarketData(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _marketId) view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) GetOracle(opts *bind.CallOpts, _marketId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "getOracle", _marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _marketId) view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) GetOracle(_marketId [32]byte) (common.Address, error) {
	return _YBNegRiskAdapter.Contract.GetOracle(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _marketId) view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) GetOracle(_marketId [32]byte) (common.Address, error) {
	return _YBNegRiskAdapter.Contract.GetOracle(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetPositionId is a free data retrieval call binding the contract method 0x752b5ba5.
//
// Solidity: function getPositionId(bytes32 _questionId, bool _outcome) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) GetPositionId(opts *bind.CallOpts, _questionId [32]byte, _outcome bool) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "getPositionId", _questionId, _outcome)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionId is a free data retrieval call binding the contract method 0x752b5ba5.
//
// Solidity: function getPositionId(bytes32 _questionId, bool _outcome) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) GetPositionId(_questionId [32]byte, _outcome bool) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.GetPositionId(&_YBNegRiskAdapter.CallOpts, _questionId, _outcome)
}

// GetPositionId is a free data retrieval call binding the contract method 0x752b5ba5.
//
// Solidity: function getPositionId(bytes32 _questionId, bool _outcome) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) GetPositionId(_questionId [32]byte, _outcome bool) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.GetPositionId(&_YBNegRiskAdapter.CallOpts, _questionId, _outcome)
}

// GetQuestionCount is a free data retrieval call binding the contract method 0xb7f75d2c.
//
// Solidity: function getQuestionCount(bytes32 _marketId) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) GetQuestionCount(opts *bind.CallOpts, _marketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "getQuestionCount", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuestionCount is a free data retrieval call binding the contract method 0xb7f75d2c.
//
// Solidity: function getQuestionCount(bytes32 _marketId) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) GetQuestionCount(_marketId [32]byte) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.GetQuestionCount(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetQuestionCount is a free data retrieval call binding the contract method 0xb7f75d2c.
//
// Solidity: function getQuestionCount(bytes32 _marketId) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) GetQuestionCount(_marketId [32]byte) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.GetQuestionCount(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _marketId) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) GetResult(opts *bind.CallOpts, _marketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "getResult", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _marketId) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) GetResult(_marketId [32]byte) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.GetResult(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _marketId) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) GetResult(_marketId [32]byte) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.GetResult(&_YBNegRiskAdapter.CallOpts, _marketId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) IsAdmin(opts *bind.CallOpts, usr common.Address) (bool, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "isAdmin", usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) IsAdmin(usr common.Address) (bool, error) {
	return _YBNegRiskAdapter.Contract.IsAdmin(&_YBNegRiskAdapter.CallOpts, usr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) IsAdmin(usr common.Address) (bool, error) {
	return _YBNegRiskAdapter.Contract.IsAdmin(&_YBNegRiskAdapter.CallOpts, usr)
}

// IsConvertPositionsGated is a free data retrieval call binding the contract method 0x1fe2a01a.
//
// Solidity: function isConvertPositionsGated() view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) IsConvertPositionsGated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "isConvertPositionsGated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConvertPositionsGated is a free data retrieval call binding the contract method 0x1fe2a01a.
//
// Solidity: function isConvertPositionsGated() view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) IsConvertPositionsGated() (bool, error) {
	return _YBNegRiskAdapter.Contract.IsConvertPositionsGated(&_YBNegRiskAdapter.CallOpts)
}

// IsConvertPositionsGated is a free data retrieval call binding the contract method 0x1fe2a01a.
//
// Solidity: function isConvertPositionsGated() view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) IsConvertPositionsGated() (bool, error) {
	return _YBNegRiskAdapter.Contract.IsConvertPositionsGated(&_YBNegRiskAdapter.CallOpts)
}

// IsMergePositionsGated is a free data retrieval call binding the contract method 0x37d054bf.
//
// Solidity: function isMergePositionsGated() view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) IsMergePositionsGated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "isMergePositionsGated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMergePositionsGated is a free data retrieval call binding the contract method 0x37d054bf.
//
// Solidity: function isMergePositionsGated() view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) IsMergePositionsGated() (bool, error) {
	return _YBNegRiskAdapter.Contract.IsMergePositionsGated(&_YBNegRiskAdapter.CallOpts)
}

// IsMergePositionsGated is a free data retrieval call binding the contract method 0x37d054bf.
//
// Solidity: function isMergePositionsGated() view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) IsMergePositionsGated() (bool, error) {
	return _YBNegRiskAdapter.Contract.IsMergePositionsGated(&_YBNegRiskAdapter.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) IsOperator(opts *bind.CallOpts, usr common.Address) (bool, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "isOperator", usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) IsOperator(usr common.Address) (bool, error) {
	return _YBNegRiskAdapter.Contract.IsOperator(&_YBNegRiskAdapter.CallOpts, usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) IsOperator(usr common.Address) (bool, error) {
	return _YBNegRiskAdapter.Contract.IsOperator(&_YBNegRiskAdapter.CallOpts, usr)
}

// IsSplitPositionGated is a free data retrieval call binding the contract method 0x9d4913aa.
//
// Solidity: function isSplitPositionGated() view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) IsSplitPositionGated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "isSplitPositionGated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSplitPositionGated is a free data retrieval call binding the contract method 0x9d4913aa.
//
// Solidity: function isSplitPositionGated() view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) IsSplitPositionGated() (bool, error) {
	return _YBNegRiskAdapter.Contract.IsSplitPositionGated(&_YBNegRiskAdapter.CallOpts)
}

// IsSplitPositionGated is a free data retrieval call binding the contract method 0x9d4913aa.
//
// Solidity: function isSplitPositionGated() view returns(bool)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) IsSplitPositionGated() (bool, error) {
	return _YBNegRiskAdapter.Contract.IsSplitPositionGated(&_YBNegRiskAdapter.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.Operators(&_YBNegRiskAdapter.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _YBNegRiskAdapter.Contract.Operators(&_YBNegRiskAdapter.CallOpts, arg0)
}

// UserPermissions is a free data retrieval call binding the contract method 0x85fb9422.
//
// Solidity: function userPermissions(address ) view returns(bool canSplitPosition, bool canMergePositions, bool canConvertPositions)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) UserPermissions(opts *bind.CallOpts, arg0 common.Address) (struct {
	CanSplitPosition    bool
	CanMergePositions   bool
	CanConvertPositions bool
}, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "userPermissions", arg0)

	outstruct := new(struct {
		CanSplitPosition    bool
		CanMergePositions   bool
		CanConvertPositions bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CanSplitPosition = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CanMergePositions = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.CanConvertPositions = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// UserPermissions is a free data retrieval call binding the contract method 0x85fb9422.
//
// Solidity: function userPermissions(address ) view returns(bool canSplitPosition, bool canMergePositions, bool canConvertPositions)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) UserPermissions(arg0 common.Address) (struct {
	CanSplitPosition    bool
	CanMergePositions   bool
	CanConvertPositions bool
}, error) {
	return _YBNegRiskAdapter.Contract.UserPermissions(&_YBNegRiskAdapter.CallOpts, arg0)
}

// UserPermissions is a free data retrieval call binding the contract method 0x85fb9422.
//
// Solidity: function userPermissions(address ) view returns(bool canSplitPosition, bool canMergePositions, bool canConvertPositions)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) UserPermissions(arg0 common.Address) (struct {
	CanSplitPosition    bool
	CanMergePositions   bool
	CanConvertPositions bool
}, error) {
	return _YBNegRiskAdapter.Contract.UserPermissions(&_YBNegRiskAdapter.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) Vault() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.Vault(&_YBNegRiskAdapter.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) Vault() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.Vault(&_YBNegRiskAdapter.CallOpts)
}

// Wcol is a free data retrieval call binding the contract method 0x7e3b74c3.
//
// Solidity: function wcol() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCaller) Wcol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRiskAdapter.contract.Call(opts, &out, "wcol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wcol is a free data retrieval call binding the contract method 0x7e3b74c3.
//
// Solidity: function wcol() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) Wcol() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.Wcol(&_YBNegRiskAdapter.CallOpts)
}

// Wcol is a free data retrieval call binding the contract method 0x7e3b74c3.
//
// Solidity: function wcol() view returns(address)
func (_YBNegRiskAdapter *YBNegRiskAdapterCallerSession) Wcol() (common.Address, error) {
	return _YBNegRiskAdapter.Contract.Wcol(&_YBNegRiskAdapter.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) AddAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "addAdmin", admin_)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) AddAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.AddAdmin(&_YBNegRiskAdapter.TransactOpts, admin_)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) AddAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.AddAdmin(&_YBNegRiskAdapter.TransactOpts, admin_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) AddOperator(opts *bind.TransactOpts, operator_ common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "addOperator", operator_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) AddOperator(operator_ common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.AddOperator(&_YBNegRiskAdapter.TransactOpts, operator_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) AddOperator(operator_ common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.AddOperator(&_YBNegRiskAdapter.TransactOpts, operator_)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) ConvertPositions(opts *bind.TransactOpts, _marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "convertPositions", _marketId, _indexSet, _amount)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) ConvertPositions(_marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.ConvertPositions(&_YBNegRiskAdapter.TransactOpts, _marketId, _indexSet, _amount)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) ConvertPositions(_marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.ConvertPositions(&_YBNegRiskAdapter.TransactOpts, _marketId, _indexSet, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) MergePositions(opts *bind.TransactOpts, _collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "mergePositions", _collateralToken, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) MergePositions(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.MergePositions(&_YBNegRiskAdapter.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) MergePositions(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.MergePositions(&_YBNegRiskAdapter.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// MergePositions0 is a paid mutator transaction binding the contract method 0xb10c5c17.
//
// Solidity: function mergePositions(bytes32 _conditionId, uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) MergePositions0(opts *bind.TransactOpts, _conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "mergePositions0", _conditionId, _amount)
}

// MergePositions0 is a paid mutator transaction binding the contract method 0xb10c5c17.
//
// Solidity: function mergePositions(bytes32 _conditionId, uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) MergePositions0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.MergePositions0(&_YBNegRiskAdapter.TransactOpts, _conditionId, _amount)
}

// MergePositions0 is a paid mutator transaction binding the contract method 0xb10c5c17.
//
// Solidity: function mergePositions(bytes32 _conditionId, uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) MergePositions0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.MergePositions0(&_YBNegRiskAdapter.TransactOpts, _conditionId, _amount)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.OnERC1155BatchReceived(&_YBNegRiskAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.OnERC1155BatchReceived(&_YBNegRiskAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.OnERC1155Received(&_YBNegRiskAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.OnERC1155Received(&_YBNegRiskAdapter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PrepareMarket is a paid mutator transaction binding the contract method 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _metadata) returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) PrepareMarket(opts *bind.TransactOpts, _feeBips *big.Int, _metadata []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "prepareMarket", _feeBips, _metadata)
}

// PrepareMarket is a paid mutator transaction binding the contract method 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _metadata) returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) PrepareMarket(_feeBips *big.Int, _metadata []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.PrepareMarket(&_YBNegRiskAdapter.TransactOpts, _feeBips, _metadata)
}

// PrepareMarket is a paid mutator transaction binding the contract method 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _metadata) returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) PrepareMarket(_feeBips *big.Int, _metadata []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.PrepareMarket(&_YBNegRiskAdapter.TransactOpts, _feeBips, _metadata)
}

// PrepareQuestion is a paid mutator transaction binding the contract method 0x1d69b48d.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _metadata) returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) PrepareQuestion(opts *bind.TransactOpts, _marketId [32]byte, _metadata []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "prepareQuestion", _marketId, _metadata)
}

// PrepareQuestion is a paid mutator transaction binding the contract method 0x1d69b48d.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _metadata) returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) PrepareQuestion(_marketId [32]byte, _metadata []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.PrepareQuestion(&_YBNegRiskAdapter.TransactOpts, _marketId, _metadata)
}

// PrepareQuestion is a paid mutator transaction binding the contract method 0x1d69b48d.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _metadata) returns(bytes32)
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) PrepareQuestion(_marketId [32]byte, _metadata []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.PrepareQuestion(&_YBNegRiskAdapter.TransactOpts, _marketId, _metadata)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0xdbeccb23.
//
// Solidity: function redeemPositions(bytes32 _conditionId, uint256[] _amounts) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) RedeemPositions(opts *bind.TransactOpts, _conditionId [32]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "redeemPositions", _conditionId, _amounts)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0xdbeccb23.
//
// Solidity: function redeemPositions(bytes32 _conditionId, uint256[] _amounts) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) RedeemPositions(_conditionId [32]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RedeemPositions(&_YBNegRiskAdapter.TransactOpts, _conditionId, _amounts)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0xdbeccb23.
//
// Solidity: function redeemPositions(bytes32 _conditionId, uint256[] _amounts) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) RedeemPositions(_conditionId [32]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RedeemPositions(&_YBNegRiskAdapter.TransactOpts, _conditionId, _amounts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RemoveAdmin(&_YBNegRiskAdapter.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RemoveAdmin(&_YBNegRiskAdapter.TransactOpts, admin)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RemoveOperator(&_YBNegRiskAdapter.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RemoveOperator(&_YBNegRiskAdapter.TransactOpts, operator)
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) RenounceAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "renounceAdminRole")
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) RenounceAdminRole() (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RenounceAdminRole(&_YBNegRiskAdapter.TransactOpts)
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) RenounceAdminRole() (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RenounceAdminRole(&_YBNegRiskAdapter.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) RenounceOperatorRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "renounceOperatorRole")
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RenounceOperatorRole(&_YBNegRiskAdapter.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.RenounceOperatorRole(&_YBNegRiskAdapter.TransactOpts)
}

// ReportOutcome is a paid mutator transaction binding the contract method 0xe200affd.
//
// Solidity: function reportOutcome(bytes32 _questionId, bool _outcome) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) ReportOutcome(opts *bind.TransactOpts, _questionId [32]byte, _outcome bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "reportOutcome", _questionId, _outcome)
}

// ReportOutcome is a paid mutator transaction binding the contract method 0xe200affd.
//
// Solidity: function reportOutcome(bytes32 _questionId, bool _outcome) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) ReportOutcome(_questionId [32]byte, _outcome bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.ReportOutcome(&_YBNegRiskAdapter.TransactOpts, _questionId, _outcome)
}

// ReportOutcome is a paid mutator transaction binding the contract method 0xe200affd.
//
// Solidity: function reportOutcome(bytes32 _questionId, bool _outcome) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) ReportOutcome(_questionId [32]byte, _outcome bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.ReportOutcome(&_YBNegRiskAdapter.TransactOpts, _questionId, _outcome)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "safeTransferFrom", _from, _to, _id, _value, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SafeTransferFrom(&_YBNegRiskAdapter.TransactOpts, _from, _to, _id, _value, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SafeTransferFrom(&_YBNegRiskAdapter.TransactOpts, _from, _to, _id, _value, _data)
}

// SetIsConvertPositionsGated is a paid mutator transaction binding the contract method 0x70415370.
//
// Solidity: function setIsConvertPositionsGated(bool gated) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) SetIsConvertPositionsGated(opts *bind.TransactOpts, gated bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "setIsConvertPositionsGated", gated)
}

// SetIsConvertPositionsGated is a paid mutator transaction binding the contract method 0x70415370.
//
// Solidity: function setIsConvertPositionsGated(bool gated) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) SetIsConvertPositionsGated(gated bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SetIsConvertPositionsGated(&_YBNegRiskAdapter.TransactOpts, gated)
}

// SetIsConvertPositionsGated is a paid mutator transaction binding the contract method 0x70415370.
//
// Solidity: function setIsConvertPositionsGated(bool gated) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) SetIsConvertPositionsGated(gated bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SetIsConvertPositionsGated(&_YBNegRiskAdapter.TransactOpts, gated)
}

// SetIsMergePositionsGated is a paid mutator transaction binding the contract method 0x80d74ff9.
//
// Solidity: function setIsMergePositionsGated(bool gated) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) SetIsMergePositionsGated(opts *bind.TransactOpts, gated bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "setIsMergePositionsGated", gated)
}

// SetIsMergePositionsGated is a paid mutator transaction binding the contract method 0x80d74ff9.
//
// Solidity: function setIsMergePositionsGated(bool gated) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) SetIsMergePositionsGated(gated bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SetIsMergePositionsGated(&_YBNegRiskAdapter.TransactOpts, gated)
}

// SetIsMergePositionsGated is a paid mutator transaction binding the contract method 0x80d74ff9.
//
// Solidity: function setIsMergePositionsGated(bool gated) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) SetIsMergePositionsGated(gated bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SetIsMergePositionsGated(&_YBNegRiskAdapter.TransactOpts, gated)
}

// SetIsSplitPositionGated is a paid mutator transaction binding the contract method 0x5547ab74.
//
// Solidity: function setIsSplitPositionGated(bool gated) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) SetIsSplitPositionGated(opts *bind.TransactOpts, gated bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "setIsSplitPositionGated", gated)
}

// SetIsSplitPositionGated is a paid mutator transaction binding the contract method 0x5547ab74.
//
// Solidity: function setIsSplitPositionGated(bool gated) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) SetIsSplitPositionGated(gated bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SetIsSplitPositionGated(&_YBNegRiskAdapter.TransactOpts, gated)
}

// SetIsSplitPositionGated is a paid mutator transaction binding the contract method 0x5547ab74.
//
// Solidity: function setIsSplitPositionGated(bool gated) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) SetIsSplitPositionGated(gated bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SetIsSplitPositionGated(&_YBNegRiskAdapter.TransactOpts, gated)
}

// SetPermissions is a paid mutator transaction binding the contract method 0xae692682.
//
// Solidity: function setPermissions(address _user, bool _canSplitPosition, bool _canMergePositions, bool _canConvertPositions) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) SetPermissions(opts *bind.TransactOpts, _user common.Address, _canSplitPosition bool, _canMergePositions bool, _canConvertPositions bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "setPermissions", _user, _canSplitPosition, _canMergePositions, _canConvertPositions)
}

// SetPermissions is a paid mutator transaction binding the contract method 0xae692682.
//
// Solidity: function setPermissions(address _user, bool _canSplitPosition, bool _canMergePositions, bool _canConvertPositions) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) SetPermissions(_user common.Address, _canSplitPosition bool, _canMergePositions bool, _canConvertPositions bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SetPermissions(&_YBNegRiskAdapter.TransactOpts, _user, _canSplitPosition, _canMergePositions, _canConvertPositions)
}

// SetPermissions is a paid mutator transaction binding the contract method 0xae692682.
//
// Solidity: function setPermissions(address _user, bool _canSplitPosition, bool _canMergePositions, bool _canConvertPositions) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) SetPermissions(_user common.Address, _canSplitPosition bool, _canMergePositions bool, _canConvertPositions bool) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SetPermissions(&_YBNegRiskAdapter.TransactOpts, _user, _canSplitPosition, _canMergePositions, _canConvertPositions)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) SplitPosition(opts *bind.TransactOpts, _collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "splitPosition", _collateralToken, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) SplitPosition(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SplitPosition(&_YBNegRiskAdapter.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) SplitPosition(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SplitPosition(&_YBNegRiskAdapter.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// SplitPosition0 is a paid mutator transaction binding the contract method 0xa3d7da1d.
//
// Solidity: function splitPosition(bytes32 _conditionId, uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactor) SplitPosition0(opts *bind.TransactOpts, _conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.contract.Transact(opts, "splitPosition0", _conditionId, _amount)
}

// SplitPosition0 is a paid mutator transaction binding the contract method 0xa3d7da1d.
//
// Solidity: function splitPosition(bytes32 _conditionId, uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterSession) SplitPosition0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SplitPosition0(&_YBNegRiskAdapter.TransactOpts, _conditionId, _amount)
}

// SplitPosition0 is a paid mutator transaction binding the contract method 0xa3d7da1d.
//
// Solidity: function splitPosition(bytes32 _conditionId, uint256 _amount) returns()
func (_YBNegRiskAdapter *YBNegRiskAdapterTransactorSession) SplitPosition0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _YBNegRiskAdapter.Contract.SplitPosition0(&_YBNegRiskAdapter.TransactOpts, _conditionId, _amount)
}

// YBNegRiskAdapterIsConvertPositionsGatedUpdatedIterator is returned from FilterIsConvertPositionsGatedUpdated and is used to iterate over the raw logs and unpacked data for IsConvertPositionsGatedUpdated events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterIsConvertPositionsGatedUpdatedIterator struct {
	Event *YBNegRiskAdapterIsConvertPositionsGatedUpdated // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterIsConvertPositionsGatedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterIsConvertPositionsGatedUpdated)
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
		it.Event = new(YBNegRiskAdapterIsConvertPositionsGatedUpdated)
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
func (it *YBNegRiskAdapterIsConvertPositionsGatedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterIsConvertPositionsGatedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterIsConvertPositionsGatedUpdated represents a IsConvertPositionsGatedUpdated event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterIsConvertPositionsGatedUpdated struct {
	Gated bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIsConvertPositionsGatedUpdated is a free log retrieval operation binding the contract event 0x2d71a31d83aa6e18948020887263f570d642d621f8e94c6eb77cf8686e43c36f.
//
// Solidity: event IsConvertPositionsGatedUpdated(bool gated)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterIsConvertPositionsGatedUpdated(opts *bind.FilterOpts) (*YBNegRiskAdapterIsConvertPositionsGatedUpdatedIterator, error) {

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "IsConvertPositionsGatedUpdated")
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterIsConvertPositionsGatedUpdatedIterator{contract: _YBNegRiskAdapter.contract, event: "IsConvertPositionsGatedUpdated", logs: logs, sub: sub}, nil
}

// WatchIsConvertPositionsGatedUpdated is a free log subscription operation binding the contract event 0x2d71a31d83aa6e18948020887263f570d642d621f8e94c6eb77cf8686e43c36f.
//
// Solidity: event IsConvertPositionsGatedUpdated(bool gated)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchIsConvertPositionsGatedUpdated(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterIsConvertPositionsGatedUpdated) (event.Subscription, error) {

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "IsConvertPositionsGatedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterIsConvertPositionsGatedUpdated)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "IsConvertPositionsGatedUpdated", log); err != nil {
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

// ParseIsConvertPositionsGatedUpdated is a log parse operation binding the contract event 0x2d71a31d83aa6e18948020887263f570d642d621f8e94c6eb77cf8686e43c36f.
//
// Solidity: event IsConvertPositionsGatedUpdated(bool gated)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseIsConvertPositionsGatedUpdated(log types.Log) (*YBNegRiskAdapterIsConvertPositionsGatedUpdated, error) {
	event := new(YBNegRiskAdapterIsConvertPositionsGatedUpdated)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "IsConvertPositionsGatedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterIsMergePositionsGatedUpdatedIterator is returned from FilterIsMergePositionsGatedUpdated and is used to iterate over the raw logs and unpacked data for IsMergePositionsGatedUpdated events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterIsMergePositionsGatedUpdatedIterator struct {
	Event *YBNegRiskAdapterIsMergePositionsGatedUpdated // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterIsMergePositionsGatedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterIsMergePositionsGatedUpdated)
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
		it.Event = new(YBNegRiskAdapterIsMergePositionsGatedUpdated)
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
func (it *YBNegRiskAdapterIsMergePositionsGatedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterIsMergePositionsGatedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterIsMergePositionsGatedUpdated represents a IsMergePositionsGatedUpdated event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterIsMergePositionsGatedUpdated struct {
	Gated bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIsMergePositionsGatedUpdated is a free log retrieval operation binding the contract event 0x25a93d19f0753a683e9d386cb37d4b9b31cce0fffaa928711b9a7cd09ca0eb65.
//
// Solidity: event IsMergePositionsGatedUpdated(bool gated)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterIsMergePositionsGatedUpdated(opts *bind.FilterOpts) (*YBNegRiskAdapterIsMergePositionsGatedUpdatedIterator, error) {

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "IsMergePositionsGatedUpdated")
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterIsMergePositionsGatedUpdatedIterator{contract: _YBNegRiskAdapter.contract, event: "IsMergePositionsGatedUpdated", logs: logs, sub: sub}, nil
}

// WatchIsMergePositionsGatedUpdated is a free log subscription operation binding the contract event 0x25a93d19f0753a683e9d386cb37d4b9b31cce0fffaa928711b9a7cd09ca0eb65.
//
// Solidity: event IsMergePositionsGatedUpdated(bool gated)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchIsMergePositionsGatedUpdated(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterIsMergePositionsGatedUpdated) (event.Subscription, error) {

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "IsMergePositionsGatedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterIsMergePositionsGatedUpdated)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "IsMergePositionsGatedUpdated", log); err != nil {
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

// ParseIsMergePositionsGatedUpdated is a log parse operation binding the contract event 0x25a93d19f0753a683e9d386cb37d4b9b31cce0fffaa928711b9a7cd09ca0eb65.
//
// Solidity: event IsMergePositionsGatedUpdated(bool gated)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseIsMergePositionsGatedUpdated(log types.Log) (*YBNegRiskAdapterIsMergePositionsGatedUpdated, error) {
	event := new(YBNegRiskAdapterIsMergePositionsGatedUpdated)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "IsMergePositionsGatedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterIsSplitPositionGatedUpdatedIterator is returned from FilterIsSplitPositionGatedUpdated and is used to iterate over the raw logs and unpacked data for IsSplitPositionGatedUpdated events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterIsSplitPositionGatedUpdatedIterator struct {
	Event *YBNegRiskAdapterIsSplitPositionGatedUpdated // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterIsSplitPositionGatedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterIsSplitPositionGatedUpdated)
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
		it.Event = new(YBNegRiskAdapterIsSplitPositionGatedUpdated)
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
func (it *YBNegRiskAdapterIsSplitPositionGatedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterIsSplitPositionGatedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterIsSplitPositionGatedUpdated represents a IsSplitPositionGatedUpdated event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterIsSplitPositionGatedUpdated struct {
	Gated bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIsSplitPositionGatedUpdated is a free log retrieval operation binding the contract event 0x9b8cf6bb7e5e5e77eac6663deee1eb5c348d53be981ee51c355e6f215b9dbd0e.
//
// Solidity: event IsSplitPositionGatedUpdated(bool gated)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterIsSplitPositionGatedUpdated(opts *bind.FilterOpts) (*YBNegRiskAdapterIsSplitPositionGatedUpdatedIterator, error) {

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "IsSplitPositionGatedUpdated")
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterIsSplitPositionGatedUpdatedIterator{contract: _YBNegRiskAdapter.contract, event: "IsSplitPositionGatedUpdated", logs: logs, sub: sub}, nil
}

// WatchIsSplitPositionGatedUpdated is a free log subscription operation binding the contract event 0x9b8cf6bb7e5e5e77eac6663deee1eb5c348d53be981ee51c355e6f215b9dbd0e.
//
// Solidity: event IsSplitPositionGatedUpdated(bool gated)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchIsSplitPositionGatedUpdated(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterIsSplitPositionGatedUpdated) (event.Subscription, error) {

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "IsSplitPositionGatedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterIsSplitPositionGatedUpdated)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "IsSplitPositionGatedUpdated", log); err != nil {
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

// ParseIsSplitPositionGatedUpdated is a log parse operation binding the contract event 0x9b8cf6bb7e5e5e77eac6663deee1eb5c348d53be981ee51c355e6f215b9dbd0e.
//
// Solidity: event IsSplitPositionGatedUpdated(bool gated)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseIsSplitPositionGatedUpdated(log types.Log) (*YBNegRiskAdapterIsSplitPositionGatedUpdated, error) {
	event := new(YBNegRiskAdapterIsSplitPositionGatedUpdated)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "IsSplitPositionGatedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterMarketPreparedIterator is returned from FilterMarketPrepared and is used to iterate over the raw logs and unpacked data for MarketPrepared events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterMarketPreparedIterator struct {
	Event *YBNegRiskAdapterMarketPrepared // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterMarketPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterMarketPrepared)
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
		it.Event = new(YBNegRiskAdapterMarketPrepared)
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
func (it *YBNegRiskAdapterMarketPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterMarketPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterMarketPrepared represents a MarketPrepared event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterMarketPrepared struct {
	MarketId [32]byte
	Oracle   common.Address
	FeeBips  *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketPrepared is a free log retrieval operation binding the contract event 0xf059ab16d1ca60e123eab60e3c02b68faf060347c701a5d14885a8e1def7b3a8.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterMarketPrepared(opts *bind.FilterOpts, marketId [][32]byte, oracle []common.Address) (*YBNegRiskAdapterMarketPreparedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "MarketPrepared", marketIdRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterMarketPreparedIterator{contract: _YBNegRiskAdapter.contract, event: "MarketPrepared", logs: logs, sub: sub}, nil
}

// WatchMarketPrepared is a free log subscription operation binding the contract event 0xf059ab16d1ca60e123eab60e3c02b68faf060347c701a5d14885a8e1def7b3a8.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchMarketPrepared(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterMarketPrepared, marketId [][32]byte, oracle []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "MarketPrepared", marketIdRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterMarketPrepared)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "MarketPrepared", log); err != nil {
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

// ParseMarketPrepared is a log parse operation binding the contract event 0xf059ab16d1ca60e123eab60e3c02b68faf060347c701a5d14885a8e1def7b3a8.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseMarketPrepared(log types.Log) (*YBNegRiskAdapterMarketPrepared, error) {
	event := new(YBNegRiskAdapterMarketPrepared)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "MarketPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterNewAdminIterator struct {
	Event *YBNegRiskAdapterNewAdmin // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterNewAdmin)
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
		it.Event = new(YBNegRiskAdapterNewAdmin)
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
func (it *YBNegRiskAdapterNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterNewAdmin represents a NewAdmin event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterNewAdmin struct {
	NewAdminAddress common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterNewAdmin(opts *bind.FilterOpts, newAdminAddress []common.Address, admin []common.Address) (*YBNegRiskAdapterNewAdminIterator, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterNewAdminIterator{contract: _YBNegRiskAdapter.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterNewAdmin, newAdminAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterNewAdmin)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseNewAdmin(log types.Log) (*YBNegRiskAdapterNewAdmin, error) {
	event := new(YBNegRiskAdapterNewAdmin)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterNewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterNewOperatorIterator struct {
	Event *YBNegRiskAdapterNewOperator // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterNewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterNewOperator)
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
		it.Event = new(YBNegRiskAdapterNewOperator)
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
func (it *YBNegRiskAdapterNewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterNewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterNewOperator represents a NewOperator event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterNewOperator struct {
	NewOperatorAddress common.Address
	Admin              common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterNewOperator(opts *bind.FilterOpts, newOperatorAddress []common.Address, admin []common.Address) (*YBNegRiskAdapterNewOperatorIterator, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterNewOperatorIterator{contract: _YBNegRiskAdapter.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterNewOperator, newOperatorAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterNewOperator)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "NewOperator", log); err != nil {
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

// ParseNewOperator is a log parse operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseNewOperator(log types.Log) (*YBNegRiskAdapterNewOperator, error) {
	event := new(YBNegRiskAdapterNewOperator)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterOutcomeReportedIterator is returned from FilterOutcomeReported and is used to iterate over the raw logs and unpacked data for OutcomeReported events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterOutcomeReportedIterator struct {
	Event *YBNegRiskAdapterOutcomeReported // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterOutcomeReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterOutcomeReported)
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
		it.Event = new(YBNegRiskAdapterOutcomeReported)
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
func (it *YBNegRiskAdapterOutcomeReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterOutcomeReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterOutcomeReported represents a OutcomeReported event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterOutcomeReported struct {
	MarketId   [32]byte
	QuestionId [32]byte
	Outcome    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOutcomeReported is a free log retrieval operation binding the contract event 0x9e9fa7fd355160bd4cd3f22d4333519354beff1f5689bde87f2c5e63d8d484b2.
//
// Solidity: event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterOutcomeReported(opts *bind.FilterOpts, marketId [][32]byte, questionId [][32]byte) (*YBNegRiskAdapterOutcomeReportedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "OutcomeReported", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterOutcomeReportedIterator{contract: _YBNegRiskAdapter.contract, event: "OutcomeReported", logs: logs, sub: sub}, nil
}

// WatchOutcomeReported is a free log subscription operation binding the contract event 0x9e9fa7fd355160bd4cd3f22d4333519354beff1f5689bde87f2c5e63d8d484b2.
//
// Solidity: event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchOutcomeReported(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterOutcomeReported, marketId [][32]byte, questionId [][32]byte) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "OutcomeReported", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterOutcomeReported)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "OutcomeReported", log); err != nil {
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

// ParseOutcomeReported is a log parse operation binding the contract event 0x9e9fa7fd355160bd4cd3f22d4333519354beff1f5689bde87f2c5e63d8d484b2.
//
// Solidity: event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseOutcomeReported(log types.Log) (*YBNegRiskAdapterOutcomeReported, error) {
	event := new(YBNegRiskAdapterOutcomeReported)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "OutcomeReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterPayoutRedemptionIterator is returned from FilterPayoutRedemption and is used to iterate over the raw logs and unpacked data for PayoutRedemption events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPayoutRedemptionIterator struct {
	Event *YBNegRiskAdapterPayoutRedemption // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterPayoutRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterPayoutRedemption)
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
		it.Event = new(YBNegRiskAdapterPayoutRedemption)
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
func (it *YBNegRiskAdapterPayoutRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterPayoutRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterPayoutRedemption represents a PayoutRedemption event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPayoutRedemption struct {
	Redeemer    common.Address
	ConditionId [32]byte
	Amounts     []*big.Int
	Payout      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayoutRedemption is a free log retrieval operation binding the contract event 0x9140a6a270ef945260c03894b3c6b3b2695e9d5101feef0ff24fec960cfd3224.
//
// Solidity: event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterPayoutRedemption(opts *bind.FilterOpts, redeemer []common.Address, conditionId [][32]byte) (*YBNegRiskAdapterPayoutRedemptionIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "PayoutRedemption", redeemerRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterPayoutRedemptionIterator{contract: _YBNegRiskAdapter.contract, event: "PayoutRedemption", logs: logs, sub: sub}, nil
}

// WatchPayoutRedemption is a free log subscription operation binding the contract event 0x9140a6a270ef945260c03894b3c6b3b2695e9d5101feef0ff24fec960cfd3224.
//
// Solidity: event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchPayoutRedemption(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterPayoutRedemption, redeemer []common.Address, conditionId [][32]byte) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "PayoutRedemption", redeemerRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterPayoutRedemption)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
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

// ParsePayoutRedemption is a log parse operation binding the contract event 0x9140a6a270ef945260c03894b3c6b3b2695e9d5101feef0ff24fec960cfd3224.
//
// Solidity: event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParsePayoutRedemption(log types.Log) (*YBNegRiskAdapterPayoutRedemption, error) {
	event := new(YBNegRiskAdapterPayoutRedemption)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterPermissionsUpdatedIterator is returned from FilterPermissionsUpdated and is used to iterate over the raw logs and unpacked data for PermissionsUpdated events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPermissionsUpdatedIterator struct {
	Event *YBNegRiskAdapterPermissionsUpdated // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterPermissionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterPermissionsUpdated)
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
		it.Event = new(YBNegRiskAdapterPermissionsUpdated)
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
func (it *YBNegRiskAdapterPermissionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterPermissionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterPermissionsUpdated represents a PermissionsUpdated event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPermissionsUpdated struct {
	User                common.Address
	CanSplitPosition    bool
	CanMergePositions   bool
	CanConvertPositions bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPermissionsUpdated is a free log retrieval operation binding the contract event 0x335ef3609facbbdb9893b4c959c9371a213cf7615046e0b96af7fddde2269e9f.
//
// Solidity: event PermissionsUpdated(address indexed user, bool canSplitPosition, bool canMergePositions, bool canConvertPositions)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterPermissionsUpdated(opts *bind.FilterOpts, user []common.Address) (*YBNegRiskAdapterPermissionsUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "PermissionsUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterPermissionsUpdatedIterator{contract: _YBNegRiskAdapter.contract, event: "PermissionsUpdated", logs: logs, sub: sub}, nil
}

// WatchPermissionsUpdated is a free log subscription operation binding the contract event 0x335ef3609facbbdb9893b4c959c9371a213cf7615046e0b96af7fddde2269e9f.
//
// Solidity: event PermissionsUpdated(address indexed user, bool canSplitPosition, bool canMergePositions, bool canConvertPositions)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchPermissionsUpdated(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterPermissionsUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "PermissionsUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterPermissionsUpdated)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PermissionsUpdated", log); err != nil {
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

// ParsePermissionsUpdated is a log parse operation binding the contract event 0x335ef3609facbbdb9893b4c959c9371a213cf7615046e0b96af7fddde2269e9f.
//
// Solidity: event PermissionsUpdated(address indexed user, bool canSplitPosition, bool canMergePositions, bool canConvertPositions)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParsePermissionsUpdated(log types.Log) (*YBNegRiskAdapterPermissionsUpdated, error) {
	event := new(YBNegRiskAdapterPermissionsUpdated)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PermissionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterPositionSplitIterator is returned from FilterPositionSplit and is used to iterate over the raw logs and unpacked data for PositionSplit events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPositionSplitIterator struct {
	Event *YBNegRiskAdapterPositionSplit // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterPositionSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterPositionSplit)
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
		it.Event = new(YBNegRiskAdapterPositionSplit)
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
func (it *YBNegRiskAdapterPositionSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterPositionSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterPositionSplit represents a PositionSplit event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPositionSplit struct {
	Stakeholder common.Address
	ConditionId [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionSplit is a free log retrieval operation binding the contract event 0xbbed930dbfb7907ae2d60ddf78345610214f26419a0128df39b6cc3d9e5df9b0.
//
// Solidity: event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterPositionSplit(opts *bind.FilterOpts, stakeholder []common.Address, conditionId [][32]byte) (*YBNegRiskAdapterPositionSplitIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "PositionSplit", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterPositionSplitIterator{contract: _YBNegRiskAdapter.contract, event: "PositionSplit", logs: logs, sub: sub}, nil
}

// WatchPositionSplit is a free log subscription operation binding the contract event 0xbbed930dbfb7907ae2d60ddf78345610214f26419a0128df39b6cc3d9e5df9b0.
//
// Solidity: event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchPositionSplit(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterPositionSplit, stakeholder []common.Address, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "PositionSplit", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterPositionSplit)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PositionSplit", log); err != nil {
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

// ParsePositionSplit is a log parse operation binding the contract event 0xbbed930dbfb7907ae2d60ddf78345610214f26419a0128df39b6cc3d9e5df9b0.
//
// Solidity: event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParsePositionSplit(log types.Log) (*YBNegRiskAdapterPositionSplit, error) {
	event := new(YBNegRiskAdapterPositionSplit)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PositionSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterPositionsConvertedIterator is returned from FilterPositionsConverted and is used to iterate over the raw logs and unpacked data for PositionsConverted events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPositionsConvertedIterator struct {
	Event *YBNegRiskAdapterPositionsConverted // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterPositionsConvertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterPositionsConverted)
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
		it.Event = new(YBNegRiskAdapterPositionsConverted)
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
func (it *YBNegRiskAdapterPositionsConvertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterPositionsConvertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterPositionsConverted represents a PositionsConverted event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPositionsConverted struct {
	Stakeholder common.Address
	MarketId    [32]byte
	IndexSet    *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionsConverted is a free log retrieval operation binding the contract event 0xb03d19dddbc72a87e735ff0ea3b57bef133ebe44e1894284916a84044deb367e.
//
// Solidity: event PositionsConverted(address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterPositionsConverted(opts *bind.FilterOpts, stakeholder []common.Address, marketId [][32]byte, indexSet []*big.Int) (*YBNegRiskAdapterPositionsConvertedIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var indexSetRule []interface{}
	for _, indexSetItem := range indexSet {
		indexSetRule = append(indexSetRule, indexSetItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "PositionsConverted", stakeholderRule, marketIdRule, indexSetRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterPositionsConvertedIterator{contract: _YBNegRiskAdapter.contract, event: "PositionsConverted", logs: logs, sub: sub}, nil
}

// WatchPositionsConverted is a free log subscription operation binding the contract event 0xb03d19dddbc72a87e735ff0ea3b57bef133ebe44e1894284916a84044deb367e.
//
// Solidity: event PositionsConverted(address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchPositionsConverted(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterPositionsConverted, stakeholder []common.Address, marketId [][32]byte, indexSet []*big.Int) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var indexSetRule []interface{}
	for _, indexSetItem := range indexSet {
		indexSetRule = append(indexSetRule, indexSetItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "PositionsConverted", stakeholderRule, marketIdRule, indexSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterPositionsConverted)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PositionsConverted", log); err != nil {
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

// ParsePositionsConverted is a log parse operation binding the contract event 0xb03d19dddbc72a87e735ff0ea3b57bef133ebe44e1894284916a84044deb367e.
//
// Solidity: event PositionsConverted(address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParsePositionsConverted(log types.Log) (*YBNegRiskAdapterPositionsConverted, error) {
	event := new(YBNegRiskAdapterPositionsConverted)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PositionsConverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterPositionsMergeIterator is returned from FilterPositionsMerge and is used to iterate over the raw logs and unpacked data for PositionsMerge events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPositionsMergeIterator struct {
	Event *YBNegRiskAdapterPositionsMerge // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterPositionsMergeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterPositionsMerge)
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
		it.Event = new(YBNegRiskAdapterPositionsMerge)
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
func (it *YBNegRiskAdapterPositionsMergeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterPositionsMergeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterPositionsMerge represents a PositionsMerge event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterPositionsMerge struct {
	Stakeholder common.Address
	ConditionId [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionsMerge is a free log retrieval operation binding the contract event 0xba33ac50d8894676597e6e35dc09cff59854708b642cd069d21eb9c7ca072a04.
//
// Solidity: event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterPositionsMerge(opts *bind.FilterOpts, stakeholder []common.Address, conditionId [][32]byte) (*YBNegRiskAdapterPositionsMergeIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "PositionsMerge", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterPositionsMergeIterator{contract: _YBNegRiskAdapter.contract, event: "PositionsMerge", logs: logs, sub: sub}, nil
}

// WatchPositionsMerge is a free log subscription operation binding the contract event 0xba33ac50d8894676597e6e35dc09cff59854708b642cd069d21eb9c7ca072a04.
//
// Solidity: event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchPositionsMerge(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterPositionsMerge, stakeholder []common.Address, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "PositionsMerge", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterPositionsMerge)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
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

// ParsePositionsMerge is a log parse operation binding the contract event 0xba33ac50d8894676597e6e35dc09cff59854708b642cd069d21eb9c7ca072a04.
//
// Solidity: event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParsePositionsMerge(log types.Log) (*YBNegRiskAdapterPositionsMerge, error) {
	event := new(YBNegRiskAdapterPositionsMerge)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterQuestionPreparedIterator is returned from FilterQuestionPrepared and is used to iterate over the raw logs and unpacked data for QuestionPrepared events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterQuestionPreparedIterator struct {
	Event *YBNegRiskAdapterQuestionPrepared // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterQuestionPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterQuestionPrepared)
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
		it.Event = new(YBNegRiskAdapterQuestionPrepared)
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
func (it *YBNegRiskAdapterQuestionPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterQuestionPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterQuestionPrepared represents a QuestionPrepared event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterQuestionPrepared struct {
	MarketId   [32]byte
	QuestionId [32]byte
	Index      *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionPrepared is a free log retrieval operation binding the contract event 0xaac410f87d423a922a7b226ac68f0c2eaf5bf6d15e644ac0758c7f96e2c253f7.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterQuestionPrepared(opts *bind.FilterOpts, marketId [][32]byte, questionId [][32]byte) (*YBNegRiskAdapterQuestionPreparedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "QuestionPrepared", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterQuestionPreparedIterator{contract: _YBNegRiskAdapter.contract, event: "QuestionPrepared", logs: logs, sub: sub}, nil
}

// WatchQuestionPrepared is a free log subscription operation binding the contract event 0xaac410f87d423a922a7b226ac68f0c2eaf5bf6d15e644ac0758c7f96e2c253f7.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchQuestionPrepared(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterQuestionPrepared, marketId [][32]byte, questionId [][32]byte) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "QuestionPrepared", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterQuestionPrepared)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "QuestionPrepared", log); err != nil {
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

// ParseQuestionPrepared is a log parse operation binding the contract event 0xaac410f87d423a922a7b226ac68f0c2eaf5bf6d15e644ac0758c7f96e2c253f7.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseQuestionPrepared(log types.Log) (*YBNegRiskAdapterQuestionPrepared, error) {
	event := new(YBNegRiskAdapterQuestionPrepared)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "QuestionPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterRemovedAdminIterator struct {
	Event *YBNegRiskAdapterRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterRemovedAdmin)
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
		it.Event = new(YBNegRiskAdapterRemovedAdmin)
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
func (it *YBNegRiskAdapterRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterRemovedAdmin represents a RemovedAdmin event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterRemovedAdmin struct {
	RemovedAdmin common.Address
	Admin        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, removedAdmin []common.Address, admin []common.Address) (*YBNegRiskAdapterRemovedAdminIterator, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterRemovedAdminIterator{contract: _YBNegRiskAdapter.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterRemovedAdmin, removedAdmin []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterRemovedAdmin)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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

// ParseRemovedAdmin is a log parse operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseRemovedAdmin(log types.Log) (*YBNegRiskAdapterRemovedAdmin, error) {
	event := new(YBNegRiskAdapterRemovedAdmin)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskAdapterRemovedOperatorIterator is returned from FilterRemovedOperator and is used to iterate over the raw logs and unpacked data for RemovedOperator events raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterRemovedOperatorIterator struct {
	Event *YBNegRiskAdapterRemovedOperator // Event containing the contract specifics and raw log

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
func (it *YBNegRiskAdapterRemovedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskAdapterRemovedOperator)
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
		it.Event = new(YBNegRiskAdapterRemovedOperator)
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
func (it *YBNegRiskAdapterRemovedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskAdapterRemovedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskAdapterRemovedOperator represents a RemovedOperator event raised by the YBNegRiskAdapter contract.
type YBNegRiskAdapterRemovedOperator struct {
	RemovedOperator common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRemovedOperator is a free log retrieval operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) FilterRemovedOperator(opts *bind.FilterOpts, removedOperator []common.Address, admin []common.Address) (*YBNegRiskAdapterRemovedOperatorIterator, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.FilterLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskAdapterRemovedOperatorIterator{contract: _YBNegRiskAdapter.contract, event: "RemovedOperator", logs: logs, sub: sub}, nil
}

// WatchRemovedOperator is a free log subscription operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) WatchRemovedOperator(opts *bind.WatchOpts, sink chan<- *YBNegRiskAdapterRemovedOperator, removedOperator []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRiskAdapter.contract.WatchLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskAdapterRemovedOperator)
				if err := _YBNegRiskAdapter.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
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

// ParseRemovedOperator is a log parse operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_YBNegRiskAdapter *YBNegRiskAdapterFilterer) ParseRemovedOperator(log types.Log) (*YBNegRiskAdapterRemovedOperator, error) {
	event := new(YBNegRiskAdapterRemovedOperator)
	if err := _YBNegRiskAdapter.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
