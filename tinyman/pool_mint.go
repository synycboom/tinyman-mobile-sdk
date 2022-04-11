package tinyman

import (
	"context"
	"fmt"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// PrepareMintTransactions prepares mint transaction and returns a transaction group
func (p *Pool) PrepareMintTransactions(
	amountsInIter *AssetAmountIterator,
	liquidityAssetAmount *AssetAmount,
	minterAddress string,
) (*TransactionGroup, error) {
	if amountsInIter == nil {
		return nil, fmt.Errorf("amountsInIter is required")
	}
	if liquidityAssetAmount == nil {
		return nil, fmt.Errorf("liquidityAssetAmount is required")
	}

	var amountsIn map[uint64]types.AssetAmount
	for amountsInIter.HasNext() {
		item := amountsInIter.Next()
		amountsIn[item.wrapped.Asset.ID] = *item.wrapped
	}

	txGroup, err := p.wrapped.PrepareMintTransactions(context.Background(), amountsIn, *liquidityAssetAmount.wrapped, minterAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
