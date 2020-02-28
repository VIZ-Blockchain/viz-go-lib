package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//ClaimInviteBalanceOperation represents claim_invite_balance operation data.
type ClaimInviteBalanceOperation struct {
	Initiator    string `json:"initiator"`
	Receiver     string `json:"receiver"`
	InviteSecret string `json:"invite_secret"`
}

//Type function that defines the type of operation ClaimInviteBalanceOperation.
func (op *ClaimInviteBalanceOperation) Type() OpType {
	return TypeClaimInviteBalance
}

//Data returns the operation data ClaimInviteBalanceOperation.
func (op *ClaimInviteBalanceOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ClaimInviteBalanceOperation to bytes.
func (op *ClaimInviteBalanceOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeClaimInviteBalance.Code()))
	enc.Encode(op.Initiator)
	enc.Encode(op.Receiver)
	enc.Encode(op.InviteSecret)
	return enc.Err()
}
