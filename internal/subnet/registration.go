package subnet

import (
    "math/big"
    "time"
)

// RegistrationConfig holds subnet registration settings
type RegistrationConfig struct {
    BaseRegistrationCost *big.Int
    NetworkUtilization  uint64
    LastUpdateTime      time.Time
}

// RegistrationManager handles subnet registration costs
type RegistrationManager struct {
    config   *RegistrationConfig
    maxCost  *big.Int
    minCost  *big.Int
}

func NewRegistrationManager(baseRegistrationCost *big.Int) *RegistrationManager {
    return &RegistrationManager{
        config: &RegistrationConfig{
            BaseRegistrationCost: baseRegistrationCost,
            NetworkUtilization:  0,
            LastUpdateTime:      time.Now(),
        },
        maxCost: new(big.Int).Mul(baseRegistrationCost, big.NewInt(5)),  // 5x max multiplier
        minCost: new(big.Int).Div(baseRegistrationCost, big.NewInt(2)),  // 0.5x min multiplier
    }
}

// CalculateRegistrationCost calculates optimized registration cost
func (rm *RegistrationManager) CalculateRegistrationCost() *big.Int {
    cost := new(big.Int).Set(rm.config.BaseRegistrationCost)
    
    // Adjust based on network utilization
    if rm.config.NetworkUtilization > 80 {
        // High utilization: increase cost
        cost.Mul(cost, big.NewInt(2))
    } else if rm.config.NetworkUtilization < 20 {
        // Low utilization: decrease cost
        cost.Div(cost, big.NewInt(2))
    }
    
    // Ensure cost stays within bounds
    if cost.Cmp(rm.maxCost) > 0 {
        return new(big.Int).Set(rm.maxCost)
    }
    if cost.Cmp(rm.minCost) < 0 {
        return new(big.Int).Set(rm.minCost)
    }
    
    return cost
}

// UpdateNetworkUtilization updates the current network utilization
func (rm *RegistrationManager) UpdateNetworkUtilization(utilization uint64) {
    rm.config.NetworkUtilization = utilization
    rm.config.LastUpdateTime = time.Now()
}
