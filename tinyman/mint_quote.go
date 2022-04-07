package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// MintQuote represents a mint quote
type MintQuote struct {
	wrappedMintQuote *types.MintQuote
}

// AssetAmountsInIterator returns an iterator for iterating input asset amounts
func (b *MintQuote) AssetAmountsInIterator() *AssetAmountIterator {
	var aa []*AssetAmount
	for _, v := range b.wrappedMintQuote.AmountsIn {
		aa = append(aa, unwrapAssetAmount(&v))
	}

	return &AssetAmountIterator{
		curr:   0,
		values: aa,
	}
}

// LiquidityAssetAmount returns a liquidity asset amount
func (m *MintQuote) LiquidityAssetAmount() *AssetAmount {
	return unwrapAssetAmount(&m.wrappedMintQuote.LiquidityAssetAmount)
}

// LiquidityAssetAmountWithSlippage calculates liquidity asset after applying the slippage
func (m *MintQuote) LiquidityAssetAmountWithSlippage() (*AssetAmount, error) {
	assetAmount, err := m.wrappedMintQuote.LiquidityAssetAmountWithSlippage()
	if err != nil {
		return nil, err
	}

	return unwrapAssetAmount(assetAmount), nil
}

// Slippage returns a slippage
func (b *MintQuote) Slippage() string {
	return strconv.FormatFloat(b.wrappedMintQuote.Slippage, 'f', -1, 64)
}
