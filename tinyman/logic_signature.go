package tinyman

import (
	algoTypes "github.com/algorand/go-algorand-sdk/types"
)

// LogicSigArgsIterator is a logic signature argument iterator
type LogicSigArgsIterator struct {
	curr   int
	values [][]byte
}

// HasNext return true if there are items to be iterated
func (l *LogicSigArgsIterator) HasNext() bool {
	return l.curr < len(l.values)
}

// Next returns the next item
func (l *LogicSigArgsIterator) Next() []byte {
	if l.HasNext() {
		idx := l.curr
		l.curr += 1

		return l.values[idx]
	}

	return nil
}

// Reset resets the iterator
func (l *LogicSigArgsIterator) Reset() {
	l.curr = 0
}

// Add adds an item to the iterator
func (l *LogicSigArgsIterator) Add(item []byte) {
	l.values = append(l.values, item)
}

// LogicSig contains logic for validating a transaction.
// LogicSig is signed by an account, allowing delegation of operations.
// OR
// LogicSig defines a contract account.
type LogicSig struct {
	wrapped *algoTypes.LogicSig
}

// IsBlank returns true iff the lsig is empty. We need this instead of just
// comparing with == LogicSig{}, because it contains slices.
func (lsig *LogicSig) IsBlank() bool {
	return lsig.wrapped.Blank()
}

// GetLogic returns Logic signed by GetSig or GetMsig
// OR hashed to be the GetAddress of an account.
func (lsig *LogicSig) GetLogic() []byte {
	return lsig.wrapped.Logic
}

// GetSig returns the signature of the account that has delegated to this LogicSig, if any
func (lsig *LogicSig) GetSig() *Signature {
	return wrapSig(&lsig.wrapped.Sig)
}

// GetMsig returns the signature of the multisig account that has delegated to this LogicSig, if any
func (lsig *LogicSig) GetMsig() *MultisigSig {
	return wrapMultisigSigsig(&lsig.wrapped.Msig)
}

// GetArgs returns LogicSigArgsIterator
func (lsig *LogicSig) GetArgs() *LogicSigArgsIterator {
	var values [][]byte
	for _, value := range lsig.wrapped.Args {
		value := value
		values = append(values, value)
	}

	return &LogicSigArgsIterator{
		curr:   0,
		values: values,
	}
}

// SetLogic sets a logic
func (lsig *LogicSig) SetLogic(value []byte) {
	lsig.wrapped.Logic = value
}

// SetSig sets the signature
func (lsig *LogicSig) SetSig(sig *Signature) {
	lsig.wrapped.Sig = *sig.wrapped
}

// SetMsig sets the signature of the multisig account
func (lsig *LogicSig) SetMsig(msig *MultisigSig) {
	lsig.wrapped.Msig = *msig.wrapped
}

// SetArgs sets args from LogicSigArgsIterator
func (lsig *LogicSig) SetArgs(iter *LogicSigArgsIterator) {
	var values [][]byte
	for iter.HasNext() {
		item := iter.Next()
		values = append(values, item)
	}

	lsig.wrapped.Args = values
}

func wrapLogicSig(lsig *algoTypes.LogicSig) *LogicSig {
	return &LogicSig{wrapped: lsig}
}
