package tinyman

import "github.com/algorand/go-algorand-sdk/types"

// SuggestedParams wraps the transaction parameters common to all transactions,
// typically received from the SuggestedParams endpoint of algod.
// This struct itself is not sent over the wire to or from algod: see models.TransactionParams.
type SuggestedParams struct {
	wrapped *types.SuggestedParams
}
