package predictcontracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ContractConfig holds all contract addresses for Predict platform
//
// IMPORTANT: ConditionalTokens Access Control
// The ConditionalTokens contract has role-based access control for split/merge operations.
// When the access control switch is enabled, only addresses with SPLITTER_ROLE or MERGER_ROLE
// can perform splitPosition/mergePositions operations. This is a security feature but may
// cause transaction failures if your address doesn't have the required role.
// Check contract's accessControlEnabled() and hasRole() before calling split/merge functions.
type ContractConfig struct {
	// Common
	Collateral common.Address // USDT on BNB

	// Yield Bearing Prediction Market
	// YieldBearingConditionalTokens is fully compatible with ConditionalTokens, but automatically
	// starts earning yield when splitPosition is called, solving the idle capital problem in
	// traditional prediction markets. The collateral is deposited into Venus protocol to earn yield.
	// NOTE: YieldBearingConditionalTokens and YieldBearingNegRiskConditionalTokens have
	// split/merge role-based access control (SPLIT_POSITION_ROLE, MERGE_POSITIONS_ROLE).
	// Check isSplitPositionGated()/isMergePositionsGated() and ensure your address has the required role.
	YieldBearingConditionalTokens common.Address
	YieldBearingCtfAdapter               common.Address
	YieldBearingExchange                 common.Address
	YieldBearingNegRiskConditionalTokens common.Address
	YieldBearingNegRiskAdapter           common.Address
	YieldBearingWrappedCollateral        common.Address
	YieldBearingNegRiskExchange          common.Address
	YieldBearingNegRiskOperator          common.Address
	YieldBearingNegRiskCtfAdapter        common.Address
	YieldBearingFeeModule                common.Address
	YieldBearingNegRiskFeeModule         common.Address
	YieldBearingFeesHandler              common.Address
	YieldBearingRegisterTokenHelper      common.Address

	// Non Yield Bearing Prediction Market
	// NOTE: ConditionalTokens and NegRiskConditionalTokens have split/merge role-based
	// access control. Ensure your address has the required role when access control is enabled.
	ConditionalTokens common.Address
	CtfAdapter               common.Address
	Exchange                 common.Address
	NegRiskConditionalTokens common.Address
	NegRiskAdapter           common.Address
	WrappedCollateral        common.Address
	NegRiskExchange          common.Address
	NegRiskOperator          common.Address
	NegRiskCtfAdapter        common.Address
	FeeModule                common.Address
	NegRiskFeeModule         common.Address
	FeesHandler              common.Address
	RegisterTokenHelper      common.Address

	// Oracle & Utils
	OptimisticOracle        common.Address
	Vault                   common.Address
	ZeroDevWithdrawalHelper common.Address
	RewardDistributor       common.Address
}

