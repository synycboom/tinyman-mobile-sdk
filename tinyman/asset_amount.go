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

// Reset resets the iterator
func (a *AssetAmountIterator) Reset() {
	a.curr = 0
}

// Add adds an item to the iterator
func (a *AssetAmountIterator) Add(item *AssetAmount) {
	a.values = append(a.values, item)
}

// AssetAmount represents an asset amount
type AssetAmount struct {
	wrapped *types.AssetAmount
}

// NewAssetAmount creates a new asset amount. Note that amount is a string here
// Eventually it will be converted to 64-bit unsigned integer
func NewAssetAmount() *AssetAmount {
	return &AssetAmount{
		wrapped: &types.AssetAmount{},
	}
}

// SetAsset sets an asset
func (a *AssetAmount) SetAsset(asset *Asset) error {
	if asset == nil {
		return fmt.Errorf("asset is required")
	}

	a.wrapped.Asset = asset.wrapped

	return nil
}

// SetAmount sets an amount the value will be converted to 64-bit unsigned integer
func (a *AssetAmount) SetAmount(value string) error {
	amount, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	a.wrapped.Amount = amount

	return nil
}

// GetAsset returns an asset
func (a *AssetAmount) GetAsset() *Asset {
	return wrapAsset(a.wrapped.Asset)
}

// GetAmount returns an asset by converting the underlying 64-bit unsigned integer to a string
func (a *AssetAmount) GetAmount() string {
	return strconv.FormatUint(a.wrapped.Amount, 10)
}

func wrapAsset(asset *types.Asset) *Asset {
	return &Asset{wrapped: asset}
}

func wrapAssetAmount(assetAmount *types.AssetAmount) *AssetAmount {
	return &AssetAmount{wrapped: assetAmount}
}
