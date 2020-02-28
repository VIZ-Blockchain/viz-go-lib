package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//AccountUpdateOperation represents account_update operation data.
type AccountUpdateOperation struct {
	Account      string           `json:"account"`
	Master       *types.Authority `json:"master,omitempty"`
	Active       *types.Authority `json:"active,omitempty"`
	Regular      *types.Authority `json:"regular,omitempty"`
	MemoKey      string           `json:"memo_key"`
	JSONMetadata string           `json:"json_metadata"`
}

//Type function that defines the type of operation AccountUpdateOperation.
func (op *AccountUpdateOperation) Type() OpType {
	return TypeAccountUpdate
}

//Data returns the operation data AccountUpdateOperation.
func (op *AccountUpdateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountUpdateOperation to bytes.
func (op *AccountUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountUpdate.Code()))
	enc.EncodeString(op.Account)
	if op.Master != nil {
		enc.Encode(byte(1))
		enc.Encode(op.Master)
	} else {
		enc.Encode(byte(0))
	}
	if op.Active != nil {
		enc.Encode(byte(1))
		enc.Encode(op.Active)
	} else {
		enc.Encode(byte(0))
	}
	if op.Regular != nil {
		enc.Encode(byte(1))
		enc.Encode(op.Regular)
	} else {
		enc.Encode(byte(0))
	}
	enc.EncodePubKey(op.MemoKey)
	enc.Encode(op.JSONMetadata)
	return enc.Err()
}
