# Predict Go Contracts

A Go library for interacting with Predict.fun smart contracts on BNB Chain. This library provides a convenient interface for trading on Predict.fun using EOA (Externally Owned Account) wallets.

## Features

- 🔐 Support for EOA signers
- 🔄 Complete trading operations (enable trading, split, merge, redeem)
- ⚡ Built on go-ethereum
- 🚀 Enhanced transaction inclusion speed (1.3x gas price multiplier)

## Installation

```bash
go get github.com/ivanzzeth/predict-go-contracts
```

## Quick Start

### EOA Signer (Direct Private Key)

EOA signer allows you to interact with Predict.fun directly using a private key. This is the simplest approach for individual accounts.

```go
package main

import (
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ivanzzeth/ethsig"
    predictcontracts "github.com/ivanzzeth/predict-go-contracts"
)

func main() {
    // Connect to BNB Chain RPC
    client, _ := ethclient.Dial(os.Getenv("RPC_URL"))

    // Load private key
    privateKey, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

    // Create EOA signer
    eoaSigner := ethsig.NewEthPrivateKeySigner(privateKey)

    // Initialize contract interface
    config := predictcontracts.GetContractConfig(big.NewInt(56)) // BNB Chain mainnet
    predictInterface, _ := predictcontracts.NewContractInterface(
        client,
        predictcontracts.WithContractConfig(config),
        predictcontracts.WithEOASigner(eoaSigner),
    )

    // Enable trading
    ctx := context.Background()
    txHashes, _ := predictInterface.EnableTrading(ctx)
}
```

## Environment Variables

```bash
# Required for all examples
RPC_URL=https://bsc-dataseed1.binance.org
PRIVATE_KEY=your_private_key_hex

# Required for split/merge examples
CONDITION_ID=0x... # The condition ID from Predict.fun
AMOUNT=100         # Amount in USDT (for split) or tokens (for merge)
```

## Examples

For complete, runnable examples, see the [`examples/`](./examples) directory:

### EOA Examples

- **[`enable_trading`](./examples/enable_trading)** - Enable trading for an EOA account
  - Direct USDT and CTF token approvals
  - Simplest setup for individual traders

## Running Examples

### Basic Example (Enable Trading)

```bash
# Navigate to an example directory
cd examples/enable_trading

# Set environment variables
export RPC_URL=https://bsc-dataseed1.binance.org
export PRIVATE_KEY=your_private_key_hex

# Run the example
go run main.go
```


## Core Operations

### Enable Trading

Sets up all required token approvals for trading on Predict.fun:
- USDT approvals for Exchange, NegRiskAdapter, and NegRiskExchange
- CTF token approvals for the same contracts

```go
txHashes, err := predictInterface.EnableTrading(ctx)
```

### Split Position

Split collateral tokens into conditional tokens. The partition parameter is automatically set to [1, 2] for Predict.fun's binary markets (Yes/No outcomes):

```go
txHash, err := predictInterface.Split(ctx, conditionId, amount)
```

### Merge Positions

Merge conditional tokens back to collateral. The partition parameter is automatically set to [1, 2] for Predict.fun's binary markets:

```go
txHash, err := predictInterface.Merge(ctx, conditionId, amount)
```

### Redeem Positions

Redeem conditional tokens after market resolution. The indexSets parameter is automatically set to [1, 2] for Predict.fun's binary markets:

```go
txHash, err := predictInterface.Redeem(ctx, conditionId)
```

## Architecture

```
predict-go-contracts/
├── interface.go              # Main contract interface
├── config.go                 # Contract addresses and configs
├── types.go                  # Type definitions
├── signer/                   # Signing implementations
│   ├── eoa_trading_signer.go     # EOA signer interface
│   └── transaction_sender.go     # Transaction sending logic
├── sender/                   # Transaction sender interface
├── contracts/                # Generated contract bindings
└── examples/                 # Complete usage examples
```

## Supported Networks

- **BNB Chain Mainnet** (Chain ID: 56)
- **BNB Chain Testnet** (Chain ID: 97)

Pre-configured contract addresses are available via `GetContractConfig()`.

## Security Considerations

- **Private Keys**: Never commit private keys to version control
- **Environment Variables**: Use `.env` files or secure secret management
- **Transaction Inclusion**: The library automatically applies a 1.3x multiplier to suggested gas prices to improve transaction inclusion speed on BNB Chain
- **Gas Estimation**: Always verify gas costs before mainnet deployment

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

For issues and questions:
- Open an issue on [GitHub](https://github.com/ivanzzeth/predict-go-contracts/issues)
- Check the [`examples/`](./examples) directory for working code

## Links

- [Predict.fun](https://predict.fun/)
