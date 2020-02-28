package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//AuthorRewardOperation represents author_reward operation data.
type AuthorRewardOperation struct {
	Author        string      `json:"author"`
	Permlink      string      `json:"permlink"`
	TokenPayout   types.Asset `json:"token_payout"`
	VestingPayout types.Asset `json:"vesting_payout"`
}

//Type function that defines the type of operation AuthorRewardOperation.
func (op *AuthorRewardOperation) Type() OpType {
	return TypeAuthorReward
}

//Data returns the operation data AuthorRewardOperation.
func (op *AuthorRewardOperation) Data() interface{} {
	return op
}
