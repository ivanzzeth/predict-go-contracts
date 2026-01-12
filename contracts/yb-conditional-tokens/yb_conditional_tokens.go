// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yb_conditional_tokens

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

// YBConditionalTokensMetaData contains all meta data concerning the YBConditionalTokens contract.
var YBConditionalTokensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Insolvent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUnderlyingToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoYield\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoYieldBearingToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotYieldBearingToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnderlyingTokenAlreadyConnected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnderlyingTokenAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnderlyingTokenAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"err\",\"type\":\"uint256\"}],\"name\":\"VTokenCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VTokenTreasuryPercentMustBeZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"ConditionPreparation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"ConditionResolution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"IsMergePositionsGated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"IsSplitPositionGated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"IsTransferAllowedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"indexSets\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"}],\"name\":\"PayoutRedemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"partition\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionSplit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"partition\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionsMerge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"TransferControlEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"UnderlyingTokenConnectedToVToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"yieldRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldClaimed\",\"type\":\"uint256\"}],\"name\":\"UnderlyingTokenDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnderlyingTokenEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"name\":\"UnderlyingTokenRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"}],\"name\":\"VTokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"}],\"name\":\"YieldClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MERGE_POSITIONS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPLIT_POSITION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"YIELD_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimYield\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"connectVTokenToUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"depositedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"yieldRecipient\",\"type\":\"address\"}],\"name\":\"disableUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"enableUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"indexSet\",\"type\":\"uint256\"}],\"name\":\"getCollectionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"getConditionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"getOutcomeSlotCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"collectionId\",\"type\":\"bytes32\"}],\"name\":\"getPositionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMergePositionsGated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSplitPositionGated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isTransferAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"partition\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mergePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"payoutDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payoutNumerators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"prepareCondition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"indexSets\",\"type\":\"uint256[]\"}],\"name\":\"redeemPositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"questionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"payouts\",\"type\":\"uint256[]\"}],\"name\":\"reportPayouts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"setIsMergePositionsGated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"gated\",\"type\":\"bool\"}],\"name\":\"setIsSplitPositionGated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setTransferControlEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"partition\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"splitPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"splitPrincipalAndYield\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"yield\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferControlEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"underlyingIsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"underlyingToVToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"updateIsTransferAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YBConditionalTokensABI is the input ABI used to generate the binding from.
// Deprecated: Use YBConditionalTokensMetaData.ABI instead.
var YBConditionalTokensABI = YBConditionalTokensMetaData.ABI

// YBConditionalTokens is an auto generated Go binding around an Ethereum contract.
type YBConditionalTokens struct {
	YBConditionalTokensCaller     // Read-only binding to the contract
	YBConditionalTokensTransactor // Write-only binding to the contract
	YBConditionalTokensFilterer   // Log filterer for contract events
}

// YBConditionalTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type YBConditionalTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBConditionalTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YBConditionalTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBConditionalTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YBConditionalTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBConditionalTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YBConditionalTokensSession struct {
	Contract     *YBConditionalTokens // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// YBConditionalTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YBConditionalTokensCallerSession struct {
	Contract *YBConditionalTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// YBConditionalTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YBConditionalTokensTransactorSession struct {
	Contract     *YBConditionalTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// YBConditionalTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type YBConditionalTokensRaw struct {
	Contract *YBConditionalTokens // Generic contract binding to access the raw methods on
}

// YBConditionalTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YBConditionalTokensCallerRaw struct {
	Contract *YBConditionalTokensCaller // Generic read-only contract binding to access the raw methods on
}

// YBConditionalTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YBConditionalTokensTransactorRaw struct {
	Contract *YBConditionalTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYBConditionalTokens creates a new instance of YBConditionalTokens, bound to a specific deployed contract.
func NewYBConditionalTokens(address common.Address, backend bind.ContractBackend) (*YBConditionalTokens, error) {
	contract, err := bindYBConditionalTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokens{YBConditionalTokensCaller: YBConditionalTokensCaller{contract: contract}, YBConditionalTokensTransactor: YBConditionalTokensTransactor{contract: contract}, YBConditionalTokensFilterer: YBConditionalTokensFilterer{contract: contract}}, nil
}

// NewYBConditionalTokensCaller creates a new read-only instance of YBConditionalTokens, bound to a specific deployed contract.
func NewYBConditionalTokensCaller(address common.Address, caller bind.ContractCaller) (*YBConditionalTokensCaller, error) {
	contract, err := bindYBConditionalTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensCaller{contract: contract}, nil
}

// NewYBConditionalTokensTransactor creates a new write-only instance of YBConditionalTokens, bound to a specific deployed contract.
func NewYBConditionalTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*YBConditionalTokensTransactor, error) {
	contract, err := bindYBConditionalTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensTransactor{contract: contract}, nil
}

// NewYBConditionalTokensFilterer creates a new log filterer instance of YBConditionalTokens, bound to a specific deployed contract.
func NewYBConditionalTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*YBConditionalTokensFilterer, error) {
	contract, err := bindYBConditionalTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensFilterer{contract: contract}, nil
}

// bindYBConditionalTokens binds a generic wrapper to an already deployed contract.
func bindYBConditionalTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YBConditionalTokensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YBConditionalTokens *YBConditionalTokensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YBConditionalTokens.Contract.YBConditionalTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YBConditionalTokens *YBConditionalTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.YBConditionalTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YBConditionalTokens *YBConditionalTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.YBConditionalTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YBConditionalTokens *YBConditionalTokensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YBConditionalTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YBConditionalTokens *YBConditionalTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YBConditionalTokens *YBConditionalTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _YBConditionalTokens.Contract.DEFAULTADMINROLE(&_YBConditionalTokens.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _YBConditionalTokens.Contract.DEFAULTADMINROLE(&_YBConditionalTokens.CallOpts)
}

// MERGEPOSITIONSROLE is a free data retrieval call binding the contract method 0x7999c003.
//
// Solidity: function MERGE_POSITIONS_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCaller) MERGEPOSITIONSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "MERGE_POSITIONS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MERGEPOSITIONSROLE is a free data retrieval call binding the contract method 0x7999c003.
//
// Solidity: function MERGE_POSITIONS_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensSession) MERGEPOSITIONSROLE() ([32]byte, error) {
	return _YBConditionalTokens.Contract.MERGEPOSITIONSROLE(&_YBConditionalTokens.CallOpts)
}

// MERGEPOSITIONSROLE is a free data retrieval call binding the contract method 0x7999c003.
//
// Solidity: function MERGE_POSITIONS_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) MERGEPOSITIONSROLE() ([32]byte, error) {
	return _YBConditionalTokens.Contract.MERGEPOSITIONSROLE(&_YBConditionalTokens.CallOpts)
}

// SPLITPOSITIONROLE is a free data retrieval call binding the contract method 0x84c4798b.
//
// Solidity: function SPLIT_POSITION_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCaller) SPLITPOSITIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "SPLIT_POSITION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SPLITPOSITIONROLE is a free data retrieval call binding the contract method 0x84c4798b.
//
// Solidity: function SPLIT_POSITION_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensSession) SPLITPOSITIONROLE() ([32]byte, error) {
	return _YBConditionalTokens.Contract.SPLITPOSITIONROLE(&_YBConditionalTokens.CallOpts)
}

// SPLITPOSITIONROLE is a free data retrieval call binding the contract method 0x84c4798b.
//
// Solidity: function SPLIT_POSITION_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) SPLITPOSITIONROLE() ([32]byte, error) {
	return _YBConditionalTokens.Contract.SPLITPOSITIONROLE(&_YBConditionalTokens.CallOpts)
}

// YIELDMANAGERROLE is a free data retrieval call binding the contract method 0xfe9e9640.
//
// Solidity: function YIELD_MANAGER_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCaller) YIELDMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "YIELD_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// YIELDMANAGERROLE is a free data retrieval call binding the contract method 0xfe9e9640.
//
// Solidity: function YIELD_MANAGER_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensSession) YIELDMANAGERROLE() ([32]byte, error) {
	return _YBConditionalTokens.Contract.YIELDMANAGERROLE(&_YBConditionalTokens.CallOpts)
}

// YIELDMANAGERROLE is a free data retrieval call binding the contract method 0xfe9e9640.
//
// Solidity: function YIELD_MANAGER_ROLE() view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) YIELDMANAGERROLE() ([32]byte, error) {
	return _YBConditionalTokens.Contract.YIELDMANAGERROLE(&_YBConditionalTokens.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _YBConditionalTokens.Contract.BalanceOf(&_YBConditionalTokens.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _YBConditionalTokens.Contract.BalanceOf(&_YBConditionalTokens.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_YBConditionalTokens *YBConditionalTokensCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_YBConditionalTokens *YBConditionalTokensSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _YBConditionalTokens.Contract.BalanceOfBatch(&_YBConditionalTokens.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_YBConditionalTokens *YBConditionalTokensCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _YBConditionalTokens.Contract.BalanceOfBatch(&_YBConditionalTokens.CallOpts, accounts, ids)
}

// DepositedAmount is a free data retrieval call binding the contract method 0x4a4643f7.
//
// Solidity: function depositedAmount(address underlying) view returns(uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensCaller) DepositedAmount(opts *bind.CallOpts, underlying common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "depositedAmount", underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositedAmount is a free data retrieval call binding the contract method 0x4a4643f7.
//
// Solidity: function depositedAmount(address underlying) view returns(uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensSession) DepositedAmount(underlying common.Address) (*big.Int, error) {
	return _YBConditionalTokens.Contract.DepositedAmount(&_YBConditionalTokens.CallOpts, underlying)
}

// DepositedAmount is a free data retrieval call binding the contract method 0x4a4643f7.
//
// Solidity: function depositedAmount(address underlying) view returns(uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) DepositedAmount(underlying common.Address) (*big.Int, error) {
	return _YBConditionalTokens.Contract.DepositedAmount(&_YBConditionalTokens.CallOpts, underlying)
}

// ExpScale is a free data retrieval call binding the contract method 0x69aa3ac6.
//
// Solidity: function expScale() view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCaller) ExpScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "expScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpScale is a free data retrieval call binding the contract method 0x69aa3ac6.
//
// Solidity: function expScale() view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensSession) ExpScale() (*big.Int, error) {
	return _YBConditionalTokens.Contract.ExpScale(&_YBConditionalTokens.CallOpts)
}

