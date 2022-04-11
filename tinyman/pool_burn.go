package tinyman

import (
	"context"
	"fmt"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// PrepareBurnTransactions prepares burn transaction and returns a transaction group
func (p *Pool) PrepareBurnTransactions(
	assetsOutIter *AssetAmountIterator,
	liquidityAssetAmount *AssetAmount,
	burnerAddress string,
) (*TransactionGroup, error) {
	if assetsOutIter == nil {
		return nil, fmt.Errorf("assetsOutIter is required")
	}
	if liquidityAssetAmount == nil {
		return nil, fmt.Errorf("liquidityAssetAmount is required")
	}

	var assetsOut map[uint64]types.AssetAmount
	for assetsOutIter.HasNext() {
		item := assetsOutIter.Next()
		assetsOut[item.wrapped.Asset.ID] = *item.wrapped
	}

	txGroup, err := p.wrapped.PrepareBurnTransactions(context.Background(), assetsOut, liquidityAssetAmount.wrapped, burnerAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}

// PrepareBurnTransactionsFromQuote prepares burn transaction from a given burn quote and returns a transaction group
func (p *Pool) PrepareBurnTransactionsFromQuote(quote *BurnQuote, burnerAddress string) (*TransactionGroup, error) {
	if quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	txGroup, err := p.wrapped.PrepareBurnTransactionsFromQuote(context.Background(), quote.wrapped, burnerAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
