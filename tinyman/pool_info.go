package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// PoolInfo represents pool information
type PoolInfo struct {
	wrappedPoolInfo *types.PoolInfo
}

// NewPoolInfo creates a pool information
func NewPoolInfo() *PoolInfo {
	return &PoolInfo{
		wrappedPoolInfo: &types.PoolInfo{
			Address:                         "",
			Asset1ID:                        0,
			Asset2ID:                        0,
			Asset1UnitName:                  "",
			Asset2UnitName:                  "",
			LiquidityAssetID:                0,
			LiquidityAssetName:              "",
			Asset1Reserves:                  0,
			Asset2Reserves:                  0,
			IssuedLiquidity:                 0,
			UnclaimedProtocolFee:            0,
			OutstandingAsset1Amount:         0,
			OutstandingAsset2Amount:         0,
			OutstandingLiquidityAssetAmount: 0,
			ValidatorAppID:                  0,
			AlgoBalance:                     0,
			Round:                           0,
		},
	}
}

// SetAddress sets an address of the pool
func (p *PoolInfo) SetAddress(value string) {
	p.wrappedPoolInfo.Address = value
}

// SetAsset1ID sets an asset1 id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAsset1ID(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.Asset1ID = pv

	return nil
}

// SetAsset2ID sets an asset2 id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAsset2ID(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.Asset2ID = pv

	return nil
}

// SetAsset1UnitName sets an asset1 unit name
func (p *PoolInfo) SetAsset1UnitName(value string) {
	p.wrappedPoolInfo.Asset1UnitName = value
}

// SetAsset2UnitName sets an asset1 unit name
func (p *PoolInfo) SetAsset2UnitName(value string) {
	p.wrappedPoolInfo.Asset2UnitName = value
}

// SetLiquidityAssetID sets a liquidity asset id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetLiquidityAssetID(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.LiquidityAssetID = pv

	return nil
}

// SetLiquidityAssetName sets a liquidity name
func (p *PoolInfo) SetLiquidityAssetName(value string) {
	p.wrappedPoolInfo.LiquidityAssetName = value
}

// SetAsset1Reserves sets a liquidity asset id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAsset1Reserves(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.Asset1Reserves = pv

	return nil
}

// SetAsset2Reserves sets a liquidity asset id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAsset2Reserves(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.Asset2Reserves = pv

	return nil
}

// SetIssuedLiquidity sets issued liquidity amount, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetIssuedLiquidity(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.IssuedLiquidity = pv

	return nil
}

// SetUnclaimedProtocolFee sets unclaimed protocol fee, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetUnclaimedProtocolFee(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.UnclaimedProtocolFee = pv

	return nil
}

// SetOutstandingAsset1Amount sets outstanding asset1 amount, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetOutstandingAsset1Amount(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.OutstandingAsset1Amount = pv

	return nil
}

// SetOutstandingAsset2Amount sets outstanding asset2 amount, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetOutstandingAsset2Amount(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.OutstandingAsset2Amount = pv

	return nil
}

// SetOutstandingLiquidityAssetAmount sets outstanding liquidity asset amount, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetOutstandingLiquidityAssetAmount(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.OutstandingLiquidityAssetAmount = pv

	return nil
}

// SetValidatorAppID sets the validator app id, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetValidatorAppID(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.ValidatorAppID = pv

	return nil
}

// SetAlgoBalance sets algo balance, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetAlgoBalance(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.AlgoBalance = pv

	return nil
}

// SetRound sets round, the value will be converted to 64-bit unsigned integer
func (p *PoolInfo) SetRound(value string) error {
	pv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	p.wrappedPoolInfo.Round = pv

	return nil
}

// Address is a pool address
func (p *PoolInfo) Address() string {
	return p.wrappedPoolInfo.Address
}

// Asset1ID is an asset1 id
func (p *PoolInfo) Asset1ID() string {
	return strconv.FormatUint(p.wrappedPoolInfo.Asset1ID, 10)
}

// Asset2ID is an asset2 id
func (p *PoolInfo) Asset2ID() string {
	return strconv.FormatUint(p.wrappedPoolInfo.Asset2ID, 10)
}

// Asset1UnitName is an asset1 unit name
func (p *PoolInfo) Asset1UnitName() string {
	return p.wrappedPoolInfo.Asset1UnitName
}

// Asset2UnitName is an asset2 unit name
func (p *PoolInfo) Asset2UnitName() string {
	return p.wrappedPoolInfo.Asset2UnitName
}

// LiquidityAssetID is an asset id for the liquidity
func (p *PoolInfo) LiquidityAssetID() string {
	return strconv.FormatUint(p.wrappedPoolInfo.LiquidityAssetID, 10)
}

// LiquidityAssetName is an asset name for the liquidity
func (p *PoolInfo) LiquidityAssetName() string {
	return p.wrappedPoolInfo.LiquidityAssetName
}

// Asset1Reserves is an asset1's reserves value
func (p *PoolInfo) Asset1Reserves() string {
	return strconv.FormatUint(p.wrappedPoolInfo.Asset1Reserves, 10)
}

// Asset2Reserves is an asset2's reserves value
func (p *PoolInfo) Asset2Reserves() string {
	return strconv.FormatUint(p.wrappedPoolInfo.Asset2Reserves, 10)
}

// IssuedLiquidity is the total issued liquidity
func (p *PoolInfo) IssuedLiquidity() string {
	return strconv.FormatUint(p.wrappedPoolInfo.IssuedLiquidity, 10)
}

// UnclaimedProtocolFee is an unclaimed protocol fee
func (p *PoolInfo) UnclaimedProtocolFee() string {
	return strconv.FormatUint(p.wrappedPoolInfo.UnclaimedProtocolFee, 10)
}

// OutstandingAsset1Amount is an outstanding asset1 amount
func (p *PoolInfo) OutstandingAsset1Amount() string {
	return strconv.FormatUint(p.wrappedPoolInfo.OutstandingAsset1Amount, 10)
}

// OutstandingAsset2Amount is an outstanding asset2 amount
func (p *PoolInfo) OutstandingAsset2Amount() string {
	return strconv.FormatUint(p.wrappedPoolInfo.OutstandingAsset2Amount, 10)
}

// OutstandingLiquidityAssetAmount is an outstanding liquidity asset amount
func (p *PoolInfo) OutstandingLiquidityAssetAmount() string {
	return strconv.FormatUint(p.wrappedPoolInfo.OutstandingLiquidityAssetAmount, 10)
}

// ValidatorAppID is the validator app id
func (p *PoolInfo) ValidatorAppID() string {
	return strconv.FormatUint(p.wrappedPoolInfo.ValidatorAppID, 10)
}

// AlgoBalance is a balance of the pool
func (p *PoolInfo) AlgoBalance() string {
	return strconv.FormatUint(p.wrappedPoolInfo.AlgoBalance, 10)
}

// Round is the latest fetch round
func (p *PoolInfo) Round() string {
	return strconv.FormatUint(p.wrappedPoolInfo.Round, 10)
}
