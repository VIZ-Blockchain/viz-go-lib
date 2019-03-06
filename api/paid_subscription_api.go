package api

//paid_subscription_api

//GetPaidSubscriptionOptions api request get_paid_subscription_options
func (api *API) GetPaidSubscriptionOptions(account string) (*SubscriptionState, error) {
	var resp SubscriptionState
	err := api.call("paid_subscription_api", "get_paid_subscription_options", []interface{}{account}, &resp)
	return &resp, err
}

//GetPaidSubscriptionStatus api request get_paid_subscription_status
func (api *API) GetPaidSubscriptionStatus(subscriber, account string) (*SubscribeState, error) {
	var resp SubscribeState
	err := api.call("paid_subscription_api", "get_paid_subscription_status", []interface{}{subscriber, account}, &resp)
	return &resp, err
}

//GetActivePaidSubscriptions api request get_active_paid_subscriptions
func (api *API) GetActivePaidSubscriptions(subscriber string) ([]*string, error) {
	var resp []*string
	err := api.call("paid_subscription_api", "get_active_paid_subscriptions", []interface{}{subscriber}, &resp)
	return resp, err
}

//GetInactivePaidSubscriptions api request get_inactive_paid_subscriptions
func (api *API) GetInactivePaidSubscriptions(subscriber string) ([]*string, error) {
	var resp []*string
	err := api.call("paid_subscription_api", "get_inactive_paid_subscriptions", []interface{}{subscriber}, &resp)
	return resp, err
}