// BNB Mainnet (Chain ID: 56)
var BNB_CONTRACTS = &ContractConfig{
	// Common
	Collateral: common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), // USDT on BNB

	// Yield Bearing Prediction Market
	YieldBearingConditionalTokens:        common.HexToAddress("0x9400F8Ad57e9e0F352345935d6D3175975eb1d9F"),
	YieldBearingCtfAdapter:               common.HexToAddress("0x947cc06D38d3cB0a2BB5AdFB668b99B4FF53d7B4"),
	YieldBearingExchange:                 common.HexToAddress("0x6bEb5a40C032AFc305961162d8204CDA16DECFa5"),
	YieldBearingNegRiskConditionalTokens: common.HexToAddress("0xF64b0b318AAf83BD9071110af24D24445719A07F"),
	YieldBearingNegRiskAdapter:           common.HexToAddress("0x41dCe1A4B8FB5e6327701750aF6231B7CD0B2A40"),
	YieldBearingWrappedCollateral:        common.HexToAddress("0xCfb9beF5F7B748aC72311F057f3a888BC73334D9"),
	YieldBearingNegRiskExchange:          common.HexToAddress("0x8A289d458f5a134bA40015085A8F50Ffb681B41d"),
	YieldBearingNegRiskOperator:          common.HexToAddress("0xBB7250101e0e3611D7e136fFE73Bc24b98E3e175"),
	YieldBearingNegRiskCtfAdapter:        common.HexToAddress("0x26B366Ab634C43BdA6D784fDCe34F24A37DF8172"),
	YieldBearingFeeModule:                common.HexToAddress("0xFbC2259aBB3F01c019ECE1d0200Ee673BB7BA34F"),
	YieldBearingNegRiskFeeModule:         common.HexToAddress("0xD172f3FBabe763Ee8E52D8b32421574236dA6057"),
	YieldBearingFeesHandler:              common.HexToAddress("0xb4D9F13738a50E88E0ade2ecCc89254EF1645f6E"),
	YieldBearingRegisterTokenHelper:      common.HexToAddress("0xa48C26abd9024a5CC5a869bBd97A6a3D6B9C2089"),

	// Non Yield Bearing Prediction Market
	ConditionalTokens:        common.HexToAddress("0x22DA1810B194ca018378464a58f6Ac2B10C9d244"),
	CtfAdapter:               common.HexToAddress("0x242E1Ba24f6fC524bfb410062Ca5689A9622613d"),
	Exchange:                 common.HexToAddress("0x8BC070BEdAB741406F4B1Eb65A72bee27894B689"),
	NegRiskConditionalTokens: common.HexToAddress("0x22DA1810B194ca018378464a58f6Ac2B10C9d244"),
	NegRiskAdapter:           common.HexToAddress("0xc3Cf7c252f65E0d8D88537dF96569AE94a7F1A6E"),
	WrappedCollateral:        common.HexToAddress("0x66239b70133773A72A0D589E5564E88a50Cd39e7"),
	NegRiskExchange:          common.HexToAddress("0x365fb81bd4A24D6303cd2F19c349dE6894D8d58A"),
	NegRiskOperator:          common.HexToAddress("0x56020F5024641d577Cb54032aF70a23a986ECfFD"),
	NegRiskCtfAdapter:        common.HexToAddress("0xf61198a64C2e4CAD8CCAf218f3f2ECeFb017902F"),
	FeeModule:                common.HexToAddress("0xF1f8F5C641F20C48526269EF7DFF19172Efa9783"),
	NegRiskFeeModule:         common.HexToAddress("0xF2311C668aAA8dEc48D5da577d3018eb94b3132F"),
	FeesHandler:              common.HexToAddress("0xD63206243192f1AF3d6fC4442db4e3cf25E64030"),
	RegisterTokenHelper:      common.HexToAddress("0x89f92C3c27F18080AF1361024C6a892144fD8e5E"),

	// Oracle & Utils
	OptimisticOracle:        common.HexToAddress("0x76F42e5520E62AD88f8fE583cBb4BfF27eeC2531"),
	Vault:                   common.HexToAddress("0x09F683d8a144c4ac296D770F839098c3377410c5"),
	ZeroDevWithdrawalHelper: common.HexToAddress("0xf4aa30b537882eca7e69defb68d6f631cda77b00"),
	RewardDistributor:       common.HexToAddress("0x14e3cB02F48818a8FeF6BC257059767cA9d436Ae"),
}

