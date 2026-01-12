// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yb_negrisk

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

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Salt          *big.Int
	Maker         common.Address
	Signer        common.Address
	Taker         common.Address
	TokenId       *big.Int
	MakerAmount   *big.Int
	TakerAmount   *big.Int
	Expiration    *big.Int
	Nonce         *big.Int
	FeeRateBps    *big.Int
	Side          uint8
	SignatureType uint8
	Signature     []byte
}

// OrderStatus is an auto generated low-level Go binding around an user-defined struct.
type OrderStatus struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}

// YBNegRiskMetaData contains all meta data concerning the YBNegRisk contract.
var YBNegRiskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_negRiskAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_safeFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidComplement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingGtRemaining\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedTokenIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCrossing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderFilledOrCancelled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLittleTokensReceived\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"takerOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"takerOrderMaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProxyFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProxyFactory\",\"type\":\"address\"}],\"name\":\"ProxyFactoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSafeFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSafeFactory\",\"type\":\"address\"}],\"name\":\"SafeFactoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"token0\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"token1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingUnpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fillAmounts\",\"type\":\"uint256[]\"}],\"name\":\"fillOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"getComplement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"getConditionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCtf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isFilledOrCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"internalType\":\"structOrderStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolyProxyFactoryImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPolyProxyWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getSafeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeFactoryImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"isValidNonce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"makerOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFillAmounts\",\"type\":\"uint256[]\"}],\"name\":\"matchOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFilledOrCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentCollectionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProxyFactory\",\"type\":\"address\"}],\"name\":\"setProxyFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSafeFactory\",\"type\":\"address\"}],\"name\":\"setSafeFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"}],\"name\":\"validateComplement\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"validateTokenId\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YBNegRiskABI is the input ABI used to generate the binding from.
// Deprecated: Use YBNegRiskMetaData.ABI instead.
var YBNegRiskABI = YBNegRiskMetaData.ABI

// YBNegRisk is an auto generated Go binding around an Ethereum contract.
type YBNegRisk struct {
	YBNegRiskCaller     // Read-only binding to the contract
	YBNegRiskTransactor // Write-only binding to the contract
	YBNegRiskFilterer   // Log filterer for contract events
}

