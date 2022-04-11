package tinyman

import (
	"context"
	"strconv"
)

// FetchBurnQuote returns a burn quote
// slippage must be a string-formatted float64
func (p *Pool) FetchBurnQuote(liquidityAsset *AssetAmount, slippage string) (*BurnQuote, error) {
	floatSlippage, err := strconv.ParseFloat(slippage, 64)
	if err != nil {
		return nil, err
	}

	quote, err := p.wrapped.FetchBurnQuote(context.Background(), *liquidityAsset.wrapped, floatSlippage)
	if err != nil {
		return nil, err
	}

	return &BurnQuote{wrapped: quote}, nil
}
