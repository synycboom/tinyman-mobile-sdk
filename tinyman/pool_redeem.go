package tinyman

import (
	"context"
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// PrepareRedeemTransactions prepares redeem transaction and returns a transaction group
func (p *Pool) PrepareRedeemTransactions(amountOut *AssetAmount, redeemerAddress string) (*TransactionGroup, error) {
	if amountOut == nil {
		return nil, fmt.Errorf("amountOut is required")
	}

	txGroup, err := p.wrapped.PrepareRedeemTransactions(context.Background(), amountOut.wrapped, redeemerAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}

// PrepareRedeemTransactionsFromQuote prepares redeem transactions and return a transaction group from quote
func (p *Pool) PrepareRedeemTransactionsFromQuote(quote *RedeemQuote, redeemerAddress string) (*TransactionGroup, error) {
	if quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	txGroup, err := p.wrapped.PrepareRedeemTransactionsFromQuote(context.Background(), quote.wrapped, redeemerAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}

// FilterRedeemQuotes filters only redeem quote that belonging to this pool
// It returns a new RedeemQuoteIterator
func (p *Pool) FilterRedeemQuotes(iter *RedeemQuoteIterator) (*RedeemQuoteIterator, error) {
	if iter == nil {
		return nil, fmt.Errorf("iter is required")
	}

	var quotes []types.RedeemQuote
	for iter.HasNext() {
		item := iter.Next()
		quotes = append(quotes, *item.wrapped)
	}

	quotes, err := p.wrapped.FilterRedeemQuotes(quotes)
	if err != nil {
		return nil, err
	}

	var wrappedQuotes []*RedeemQuote
	for _, quote := range quotes {
		wrappedQuotes = append(wrappedQuotes, &RedeemQuote{wrapped: &quote})
	}

	return &RedeemQuoteIterator{
		values: wrappedQuotes,
	}, nil
}

// GetRedeemQuoteMatchesAssetID filters redeem quote belonging to this pool which matches an asset id
// assetID is converted to uint64
func (p *Pool) GetRedeemQuoteMatchesAssetID(assetID string, iter *RedeemQuoteIterator) (*RedeemQuote, error) {
	if iter == nil {
		return nil, fmt.Errorf("iter is required")
	}

	uintAssetID, err := strconv.ParseUint(assetID, 10, 64)
	if err != nil {
		return nil, err
	}

	var quotes []types.RedeemQuote
	for iter.HasNext() {
		item := iter.Next()
		quotes = append(quotes, *item.wrapped)
	}

	quote, err := p.wrapped.GetRedeemQuoteMatchesAssetID(uintAssetID, quotes)
	if err != nil {
		return nil, err
	}

	return &RedeemQuote{
		wrapped: quote,
	}, nil
}
