package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/v1/prepare"
)

// PrepareSwapTransactions Prepare a transaction group to swap assets.
// validatorAppID, asset1ID, asset2ID, liquidityAssetID, assetInID, assetInAmount, assetOutAmount are converted to uint64,
func PrepareSwapTransactions(
	validatorAppID,
	asset1ID,
	asset2ID,
	liquidityAssetID,
	assetInID,
	assetInAmount,
	assetOutAmount,
	swapType,
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

	uintAssetInID, err := strconv.ParseUint(assetInID, 10, 64)
	if err != nil {
		return nil, err
	}

	uintAssetInAmount, err := strconv.ParseUint(assetInAmount, 10, 64)
	if err != nil {
		return nil, err
	}

	uintAssetOutAmount, err := strconv.ParseUint(assetOutAmount, 10, 64)
	if err != nil {
		return nil, err
	}

	txGroup, err := prepare.SwapTransactions(
		uintValidatorAppID,
		uintAsset1ID,
		uintAsset2ID,
		uintLiquidityAssetID,
		uintAssetInID,
		uintAssetInAmount,
		uintAssetOutAmount,
		swapType,
		senderAddress,
		*suggestedParams.wrapped,
	)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
