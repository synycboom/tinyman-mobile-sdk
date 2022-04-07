package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/v1/prepare"
)

// PrepareAppOptOutTransactions prepares a transaction group to opt-out of Tinyman
// validatorAppID is converted to uint64
func PrepareAppOptOutTransactions(validatorAppID string, senderAddress string, suggestedParams *SuggestedParams) (*TransactionGroup, error) {
	if suggestedParams == nil {
		return nil, fmt.Errorf("suggestedParams is required")
	}

	uintValidatorAppID, err := strconv.ParseUint(validatorAppID, 10, 64)
	if err != nil {
		return nil, err
	}

	txGroup, err := prepare.AppOptOutTransactions(uintValidatorAppID, senderAddress, *suggestedParams.wrapped)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
