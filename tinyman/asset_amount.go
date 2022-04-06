package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// AssetAmount represents an asset amount
type AssetAmount struct {
	asset              *Asset
	wrappedAssetAmount *types.AssetAmount
}

//
// NewAssetAmount creates a new asset amount. Note that amount is a string here
// Eventually it will be converted to 64-bit unsigned integer
func NewAssetAmount(asset *Asset, amount string) (*AssetAmount, error) {
	if asset == nil {
		return nil, fmt.Errorf("an asset cannot be nil")
	}

	a := &AssetAmount{
		asset: asset,
		wrappedAssetAmount: &types.AssetAmount{
			Asset:  asset.wrappedAsset,
			Amount: 0,
		},
	}

	if err := a.SetAmount(amount); err != nil {
		return nil, err
	}

	return a, nil
}

// SetAsset sets an asset
func (a *AssetAmount) SetAsset(asset *Asset) error {
	if asset == nil {
		return fmt.Errorf("an asset cannot be nil")
	}

	a.asset = asset
	a.wrappedAssetAmount.Asset = asset.wrappedAsset

	return nil
}

// SetAmount sets an amount the value will be converted to 64-bit unsigned integer
func (a *AssetAmount) SetAmount(value string) error {
	amount, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	a.wrappedAssetAmount.Amount = amount

	return nil
}

// Asset returns an asset
func (a *AssetAmount) Asset() *Asset {
	return a.asset
}

// AssetAmount returns an asset by converting the underlying 64-bit unsigned integer to a string
func (a *AssetAmount) AssetAmount() string {
	return strconv.FormatUint(a.wrappedAssetAmount.Amount, 10)
}
