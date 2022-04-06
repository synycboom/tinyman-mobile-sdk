package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// Asset is an Algorand token
type Asset struct {
	asset *types.Asset
}

// NewAsset creates an asset
func NewAsset(id, decimals, name, unitName string) (*Asset, error) {
	a := &Asset{
		asset: &types.Asset{
			Name:     name,
			UnitName: unitName,
		},
	}

	if err := a.SetID(id); err != nil {
		return nil, err
	}

	if err := a.SetDecimals(decimals); err != nil {
		return nil, err
	}

	return a, nil
}

// SetID sets an id of the asset by converting a string to 64-bit unsigned integer
func (a *Asset) SetID(value string) error {
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	a.asset.ID = id

	return nil
}

// SetDecimals sets a decimals of the asset by converting a string to 64-bit unsigned integer
func (a *Asset) SetDecimals(value string) error {
	decimals, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}

	a.asset.Decimals = decimals

	return nil
}

// SetName sets a name of the asset
func (a *Asset) SetName(value string) {
	a.asset.Name = value
}

// SetUnitName sets a unit name of the asset
func (a *Asset) SetUnitName(value string) {
	a.asset.UnitName = value
}

// ID returns an id of the asset by converting a 64-bit unsigned integer to a string
func (a *Asset) ID() string {
	return strconv.FormatUint(a.asset.ID, 10)
}

// Decimals returns a decimals of the asset by converting a 64-bit unsigned integer to a string
func (a *Asset) Decimals() string {
	return strconv.FormatUint(a.asset.Decimals, 10)
}

// Name returns a name of the asset
func (a *Asset) Name() string {
	return a.asset.Name
}

// UnitName returns a unit name of the asset
func (a *Asset) UnitName() string {
	return a.asset.UnitName
}
