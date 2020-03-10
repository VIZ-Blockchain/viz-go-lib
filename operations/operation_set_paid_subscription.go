package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//SetPaidSubscriptionOperation represents set_paid_subscription operation data.
type SetPaidSubscriptionOperation struct {
	Account string       `json:"account"`
	URL     string       `json:"url"`
	Levels  uint16       `json:"levels"`
	Amount  *types.Asset `json:"amount"`
	Period  uint16       `json:"period"`
}

//Type function that defines the type of operation SetPaidSubscriptionOperation.
func (op *SetPaidSubscriptionOperation) Type() OpType {
	return TypeSetPaidSubscription
}

//Data returns the operation data SetPaidSubscriptionOperation.
func (op *SetPaidSubscriptionOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type SetPaidSubscriptionOperation to bytes.
func (op *SetPaidSubscriptionOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetPaidSubscription.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.URL)
	enc.Encode(op.Levels)
	enc.Encode(op.Amount)
	enc.Encode(op.Period)
	return enc.Err()
}
