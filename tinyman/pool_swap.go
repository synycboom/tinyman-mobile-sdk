package tinyman

import (
	"context"
	"fmt"
)

// PrepareSwapTransactions prepares swap transaction and returns a transaction group
func (p *Pool) PrepareSwapTransactions(
	assetAmountIn,
	assetAmountOut *AssetAmount,
	swapType,
	swapperAddress string,
) (*TransactionGroup, error) {
	txGroup, err := p.wrapped.PrepareSwapTransactions(context.Background(), assetAmountIn.wrapped, assetAmountOut.wrapped, swapType, swapperAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}

// PrepareSwapTransactionsFromQuote prepares swap transaction from a given swap quote and returns a transaction group
func (p *Pool) PrepareSwapTransactionsFromQuote(quote *SwapQuote, swapperAddress string) (*TransactionGroup, error) {
	if quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	txGroup, err := p.wrapped.PrepareSwapTransactionsFromQuote(context.Background(), quote.wrapped, swapperAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
