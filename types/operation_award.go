package types

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//AwardOperation represents award operation data.
type AwardOperation struct {
	Initiator      string        `json:"initiator"`
	Receiver       string        `json:"receiver"`
	Energy         uint16        `json:"energy"`
	CustomSequence uint64        `json:"custom_sequence"`
	Memo           string        `json:"memo"`
	Beneficiaries  []Beneficiary `json:"beneficiaries"`
}

//Type function that defines the type of operation AwardOperation.
func (op *AwardOperation) Type() OpType {
	return TypeAward
}

//Data returns the operation data AwardOperation.
func (op *AwardOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AwardOperation to bytes.
func (op *AwardOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAward.Code()))
	enc.Encode(op.Initiator)
	enc.Encode(op.Receiver)
	enc.Encode(op.Energy)
	enc.Encode(op.CustomSequence)
	enc.Encode(op.Memo)
	enc.Encode(op.Beneficiaries)
	return enc.Err()
}
