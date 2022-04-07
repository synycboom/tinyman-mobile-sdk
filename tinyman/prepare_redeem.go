package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/v1/prepare"
)

// PrepareRedeemTransactions prepares a transaction group to redeem a specified excess asset amount from a pool.
// validatorAppID, asset1ID, asset2ID, liquidityAssetID, assetID, assetAmount are converted to uint64
func PrepareRedeemTransactions(
	validatorAppID,
	asset1ID,
	asset2ID,
	liquidityAssetID,
	assetID,
	assetAmount,
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

	uintAssetID, err := strconv.ParseUint(assetID, 10, 64)
	if err != nil {
		return nil, err
	}

	uintAssetAmount, err := strconv.ParseUint(assetAmount, 10, 64)
	if err != nil {
		return nil, err
	}

	txGroup, err := prepare.RedeemTransactions(
		uintValidatorAppID,
		uintAsset1ID,
		uintAsset2ID,
		uintLiquidityAssetID,
		uintAssetID,
		uintAssetAmount,
		senderAddress,
		*suggestedParams.wrapped,
	)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
