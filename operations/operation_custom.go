package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//CustomOperation represents custom operation data.
type CustomOperation struct {
	RequiredActiveAuths  []string `json:"required_active_auths"`
	RequiredRegularAuths []string `json:"required_regular_auths"`
	ID                   string   `json:"id"`
	JSON                 string   `json:"json"`
}

//Type function that defines the type of operation.
func (op *CustomOperation) Type() OpType {
	return TypeCustom
}

//Data returns the operation data.
func (op *CustomOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type CustomOperation to bytes.
func (op *CustomOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCustom.Code()))
	enc.EncodeArrString(op.RequiredActiveAuths)
	enc.EncodeArrString(op.RequiredRegularAuths)
	enc.Encode(op.ID)
	enc.Encode(op.JSON)
	return enc.Err()
}
