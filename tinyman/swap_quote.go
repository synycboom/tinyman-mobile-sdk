package tinyman

import (
	"strconv"

	"github.com/synycboom/tinyman-go-sdk/types"
)

// SwapQuote represents a swap quote
type SwapQuote struct {
	wrapped *types.SwapQuote
}

// GetSwapType returns a swap type
func (s *SwapQuote) GetSwapType() string {
	return s.wrapped.SwapType
}

// GetAssetAmountIn returns an input asset amount
func (s *SwapQuote) GetAssetAmountIn() *AssetAmount {
	return wrapAssetAmount(s.wrapped.AmountIn)
}

// GetAssetAmountOut returns an output asset amount
func (s *SwapQuote) GetAssetAmountOut() *AssetAmount {
	return wrapAssetAmount(s.wrapped.AmountOut)
}

// GetSwapFeeAssetAmount returns a swap fee asset amount
func (s *SwapQuote) GetSwapFeeAssetAmount() *AssetAmount {
	return wrapAssetAmount(s.wrapped.SwapFee)
}

// GetSlippage returns a slippage
func (s *SwapQuote) GetSlippage() string {
	return strconv.FormatFloat(s.wrapped.Slippage, 'f', -1, 64)
}

// GetAssetAmountOutWithSlippage returns a calculated output asset amount after applying the slippage
func (s *SwapQuote) GetAssetAmountOutWithSlippage() (*AssetAmount, error) {
	a, err := s.wrapped.AmountOutWithSlippage()
	if err != nil {
		return nil, err
	}

	return wrapAssetAmount(a), nil
}

// GetAssetAmountInWithSlippage returns a calculated input asset amount after applying the slippage
func (s *SwapQuote) GetAssetAmountInWithSlippage() (*AssetAmount, error) {
	a, err := s.wrapped.AmountInWithSlippage()
	if err != nil {
		return nil, err
	}

	return wrapAssetAmount(a), nil
}

// GetPrice returns the price, the value is converted from float64 to string
func (s *SwapQuote) GetPrice() string {
	return strconv.FormatFloat(s.wrapped.Price(), 'f', -1, 64)
}

// GetPriceWithSlippage returns the price after applying the slippage, the value is converted from float64 to string
func (s *SwapQuote) GetPriceWithSlippage() (string, error) {
	value, err := s.wrapped.PriceWithSlippage()
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(value, 'f', -1, 64), nil
}
