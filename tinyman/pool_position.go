package tinyman

import (
	"context"
	"fmt"
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// PoolPosition represents a user position in the pool
type PoolPosition struct {
	wrapped *types.PoolPosition
}

// NewPoolPosition createa a pool position
func NewPoolPosition() *PoolPosition {
	return &PoolPosition{
		wrapped: &types.PoolPosition{
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

	a.wrapped.Asset1 = *asset.wrapped

	return nil
}

// SetAssetAmount2 sets an asset amount 2
func (a *PoolPosition) SetAssetAmount2(asset *AssetAmount) error {
	if asset == nil {
		return fmt.Errorf("an asset amount cannot be nil")
	}

	a.wrapped.Asset2 = *asset.wrapped

	return nil
}

// SetLiquidityAssetAmount sets a liquidity asset amount
func (a *PoolPosition) SetLiquidityAssetAmount(asset *AssetAmount) error {
	if asset == nil {
		return fmt.Errorf("an asset amount cannot be nil")
	}

	a.wrapped.LiquidityAsset = *asset.wrapped

	return nil
}

// SetShare sets share of the liquidity pool
func (a *PoolPosition) SetShare(value string) error {
	pv, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	a.wrapped.Share = pv

	return nil
}

// AssetAmount1 is an asset amount 1
func (a *PoolPosition) AssetAmount1() *AssetAmount {
	return wrapAssetAmount(&a.wrapped.Asset1)
}

// AssetAmount2 is an asset amount 2
func (a *PoolPosition) AssetAmount2() *AssetAmount {
	return wrapAssetAmount(&a.wrapped.Asset2)
}

// LiquidityAssetAmount is a asset asset amount
func (a *PoolPosition) LiquidityAssetAmount() *AssetAmount {
	return wrapAssetAmount(&a.wrapped.LiquidityAsset)
}

// Share is a share of user which can be calculated as a percentage by (share * 100)
func (a *PoolPosition) Share() string {
	return strconv.FormatFloat(a.wrapped.Share, 'f', -1, 64)
}

// FetchPoolPosition fetches pool position of a user
func (p *Pool) FetchPoolPosition(userAddress string) (*PoolPosition, error) {
	position, err := p.wrapped.FetchPoolPosition(context.Background(), userAddress)
	if err != nil {
		return nil, err
	}

	return &PoolPosition{wrapped: position}, nil
}
