package operations

import (
	"github.com/biter777/viz-go-lib/types"
)

// PaidSubscriptionActionOperation represents paid_subscription_action operation data.
type PaidSubscriptionActionOperation struct {
	Subscriber         string      `json:"subscriber"`
	Account            string      `json:"account"`
	Level              uint16      `json:"level"`
	Amount             types.Asset `json:"amount"`
	Period             uint16      `json:"period"`
	SummaryDurationSec uint64      `json:"summary_duration_sec"`
	SummaryAmount      types.Asset `json:"summary_amount"`
}

// Type function that defines the type of operation PaidSubscriptionActionOperation.
func (op *PaidSubscriptionActionOperation) Type() OpType {
	return TypePaidSubscriptionAction
}

// Data returns the operation data PaidSubscriptionActionOperation.
func (op *PaidSubscriptionActionOperation) Data() interface{} {
	return op
}
