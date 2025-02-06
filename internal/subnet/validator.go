package subnet

import (
    "errors"
    "math/big"
)

// ValidatorManager handles validator operations
type ValidatorManager struct {
    subnet *Subnet
}

// AddValidator adds a new validator to the subnet
func (s *Subnet) AddValidator(address string, stake *big.Int) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Validation checks
    if stake.Cmp(s.config.MinStake) < 0 {
        return errors.New("stake amount below minimum requirement")
    }

    if len(s.validators) >= int(s.config.MaxValidators) {
        return errors.New("maximum validator limit reached")
    }

    // Add validator
    s.validators[address] = &Validator{
        Address:     address,
        StakeAmount: stake,
        IsActive:    true,
    }

    return nil
}

// RemoveValidator removes a validator from the subnet
func (s *Subnet) RemoveValidator(address string) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, exists := s.validators[address]; !exists {
        return errors.New("validator not found")
    }

    delete(s.validators, address)
    return nil
}
