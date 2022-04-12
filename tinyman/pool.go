package tinyman

import (
	"context"
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
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
	if ac == nil {
		return nil, fmt.Errorf("algodClient is required")
	}
	if assetA == nil || assetB == nil {
		return nil, fmt.Errorf("both assetA and assetB are required")
	}

	uintValidatorAppID, err := strconv.ParseUint(validatorAppID, 10, 64)
	if err != nil {
		return nil, err
	}

	var poolInfo *types.PoolInfo
	if info != nil {
		poolInfo = info.wrapped
	}

	wrapped, err := pools.NewPool(
		context.Background(),
		ac.wrapped,
		assetA.wrapped,
		assetB.wrapped,
		poolInfo,
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

// GetMinimumBalance calculates minimum balance
// The returned value is in uint64 formatted string
func (p *Pool) GetMinimumBalance() string {
	value := p.wrapped.MinimumBalance()

	return strconv.FormatUint(value, 10)
}

// GetLogicSig returns a logic signature account
func (p *Pool) GetLogicSig() (*LogicSigAccount, error) {
	return poolLogicSigAccount(p.wrapped.ValidatorAppID, p.wrapped.Asset1.ID, p.wrapped.Asset2.ID)
}

// GetLiquidityAsset returns a liquidity asset
func (p *Pool) GetLiquidityAsset() *Asset {
	return wrapAsset(p.wrapped.LiquidityAsset)
}

// GetAsset1 returns an asset 1
func (p *Pool) GetAsset1() *Asset {
	return wrapAsset(p.wrapped.Asset1)
}

// GetAsset2 returns an asset 2
func (p *Pool) GetAsset2() *Asset {
	return wrapAsset(p.wrapped.Asset2)
}

// GetAddress returns a logic signature address (pool address)
func (p *Pool) GetAddress() (string, error) {
	return p.wrapped.Address()
}

// GetAsset1Price returns asset1 price in string-formatted float64
func (p *Pool) GetAsset1Price() string {
	return strconv.FormatFloat(p.wrapped.Asset1Price(), 'f', -1, 64)
}

// GetAsset2Price returns asset2 price in string-formatted float64
func (p *Pool) GetAsset2Price() string {
	return strconv.FormatFloat(p.wrapped.Asset2Price(), 'f', -1, 64)
}

// GetInfo returns pool information
func (p *Pool) GetInfo() (*PoolInfo, error) {
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
