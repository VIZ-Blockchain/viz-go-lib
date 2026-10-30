package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

// ExpireEscrowRatificationOperation represents expire_escrow_ratification operation data.
type ExpireEscrowRatificationOperation struct {
	From                 string       `json:"from"`
	To                   string       `json:"to"`
	Agent                string       `json:"agent"`
	EscrowID             uint32       `json:"escrow_id"`
	TokenAmount          *types.Asset `json:"token_amount"`
	Fee                  *types.Asset `json:"fee"`
	RatificationDeadline *types.Time  `json:"ratification_deadline"`
}

// Type function that defines the type of operation ExpireEscrowRatificationOperation.
func (op *ExpireEscrowRatificationOperation) Type() OpType {
	return TypeExpireEscrowRatification
}

// Data returns the operation data ExpireEscrowRatificationOperation.
func (op *ExpireEscrowRatificationOperation) Data() interface{} {
	return op
}

// BidOperation represents bid operation data.
type BidOperation struct {
	Account string       `json:"account"`
	Bidder  string       `json:"bidder"`
	Bid     *types.Asset `json:"bid"`
}

// Type function that defines the type of operation BidOperation.
func (op *BidOperation) Type() OpType {
	return TypeBid
}

// Data returns the operation data BidOperation.
func (op *BidOperation) Data() interface{} {
	return op
}

// OutbidOperation represents outbid operation data.
type OutbidOperation struct {
	Account string       `json:"account"`
	Bidder  string       `json:"bidder"`
	Bid     *types.Asset `json:"bid"`
}

// Type function that defines the type of operation OutbidOperation.
func (op *OutbidOperation) Type() OpType {
	return TypeOutbid
}

// Data returns the operation data OutbidOperation.
func (op *OutbidOperation) Data() interface{} {
	return op
}

// StakeholderRewardOperation represents stakeholder_reward operation data.
type StakeholderRewardOperation struct {
	Validator   string       `json:"validator"`
	Stakeholder string       `json:"stakeholder"`
	Shares      *types.Asset `json:"shares"`
}

// Type function that defines the type of operation StakeholderRewardOperation.
func (op *StakeholderRewardOperation) Type() OpType {
	return TypeStakeholderReward
}

// Data returns the operation data StakeholderRewardOperation.
func (op *StakeholderRewardOperation) Data() interface{} {
	return op
}
