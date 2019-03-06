package api

//invite_api

//GetInvitesList api request get_invites_list
/*
status:
0 - open
2 - close
*/
func (api *API) GetInvitesList(status uint16) ([]*int64, error) {
	var resp []*int64
	err := api.call("invite_api", "get_invites_list", []interface{}{status}, &resp)
	return resp, err
}

//GetInviteById api request get_invite_by_id
func (api *API) GetInviteById(id int64) (*InviteObject, error) {
	var resp InviteObject
	err := api.call("invite_api", "get_invite_by_id", []interface{}{id}, &resp)
	return &resp, err
}

//GetInviteByKey api request get_invite_by_key
func (api *API) GetInviteByKey(key string) (*InviteObject, error) {
	var resp InviteObject
	err := api.call("invite_api", "get_invite_by_key", []interface{}{key}, &resp)
	return &resp, err
}