// YBNegRiskCaller is an auto generated read-only Go binding around an Ethereum contract.
type YBNegRiskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBNegRiskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YBNegRiskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBNegRiskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YBNegRiskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YBNegRiskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YBNegRiskSession struct {
	Contract     *YBNegRisk        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YBNegRiskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YBNegRiskCallerSession struct {
	Contract *YBNegRiskCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// YBNegRiskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YBNegRiskTransactorSession struct {
	Contract     *YBNegRiskTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// YBNegRiskRaw is an auto generated low-level Go binding around an Ethereum contract.
type YBNegRiskRaw struct {
	Contract *YBNegRisk // Generic contract binding to access the raw methods on
}

// YBNegRiskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YBNegRiskCallerRaw struct {
	Contract *YBNegRiskCaller // Generic read-only contract binding to access the raw methods on
}

// YBNegRiskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YBNegRiskTransactorRaw struct {
	Contract *YBNegRiskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYBNegRisk creates a new instance of YBNegRisk, bound to a specific deployed contract.
func NewYBNegRisk(address common.Address, backend bind.ContractBackend) (*YBNegRisk, error) {
	contract, err := bindYBNegRisk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YBNegRisk{YBNegRiskCaller: YBNegRiskCaller{contract: contract}, YBNegRiskTransactor: YBNegRiskTransactor{contract: contract}, YBNegRiskFilterer: YBNegRiskFilterer{contract: contract}}, nil
}

// NewYBNegRiskCaller creates a new read-only instance of YBNegRisk, bound to a specific deployed contract.
func NewYBNegRiskCaller(address common.Address, caller bind.ContractCaller) (*YBNegRiskCaller, error) {
	contract, err := bindYBNegRisk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskCaller{contract: contract}, nil
}

// NewYBNegRiskTransactor creates a new write-only instance of YBNegRisk, bound to a specific deployed contract.
func NewYBNegRiskTransactor(address common.Address, transactor bind.ContractTransactor) (*YBNegRiskTransactor, error) {
	contract, err := bindYBNegRisk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskTransactor{contract: contract}, nil
}

// NewYBNegRiskFilterer creates a new log filterer instance of YBNegRisk, bound to a specific deployed contract.
func NewYBNegRiskFilterer(address common.Address, filterer bind.ContractFilterer) (*YBNegRiskFilterer, error) {
	contract, err := bindYBNegRisk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskFilterer{contract: contract}, nil
}

// bindYBNegRisk binds a generic wrapper to an already deployed contract.
func bindYBNegRisk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YBNegRiskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YBNegRisk *YBNegRiskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YBNegRisk.Contract.YBNegRiskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YBNegRisk *YBNegRiskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRisk.Contract.YBNegRiskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YBNegRisk *YBNegRiskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YBNegRisk.Contract.YBNegRiskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YBNegRisk *YBNegRiskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YBNegRisk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YBNegRisk *YBNegRiskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRisk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YBNegRisk *YBNegRiskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YBNegRisk.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_YBNegRisk *YBNegRiskCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_YBNegRisk *YBNegRiskSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _YBNegRisk.Contract.Admins(&_YBNegRisk.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_YBNegRisk *YBNegRiskCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _YBNegRisk.Contract.Admins(&_YBNegRisk.CallOpts, arg0)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_YBNegRisk *YBNegRiskCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_YBNegRisk *YBNegRiskSession) DomainSeparator() ([32]byte, error) {
	return _YBNegRisk.Contract.DomainSeparator(&_YBNegRisk.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_YBNegRisk *YBNegRiskCallerSession) DomainSeparator() ([32]byte, error) {
	return _YBNegRisk.Contract.DomainSeparator(&_YBNegRisk.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_YBNegRisk *YBNegRiskCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_YBNegRisk *YBNegRiskSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _YBNegRisk.Contract.Eip712Domain(&_YBNegRisk.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_YBNegRisk *YBNegRiskCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _YBNegRisk.Contract.Eip712Domain(&_YBNegRisk.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_YBNegRisk *YBNegRiskCaller) GetCollateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getCollateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_YBNegRisk *YBNegRiskSession) GetCollateral() (common.Address, error) {
	return _YBNegRisk.Contract.GetCollateral(&_YBNegRisk.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) GetCollateral() (common.Address, error) {
	return _YBNegRisk.Contract.GetCollateral(&_YBNegRisk.CallOpts)
}

// GetComplement is a free data retrieval call binding the contract method 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (_YBNegRisk *YBNegRiskCaller) GetComplement(opts *bind.CallOpts, token *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getComplement", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetComplement is a free data retrieval call binding the contract method 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (_YBNegRisk *YBNegRiskSession) GetComplement(token *big.Int) (*big.Int, error) {
	return _YBNegRisk.Contract.GetComplement(&_YBNegRisk.CallOpts, token)
}

// GetComplement is a free data retrieval call binding the contract method 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (_YBNegRisk *YBNegRiskCallerSession) GetComplement(token *big.Int) (*big.Int, error) {
	return _YBNegRisk.Contract.GetComplement(&_YBNegRisk.CallOpts, token)
}

// GetConditionId is a free data retrieval call binding the contract method 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (_YBNegRisk *YBNegRiskCaller) GetConditionId(opts *bind.CallOpts, token *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getConditionId", token)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConditionId is a free data retrieval call binding the contract method 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (_YBNegRisk *YBNegRiskSession) GetConditionId(token *big.Int) ([32]byte, error) {
	return _YBNegRisk.Contract.GetConditionId(&_YBNegRisk.CallOpts, token)
}

// GetConditionId is a free data retrieval call binding the contract method 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (_YBNegRisk *YBNegRiskCallerSession) GetConditionId(token *big.Int) ([32]byte, error) {
	return _YBNegRisk.Contract.GetConditionId(&_YBNegRisk.CallOpts, token)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_YBNegRisk *YBNegRiskCaller) GetCtf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getCtf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_YBNegRisk *YBNegRiskSession) GetCtf() (common.Address, error) {
	return _YBNegRisk.Contract.GetCtf(&_YBNegRisk.CallOpts)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) GetCtf() (common.Address, error) {
	return _YBNegRisk.Contract.GetCtf(&_YBNegRisk.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (_YBNegRisk *YBNegRiskCaller) GetMaxFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getMaxFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (_YBNegRisk *YBNegRiskSession) GetMaxFeeRate() (*big.Int, error) {
	return _YBNegRisk.Contract.GetMaxFeeRate(&_YBNegRisk.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (_YBNegRisk *YBNegRiskCallerSession) GetMaxFeeRate() (*big.Int, error) {
	return _YBNegRisk.Contract.GetMaxFeeRate(&_YBNegRisk.CallOpts)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (_YBNegRisk *YBNegRiskCaller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (OrderStatus, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getOrderStatus", orderHash)

	if err != nil {
		return *new(OrderStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderStatus)).(*OrderStatus)

	return out0, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (_YBNegRisk *YBNegRiskSession) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _YBNegRisk.Contract.GetOrderStatus(&_YBNegRisk.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (_YBNegRisk *YBNegRiskCallerSession) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _YBNegRisk.Contract.GetOrderStatus(&_YBNegRisk.CallOpts, orderHash)
}

// GetPolyProxyFactoryImplementation is a free data retrieval call binding the contract method 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (_YBNegRisk *YBNegRiskCaller) GetPolyProxyFactoryImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getPolyProxyFactoryImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPolyProxyFactoryImplementation is a free data retrieval call binding the contract method 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (_YBNegRisk *YBNegRiskSession) GetPolyProxyFactoryImplementation() (common.Address, error) {
	return _YBNegRisk.Contract.GetPolyProxyFactoryImplementation(&_YBNegRisk.CallOpts)
}

// GetPolyProxyFactoryImplementation is a free data retrieval call binding the contract method 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) GetPolyProxyFactoryImplementation() (common.Address, error) {
	return _YBNegRisk.Contract.GetPolyProxyFactoryImplementation(&_YBNegRisk.CallOpts)
}

// GetPolyProxyWalletAddress is a free data retrieval call binding the contract method 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (_YBNegRisk *YBNegRiskCaller) GetPolyProxyWalletAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getPolyProxyWalletAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPolyProxyWalletAddress is a free data retrieval call binding the contract method 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (_YBNegRisk *YBNegRiskSession) GetPolyProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _YBNegRisk.Contract.GetPolyProxyWalletAddress(&_YBNegRisk.CallOpts, _addr)
}

// GetPolyProxyWalletAddress is a free data retrieval call binding the contract method 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) GetPolyProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _YBNegRisk.Contract.GetPolyProxyWalletAddress(&_YBNegRisk.CallOpts, _addr)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_YBNegRisk *YBNegRiskCaller) GetProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getProxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_YBNegRisk *YBNegRiskSession) GetProxyFactory() (common.Address, error) {
	return _YBNegRisk.Contract.GetProxyFactory(&_YBNegRisk.CallOpts)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) GetProxyFactory() (common.Address, error) {
	return _YBNegRisk.Contract.GetProxyFactory(&_YBNegRisk.CallOpts)
}

// GetSafeAddress is a free data retrieval call binding the contract method 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (_YBNegRisk *YBNegRiskCaller) GetSafeAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getSafeAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeAddress is a free data retrieval call binding the contract method 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (_YBNegRisk *YBNegRiskSession) GetSafeAddress(_addr common.Address) (common.Address, error) {
	return _YBNegRisk.Contract.GetSafeAddress(&_YBNegRisk.CallOpts, _addr)
}

// GetSafeAddress is a free data retrieval call binding the contract method 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) GetSafeAddress(_addr common.Address) (common.Address, error) {
	return _YBNegRisk.Contract.GetSafeAddress(&_YBNegRisk.CallOpts, _addr)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_YBNegRisk *YBNegRiskCaller) GetSafeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getSafeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_YBNegRisk *YBNegRiskSession) GetSafeFactory() (common.Address, error) {
	return _YBNegRisk.Contract.GetSafeFactory(&_YBNegRisk.CallOpts)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) GetSafeFactory() (common.Address, error) {
	return _YBNegRisk.Contract.GetSafeFactory(&_YBNegRisk.CallOpts)
}

// GetSafeFactoryImplementation is a free data retrieval call binding the contract method 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (_YBNegRisk *YBNegRiskCaller) GetSafeFactoryImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "getSafeFactoryImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeFactoryImplementation is a free data retrieval call binding the contract method 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (_YBNegRisk *YBNegRiskSession) GetSafeFactoryImplementation() (common.Address, error) {
	return _YBNegRisk.Contract.GetSafeFactoryImplementation(&_YBNegRisk.CallOpts)
}

// GetSafeFactoryImplementation is a free data retrieval call binding the contract method 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) GetSafeFactoryImplementation() (common.Address, error) {
	return _YBNegRisk.Contract.GetSafeFactoryImplementation(&_YBNegRisk.CallOpts)
}

