package tinyman

import "context"

// PrepareLiquidityAssetOptInTransactions prepares liquidity asset opt-in transaction and returns a transaction group
func (p *Pool) PrepareLiquidityAssetOptInTransactions(userAddress string) (*TransactionGroup, error) {
	txGroup, err := p.wrapped.PrepareLiquidityAssetOptInTransactions(context.Background(), userAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
