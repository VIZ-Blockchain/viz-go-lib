package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//AccountValidatorProxyOperation represents account_validator_proxy operation data.
type AccountValidatorProxyOperation struct {
	Account string `json:"account"`
	Proxy   string `json:"proxy"`
}

//Type function that defines the type of operation AccountValidatorProxyOperation.
func (op *AccountValidatorProxyOperation) Type() OpType {
	return TypeAccountValidatorProxy
}

//Data returns the operation data AccountValidatorProxyOperation.
func (op *AccountValidatorProxyOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountValidatorProxyOperation to bytes.
func (op *AccountValidatorProxyOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountValidatorProxy.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Proxy)
	return enc.Err()
}