// ExpScale is a free data retrieval call binding the contract method 0x69aa3ac6.
//
// Solidity: function expScale() view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) ExpScale() (*big.Int, error) {
	return _YBConditionalTokens.Contract.ExpScale(&_YBConditionalTokens.CallOpts)
}

// GetCollectionId is a free data retrieval call binding the contract method 0x856296f7.
//
// Solidity: function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet) view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCaller) GetCollectionId(opts *bind.CallOpts, parentCollectionId [32]byte, conditionId [32]byte, indexSet *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "getCollectionId", parentCollectionId, conditionId, indexSet)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCollectionId is a free data retrieval call binding the contract method 0x856296f7.
//
// Solidity: function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet) view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensSession) GetCollectionId(parentCollectionId [32]byte, conditionId [32]byte, indexSet *big.Int) ([32]byte, error) {
	return _YBConditionalTokens.Contract.GetCollectionId(&_YBConditionalTokens.CallOpts, parentCollectionId, conditionId, indexSet)
}

// GetCollectionId is a free data retrieval call binding the contract method 0x856296f7.
//
// Solidity: function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet) view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) GetCollectionId(parentCollectionId [32]byte, conditionId [32]byte, indexSet *big.Int) ([32]byte, error) {
	return _YBConditionalTokens.Contract.GetCollectionId(&_YBConditionalTokens.CallOpts, parentCollectionId, conditionId, indexSet)
}

// GetConditionId is a free data retrieval call binding the contract method 0x852c6ae2.
//
// Solidity: function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCaller) GetConditionId(opts *bind.CallOpts, oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "getConditionId", oracle, questionId, outcomeSlotCount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConditionId is a free data retrieval call binding the contract method 0x852c6ae2.
//
// Solidity: function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensSession) GetConditionId(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([32]byte, error) {
	return _YBConditionalTokens.Contract.GetConditionId(&_YBConditionalTokens.CallOpts, oracle, questionId, outcomeSlotCount)
}

// GetConditionId is a free data retrieval call binding the contract method 0x852c6ae2.
//
// Solidity: function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) GetConditionId(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([32]byte, error) {
	return _YBConditionalTokens.Contract.GetConditionId(&_YBConditionalTokens.CallOpts, oracle, questionId, outcomeSlotCount)
}

// GetOutcomeSlotCount is a free data retrieval call binding the contract method 0xd42dc0c2.
//
// Solidity: function getOutcomeSlotCount(bytes32 conditionId) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCaller) GetOutcomeSlotCount(opts *bind.CallOpts, conditionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "getOutcomeSlotCount", conditionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOutcomeSlotCount is a free data retrieval call binding the contract method 0xd42dc0c2.
//
// Solidity: function getOutcomeSlotCount(bytes32 conditionId) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensSession) GetOutcomeSlotCount(conditionId [32]byte) (*big.Int, error) {
	return _YBConditionalTokens.Contract.GetOutcomeSlotCount(&_YBConditionalTokens.CallOpts, conditionId)
}

// GetOutcomeSlotCount is a free data retrieval call binding the contract method 0xd42dc0c2.
//
// Solidity: function getOutcomeSlotCount(bytes32 conditionId) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) GetOutcomeSlotCount(conditionId [32]byte) (*big.Int, error) {
	return _YBConditionalTokens.Contract.GetOutcomeSlotCount(&_YBConditionalTokens.CallOpts, conditionId)
}

// GetPositionId is a free data retrieval call binding the contract method 0x39dd7530.
//
// Solidity: function getPositionId(address collateralToken, bytes32 collectionId) pure returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCaller) GetPositionId(opts *bind.CallOpts, collateralToken common.Address, collectionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "getPositionId", collateralToken, collectionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionId is a free data retrieval call binding the contract method 0x39dd7530.
//
// Solidity: function getPositionId(address collateralToken, bytes32 collectionId) pure returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensSession) GetPositionId(collateralToken common.Address, collectionId [32]byte) (*big.Int, error) {
	return _YBConditionalTokens.Contract.GetPositionId(&_YBConditionalTokens.CallOpts, collateralToken, collectionId)
}

// GetPositionId is a free data retrieval call binding the contract method 0x39dd7530.
//
// Solidity: function getPositionId(address collateralToken, bytes32 collectionId) pure returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) GetPositionId(collateralToken common.Address, collectionId [32]byte) (*big.Int, error) {
	return _YBConditionalTokens.Contract.GetPositionId(&_YBConditionalTokens.CallOpts, collateralToken, collectionId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _YBConditionalTokens.Contract.GetRoleAdmin(&_YBConditionalTokens.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _YBConditionalTokens.Contract.GetRoleAdmin(&_YBConditionalTokens.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _YBConditionalTokens.Contract.HasRole(&_YBConditionalTokens.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _YBConditionalTokens.Contract.HasRole(&_YBConditionalTokens.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _YBConditionalTokens.Contract.IsApprovedForAll(&_YBConditionalTokens.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _YBConditionalTokens.Contract.IsApprovedForAll(&_YBConditionalTokens.CallOpts, account, operator)
}

// IsMergePositionsGated is a free data retrieval call binding the contract method 0x37d054bf.
//
// Solidity: function isMergePositionsGated() view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCaller) IsMergePositionsGated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "isMergePositionsGated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMergePositionsGated is a free data retrieval call binding the contract method 0x37d054bf.
//
// Solidity: function isMergePositionsGated() view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensSession) IsMergePositionsGated() (bool, error) {
	return _YBConditionalTokens.Contract.IsMergePositionsGated(&_YBConditionalTokens.CallOpts)
}

// IsMergePositionsGated is a free data retrieval call binding the contract method 0x37d054bf.
//
// Solidity: function isMergePositionsGated() view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) IsMergePositionsGated() (bool, error) {
	return _YBConditionalTokens.Contract.IsMergePositionsGated(&_YBConditionalTokens.CallOpts)
}

// IsSplitPositionGated is a free data retrieval call binding the contract method 0x9d4913aa.
//
// Solidity: function isSplitPositionGated() view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCaller) IsSplitPositionGated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "isSplitPositionGated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSplitPositionGated is a free data retrieval call binding the contract method 0x9d4913aa.
//
// Solidity: function isSplitPositionGated() view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensSession) IsSplitPositionGated() (bool, error) {
	return _YBConditionalTokens.Contract.IsSplitPositionGated(&_YBConditionalTokens.CallOpts)
}

// IsSplitPositionGated is a free data retrieval call binding the contract method 0x9d4913aa.
//
// Solidity: function isSplitPositionGated() view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) IsSplitPositionGated() (bool, error) {
	return _YBConditionalTokens.Contract.IsSplitPositionGated(&_YBConditionalTokens.CallOpts)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x8822048e.
//
// Solidity: function isTransferAllowed(address addr) view returns(bool isAllowed)
func (_YBConditionalTokens *YBConditionalTokensCaller) IsTransferAllowed(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "isTransferAllowed", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x8822048e.
//
// Solidity: function isTransferAllowed(address addr) view returns(bool isAllowed)
func (_YBConditionalTokens *YBConditionalTokensSession) IsTransferAllowed(addr common.Address) (bool, error) {
	return _YBConditionalTokens.Contract.IsTransferAllowed(&_YBConditionalTokens.CallOpts, addr)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x8822048e.
//
// Solidity: function isTransferAllowed(address addr) view returns(bool isAllowed)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) IsTransferAllowed(addr common.Address) (bool, error) {
	return _YBConditionalTokens.Contract.IsTransferAllowed(&_YBConditionalTokens.CallOpts, addr)
}

// PayoutDenominator is a free data retrieval call binding the contract method 0xdd34de67.
//
// Solidity: function payoutDenominator(bytes32 ) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCaller) PayoutDenominator(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "payoutDenominator", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutDenominator is a free data retrieval call binding the contract method 0xdd34de67.
//
// Solidity: function payoutDenominator(bytes32 ) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensSession) PayoutDenominator(arg0 [32]byte) (*big.Int, error) {
	return _YBConditionalTokens.Contract.PayoutDenominator(&_YBConditionalTokens.CallOpts, arg0)
}

// PayoutDenominator is a free data retrieval call binding the contract method 0xdd34de67.
//
// Solidity: function payoutDenominator(bytes32 ) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) PayoutDenominator(arg0 [32]byte) (*big.Int, error) {
	return _YBConditionalTokens.Contract.PayoutDenominator(&_YBConditionalTokens.CallOpts, arg0)
}

// PayoutNumerators is a free data retrieval call binding the contract method 0x0504c814.
//
// Solidity: function payoutNumerators(bytes32 , uint256 ) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCaller) PayoutNumerators(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "payoutNumerators", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutNumerators is a free data retrieval call binding the contract method 0x0504c814.
//
// Solidity: function payoutNumerators(bytes32 , uint256 ) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensSession) PayoutNumerators(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _YBConditionalTokens.Contract.PayoutNumerators(&_YBConditionalTokens.CallOpts, arg0, arg1)
}