// BNB Testnet (Chain ID: 97)
var BNB_TESTNET_CONTRACTS = &ContractConfig{
	// Common
	Collateral: common.HexToAddress("0x337610d27c682E347C9cD60BD4b3b107C9d34dDd"), // USDT on BNB Testnet

	// Yield Bearing Prediction Market
	YieldBearingConditionalTokens:        common.HexToAddress("0x38BF1cbD66d174bb5F3037d7068E708861D68D7f"),
	YieldBearingCtfAdapter:               common.HexToAddress("0xF2DBA0C0d306bD8480653Ca9faa0cd2Ef32ca121"),
	YieldBearingExchange:                 common.HexToAddress("0x8a6B4Fa700A1e310b106E7a48bAFa29111f66e89"),
	YieldBearingNegRiskConditionalTokens: common.HexToAddress("0x26e865CbaAe99b62fbF9D18B55c25B5E079A93D5"),
	YieldBearingNegRiskAdapter:           common.HexToAddress("0xb74aea04bdeBE912Aa425bC9173F9668e6f11F99"),
	YieldBearingWrappedCollateral:        common.HexToAddress("0xd5647Df4f2e884bb21167C2aBB99971693094599"),
	YieldBearingNegRiskExchange:          common.HexToAddress("0x95D5113bc50eD201e319101bbca3e0E250662fCC"),
	YieldBearingNegRiskOperator:          common.HexToAddress("0x7Ab1B03319cD292A787090337E9F5db839E662F4"),
	YieldBearingNegRiskCtfAdapter:        common.HexToAddress("0x42bA93dDF195D93dEAABA06954bBbad5C78B7128"),
	YieldBearingFeeModule:                common.HexToAddress("0x66569E776f247C0e270Bb71259755DC31869c3c7"),
	YieldBearingNegRiskFeeModule:         common.HexToAddress("0xD803Ab73B8598305D7256BF3ba7143a2b5C39Cdf"),
	YieldBearingFeesHandler:              common.HexToAddress("0x749c677A7DA13f1D4bdF5853071FA2f1FB1c2587"),
	YieldBearingRegisterTokenHelper:      common.HexToAddress("0x20526fA3Ef788A0898b9DFcf390c16B1391d84e7"),

	// Non Yield Bearing Prediction Market
	ConditionalTokens:        common.HexToAddress("0x2827AAef52D71910E8FBad2FfeBC1B6C2DA37743"),
	CtfAdapter:               common.HexToAddress("0x8403922ad0c5d39B4148e17C86095cAf10B37E50"),
	Exchange:                 common.HexToAddress("0x2A6413639BD3d73a20ed8C95F634Ce198ABbd2d7"),
	NegRiskConditionalTokens: common.HexToAddress("0x2827AAef52D71910E8FBad2FfeBC1B6C2DA37743"),
	NegRiskAdapter:           common.HexToAddress("0x285c1B939380B130D7EBd09467b93faD4BA623Ed"),
	WrappedCollateral:        common.HexToAddress("0x61E0b85207393A3804c79Ef35A155A676238C469"),
	NegRiskExchange:          common.HexToAddress("0xd690b2bd441bE36431F6F6639D7Ad351e7B29680"),
	NegRiskOperator:          common.HexToAddress("0xD02331CdbEf06F54aeEddA76A2E61Df2cef1C8e8"),
	NegRiskCtfAdapter:        common.HexToAddress("0x52DA245ac170155391e7607c67b77D549005002d"),
	FeeModule:                common.HexToAddress("0xD479C5c6A3B98C6db8503cb6530519Bb19249CC2"),
	NegRiskFeeModule:         common.HexToAddress("0x7F5ED9CD275DD6b88C5302262DbA9FD8c407Dd87"),
	FeesHandler:              common.HexToAddress("0xD63D7F7AeAD49C1711b5Cd00Ed4CA26e9a83ff6C"),
	RegisterTokenHelper:      common.HexToAddress("0x450bfd61d19b0a02dD87056F04914a04749d5bc4"),

	// Oracle & Utils
	OptimisticOracle: common.HexToAddress("0xb5ae73B1dB6659b741d14DEDCeF6B73c88A8fe4c"),
	Vault:            common.HexToAddress("0x415bdd0F4e5eE9A50B2394ff8B6b20319e77255d"),
	// ZeroDevWithdrawalHelper: N/A on testnet
	// RewardDistributor: N/A on testnet
}

func GetContractConfig(chainID *big.Int) *ContractConfig {
	switch chainID.Int64() {
	case 56:
		return BNB_CONTRACTS
	case 97:
		return BNB_TESTNET_CONTRACTS
	default:
		panic("Invalid network: only BNB (56) and BNB Testnet (97) are supported")
	}
}
