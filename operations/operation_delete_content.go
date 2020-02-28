package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//DeleteContentOperation represents delete_content operation data.
type DeleteContentOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//Type function that defines the type of operation DeleteContentOperation.
func (op *DeleteContentOperation) Type() OpType {
	return TypeDeleteContent
}

//Data returns the operation data DeleteContentOperation.
func (op *DeleteContentOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type DeleteContentOperation to bytes.
func (op *DeleteContentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeleteContent.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	return enc.Err()
}
