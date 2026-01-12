package predictcontracts

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ivanzzeth/ethclient"
	conditional_tokens "github.com/ivanzzeth/predict-go-contracts/contracts/conditional-tokens"
	"github.com/ivanzzeth/predict-go-contracts/contracts/erc20"
	"github.com/ivanzzeth/predict-go-contracts/contracts/exchange"
	negrisk "github.com/ivanzzeth/predict-go-contracts/contracts/neg-risk"
	negriskadapter "github.com/ivanzzeth/predict-go-contracts/contracts/neg-risk-adapter"
	"github.com/ivanzzeth/predict-go-contracts/sender"
	"github.com/ivanzzeth/predict-go-contracts/signer"
)

// BalanceAllowanceInfo holds balance and allowance information
type BalanceAllowanceInfo struct {
	Balance                    *big.Int
	AllowanceExchange          *big.Int
	AllowanceNegRiskAdapter    *big.Int
	AllowanceNegRiskExchange   *big.Int
	CTFApprovedExchange        bool
	CTFApprovedNegRiskAdapter  bool
	CTFApprovedNegRiskExchange bool
}

type ContractInterface struct {
	chainID          *big.Int
	client           ethclient.EthClientInterface
	contractConfig   *ContractConfig
	eoaTradingSigner signer.EOATradingSigner
	txSender         sender.TransactionSender

	collateralContract        *erc20.Erc20
	conditionalTokensContract *conditional_tokens.ConditionalTokens
	exchangeContract          *exchange.Exchange
	negRiskAdapterContract    *negriskadapter.NegRiskAdapter
	negRiskContract           *negrisk.NegRisk
}

type ContractInterfaceConfig struct {
	TxSender         sender.TransactionSender
	EOATradingSigner signer.EOATradingSigner
	ContractConfig   *ContractConfig
}

type ContractInterfaceOption func(c *ContractInterfaceConfig)

func WithContractConfig(cc *ContractConfig) ContractInterfaceOption {
	return func(c *ContractInterfaceConfig) {
		c.ContractConfig = cc
	}
}

func WithEOASigner(eoaSigner signer.EOATradingSigner) ContractInterfaceOption {
	return func(c *ContractInterfaceConfig) {
		c.EOATradingSigner = eoaSigner
	}
}

func WithTransactionSender(txSender sender.TransactionSender) ContractInterfaceOption {
	return func(c *ContractInterfaceConfig) {
		c.TxSender = txSender
	}
}

func NewContractInterface(
	client ethclient.EthClientInterface,
	options ...ContractInterfaceOption,
) (*ContractInterface, error) {
	defaultOptions := &ContractInterfaceConfig{
		ContractConfig: BNB_CONTRACTS,
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %w", err)
	}

	for _, opFn := range options {
		opFn(defaultOptions)
	}

	if defaultOptions.TxSender == nil && defaultOptions.EOATradingSigner != nil {
		txSender, err := signer.GetTransactionSenderBySigner(chainID, client, defaultOptions.EOATradingSigner)
		if err != nil {
			return nil, fmt.Errorf("failed to get txSender for EOA signer: %w", err)
		}
		defaultOptions.TxSender = txSender
	}

	// Initialize Collateral (USDT) contract
	collateralContract, err := erc20.NewErc20(defaultOptions.ContractConfig.Collateral, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup Collateral: %w", err)
	}

	// Initialize Conditional Tokens Framework contract
	ctfContract, err := conditional_tokens.NewConditionalTokens(defaultOptions.ContractConfig.ConditionalTokens, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup ConditionalTokens: %w", err)
	}

	// Initialize Exchange contract
	exchangeContract, err := exchange.NewExchange(defaultOptions.ContractConfig.Exchange, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup Exchange: %w", err)
	}

	// Initialize NegRisk Adapter contract
	negRiskAdapterContract, err := negriskadapter.NewNegRiskAdapter(defaultOptions.ContractConfig.NegRiskAdapter, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup NegRiskAdapter: %w", err)
	}

	// Initialize NegRisk Exchange contract
	negRiskContract, err := negrisk.NewNegRisk(defaultOptions.ContractConfig.NegRiskExchange, client)
	if err != nil {
		return nil, fmt.Errorf("failed to setup NegRisk: %w", err)
	}

	return &ContractInterface{
		chainID:          chainID,
		client:           client,
		contractConfig:   defaultOptions.ContractConfig,
		txSender:         defaultOptions.TxSender,
		eoaTradingSigner: defaultOptions.EOATradingSigner,

		collateralContract:        collateralContract,
		conditionalTokensContract: ctfContract,
		exchangeContract:          exchangeContract,
		negRiskAdapterContract:    negRiskAdapterContract,
		negRiskContract:           negRiskContract,
	}, nil
}

