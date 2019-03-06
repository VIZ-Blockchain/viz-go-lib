package types

//CancelPaidSubscriptionOperation represents cancel_paid_subscription operation data.
type CancelPaidSubscriptionOperation struct {
	Subscriber string `json:"subscriber"`
	Account    string `json:"account"`
}

//Type function that defines the type of operation CancelPaidSubscriptionOperation.
func (op *CancelPaidSubscriptionOperation) Type() OpType {
	return TypeCancelPaidSubscription
}

//Data returns the operation data CancelPaidSubscriptionOperation.
func (op *CancelPaidSubscriptionOperation) Data() interface{} {
	return op
}
