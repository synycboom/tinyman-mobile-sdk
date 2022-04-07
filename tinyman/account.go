package tinyman

import (
	"encoding/base64"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

// Account represents an account
type Account struct {
	wrapped *crypto.Account
}

// NewAccountFromPrivateKey creates an account from a base64 encoded private key
func NewAccountFromPrivateKey(value string) (*Account, error) {
	pk, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}

	account, err := crypto.AccountFromPrivateKey(pk)
	if err != nil {
		return nil, err
	}

	return &Account{wrapped: &account}, nil
}

// NewAccountFromMnemonic creates an account from a mnemonic
func NewAccountFromMnemonic(value string) (*Account, error) {
	pk, err := mnemonic.ToPrivateKey(value)
	if err != nil {
		return nil, err
	}

	account, err := crypto.AccountFromPrivateKey(pk)
	if err != nil {
		return nil, err
	}

	return &Account{wrapped: &account}, nil
}

// Address returns an address
func (a *Account) Address() string {
	return a.wrapped.Address.String()
}

// PrivateKey returns a private key
func (a *Account) PrivateKey() []byte {
	return a.wrapped.PrivateKey
}

// PublicKey returns a public key
func (a *Account) PublicKey() []byte {
	return a.wrapped.PublicKey
}

func unwrapAccount(acc *Account) (*crypto.Account, error) {
	unwrapped, err := crypto.AccountFromPrivateKey(acc.PrivateKey())
	if err != nil {
		return nil, err
	}

	return &unwrapped, nil
}
