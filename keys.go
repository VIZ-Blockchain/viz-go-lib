package viz

import (
	"errors"

	"github.com/VIZ-Blockchain/viz-go-lib/encoding/wif"
	"github.com/VIZ-Blockchain/viz-go-lib/operations"
)

var (
	//OpTypeKey include a description of the operation and the key needed to sign it
	OpTypeKey = make(map[operations.OpType][]string)
)

//Keys is used as a keystroke for a specific user.
//Only a few keys can be set.
type Keys struct {
	PKey []string
	AKey []string
	OKey []string
	MKey []string
}

func init() {
	OpTypeKey["vote"] = []string{"regular"}
	OpTypeKey["comment"] = []string{"regular"}
	OpTypeKey["transfer"] = []string{"active"}
	OpTypeKey["transfer_to_vesting"] = []string{"active"}
	OpTypeKey["withdraw_vesting"] = []string{"active"}
	OpTypeKey["limit_order_create"] = []string{"active"}
	OpTypeKey["limit_order_cancel"] = []string{"active"}
	OpTypeKey["feed_publish"] = []string{"active"}
	OpTypeKey["convert"] = []string{"active"}
	OpTypeKey["account_create"] = []string{"active"}
	OpTypeKey["account_update"] = []string{"active"}
	OpTypeKey["witness_update"] = []string{"active"}
	OpTypeKey["account_witness_vote"] = []string{"regular"}
	OpTypeKey["account_witness_proxy"] = []string{"regular"}
	OpTypeKey["delete_comment"] = []string{"regular"}
	OpTypeKey["custom_json"] = []string{"regular"}
	OpTypeKey["comment_options"] = []string{"regular"}
	OpTypeKey["set_withdraw_vesting_route"] = []string{"active"}
	OpTypeKey["limit_order_create2"] = []string{"active"}
	OpTypeKey["challenge_authority"] = []string{"master"}
	OpTypeKey["prove_authority"] = []string{"active"}
	OpTypeKey["request_account_recovery"] = []string{"active"}
	OpTypeKey["recover_account"] = []string{"master"}
	OpTypeKey["change_recovery_account"] = []string{"master"}
	OpTypeKey["escrow_transfer"] = []string{"active"}
	OpTypeKey["escrow_dispute"] = []string{"active"}
	OpTypeKey["escrow_release"] = []string{"active"}
	OpTypeKey["escrow_approve"] = []string{"active"}
	OpTypeKey["transfer_to_savings"] = []string{"active"}
	OpTypeKey["transfer_from_savings"] = []string{"active"}
	OpTypeKey["cancel_transfer_from_savings"] = []string{"active"}
	OpTypeKey["decline_voting_rights"] = []string{"master"}
	OpTypeKey["reset_account"] = []string{"active"}
	OpTypeKey["set_reset_account"] = []string{"active"}
	OpTypeKey["delegate_vesting_shares"] = []string{"active"}
	OpTypeKey["account_create_with_delegation"] = []string{"active"}
	OpTypeKey["account_metadata"] = []string{"regular"}
	OpTypeKey["proposal_create"] = []string{"active"}
	OpTypeKey["proposal_update"] = []string{"active"}
	OpTypeKey["proposal_delete"] = []string{"active"}
	OpTypeKey["chain_properties_update"] = []string{"active"}
	OpTypeKey["break_free_referral"] = []string{"active"}
	OpTypeKey["delegate_vesting_shares_with_interest"] = []string{"active"}
	OpTypeKey["reject_vesting_shares_delegation"] = []string{"active"}
	OpTypeKey["transit_to_cyberway"] = []string{"active"}
	OpTypeKey["worker_request"] = []string{"active"}
	OpTypeKey["worker_request_delete"] = []string{"active"}
	OpTypeKey["worker_request_vote"] = []string{"active"}
}

//SigningKeys returns the key from the CurrentKeys
func (client *Client) SigningKeys(trx operations.Operation) ([][]byte, error) {
	var keys [][]byte

	if client.CurrentKeys == nil {
		return nil, errors.New("Client Keys not initialized. Use SetKeys method")
	}

	opKeys := OpTypeKey[trx.Type()]
	for _, val := range opKeys {
		switch val {
		case "regular":
			for _, keyStr := range client.CurrentKeys.PKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Regular Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "active":
			for _, keyStr := range client.CurrentKeys.AKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Active Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "master":
			for _, keyStr := range client.CurrentKeys.OKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Master Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "memo":
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
