package tinyman

import (
	"context"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
	"github.com/synycboom/tinyman-go-sdk/v1/pools"
)

// PoolInfo represents pool information
type PoolInfo struct {
	wrapped *types.PoolInfo
}

// NewPoolInfo creates a pool information
func NewPoolInfo() *PoolInfo {
	return &PoolInfo{
		wrapped: &types.PoolInfo{},
	}
}

// SetAddress sets an address of the pool
func (p *PoolInfo) SetAddress(value string) {
	p.wrapped.Address = value
}

// SetAsset1ID sets an asset1 id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAsset1ID(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.Asset1ID = pv

	return nil
}

// SetAsset2ID sets an asset2 id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAsset2ID(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.Asset2ID = pv

	return nil
}

// SetAsset1UnitName sets an asset1 unit name
func (p *PoolInfo) SetAsset1UnitName(value string) {
	p.wrapped.Asset1UnitName = value
}

// SetAsset2UnitName sets an asset1 unit name
func (p *PoolInfo) SetAsset2UnitName(value string) {
	p.wrapped.Asset2UnitName = value
}

// SetLiquidityAssetID sets a liquidity asset id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetLiquidityAssetID(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.LiquidityAssetID = pv

	return nil
}

// SetLiquidityAssetName sets a liquidity name
func (p *PoolInfo) SetLiquidityAssetName(value string) {
	p.wrapped.LiquidityAssetName = value
}

// SetAsset1Reserves sets a liquidity asset id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAsset1Reserves(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.Asset1Reserves = pv

	return nil
}

// SetAsset2Reserves sets a liquidity asset id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAsset2Reserves(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.Asset2Reserves = pv

	return nil
}

// SetIssuedLiquidity sets issued liquidity amount, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetIssuedLiquidity(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.IssuedLiquidity = pv

	return nil
}

// SetUnclaimedProtocolFee sets unclaimed protocol fee, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetUnclaimedProtocolFee(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.UnclaimedProtocolFee = pv

	return nil
}

// SetOutstandingAsset1Amount sets outstanding asset1 amount, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetOutstandingAsset1Amount(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.OutstandingAsset1Amount = pv

	return nil
}

// SetOutstandingAsset2Amount sets outstanding asset2 amount, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetOutstandingAsset2Amount(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.OutstandingAsset2Amount = pv

	return nil
}

// SetOutstandingLiquidityAssetAmount sets outstanding liquidity asset amount, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetOutstandingLiquidityAssetAmount(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.OutstandingLiquidityAssetAmount = pv

	return nil
}

// SetValidatorAppID sets the validator app id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetValidatorAppID(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.ValidatorAppID = pv

	return nil
}

// SetAlgoBalance sets algo balance, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAlgoBalance(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.AlgoBalance = pv

	return nil
}

// SetRound sets round, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetRound(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrapped.Round = pv

	return nil
}

// GetAddress is a pool address
func (p *PoolInfo) GetAddress() string {
	return p.wrapped.Address
}

// GetAsset1ID is an asset1 id
func (p *PoolInfo) GetAsset1ID() string {
	return strconv.FormatUint(p.wrapped.Asset1ID, 10)
}

// GetAsset2ID is an asset2 id
func (p *PoolInfo) GetAsset2ID() string {
	return strconv.FormatUint(p.wrapped.Asset2ID, 10)
}

// GetAsset1UnitName is an asset1 unit name
func (p *PoolInfo) GetAsset1UnitName() string {
	return p.wrapped.Asset1UnitName
}

// GetAsset2UnitName is an asset2 unit name
func (p *PoolInfo) GetAsset2UnitName() string {
	return p.wrapped.Asset2UnitName
}

// GetLiquidityAssetID is an asset id for the liquidity
func (p *PoolInfo) GetLiquidityAssetID() string {
	return strconv.FormatUint(p.wrapped.LiquidityAssetID, 10)
}

// GetLiquidityAssetName is an asset name for the liquidity
func (p *PoolInfo) GetLiquidityAssetName() string {
	return p.wrapped.LiquidityAssetName
}

// GetAsset1Reserves is an asset1's reserves value
func (p *PoolInfo) GetAsset1Reserves() string {
	return strconv.FormatUint(p.wrapped.Asset1Reserves, 10)
}

// GetAsset2Reserves is an asset2's reserves value
func (p *PoolInfo) GetAsset2Reserves() string {
	return strconv.FormatUint(p.wrapped.Asset2Reserves, 10)
}

// GetIssuedLiquidity is the total issued liquidity
func (p *PoolInfo) GetIssuedLiquidity() string {
	return strconv.FormatUint(p.wrapped.IssuedLiquidity, 10)
}

// GetUnclaimedProtocolFee is an unclaimed protocol fee
func (p *PoolInfo) GetUnclaimedProtocolFee() string {
	return strconv.FormatUint(p.wrapped.UnclaimedProtocolFee, 10)
}

// GetOutstandingAsset1Amount is an outstanding asset1 amount
func (p *PoolInfo) GetOutstandingAsset1Amount() string {
	return strconv.FormatUint(p.wrapped.OutstandingAsset1Amount, 10)
}

// GetOutstandingAsset2Amount is an outstanding asset2 amount
func (p *PoolInfo) GetOutstandingAsset2Amount() string {
	return strconv.FormatUint(p.wrapped.OutstandingAsset2Amount, 10)
}

// GetOutstandingLiquidityAssetAmount is an outstanding liquidity asset amount
func (p *PoolInfo) GetOutstandingLiquidityAssetAmount() string {
	return strconv.FormatUint(p.wrapped.OutstandingLiquidityAssetAmount, 10)
}

// GetValidatorAppID is the validator app id
func (p *PoolInfo) GetValidatorAppID() string {
	return strconv.FormatUint(p.wrapped.ValidatorAppID, 10)
}

// GetAlgoBalance is a balance of the pool
func (p *PoolInfo) GetAlgoBalance() string {
	return strconv.FormatUint(p.wrapped.AlgoBalance, 10)
}

// GetRound is the latest fetch round
func (p *PoolInfo) GetRound() string {
	return strconv.FormatUint(p.wrapped.Round, 10)
}

// FetchPoolInfo returns pool information for the given asset1 and asset2
// validatorAppID, asset1ID, asset2ID are converted to uint64
func FetchPoolInfo(ac *AlgodClient, validatorAppID, asset1ID, asset2ID string) (*PoolInfo, error) {
	uintValidatorAppID, err := strconv.ParseUint(validatorAppID, 10, 64)
	if err != nil {
		return nil, err
	}
	uintAsset1ID, err := strconv.ParseUint(asset1ID, 10, 64)
	if err != nil {
		return nil, err
	}
	uintAsset2ID, err := strconv.ParseUint(asset2ID, 10, 64)
	if err != nil {
		return nil, err
	}

	info, err := pools.PoolInfo(context.Background(), ac.wrapped, uintValidatorAppID, uintAsset1ID, uintAsset2ID)
	if err != nil {
		return nil, err
	}

	return &PoolInfo{wrapped: info}, nil
}
