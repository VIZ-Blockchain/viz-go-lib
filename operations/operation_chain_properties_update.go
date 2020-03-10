package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//ChainPropertiesUpdateOperation represents chain_properties_update operation data.
type ChainPropertiesUpdateOperation struct {
	Owner string                    `json:"owner"`
	Props *types.ChainPropertiesOLD `json:"props"`
}

//Type function that defines the type of operation ChainPropertiesUpdateOperation.
func (op *ChainPropertiesUpdateOperation) Type() OpType {
	return TypeChainPropertiesUpdate
}

//Data returns the operation data ChainPropertiesUpdateOperation.
func (op *ChainPropertiesUpdateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ChainPropertiesUpdateOperation to bytes.
func (op *ChainPropertiesUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChainPropertiesUpdate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.Props)
	return enc.Err()
}
