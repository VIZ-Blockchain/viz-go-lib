package types

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//VersionedChainPropertiesUpdateOperation represents versioned_chain_properties_update operation data.
type VersionedChainPropertiesUpdateOperation struct {
	Owner string           `json:"owner"`
	Props *ChainProperties `json:"props"`
}

//Type function that defines the type of operation VersionedChainPropertiesUpdateOperation.
func (op *VersionedChainPropertiesUpdateOperation) Type() OpType {
	return TypeVersionedChainPropertiesUpdate
}

//Data returns the operation data VersionedChainPropertiesUpdateOperation.
func (op *VersionedChainPropertiesUpdateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type VersionedChainPropertiesUpdateOperation to bytes.
func (op *VersionedChainPropertiesUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeVersionedChainPropertiesUpdate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.Props)
	return enc.Err()
}
