package tinyman

import (
	"encoding/base64"
	"fmt"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

// Account represents an account
type Account struct {
	wrapped *crypto.Account
}

// NewAccount creates an account from a given type either a base64 encoded private key or a mnemonic
func NewAccount(from string, value string) (*Account, error) {
	var privateKey []byte
	if from == AccountFromPrivateKey {
		pk, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return nil, err
		}

		privateKey = pk
	} else if from == AccountFromMnemonic {
		pk, err := mnemonic.ToPrivateKey(value)
		if err != nil {
			return nil, err
		}

		privateKey = pk
	} else {
		return nil, fmt.Errorf("wrong account creation type %s", from)
	}

	account, err := crypto.AccountFromPrivateKey(privateKey)
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
