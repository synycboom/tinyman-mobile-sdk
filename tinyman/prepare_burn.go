package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/v1/prepare"
)

// BurnTransactions prepares a transaction group to burn the liquidity pool asset amount in exchange for pool assets.
// validatorappID, asset1ID, asset2ID, liquidityAssetID, asset1Amount, asset2Amount, liquidityAssetAmount are converted to uint64
func BurnTransactions(
	validatorAppID,
	asset1ID,
	asset2ID,
	liquidityAssetID,
	asset1Amount,
	asset2Amount,
	liquidityAssetAmount,
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

	uintLiquidityAssetID, err := strconv.ParseUint(liquidityAssetID, 10, 64)
	if err != nil {
		return nil, err
	}

	uintAsset1Amount, err := strconv.ParseUint(asset1Amount, 10, 64)
	if err != nil {
		return nil, err
	}

	uintAsset2Amount, err := strconv.ParseUint(asset2Amount, 10, 64)
	if err != nil {
		return nil, err
	}

	uintLiquidityAssetAmount, err := strconv.ParseUint(liquidityAssetAmount, 10, 64)
	if err != nil {
		return nil, err
	}

	txGroup, err := prepare.BurnTransactions(
		uintValidatorAppID,
		uintAsset1ID,
		uintAsset2ID,
		uintLiquidityAssetID,
		uintAsset1Amount,
		uintAsset2Amount,
		uintLiquidityAssetAmount,
		senderAddress,
		*suggestedParams.wrapped,
	)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
