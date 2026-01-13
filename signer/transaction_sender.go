package signer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ivanzzeth/ethsig"
	"github.com/ivanzzeth/predict-go-contracts/sender"
)

// GetTransactionSenderBySigner creates a TransactionSender based on the signer type
func GetTransactionSenderBySigner(chainId *big.Int, client bind.ContractBackend, signerInstance any) (sender.TransactionSender, error) {
	switch s := signerInstance.(type) {
	case TransactionSignerAndAddrGetter:
		return GetTransactionSenderByTransactionSignerAndAddrGetter(chainId, client, s)
	default:
		return nil, fmt.Errorf("unsupported signer type %T", signerInstance)
	}
}

// TransactionSignerAndAddrGetter combines transaction signing and address retrieval
type TransactionSignerAndAddrGetter interface {
	ethsig.AddressGetter
	ethsig.TransactionSigner
}

// TransactionSenderByTransactionSigner implements TransactionSender using a transaction signer
type TransactionSenderByTransactionSigner struct {
	chainId  *big.Int
	client   bind.ContractBackend
	txSigner TransactionSignerAndAddrGetter
}

// GetTransactionSenderByTransactionSignerAndAddrGetter creates a TransactionSender from a transaction signer
func GetTransactionSenderByTransactionSignerAndAddrGetter(chainId *big.Int, client bind.ContractBackend, txSigner TransactionSignerAndAddrGetter) (sender.TransactionSender, error) {
	return &TransactionSenderByTransactionSigner{chainId: chainId, client: client, txSigner: txSigner}, nil
}

// SendEthereumTransaction sends an Ethereum transaction using the transaction signer
func (s *TransactionSenderByTransactionSigner) SendEthereumTransaction(to common.Address, data []byte, value *big.Int) (common.Hash, error) {
	ctx := context.Background()

	// Get nonce
	nonce, err := s.client.PendingNonceAt(ctx, s.txSigner.GetAddress())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get gas price and increase by 30% to improve transaction inclusion speed
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get gas price: %w", err)
	}
	gasPrice.Mul(gasPrice, big.NewInt(13))
	gasPrice.Div(gasPrice, big.NewInt(10))

	// Estimate gas limit
	msg := ethereum.CallMsg{
		From:     s.txSigner.GetAddress(),
		To:       &to,
		Value:    value,
		Data:     data,
		GasPrice: gasPrice,
	}
	gasLimit, err := s.client.EstimateGas(ctx, msg)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// Create and sign the transaction
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := s.txSigner.SignTransactionWithChainID(tx, s.chainId)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send the signed transaction
	err = s.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	return signedTx.Hash(), nil
}
