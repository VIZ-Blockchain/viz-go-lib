package types

//CommitteePayRequestOperation represents committee_pay_request operation data.
type CommitteePayRequestOperation struct {
	Worker    string `json:"worker"`
	RequestID uint32 `json:"request_id"`
	Tokens    *Asset `json:"tokens"`
}

//Type function that defines the type of operation CommitteePayRequestOperation.
func (op *CommitteePayRequestOperation) Type() OpType {
	return TypeCommitteePayRequest
}

//Data returns the operation data CommitteePayRequestOperation.
func (op *CommitteePayRequestOperation) Data() interface{} {
	return op
}
