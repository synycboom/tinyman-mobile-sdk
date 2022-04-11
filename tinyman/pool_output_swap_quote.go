package tinyman

import (
	"context"
	"fmt"
	"strconv"
)

// FetchFixedOutputSwapQuote returns a fixed input swap quote
// slippage is converted to float64
func (p *Pool) FetchFixedOutputSwapQuote(amountOut *AssetAmount, slippage string) (*SwapQuote, error) {
	if amountOut == nil {
		return nil, fmt.Errorf("amountOut is required")
	}

	floatSlippage, err := strconv.ParseFloat(slippage, 64)
	if err != nil {
		return nil, err
	}

	quote, err := p.wrapped.FetchFixedOutputSwapQuote(context.Background(), *amountOut.wrapped, floatSlippage)
	if err != nil {
		return nil, err
	}

	return &SwapQuote{wrapped: quote}, nil
}
