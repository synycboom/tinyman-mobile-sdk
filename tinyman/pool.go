package tinyman

import (
	"context"
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/v1/pools"
)

// Pool represents a liquidity pool
type Pool struct {
	wrapped *pools.Pool
}

// NewPool initializes a new liquidity pool
// validatorAppID will be converted to uint64
func NewPool(
	ac *AlgodClient,
	assetA,
	assetB *Asset,
	info *PoolInfo,
	validatorAppID string,
	userAddress string,
	fetch bool,
) (*Pool, error) {
	uintValidatorAppID, err := strconv.ParseUint(validatorAppID, 10, 64)
	if err != nil {
		return nil, err
	}

	wrapped, err := pools.NewPool(
		context.Background(),
		ac.wrapped,
		assetA.wrapped,
		assetB.wrapped,
		info.wrapped,
		uintValidatorAppID,
		userAddress,
		fetch,
	)
	if err != nil {
		return nil, err
	}

	return &Pool{
		wrapped: wrapped,
	}, nil
}

// FromAccountInfo create a pool from an account information
func FromAccountInfo(account *AccountInformation, ac *AlgodClient, userAddress string) (*Pool, error) {
	if account == nil {
		return nil, fmt.Errorf("account is required")
	}

	if ac == nil {
		return nil, fmt.Errorf("ac is required")
	}

	p, err := pools.FromAccountInfo(context.Background(), *account.wrapped, ac.wrapped, userAddress)
	if err != nil {
		return nil, err
	}

	return &Pool{wrapped: p}, nil
}

// Refresh refreshes pool information
func (p *Pool) Refresh(info *PoolInfo) error {
	return p.wrapped.Refresh(context.Background(), info.wrapped)
}

// UpdateFromInfo updates pool information from a given pool info
func (p *Pool) UpdateFromInfo(info *PoolInfo) error {
	return p.wrapped.UpdateFromInfo(context.Background(), info.wrapped)
}

// MinimumBalance calculates minimum balance
// The returned value is in uint64 formatted string
func (p *Pool) MinimumBalance() string {
	value := p.wrapped.MinimumBalance()

	return strconv.FormatUint(value, 10)
}

// LogicSig returns a logic signature account
func (p *Pool) LogicSig() (*LogicSigAccount, error) {
	return poolLogicSigAccount(p.wrapped.ValidatorAppID, p.wrapped.Asset1.ID, p.wrapped.Asset2.ID)
}

// Address returns a logic signature address (pool address)
func (p *Pool) Address() (string, error) {
	return p.wrapped.Address()
}

// Asset1Price returns asset1 price in string-formatted float64
func (p *Pool) Asset1Price() string {
	return strconv.FormatFloat(p.wrapped.Asset1Price(), 'f', -1, 64)
}

// Asset2Price returns asset2 price in string-formatted float64
func (p *Pool) Asset2Price() string {
	return strconv.FormatFloat(p.wrapped.Asset2Price(), 'f', -1, 64)
}

// Info returns pool information
func (p *Pool) Info() (*PoolInfo, error) {
	info, err := p.wrapped.Info()
	if err != nil {
		return nil, err
	}

	return &PoolInfo{wrapped: info}, nil
}

// Convert converts one asset amount to another
func (p *Pool) Convert(amount *AssetAmount) (*AssetAmount, error) {
	a, err := p.wrapped.Convert(amount.wrapped)
	if err != nil {
		return nil, err
	}

	return &AssetAmount{wrapped: a}, nil
}

// FetchStateInt returns an application state int value of the pool by a given key
// The returned value is converted to string-formatted uint64
func (p *Pool) FetchStateInt(key string) (string, error) {
	val, err := p.wrapped.FetchStateInt(context.Background(), key)
	if err != nil {
		return "", err
	}

	return strconv.FormatUint(val, 10), nil
}

// FetchStateBytes returns an application state bytes value of the pool by a given key
// The returned value is converted to string-formatted uint64
func (p *Pool) FetchStateBytes(key string) ([]byte, error) {
	return p.wrapped.FetchStateBytes(context.Background(), key)
}
