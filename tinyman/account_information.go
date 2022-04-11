package tinyman

import "github.com/algorand/go-algorand-sdk/client/v2/common/models"

// AccountInformation represents an account information at a given round
// This can be initiated only by fetching via the algorand sdk
type AccountInformation struct {
	wrapped *models.Account
}
