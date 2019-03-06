package types

//CommitteeCancelRequestOperation represents committee_cancel_request operation data.
type CommitteeCancelRequestOperation struct {
	RequestID uint32 `json:"request_id"`
}

//Type function that defines the type of operation CommitteeCancelRequestOperation.
func (op *CommitteeCancelRequestOperation) Type() OpType {
	return TypeCommitteeCancelRequest
}

//Data returns the operation data CommitteeCancelRequestOperation.
func (op *CommitteeCancelRequestOperation) Data() interface{} {
	return op
}
