package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/v1/prepare"
)

// BootstrapTransactions prepares a transaction group to bootstrap a new pool
// validatorAppID, asset1ID, asset2ID are converted to uint64
func BootstrapTransactions(
	validatorAppID,
	asset1ID,
	asset2ID string,
	asset1UnitName,
	asset2UnitName,
	senderAddress string,
	suggestedParams *SuggestedParams,
) (*TransactionGroup, error) {
	if suggestedParams == nil {
		return nil, fmt.Errorf("suggestedParams is required")
	}
	uintValidatorAppID, err := strconv.ParseUint(validatorAppID, 10, 64)
	if err != nil {
		return nil, err
	}
	uintAsset1ID, err := strconv.ParseUint(asset1ID, 10, 64)
	if err != nil {
		return nil, err
	}
	uintAsset2ID, err := strconv.ParseUint(asset2ID, 10, 64)
	if err != nil {
		return nil, err
	}

	txGroup, err := prepare.BootstrapTransactions(uintValidatorAppID, uintAsset1ID, uintAsset2ID, asset1UnitName, asset2UnitName, senderAddress, *suggestedParams.wrapped)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
