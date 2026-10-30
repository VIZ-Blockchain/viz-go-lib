package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//ValidatorRewardOperation represents validator_reward operation data.
type ValidatorRewardOperation struct {
	Validator string    `json:"validator"`
	Shares  types.Asset `json:"shares"`
}

//Type function that defines the type of operation ValidatorRewardOperation.
func (op *ValidatorRewardOperation) Type() OpType {
	return TypeValidatorReward
}

//Data returns the operation data ValidatorRewardOperation.
func (op *ValidatorRewardOperation) Data() interface{} {
	return op
}
