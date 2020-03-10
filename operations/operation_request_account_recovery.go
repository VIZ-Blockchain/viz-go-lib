package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//RequestAccountRecoveryOperation represents request_account_recovery operation data.
type RequestAccountRecoveryOperation struct {
	RecoveryAccount    string           `json:"recovery_account"`
	AccountToRecover   string           `json:"account_to_recover"`
	NewMasterAuthority *types.Authority `json:"new_master_authority"`
	Extensions         []interface{}    `json:"extensions"`
}

//Type function that defines the type of operation RequestAccountRecoveryOperation.
func (op *RequestAccountRecoveryOperation) Type() OpType {
	return TypeRequestAccountRecovery
}

//Data returns the operation data RequestAccountRecoveryOperation.
func (op *RequestAccountRecoveryOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type RequestAccountRecoveryOperation to bytes.
func (op *RequestAccountRecoveryOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeRequestAccountRecovery.Code()))
	enc.Encode(op.RecoveryAccount)
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewMasterAuthority)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
