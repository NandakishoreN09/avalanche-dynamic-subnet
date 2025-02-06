# Custom Avalanche Subnet Implementation

This project implements a custom Avalanche subnet with the following features:
- Dynamic fee structure for optimized transaction costs
- Cross-subnet communication using Avalanche Warp Messaging (AWM)
- Validator management system
- Optimized subnet registration costs

## Project Structure
## Features

### 1. Dynamic Fee Structure
- Base fee adjustment based on network load
- Multiplier system for peak usage periods
- Fee optimization for different transaction types

### 2. Cross-Subnet Communication (AWM)
- Secure message passing between subnets
- Support for cross-chain asset transfers
- Message verification and status tracking

### 3. Validator Management
- Stake-based validator system
- Dynamic validator set updates
- Performance monitoring

### 4. Registration Cost Optimization
- Network utilization-based cost adjustment
- Automatic price discovery
- Protection against network spam

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Avalanche CLI
- Solidity Compiler

### Installation
```bash
git clone <your-repo-url>
cd avalanche-subnet
go mod tidy
go run cmd/main/main.go
config := subnet.SubnetConfig{
    MinStake:       big.NewInt(1000),
    MaxValidators:  100,
    BaseFee:        big.NewInt(1),
    DynamicFeeMult: 2,
}

mySubnet := subnet.NewSubnet(config)
err := mySubnet.AddValidator(validatorAddr, stake)
warp := subnet.NewWarpMessenger(mySubnet)
messageID, err := warp.SendMessage(
    "destination-chain-123",
    payload,
    amount,
)
2. Create a .gitignore:
```bash
cat > .gitignore << 'EOL'
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
vendor/

# Go workspace file
go.work

# IDE specific files
.idea/
.vscode/
*.swp
*.swo

# OS specific files
.DS_Store
