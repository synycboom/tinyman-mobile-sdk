package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// Reset resets the iterator
func (a *AssetAmountIterator) Reset() {
	a.curr = 0
}

// BurnQuote represents a burn quote
type BurnQuote struct {
	wrappedBurnQuote *types.BurnQuote
}

// AssetAmountsOutIterator returns an iterator for iterating output asset amounts
func (b *BurnQuote) AssetAmountsOutIterator() *AssetAmountIterator {
	var aa []*AssetAmount
	for _, v := range b.wrappedBurnQuote.AmountsOut {
		aa = append(aa, unwrapAssetAmount(&v))
	}

	return &AssetAmountIterator{
		curr:   0,
		values: aa,
	}
}

// AssetAmountsOutWithSlippageIterator returns an iterator for iterating out asset amounts after applying the slippage
func (b *BurnQuote) AssetAmountsOutWithSlippageIterator() (*AssetAmountIterator, error) {
	res, err := b.wrappedBurnQuote.AmountsOutWithSlippage()
	if err != nil {
		return nil, err
	}

	var aa []*AssetAmount
	for _, v := range res {
		aa = append(aa, unwrapAssetAmount(&v))
	}

	return &AssetAmountIterator{
		curr:   0,
		values: aa,
	}, nil
}

// LiquidityAssetAmount returns a liquidity asset amount
func (b *BurnQuote) LiquidityAssetAmount() *AssetAmount {
	return unwrapAssetAmount(&b.wrappedBurnQuote.LiquidityAssetAmount)
}

// Slippage returns a slippage
func (b *BurnQuote) Slippage() string {
	return strconv.FormatFloat(b.wrappedBurnQuote.Slippage, 'f', -1, 64)
}
