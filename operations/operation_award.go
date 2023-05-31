package operations

import (
	"github.com/biter777/viz-go-lib/encoding/transaction"
	"github.com/biter777/viz-go-lib/types"
)

// AwardOperation represents award operation data.
type AwardOperation struct {
	Initiator      string              `json:"initiator"`
	Receiver       string              `json:"receiver"`
	Energy         uint16              `json:"energy"`
	CustomSequence uint64              `json:"custom_sequence"`
	Memo           string              `json:"memo"`
	Beneficiaries  []types.Beneficiary `json:"beneficiaries"`
}

// Type function that defines the type of operation AwardOperation.
func (op *AwardOperation) Type() OpType {
	return TypeAward
}

// Data returns the operation data AwardOperation.
func (op *AwardOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type AwardOperation to bytes.
func (op *AwardOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAward.Code()))
	enc.Encode(op.Initiator)
	enc.Encode(op.Receiver)
	enc.Encode(op.Energy)
	enc.Encode(op.CustomSequence)
	enc.Encode(op.Memo)
	if len(op.Beneficiaries) > 0 {
		enc.EncodeUVarint(uint64(len(op.Beneficiaries)))
		for _, val := range op.Beneficiaries {
			enc.Encode(val.Account)
			enc.Encode(val.Weight)
		}
	} else {
		enc.Encode(byte(0))
	}
	return enc.Err()
}
