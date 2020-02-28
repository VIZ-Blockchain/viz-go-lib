package api

import (
	"errors"
)

//paid_subscription_api

//GetPaidSubscriptions api request get_paid_subscriptions
func (api *API) GetPaidSubscriptions(from, limit uint32) ([]*PaidSubscription, error) {
	if limit > 1000 {
		return nil, errors.New("database_api: get_paid_subscriptions -> limit must not exceed 1000")
	}
	var resp []*PaidSubscription
	err := api.call("paid_subscription_api", "get_paid_subscriptions", []interface{}{from, limit}, &resp)
	return resp, err
}

//GetPaidSubscriptionOptions api request get_paid_subscription_options
func (api *API) GetPaidSubscriptionOptions(account string) (*PaidSubscriptionState, error) {
	var resp PaidSubscriptionState
	err := api.call("paid_subscription_api", "get_paid_subscription_options", []interface{}{account}, &resp)
	return &resp, err
}

//GetPaidSubscriptionStatus api request get_paid_subscription_status
func (api *API) GetPaidSubscriptionStatus(subscriber, account string) (*PaidSubscribeState, error) {
	var resp PaidSubscribeState
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
