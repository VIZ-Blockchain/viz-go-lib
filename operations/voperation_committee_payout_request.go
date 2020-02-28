package operations

//CommitteePayoutRequestOperation represents committee_payout_request operation data.
type CommitteePayoutRequestOperation struct {
	RequestID uint32 `json:"request_id"`
}

//Type function that defines the type of operation CommitteePayoutRequestOperation.
func (op *CommitteePayoutRequestOperation) Type() OpType {
	return TypeCommitteePayoutRequest
}

//Data returns the operation data CommitteePayoutRequestOperation.
func (op *CommitteePayoutRequestOperation) Data() interface{} {
	return op
}
