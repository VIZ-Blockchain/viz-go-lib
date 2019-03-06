package types

//CurationRewardOperation represents curation_reward operation data.
type CurationRewardOperation struct {
	Curator         string `json:"curator"`
	Reward          *Asset `json:"reward"`
	ContentAuthor   string `json:"content_author"`
	ContentPermlink string `json:"content_permlink"`
}

//Type function that defines the type of operation CurationRewardOperation.
func (op *CurationRewardOperation) Type() OpType {
	return TypeCurationReward
}

//Data returns the operation data CurationRewardOperation.
func (op *CurationRewardOperation) Data() interface{} {
	return op
}