// PayoutNumerators is a free data retrieval call binding the contract method 0x0504c814.
//
// Solidity: function payoutNumerators(bytes32 , uint256 ) view returns(uint256)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) PayoutNumerators(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _YBConditionalTokens.Contract.PayoutNumerators(&_YBConditionalTokens.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _YBConditionalTokens.Contract.SupportsInterface(&_YBConditionalTokens.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _YBConditionalTokens.Contract.SupportsInterface(&_YBConditionalTokens.CallOpts, interfaceId)
}

// TransferControlEnabled is a free data retrieval call binding the contract method 0xfd370543.
//
// Solidity: function transferControlEnabled() view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCaller) TransferControlEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "transferControlEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferControlEnabled is a free data retrieval call binding the contract method 0xfd370543.
//
// Solidity: function transferControlEnabled() view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensSession) TransferControlEnabled() (bool, error) {
	return _YBConditionalTokens.Contract.TransferControlEnabled(&_YBConditionalTokens.CallOpts)
}

// TransferControlEnabled is a free data retrieval call binding the contract method 0xfd370543.
//
// Solidity: function transferControlEnabled() view returns(bool)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) TransferControlEnabled() (bool, error) {
	return _YBConditionalTokens.Contract.TransferControlEnabled(&_YBConditionalTokens.CallOpts)
}

// UnderlyingIsEnabled is a free data retrieval call binding the contract method 0x87aea3ad.
//
// Solidity: function underlyingIsEnabled(address underlying) view returns(bool isEnabled)
func (_YBConditionalTokens *YBConditionalTokensCaller) UnderlyingIsEnabled(opts *bind.CallOpts, underlying common.Address) (bool, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "underlyingIsEnabled", underlying)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnderlyingIsEnabled is a free data retrieval call binding the contract method 0x87aea3ad.
//
// Solidity: function underlyingIsEnabled(address underlying) view returns(bool isEnabled)
func (_YBConditionalTokens *YBConditionalTokensSession) UnderlyingIsEnabled(underlying common.Address) (bool, error) {
	return _YBConditionalTokens.Contract.UnderlyingIsEnabled(&_YBConditionalTokens.CallOpts, underlying)
}

// UnderlyingIsEnabled is a free data retrieval call binding the contract method 0x87aea3ad.
//
// Solidity: function underlyingIsEnabled(address underlying) view returns(bool isEnabled)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) UnderlyingIsEnabled(underlying common.Address) (bool, error) {
	return _YBConditionalTokens.Contract.UnderlyingIsEnabled(&_YBConditionalTokens.CallOpts, underlying)
}

// UnderlyingToVToken is a free data retrieval call binding the contract method 0x7feaa2ac.
//
// Solidity: function underlyingToVToken(address underlying) view returns(address vToken)
func (_YBConditionalTokens *YBConditionalTokensCaller) UnderlyingToVToken(opts *bind.CallOpts, underlying common.Address) (common.Address, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "underlyingToVToken", underlying)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToVToken is a free data retrieval call binding the contract method 0x7feaa2ac.
//
// Solidity: function underlyingToVToken(address underlying) view returns(address vToken)
func (_YBConditionalTokens *YBConditionalTokensSession) UnderlyingToVToken(underlying common.Address) (common.Address, error) {
	return _YBConditionalTokens.Contract.UnderlyingToVToken(&_YBConditionalTokens.CallOpts, underlying)
}

// UnderlyingToVToken is a free data retrieval call binding the contract method 0x7feaa2ac.
//
// Solidity: function underlyingToVToken(address underlying) view returns(address vToken)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) UnderlyingToVToken(underlying common.Address) (common.Address, error) {
	return _YBConditionalTokens.Contract.UnderlyingToVToken(&_YBConditionalTokens.CallOpts, underlying)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_YBConditionalTokens *YBConditionalTokensCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _YBConditionalTokens.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_YBConditionalTokens *YBConditionalTokensSession) Uri(arg0 *big.Int) (string, error) {
	return _YBConditionalTokens.Contract.Uri(&_YBConditionalTokens.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_YBConditionalTokens *YBConditionalTokensCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _YBConditionalTokens.Contract.Uri(&_YBConditionalTokens.CallOpts, arg0)
}

// ClaimYield is a paid mutator transaction binding the contract method 0xf1b0e8eb.
//
// Solidity: function claimYield(address underlying, uint256 vTokenAmount, address recipient) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) ClaimYield(opts *bind.TransactOpts, underlying common.Address, vTokenAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "claimYield", underlying, vTokenAmount, recipient)
}

// ClaimYield is a paid mutator transaction binding the contract method 0xf1b0e8eb.
//
// Solidity: function claimYield(address underlying, uint256 vTokenAmount, address recipient) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) ClaimYield(underlying common.Address, vTokenAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.ClaimYield(&_YBConditionalTokens.TransactOpts, underlying, vTokenAmount, recipient)
}

// ClaimYield is a paid mutator transaction binding the contract method 0xf1b0e8eb.
//
// Solidity: function claimYield(address underlying, uint256 vTokenAmount, address recipient) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) ClaimYield(underlying common.Address, vTokenAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.ClaimYield(&_YBConditionalTokens.TransactOpts, underlying, vTokenAmount, recipient)
}

// ConnectVTokenToUnderlying is a paid mutator transaction binding the contract method 0x7ec239a7.
//
// Solidity: function connectVTokenToUnderlying(address vToken) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) ConnectVTokenToUnderlying(opts *bind.TransactOpts, vToken common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "connectVTokenToUnderlying", vToken)
}

// ConnectVTokenToUnderlying is a paid mutator transaction binding the contract method 0x7ec239a7.
//
// Solidity: function connectVTokenToUnderlying(address vToken) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) ConnectVTokenToUnderlying(vToken common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.ConnectVTokenToUnderlying(&_YBConditionalTokens.TransactOpts, vToken)
}

// ConnectVTokenToUnderlying is a paid mutator transaction binding the contract method 0x7ec239a7.
//
// Solidity: function connectVTokenToUnderlying(address vToken) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) ConnectVTokenToUnderlying(vToken common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.ConnectVTokenToUnderlying(&_YBConditionalTokens.TransactOpts, vToken)
}

// DisableUnderlying is a paid mutator transaction binding the contract method 0xcf314cb1.
//
// Solidity: function disableUnderlying(address underlying, address yieldRecipient) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) DisableUnderlying(opts *bind.TransactOpts, underlying common.Address, yieldRecipient common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "disableUnderlying", underlying, yieldRecipient)
}

// DisableUnderlying is a paid mutator transaction binding the contract method 0xcf314cb1.
//
// Solidity: function disableUnderlying(address underlying, address yieldRecipient) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) DisableUnderlying(underlying common.Address, yieldRecipient common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.DisableUnderlying(&_YBConditionalTokens.TransactOpts, underlying, yieldRecipient)
}

// DisableUnderlying is a paid mutator transaction binding the contract method 0xcf314cb1.
//
// Solidity: function disableUnderlying(address underlying, address yieldRecipient) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) DisableUnderlying(underlying common.Address, yieldRecipient common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.DisableUnderlying(&_YBConditionalTokens.TransactOpts, underlying, yieldRecipient)
}

// EnableUnderlying is a paid mutator transaction binding the contract method 0xe74784e4.
//
// Solidity: function enableUnderlying(address underlying) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) EnableUnderlying(opts *bind.TransactOpts, underlying common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "enableUnderlying", underlying)
}

// EnableUnderlying is a paid mutator transaction binding the contract method 0xe74784e4.
//
// Solidity: function enableUnderlying(address underlying) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) EnableUnderlying(underlying common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.EnableUnderlying(&_YBConditionalTokens.TransactOpts, underlying)
}

// EnableUnderlying is a paid mutator transaction binding the contract method 0xe74784e4.
//
// Solidity: function enableUnderlying(address underlying) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) EnableUnderlying(underlying common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.EnableUnderlying(&_YBConditionalTokens.TransactOpts, underlying)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.GrantRole(&_YBConditionalTokens.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.GrantRole(&_YBConditionalTokens.TransactOpts, role, account)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) MergePositions(opts *bind.TransactOpts, collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "mergePositions", collateralToken, parentCollectionId, conditionId, partition, amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) MergePositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.MergePositions(&_YBConditionalTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) MergePositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.MergePositions(&_YBConditionalTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// PrepareCondition is a paid mutator transaction binding the contract method 0x495f17f9.
//
// Solidity: function prepareCondition(bytes32 questionId, uint256 outcomeSlotCount) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) PrepareCondition(opts *bind.TransactOpts, questionId [32]byte, outcomeSlotCount *big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "prepareCondition", questionId, outcomeSlotCount)
}

// PrepareCondition is a paid mutator transaction binding the contract method 0x495f17f9.
//
// Solidity: function prepareCondition(bytes32 questionId, uint256 outcomeSlotCount) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) PrepareCondition(questionId [32]byte, outcomeSlotCount *big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.PrepareCondition(&_YBConditionalTokens.TransactOpts, questionId, outcomeSlotCount)
}

