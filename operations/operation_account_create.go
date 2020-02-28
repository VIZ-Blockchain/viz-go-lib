package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//AccountCreateOperation represents account_create operation data.
type AccountCreateOperation struct {
	Fee            types.Asset     `json:"fee"`
	Delegation     types.Asset     `json:"delegation"`
	Creator        string          `json:"creator"`
	NewAccountName string          `json:"new_account_name"`
	Master         types.Authority `json:"master"`
	Active         types.Authority `json:"active"`
	Regular        types.Authority `json:"regular"`
	MemoKey        string          `json:"memo_key"`
	JSONMetadata   string          `json:"json_metadata"`
	Referrer       string          `json:"referrer"`
	Extensions     []interface{}   `json:"extensions"`
}

//Type function that defines the type of operation AccountCreateOperation.
func (op *AccountCreateOperation) Type() OpType {
	return TypeAccountCreate
}

//Data returns the operation data AccountCreateOperation.
func (op *AccountCreateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountCreateOperation to bytes.
func (op *AccountCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountCreate.Code()))
	enc.Encode(op.Fee)
	enc.Encode(op.Delegation)
	enc.Encode(op.Creator)
	enc.Encode(op.NewAccountName)
	enc.Encode(op.Master)
	enc.Encode(op.Active)
	enc.Encode(op.Regular)
	enc.EncodePubKey(op.MemoKey)
	enc.Encode(op.JSONMetadata)
	enc.Encode(op.Referrer)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
