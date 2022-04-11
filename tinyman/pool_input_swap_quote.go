package tinyman

import (
	"context"
	"fmt"
	"strconv"
)

// FetchFixedInputSwapQuote returns a fixed input swap quote
// slippage is converted to float64
func (p *Pool) FetchFixedInputSwapQuote(amountIn *AssetAmount, slippage string) (*SwapQuote, error) {
	if amountIn == nil {
		return nil, fmt.Errorf("amountIn is required")
	}

	floatSlippage, err := strconv.ParseFloat(slippage, 64)
	if err != nil {
		return nil, err
	}

	quote, err := p.wrapped.FetchFixedInputSwapQuote(context.Background(), amountIn.wrapped, floatSlippage)
	if err != nil {
		return nil, err
	}

	return &SwapQuote{wrapped: quote}, nil
}
