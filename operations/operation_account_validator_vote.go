package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//AccountValidatorVoteOperation represents account_validator_vote operation data.
type AccountValidatorVoteOperation struct {
	Account string `json:"account"`
	Validator string `json:"validator"`
	Approve bool   `json:"approve"`
}

//Type function that defines the type of operation AccountValidatorVoteOperation.
func (op *AccountValidatorVoteOperation) Type() OpType {
	return TypeAccountValidatorVote
}

//Data returns the operation data AccountValidatorVoteOperation.
func (op *AccountValidatorVoteOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountValidatorVoteOperation to bytes.
func (op *AccountValidatorVoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountValidatorVote.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Validator)
	enc.Encode(op.Approve)
	return enc.Err()
}
