package tinyman

import "github.com/synycboom/tinyman-go-sdk/types"

// RedeemQuoteIterator is an iterator for iterating redeem quotes
type RedeemQuoteIterator struct {
	curr   int
	values []*RedeemQuote
}

// HasNext return true if there are items to be iterated
func (i *RedeemQuoteIterator) HasNext() bool {
	return i.curr < len(i.values)
}

// Next returns the next a item, returns nil if no item left
func (i *RedeemQuoteIterator) Next() *RedeemQuote {
	if i.HasNext() {
		idx := i.curr
		i.curr += 1

		return i.values[idx]
	}

	return nil
}

// Reset resets the iterator
func (i *RedeemQuoteIterator) Reset() {
	i.curr = 0
}

// Add adds an item to the iterator
func (i *RedeemQuoteIterator) Add(item *RedeemQuote) {
	i.values = append(i.values, item)
}

// RedeemQuote represents a redeem quote
type RedeemQuote struct {
	wrapped *types.RedeemQuote
}

// GetAssetAmount returns an asset amount
func (r *RedeemQuote) GetAssetAmount() *AssetAmount {
	return wrapAssetAmount(&r.wrapped.Amount)
}

// GetPoolAddress returns an address of the pool
func (r *RedeemQuote) GetPoolAddress() string {
	return r.wrapped.PoolAddress
}
