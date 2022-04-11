package tinyman

import (
	"context"
	"strconv"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/future"
)

// CreateAnAsset is a function to create an asset, only testing purpose
func CreateAnAsset(
	assetName,
	unitName,
	decimals,
	totalIssuance,
	userAddress string,
	account *Account,
	ac *AlgodClient,
) (string, error) {
	txParams, err := ac.SuggestedParams()
	if err != nil {
		return "", err
	}

	var note []byte = nil
	addr := userAddress
	reserve := addr
	freeze := addr
	clawback := addr
	manager := addr
	assetURL := "http://someurl"
	assetMetadataHash := "thisIsSomeLength32HashCommitment"

	uintTotalIssuance, err := strconv.ParseUint(totalIssuance, 10, 64)
	if err != nil {
		return "", err
	}
	uintDecimals, err := strconv.ParseUint(decimals, 10, 64)
	if err != nil {
		return "", err
	}

	txn, err := future.MakeAssetCreateTxn(addr, note, *txParams.wrapped,
		uintTotalIssuance, uint32(uintDecimals), false, manager, reserve, freeze, clawback,
		unitName, assetName, assetURL, assetMetadataHash,
	)
	if err != nil {
		return "", err
	}

	txid, stx, err := crypto.SignTransaction(account.wrapped.PrivateKey, txn)
	if err != nil {
		return "", err
	}

	if err := ac.SendRawTransaction(stx); err != nil {
		return "", err
	}

	confirmedTxn, err := future.WaitForConfirmation(ac.wrapped, txid, 4, context.Background())
	if err != nil {
		return "", err
	}

	return strconv.FormatUint(confirmedTxn.AssetIndex, 10), nil
}
