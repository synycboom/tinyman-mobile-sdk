package tinyman

import (
	"fmt"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/types"
	algoTypes "github.com/algorand/go-algorand-sdk/types"
	"golang.org/x/crypto/ed25519"
)

// AddressIterator is an address iterator
type AddressIterator struct {
	curr   int
	values []string
}

// HasNext return true if there are items to be iterated
func (i *AddressIterator) HasNext() bool {
	return i.curr < len(i.values)
}

// Next returns the next item
func (i *AddressIterator) Next() string {
	if i.HasNext() {
		idx := i.curr
		i.curr += 1

		return i.values[idx]
	}

	return ""
}

// Reset resets the iterator
func (i *AddressIterator) Reset() {
	i.curr = 0
}

// Add adds an item to the iterator
func (i *AddressIterator) Add(item string) {
	i.values = append(i.values, item)
}

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

// MultiSigAccount is a convenience type for holding multisig preimage data
type MultiSigAccount struct {
	wrapped *crypto.MultisigAccount
}

// MultiSigAccountWithParams creates a MultiSigAccount with the given parameters
func MultiSigAccountWithParams(version int, threshold int, addressIter *AddressIterator) (*MultiSigAccount, error) {
	if addressIter == nil {
		return nil, fmt.Errorf("addressIter is required")
	}

	var addresses []types.Address
	for addressIter.HasNext() {
		item := addressIter.Next()
		decoded, err := algoTypes.DecodeAddress(item)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, decoded)
	}

	acc, err := crypto.MultisigAccountWithParams(uint8(version), uint8(threshold), addresses)
	if err != nil {
		return nil, err
	}

	return &MultiSigAccount{wrapped: &acc}, nil
}

// MultiSigAccountFromSig is a convenience method that creates an account
// from a sig in a signed tx. Useful for getting addresses from signed msig txs, etc.
func MultiSigAccountFromSig(sig *MultisigSig) (*MultiSigAccount, error) {
	if sig == nil {
		return nil, fmt.Errorf("sig is required")
	}

	acc, err := crypto.MultisigAccountFromSig(*sig.wrapped)
	if err != nil {
		return nil, err
	}

	return &MultiSigAccount{wrapped: &acc}, nil
}

// Address takes this multisig preimage data, and generates the corresponding identifying
// address, committing to the exact group, version, and public keys that it requires to sign.
// Hash("MultiSigAddr" || version uint8 || threshold uint8 || PK1 || PK2 || ...)
func (ma *MultiSigAccount) Address() (string, error) {
	addr, err := ma.wrapped.Address()
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}

// Validate ensures that this multisig setup is a valid multisig account
func (ma *MultiSigAccount) Validate() (err error) {
	return ma.wrapped.Validate()
}

// Blank return true if MultiSigAccount is empty
// struct containing []ed25519.GetPublicKey cannot be compared
func (ma *MultiSigAccount) Blank() bool {
	return ma.wrapped.Blank()
}

// Version return a version which is the version of this multisig
func (ma *MultiSigAccount) Version() int {
	return int(ma.wrapped.Version)
}

// Threshold returns a threshold which is how many signatures are needed to fully sign as this address
func (ma *MultiSigAccount) Threshold() int {
	return int(ma.wrapped.Threshold)
}

// Pks returns a public key iterator which is an ordered list of public keys that could potentially sign a message
func (ma *MultiSigAccount) Pks() *PublicKeyIterator {
	var values [][]byte
	for _, pk := range ma.wrapped.Pks {
		pk := pk
		values = append(values, pk)
	}

	return &PublicKeyIterator{curr: 0, values: values}
}

// SetVersion sets a version
func (ma *MultiSigAccount) SetVersion(value int) {
	ma.wrapped.Version = uint8(value)
}

// SetThreshold set a threshold
func (ma *MultiSigAccount) SetThreshold(value int) {
	ma.wrapped.Threshold = uint8(value)
}

// SetPks set public keys by PublicKeyIterator
func (ma *MultiSigAccount) SetPks(iter *PublicKeyIterator) {
	var values []ed25519.PublicKey
	for iter.HasNext() {
		item := iter.Next()
		values = append(values, item)
	}

	ma.wrapped.Pks = values
}
