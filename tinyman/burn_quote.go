package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// BurnQuote represents a burn quote
type BurnQuote struct {
	wrapped *types.BurnQuote
}

// AssetAmountsOutIterator returns an iterator for iterating output asset amounts
func (b *BurnQuote) AssetAmountsOutIterator() *AssetAmountIterator {
	var aa []*AssetAmount
	for _, v := range b.wrapped.AmountsOut {
		v := v
		aa = append(aa, wrapAssetAmount(&v))
	}

	return &AssetAmountIterator{
		curr:   0,
		values: aa,
	}
}

// AssetAmountsOutWithSlippageIterator returns an iterator for iterating out asset amounts after applying the slippage
func (b *BurnQuote) AssetAmountsOutWithSlippageIterator() (*AssetAmountIterator, error) {
	res, err := b.wrapped.AmountsOutWithSlippage()
	if err != nil {
		return nil, err
	}

	var aa []*AssetAmount
	for _, v := range res {
		v := v
		aa = append(aa, wrapAssetAmount(&v))
	}

	return &AssetAmountIterator{
		curr:   0,
		values: aa,
	}, nil
}

// LiquidityAssetAmount returns a liquidity asset amount
func (b *BurnQuote) LiquidityAssetAmount() *AssetAmount {
	return wrapAssetAmount(&b.wrapped.LiquidityAssetAmount)
}

// Slippage returns a slippage
func (b *BurnQuote) Slippage() string {
	return strconv.FormatFloat(b.wrapped.Slippage, 'f', -1, 64)
}