// GetCollateral returns the Collateral contract instance
func (b *ContractInterface) GetCollateral() *erc20.Erc20 {
	return b.collateralContract
}

// GetConditionalTokens returns the Conditional Tokens Framework contract instance
func (b *ContractInterface) GetConditionalTokens() *conditional_tokens.ConditionalTokens {
	return b.conditionalTokensContract
}

// GetExchange returns the Exchange contract instance
func (b *ContractInterface) GetExchange() *exchange.Exchange {
	return b.exchangeContract
}

// GetNegRiskAdapter returns the NegRisk Adapter contract instance
func (b *ContractInterface) GetNegRiskAdapter() *negriskadapter.NegRiskAdapter {
	return b.negRiskAdapterContract
}

// GetNegRisk returns the NegRisk Exchange contract instance
func (b *ContractInterface) GetNegRisk() *negrisk.NegRisk {
	return b.negRiskContract
}

// GetConfig returns the contract configuration
func (b *ContractInterface) GetConfig() *ContractConfig {
	return b.contractConfig
}

// GetClient returns the Ethereum client
func (b *ContractInterface) GetClient() ethclient.EthClientInterface {
	return b.client
}

// GetTxSender returns the transaction sender
func (b *ContractInterface) GetTxSender() sender.TransactionSender {
	return b.txSender
}

// GetEOATradingSigner returns the EOA trading signer
func (b *ContractInterface) GetEOATradingSigner() signer.EOATradingSigner {
	return b.eoaTradingSigner
}

// CheckBalanceAndAllowance checks the collateral balance and all allowances for the given address
func (b *ContractInterface) CheckBalanceAndAllowance(ctx context.Context, address common.Address) (*BalanceAllowanceInfo, error) {
	return b.CheckBalanceAndAllowanceAtBlock(ctx, address, nil)
}

// CheckBalanceAndAllowanceAtBlock checks the collateral balance and all allowances for the given address at a specific block
func (b *ContractInterface) CheckBalanceAndAllowanceAtBlock(ctx context.Context, address common.Address, blockNumber *big.Int) (*BalanceAllowanceInfo, error) {
	info := &BalanceAllowanceInfo{}
	callOpts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}

	// Check collateral balance
	balance, err := b.collateralContract.BalanceOf(callOpts, address)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral balance: %w", err)
	}
	info.Balance = balance

	// Check collateral allowances
	allowanceExchange, err := b.collateralContract.Allowance(callOpts, address, b.contractConfig.Exchange)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral allowance for Exchange: %w", err)
	}
	info.AllowanceExchange = allowanceExchange

	allowanceNegRiskAdapter, err := b.collateralContract.Allowance(callOpts, address, b.contractConfig.NegRiskAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral allowance for NegRiskAdapter: %w", err)
	}
	info.AllowanceNegRiskAdapter = allowanceNegRiskAdapter

	allowanceNegRiskExchange, err := b.collateralContract.Allowance(callOpts, address, b.contractConfig.NegRiskExchange)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral allowance for NegRiskExchange: %w", err)
	}
	info.AllowanceNegRiskExchange = allowanceNegRiskExchange

	// Check CTF approvals
	ctfApprovedExchange, err := b.conditionalTokensContract.IsApprovedForAll(callOpts, address, b.contractConfig.Exchange)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for Exchange: %w", err)
	}
	info.CTFApprovedExchange = ctfApprovedExchange

	ctfApprovedNegRiskAdapter, err := b.conditionalTokensContract.IsApprovedForAll(callOpts, address, b.contractConfig.NegRiskAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for NegRiskAdapter: %w", err)
	}
	info.CTFApprovedNegRiskAdapter = ctfApprovedNegRiskAdapter

	ctfApprovedNegRiskExchange, err := b.conditionalTokensContract.IsApprovedForAll(callOpts, address, b.contractConfig.NegRiskExchange)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for NegRiskExchange: %w", err)
	}
	info.CTFApprovedNegRiskExchange = ctfApprovedNegRiskExchange

	return info, nil
}

