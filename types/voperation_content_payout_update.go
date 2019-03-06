package types

//ContentPayoutUpdateOperation represents content_payout_update operation data.
type ContentPayoutUpdateOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//Type function that defines the type of operation ContentPayoutUpdateOperation.
func (op *ContentPayoutUpdateOperation) Type() OpType {
	return TypeContentPayoutUpdate
}

//Data returns the operation data ContentPayoutUpdateOperation.
func (op *ContentPayoutUpdateOperation) Data() interface{} {
	return op
}
