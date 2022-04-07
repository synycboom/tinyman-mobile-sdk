package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// SwapQuote represents a swap quote
type SwapQuote struct {
	wrapped *types.SwapQuote
	// // Slippage is a slippage
	// Slippage float64
}

// SwapType returns a swap type
func (s *SwapQuote) SwapType() string {
	return s.wrapped.SwapType
}

// AmountIn returns an input asset amount
func (s *SwapQuote) AmountIn() *AssetAmount {
	return unwrapAssetAmount(s.wrapped.AmountIn)
}

// AmountOut returns an output asset amount
func (s *SwapQuote) AmountOut() *AssetAmount {
	return unwrapAssetAmount(s.wrapped.AmountOut)
}

// SwapFee returns a swap fee asset amount
func (s *SwapQuote) SwapFee() *AssetAmount {
	return unwrapAssetAmount(s.wrapped.SwapFee)
}

// Slippage returns a slippage
func (s *SwapQuote) Slippage() string {
	return strconv.FormatFloat(s.wrapped.Slippage, 'f', -1, 64)
}

// AssetAmountOutWithSlippage returns a calculated output asset amount after applying the slippage
func (s *SwapQuote) AssetAmountOutWithSlippage() (*AssetAmount, error) {
	a, err := s.wrapped.AmountOutWithSlippage()
	if err != nil {
		return nil, err
	}

	return unwrapAssetAmount(a), nil
}

// AssetAmountInWithSlippage returns a calculated input asset amount after applying the slippage
func (s *SwapQuote) AssetAmountInWithSlippage() (*AssetAmount, error) {
	a, err := s.wrapped.AmountInWithSlippage()
	if err != nil {
		return nil, err
	}

	return unwrapAssetAmount(a), nil
}

// Price returns the price, the value is converted from float64 to string
func (s *SwapQuote) Price() string {
	return strconv.FormatFloat(s.wrapped.Price(), 'f', -1, 64)
}

// PriceWithSlippage returns the price after applying the slippage, the value is converted from float64 to string
func (s *SwapQuote) PriceWithSlippage() (string, error) {
	value, err := s.wrapped.PriceWithSlippage()
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(value, 'f', -1, 64), nil
}