// PrintBalanceAndAllowance prints the balance and allowance information in a user-friendly format
func (b *ContractInterface) PrintBalanceAndAllowance(ctx context.Context, address common.Address) error {
	info, err := b.CheckBalanceAndAllowance(ctx, address)
	if err != nil {
		return err
	}

	fmt.Println("=== Balance and Allowance Status ===")

	// Print Collateral Balance
	fmt.Printf("Collateral Balance: %s\n\n", formatUSDC(info.Balance))

	// Print Allowances
	fmt.Println("Allowances:")
	fmt.Printf("  Collateral → Exchange: %s\n", checkmark(info.AllowanceExchange.Cmp(big.NewInt(0)) > 0))
	fmt.Printf("  Collateral → NegRiskAdapter: %s\n", checkmark(info.AllowanceNegRiskAdapter.Cmp(big.NewInt(0)) > 0))
	fmt.Printf("  Collateral → NegRiskExchange: %s\n", checkmark(info.AllowanceNegRiskExchange.Cmp(big.NewInt(0)) > 0))
	fmt.Printf("  CTF → Exchange: %s\n", checkmark(info.CTFApprovedExchange))
	fmt.Printf("  CTF → NegRiskAdapter: %s\n", checkmark(info.CTFApprovedNegRiskAdapter))
	fmt.Printf("  CTF → NegRiskExchange: %s\n\n", checkmark(info.CTFApprovedNegRiskExchange))

	return nil
}

// EnableTrading enables trading for the EOA on both YB and NYB contracts
func (b *ContractInterface) EnableTrading(ctx context.Context) ([]common.Hash, error) {
	if b.eoaTradingSigner == nil {
		return nil, fmt.Errorf("EOA trading signer not set")
	}

	var allTxHashes []common.Hash

	// Enable trading for Non Yield Bearing contracts
	fmt.Println("Enabling trading for Non Yield Bearing (NYB) contracts...")
	nybTxHashes, err := b.enableTradingForContracts(ctx,
		b.contractConfig.ConditionalTokens,
		b.contractConfig.Exchange,
		b.contractConfig.NegRiskAdapter,
		b.contractConfig.NegRiskExchange,
	)
	if err != nil {
		return allTxHashes, fmt.Errorf("failed to enable trading for NYB contracts: %w", err)
	}
	allTxHashes = append(allTxHashes, nybTxHashes...)

	// Enable trading for Yield Bearing contracts
	fmt.Println("Enabling trading for Yield Bearing (YB) contracts...")
	ybTxHashes, err := b.enableTradingForContracts(ctx,
		b.contractConfig.YieldBearingConditionalTokens,
		b.contractConfig.YieldBearingExchange,
		b.contractConfig.YieldBearingNegRiskAdapter,
		b.contractConfig.YieldBearingNegRiskExchange,
	)
	if err != nil {
		return allTxHashes, fmt.Errorf("failed to enable trading for YB contracts: %w", err)
	}
	allTxHashes = append(allTxHashes, ybTxHashes...)

	return allTxHashes, nil
}

// Redeem redeems conditional tokens for a resolved market
// ctfAddress: the ConditionalTokens contract address (use YieldBearingConditionalTokens or ConditionalTokens based on market type)
func (b *ContractInterface) Redeem(ctx context.Context, ctfAddress common.Address, conditionId [32]byte) (common.Hash, error) {
	// For binary markets, indexSets is always [1, 2] representing Yes/No outcomes
	indexSets := []*big.Int{big.NewInt(1), big.NewInt(2)}
	return b.RedeemPositionsForEOA(ctx, ctfAddress, conditionId, indexSets)
}

// RedeemNegRisk redeems NegRisk market positions
// negRiskAdapterAddress: the NegRiskAdapter contract address (use YieldBearingNegRiskAdapter or NegRiskAdapter based on market type)
func (b *ContractInterface) RedeemNegRisk(ctx context.Context, negRiskAdapterAddress common.Address, conditionId [32]byte, amounts []*big.Int) (common.Hash, error) {
	return b.RedeemPositionsNegRiskForEOA(ctx, negRiskAdapterAddress, conditionId, amounts)
}

// Split splits collateral into conditional tokens
// ctfAddress: the ConditionalTokens contract address (use YieldBearingConditionalTokens or ConditionalTokens based on market type)
func (b *ContractInterface) Split(ctx context.Context, ctfAddress common.Address, conditionId [32]byte, amount *big.Int) (common.Hash, error) {
	// For binary markets, partition is always [1, 2] representing Yes/No outcomes
	partition := []*big.Int{big.NewInt(1), big.NewInt(2)}
	return b.SplitPositionForEOA(ctx, ctfAddress, conditionId, partition, amount)
}

