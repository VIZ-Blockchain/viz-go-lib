package types

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//CommitteeVoteRequestOperation represents committee_vote_request operation data.
type CommitteeVoteRequestOperation struct {
	Voter       string `json:"voter"`
	RequestID   uint32 `json:"request_id"`
	VotePercent int16  `json:"vote_percent"`
}

//Type function that defines the type of operation CommitteeVoteRequestOperation.
func (op *CommitteeVoteRequestOperation) Type() OpType {
	return TypeCommitteeVoteRequest
}

//Data returns the operation data CommitteeVoteRequestOperation.
func (op *CommitteeVoteRequestOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type CommitteeVoteRequestOperation to bytes.
func (op *CommitteeVoteRequestOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCommitteeVoteRequest.Code()))
	enc.Encode(op.Voter)
	enc.Encode(op.RequestID)
	enc.Encode(op.VotePercent)
	return enc.Err()
}
