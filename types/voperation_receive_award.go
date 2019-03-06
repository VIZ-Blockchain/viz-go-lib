package types

//ReceiveAwardOperation represents receive_award operation data.
type ReceiveAwardOperation struct {
	Receiver       string `json:"receiver"`
	CustomSequence uint64 `json:"custom_sequence"`
	Memo           string `json:"memo"`
	Shares         *Asset `json:"shares"`
}

//Type function that defines the type of operation ReceiveAwardOperation.
func (op *ReceiveAwardOperation) Type() OpType {
	return TypeReceiveAward
}

//Data returns the operation data ReceiveAwardOperation.
func (op *ReceiveAwardOperation) Data() interface{} {
	return op
}
