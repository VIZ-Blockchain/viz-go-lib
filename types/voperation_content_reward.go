package types

//ContentRewardOperation represents content_reward operation data.
type ContentRewardOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Payout   *Asset `json:"payout"`
}

//Type function that defines the type of operation ContentRewardOperation.
func (op *ContentRewardOperation) Type() OpType {
	return TypeContentReward
}

//Data returns the operation data ContentRewardOperation.
func (op *ContentRewardOperation) Data() interface{} {
	return op
}
