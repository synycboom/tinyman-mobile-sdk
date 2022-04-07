package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// AmountsIterator is an iterator for iterating amounts
type AmountsIterator struct {
	curr   int
	values []*AssetAmount
}

// HasNext return true if there are asset amounts to be iterated
func (a *AmountsIterator) HasNext() bool {
	return a.curr < len(a.values)
}

// Next returns the next a asset amount, returns nil if no asset amounts left
func (a *AmountsIterator) Next() *AssetAmount {
	if a.HasNext() {
		idx := a.curr
		a.curr += 1

		return a.values[idx]
	}

	return nil
}

// Reset resets the iterator
func (a *AmountsIterator) Reset() {
	a.curr = 0
}

// BurnQuote represents a burn quote
type BurnQuote struct {
	wrappedBurnQuote *types.BurnQuote
}

// AmountsIterator returns an iterator for iterating out asset amounts
func (b *BurnQuote) AmountsIterator() *AmountsIterator {
	var aa []*AssetAmount
	for _, v := range b.wrappedBurnQuote.AmountsOut {
		aa = append(aa, unwrapAssetAmount(&v))
	}

	return &AmountsIterator{
		curr:   0,
		values: aa,
	}
}

// AmountsOutWithSlippageIterator returns an iterator for iterating out asset amounts after applying the slippage
func (b *BurnQuote) AmountsOutWithSlippageIterator() (*AmountsIterator, error) {
	res, err := b.wrappedBurnQuote.AmountsOutWithSlippage()
	if err != nil {
		return nil, err
	}

	var aa []*AssetAmount
	for _, v := range res {
		aa = append(aa, unwrapAssetAmount(&v))
	}

	return &AmountsIterator{
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