// Merge merges conditional tokens back into collateral
// ctfAddress: the ConditionalTokens contract address (use YieldBearingConditionalTokens or ConditionalTokens based on market type)
func (b *ContractInterface) Merge(ctx context.Context, ctfAddress common.Address, conditionId [32]byte, amount *big.Int) (common.Hash, error) {
	// For binary markets, partition is always [1, 2] representing Yes/No outcomes
	partition := []*big.Int{big.NewInt(1), big.NewInt(2)}
	return b.MergePositionsForEOA(ctx, ctfAddress, conditionId, partition, amount)
}

// SplitNegRisk splits collateral into NegRisk conditional tokens
// negRiskAdapterAddress: the NegRiskAdapter contract address (use YieldBearingNegRiskAdapter or NegRiskAdapter based on market type)
func (b *ContractInterface) SplitNegRisk(ctx context.Context, negRiskAdapterAddress common.Address, conditionId [32]byte, amount *big.Int) (common.Hash, error) {
	return b.SplitPositionNegRiskForEOA(ctx, negRiskAdapterAddress, conditionId, amount)
}

// MergeNegRisk merges NegRisk conditional tokens back into collateral
// negRiskAdapterAddress: the NegRiskAdapter contract address (use YieldBearingNegRiskAdapter or NegRiskAdapter based on market type)
func (b *ContractInterface) MergeNegRisk(ctx context.Context, negRiskAdapterAddress common.Address, conditionId [32]byte, amount *big.Int) (common.Hash, error) {
	return b.MergePositionsNegRiskForEOA(ctx, negRiskAdapterAddress, conditionId, amount)
}

