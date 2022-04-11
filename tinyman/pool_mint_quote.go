package tinyman

import (
	"context"
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// FetchMintQuote returns a mint quote
// slippage is converted to float64
func (p *Pool) FetchMintQuote(
	amountA *AssetAmount,
	amountB *AssetAmount,
	slippage string,
) (*MintQuote, error) {
	if amountA == nil {
		return nil, fmt.Errorf("amountA is required")
	}

	var assetAmountB *types.AssetAmount
	if amountB != nil {
		assetAmountB = amountB.wrapped
	}

	floatSlippage, err := strconv.ParseFloat(slippage, 64)
	if err != nil {
		return nil, err
	}

	quote, err := p.wrapped.FetchMintQuote(context.Background(), amountA.wrapped, assetAmountB, floatSlippage)
	if err != nil {
		return nil, err
	}

	return &MintQuote{wrapped: quote}, nil
}
