package tinyman

import (
	"fmt"

	algoTypes "github.com/algorand/go-algorand-sdk/types"
)

// MultisigSubsigIterator is an iterator
type MultisigSubsigIterator struct {
	curr   int
	values []*MultisigSubsig
}

// HasNext return true if there are items to be iterated
func (m *MultisigSubsigIterator) HasNext() bool {
	return m.curr < len(m.values)
}

// Next returns the next item
func (m *MultisigSubsigIterator) Next() *MultisigSubsig {
	if m.HasNext() {
		idx := m.curr
		m.curr += 1

		return m.values[idx]
	}

	return nil
}

// Reset resets the iterator
func (m *MultisigSubsigIterator) Reset() {
	m.curr = 0
}

// Add adds an item to the iterator
func (m *MultisigSubsigIterator) Add(item *MultisigSubsig) {
	m.values = append(m.values, item)
}

// MultisigSubsig contains a single public key and, optionally, a signature
type MultisigSubsig struct {
	wrapped *algoTypes.MultisigSubsig
}

// SetKey a key
func (m *MultisigSubsig) SetKey(value []byte) {
	m.wrapped.Key = value
}

// SetSignature sets a signature
func (m *MultisigSubsig) SetSignature(sig *Signature) {
	m.wrapped.Sig = *sig.wrapped
}

// Key returns a key
func (m *MultisigSubsig) Key() []byte {
	return m.wrapped.Key
}

// Signature returns a signature
func (m *MultisigSubsig) Signature() *Signature {
	return wrapSig(&m.wrapped.Sig)
}

// MultisigSig holds multiple Subsigs, as well as threshold and version info
type MultisigSig struct {
	wrapped *algoTypes.MultisigSig
}

// Blank returns true iff the msig is empty. We need this instead of just
// comparing with == MultisigSig{}, because Subsigs is a slice.
func (msig *MultisigSig) Blank() bool {
	return msig.wrapped.Blank()
}

// Version returns the version which is converted from uint8
func (msig *MultisigSig) Version() int {
	return int(msig.wrapped.Version)
}

// Threshold returns the threshold which is converted from uint8
func (msig *MultisigSig) Threshold() int {
	return int(msig.wrapped.Threshold)
}

// Subsigs returns a new MultisigSubsig iterator
func (msig *MultisigSig) Subsigs() *MultisigSubsigIterator {
	var ss []*MultisigSubsig
	for _, s := range msig.wrapped.Subsigs {
		s := s
		ss = append(ss, wrapMultisigSubsig(&s))
	}

	return &MultisigSubsigIterator{
		curr:   0,
		values: ss,
	}
}

// SetVersion sets the varsion
func (msig *MultisigSig) SetVersion(value int) {
	msig.wrapped.Version = uint8(value)
}

// SetThreshold sets the varsion
func (msig *MultisigSig) SetThreshold(value int) {
	msig.wrapped.Threshold = uint8(value)
}

// SetSubsigs sets Subsigs from MultisigSubsigIterator
func (msig *MultisigSig) SetSubsigs(iter *MultisigSubsigIterator) error {
	if iter == nil {
		return fmt.Errorf("iter is required")
	}

	var subsigs []algoTypes.MultisigSubsig
	for iter.HasNext() {
		item := iter.Next()
		subsigs = append(subsigs, *item.wrapped)
	}

	msig.wrapped.Subsigs = subsigs

	return nil
}

func wrapMultisigSubsig(s *algoTypes.MultisigSubsig) *MultisigSubsig {
	return &MultisigSubsig{wrapped: s}
}

func wrapMultisigSigsig(s *algoTypes.MultisigSig) *MultisigSig {
	return &MultisigSig{wrapped: s}
}
