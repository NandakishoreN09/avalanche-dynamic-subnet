package subnet

import (
    "encoding/hex"
    "math/big"
    "errors"
)

// WarpMessage represents a cross-subnet message
type WarpMessage struct {
    SourceChainID      string
    DestinationChainID string
    Payload            []byte
    Amount             *big.Int
    Status             string
}

// WarpMessenger handles cross-subnet communication
type WarpMessenger struct {
    subnet          *Subnet
    pendingMessages map[string]*WarpMessage
    processedMessages map[string]bool
}

// NewWarpMessenger creates a new warp messenger instance
func NewWarpMessenger(s *Subnet) *WarpMessenger {
    return &WarpMessenger{
        subnet:            s,
        pendingMessages:   make(map[string]*WarpMessage),
        processedMessages: make(map[string]bool),
    }
}

// SendMessage sends a warp message to another subnet
func (w *WarpMessenger) SendMessage(destChainID string, payload []byte, amount *big.Int) (string, error) {
    if len(payload) < 20 {
        return "", errors.New("payload too short")
    }

    msg := &WarpMessage{
        SourceChainID:      "current-chain",
        DestinationChainID: destChainID,
        Payload:            payload,
        Amount:             amount,
        Status:             "pending",
    }
    
    messageID := hex.EncodeToString(payload[:20])
    w.pendingMessages[messageID] = msg
    
    return messageID, nil
}

// GetMessageStatus retrieves the status of a message
func (w *WarpMessenger) GetMessageStatus(messageID string) (string, error) {
    if msg, exists := w.pendingMessages[messageID]; exists {
        return msg.Status, nil
    }
    if w.processedMessages[messageID] {
        return "completed", nil
    }
    return "", errors.New("message not found")
}
