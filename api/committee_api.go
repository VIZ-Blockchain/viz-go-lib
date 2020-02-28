package api

//committee_api

//GetCommitteeRequest api request get_committee_request
func (api *API) GetCommitteeRequest(id uint32, count int32) (*CommitteeObject, error) {
	var resp CommitteeObject
	err := api.call("committee_api", "get_committee_request", []interface{}{id, count}, &resp)
	return &resp, err
}

//GetCommitteeRequestVotes api request get_committee_request_votes
func (api *API) GetCommitteeRequestVotes(id uint32) ([]*CommitteeVoteState, error) {
	var resp []*CommitteeVoteState
	err := api.call("committee_api", "get_committee_request_votes", []interface{}{id}, &resp)
	return resp, err
}

//GetCommitteeRequestsList api request get_committee_requests_list
func (api *API) GetCommitteeRequestsList(status uint16) ([]*uint16, error) {
	var resp []*uint16
	err := api.call("committee_api", "get_committee_requests_list", []interface{}{status}, &resp)
	return resp, err
}