// PrepareCondition is a paid mutator transaction binding the contract method 0x495f17f9.
//
// Solidity: function prepareCondition(bytes32 questionId, uint256 outcomeSlotCount) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) PrepareCondition(questionId [32]byte, outcomeSlotCount *big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.PrepareCondition(&_YBConditionalTokens.TransactOpts, questionId, outcomeSlotCount)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] indexSets) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) RedeemPositions(opts *bind.TransactOpts, collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, indexSets []*big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "redeemPositions", collateralToken, parentCollectionId, conditionId, indexSets)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] indexSets) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) RedeemPositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, indexSets []*big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.RedeemPositions(&_YBConditionalTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, indexSets)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] indexSets) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) RedeemPositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, indexSets []*big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.RedeemPositions(&_YBConditionalTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, indexSets)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.RenounceRole(&_YBConditionalTokens.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.RenounceRole(&_YBConditionalTokens.TransactOpts, role, callerConfirmation)
}

// ReportPayouts is a paid mutator transaction binding the contract method 0xc49298ac.
//
// Solidity: function reportPayouts(bytes32 questionId, uint256[] payouts) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) ReportPayouts(opts *bind.TransactOpts, questionId [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "reportPayouts", questionId, payouts)
}

// ReportPayouts is a paid mutator transaction binding the contract method 0xc49298ac.
//
// Solidity: function reportPayouts(bytes32 questionId, uint256[] payouts) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) ReportPayouts(questionId [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.ReportPayouts(&_YBConditionalTokens.TransactOpts, questionId, payouts)
}

// ReportPayouts is a paid mutator transaction binding the contract method 0xc49298ac.
//
// Solidity: function reportPayouts(bytes32 questionId, uint256[] payouts) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) ReportPayouts(questionId [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.ReportPayouts(&_YBConditionalTokens.TransactOpts, questionId, payouts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.RevokeRole(&_YBConditionalTokens.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.RevokeRole(&_YBConditionalTokens.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SafeBatchTransferFrom(&_YBConditionalTokens.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SafeBatchTransferFrom(&_YBConditionalTokens.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SafeTransferFrom(&_YBConditionalTokens.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SafeTransferFrom(&_YBConditionalTokens.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SetApprovalForAll(&_YBConditionalTokens.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SetApprovalForAll(&_YBConditionalTokens.TransactOpts, operator, approved)
}

// SetIsMergePositionsGated is a paid mutator transaction binding the contract method 0x80d74ff9.
//
// Solidity: function setIsMergePositionsGated(bool gated) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) SetIsMergePositionsGated(opts *bind.TransactOpts, gated bool) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "setIsMergePositionsGated", gated)
}

// SetIsMergePositionsGated is a paid mutator transaction binding the contract method 0x80d74ff9.
//
// Solidity: function setIsMergePositionsGated(bool gated) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) SetIsMergePositionsGated(gated bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SetIsMergePositionsGated(&_YBConditionalTokens.TransactOpts, gated)
}

// SetIsMergePositionsGated is a paid mutator transaction binding the contract method 0x80d74ff9.
//
// Solidity: function setIsMergePositionsGated(bool gated) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) SetIsMergePositionsGated(gated bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SetIsMergePositionsGated(&_YBConditionalTokens.TransactOpts, gated)
}

// SetIsSplitPositionGated is a paid mutator transaction binding the contract method 0x5547ab74.
//
// Solidity: function setIsSplitPositionGated(bool gated) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) SetIsSplitPositionGated(opts *bind.TransactOpts, gated bool) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "setIsSplitPositionGated", gated)
}

// SetIsSplitPositionGated is a paid mutator transaction binding the contract method 0x5547ab74.
//
// Solidity: function setIsSplitPositionGated(bool gated) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) SetIsSplitPositionGated(gated bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SetIsSplitPositionGated(&_YBConditionalTokens.TransactOpts, gated)
}

// SetIsSplitPositionGated is a paid mutator transaction binding the contract method 0x5547ab74.
//
// Solidity: function setIsSplitPositionGated(bool gated) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) SetIsSplitPositionGated(gated bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SetIsSplitPositionGated(&_YBConditionalTokens.TransactOpts, gated)
}

// SetTransferControlEnabled is a paid mutator transaction binding the contract method 0xd4e4c30e.
//
// Solidity: function setTransferControlEnabled(bool enabled) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) SetTransferControlEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "setTransferControlEnabled", enabled)
}

// SetTransferControlEnabled is a paid mutator transaction binding the contract method 0xd4e4c30e.
//
// Solidity: function setTransferControlEnabled(bool enabled) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) SetTransferControlEnabled(enabled bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SetTransferControlEnabled(&_YBConditionalTokens.TransactOpts, enabled)
}

// SetTransferControlEnabled is a paid mutator transaction binding the contract method 0xd4e4c30e.
//
// Solidity: function setTransferControlEnabled(bool enabled) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) SetTransferControlEnabled(enabled bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SetTransferControlEnabled(&_YBConditionalTokens.TransactOpts, enabled)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) SplitPosition(opts *bind.TransactOpts, collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "splitPosition", collateralToken, parentCollectionId, conditionId, partition, amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) SplitPosition(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SplitPosition(&_YBConditionalTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) SplitPosition(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SplitPosition(&_YBConditionalTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// SplitPrincipalAndYield is a paid mutator transaction binding the contract method 0x0d7468bc.
//
// Solidity: function splitPrincipalAndYield(address underlying) returns(uint256 principal, uint256 yield)
func (_YBConditionalTokens *YBConditionalTokensTransactor) SplitPrincipalAndYield(opts *bind.TransactOpts, underlying common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "splitPrincipalAndYield", underlying)
}

// SplitPrincipalAndYield is a paid mutator transaction binding the contract method 0x0d7468bc.
//
// Solidity: function splitPrincipalAndYield(address underlying) returns(uint256 principal, uint256 yield)
func (_YBConditionalTokens *YBConditionalTokensSession) SplitPrincipalAndYield(underlying common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SplitPrincipalAndYield(&_YBConditionalTokens.TransactOpts, underlying)
}

// SplitPrincipalAndYield is a paid mutator transaction binding the contract method 0x0d7468bc.
//
// Solidity: function splitPrincipalAndYield(address underlying) returns(uint256 principal, uint256 yield)
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) SplitPrincipalAndYield(underlying common.Address) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.SplitPrincipalAndYield(&_YBConditionalTokens.TransactOpts, underlying)
}

// UpdateIsTransferAllowed is a paid mutator transaction binding the contract method 0x5dd47e2e.
//
// Solidity: function updateIsTransferAllowed(address addr, bool isAllowed) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactor) UpdateIsTransferAllowed(opts *bind.TransactOpts, addr common.Address, isAllowed bool) (*types.Transaction, error) {
	return _YBConditionalTokens.contract.Transact(opts, "updateIsTransferAllowed", addr, isAllowed)
}

// UpdateIsTransferAllowed is a paid mutator transaction binding the contract method 0x5dd47e2e.
//
// Solidity: function updateIsTransferAllowed(address addr, bool isAllowed) returns()
func (_YBConditionalTokens *YBConditionalTokensSession) UpdateIsTransferAllowed(addr common.Address, isAllowed bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.UpdateIsTransferAllowed(&_YBConditionalTokens.TransactOpts, addr, isAllowed)
}

// UpdateIsTransferAllowed is a paid mutator transaction binding the contract method 0x5dd47e2e.
//
// Solidity: function updateIsTransferAllowed(address addr, bool isAllowed) returns()
func (_YBConditionalTokens *YBConditionalTokensTransactorSession) UpdateIsTransferAllowed(addr common.Address, isAllowed bool) (*types.Transaction, error) {
	return _YBConditionalTokens.Contract.UpdateIsTransferAllowed(&_YBConditionalTokens.TransactOpts, addr, isAllowed)
}

// YBConditionalTokensApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the YBConditionalTokens contract.
type YBConditionalTokensApprovalForAllIterator struct {
	Event *YBConditionalTokensApprovalForAll // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensApprovalForAll)
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
		it.Event = new(YBConditionalTokensApprovalForAll)
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
func (it *YBConditionalTokensApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensApprovalForAll represents a ApprovalForAll event raised by the YBConditionalTokens contract.
type YBConditionalTokensApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*YBConditionalTokensApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensApprovalForAllIterator{contract: _YBConditionalTokens.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensApprovalForAll)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseApprovalForAll(log types.Log) (*YBConditionalTokensApprovalForAll, error) {
	event := new(YBConditionalTokensApprovalForAll)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensConditionPreparationIterator is returned from FilterConditionPreparation and is used to iterate over the raw logs and unpacked data for ConditionPreparation events raised by the YBConditionalTokens contract.
type YBConditionalTokensConditionPreparationIterator struct {
	Event *YBConditionalTokensConditionPreparation // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensConditionPreparationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensConditionPreparation)
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
		it.Event = new(YBConditionalTokensConditionPreparation)
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
func (it *YBConditionalTokensConditionPreparationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensConditionPreparationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensConditionPreparation represents a ConditionPreparation event raised by the YBConditionalTokens contract.
type YBConditionalTokensConditionPreparation struct {
	ConditionId      [32]byte
	Oracle           common.Address
	QuestionId       [32]byte
	OutcomeSlotCount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConditionPreparation is a free log retrieval operation binding the contract event 0xab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177.
//
// Solidity: event ConditionPreparation(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterConditionPreparation(opts *bind.FilterOpts, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (*YBConditionalTokensConditionPreparationIterator, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "ConditionPreparation", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensConditionPreparationIterator{contract: _YBConditionalTokens.contract, event: "ConditionPreparation", logs: logs, sub: sub}, nil
}

// WatchConditionPreparation is a free log subscription operation binding the contract event 0xab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177.
//
// Solidity: event ConditionPreparation(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchConditionPreparation(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensConditionPreparation, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (event.Subscription, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "ConditionPreparation", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensConditionPreparation)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "ConditionPreparation", log); err != nil {
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

// ParseConditionPreparation is a log parse operation binding the contract event 0xab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177.
//
// Solidity: event ConditionPreparation(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseConditionPreparation(log types.Log) (*YBConditionalTokensConditionPreparation, error) {
	event := new(YBConditionalTokensConditionPreparation)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "ConditionPreparation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensConditionResolutionIterator is returned from FilterConditionResolution and is used to iterate over the raw logs and unpacked data for ConditionResolution events raised by the YBConditionalTokens contract.
type YBConditionalTokensConditionResolutionIterator struct {
	Event *YBConditionalTokensConditionResolution // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensConditionResolutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensConditionResolution)
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
		it.Event = new(YBConditionalTokensConditionResolution)
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
func (it *YBConditionalTokensConditionResolutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensConditionResolutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensConditionResolution represents a ConditionResolution event raised by the YBConditionalTokens contract.
type YBConditionalTokensConditionResolution struct {
	ConditionId      [32]byte
	Oracle           common.Address
	QuestionId       [32]byte
	OutcomeSlotCount *big.Int
	PayoutNumerators []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConditionResolution is a free log retrieval operation binding the contract event 0xb44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894.
//
// Solidity: event ConditionResolution(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount, uint256[] payoutNumerators)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterConditionResolution(opts *bind.FilterOpts, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (*YBConditionalTokensConditionResolutionIterator, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "ConditionResolution", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensConditionResolutionIterator{contract: _YBConditionalTokens.contract, event: "ConditionResolution", logs: logs, sub: sub}, nil
}

// WatchConditionResolution is a free log subscription operation binding the contract event 0xb44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894.
//
// Solidity: event ConditionResolution(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount, uint256[] payoutNumerators)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchConditionResolution(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensConditionResolution, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (event.Subscription, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "ConditionResolution", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensConditionResolution)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "ConditionResolution", log); err != nil {
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

// ParseConditionResolution is a log parse operation binding the contract event 0xb44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894.
//
// Solidity: event ConditionResolution(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount, uint256[] payoutNumerators)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseConditionResolution(log types.Log) (*YBConditionalTokensConditionResolution, error) {
	event := new(YBConditionalTokensConditionResolution)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "ConditionResolution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensIsMergePositionsGatedIterator is returned from FilterIsMergePositionsGated and is used to iterate over the raw logs and unpacked data for IsMergePositionsGated events raised by the YBConditionalTokens contract.
type YBConditionalTokensIsMergePositionsGatedIterator struct {
	Event *YBConditionalTokensIsMergePositionsGated // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensIsMergePositionsGatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensIsMergePositionsGated)
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
		it.Event = new(YBConditionalTokensIsMergePositionsGated)
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
func (it *YBConditionalTokensIsMergePositionsGatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensIsMergePositionsGatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensIsMergePositionsGated represents a IsMergePositionsGated event raised by the YBConditionalTokens contract.
type YBConditionalTokensIsMergePositionsGated struct {
	Gated bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIsMergePositionsGated is a free log retrieval operation binding the contract event 0x32acd5cb959296624fe98cabbebabe6f94bf2fccb2cf25d5986545eaeb4fda1f.
//
// Solidity: event IsMergePositionsGated(bool gated)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterIsMergePositionsGated(opts *bind.FilterOpts) (*YBConditionalTokensIsMergePositionsGatedIterator, error) {

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "IsMergePositionsGated")
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensIsMergePositionsGatedIterator{contract: _YBConditionalTokens.contract, event: "IsMergePositionsGated", logs: logs, sub: sub}, nil
}

// WatchIsMergePositionsGated is a free log subscription operation binding the contract event 0x32acd5cb959296624fe98cabbebabe6f94bf2fccb2cf25d5986545eaeb4fda1f.
//
// Solidity: event IsMergePositionsGated(bool gated)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchIsMergePositionsGated(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensIsMergePositionsGated) (event.Subscription, error) {

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "IsMergePositionsGated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensIsMergePositionsGated)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "IsMergePositionsGated", log); err != nil {
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

// ParseIsMergePositionsGated is a log parse operation binding the contract event 0x32acd5cb959296624fe98cabbebabe6f94bf2fccb2cf25d5986545eaeb4fda1f.
//
// Solidity: event IsMergePositionsGated(bool gated)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseIsMergePositionsGated(log types.Log) (*YBConditionalTokensIsMergePositionsGated, error) {
	event := new(YBConditionalTokensIsMergePositionsGated)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "IsMergePositionsGated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensIsSplitPositionGatedIterator is returned from FilterIsSplitPositionGated and is used to iterate over the raw logs and unpacked data for IsSplitPositionGated events raised by the YBConditionalTokens contract.
type YBConditionalTokensIsSplitPositionGatedIterator struct {
	Event *YBConditionalTokensIsSplitPositionGated // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensIsSplitPositionGatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensIsSplitPositionGated)
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
		it.Event = new(YBConditionalTokensIsSplitPositionGated)
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
func (it *YBConditionalTokensIsSplitPositionGatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensIsSplitPositionGatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensIsSplitPositionGated represents a IsSplitPositionGated event raised by the YBConditionalTokens contract.
type YBConditionalTokensIsSplitPositionGated struct {
	Gated bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIsSplitPositionGated is a free log retrieval operation binding the contract event 0x037fb73c20dedbdf8890e142b309dce4c6327f39814173f5278183d0cb78d741.
//
// Solidity: event IsSplitPositionGated(bool gated)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterIsSplitPositionGated(opts *bind.FilterOpts) (*YBConditionalTokensIsSplitPositionGatedIterator, error) {

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "IsSplitPositionGated")
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensIsSplitPositionGatedIterator{contract: _YBConditionalTokens.contract, event: "IsSplitPositionGated", logs: logs, sub: sub}, nil
}

// WatchIsSplitPositionGated is a free log subscription operation binding the contract event 0x037fb73c20dedbdf8890e142b309dce4c6327f39814173f5278183d0cb78d741.
//
// Solidity: event IsSplitPositionGated(bool gated)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchIsSplitPositionGated(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensIsSplitPositionGated) (event.Subscription, error) {

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "IsSplitPositionGated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensIsSplitPositionGated)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "IsSplitPositionGated", log); err != nil {
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

// ParseIsSplitPositionGated is a log parse operation binding the contract event 0x037fb73c20dedbdf8890e142b309dce4c6327f39814173f5278183d0cb78d741.
//
// Solidity: event IsSplitPositionGated(bool gated)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseIsSplitPositionGated(log types.Log) (*YBConditionalTokensIsSplitPositionGated, error) {
	event := new(YBConditionalTokensIsSplitPositionGated)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "IsSplitPositionGated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensIsTransferAllowedUpdatedIterator is returned from FilterIsTransferAllowedUpdated and is used to iterate over the raw logs and unpacked data for IsTransferAllowedUpdated events raised by the YBConditionalTokens contract.
type YBConditionalTokensIsTransferAllowedUpdatedIterator struct {
	Event *YBConditionalTokensIsTransferAllowedUpdated // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensIsTransferAllowedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensIsTransferAllowedUpdated)
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
		it.Event = new(YBConditionalTokensIsTransferAllowedUpdated)
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
func (it *YBConditionalTokensIsTransferAllowedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensIsTransferAllowedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensIsTransferAllowedUpdated represents a IsTransferAllowedUpdated event raised by the YBConditionalTokens contract.
type YBConditionalTokensIsTransferAllowedUpdated struct {
	Addr      common.Address
	IsAllowed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIsTransferAllowedUpdated is a free log retrieval operation binding the contract event 0x89ed7823a4b81a8c6f8f4bcbefe7d44715aada5fd92c54c06ee0ae8cee6ae1cd.
//
// Solidity: event IsTransferAllowedUpdated(address indexed addr, bool isAllowed)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterIsTransferAllowedUpdated(opts *bind.FilterOpts, addr []common.Address) (*YBConditionalTokensIsTransferAllowedUpdatedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "IsTransferAllowedUpdated", addrRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensIsTransferAllowedUpdatedIterator{contract: _YBConditionalTokens.contract, event: "IsTransferAllowedUpdated", logs: logs, sub: sub}, nil
}

// WatchIsTransferAllowedUpdated is a free log subscription operation binding the contract event 0x89ed7823a4b81a8c6f8f4bcbefe7d44715aada5fd92c54c06ee0ae8cee6ae1cd.
//
// Solidity: event IsTransferAllowedUpdated(address indexed addr, bool isAllowed)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchIsTransferAllowedUpdated(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensIsTransferAllowedUpdated, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "IsTransferAllowedUpdated", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensIsTransferAllowedUpdated)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "IsTransferAllowedUpdated", log); err != nil {
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

// ParseIsTransferAllowedUpdated is a log parse operation binding the contract event 0x89ed7823a4b81a8c6f8f4bcbefe7d44715aada5fd92c54c06ee0ae8cee6ae1cd.
//
// Solidity: event IsTransferAllowedUpdated(address indexed addr, bool isAllowed)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseIsTransferAllowedUpdated(log types.Log) (*YBConditionalTokensIsTransferAllowedUpdated, error) {
	event := new(YBConditionalTokensIsTransferAllowedUpdated)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "IsTransferAllowedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensPayoutRedemptionIterator is returned from FilterPayoutRedemption and is used to iterate over the raw logs and unpacked data for PayoutRedemption events raised by the YBConditionalTokens contract.
type YBConditionalTokensPayoutRedemptionIterator struct {
	Event *YBConditionalTokensPayoutRedemption // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensPayoutRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensPayoutRedemption)
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
		it.Event = new(YBConditionalTokensPayoutRedemption)
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
func (it *YBConditionalTokensPayoutRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensPayoutRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensPayoutRedemption represents a PayoutRedemption event raised by the YBConditionalTokens contract.
type YBConditionalTokensPayoutRedemption struct {
	Redeemer           common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	IndexSets          []*big.Int
	Payout             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPayoutRedemption is a free log retrieval operation binding the contract event 0x2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d.
//
// Solidity: event PayoutRedemption(address indexed redeemer, address indexed collateralToken, bytes32 indexed parentCollectionId, bytes32 conditionId, uint256[] indexSets, uint256 payout)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterPayoutRedemption(opts *bind.FilterOpts, redeemer []common.Address, collateralToken []common.Address, parentCollectionId [][32]byte) (*YBConditionalTokensPayoutRedemptionIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var collateralTokenRule []interface{}
	for _, collateralTokenItem := range collateralToken {
		collateralTokenRule = append(collateralTokenRule, collateralTokenItem)
	}
	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "PayoutRedemption", redeemerRule, collateralTokenRule, parentCollectionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensPayoutRedemptionIterator{contract: _YBConditionalTokens.contract, event: "PayoutRedemption", logs: logs, sub: sub}, nil
}

// WatchPayoutRedemption is a free log subscription operation binding the contract event 0x2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d.
//
// Solidity: event PayoutRedemption(address indexed redeemer, address indexed collateralToken, bytes32 indexed parentCollectionId, bytes32 conditionId, uint256[] indexSets, uint256 payout)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchPayoutRedemption(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensPayoutRedemption, redeemer []common.Address, collateralToken []common.Address, parentCollectionId [][32]byte) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var collateralTokenRule []interface{}
	for _, collateralTokenItem := range collateralToken {
		collateralTokenRule = append(collateralTokenRule, collateralTokenItem)
	}
	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "PayoutRedemption", redeemerRule, collateralTokenRule, parentCollectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensPayoutRedemption)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
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

// ParsePayoutRedemption is a log parse operation binding the contract event 0x2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d.
//
// Solidity: event PayoutRedemption(address indexed redeemer, address indexed collateralToken, bytes32 indexed parentCollectionId, bytes32 conditionId, uint256[] indexSets, uint256 payout)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParsePayoutRedemption(log types.Log) (*YBConditionalTokensPayoutRedemption, error) {
	event := new(YBConditionalTokensPayoutRedemption)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensPositionSplitIterator is returned from FilterPositionSplit and is used to iterate over the raw logs and unpacked data for PositionSplit events raised by the YBConditionalTokens contract.
type YBConditionalTokensPositionSplitIterator struct {
	Event *YBConditionalTokensPositionSplit // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensPositionSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensPositionSplit)
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
		it.Event = new(YBConditionalTokensPositionSplit)
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
func (it *YBConditionalTokensPositionSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensPositionSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensPositionSplit represents a PositionSplit event raised by the YBConditionalTokens contract.
type YBConditionalTokensPositionSplit struct {
	Stakeholder        common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	Partition          []*big.Int
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPositionSplit is a free log retrieval operation binding the contract event 0x2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298.
//
// Solidity: event PositionSplit(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterPositionSplit(opts *bind.FilterOpts, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (*YBConditionalTokensPositionSplitIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "PositionSplit", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensPositionSplitIterator{contract: _YBConditionalTokens.contract, event: "PositionSplit", logs: logs, sub: sub}, nil
}

// WatchPositionSplit is a free log subscription operation binding the contract event 0x2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298.
//
// Solidity: event PositionSplit(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchPositionSplit(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensPositionSplit, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "PositionSplit", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensPositionSplit)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "PositionSplit", log); err != nil {
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

// ParsePositionSplit is a log parse operation binding the contract event 0x2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298.
//
// Solidity: event PositionSplit(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParsePositionSplit(log types.Log) (*YBConditionalTokensPositionSplit, error) {
	event := new(YBConditionalTokensPositionSplit)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "PositionSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensPositionsMergeIterator is returned from FilterPositionsMerge and is used to iterate over the raw logs and unpacked data for PositionsMerge events raised by the YBConditionalTokens contract.
type YBConditionalTokensPositionsMergeIterator struct {
	Event *YBConditionalTokensPositionsMerge // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensPositionsMergeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensPositionsMerge)
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
		it.Event = new(YBConditionalTokensPositionsMerge)
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
func (it *YBConditionalTokensPositionsMergeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensPositionsMergeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensPositionsMerge represents a PositionsMerge event raised by the YBConditionalTokens contract.
type YBConditionalTokensPositionsMerge struct {
	Stakeholder        common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	Partition          []*big.Int
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPositionsMerge is a free log retrieval operation binding the contract event 0x6f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca.
//
// Solidity: event PositionsMerge(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterPositionsMerge(opts *bind.FilterOpts, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (*YBConditionalTokensPositionsMergeIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "PositionsMerge", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensPositionsMergeIterator{contract: _YBConditionalTokens.contract, event: "PositionsMerge", logs: logs, sub: sub}, nil
}

// WatchPositionsMerge is a free log subscription operation binding the contract event 0x6f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca.
//
// Solidity: event PositionsMerge(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchPositionsMerge(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensPositionsMerge, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "PositionsMerge", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensPositionsMerge)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
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

// ParsePositionsMerge is a log parse operation binding the contract event 0x6f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca.
//
// Solidity: event PositionsMerge(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParsePositionsMerge(log types.Log) (*YBConditionalTokensPositionsMerge, error) {
	event := new(YBConditionalTokensPositionsMerge)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the YBConditionalTokens contract.
type YBConditionalTokensRoleAdminChangedIterator struct {
	Event *YBConditionalTokensRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensRoleAdminChanged)
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
		it.Event = new(YBConditionalTokensRoleAdminChanged)
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
func (it *YBConditionalTokensRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensRoleAdminChanged represents a RoleAdminChanged event raised by the YBConditionalTokens contract.
type YBConditionalTokensRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*YBConditionalTokensRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensRoleAdminChangedIterator{contract: _YBConditionalTokens.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensRoleAdminChanged)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseRoleAdminChanged(log types.Log) (*YBConditionalTokensRoleAdminChanged, error) {
	event := new(YBConditionalTokensRoleAdminChanged)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the YBConditionalTokens contract.
type YBConditionalTokensRoleGrantedIterator struct {
	Event *YBConditionalTokensRoleGranted // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensRoleGranted)
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
		it.Event = new(YBConditionalTokensRoleGranted)
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
func (it *YBConditionalTokensRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensRoleGranted represents a RoleGranted event raised by the YBConditionalTokens contract.
type YBConditionalTokensRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*YBConditionalTokensRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensRoleGrantedIterator{contract: _YBConditionalTokens.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensRoleGranted)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseRoleGranted(log types.Log) (*YBConditionalTokensRoleGranted, error) {
	event := new(YBConditionalTokensRoleGranted)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the YBConditionalTokens contract.
type YBConditionalTokensRoleRevokedIterator struct {
	Event *YBConditionalTokensRoleRevoked // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensRoleRevoked)
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
		it.Event = new(YBConditionalTokensRoleRevoked)
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
func (it *YBConditionalTokensRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensRoleRevoked represents a RoleRevoked event raised by the YBConditionalTokens contract.
type YBConditionalTokensRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*YBConditionalTokensRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensRoleRevokedIterator{contract: _YBConditionalTokens.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensRoleRevoked)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseRoleRevoked(log types.Log) (*YBConditionalTokensRoleRevoked, error) {
	event := new(YBConditionalTokensRoleRevoked)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the YBConditionalTokens contract.
type YBConditionalTokensTransferBatchIterator struct {
	Event *YBConditionalTokensTransferBatch // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensTransferBatch)
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
		it.Event = new(YBConditionalTokensTransferBatch)
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
func (it *YBConditionalTokensTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensTransferBatch represents a TransferBatch event raised by the YBConditionalTokens contract.
type YBConditionalTokensTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*YBConditionalTokensTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensTransferBatchIterator{contract: _YBConditionalTokens.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensTransferBatch)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseTransferBatch(log types.Log) (*YBConditionalTokensTransferBatch, error) {
	event := new(YBConditionalTokensTransferBatch)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensTransferControlEnabledIterator is returned from FilterTransferControlEnabled and is used to iterate over the raw logs and unpacked data for TransferControlEnabled events raised by the YBConditionalTokens contract.
type YBConditionalTokensTransferControlEnabledIterator struct {
	Event *YBConditionalTokensTransferControlEnabled // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensTransferControlEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensTransferControlEnabled)
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
		it.Event = new(YBConditionalTokensTransferControlEnabled)
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
func (it *YBConditionalTokensTransferControlEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensTransferControlEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensTransferControlEnabled represents a TransferControlEnabled event raised by the YBConditionalTokens contract.
type YBConditionalTokensTransferControlEnabled struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferControlEnabled is a free log retrieval operation binding the contract event 0x7c9b1f40d342591d9ca85220c58275fc1b8825dc2216016cdff2b9cff1eb0c09.
//
// Solidity: event TransferControlEnabled(bool enabled)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterTransferControlEnabled(opts *bind.FilterOpts) (*YBConditionalTokensTransferControlEnabledIterator, error) {

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "TransferControlEnabled")
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensTransferControlEnabledIterator{contract: _YBConditionalTokens.contract, event: "TransferControlEnabled", logs: logs, sub: sub}, nil
}

// WatchTransferControlEnabled is a free log subscription operation binding the contract event 0x7c9b1f40d342591d9ca85220c58275fc1b8825dc2216016cdff2b9cff1eb0c09.
//
// Solidity: event TransferControlEnabled(bool enabled)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchTransferControlEnabled(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensTransferControlEnabled) (event.Subscription, error) {

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "TransferControlEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensTransferControlEnabled)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "TransferControlEnabled", log); err != nil {
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

// ParseTransferControlEnabled is a log parse operation binding the contract event 0x7c9b1f40d342591d9ca85220c58275fc1b8825dc2216016cdff2b9cff1eb0c09.
//
// Solidity: event TransferControlEnabled(bool enabled)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseTransferControlEnabled(log types.Log) (*YBConditionalTokensTransferControlEnabled, error) {
	event := new(YBConditionalTokensTransferControlEnabled)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "TransferControlEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the YBConditionalTokens contract.
type YBConditionalTokensTransferSingleIterator struct {
	Event *YBConditionalTokensTransferSingle // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensTransferSingle)
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
		it.Event = new(YBConditionalTokensTransferSingle)
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
func (it *YBConditionalTokensTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensTransferSingle represents a TransferSingle event raised by the YBConditionalTokens contract.
type YBConditionalTokensTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*YBConditionalTokensTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensTransferSingleIterator{contract: _YBConditionalTokens.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensTransferSingle)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseTransferSingle(log types.Log) (*YBConditionalTokensTransferSingle, error) {
	event := new(YBConditionalTokensTransferSingle)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the YBConditionalTokens contract.
type YBConditionalTokensURIIterator struct {
	Event *YBConditionalTokensURI // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensURI)
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
		it.Event = new(YBConditionalTokensURI)
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
func (it *YBConditionalTokensURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensURI represents a URI event raised by the YBConditionalTokens contract.
type YBConditionalTokensURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*YBConditionalTokensURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensURIIterator{contract: _YBConditionalTokens.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensURI)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseURI(log types.Log) (*YBConditionalTokensURI, error) {
	event := new(YBConditionalTokensURI)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensUnderlyingTokenConnectedToVTokenIterator is returned from FilterUnderlyingTokenConnectedToVToken and is used to iterate over the raw logs and unpacked data for UnderlyingTokenConnectedToVToken events raised by the YBConditionalTokens contract.
type YBConditionalTokensUnderlyingTokenConnectedToVTokenIterator struct {
	Event *YBConditionalTokensUnderlyingTokenConnectedToVToken // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensUnderlyingTokenConnectedToVTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensUnderlyingTokenConnectedToVToken)
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
		it.Event = new(YBConditionalTokensUnderlyingTokenConnectedToVToken)
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
func (it *YBConditionalTokensUnderlyingTokenConnectedToVTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensUnderlyingTokenConnectedToVTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensUnderlyingTokenConnectedToVToken represents a UnderlyingTokenConnectedToVToken event raised by the YBConditionalTokens contract.
type YBConditionalTokensUnderlyingTokenConnectedToVToken struct {
	Underlying common.Address
	VToken     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnderlyingTokenConnectedToVToken is a free log retrieval operation binding the contract event 0xb05af722d3d9d87d1357b327a71ea82ddac5f44b037563dd739c1d0302d85724.
//
// Solidity: event UnderlyingTokenConnectedToVToken(address underlying, address vToken)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterUnderlyingTokenConnectedToVToken(opts *bind.FilterOpts) (*YBConditionalTokensUnderlyingTokenConnectedToVTokenIterator, error) {

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "UnderlyingTokenConnectedToVToken")
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensUnderlyingTokenConnectedToVTokenIterator{contract: _YBConditionalTokens.contract, event: "UnderlyingTokenConnectedToVToken", logs: logs, sub: sub}, nil
}

// WatchUnderlyingTokenConnectedToVToken is a free log subscription operation binding the contract event 0xb05af722d3d9d87d1357b327a71ea82ddac5f44b037563dd739c1d0302d85724.
//
// Solidity: event UnderlyingTokenConnectedToVToken(address underlying, address vToken)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchUnderlyingTokenConnectedToVToken(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensUnderlyingTokenConnectedToVToken) (event.Subscription, error) {

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "UnderlyingTokenConnectedToVToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensUnderlyingTokenConnectedToVToken)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "UnderlyingTokenConnectedToVToken", log); err != nil {
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

// ParseUnderlyingTokenConnectedToVToken is a log parse operation binding the contract event 0xb05af722d3d9d87d1357b327a71ea82ddac5f44b037563dd739c1d0302d85724.
//
// Solidity: event UnderlyingTokenConnectedToVToken(address underlying, address vToken)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseUnderlyingTokenConnectedToVToken(log types.Log) (*YBConditionalTokensUnderlyingTokenConnectedToVToken, error) {
	event := new(YBConditionalTokensUnderlyingTokenConnectedToVToken)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "UnderlyingTokenConnectedToVToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensUnderlyingTokenDisabledIterator is returned from FilterUnderlyingTokenDisabled and is used to iterate over the raw logs and unpacked data for UnderlyingTokenDisabled events raised by the YBConditionalTokens contract.
type YBConditionalTokensUnderlyingTokenDisabledIterator struct {
	Event *YBConditionalTokensUnderlyingTokenDisabled // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensUnderlyingTokenDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensUnderlyingTokenDisabled)
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
		it.Event = new(YBConditionalTokensUnderlyingTokenDisabled)
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
func (it *YBConditionalTokensUnderlyingTokenDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensUnderlyingTokenDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensUnderlyingTokenDisabled represents a UnderlyingTokenDisabled event raised by the YBConditionalTokens contract.
type YBConditionalTokensUnderlyingTokenDisabled struct {
	Underlying     common.Address
	AmountRedeemed *big.Int
	YieldRecipient common.Address
	YieldClaimed   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUnderlyingTokenDisabled is a free log retrieval operation binding the contract event 0x967c37c135825f4c088eaa8f9237e3977ba0c5c8dddf5d9170e833c0109e247f.
//
// Solidity: event UnderlyingTokenDisabled(address indexed underlying, uint256 amountRedeemed, address yieldRecipient, uint256 yieldClaimed)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterUnderlyingTokenDisabled(opts *bind.FilterOpts, underlying []common.Address) (*YBConditionalTokensUnderlyingTokenDisabledIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "UnderlyingTokenDisabled", underlyingRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensUnderlyingTokenDisabledIterator{contract: _YBConditionalTokens.contract, event: "UnderlyingTokenDisabled", logs: logs, sub: sub}, nil
}

// WatchUnderlyingTokenDisabled is a free log subscription operation binding the contract event 0x967c37c135825f4c088eaa8f9237e3977ba0c5c8dddf5d9170e833c0109e247f.
//
// Solidity: event UnderlyingTokenDisabled(address indexed underlying, uint256 amountRedeemed, address yieldRecipient, uint256 yieldClaimed)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchUnderlyingTokenDisabled(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensUnderlyingTokenDisabled, underlying []common.Address) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "UnderlyingTokenDisabled", underlyingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensUnderlyingTokenDisabled)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "UnderlyingTokenDisabled", log); err != nil {
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

// ParseUnderlyingTokenDisabled is a log parse operation binding the contract event 0x967c37c135825f4c088eaa8f9237e3977ba0c5c8dddf5d9170e833c0109e247f.
//
// Solidity: event UnderlyingTokenDisabled(address indexed underlying, uint256 amountRedeemed, address yieldRecipient, uint256 yieldClaimed)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseUnderlyingTokenDisabled(log types.Log) (*YBConditionalTokensUnderlyingTokenDisabled, error) {
	event := new(YBConditionalTokensUnderlyingTokenDisabled)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "UnderlyingTokenDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensUnderlyingTokenEnabledIterator is returned from FilterUnderlyingTokenEnabled and is used to iterate over the raw logs and unpacked data for UnderlyingTokenEnabled events raised by the YBConditionalTokens contract.
type YBConditionalTokensUnderlyingTokenEnabledIterator struct {
	Event *YBConditionalTokensUnderlyingTokenEnabled // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensUnderlyingTokenEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensUnderlyingTokenEnabled)
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
		it.Event = new(YBConditionalTokensUnderlyingTokenEnabled)
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
func (it *YBConditionalTokensUnderlyingTokenEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensUnderlyingTokenEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensUnderlyingTokenEnabled represents a UnderlyingTokenEnabled event raised by the YBConditionalTokens contract.
type YBConditionalTokensUnderlyingTokenEnabled struct {
	Underlying common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnderlyingTokenEnabled is a free log retrieval operation binding the contract event 0x6ffa63b7792250c03119c3f600c10502d6bb232967842053a0510ac4ec1140f3.
//
// Solidity: event UnderlyingTokenEnabled(address indexed underlying, uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterUnderlyingTokenEnabled(opts *bind.FilterOpts, underlying []common.Address) (*YBConditionalTokensUnderlyingTokenEnabledIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "UnderlyingTokenEnabled", underlyingRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensUnderlyingTokenEnabledIterator{contract: _YBConditionalTokens.contract, event: "UnderlyingTokenEnabled", logs: logs, sub: sub}, nil
}

// WatchUnderlyingTokenEnabled is a free log subscription operation binding the contract event 0x6ffa63b7792250c03119c3f600c10502d6bb232967842053a0510ac4ec1140f3.
//
// Solidity: event UnderlyingTokenEnabled(address indexed underlying, uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchUnderlyingTokenEnabled(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensUnderlyingTokenEnabled, underlying []common.Address) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "UnderlyingTokenEnabled", underlyingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensUnderlyingTokenEnabled)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "UnderlyingTokenEnabled", log); err != nil {
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

// ParseUnderlyingTokenEnabled is a log parse operation binding the contract event 0x6ffa63b7792250c03119c3f600c10502d6bb232967842053a0510ac4ec1140f3.
//
// Solidity: event UnderlyingTokenEnabled(address indexed underlying, uint256 amount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseUnderlyingTokenEnabled(log types.Log) (*YBConditionalTokensUnderlyingTokenEnabled, error) {
	event := new(YBConditionalTokensUnderlyingTokenEnabled)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "UnderlyingTokenEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensUnderlyingTokenRedeemedIterator is returned from FilterUnderlyingTokenRedeemed and is used to iterate over the raw logs and unpacked data for UnderlyingTokenRedeemed events raised by the YBConditionalTokens contract.
type YBConditionalTokensUnderlyingTokenRedeemedIterator struct {
	Event *YBConditionalTokensUnderlyingTokenRedeemed // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensUnderlyingTokenRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensUnderlyingTokenRedeemed)
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
		it.Event = new(YBConditionalTokensUnderlyingTokenRedeemed)
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
func (it *YBConditionalTokensUnderlyingTokenRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensUnderlyingTokenRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensUnderlyingTokenRedeemed represents a UnderlyingTokenRedeemed event raised by the YBConditionalTokens contract.
type YBConditionalTokensUnderlyingTokenRedeemed struct {
	Underlying     common.Address
	VToken         common.Address
	AmountRedeemed *big.Int
	AmountReceived *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUnderlyingTokenRedeemed is a free log retrieval operation binding the contract event 0x1a90e48fb93dffb289624eeb1a2fb213e9d5e4a317088d8cf994d524139b62bc.
//
// Solidity: event UnderlyingTokenRedeemed(address indexed underlying, address indexed vToken, uint256 amountRedeemed, uint256 amountReceived)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterUnderlyingTokenRedeemed(opts *bind.FilterOpts, underlying []common.Address, vToken []common.Address) (*YBConditionalTokensUnderlyingTokenRedeemedIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "UnderlyingTokenRedeemed", underlyingRule, vTokenRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensUnderlyingTokenRedeemedIterator{contract: _YBConditionalTokens.contract, event: "UnderlyingTokenRedeemed", logs: logs, sub: sub}, nil
}

// WatchUnderlyingTokenRedeemed is a free log subscription operation binding the contract event 0x1a90e48fb93dffb289624eeb1a2fb213e9d5e4a317088d8cf994d524139b62bc.
//
// Solidity: event UnderlyingTokenRedeemed(address indexed underlying, address indexed vToken, uint256 amountRedeemed, uint256 amountReceived)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchUnderlyingTokenRedeemed(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensUnderlyingTokenRedeemed, underlying []common.Address, vToken []common.Address) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "UnderlyingTokenRedeemed", underlyingRule, vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensUnderlyingTokenRedeemed)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "UnderlyingTokenRedeemed", log); err != nil {
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

// ParseUnderlyingTokenRedeemed is a log parse operation binding the contract event 0x1a90e48fb93dffb289624eeb1a2fb213e9d5e4a317088d8cf994d524139b62bc.
//
// Solidity: event UnderlyingTokenRedeemed(address indexed underlying, address indexed vToken, uint256 amountRedeemed, uint256 amountReceived)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseUnderlyingTokenRedeemed(log types.Log) (*YBConditionalTokensUnderlyingTokenRedeemed, error) {
	event := new(YBConditionalTokensUnderlyingTokenRedeemed)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "UnderlyingTokenRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensVTokenMintedIterator is returned from FilterVTokenMinted and is used to iterate over the raw logs and unpacked data for VTokenMinted events raised by the YBConditionalTokens contract.
type YBConditionalTokensVTokenMintedIterator struct {
	Event *YBConditionalTokensVTokenMinted // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensVTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensVTokenMinted)
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
		it.Event = new(YBConditionalTokensVTokenMinted)
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
func (it *YBConditionalTokensVTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensVTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensVTokenMinted represents a VTokenMinted event raised by the YBConditionalTokens contract.
type YBConditionalTokensVTokenMinted struct {
	Underlying       common.Address
	VToken           common.Address
	UnderlyingAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVTokenMinted is a free log retrieval operation binding the contract event 0x2d3f40f5e9be9fe18e61a9d2f160a2bd4b75533ca481194f7c8ef886cbc1cfb9.
//
// Solidity: event VTokenMinted(address indexed underlying, address indexed vToken, uint256 underlyingAmount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterVTokenMinted(opts *bind.FilterOpts, underlying []common.Address, vToken []common.Address) (*YBConditionalTokensVTokenMintedIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "VTokenMinted", underlyingRule, vTokenRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensVTokenMintedIterator{contract: _YBConditionalTokens.contract, event: "VTokenMinted", logs: logs, sub: sub}, nil
}

// WatchVTokenMinted is a free log subscription operation binding the contract event 0x2d3f40f5e9be9fe18e61a9d2f160a2bd4b75533ca481194f7c8ef886cbc1cfb9.
//
// Solidity: event VTokenMinted(address indexed underlying, address indexed vToken, uint256 underlyingAmount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchVTokenMinted(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensVTokenMinted, underlying []common.Address, vToken []common.Address) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "VTokenMinted", underlyingRule, vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensVTokenMinted)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "VTokenMinted", log); err != nil {
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

// ParseVTokenMinted is a log parse operation binding the contract event 0x2d3f40f5e9be9fe18e61a9d2f160a2bd4b75533ca481194f7c8ef886cbc1cfb9.
//
// Solidity: event VTokenMinted(address indexed underlying, address indexed vToken, uint256 underlyingAmount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseVTokenMinted(log types.Log) (*YBConditionalTokensVTokenMinted, error) {
	event := new(YBConditionalTokensVTokenMinted)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "VTokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBConditionalTokensYieldClaimedIterator is returned from FilterYieldClaimed and is used to iterate over the raw logs and unpacked data for YieldClaimed events raised by the YBConditionalTokens contract.
type YBConditionalTokensYieldClaimedIterator struct {
	Event *YBConditionalTokensYieldClaimed // Event containing the contract specifics and raw log

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
func (it *YBConditionalTokensYieldClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBConditionalTokensYieldClaimed)
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
		it.Event = new(YBConditionalTokensYieldClaimed)
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
func (it *YBConditionalTokensYieldClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBConditionalTokensYieldClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBConditionalTokensYieldClaimed represents a YieldClaimed event raised by the YBConditionalTokens contract.
type YBConditionalTokensYieldClaimed struct {
	Underlying       common.Address
	VToken           common.Address
	VTokenAmount     *big.Int
	UnderlyingAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterYieldClaimed is a free log retrieval operation binding the contract event 0xb10c79aa8ce52424b7676b33fe7b45ea2516b75f43f3b5c0c730e52ca8ef36a5.
//
// Solidity: event YieldClaimed(address indexed underlying, address indexed vToken, uint256 vTokenAmount, uint256 underlyingAmount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) FilterYieldClaimed(opts *bind.FilterOpts, underlying []common.Address, vToken []common.Address) (*YBConditionalTokensYieldClaimedIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.FilterLogs(opts, "YieldClaimed", underlyingRule, vTokenRule)
	if err != nil {
		return nil, err
	}
	return &YBConditionalTokensYieldClaimedIterator{contract: _YBConditionalTokens.contract, event: "YieldClaimed", logs: logs, sub: sub}, nil
}

// WatchYieldClaimed is a free log subscription operation binding the contract event 0xb10c79aa8ce52424b7676b33fe7b45ea2516b75f43f3b5c0c730e52ca8ef36a5.
//
// Solidity: event YieldClaimed(address indexed underlying, address indexed vToken, uint256 vTokenAmount, uint256 underlyingAmount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) WatchYieldClaimed(opts *bind.WatchOpts, sink chan<- *YBConditionalTokensYieldClaimed, underlying []common.Address, vToken []common.Address) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _YBConditionalTokens.contract.WatchLogs(opts, "YieldClaimed", underlyingRule, vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBConditionalTokensYieldClaimed)
				if err := _YBConditionalTokens.contract.UnpackLog(event, "YieldClaimed", log); err != nil {
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

// ParseYieldClaimed is a log parse operation binding the contract event 0xb10c79aa8ce52424b7676b33fe7b45ea2516b75f43f3b5c0c730e52ca8ef36a5.
//
// Solidity: event YieldClaimed(address indexed underlying, address indexed vToken, uint256 vTokenAmount, uint256 underlyingAmount)
func (_YBConditionalTokens *YBConditionalTokensFilterer) ParseYieldClaimed(log types.Log) (*YBConditionalTokensYieldClaimed, error) {
	event := new(YBConditionalTokensYieldClaimed)
	if err := _YBConditionalTokens.contract.UnpackLog(event, "YieldClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
