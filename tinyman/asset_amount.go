package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// AssetAmountIterator is an iterator for iterating amounts
type AssetAmountIterator struct {
	curr   int
	values []*AssetAmount
}

// HasNext return true if there are asset amounts to be iterated
func (a *AssetAmountIterator) HasNext() bool {
	return a.curr < len(a.values)
}

// Next returns the next a asset amount, returns nil if no asset amounts left
func (a *AssetAmountIterator) Next() *AssetAmount {
	if a.HasNext() {
		idx := a.curr
		a.curr += 1

		return a.values[idx]
	}

	return nil
}

// AssetAmount represents an asset amount
type AssetAmount struct {
	wrappedAssetAmount *types.AssetAmount
}

// NewAssetAmount creates a new asset amount. Note that amount is a string here
// Eventually it will be converted to 64-bit unsigned integer
func NewAssetAmount() *AssetAmount {
	return &AssetAmount{
		wrappedAssetAmount: &types.AssetAmount{},
	}
}

// SetAsset sets an asset
func (a *AssetAmount) SetAsset(asset *Asset) error {
	if asset == nil {
		return fmt.Errorf("asset is required")
	}

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
	return unwrapAsset(a.wrappedAssetAmount.Asset)
}

// AssetAmount returns an asset by converting the underlying 64-bit unsigned integer to a string
func (a *AssetAmount) AssetAmount() string {
	return strconv.FormatUint(a.wrappedAssetAmount.Amount, 10)
}

func unwrapAsset(wrappedAsset *types.Asset) *Asset {
	asset := Asset{}
	asset.SetDecimals(strconv.FormatUint(wrappedAsset.Decimals, 10))
	asset.SetID(strconv.FormatUint(wrappedAsset.ID, 10))
	asset.SetName(wrappedAsset.Name)
	asset.SetUnitName(wrappedAsset.UnitName)

	return &asset
}

func unwrapAssetAmount(wrapped *types.AssetAmount) *AssetAmount {
	assetAmount := NewAssetAmount()
	assetAmount.SetAmount(strconv.FormatUint(wrapped.Amount, 10))
	assetAmount.SetAsset(unwrapAsset(wrapped.Asset))

	return assetAmount
}
