package operations

//CommitteeApproveRequestOperation represents committee_approve_request operation data.
type CommitteeApproveRequestOperation struct {
	RequestID uint32 `json:"request_id"`
}

//Type function that defines the type of operation CommitteeApproveRequestOperation.
func (op *CommitteeApproveRequestOperation) Type() OpType {
	return TypeCommitteeApproveRequest
}

//Data returns the operation data CommitteeApproveRequestOperation.
func (op *CommitteeApproveRequestOperation) Data() interface{} {
	return op
}
