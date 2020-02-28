package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//ReceiveAwardOperation represents receive_award operation data.
type ReceiveAwardOperation struct {
	Initiator      string      `json:"initiator"`
	Receiver       string      `json:"receiver"`
	CustomSequence uint64      `json:"custom_sequence"`
	Memo           string      `json:"memo"`
	Shares         types.Asset `json:"shares"`
}

//Type function that defines the type of operation ReceiveAwardOperation.
func (op *ReceiveAwardOperation) Type() OpType {
	return TypeReceiveAward
}

//Data returns the operation data ReceiveAwardOperation.
func (op *ReceiveAwardOperation) Data() interface{} {
	return op
}
