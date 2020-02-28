package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//CommitteeWorkerCancelRequestOperation represents committee_worker_cancel_request operation data.
type CommitteeWorkerCancelRequestOperation struct {
	Creator   string `json:"creator"`
	RequestID uint32 `json:"request_id"`
}

//Type function that defines the type of operation CommitteeWorkerCancelRequestOperation.
func (op *CommitteeWorkerCancelRequestOperation) Type() OpType {
	return TypeCommitteeWorkerCancelRequest
}

//Data returns the operation data CommitteeWorkerCancelRequestOperation.
func (op *CommitteeWorkerCancelRequestOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type CommitteeWorkerCancelRequestOperation to bytes.
func (op *CommitteeWorkerCancelRequestOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCommitteeWorkerCancelRequest.Code()))
	enc.Encode(op.Creator)
	enc.Encode(op.RequestID)
	return enc.Err()
}
