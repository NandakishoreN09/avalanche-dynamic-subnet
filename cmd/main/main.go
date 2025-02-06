package main

import (
    "fmt"
    "math/big"
    "crypto/rand"
    "avalanche-subnet/internal/subnet"
)

func generateRandomPayload(size int) []byte {
    payload := make([]byte, size)
    rand.Read(payload)
    return payload
}

func main() {
    // Initialize subnet configuration
    config := subnet.SubnetConfig{
        MinStake:       big.NewInt(1000),
        MaxValidators:  100,
        BaseFee:        big.NewInt(1),
        DynamicFeeMult: 2,
    }

    // Create new subnet
    mySubnet := subnet.NewSubnet(config)

    // Initialize registration manager
    regManager := subnet.NewRegistrationManager(big.NewInt(1000))

    // Test different network conditions
    fmt.Println("\n--- Testing Registration Costs ---")
    
    // Low utilization
    regManager.UpdateNetworkUtilization(10)
    lowCost := regManager.CalculateRegistrationCost()
    fmt.Printf("Registration cost at 10%% utilization: %v\n", lowCost)

    // Normal utilization
    regManager.UpdateNetworkUtilization(50)
    normalCost := regManager.CalculateRegistrationCost()
    fmt.Printf("Registration cost at 50%% utilization: %v\n", normalCost)

    // High utilization
    regManager.UpdateNetworkUtilization(90)
    highCost := regManager.CalculateRegistrationCost()
    fmt.Printf("Registration cost at 90%% utilization: %v\n", highCost)

    // Test other functionality
    validatorAddr := "0x1234567890123456789012345678901234567890"
    stake := big.NewInt(2000)
    
    err := mySubnet.AddValidator(validatorAddr, stake)
    if err != nil {
        fmt.Printf("Error adding validator: %v\n", err)
        return
    }
    fmt.Printf("\nValidator added successfully with stake: %v\n", stake)

    fee := mySubnet.FeeCalculator().CalculateFee(50)
    fmt.Printf("Current fee with 50%% network load: %v\n", fee)

    warp := subnet.NewWarpMessenger(mySubnet)
    payload := generateRandomPayload(32)
    messageID, err := warp.SendMessage(
        "destination-chain-123",
        payload,
        big.NewInt(100),
    )
    if err != nil {
        fmt.Printf("Error sending warp message: %v\n", err)
        return
    }
    fmt.Printf("Warp message sent with ID: %s\n", messageID)

    status, err := warp.GetMessageStatus(messageID)
    if err != nil {
        fmt.Printf("Error checking message status: %v\n", err)
        return
    }
    fmt.Printf("Message status: %s\n", status)
}
