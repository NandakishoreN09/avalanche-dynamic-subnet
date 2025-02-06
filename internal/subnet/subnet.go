package subnet

import (
    "math/big"
    "sync"
)

// SubnetConfig holds the configuration for our custom subnet
type SubnetConfig struct {
    MinStake       *big.Int
    MaxValidators  uint
    BaseFee        *big.Int
    DynamicFeeMult uint64
}

// Subnet represents our custom subnet implementation
type Subnet struct {
    config         SubnetConfig
    validators     map[string]*Validator
    feeCalculator  *FeeCalculator
    mu            sync.RWMutex
}

// Validator represents a subnet validator
type Validator struct {
    Address     string
    StakeAmount *big.Int
    IsActive    bool
}

// FeeCalculator handles dynamic fee calculations
type FeeCalculator struct {
    baseFee    *big.Int
    multiplier uint64
}

// NewSubnet creates a new subnet instance
func NewSubnet(config SubnetConfig) *Subnet {
    return &Subnet{
        config:     config,
        validators: make(map[string]*Validator),
        feeCalculator: &FeeCalculator{
            baseFee:    config.BaseFee,
            multiplier: config.DynamicFeeMult,
        },
    }
}

// FeeCalculator returns the fee calculator instance
func (s *Subnet) FeeCalculator() *FeeCalculator {
    return s.feeCalculator
}
