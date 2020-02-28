package api

import (
	"encoding/json"
	"errors"

	"github.com/VIZ-Blockchain/viz-go-lib/operations"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//database_api

//GetAccountCount returns the number of registered users.
func (api *API) GetAccountCount() (*uint64, error) {
	var resp uint64
	err := api.call("database_api", "get_account_count", EmptyParams, &resp)
	return &resp, err
}

//GetAccounts api request get_accounts
func (api *API) GetAccounts(accountNames ...string) ([]*Account, error) {
	var resp []*Account
	err := api.call("database_api", "get_accounts", []interface{}{accountNames}, &resp)
	return resp, err
}

//GetBlock api request get_block
func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	var resp Block
	err := api.call("database_api", "get_block", []uint32{blockNum}, &resp)
	resp.Number = blockNum
	return &resp, err
}

//GetBlockHeader api request get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	var resp BlockHeader
	err := api.call("database_api", "get_block_header", []uint32{blockNum}, &resp)
	resp.Number = blockNum
	return &resp, err
}

//GetChainProperties api request get_chain_properties
func (api *API) GetChainProperties() (*types.ChainProperties, error) {
	var resp types.ChainProperties
	err := api.call("database_api", "get_chain_properties", EmptyParams, &resp)
	return &resp, err
}

//GetConfig api request get_config
func (api *API) GetConfig() (*Config, error) {
	var resp Config
	err := api.call("database_api", "get_config", EmptyParams, &resp)
	return &resp, err
}

//GetDatabaseInfo api request get_database_info
func (api *API) GetDatabaseInfo() (*DatabaseInfo, error) {
	var resp DatabaseInfo
	err := api.call("database_api", "get_database_info", EmptyParams, &resp)
	return &resp, err
}

//GetDynamicGlobalProperties api request get_dynamic_global_properties
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	err := api.call("database_api", "get_dynamic_global_properties", EmptyParams, &resp)
	return &resp, err
}

//GetEscrow api request get_escrow
func (api *API) GetEscrow(from string, escrowID uint32) (*Escrow, error) {
	var resp Escrow
	err := api.call("database_api", "get_escrow", []interface{}{from, escrowID}, &resp)
	return &resp, err
}

//GetExpiringVestingDelegations api request get_expiring_vesting_delegations
func (api *API) GetExpiringVestingDelegations(account string, from types.Time, limit ...uint32) ([]*VestingDelegationExpiration, error) {
	params := []interface{}{account, from}
	switch len(limit) {
	case 0:
		params = append(params, 100)
	default:
		lm := limit[0]
		if lm > 1000 {
			return nil, errors.New("database_api: get_expiring_vesting_delegations -> limit must not exceed 1000")
		}
		params = append(params, lm)
	}
	var resp []*VestingDelegationExpiration
	err := api.call("database_api", "get_expiring_vesting_delegations", params, &resp)
	return resp, err
}

//GetHardforkVersion api request get_hardfork_version
func (api *API) GetHardforkVersion() (*string, error) {
	var resp string
	err := api.call("database_api", "get_hardfork_version", EmptyParams, &resp)
	return &resp, err
}

//GetNextScheduledHardfork api request get_next_scheduled_hardfork
func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	var resp NextScheduledHardfork
	err := api.call("database_api", "get_next_scheduled_hardfork", EmptyParams, &resp)
	return &resp, err
}

//GetMasterHistory api request get_master_history
func (api *API) GetMasterHistory(accountName string) ([]*MasterHistory, error) {
	var resp []*MasterHistory
	err := api.call("database_api", "get_master_history", []interface{}{accountName}, &resp)
	return resp, err
}

//GetPotentialSignatures api request get_potential_signatures
func (api *API) GetPotentialSignatures(trx *operations.Transaction) ([]*string, error) {
	var resp []*string
	err := api.call("database_api", "get_potential_signatures", []interface{}{&trx}, &resp)
	return resp, err
}

//GetProposedTransaction api request get_proposed_transactions
func (api *API) GetProposedTransaction(account string, from, limit uint32) ([]*ProposalObject, error) {
	var resp []*ProposalObject
	err := api.call("database_api", "get_proposed_transactions", []interface{}{account, from, limit}, &resp)
	return resp, err
}

