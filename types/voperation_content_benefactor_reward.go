package types

//ContentBenefactorRewardOperation represents content_benefactor_reward operation data.
type ContentBenefactorRewardOperation struct {
	Benefactor string `json:"benefactor"`
	Author     string `json:"author"`
	Permlink   string `json:"permlink"`
	Reward     *Asset `json:"reward"`
}

//Type function that defines the type of operation ContentBenefactorRewardOperation.
func (op *ContentBenefactorRewardOperation) Type() OpType {
	return TypeContentBenefactorReward
}

//Data returns the operation data ContentBenefactorRewardOperation.
func (op *ContentBenefactorRewardOperation) Data() interface{} {
	return op
}
