package tinyman

import (
	"context"
	"fmt"
)

// PrepareRedeemTransactions prepares redeem transaction and returns a transaction group
func (p *Pool) PrepareRedeemTransactions(amountOut *AssetAmount, redeemerAddress string) (*TransactionGroup, error) {
	if amountOut == nil {
		return nil, fmt.Errorf("amountOut is required")
	}

	txGroup, err := p.wrapped.PrepareRedeemTransactions(context.Background(), *amountOut.wrapped, redeemerAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