//GetRecoveryRequest api request get_recovery_request
func (api *API) GetRecoveryRequest(accountName string) (*AccountRecoveryRequest, error) {
	var resp AccountRecoveryRequest
	err := api.call("database_api", "get_recovery_request", []interface{}{accountName}, &resp)
	return &resp, err
}

//GetRequiredSignatures api request get_required_signatures
func (api *API) GetRequiredSignatures(trx *operations.Transaction, keys ...string) ([]*string, error) {
	var resp []*string
	err := api.call("database_api", "get_required_signatures", []interface{}{trx, keys}, &resp)
	return resp, err
}

//GetTransactionHex api request get_transaction_hex
func (api *API) GetTransactionHex(trx *operations.Transaction) (*string, error) {
	var resp string
	err := api.call("database_api", "get_transaction_hex", []interface{}{&trx}, &resp)
	return &resp, err
}

//GetVestingDelegations api request get_vesting_delegations
//dtype:
//  delegated
//  any other
func (api *API) GetVestingDelegations(account, from, dtype string, limit ...uint32) ([]*VestingDelegation, error) {
	params := []interface{}{account, from}
	switch len(limit) {
	case 0:
		params = append(params, 100)
	default:
		lm := limit[0]
		if lm > 1000 {
			return nil, errors.New("database_api: get_vesting_delegations -> limit must not exceed 1000")
		}
		params = append(params, lm)
	}
	params = append(params, dtype)
	var resp []*VestingDelegation
	err := api.call("database_api", "get_vesting_delegations", params, &resp)
	return resp, err
}

//GetWithdrawRoutes api request get_withdraw_routes
/*
withdrawRouteType:
0 = incoming
1 = outgoing
2 = all
*/
func (api *API) GetWithdrawRoutes(accountName string, withdrawRouteType uint32) ([]*WithdrawVestingRoutes, error) {
	var resp []*WithdrawVestingRoutes
	err := api.call("database_api", "get_withdraw_routes", []interface{}{accountName, withdrawRouteType}, &resp)
	return resp, err
}

//LookupAccountNames api request lookup_account_names
func (api *API) LookupAccountNames(accountNames ...string) ([]*Account, error) {
	var resp []*Account
	err := api.call("database_api", "lookup_account_names", []interface{}{accountNames}, &resp)
	return resp, err
}

//LookupAccounts api request lookup_accounts
func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]*string, error) {
	var resp []*string
	err := api.call("database_api", "lookup_accounts", []interface{}{lowerBoundName, limit}, &resp)
	return resp, err
}

//GetVerifyAccountAuthority api request verify_account_authority
func (api *API) GetVerifyAccountAuthority(accountName string, keys ...string) (*bool, error) {
	var resp bool
	err := api.call("database_api", "verify_account_authority", []interface{}{accountName, keys}, &resp)
	return &resp, err
}

//GetVerifyAuthority api request verify_authority
func (api *API) GetVerifyAuthority(trx *operations.Transaction) (*bool, error) {
	var resp bool
	err := api.call("database_api", "verify_authority", []interface{}{&trx}, &resp)
	return &resp, err
}

//GetAccountsOnSale api request get_accounts_on_sale
func (api *API) GetAccountsOnSale(from, limit uint32) ([]*AccountOnSale, error) {
	if limit > 1000 {
		return nil, errors.New("database_api: get_accounts_on_sale -> limit must not exceed 1000")
	}
	var resp []*AccountOnSale
	err := api.call("database_api", "get_accounts_on_sale", []interface{}{from, limit}, &resp)
	return resp, err
}

//GetSubAccountsOnSale api request get_subaccounts_on_sale
func (api *API) GetSubAccountsOnSale(from, limit uint32) ([]*SubAccountOnSale, error) {
	if limit > 1000 {
		return nil, errors.New("database_api: get_subaccounts_on_sale -> limit must not exceed 1000")
	}
	var resp []*SubAccountOnSale
	err := api.call("database_api", "get_subaccounts_on_sale", []interface{}{from, limit}, &resp)
	return resp, err
}

// Set callback to invoke as soon as a new block is applied
func (api *API) SetBlockAppliedCallback(notice func(header *CallbackBlockResponse, error error)) (err error) {
	err = api.setCallback("database_api", "set_block_applied_callback", func(raw json.RawMessage) {
		var header CallbackBlockResponse
		if err := json.Unmarshal(raw, &header); err != nil {
			notice(nil, err)
		}
		notice(&header, nil)
	})
	return
}
