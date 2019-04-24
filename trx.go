package client

import (
	"time"

	"github.com/VIZ-Blockchain/viz-go-lib/api"
	"github.com/VIZ-Blockchain/viz-go-lib/transactions"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//BResp of response when sending a transaction.
type BResp struct {
	ID       string
	BlockNum uint32
	TrxNum   string
	Expired  bool
	JSONTrx  string
}

//SendTrx generates and sends an array of transactions to VIZ.
func (client *Client) SendTrx(username string, strx []types.Operation) (*BResp, error) {
	var bresp BResp

	// Getting the necessary parameters
	props, err := client.API.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}

	// Creating a Transaction
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := transactions.NewSignedTransaction(&types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	})

	// Adding Operations to a Transaction
	for _, val := range strx {
		tx.PushOperation(val)
	}

	expTime := time.Now().Add(59 * time.Minute).UTC()
	tm := types.Time{
		Time: &expTime,
	}
	tx.Expiration = &tm

	// Obtain the key required for signing
	privKeys, err := client.SigningKeys(strx[0])
	if err != nil {
		return nil, err
	}

	// Sign the transaction
	if err := tx.Sign(privKeys, client.chainID); err != nil {
		return nil, err
	}

	// Sending a transaction
	var resp *api.BroadcastResponse
	if client.AsyncProtocol {
		err = client.API.BroadcastTransaction(tx.Transaction)
	} else {
		resp, err = client.API.BroadcastTransactionSynchronous(tx.Transaction)
	}

	bresp.JSONTrx, _ = JSONTrxString(tx)

	if err != nil {
		return &bresp, err
	}
	if resp != nil && !client.AsyncProtocol {
		bresp.ID = resp.ID
		bresp.BlockNum = resp.BlockNum
		bresp.TrxNum = resp.TrxNum
		bresp.Expired = resp.Expired

		return &bresp, nil
	}

	return &bresp, nil
}

func (client *Client) GetTrx(strx []types.Operation) (*types.Transaction, error) {
	// Getting the necessary parameters
	props, err := client.API.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}

	// Creating a Transaction
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}
	tx := &types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	}

	// Adding Operations to a Transaction
	for _, val := range strx {
		tx.PushOperation(val)
	}

	expTime := time.Now().Add(59 * time.Minute).UTC()
	tm := types.Time{
		Time: &expTime,
	}
	tx.Expiration = &tm

	return tx, nil
}
