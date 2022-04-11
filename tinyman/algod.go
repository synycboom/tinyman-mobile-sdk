package tinyman

import (
	"context"
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/common"
	"github.com/algorand/go-algorand-sdk/crypto"
)

// Header is the Algorand client header
type Header struct {
	Key   string
	Value string
}

// HeaderIterator is the header iterator
type HeaderIterator struct {
	curr   int
	values []*Header
}

// HasNext return true if there are items to be iterated
func (h *HeaderIterator) HasNext() bool {
	return h.curr < len(h.values)
}

// Next returns the next item
func (h *HeaderIterator) Next() *Header {
	if h.HasNext() {
		idx := h.curr
		h.curr += 1

		return h.values[idx]
	}

	return nil
}

// Reset resets the iterator
func (h *HeaderIterator) Reset() {
	h.curr = 0
}

// Add adds an item to the iterator
func (h *HeaderIterator) Add(item *Header) {
	h.values = append(h.values, item)
}

// AlgodClient is a wrapper for algod client
type AlgodClient struct {
	wrapped *algod.Client
}

// MakeAlgodClient is the factory for constructing the Algorand client for a given endpoint.
func MakeAlgodClient(address string, apiToken string) (*AlgodClient, error) {
	c, err := algod.MakeClient(address, apiToken)
	if err != nil {
		return nil, err
	}

	return &AlgodClient{wrapped: c}, nil
}

// MakeAlgodClientWithHeaders is the factory for constructing the Algorand client for a given endpoint with additional user defined headers.
func MakeAlgodClientWithHeaders(address string, apiToken string, headerIter *HeaderIterator) (*AlgodClient, error) {
	var headers []*common.Header
	for headerIter.HasNext() {
		item := headerIter.Next()
		headers = append(headers, &common.Header{Key: item.Key, Value: item.Value})
	}

	c, err := algod.MakeClientWithHeaders(address, apiToken, headers)
	if err != nil {
		return nil, err
	}

	return &AlgodClient{wrapped: c}, nil
}

// SuggestedParams returns suggested params
func (c *AlgodClient) SuggestedParams() (*SuggestedParams, error) {
	sp, err := c.wrapped.SuggestedParams().Do(context.Background())
	if err != nil {
		return nil, err
	}

	return &SuggestedParams{wrapped: &sp}, nil
}

// AccountInformation fetches account information
func (c *AlgodClient) AccountInformation(address string) (*AccountInformation, error) {
	a, err := c.wrapped.AccountInformation(address).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return &AccountInformation{wrapped: &a}, nil
}

// SendRawTransaction sends a transaction to the blockchain
func (c *AlgodClient) SendRawTransaction(rawTx []byte) error {
	if _, err := c.wrapped.SendRawTransaction(rawTx).Do(context.Background()); err != nil {
		return err
	}

	return nil
}

// SignTransactionWithPrivateKey signs a transaction with a given private key
func SignTransactionWithPrivateKey(privateKey []byte, tx *Transaction) (*SignedTransaction, error) {
	if tx == nil {
		return nil, fmt.Errorf("tx is required")
	}

	txID, stxBytes, err := crypto.SignTransaction(privateKey, *tx.wrapped)
	if err != nil {
		return nil, err
	}

	return &SignedTransaction{
		TxID:     txID,
		StxBytes: stxBytes,
	}, nil
}

// IMPROVEMENT: Implement wrapper methods for algod client
