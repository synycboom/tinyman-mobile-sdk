package tinyman

import (
	"context"

	"github.com/algorand/go-algorand-sdk/types"
	"github.com/synycboom/tinyman-go-sdk/utils"
)

// Transaction is an Algorand transaction wrapper
type Transaction struct {
	wrapped *types.Transaction
}

// TransactionIterator is a transaction iterator
type TransactionIterator struct {
	curr   int
	values []*Transaction
}

// HasNext return true if there are items to be iterated
func (iter *TransactionIterator) HasNext() bool {
	return iter.curr < len(iter.values)
}

// Next returns the next item
func (iter *TransactionIterator) Next() *Transaction {
	if iter.HasNext() {
		idx := iter.curr
		iter.curr += 1

		return iter.values[idx]
	}

	return nil
}

// Reset resets the iterator
func (iter *TransactionIterator) Reset() {
	iter.curr = 0
}

// TransactionGroup is a group of transaction that can be executed atomically after signing
type TransactionGroup struct {
	wrapped *utils.TransactionGroup
}

// Sign signs a transaction group with an account
func (tg *TransactionGroup) Sign(acc *Account) error {
	unwrappedAcc, err := unwrapAccount(acc)
	if err != nil {
		return err
	}

	return tg.wrapped.Sign(unwrappedAcc)
}

// SignWithLogicSig signs a transaction group with logic sig account
func (tg *TransactionGroup) SignWithLogicSig(account *LogicSigAccount) error {
	return tg.wrapped.SignWithLogicSig(account.wrapped)
}

// Submit sends a signed transaction group to the blockchain
func (tg *TransactionGroup) Submit(client *AlgodClient, wait bool) (string, error) {
	return tg.wrapped.Submit(context.Background(), client.wrapped, wait)
}

// IMPROVEMENT: Add transaction iterator
// // TransactionIterator returns a transaciton iterator
// func (tg *TransactionGroup) TransactionIterator() *TransactionIterator {
// }
