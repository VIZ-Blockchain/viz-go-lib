package operations

import (
	"github.com/biter777/viz-go-lib/encoding/transaction"
	"github.com/biter777/viz-go-lib/types"
)

// CreateInviteOperation represents create_invite operation data.
type CreateInviteOperation struct {
	Creator   string       `json:"creator"`
	Balance   *types.Asset `json:"balance"`
	InviteKey string       `json:"invite_key"`
}

// Type function that defines the type of operation CreateInviteOperation.
func (op *CreateInviteOperation) Type() OpType {
	return TypeCreateInvite
}

// Data returns the operation data CreateInviteOperation.
func (op *CreateInviteOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type CreateInviteOperation to bytes.
func (op *CreateInviteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCreateInvite.Code()))
	enc.Encode(op.Creator)
	enc.Encode(op.Balance)
	enc.Encode(op.InviteKey)
	return enc.Err()
}
