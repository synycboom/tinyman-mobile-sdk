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

// GetKey returns a key
func (m *MultisigSubsig) GetKey() []byte {
	return m.wrapped.Key
}

// GetSignature returns a signature
func (m *MultisigSubsig) GetSignature() *Signature {
	return wrapSig(&m.wrapped.Sig)
}

// MultisigSig holds multiple GetSubsigs, as well as threshold and version info
type MultisigSig struct {
	wrapped *algoTypes.MultisigSig
}

// IsBlank returns true iff the msig is empty. We need this instead of just
// comparing with == MultisigSig{}, because GetSubsigs is a slice.
func (msig *MultisigSig) IsBlank() bool {
	return msig.wrapped.Blank()
}

// GetVersion returns the version which is converted from uint8
func (msig *MultisigSig) GetVersion() int {
	return int(msig.wrapped.Version)
}

// GetThreshold returns the threshold which is converted from uint8
func (msig *MultisigSig) GetThreshold() int {
	return int(msig.wrapped.Threshold)
}

// GetSubsigs returns a new MultisigSubsig iterator
func (msig *MultisigSig) GetSubsigs() *MultisigSubsigIterator {
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

// SetSubsigs sets GetSubsigs from MultisigSubsigIterator
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
