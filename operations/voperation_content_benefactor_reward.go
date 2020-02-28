package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//ContentBenefactorRewardOperation represents content_benefactor_reward operation data.
type ContentBenefactorRewardOperation struct {
	Benefactor string      `json:"benefactor"`
	Author     string      `json:"author"`
	Permlink   string      `json:"permlink"`
	Reward     types.Asset `json:"reward"`
}

//Type function that defines the type of operation ContentBenefactorRewardOperation.
func (op *ContentBenefactorRewardOperation) Type() OpType {
	return TypeContentBenefactorReward
}

//Data returns the operation data ContentBenefactorRewardOperation.
func (op *ContentBenefactorRewardOperation) Data() interface{} {
	return op
}