// HashOrder is a free data retrieval call binding the contract method 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (_YBNegRisk *YBNegRiskCaller) HashOrder(opts *bind.CallOpts, order Order) ([32]byte, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (_YBNegRisk *YBNegRiskSession) HashOrder(order Order) ([32]byte, error) {
	return _YBNegRisk.Contract.HashOrder(&_YBNegRisk.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (_YBNegRisk *YBNegRiskCallerSession) HashOrder(order Order) ([32]byte, error) {
	return _YBNegRisk.Contract.HashOrder(&_YBNegRisk.CallOpts, order)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_YBNegRisk *YBNegRiskCaller) IsAdmin(opts *bind.CallOpts, usr common.Address) (bool, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "isAdmin", usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_YBNegRisk *YBNegRiskSession) IsAdmin(usr common.Address) (bool, error) {
	return _YBNegRisk.Contract.IsAdmin(&_YBNegRisk.CallOpts, usr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_YBNegRisk *YBNegRiskCallerSession) IsAdmin(usr common.Address) (bool, error) {
	return _YBNegRisk.Contract.IsAdmin(&_YBNegRisk.CallOpts, usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_YBNegRisk *YBNegRiskCaller) IsOperator(opts *bind.CallOpts, usr common.Address) (bool, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "isOperator", usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_YBNegRisk *YBNegRiskSession) IsOperator(usr common.Address) (bool, error) {
	return _YBNegRisk.Contract.IsOperator(&_YBNegRisk.CallOpts, usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_YBNegRisk *YBNegRiskCallerSession) IsOperator(usr common.Address) (bool, error) {
	return _YBNegRisk.Contract.IsOperator(&_YBNegRisk.CallOpts, usr)
}

// IsValidNonce is a free data retrieval call binding the contract method 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (_YBNegRisk *YBNegRiskCaller) IsValidNonce(opts *bind.CallOpts, usr common.Address, nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "isValidNonce", usr, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidNonce is a free data retrieval call binding the contract method 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (_YBNegRisk *YBNegRiskSession) IsValidNonce(usr common.Address, nonce *big.Int) (bool, error) {
	return _YBNegRisk.Contract.IsValidNonce(&_YBNegRisk.CallOpts, usr, nonce)
}

// IsValidNonce is a free data retrieval call binding the contract method 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (_YBNegRisk *YBNegRiskCallerSession) IsValidNonce(usr common.Address, nonce *big.Int) (bool, error) {
	return _YBNegRisk.Contract.IsValidNonce(&_YBNegRisk.CallOpts, usr, nonce)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_YBNegRisk *YBNegRiskCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_YBNegRisk *YBNegRiskSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _YBNegRisk.Contract.Nonces(&_YBNegRisk.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_YBNegRisk *YBNegRiskCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _YBNegRisk.Contract.Nonces(&_YBNegRisk.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_YBNegRisk *YBNegRiskCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_YBNegRisk *YBNegRiskSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _YBNegRisk.Contract.Operators(&_YBNegRisk.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_YBNegRisk *YBNegRiskCallerSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _YBNegRisk.Contract.Operators(&_YBNegRisk.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (_YBNegRisk *YBNegRiskCaller) OrderStatus(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "orderStatus", arg0)

	outstruct := new(struct {
		IsFilledOrCancelled bool
		Remaining           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsFilledOrCancelled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Remaining = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (_YBNegRisk *YBNegRiskSession) OrderStatus(arg0 [32]byte) (struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}, error) {
	return _YBNegRisk.Contract.OrderStatus(&_YBNegRisk.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (_YBNegRisk *YBNegRiskCallerSession) OrderStatus(arg0 [32]byte) (struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}, error) {
	return _YBNegRisk.Contract.OrderStatus(&_YBNegRisk.CallOpts, arg0)
}

// ParentCollectionId is a free data retrieval call binding the contract method 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (_YBNegRisk *YBNegRiskCaller) ParentCollectionId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "parentCollectionId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParentCollectionId is a free data retrieval call binding the contract method 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (_YBNegRisk *YBNegRiskSession) ParentCollectionId() ([32]byte, error) {
	return _YBNegRisk.Contract.ParentCollectionId(&_YBNegRisk.CallOpts)
}

// ParentCollectionId is a free data retrieval call binding the contract method 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (_YBNegRisk *YBNegRiskCallerSession) ParentCollectionId() ([32]byte, error) {
	return _YBNegRisk.Contract.ParentCollectionId(&_YBNegRisk.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YBNegRisk *YBNegRiskCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YBNegRisk *YBNegRiskSession) Paused() (bool, error) {
	return _YBNegRisk.Contract.Paused(&_YBNegRisk.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_YBNegRisk *YBNegRiskCallerSession) Paused() (bool, error) {
	return _YBNegRisk.Contract.Paused(&_YBNegRisk.CallOpts)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_YBNegRisk *YBNegRiskCaller) ProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "proxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_YBNegRisk *YBNegRiskSession) ProxyFactory() (common.Address, error) {
	return _YBNegRisk.Contract.ProxyFactory(&_YBNegRisk.CallOpts)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) ProxyFactory() (common.Address, error) {
	return _YBNegRisk.Contract.ProxyFactory(&_YBNegRisk.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (_YBNegRisk *YBNegRiskCaller) Registry(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Complement  *big.Int
	ConditionId [32]byte
}, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "registry", arg0)

	outstruct := new(struct {
		Complement  *big.Int
		ConditionId [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Complement = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ConditionId = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Registry is a free data retrieval call binding the contract method 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (_YBNegRisk *YBNegRiskSession) Registry(arg0 *big.Int) (struct {
	Complement  *big.Int
	ConditionId [32]byte
}, error) {
	return _YBNegRisk.Contract.Registry(&_YBNegRisk.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (_YBNegRisk *YBNegRiskCallerSession) Registry(arg0 *big.Int) (struct {
	Complement  *big.Int
	ConditionId [32]byte
}, error) {
	return _YBNegRisk.Contract.Registry(&_YBNegRisk.CallOpts, arg0)
}

// SafeFactory is a free data retrieval call binding the contract method 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (_YBNegRisk *YBNegRiskCaller) SafeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "safeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeFactory is a free data retrieval call binding the contract method 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (_YBNegRisk *YBNegRiskSession) SafeFactory() (common.Address, error) {
	return _YBNegRisk.Contract.SafeFactory(&_YBNegRisk.CallOpts)
}

// SafeFactory is a free data retrieval call binding the contract method 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (_YBNegRisk *YBNegRiskCallerSession) SafeFactory() (common.Address, error) {
	return _YBNegRisk.Contract.SafeFactory(&_YBNegRisk.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YBNegRisk *YBNegRiskCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YBNegRisk *YBNegRiskSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _YBNegRisk.Contract.SupportsInterface(&_YBNegRisk.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_YBNegRisk *YBNegRiskCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _YBNegRisk.Contract.SupportsInterface(&_YBNegRisk.CallOpts, interfaceId)
}

// ValidateComplement is a free data retrieval call binding the contract method 0xd82da838.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (_YBNegRisk *YBNegRiskCaller) ValidateComplement(opts *bind.CallOpts, token *big.Int, complement *big.Int) error {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "validateComplement", token, complement)

	if err != nil {
		return err
	}

	return err

}

// ValidateComplement is a free data retrieval call binding the contract method 0xd82da838.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (_YBNegRisk *YBNegRiskSession) ValidateComplement(token *big.Int, complement *big.Int) error {
	return _YBNegRisk.Contract.ValidateComplement(&_YBNegRisk.CallOpts, token, complement)
}

// ValidateComplement is a free data retrieval call binding the contract method 0xd82da838.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (_YBNegRisk *YBNegRiskCallerSession) ValidateComplement(token *big.Int, complement *big.Int) error {
	return _YBNegRisk.Contract.ValidateComplement(&_YBNegRisk.CallOpts, token, complement)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x654f0ce4.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_YBNegRisk *YBNegRiskCaller) ValidateOrder(opts *bind.CallOpts, order Order) error {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "validateOrder", order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrder is a free data retrieval call binding the contract method 0x654f0ce4.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_YBNegRisk *YBNegRiskSession) ValidateOrder(order Order) error {
	return _YBNegRisk.Contract.ValidateOrder(&_YBNegRisk.CallOpts, order)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x654f0ce4.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_YBNegRisk *YBNegRiskCallerSession) ValidateOrder(order Order) error {
	return _YBNegRisk.Contract.ValidateOrder(&_YBNegRisk.CallOpts, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xe2eec405.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_YBNegRisk *YBNegRiskCaller) ValidateOrderSignature(opts *bind.CallOpts, orderHash [32]byte, order Order) error {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "validateOrderSignature", orderHash, order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xe2eec405.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_YBNegRisk *YBNegRiskSession) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _YBNegRisk.Contract.ValidateOrderSignature(&_YBNegRisk.CallOpts, orderHash, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xe2eec405.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_YBNegRisk *YBNegRiskCallerSession) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _YBNegRisk.Contract.ValidateOrderSignature(&_YBNegRisk.CallOpts, orderHash, order)
}

// ValidateTokenId is a free data retrieval call binding the contract method 0x34600901.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (_YBNegRisk *YBNegRiskCaller) ValidateTokenId(opts *bind.CallOpts, tokenId *big.Int) error {
	var out []interface{}
	err := _YBNegRisk.contract.Call(opts, &out, "validateTokenId", tokenId)

	if err != nil {
		return err
	}

	return err

}

// ValidateTokenId is a free data retrieval call binding the contract method 0x34600901.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (_YBNegRisk *YBNegRiskSession) ValidateTokenId(tokenId *big.Int) error {
	return _YBNegRisk.Contract.ValidateTokenId(&_YBNegRisk.CallOpts, tokenId)
}

// ValidateTokenId is a free data retrieval call binding the contract method 0x34600901.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (_YBNegRisk *YBNegRiskCallerSession) ValidateTokenId(tokenId *big.Int) error {
	return _YBNegRisk.Contract.ValidateTokenId(&_YBNegRisk.CallOpts, tokenId)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_YBNegRisk *YBNegRiskTransactor) AddAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "addAdmin", admin_)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_YBNegRisk *YBNegRiskSession) AddAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.AddAdmin(&_YBNegRisk.TransactOpts, admin_)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) AddAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.AddAdmin(&_YBNegRisk.TransactOpts, admin_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_YBNegRisk *YBNegRiskTransactor) AddOperator(opts *bind.TransactOpts, operator_ common.Address) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "addOperator", operator_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_YBNegRisk *YBNegRiskSession) AddOperator(operator_ common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.AddOperator(&_YBNegRisk.TransactOpts, operator_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) AddOperator(operator_ common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.AddOperator(&_YBNegRisk.TransactOpts, operator_)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa6dfcf86.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (_YBNegRisk *YBNegRiskTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa6dfcf86.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (_YBNegRisk *YBNegRiskSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _YBNegRisk.Contract.CancelOrder(&_YBNegRisk.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa6dfcf86.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _YBNegRisk.Contract.CancelOrder(&_YBNegRisk.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xfa950b48.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (_YBNegRisk *YBNegRiskTransactor) CancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xfa950b48.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (_YBNegRisk *YBNegRiskSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _YBNegRisk.Contract.CancelOrders(&_YBNegRisk.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xfa950b48.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _YBNegRisk.Contract.CancelOrders(&_YBNegRisk.TransactOpts, orders)
}

// FillOrder is a paid mutator transaction binding the contract method 0xfe729aaf.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (_YBNegRisk *YBNegRiskTransactor) FillOrder(opts *bind.TransactOpts, order Order, fillAmount *big.Int) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "fillOrder", order, fillAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0xfe729aaf.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (_YBNegRisk *YBNegRiskSession) FillOrder(order Order, fillAmount *big.Int) (*types.Transaction, error) {
	return _YBNegRisk.Contract.FillOrder(&_YBNegRisk.TransactOpts, order, fillAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0xfe729aaf.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) FillOrder(order Order, fillAmount *big.Int) (*types.Transaction, error) {
	return _YBNegRisk.Contract.FillOrder(&_YBNegRisk.TransactOpts, order, fillAmount)
}

// FillOrders is a paid mutator transaction binding the contract method 0xd798eff6.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (_YBNegRisk *YBNegRiskTransactor) FillOrders(opts *bind.TransactOpts, orders []Order, fillAmounts []*big.Int) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "fillOrders", orders, fillAmounts)
}

// FillOrders is a paid mutator transaction binding the contract method 0xd798eff6.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (_YBNegRisk *YBNegRiskSession) FillOrders(orders []Order, fillAmounts []*big.Int) (*types.Transaction, error) {
	return _YBNegRisk.Contract.FillOrders(&_YBNegRisk.TransactOpts, orders, fillAmounts)
}

// FillOrders is a paid mutator transaction binding the contract method 0xd798eff6.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) FillOrders(orders []Order, fillAmounts []*big.Int) (*types.Transaction, error) {
	return _YBNegRisk.Contract.FillOrders(&_YBNegRisk.TransactOpts, orders, fillAmounts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_YBNegRisk *YBNegRiskTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_YBNegRisk *YBNegRiskSession) IncrementNonce() (*types.Transaction, error) {
	return _YBNegRisk.Contract.IncrementNonce(&_YBNegRisk.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_YBNegRisk *YBNegRiskTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _YBNegRisk.Contract.IncrementNonce(&_YBNegRisk.TransactOpts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xe60f0c05.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (_YBNegRisk *YBNegRiskTransactor) MatchOrders(opts *bind.TransactOpts, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "matchOrders", takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xe60f0c05.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (_YBNegRisk *YBNegRiskSession) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) (*types.Transaction, error) {
	return _YBNegRisk.Contract.MatchOrders(&_YBNegRisk.TransactOpts, takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xe60f0c05.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) (*types.Transaction, error) {
	return _YBNegRisk.Contract.MatchOrders(&_YBNegRisk.TransactOpts, takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_YBNegRisk *YBNegRiskTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_YBNegRisk *YBNegRiskSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRisk.Contract.OnERC1155BatchReceived(&_YBNegRisk.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_YBNegRisk *YBNegRiskTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRisk.Contract.OnERC1155BatchReceived(&_YBNegRisk.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_YBNegRisk *YBNegRiskTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_YBNegRisk *YBNegRiskSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRisk.Contract.OnERC1155Received(&_YBNegRisk.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_YBNegRisk *YBNegRiskTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _YBNegRisk.Contract.OnERC1155Received(&_YBNegRisk.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_YBNegRisk *YBNegRiskTransactor) PauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "pauseTrading")
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_YBNegRisk *YBNegRiskSession) PauseTrading() (*types.Transaction, error) {
	return _YBNegRisk.Contract.PauseTrading(&_YBNegRisk.TransactOpts)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_YBNegRisk *YBNegRiskTransactorSession) PauseTrading() (*types.Transaction, error) {
	return _YBNegRisk.Contract.PauseTrading(&_YBNegRisk.TransactOpts)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x68c7450f.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (_YBNegRisk *YBNegRiskTransactor) RegisterToken(opts *bind.TransactOpts, token *big.Int, complement *big.Int, conditionId [32]byte) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "registerToken", token, complement, conditionId)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x68c7450f.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (_YBNegRisk *YBNegRiskSession) RegisterToken(token *big.Int, complement *big.Int, conditionId [32]byte) (*types.Transaction, error) {
	return _YBNegRisk.Contract.RegisterToken(&_YBNegRisk.TransactOpts, token, complement, conditionId)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x68c7450f.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) RegisterToken(token *big.Int, complement *big.Int, conditionId [32]byte) (*types.Transaction, error) {
	return _YBNegRisk.Contract.RegisterToken(&_YBNegRisk.TransactOpts, token, complement, conditionId)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_YBNegRisk *YBNegRiskTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_YBNegRisk *YBNegRiskSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.RemoveAdmin(&_YBNegRisk.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.RemoveAdmin(&_YBNegRisk.TransactOpts, admin)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_YBNegRisk *YBNegRiskTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_YBNegRisk *YBNegRiskSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.RemoveOperator(&_YBNegRisk.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.RemoveOperator(&_YBNegRisk.TransactOpts, operator)
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_YBNegRisk *YBNegRiskTransactor) RenounceAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "renounceAdminRole")
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_YBNegRisk *YBNegRiskSession) RenounceAdminRole() (*types.Transaction, error) {
	return _YBNegRisk.Contract.RenounceAdminRole(&_YBNegRisk.TransactOpts)
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_YBNegRisk *YBNegRiskTransactorSession) RenounceAdminRole() (*types.Transaction, error) {
	return _YBNegRisk.Contract.RenounceAdminRole(&_YBNegRisk.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_YBNegRisk *YBNegRiskTransactor) RenounceOperatorRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "renounceOperatorRole")
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_YBNegRisk *YBNegRiskSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _YBNegRisk.Contract.RenounceOperatorRole(&_YBNegRisk.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_YBNegRisk *YBNegRiskTransactorSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _YBNegRisk.Contract.RenounceOperatorRole(&_YBNegRisk.TransactOpts)
}

// SetProxyFactory is a paid mutator transaction binding the contract method 0xfbddd751.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (_YBNegRisk *YBNegRiskTransactor) SetProxyFactory(opts *bind.TransactOpts, _newProxyFactory common.Address) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "setProxyFactory", _newProxyFactory)
}

// SetProxyFactory is a paid mutator transaction binding the contract method 0xfbddd751.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (_YBNegRisk *YBNegRiskSession) SetProxyFactory(_newProxyFactory common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.SetProxyFactory(&_YBNegRisk.TransactOpts, _newProxyFactory)
}

// SetProxyFactory is a paid mutator transaction binding the contract method 0xfbddd751.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) SetProxyFactory(_newProxyFactory common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.SetProxyFactory(&_YBNegRisk.TransactOpts, _newProxyFactory)
}

// SetSafeFactory is a paid mutator transaction binding the contract method 0x4544f055.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (_YBNegRisk *YBNegRiskTransactor) SetSafeFactory(opts *bind.TransactOpts, _newSafeFactory common.Address) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "setSafeFactory", _newSafeFactory)
}

// SetSafeFactory is a paid mutator transaction binding the contract method 0x4544f055.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (_YBNegRisk *YBNegRiskSession) SetSafeFactory(_newSafeFactory common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.SetSafeFactory(&_YBNegRisk.TransactOpts, _newSafeFactory)
}

// SetSafeFactory is a paid mutator transaction binding the contract method 0x4544f055.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (_YBNegRisk *YBNegRiskTransactorSession) SetSafeFactory(_newSafeFactory common.Address) (*types.Transaction, error) {
	return _YBNegRisk.Contract.SetSafeFactory(&_YBNegRisk.TransactOpts, _newSafeFactory)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_YBNegRisk *YBNegRiskTransactor) UnpauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YBNegRisk.contract.Transact(opts, "unpauseTrading")
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_YBNegRisk *YBNegRiskSession) UnpauseTrading() (*types.Transaction, error) {
	return _YBNegRisk.Contract.UnpauseTrading(&_YBNegRisk.TransactOpts)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_YBNegRisk *YBNegRiskTransactorSession) UnpauseTrading() (*types.Transaction, error) {
	return _YBNegRisk.Contract.UnpauseTrading(&_YBNegRisk.TransactOpts)
}

// YBNegRiskEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the YBNegRisk contract.
type YBNegRiskEIP712DomainChangedIterator struct {
	Event *YBNegRiskEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *YBNegRiskEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskEIP712DomainChanged)
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
		it.Event = new(YBNegRiskEIP712DomainChanged)
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
func (it *YBNegRiskEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskEIP712DomainChanged represents a EIP712DomainChanged event raised by the YBNegRisk contract.
type YBNegRiskEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_YBNegRisk *YBNegRiskFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*YBNegRiskEIP712DomainChangedIterator, error) {

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &YBNegRiskEIP712DomainChangedIterator{contract: _YBNegRisk.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_YBNegRisk *YBNegRiskFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *YBNegRiskEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskEIP712DomainChanged)
				if err := _YBNegRisk.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_YBNegRisk *YBNegRiskFilterer) ParseEIP712DomainChanged(log types.Log) (*YBNegRiskEIP712DomainChanged, error) {
	event := new(YBNegRiskEIP712DomainChanged)
	if err := _YBNegRisk.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskFeeChargedIterator is returned from FilterFeeCharged and is used to iterate over the raw logs and unpacked data for FeeCharged events raised by the YBNegRisk contract.
type YBNegRiskFeeChargedIterator struct {
	Event *YBNegRiskFeeCharged // Event containing the contract specifics and raw log

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
func (it *YBNegRiskFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskFeeCharged)
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
		it.Event = new(YBNegRiskFeeCharged)
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
func (it *YBNegRiskFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskFeeCharged represents a FeeCharged event raised by the YBNegRisk contract.
type YBNegRiskFeeCharged struct {
	Receiver common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeCharged is a free log retrieval operation binding the contract event 0xacffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (_YBNegRisk *YBNegRiskFilterer) FilterFeeCharged(opts *bind.FilterOpts, receiver []common.Address) (*YBNegRiskFeeChargedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskFeeChargedIterator{contract: _YBNegRisk.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

// WatchFeeCharged is a free log subscription operation binding the contract event 0xacffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (_YBNegRisk *YBNegRiskFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *YBNegRiskFeeCharged, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskFeeCharged)
				if err := _YBNegRisk.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

// ParseFeeCharged is a log parse operation binding the contract event 0xacffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (_YBNegRisk *YBNegRiskFilterer) ParseFeeCharged(log types.Log) (*YBNegRiskFeeCharged, error) {
	event := new(YBNegRiskFeeCharged)
	if err := _YBNegRisk.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the YBNegRisk contract.
type YBNegRiskNewAdminIterator struct {
	Event *YBNegRiskNewAdmin // Event containing the contract specifics and raw log

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
func (it *YBNegRiskNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskNewAdmin)
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
		it.Event = new(YBNegRiskNewAdmin)
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
func (it *YBNegRiskNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskNewAdmin represents a NewAdmin event raised by the YBNegRisk contract.
type YBNegRiskNewAdmin struct {
	NewAdminAddress common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_YBNegRisk *YBNegRiskFilterer) FilterNewAdmin(opts *bind.FilterOpts, newAdminAddress []common.Address, admin []common.Address) (*YBNegRiskNewAdminIterator, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskNewAdminIterator{contract: _YBNegRisk.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_YBNegRisk *YBNegRiskFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *YBNegRiskNewAdmin, newAdminAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskNewAdmin)
				if err := _YBNegRisk.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_YBNegRisk *YBNegRiskFilterer) ParseNewAdmin(log types.Log) (*YBNegRiskNewAdmin, error) {
	event := new(YBNegRiskNewAdmin)
	if err := _YBNegRisk.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskNewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the YBNegRisk contract.
type YBNegRiskNewOperatorIterator struct {
	Event *YBNegRiskNewOperator // Event containing the contract specifics and raw log

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
func (it *YBNegRiskNewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskNewOperator)
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
		it.Event = new(YBNegRiskNewOperator)
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
func (it *YBNegRiskNewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskNewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskNewOperator represents a NewOperator event raised by the YBNegRisk contract.
type YBNegRiskNewOperator struct {
	NewOperatorAddress common.Address
	Admin              common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_YBNegRisk *YBNegRiskFilterer) FilterNewOperator(opts *bind.FilterOpts, newOperatorAddress []common.Address, admin []common.Address) (*YBNegRiskNewOperatorIterator, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskNewOperatorIterator{contract: _YBNegRisk.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_YBNegRisk *YBNegRiskFilterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *YBNegRiskNewOperator, newOperatorAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskNewOperator)
				if err := _YBNegRisk.contract.UnpackLog(event, "NewOperator", log); err != nil {
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
func (_YBNegRisk *YBNegRiskFilterer) ParseNewOperator(log types.Log) (*YBNegRiskNewOperator, error) {
	event := new(YBNegRiskNewOperator)
	if err := _YBNegRisk.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskNonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the YBNegRisk contract.
type YBNegRiskNonceIncrementedIterator struct {
	Event *YBNegRiskNonceIncremented // Event containing the contract specifics and raw log

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
func (it *YBNegRiskNonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskNonceIncremented)
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
		it.Event = new(YBNegRiskNonceIncremented)
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
func (it *YBNegRiskNonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskNonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskNonceIncremented represents a NonceIncremented event raised by the YBNegRisk contract.
type YBNegRiskNonceIncremented struct {
	User     common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed user, uint256 newNonce)
func (_YBNegRisk *YBNegRiskFilterer) FilterNonceIncremented(opts *bind.FilterOpts, user []common.Address) (*YBNegRiskNonceIncrementedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "NonceIncremented", userRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskNonceIncrementedIterator{contract: _YBNegRisk.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed user, uint256 newNonce)
func (_YBNegRisk *YBNegRiskFilterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *YBNegRiskNonceIncremented, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "NonceIncremented", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskNonceIncremented)
				if err := _YBNegRisk.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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

// ParseNonceIncremented is a log parse operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed user, uint256 newNonce)
func (_YBNegRisk *YBNegRiskFilterer) ParseNonceIncremented(log types.Log) (*YBNegRiskNonceIncremented, error) {
	event := new(YBNegRiskNonceIncremented)
	if err := _YBNegRisk.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the YBNegRisk contract.
type YBNegRiskOrderCancelledIterator struct {
	Event *YBNegRiskOrderCancelled // Event containing the contract specifics and raw log

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
func (it *YBNegRiskOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskOrderCancelled)
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
		it.Event = new(YBNegRiskOrderCancelled)
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
func (it *YBNegRiskOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskOrderCancelled represents a OrderCancelled event raised by the YBNegRisk contract.
type YBNegRiskOrderCancelled struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (_YBNegRisk *YBNegRiskFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderHash [][32]byte) (*YBNegRiskOrderCancelledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "OrderCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskOrderCancelledIterator{contract: _YBNegRisk.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (_YBNegRisk *YBNegRiskFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *YBNegRiskOrderCancelled, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "OrderCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskOrderCancelled)
				if err := _YBNegRisk.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (_YBNegRisk *YBNegRiskFilterer) ParseOrderCancelled(log types.Log) (*YBNegRiskOrderCancelled, error) {
	event := new(YBNegRiskOrderCancelled)
	if err := _YBNegRisk.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the YBNegRisk contract.
type YBNegRiskOrderFilledIterator struct {
	Event *YBNegRiskOrderFilled // Event containing the contract specifics and raw log

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
func (it *YBNegRiskOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskOrderFilled)
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
		it.Event = new(YBNegRiskOrderFilled)
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
func (it *YBNegRiskOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskOrderFilled represents a OrderFilled event raised by the YBNegRisk contract.
type YBNegRiskOrderFilled struct {
	OrderHash         [32]byte
	Maker             common.Address
	Taker             common.Address
	MakerAssetId      *big.Int
	TakerAssetId      *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Fee               *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (_YBNegRisk *YBNegRiskFilterer) FilterOrderFilled(opts *bind.FilterOpts, orderHash [][32]byte, maker []common.Address, taker []common.Address) (*YBNegRiskOrderFilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskOrderFilledIterator{contract: _YBNegRisk.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (_YBNegRisk *YBNegRiskFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *YBNegRiskOrderFilled, orderHash [][32]byte, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskOrderFilled)
				if err := _YBNegRisk.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (_YBNegRisk *YBNegRiskFilterer) ParseOrderFilled(log types.Log) (*YBNegRiskOrderFilled, error) {
	event := new(YBNegRiskOrderFilled)
	if err := _YBNegRisk.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the YBNegRisk contract.
type YBNegRiskOrdersMatchedIterator struct {
	Event *YBNegRiskOrdersMatched // Event containing the contract specifics and raw log

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
func (it *YBNegRiskOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskOrdersMatched)
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
		it.Event = new(YBNegRiskOrdersMatched)
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
func (it *YBNegRiskOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskOrdersMatched represents a OrdersMatched event raised by the YBNegRisk contract.
type YBNegRiskOrdersMatched struct {
	TakerOrderHash    [32]byte
	TakerOrderMaker   common.Address
	MakerAssetId      *big.Int
	TakerAssetId      *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_YBNegRisk *YBNegRiskFilterer) FilterOrdersMatched(opts *bind.FilterOpts, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (*YBNegRiskOrdersMatchedIterator, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskOrdersMatchedIterator{contract: _YBNegRisk.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_YBNegRisk *YBNegRiskFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *YBNegRiskOrdersMatched, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (event.Subscription, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskOrdersMatched)
				if err := _YBNegRisk.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_YBNegRisk *YBNegRiskFilterer) ParseOrdersMatched(log types.Log) (*YBNegRiskOrdersMatched, error) {
	event := new(YBNegRiskOrdersMatched)
	if err := _YBNegRisk.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskProxyFactoryUpdatedIterator is returned from FilterProxyFactoryUpdated and is used to iterate over the raw logs and unpacked data for ProxyFactoryUpdated events raised by the YBNegRisk contract.
type YBNegRiskProxyFactoryUpdatedIterator struct {
	Event *YBNegRiskProxyFactoryUpdated // Event containing the contract specifics and raw log

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
func (it *YBNegRiskProxyFactoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskProxyFactoryUpdated)
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
		it.Event = new(YBNegRiskProxyFactoryUpdated)
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
func (it *YBNegRiskProxyFactoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskProxyFactoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskProxyFactoryUpdated represents a ProxyFactoryUpdated event raised by the YBNegRisk contract.
type YBNegRiskProxyFactoryUpdated struct {
	OldProxyFactory common.Address
	NewProxyFactory common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProxyFactoryUpdated is a free log retrieval operation binding the contract event 0x3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a37.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (_YBNegRisk *YBNegRiskFilterer) FilterProxyFactoryUpdated(opts *bind.FilterOpts, oldProxyFactory []common.Address, newProxyFactory []common.Address) (*YBNegRiskProxyFactoryUpdatedIterator, error) {

	var oldProxyFactoryRule []interface{}
	for _, oldProxyFactoryItem := range oldProxyFactory {
		oldProxyFactoryRule = append(oldProxyFactoryRule, oldProxyFactoryItem)
	}
	var newProxyFactoryRule []interface{}
	for _, newProxyFactoryItem := range newProxyFactory {
		newProxyFactoryRule = append(newProxyFactoryRule, newProxyFactoryItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "ProxyFactoryUpdated", oldProxyFactoryRule, newProxyFactoryRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskProxyFactoryUpdatedIterator{contract: _YBNegRisk.contract, event: "ProxyFactoryUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyFactoryUpdated is a free log subscription operation binding the contract event 0x3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a37.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (_YBNegRisk *YBNegRiskFilterer) WatchProxyFactoryUpdated(opts *bind.WatchOpts, sink chan<- *YBNegRiskProxyFactoryUpdated, oldProxyFactory []common.Address, newProxyFactory []common.Address) (event.Subscription, error) {

	var oldProxyFactoryRule []interface{}
	for _, oldProxyFactoryItem := range oldProxyFactory {
		oldProxyFactoryRule = append(oldProxyFactoryRule, oldProxyFactoryItem)
	}
	var newProxyFactoryRule []interface{}
	for _, newProxyFactoryItem := range newProxyFactory {
		newProxyFactoryRule = append(newProxyFactoryRule, newProxyFactoryItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "ProxyFactoryUpdated", oldProxyFactoryRule, newProxyFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskProxyFactoryUpdated)
				if err := _YBNegRisk.contract.UnpackLog(event, "ProxyFactoryUpdated", log); err != nil {
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

// ParseProxyFactoryUpdated is a log parse operation binding the contract event 0x3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a37.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (_YBNegRisk *YBNegRiskFilterer) ParseProxyFactoryUpdated(log types.Log) (*YBNegRiskProxyFactoryUpdated, error) {
	event := new(YBNegRiskProxyFactoryUpdated)
	if err := _YBNegRisk.contract.UnpackLog(event, "ProxyFactoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the YBNegRisk contract.
type YBNegRiskRemovedAdminIterator struct {
	Event *YBNegRiskRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *YBNegRiskRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskRemovedAdmin)
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
		it.Event = new(YBNegRiskRemovedAdmin)
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
func (it *YBNegRiskRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskRemovedAdmin represents a RemovedAdmin event raised by the YBNegRisk contract.
type YBNegRiskRemovedAdmin struct {
	RemovedAdmin common.Address
	Admin        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_YBNegRisk *YBNegRiskFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, removedAdmin []common.Address, admin []common.Address) (*YBNegRiskRemovedAdminIterator, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskRemovedAdminIterator{contract: _YBNegRisk.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_YBNegRisk *YBNegRiskFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *YBNegRiskRemovedAdmin, removedAdmin []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskRemovedAdmin)
				if err := _YBNegRisk.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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
func (_YBNegRisk *YBNegRiskFilterer) ParseRemovedAdmin(log types.Log) (*YBNegRiskRemovedAdmin, error) {
	event := new(YBNegRiskRemovedAdmin)
	if err := _YBNegRisk.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskRemovedOperatorIterator is returned from FilterRemovedOperator and is used to iterate over the raw logs and unpacked data for RemovedOperator events raised by the YBNegRisk contract.
type YBNegRiskRemovedOperatorIterator struct {
	Event *YBNegRiskRemovedOperator // Event containing the contract specifics and raw log

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
func (it *YBNegRiskRemovedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskRemovedOperator)
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
		it.Event = new(YBNegRiskRemovedOperator)
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
func (it *YBNegRiskRemovedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskRemovedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskRemovedOperator represents a RemovedOperator event raised by the YBNegRisk contract.
type YBNegRiskRemovedOperator struct {
	RemovedOperator common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRemovedOperator is a free log retrieval operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_YBNegRisk *YBNegRiskFilterer) FilterRemovedOperator(opts *bind.FilterOpts, removedOperator []common.Address, admin []common.Address) (*YBNegRiskRemovedOperatorIterator, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskRemovedOperatorIterator{contract: _YBNegRisk.contract, event: "RemovedOperator", logs: logs, sub: sub}, nil
}

// WatchRemovedOperator is a free log subscription operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_YBNegRisk *YBNegRiskFilterer) WatchRemovedOperator(opts *bind.WatchOpts, sink chan<- *YBNegRiskRemovedOperator, removedOperator []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskRemovedOperator)
				if err := _YBNegRisk.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
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
func (_YBNegRisk *YBNegRiskFilterer) ParseRemovedOperator(log types.Log) (*YBNegRiskRemovedOperator, error) {
	event := new(YBNegRiskRemovedOperator)
	if err := _YBNegRisk.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskSafeFactoryUpdatedIterator is returned from FilterSafeFactoryUpdated and is used to iterate over the raw logs and unpacked data for SafeFactoryUpdated events raised by the YBNegRisk contract.
type YBNegRiskSafeFactoryUpdatedIterator struct {
	Event *YBNegRiskSafeFactoryUpdated // Event containing the contract specifics and raw log

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
func (it *YBNegRiskSafeFactoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskSafeFactoryUpdated)
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
		it.Event = new(YBNegRiskSafeFactoryUpdated)
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
func (it *YBNegRiskSafeFactoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskSafeFactoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskSafeFactoryUpdated represents a SafeFactoryUpdated event raised by the YBNegRisk contract.
type YBNegRiskSafeFactoryUpdated struct {
	OldSafeFactory common.Address
	NewSafeFactory common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSafeFactoryUpdated is a free log retrieval operation binding the contract event 0x9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (_YBNegRisk *YBNegRiskFilterer) FilterSafeFactoryUpdated(opts *bind.FilterOpts, oldSafeFactory []common.Address, newSafeFactory []common.Address) (*YBNegRiskSafeFactoryUpdatedIterator, error) {

	var oldSafeFactoryRule []interface{}
	for _, oldSafeFactoryItem := range oldSafeFactory {
		oldSafeFactoryRule = append(oldSafeFactoryRule, oldSafeFactoryItem)
	}
	var newSafeFactoryRule []interface{}
	for _, newSafeFactoryItem := range newSafeFactory {
		newSafeFactoryRule = append(newSafeFactoryRule, newSafeFactoryItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "SafeFactoryUpdated", oldSafeFactoryRule, newSafeFactoryRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskSafeFactoryUpdatedIterator{contract: _YBNegRisk.contract, event: "SafeFactoryUpdated", logs: logs, sub: sub}, nil
}

// WatchSafeFactoryUpdated is a free log subscription operation binding the contract event 0x9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (_YBNegRisk *YBNegRiskFilterer) WatchSafeFactoryUpdated(opts *bind.WatchOpts, sink chan<- *YBNegRiskSafeFactoryUpdated, oldSafeFactory []common.Address, newSafeFactory []common.Address) (event.Subscription, error) {

	var oldSafeFactoryRule []interface{}
	for _, oldSafeFactoryItem := range oldSafeFactory {
		oldSafeFactoryRule = append(oldSafeFactoryRule, oldSafeFactoryItem)
	}
	var newSafeFactoryRule []interface{}
	for _, newSafeFactoryItem := range newSafeFactory {
		newSafeFactoryRule = append(newSafeFactoryRule, newSafeFactoryItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "SafeFactoryUpdated", oldSafeFactoryRule, newSafeFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskSafeFactoryUpdated)
				if err := _YBNegRisk.contract.UnpackLog(event, "SafeFactoryUpdated", log); err != nil {
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

// ParseSafeFactoryUpdated is a log parse operation binding the contract event 0x9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (_YBNegRisk *YBNegRiskFilterer) ParseSafeFactoryUpdated(log types.Log) (*YBNegRiskSafeFactoryUpdated, error) {
	event := new(YBNegRiskSafeFactoryUpdated)
	if err := _YBNegRisk.contract.UnpackLog(event, "SafeFactoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the YBNegRisk contract.
type YBNegRiskTokenRegisteredIterator struct {
	Event *YBNegRiskTokenRegistered // Event containing the contract specifics and raw log

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
func (it *YBNegRiskTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskTokenRegistered)
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
		it.Event = new(YBNegRiskTokenRegistered)
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
func (it *YBNegRiskTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskTokenRegistered represents a TokenRegistered event raised by the YBNegRisk contract.
type YBNegRiskTokenRegistered struct {
	Token0      *big.Int
	Token1      *big.Int
	ConditionId [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0xbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (_YBNegRisk *YBNegRiskFilterer) FilterTokenRegistered(opts *bind.FilterOpts, token0 []*big.Int, token1 []*big.Int, conditionId [][32]byte) (*YBNegRiskTokenRegisteredIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "TokenRegistered", token0Rule, token1Rule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskTokenRegisteredIterator{contract: _YBNegRisk.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (_YBNegRisk *YBNegRiskFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *YBNegRiskTokenRegistered, token0 []*big.Int, token1 []*big.Int, conditionId [][32]byte) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "TokenRegistered", token0Rule, token1Rule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskTokenRegistered)
				if err := _YBNegRisk.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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

// ParseTokenRegistered is a log parse operation binding the contract event 0xbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (_YBNegRisk *YBNegRiskFilterer) ParseTokenRegistered(log types.Log) (*YBNegRiskTokenRegistered, error) {
	event := new(YBNegRiskTokenRegistered)
	if err := _YBNegRisk.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskTradingPausedIterator is returned from FilterTradingPaused and is used to iterate over the raw logs and unpacked data for TradingPaused events raised by the YBNegRisk contract.
type YBNegRiskTradingPausedIterator struct {
	Event *YBNegRiskTradingPaused // Event containing the contract specifics and raw log

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
func (it *YBNegRiskTradingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskTradingPaused)
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
		it.Event = new(YBNegRiskTradingPaused)
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
func (it *YBNegRiskTradingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskTradingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskTradingPaused represents a TradingPaused event raised by the YBNegRisk contract.
type YBNegRiskTradingPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingPaused is a free log retrieval operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_YBNegRisk *YBNegRiskFilterer) FilterTradingPaused(opts *bind.FilterOpts, pauser []common.Address) (*YBNegRiskTradingPausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskTradingPausedIterator{contract: _YBNegRisk.contract, event: "TradingPaused", logs: logs, sub: sub}, nil
}

// WatchTradingPaused is a free log subscription operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_YBNegRisk *YBNegRiskFilterer) WatchTradingPaused(opts *bind.WatchOpts, sink chan<- *YBNegRiskTradingPaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskTradingPaused)
				if err := _YBNegRisk.contract.UnpackLog(event, "TradingPaused", log); err != nil {
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

// ParseTradingPaused is a log parse operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_YBNegRisk *YBNegRiskFilterer) ParseTradingPaused(log types.Log) (*YBNegRiskTradingPaused, error) {
	event := new(YBNegRiskTradingPaused)
	if err := _YBNegRisk.contract.UnpackLog(event, "TradingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YBNegRiskTradingUnpausedIterator is returned from FilterTradingUnpaused and is used to iterate over the raw logs and unpacked data for TradingUnpaused events raised by the YBNegRisk contract.
type YBNegRiskTradingUnpausedIterator struct {
	Event *YBNegRiskTradingUnpaused // Event containing the contract specifics and raw log

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
func (it *YBNegRiskTradingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YBNegRiskTradingUnpaused)
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
		it.Event = new(YBNegRiskTradingUnpaused)
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
func (it *YBNegRiskTradingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YBNegRiskTradingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YBNegRiskTradingUnpaused represents a TradingUnpaused event raised by the YBNegRisk contract.
type YBNegRiskTradingUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingUnpaused is a free log retrieval operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_YBNegRisk *YBNegRiskFilterer) FilterTradingUnpaused(opts *bind.FilterOpts, pauser []common.Address) (*YBNegRiskTradingUnpausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _YBNegRisk.contract.FilterLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &YBNegRiskTradingUnpausedIterator{contract: _YBNegRisk.contract, event: "TradingUnpaused", logs: logs, sub: sub}, nil
}

// WatchTradingUnpaused is a free log subscription operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_YBNegRisk *YBNegRiskFilterer) WatchTradingUnpaused(opts *bind.WatchOpts, sink chan<- *YBNegRiskTradingUnpaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _YBNegRisk.contract.WatchLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YBNegRiskTradingUnpaused)
				if err := _YBNegRisk.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
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

// ParseTradingUnpaused is a log parse operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_YBNegRisk *YBNegRiskFilterer) ParseTradingUnpaused(log types.Log) (*YBNegRiskTradingUnpaused, error) {
	event := new(YBNegRiskTradingUnpaused)
	if err := _YBNegRisk.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
