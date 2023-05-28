package operations

import (
	"github.com/biter777/viz-go-lib/types"
)

// WitnessRewardOperation represents witness_reward operation data.
type WitnessRewardOperation struct {
	Witness string      `json:"witness"`
	Shares  types.Asset `json:"shares"`
}

// Type function that defines the type of operation WitnessRewardOperation.
func (op *WitnessRewardOperation) Type() OpType {
	return TypeWitnessReward
}

// Data returns the operation data WitnessRewardOperation.
func (op *WitnessRewardOperation) Data() interface{} {
	return op
}
