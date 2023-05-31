package operations

import (
	"github.com/biter777/viz-go-lib/encoding/transaction"
	"github.com/biter777/viz-go-lib/types"
)

// CommitteeWorkerCreateRequestOperation represents committee_worker_create_request operation data.
type CommitteeWorkerCreateRequestOperation struct {
	Creator           string       `json:"creator"`
	URL               string       `json:"url"`
	Worker            string       `json:"worker"`
	RequiredAmountMin *types.Asset `json:"required_amount_min"`
	RequiredAmountMax *types.Asset `json:"required_amount_max"`
	Duration          uint32       `json:"duration"`
}

// Type function that defines the type of operation CommitteeWorkerCreateRequestOperation.
func (op *CommitteeWorkerCreateRequestOperation) Type() OpType {
	return TypeCommitteeWorkerCreateRequest
}

// Data returns the operation data CommitteeWorkerCreateRequestOperation.
func (op *CommitteeWorkerCreateRequestOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type CommitteeWorkerCreateRequestOperation to bytes.
func (op *CommitteeWorkerCreateRequestOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCommitteeWorkerCreateRequest.Code()))
	enc.Encode(op.Creator)
	enc.Encode(op.URL)
	enc.Encode(op.Worker)
	enc.Encode(op.RequiredAmountMin)
	enc.Encode(op.RequiredAmountMax)
	enc.Encode(op.Duration)
	return enc.Err()
}
