package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//RecoverAccountOperation represents recover_account operation data.
type RecoverAccountOperation struct {
	AccountToRecover      string           `json:"account_to_recover"`
	NewMasterAuthority    *types.Authority `json:"new_master_authority"`
	RecentMasterAuthority *types.Authority `json:"recent_master_authority"`
	Extensions            []interface{}    `json:"extensions"`
}

//Type function that defines the type of operation RecoverAccountOperation.
func (op *RecoverAccountOperation) Type() OpType {
	return TypeRecoverAccount
}

//Data returns the operation data RecoverAccountOperation.
func (op *RecoverAccountOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type RecoverAccountOperation to bytes.
func (op *RecoverAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeRecoverAccount.Code()))
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewMasterAuthority)
	enc.Encode(op.RecentMasterAuthority)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
