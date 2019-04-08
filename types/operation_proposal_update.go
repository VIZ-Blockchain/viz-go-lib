package types

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//ProposalUpdateOperation represents proposal_update operation data.
type ProposalUpdateOperation struct {
	Author                   string        `json:"author"`
	Title                    string        `json:"title"`
	ActiveApprovalsToAdd     []string      `json:"active_approvals_to_add"`
	ActiveApprovalsToRemove  []string      `json:"active_approvals_to_remove"`
	OwnerApprovalsToAdd      []string      `json:"owner_approvals_to_add"`
	OwnerApprovalsToRemove   []string      `json:"owner_approvals_to_remove"`
	RegularApprovalsToAdd    []string      `json:"regular_approvals_to_add"`
	RegularApprovalsToRemove []string      `json:"regular_approvals_to_remove"`
	KeyApprovalsToAdd        []string      `json:"key_approvals_to_add"`
	KeyApprovalsToRemove     []string      `json:"key_approvals_to_remove"`
	Extensions               []interface{} `json:"extensions"`
}

//Type function that defines the type of operation ProposalUpdateOperation.
func (op *ProposalUpdateOperation) Type() OpType {
	return TypeProposalUpdate
}

//Data returns the operation data ProposalUpdateOperation.
func (op *ProposalUpdateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ProposalUpdateOperation to bytes.
func (op *ProposalUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeProposalUpdate.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Title)
	enc.EncodeArrString(op.ActiveApprovalsToAdd)
	enc.EncodeArrString(op.ActiveApprovalsToRemove)
	enc.EncodeArrString(op.OwnerApprovalsToAdd)
	enc.EncodeArrString(op.OwnerApprovalsToRemove)
	enc.EncodeArrString(op.RegularApprovalsToAdd)
	enc.EncodeArrString(op.RegularApprovalsToRemove)
	enc.EncodeArrString(op.KeyApprovalsToAdd)
	enc.EncodeArrString(op.KeyApprovalsToRemove)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
