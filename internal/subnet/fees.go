package subnet

import (
    "math/big"
)

// CalculateFee calculates the dynamic fee based on network load
func (fc *FeeCalculator) CalculateFee(networkLoad uint64) *big.Int {
    // Dynamic fee calculation based on network load
    fee := new(big.Int).Set(fc.baseFee)
    loadMultiplier := new(big.Int).SetUint64((networkLoad * fc.multiplier) / 100 + 1) // +1 to avoid multiplying by zero
    return fee.Mul(fee, loadMultiplier)
}

// GetBaseFee returns the current base fee
func (fc *FeeCalculator) GetBaseFee() *big.Int {
    return fc.baseFee
}

// GetMultiplier returns the current multiplier
func (fc *FeeCalculator) GetMultiplier() uint64 {
    return fc.multiplier
}
