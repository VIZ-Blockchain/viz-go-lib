package viz

import (
	"errors"

	"github.com/biter777/viz-go-lib/encoding/wif"
	"github.com/biter777/viz-go-lib/operations"
)

var (
	//OpTypeKey include a description of the operation and the key needed to sign it
	OpTypeKey = make(map[operations.OpType][]string)
)

// Keys is used as a keystroke for a specific user.
// Only a few keys can be set.
type Keys struct {
	PKey []string
	AKey []string
	OKey []string
	MKey []string
}

func init() {
	OpTypeKey["vote"] = []string{"regular"}
	OpTypeKey["content"] = []string{"regular"}
	OpTypeKey["transfer"] = []string{"active"}
	OpTypeKey["transfer_to_vesting"] = []string{"active"}
	OpTypeKey["withdraw_vesting"] = []string{"active"}
	OpTypeKey["account_update"] = []string{"active"}
	OpTypeKey["witness_update"] = []string{"active"}
	OpTypeKey["account_witness_vote"] = []string{"regular"}
	OpTypeKey["account_witness_proxy"] = []string{"regular"}
	OpTypeKey["delete_content"] = []string{"regular"}
	OpTypeKey["custom"] = []string{"regular"}
	OpTypeKey["set_withdraw_vesting_route"] = []string{"active"}
	OpTypeKey["request_account_recovery"] = []string{"active"}
	OpTypeKey["recover_account"] = []string{"master"}
	OpTypeKey["change_recovery_account"] = []string{"master"}
	OpTypeKey["escrow_transfer"] = []string{"active"}
	OpTypeKey["escrow_dispute"] = []string{"active"}
	OpTypeKey["escrow_release"] = []string{"active"}
	OpTypeKey["escrow_approve"] = []string{"active"}
	OpTypeKey["delegate_vesting_shares"] = []string{"active"}
	OpTypeKey["account_create"] = []string{"active"}
	OpTypeKey["account_metadata"] = []string{"regular"}
	OpTypeKey["proposal_create"] = []string{"active"}
	OpTypeKey["proposal_update"] = []string{"regular"}
	OpTypeKey["proposal_delete"] = []string{"active"}
	OpTypeKey["chain_properties_update"] = []string{"active"}
	OpTypeKey["committee_worker_create_request"] = []string{"regular"}
	OpTypeKey["committee_worker_cancel_request"] = []string{"regular"}
	OpTypeKey["committee_vote_request"] = []string{"regular"}
	OpTypeKey["create_invite"] = []string{"active"}
	OpTypeKey["claim_invite_balance"] = []string{"active"}
	OpTypeKey["invite_registration"] = []string{"active"}
	OpTypeKey["versioned_chain_properties_update"] = []string{"active"}
	OpTypeKey["award"] = []string{"regular"}
	OpTypeKey["set_paid_subscription"] = []string{"active"}
	OpTypeKey["paid_subscribe"] = []string{"active"}
	OpTypeKey["set_account_price"] = []string{"master"}
	OpTypeKey["set_subaccount_price"] = []string{"master"}
	OpTypeKey["buy_account"] = []string{"active"}
}

// SigningKeys returns the key from the CurrentKeys
func (client *Client) SigningKeys(trx operations.Operation) ([][]byte, error) {
	var keys [][]byte

	if client.CurrentKeys == nil {
		return nil, errors.New("Client Keys not initialized. Use SetKeys method")
	}

	opKeys := OpTypeKey[trx.Type()]
	for _, val := range opKeys {
		switch val {
		case "regular":
			if len(client.CurrentKeys.PKey) < 1 {
				return nil, errors.New("Client Regular Key not initialized. Use SetKeys method")
			}
			for _, keyStr := range client.CurrentKeys.PKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Regular Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "active":
			if len(client.CurrentKeys.AKey) < 1 {
				return nil, errors.New("Client Active Key not initialized. Use SetKeys method")
			}
			for _, keyStr := range client.CurrentKeys.AKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Active Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "master":
			if len(client.CurrentKeys.OKey) < 1 {
				return nil, errors.New("Client Master Key not initialized. Use SetKeys method")
			}
			for _, keyStr := range client.CurrentKeys.OKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Master Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "memo":
			if len(client.CurrentKeys.MKey) < 1 {
				return nil, errors.New("Client Memo Key not initialized. Use SetKeys method")
			}
			for _, keyStr := range client.CurrentKeys.MKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Memo Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		}
	}

	return keys, nil
}
