package tinyman

import "github.com/synycboom/tinyman-go-sdk/types"

// RedeemQuote represents a redeem quote
type RedeemQuote struct {
	wrapped *types.RedeemQuote
}

// AssetAmount returns an asset amount
func (r *RedeemQuote) AssetAmount() *AssetAmount {
	return unwrapAssetAmount(&r.wrapped.Amount)
}

// AssetAmount returns an asset amount
func (r *RedeemQuote) PoolAddress() string {
	return r.wrapped.PoolAddress
}
