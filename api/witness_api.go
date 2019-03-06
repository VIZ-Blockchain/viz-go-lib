package api

import (
	"github.com/VIZ-Blockchain/viz-go-lib/transports"
	"github.com/pkg/errors"
)

//witness_api

//GetActiveWitnesses api request get_active_witnesses
func (api *API) GetActiveWitnesses() ([]*string, error) {
	var resp []*string
	err := api.call("witness_api", "get_active_witnesses", transports.EmptyParams, &resp)
	return resp, err
}

//GetWitnessByAccount api request get_witness_by_account
func (api *API) GetWitnessByAccount(author string) (*Witness, error) {
	var resp Witness
	err := api.call("witness_api", "get_witness_by_account", []string{author}, &resp)
	return &resp, err
}

//GetWitnessCount api request get_witness_count
func (api *API) GetWitnessCount() (*uint64, error) {
	var resp uint64
	err := api.call("witness_api", "get_witness_count", transports.EmptyParams, &resp)
	return &resp, err
}

//GetWitnessSchedule api request get_witness_schedule
func (api *API) GetWitnessSchedule() (*WitnessSchedule, error) {
	var resp WitnessSchedule
	err := api.call("witness_api", "get_witness_schedule", transports.EmptyParams, &resp)
	return &resp, err
}

//GetWitnesses api request get_witnesses
func (api *API) GetWitnesses(id ...uint32) ([]*Witness, error) {
	var resp []*Witness
	err := api.call("witness_api", "get_witnesses", []interface{}{id}, &resp)
	return resp, err
}

//GetWitnessByVote api request get_witnesses_by_vote
func (api *API) GetWitnessByVote(author string, limit uint32) ([]*Witness, error) {
	if limit > 1000 {
		return nil, errors.New("witness_api: get_witnesses_by_vote -> limit must not exceed 1000")
	}
	var resp []*Witness
	err := api.call("witness_api", "get_witnesses_by_vote", []interface{}{author, limit}, &resp)
	return resp, err
}

//LookupWitnessAccounts api request lookup_witness_accounts
func (api *API) LookupWitnessAccounts(author string, limit uint32) ([]*string, error) {
	if limit > 1000 {
		return nil, errors.New("witness_api: lookup_witness_accounts -> limit must not exceed 1000")
	}
	var resp []*string
	err := api.call("witness_api", "lookup_witness_accounts", []interface{}{author, limit}, &resp)
	return resp, err
}
