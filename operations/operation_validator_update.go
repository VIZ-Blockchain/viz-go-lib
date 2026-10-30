package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//ValidatorUpdateOperation represents validator_update operation data.
type ValidatorUpdateOperation struct {
	Owner           string `json:"owner"`
	URL             string `json:"url"`
	BlockSigningKey string `json:"block_signing_key"`
}

//Type function that defines the type of operation ValidatorUpdateOperation.
func (op *ValidatorUpdateOperation) Type() OpType {
	return TypeValidatorUpdate
}

//Data returns the operation data ValidatorUpdateOperation.
func (op *ValidatorUpdateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ValidatorUpdateOperation to bytes.
func (op *ValidatorUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeValidatorUpdate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.URL)
	enc.EncodePubKey(op.BlockSigningKey)
	return enc.Err()
}
