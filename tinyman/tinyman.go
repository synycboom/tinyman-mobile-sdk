package tinyman

import (
	"context"
	"fmt"
	"github.com/synycboom/tinyman-go-sdk/v1"
	"strconv"
)

// Client represents the Tinyman client
type Client struct {
	wrapped     *tinyman.Client
	algodClient *AlgodClient
}

// NewClient create a Tinyman client
// validatorAppID is converted to uint64
func NewClient(algodCli *AlgodClient, validatorAppID string, userAddress string) (*Client, error) {
	if algodCli == nil {
		return nil, fmt.Errorf("algodCli is required")
	}
	uintValidatorAppID, err := strconv.ParseUint(validatorAppID, 10, 64)
	if err != nil {
		return nil, err
	}

	return &Client{
		algodClient: algodCli,
		wrapped:     tinyman.NewClient(algodCli.wrapped, uintValidatorAppID, userAddress),
	}, nil
}

// NewTestNetClient create a test net Tinyman client
func NewTestNetClient(algodCli *AlgodClient, userAddress string) (*Client, error) {
	if algodCli == nil {
		return nil, fmt.Errorf("algodCli is required")
	}

	cli, err := tinyman.NewTestNetClient(algodCli.wrapped, userAddress)
	if err != nil {
		return nil, err
	}

	return &Client{
		algodClient: algodCli,
		wrapped:     cli,
	}, nil
}

// NewMainNetClient create a main net Tinyman client
func NewMainNetClient(algodCli *AlgodClient, userAddress string) (*Client, error) {
	if algodCli == nil {
		return nil, fmt.Errorf("algodCli is required")
	}

	cli, err := tinyman.NewMainNetClient(algodCli.wrapped, userAddress)
	if err != nil {
		return nil, err
	}

	return &Client{
		algodClient: algodCli,
		wrapped:     cli,
	}, nil
}

// FetchPool fetches a pool for given asset1 and asset2
func (c *Client) FetchPool(asset1, asset2 *Asset, fetch bool) (*Pool, error) {
	if asset1 == nil || asset2 == nil {
		return nil, fmt.Errorf("asset1 and asset2 are required")
	}

	pool, err := c.wrapped.FetchPool(context.Background(), asset1.wrapped, asset2.wrapped, fetch)
	if err != nil {
		return nil, err
	}

	return &Pool{wrapped: pool}, nil
}

// FetchAsset fetches an asset for a given asset id
// assetID is converted to uint64
func (c *Client) FetchAsset(assetID string) (*Asset, error) {
	uintAssetID, err := strconv.ParseUint(assetID, 10, 64)
	if err != nil {
		return nil, err
	}

	asset, err := c.wrapped.FetchAsset(context.Background(), uintAssetID)
	if err != nil {
		return nil, err
	}

	return &Asset{wrapped: asset}, nil
}

// Submit submits a transaction group to the blockchain
func (c *Client) Submit(txGroup *TransactionGroup, wait bool) (string, error) {
	return c.wrapped.Submit(context.Background(), txGroup.wrapped, wait)
}

// PrepareAppOptInTransaction prepares an app opt-in transaction and returns a transaction group
func (c *Client) PrepareAppOptInTransaction(userAddress string) (*TransactionGroup, error) {
	txGroup, err := c.wrapped.PrepareAppOptInTransaction(context.Background(), userAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}

// PrepareAssetOptInTransactions prepares asset opt-in transaction and returns a transaction group
// assetID is converted to uint64
func (c *Client) PrepareAssetOptInTransactions(assetID string, userAddress string) (*TransactionGroup, error) {
	uintAssetID, err := strconv.ParseUint(assetID, 10, 64)
	if err != nil {
		return nil, err
	}

	txGroup, err := c.wrapped.PrepareAssetOptInTransactions(context.Background(), uintAssetID, userAddress)
	if err != nil {
		return nil, err
	}

	return &TransactionGroup{wrapped: txGroup}, nil
}

// FetchExcessAmount fetches user's excess amounts and returns redeem quote iterator
func (c *Client) FetchExcessAmount(userAddr string) (*RedeemQuoteIterator, error) {
	quotes, err := c.wrapped.FetchExcessAmount(context.Background(), userAddr)
	if err != nil {
		return nil, err
	}

	iter := RedeemQuoteIterator{}
	for _, quote := range quotes {
		quote := quote
		iter.values = append(iter.values, &RedeemQuote{wrapped: &quote})
	}

	return &iter, nil
}

// IsOptedIn checks whether a user opted in for the application
func (c *Client) IsOptedIn(userAddress string) (bool, error) {
	return c.wrapped.IsOptedIn(context.Background(), userAddress)
}

// IsAssetOptedIn checks whether a user opted in for asset
// assetID is converted to uint64
func (c *Client) IsAssetOptedIn(assetID string, userAddress string) (bool, error) {
	uintAssetID, err := strconv.ParseUint(assetID, 10, 64)
	if err != nil {
		return false, err
	}

	return c.wrapped.IsAssetOptedIn(context.Background(), uintAssetID, userAddress)
}

// FetchBalance returns an asset balance of a user
func (c *Client) FetchBalance(asset *Asset, userAddress string) (*AssetAmount, error) {
	if asset == nil {
		return nil, fmt.Errorf("asset is required")
	}

	a, err := c.wrapped.Balance(context.Background(), asset.wrapped, userAddress)
	if err != nil {
		return nil, err
	}

	return &AssetAmount{wrapped: a}, nil
}

// GetValidatorAppID returns a validator app id
func (c *Client) GetValidatorAppID() string {
	return strconv.FormatUint(c.wrapped.ValidatorAppID, 10)
}

// GetUserAddress returns a user address attached to the client
func (c *Client) GetUserAddress() string {
	return c.wrapped.UserAddress
}
