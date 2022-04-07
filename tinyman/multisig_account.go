package tinyman

import (
	"fmt"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/types"
	algoTypes "github.com/algorand/go-algorand-sdk/types"
	"golang.org/x/crypto/ed25519"
)

// PublicKeyIterator is a public key iterator
type PublicKeyIterator struct {
	curr   int
	values [][]byte
}

// HasNext return true if there are items to be iterated
func (p *PublicKeyIterator) HasNext() bool {
	return p.curr < len(p.values)
}

// Next returns the next item
func (p *PublicKeyIterator) Next() []byte {
	if p.HasNext() {
		idx := p.curr
		p.curr += 1

		return p.values[idx]
	}

	return nil
}

// Reset resets the iterator
func (p *PublicKeyIterator) Reset() {
	p.curr = 0
}

// Add adds an item to the iterator
func (p *PublicKeyIterator) Add(item []byte) {
	p.values = append(p.values, item)
}

// MultisigAccount is a convenience type for holding multisig preimage data
type MultisigAccount struct {
	wrapped *crypto.MultisigAccount
}

// MultisigAccountWithParams creates a MultisigAccount with the given parameters
func MultisigAccountWithParams(version int, threshold int, addrs []string) (*MultisigAccount, error) {
	var addresses []types.Address
	for _, addr := range addrs {
		decoded, err := algoTypes.DecodeAddress(addr)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, decoded)
	}

	acc, err := crypto.MultisigAccountWithParams(uint8(version), uint8(threshold), addresses)
	if err != nil {
		return nil, err
	}

	return &MultisigAccount{wrapped: &acc}, nil
}

// MultisigAccountFromSig is a convenience method that creates an account
// from a sig in a signed tx. Useful for getting addresses from signed msig txs, etc.
func MultisigAccountFromSig(sig *MultisigSig) (*MultisigAccount, error) {
	if sig == nil {
		return nil, fmt.Errorf("sig is required")
	}

	acc, err := crypto.MultisigAccountFromSig(*sig.wrapped)
	if err != nil {
		return nil, err
	}

	return &MultisigAccount{wrapped: &acc}, nil
}

// Address takes this multisig preimage data, and generates the corresponding identifying
// address, committing to the exact group, version, and public keys that it requires to sign.
// Hash("MultisigAddr" || version uint8 || threshold uint8 || PK1 || PK2 || ...)
func (ma *MultisigAccount) Address() (string, error) {
	addr, err := ma.wrapped.Address()
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}

// Validate ensures that this multisig setup is a valid multisig account
func (ma *MultisigAccount) Validate() (err error) {
	return ma.wrapped.Validate()
}

// Blank return true if MultisigAccount is empty
// struct containing []ed25519.PublicKey cannot be compared
func (ma *MultisigAccount) Blank() bool {
	return ma.wrapped.Blank()
}

// Version return a version which is the version of this multisig
func (ma *MultisigAccount) Version() int {
	return int(ma.wrapped.Version)
}

// Threshold returns a threshold which is how many signatures are needed to fully sign as this address
func (ma *MultisigAccount) Threshold() int {
	return int(ma.wrapped.Threshold)
}

// Pks returns a public key iterator which is an ordered list of public keys that could potentially sign a message
func (ma *MultisigAccount) Pks() *PublicKeyIterator {
	var values [][]byte
	for _, pk := range ma.wrapped.Pks {
		values = append(values, pk)
	}

	return &PublicKeyIterator{curr: 0, values: values}
}

// SetVersion sets a version
func (ma *MultisigAccount) SetVersion(value int) {
	ma.wrapped.Version = uint8(value)
}

// SetThreshold set a threshold
func (ma *MultisigAccount) SetThreshold(value int) {
	ma.wrapped.Threshold = uint8(value)
}

// SetPks set public keys by PublicKeyIterator
func (ma *MultisigAccount) SetPks(iter *PublicKeyIterator) {
	var values []ed25519.PublicKey
	for iter.HasNext() {
		item := iter.Next()
		values = append(values, item)
	}

	ma.wrapped.Pks = values
}
