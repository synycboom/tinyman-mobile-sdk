package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/v1/prepare"
)

// PrepareAppOptInTransactions prepares a transaction group to opt-in of Tinyman
// validatorAppID is converted to uint64
func PrepareAppOptInTransactions(validatorAppID string, senderAddress string, suggestedParams *SuggestedParams) (*TransactionGroup, error) {
	if suggestedParams == nil {
		return nil, fmt.Errorf("suggestedParams is required")
	}

	uintValidatorAppID, err := strconv.ParseUint(validatorAppID, 10, 64)
	if err != nil {
		return nil, err
	}

	txGroup, err := prepare.AppOptInTransactions(uintValidatorAppID, senderAddress, *suggestedParams.wrapped)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}

// PrepareAssetOptInTransactions prepares a transaction group to opt-in an asset
// assetID is converted to uint64
func PrepareAssetOptInTransactions(assetID string, senderAddress string, suggestedParams *SuggestedParams) (*TransactionGroup, error) {
	if suggestedParams == nil {
		return nil, fmt.Errorf("suggestedParams is required")
	}

	uintAssetID, err := strconv.ParseUint(assetID, 10, 64)
	if err != nil {
		return nil, err
	}

	txGroup, err := prepare.AssetOptInTransactions(uintAssetID, senderAddress, *suggestedParams.wrapped)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
