package tinyman

import (
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// PoolPosition represents a user position in the pool
type PoolPosition struct {
	wrappedPoolPosition *types.PoolPosition
}

// NewPoolPosition createa a pool position
func NewPoolPosition() *PoolPosition {
	return &PoolPosition{
		wrappedPoolPosition: &types.PoolPosition{
			Asset1: types.AssetAmount{
				Asset: &types.Asset{
					ID:       0,
					Decimals: 0,
					Name:     "",
					UnitName: "",
				},
				Amount: 0,
			},
			Asset2: types.AssetAmount{
				Asset: &types.Asset{
					ID:       0,
					Decimals: 0,
					Name:     "",
					UnitName: "",
				},
				Amount: 0,
			},
			LiquidityAsset: types.AssetAmount{
				Asset: &types.Asset{
					ID:       0,
					Decimals: 0,
					Name:     "",
					UnitName: "",
				},
				Amount: 0,
			},
			Share: 0.0,
		},
	}
}

// SetAssetAmount1 sets an asset amount 1
func (a *PoolPosition) SetAssetAmount1(asset *AssetAmount) error {
	if asset == nil {
		return fmt.Errorf("an asset amount cannot be nil")
	}

	a.wrappedPoolPosition.Asset1 = *asset.wrappedAssetAmount

	return nil
}

// SetAssetAmount2 sets an asset amount 2
func (a *PoolPosition) SetAssetAmount2(asset *AssetAmount) error {
	if asset == nil {
		return fmt.Errorf("an asset amount cannot be nil")
	}

	a.wrappedPoolPosition.Asset2 = *asset.wrappedAssetAmount

	return nil
}

// SetLiquidityAssetAmount sets a liquidity asset amount
func (a *PoolPosition) SetLiquidityAssetAmount(asset *AssetAmount) error {
	if asset == nil {
		return fmt.Errorf("an asset amount cannot be nil")
	}

	a.wrappedPoolPosition.LiquidityAsset = *asset.wrappedAssetAmount

	return nil
}

// SetShare sets share of the liquidity pool
func (a *PoolPosition) SetShare(value string) error {
	pv, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	a.wrappedPoolPosition.Share = pv

	return nil
}

// AssetAmount1 is an asset amount 1
func (a *PoolPosition) AssetAmount1() *AssetAmount {
	return unwrapAssetAmount(&a.wrappedPoolPosition.Asset1)
}

// AssetAmount2 is an asset amount 2
func (a *PoolPosition) AssetAmount2() *AssetAmount {
	return unwrapAssetAmount(&a.wrappedPoolPosition.Asset2)
}

// LiquidityAssetAmount is a asset asset amount
func (a *PoolPosition) LiquidityAssetAmount() *AssetAmount {
	return unwrapAssetAmount(&a.wrappedPoolPosition.LiquidityAsset)
}

// Share is a share of user which can be calculated as a percentage by (share * 100)
func (a *PoolPosition) Share(value string) string {
	return strconv.FormatFloat(a.wrappedPoolPosition.Share, 'f', -1, 64)
}
