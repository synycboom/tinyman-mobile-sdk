package tinyman

import (
	"fmt"

	algoTypes "github.com/algorand/go-algorand-sdk/types"
)

// Signature is an ed25519 signature
type Signature struct {
	wrapped *algoTypes.Signature
}

// Set sets a new signature value
func (s *Signature) Set(value []byte) error {
	var newValue algoTypes.Signature
	if len(value) != len(newValue) {
		return fmt.Errorf("invalid signature length")
	}

	for idx := 0; idx < len(newValue); idx++ {
		newValue[idx] = value[idx]
	}

	s.wrapped = &newValue

	return nil
}

// Get returns a signature value
func (s *Signature) Get() []byte {
	var out []byte
	for _, b := range s.wrapped {
		out = append(out, b)
	}

	return out
}

func wrapSig(s *algoTypes.Signature) *Signature {
	return &Signature{wrapped: s}
}