// enableTradingForContracts is an internal helper that enables trading for a specific set of contracts
// ctfAddress: ConditionalTokens contract (YB or NYB)
// exchangeAddress: Exchange contract
// negRiskAdapterAddress: NegRiskAdapter contract
// negRiskExchangeAddress: NegRiskExchange contract
func (b *ContractInterface) enableTradingForContracts(
	ctx context.Context,
	ctfAddress common.Address,
	exchangeAddress common.Address,
	negRiskAdapterAddress common.Address,
	negRiskExchangeAddress common.Address,
) ([]common.Hash, error) {
	var txHashes []common.Hash

	eoaAddr := b.eoaTradingSigner.GetAddress()
	callOpts := &bind.CallOpts{Context: ctx}

	// Maximum allowance for ERC20 approvals
	maxAllowance := new(big.Int)
	maxAllowance.SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)

	// Parse ERC20 ABI for approve function
	erc20ABI, err := abi.JSON(strings.NewReader(erc20.Erc20MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ERC20 ABI: %w", err)
	}

	// Parse ConditionalTokens ABI for setApprovalForAll and isApprovedForAll
	ctfABI, err := abi.JSON(strings.NewReader(conditional_tokens.ConditionalTokensMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}

	// Check and approve collateral for Exchange
	allowanceExchange, err := b.collateralContract.Allowance(callOpts, eoaAddr, exchangeAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral allowance for Exchange: %w", err)
	}
	if allowanceExchange.Cmp(big.NewInt(0)) == 0 {
		approveData, err := erc20ABI.Pack("approve", exchangeAddress, maxAllowance)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve data for Exchange: %w", err)
		}
		txHash, err := b.txSender.SendEthereumTransaction(b.contractConfig.Collateral, approveData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send Collateral → Exchange approval transaction: %w", err)
		}
		fmt.Printf("  Collateral → Exchange: %s\n", txHash.Hex())
		txHashes = append(txHashes, txHash)
	} else {
		fmt.Printf("  Collateral → Exchange: already approved\n")
	}

	// Check and approve collateral for NegRiskAdapter
	allowanceNegRiskAdapter, err := b.collateralContract.Allowance(callOpts, eoaAddr, negRiskAdapterAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral allowance for NegRiskAdapter: %w", err)
	}
	if allowanceNegRiskAdapter.Cmp(big.NewInt(0)) == 0 {
		approveData, err := erc20ABI.Pack("approve", negRiskAdapterAddress, maxAllowance)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve data for NegRiskAdapter: %w", err)
		}
		txHash, err := b.txSender.SendEthereumTransaction(b.contractConfig.Collateral, approveData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send Collateral → NegRiskAdapter approval transaction: %w", err)
		}
		fmt.Printf("  Collateral → NegRiskAdapter: %s\n", txHash.Hex())
		txHashes = append(txHashes, txHash)
	} else {
		fmt.Printf("  Collateral → NegRiskAdapter: already approved\n")
	}

	// Check and approve collateral for NegRiskExchange
	allowanceNegRiskExchange, err := b.collateralContract.Allowance(callOpts, eoaAddr, negRiskExchangeAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get collateral allowance for NegRiskExchange: %w", err)
	}
	if allowanceNegRiskExchange.Cmp(big.NewInt(0)) == 0 {
		approveData, err := erc20ABI.Pack("approve", negRiskExchangeAddress, maxAllowance)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve data for NegRiskExchange: %w", err)
		}
		txHash, err := b.txSender.SendEthereumTransaction(b.contractConfig.Collateral, approveData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send Collateral → NegRiskExchange approval transaction: %w", err)
		}
		fmt.Printf("  Collateral → NegRiskExchange: %s\n", txHash.Hex())
		txHashes = append(txHashes, txHash)
	} else {
		fmt.Printf("  Collateral → NegRiskExchange: already approved\n")
	}

	// Create a temporary CTF contract instance to check approvals
	ctfContract, err := conditional_tokens.NewConditionalTokens(ctfAddress, b.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create CTF contract instance: %w", err)
	}

	// Check and approve CTF for Exchange
	ctfApprovedExchange, err := ctfContract.IsApprovedForAll(callOpts, eoaAddr, exchangeAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for Exchange: %w", err)
	}
	if !ctfApprovedExchange {
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", exchangeAddress, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack setApprovalForAll data for Exchange: %w", err)
		}
		txHash, err := b.txSender.SendEthereumTransaction(ctfAddress, setApprovalData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send CTF → Exchange approval transaction: %w", err)
		}
		fmt.Printf("  CTF → Exchange: %s\n", txHash.Hex())
		txHashes = append(txHashes, txHash)
	} else {
		fmt.Printf("  CTF → Exchange: already approved\n")
	}

	// Check and approve CTF for NegRiskAdapter
	ctfApprovedNegRiskAdapter, err := ctfContract.IsApprovedForAll(callOpts, eoaAddr, negRiskAdapterAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for NegRiskAdapter: %w", err)
	}
	if !ctfApprovedNegRiskAdapter {
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", negRiskAdapterAddress, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack setApprovalForAll data for NegRiskAdapter: %w", err)
		}
		txHash, err := b.txSender.SendEthereumTransaction(ctfAddress, setApprovalData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send CTF → NegRiskAdapter approval transaction: %w", err)
		}
		fmt.Printf("  CTF → NegRiskAdapter: %s\n", txHash.Hex())
		txHashes = append(txHashes, txHash)
	} else {
		fmt.Printf("  CTF → NegRiskAdapter: already approved\n")
	}

	// Check and approve CTF for NegRiskExchange
	ctfApprovedNegRiskExchange, err := ctfContract.IsApprovedForAll(callOpts, eoaAddr, negRiskExchangeAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check CTF approval for NegRiskExchange: %w", err)
	}
	if !ctfApprovedNegRiskExchange {
		setApprovalData, err := ctfABI.Pack("setApprovalForAll", negRiskExchangeAddress, true)
		if err != nil {
			return nil, fmt.Errorf("failed to pack setApprovalForAll data for NegRiskExchange: %w", err)
		}
		txHash, err := b.txSender.SendEthereumTransaction(ctfAddress, setApprovalData, big.NewInt(0))
		if err != nil {
			return nil, fmt.Errorf("failed to send CTF → NegRiskExchange approval transaction: %w", err)
		}
		fmt.Printf("  CTF → NegRiskExchange: %s\n", txHash.Hex())
		txHashes = append(txHashes, txHash)
	} else {
		fmt.Printf("  CTF → NegRiskExchange: already approved\n")
	}

	return txHashes, nil
}

