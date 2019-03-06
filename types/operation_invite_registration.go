package types

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//InviteRegistrationOperation represents invite_registration operation data.
type InviteRegistrationOperation struct {
	Initiator      string `json:"initiator"`
	NewAccountName string `json:"new_account_name"`
	InviteSecret   string `json:"invite_secret"`
	NewAccountKey  string `json:"new_account_key"`
}

//Type function that defines the type of operation InviteRegistrationOperation.
func (op *InviteRegistrationOperation) Type() OpType {
	return TypeInviteRegistration
}

//Data returns the operation data InviteRegistrationOperation.
func (op *InviteRegistrationOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type InviteRegistrationOperation to bytes.
func (op *InviteRegistrationOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeInviteRegistration.Code()))
	enc.Encode(op.Initiator)
	enc.Encode(op.NewAccountName)
	enc.Encode(op.InviteSecret)
	enc.EncodePubKey(op.NewAccountKey)
	return enc.Err()
}
