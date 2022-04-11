package tinyman

import "context"

// PrepareBootstrapTransactions prepares bootstrap transaction and returns a transaction group
func (p *Pool) PrepareBootstrapTransactions(bootstrapperAddress string) (*TransactionGroup, error) {
	txGroup, err := p.wrapped.PrepareBootstrapTransactions(context.Background(), bootstrapperAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}