// RedeemPositionsForEOA redeems conditional tokens for a resolved market using EOA
// ctfAddress: the ConditionalTokens contract address to call
func (b *ContractInterface) RedeemPositionsForEOA(
	ctx context.Context,
	ctfAddress common.Address,
	conditionId [32]byte,
	indexSets []*big.Int,
) (common.Hash, error) {
	// Prepare calldata for redeemPositions
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("redeemPositions", b.contractConfig.Collateral, parentCollectionId, conditionId, indexSets)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
	}

	// Send transaction
	txHash, err := b.txSender.SendEthereumTransaction(ctfAddress, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send redeem transaction: %w", err)
	}

	return txHash, nil
}

// RedeemPositionsNegRiskForEOA redeems NegRisk market positions using EOA
// negRiskAdapterAddress: the NegRiskAdapter contract address to call
func (b *ContractInterface) RedeemPositionsNegRiskForEOA(
	ctx context.Context,
	negRiskAdapterAddress common.Address,
	conditionId [32]byte,
	amounts []*big.Int,
) (common.Hash, error) {
	// Prepare calldata for redeemPositions (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("redeemPositions", conditionId, amounts)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack redeemPositions calldata: %w", err)
	}

	// Send transaction
	txHash, err := b.txSender.SendEthereumTransaction(negRiskAdapterAddress, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send NegRisk redeem transaction: %w", err)
	}

	return txHash, nil
}

// SplitPositionForEOA splits collateral into conditional tokens for an EOA wallet
// ctfAddress: the ConditionalTokens contract address to call
func (b *ContractInterface) SplitPositionForEOA(
	ctx context.Context,
	ctfAddress common.Address,
	conditionId [32]byte,
	partition []*big.Int,
	amount *big.Int,
) (common.Hash, error) {
	// Prepare calldata for splitPosition
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("splitPosition", b.contractConfig.Collateral, parentCollectionId, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack splitPosition calldata: %w", err)
	}

	// Send transaction
	txHash, err := b.txSender.SendEthereumTransaction(ctfAddress, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send split transaction: %w", err)
	}

	return txHash, nil
}

// MergePositionsForEOA merges conditional tokens back into collateral for an EOA wallet
// ctfAddress: the ConditionalTokens contract address to call
func (b *ContractInterface) MergePositionsForEOA(
	ctx context.Context,
	ctfAddress common.Address,
	conditionId [32]byte,
	partition []*big.Int,
	amount *big.Int,
) (common.Hash, error) {
	// Prepare calldata for mergePositions
	parsedABI, err := conditional_tokens.ConditionalTokensMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse ConditionalTokens ABI: %w", err)
	}
	parentCollectionId := [32]byte{} // empty for root collection
	calldata, err := parsedABI.Pack("mergePositions", b.contractConfig.Collateral, parentCollectionId, conditionId, partition, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack mergePositions calldata: %w", err)
	}

	// Send transaction
	txHash, err := b.txSender.SendEthereumTransaction(ctfAddress, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send merge transaction: %w", err)
	}

	return txHash, nil
}

// SplitPositionNegRiskForEOA splits collateral into NegRisk conditional tokens for an EOA wallet
// negRiskAdapterAddress: the NegRiskAdapter contract address to call
func (b *ContractInterface) SplitPositionNegRiskForEOA(
	ctx context.Context,
	negRiskAdapterAddress common.Address,
	conditionId [32]byte,
	amount *big.Int,
) (common.Hash, error) {
	// Prepare calldata for splitPosition (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("splitPosition", conditionId, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack splitPosition calldata: %w", err)
	}

	// Send transaction
	txHash, err := b.txSender.SendEthereumTransaction(negRiskAdapterAddress, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send NegRisk split transaction: %w", err)
	}

	return txHash, nil
}

// MergePositionsNegRiskForEOA merges NegRisk conditional tokens back into collateral for an EOA wallet
// negRiskAdapterAddress: the NegRiskAdapter contract address to call
func (b *ContractInterface) MergePositionsNegRiskForEOA(
	ctx context.Context,
	negRiskAdapterAddress common.Address,
	conditionId [32]byte,
	amount *big.Int,
) (common.Hash, error) {
	// Prepare calldata for mergePositions (NegRisk)
	parsedABI, err := negriskadapter.NegRiskAdapterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse NegRiskAdapter ABI: %w", err)
	}
	calldata, err := parsedABI.Pack("mergePositions", conditionId, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack mergePositions calldata: %w", err)
	}

	// Send transaction
	txHash, err := b.txSender.SendEthereumTransaction(negRiskAdapterAddress, calldata, big.NewInt(0))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send NegRisk merge transaction: %w", err)
	}

	return txHash, nil
}
