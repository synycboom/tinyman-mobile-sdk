package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// MintQuote represents a mint quote
type MintQuote struct {
	wrapped *types.MintQuote
}

// GetAssetAmountsInIterator returns an iterator for iterating input asset amounts
func (m *MintQuote) GetAssetAmountsInIterator() *AssetAmountIterator {
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

// GetLiquidityAssetAmount returns a liquidity asset amount
func (m *MintQuote) GetLiquidityAssetAmount() *AssetAmount {
	return wrapAssetAmount(&m.wrapped.LiquidityAssetAmount)
}

// GetLiquidityAssetAmountWithSlippage calculates liquidity asset after applying the slippage
func (m *MintQuote) GetLiquidityAssetAmountWithSlippage() (*AssetAmount, error) {
	assetAmount, err := m.wrapped.LiquidityAssetAmountWithSlippage()
	if err != nil {
		return nil, err
	}

	return wrapAssetAmount(assetAmount), nil
}

// GetSlippage returns a slippage
func (m *MintQuote) GetSlippage() string {
	return strconv.FormatFloat(m.wrapped.Slippage, 'f', -1, 64)
}
