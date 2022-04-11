package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// MintQuote represents a mint quote
type MintQuote struct {
	wrapped *types.MintQuote
}

// AssetAmountsInIterator returns an iterator for iterating input asset amounts
func (m *MintQuote) AssetAmountsInIterator() *AssetAmountIterator {
	var aa []*AssetAmount
	for _, v := range m.wrapped.AmountsIn {
		v := v
		aa = append(aa, wrapAssetAmount(&v))
	}

	return &AssetAmountIterator{
		curr:   0,
		values: aa,
	}
}

// LiquidityAssetAmount returns a liquidity asset amount
func (m *MintQuote) LiquidityAssetAmount() *AssetAmount {
	return wrapAssetAmount(&m.wrapped.LiquidityAssetAmount)
}

// LiquidityAssetAmountWithSlippage calculates liquidity asset after applying the slippage
func (m *MintQuote) LiquidityAssetAmountWithSlippage() (*AssetAmount, error) {
	assetAmount, err := m.wrapped.LiquidityAssetAmountWithSlippage()
	if err != nil {
		return nil, err
	}

	return wrapAssetAmount(assetAmount), nil
}

// Slippage returns a slippage
func (m *MintQuote) Slippage() string {
	return strconv.FormatFloat(m.wrapped.Slippage, 'f', -1, 64)
}
