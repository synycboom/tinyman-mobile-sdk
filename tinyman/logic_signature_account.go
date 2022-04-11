package tinyman

import (
	"crypto/ed25519"
	"fmt"

	"github.com/algorand/go-algorand-sdk/crypto"
)

// LogicSigAddress returns the contract (escrow) address for a LogicSig.
//
// NOTE: If the LogicSig is delegated to another account this will not
// return the delegated address of the LogicSig.
func LogicSigAddress(lsig *LogicSig) (string, error) {
	if lsig == nil {
		return "", fmt.Errorf("lsig is required")
	}

	return crypto.LogicSigAddress(*lsig.wrapped).String(), nil
}

// LogicSigAccount represents an account that can sign with a LogicSig program.
type LogicSigAccount struct {
	wrapped *crypto.LogicSigAccount
}

// MakeLogicSigAccountEscrow creates a new escrow LogicSigAccount. The address
// of this account will be a hash of its program.
func MakeLogicSigAccountEscrow(program []byte, iter *LogicSigArgsIterator) (*LogicSigAccount, error) {
	var args [][]byte
	if iter != nil {
		for iter.HasNext() {
			item := iter.Next()
			args = append(args, item)
		}
	}

	acc := crypto.MakeLogicSigAccountEscrow(program, args)

	return &LogicSigAccount{
		wrapped: &acc,
	}, nil
}

// MakeLogicSigAccountDelegated creates a new delegated LogicSigAccount. This
// type of LogicSig has the authority to sign transactions on behalf of another
// account, called the delegating account. If the delegating account is a
// multisig account, use MakeLogicSigAccountDelegated instead.
//
// The parameter signer is the private key of the delegating account.
func MakeLogicSigAccountDelegated(program []byte, iter *LogicSigArgsIterator, signer []byte) (*LogicSigAccount, error) {
	acc, err := crypto.MakeLogicSigAccountDelegated(program, iter.values, signer)
	if err != nil {
		return nil, err
	}

	return &LogicSigAccount{wrapped: &acc}, nil
}

// MakeLogicSigAccountDelegatedMsig creates a new delegated LogicSigAccount.
// This type of LogicSig has the authority to sign transactions on behalf of
// another account, called the delegating account. Use this function if the
// delegating account is a multisig account, otherwise use
// MakeLogicSigAccountDelegated.
//
// The parameter msigAccount is the delegating multisig account.
//
// The parameter signer is the private key of one of the members of the
// delegating multisig account. Use the method AppendMultisigSignature on the
// returned LogicSigAccount to add additional signatures from other members.
func MakeLogicSigAccountDelegatedMsig(program []byte, iter *LogicSigArgsIterator, msigAccount *MultiSigAccount, signer []byte) (*LogicSigAccount, error) {
	if msigAccount == nil {
		return nil, fmt.Errorf("msigAccount is required")
	}

	acc, err := crypto.MakeLogicSigAccountDelegatedMsig(program, iter.values, *msigAccount.wrapped, signer)
	if err != nil {
		return nil, err
	}

	return &LogicSigAccount{wrapped: &acc}, nil
}

// AppendMultisigSignature adds an additional signature from a member of the
// delegating multisig account.
//
// The LogicSigAccount must represent a delegated LogicSig backed by a multisig
// account.
func (lsa *LogicSigAccount) AppendMultisigSignature(signer []byte) error {
	return crypto.AppendMultisigToLogicSig(&lsa.wrapped.Lsig, signer)
}

// LogicSigAccountFromLogicSig creates a LogicSigAccount from an existing
// LogicSig object.
//
// The parameter signerPublicKey must be present if the LogicSig is delegated
// and the delegating account is backed by a single private key (i.e. not a
// multisig account). In this case, signerPublicKey must be the public key of
// the delegating account. In all other cases, an error will be returned if
// signerPublicKey is present.
func LogicSigAccountFromLogicSig(lsig *LogicSig, signerPublicKey []byte) (*LogicSigAccount, error) {
	if lsig == nil {
		return nil, fmt.Errorf("lsig is required")
	}

	var pk ed25519.PublicKey
	for _, value := range signerPublicKey {
		pk = append(pk, value)
	}

	acc, err := crypto.LogicSigAccountFromLogicSig(*lsig.wrapped, &pk)
	if err != nil {
		return nil, err
	}

	return &LogicSigAccount{wrapped: &acc}, nil
}

// IsDelegated returns true if and only if the LogicSig has been delegated to
// another account with a signature.
//
// Note this function only checks for the presence of a delegation signature. To
// verify the delegation signature, use VerifyLogicSig.
func (lsa *LogicSigAccount) IsDelegated() bool {
	return lsa.wrapped.IsDelegated()
}

// Address returns the address of this LogicSigAccount.
//
// If the LogicSig is delegated to another account, this will return the address
// of that account.
//
// If the LogicSig is not delegated to another account, this will return an
// escrow address that is the hash of the LogicSig's program code.
func (lsa *LogicSigAccount) Address() (string, error) {
	addr, err := lsa.wrapped.Address()
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}

// Lsig returns the underlying LogicSig object
func (lsa *LogicSigAccount) Lsig() *LogicSig {
	return wrapLogicSig(&lsa.wrapped.Lsig)
}

// SigningKey returns the key that provided Lsig.Sig, if any
func (lsa *LogicSigAccount) SigningKey() []byte {
	return lsa.wrapped.SigningKey
}

// SetLsig sets the underlying LogicSig object
func (lsa *LogicSigAccount) SetLsig(lsig *LogicSig) {
	lsa.wrapped.Lsig = lsa.wrapped.Lsig
}

// SetSigningKey set the key that provided Lsig.Sig
func (lsa *LogicSigAccount) SetSigningKey(signingKey []byte) {
	lsa.wrapped.SigningKey = signingKey
}
