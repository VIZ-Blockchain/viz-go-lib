package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//PaidSubscribeOperation represents paid_subscribe operation data.
type PaidSubscribeOperation struct {
	Subscriber  string      `json:"subscriber"`
	Account     string      `json:"account"`
	Level       uint16      `json:"level"`
	Amount      types.Asset `json:"amount"`
	Period      uint16      `json:"period"`
	AutoRenewal bool        `json:"auto_renewal"`
}

//Type function that defines the type of operation PaidSubscribeOperation.
func (op *PaidSubscribeOperation) Type() OpType {
	return TypePaidSubscribe
}

//Data returns the operation data PaidSubscribeOperation.
func (op *PaidSubscribeOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type PaidSubscribeOperation to bytes.
func (op *PaidSubscribeOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePaidSubscribe.Code()))
	enc.Encode(op.Subscriber)
	enc.Encode(op.Account)
	enc.Encode(op.Level)
	enc.Encode(op.Amount)
	enc.Encode(op.Period)
	enc.Encode(op.AutoRenewal)
	return enc.Err()
}
