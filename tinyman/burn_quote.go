package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// BurnQuote represents a burn quote
type BurnQuote struct {
	wrapped *types.BurnQuote
}

// GetAssetAmountsOutIterator returns an iterator for iterating output asset amounts
func (b *BurnQuote) GetAssetAmountsOutIterator() *AssetAmountIterator {
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

// GetAssetAmountsOutWithSlippageIterator returns an iterator for iterating out asset amounts after applying the slippage
func (b *BurnQuote) GetAssetAmountsOutWithSlippageIterator() (*AssetAmountIterator, error) {
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

// GetLiquidityAssetAmount returns a liquidity asset amount
func (b *BurnQuote) GetLiquidityAssetAmount() *AssetAmount {
	return wrapAssetAmount(&b.wrapped.LiquidityAssetAmount)
}

// GetSlippage returns a slippage
func (b *BurnQuote) GetSlippage() string {
	return strconv.FormatFloat(b.wrapped.Slippage, 'f', -1, 64)
}
