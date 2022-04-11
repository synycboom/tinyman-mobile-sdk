package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/v1/contracts"
)

// PoolLogicSigAccount creates a logic signature account of the pool
func PoolLogicSigAccount(validatorAppID, asset1ID, asset2ID string) (*LogicSigAccount, error) {
	if len(validatorAppID) == 0 {
		return nil, fmt.Errorf("validatorAppID is required")
	}
	if len(asset1ID) == 0 {
		return nil, fmt.Errorf("asset1ID is required")
	}
	if len(asset2ID) == 0 {
		return nil, fmt.Errorf("asset2ID is required")
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

	return poolLogicSigAccount(uintValidatorAppID, uintAsset1ID, uintAsset2ID)
}

func poolLogicSigAccount(validatorAppID, asset1ID, asset2ID uint64) (*LogicSigAccount, error) {
	lsigAccount, err := contracts.PoolLogicSigAccount(validatorAppID, asset1ID, asset2ID)
	if err != nil {
		return nil, err
	}

	return &LogicSigAccount{wrapped: lsigAccount}, nil
}
