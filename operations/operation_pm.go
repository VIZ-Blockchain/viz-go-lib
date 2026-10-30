// Hard fork 14 prediction-market and hard fork 15 agent-access operations.
//
// Field order mirrors FC_REFLECT in viz-cpp-node
// (libraries/protocol/include/graphene/protocol/pm_operations.hpp and agent_operations.hpp).
// That order is the consensus wire format and must not be changed.

package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

// PMOracleRegisterOperation represents pm_oracle_register operation data.
type PMOracleRegisterOperation struct {
	Owner              string       `json:"owner"`
	Insurance          *types.Asset `json:"insurance"`
	FeePercent         uint16       `json:"fee_percent"`
	FixedFee           *types.Asset `json:"fixed_fee"`
	RulesURL           string       `json:"rules_url"`
	AutoAcceptCreator  string       `json:"auto_accept_creator"`
	AutoAcceptResolver string       `json:"auto_accept_resolver"`
	AutoAccept         bool         `json:"auto_accept"`
}

// Type function that defines the type of operation PMOracleRegisterOperation.
func (op *PMOracleRegisterOperation) Type() OpType {
	return TypePMOracleRegister
}

// Data returns the operation data PMOracleRegisterOperation.
func (op *PMOracleRegisterOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMOracleRegisterOperation to bytes.
func (op *PMOracleRegisterOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMOracleRegister.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.Insurance)
	enc.Encode(op.FeePercent)
	enc.Encode(op.FixedFee)
	enc.Encode(op.RulesURL)
	enc.Encode(op.AutoAcceptCreator)
	enc.Encode(op.AutoAcceptResolver)
	enc.Encode(op.AutoAccept)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMOracleUpdateOperation represents pm_oracle_update operation data.
type PMOracleUpdateOperation struct {
	Owner              string       `json:"owner"`
	InsuranceDelta     *types.Asset `json:"insurance_delta,omitempty"`
	FeePercent         *uint16      `json:"fee_percent,omitempty"`
	FixedFee           *types.Asset `json:"fixed_fee,omitempty"`
	RulesURL           *string      `json:"rules_url,omitempty"`
	AutoAcceptCreator  *string      `json:"auto_accept_creator,omitempty"`
	AutoAcceptResolver *string      `json:"auto_accept_resolver,omitempty"`
	AutoAccept         *bool        `json:"auto_accept,omitempty"`
}

// Type function that defines the type of operation PMOracleUpdateOperation.
func (op *PMOracleUpdateOperation) Type() OpType {
	return TypePMOracleUpdate
}

// Data returns the operation data PMOracleUpdateOperation.
func (op *PMOracleUpdateOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMOracleUpdateOperation to bytes.
func (op *PMOracleUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMOracleUpdate.Code()))
	enc.Encode(op.Owner)
	if op.InsuranceDelta != nil {
		enc.EncodeNumber(byte(1))
		enc.Encode(op.InsuranceDelta)
	} else {
		enc.EncodeNumber(byte(0))
	}
	if op.FeePercent != nil {
		enc.EncodeNumber(byte(1))
		enc.Encode(*op.FeePercent)
	} else {
		enc.EncodeNumber(byte(0))
	}
	if op.FixedFee != nil {
		enc.EncodeNumber(byte(1))
		enc.Encode(op.FixedFee)
	} else {
		enc.EncodeNumber(byte(0))
	}
	if op.RulesURL != nil {
		enc.EncodeNumber(byte(1))
		enc.Encode(*op.RulesURL)
	} else {
		enc.EncodeNumber(byte(0))
	}
	if op.AutoAcceptCreator != nil {
		enc.EncodeNumber(byte(1))
		enc.Encode(*op.AutoAcceptCreator)
	} else {
		enc.EncodeNumber(byte(0))
	}
	if op.AutoAcceptResolver != nil {
		enc.EncodeNumber(byte(1))
		enc.Encode(*op.AutoAcceptResolver)
	} else {
		enc.EncodeNumber(byte(0))
	}
	if op.AutoAccept != nil {
		enc.EncodeNumber(byte(1))
		enc.Encode(*op.AutoAccept)
	} else {
		enc.EncodeNumber(byte(0))
	}
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMCreateMarketOperation represents pm_create_market operation data.
type PMCreateMarketOperation struct {
	Creator               string       `json:"creator"`
	Oracle                string       `json:"oracle"`
	MarketType            uint8        `json:"market_type"`
	Outcomes              []string     `json:"outcomes"`
	URL                   string       `json:"url"`
	OracleFeePercent      uint16       `json:"oracle_fee_percent"`
	OracleFixedFee        *types.Asset `json:"oracle_fixed_fee"`
	CreatorFeePercent     uint16       `json:"creator_fee_percent"`
	LiquidityFeePercent   uint16       `json:"liquidity_fee_percent"`
	Liquidity             *types.Asset `json:"liquidity"`
	LMSRB                 int64        `json:"lmsr_b"`
	BettingExpiration     *types.Time  `json:"betting_expiration"`
	ResultExpiration      *types.Time  `json:"result_expiration"`
	TimePenaltyType       uint8        `json:"time_penalty_type"`
	TimePenaltyValue      uint32       `json:"time_penalty_value"`
	PenaltyCurveType      uint8        `json:"penalty_curve_type"`
	AllowEarlyResolution  bool         `json:"allow_early_resolution"`
	AllowCancellation     bool         `json:"allow_cancellation"`
	AllowBatch            bool         `json:"allow_batch"`
	AllowInstantBet       bool         `json:"allow_instant_bet"`
	EndogeneityTier       uint8        `json:"endogeneity_tier"`
	DisputeMode           uint8        `json:"dispute_mode"`
	DisputeResolver       string       `json:"dispute_resolver"`
	DisputePenaltyPercent int16        `json:"dispute_penalty_percent"`
	Metadata              string       `json:"metadata"`
}

// Type function that defines the type of operation PMCreateMarketOperation.
func (op *PMCreateMarketOperation) Type() OpType {
	return TypePMCreateMarket
}

// Data returns the operation data PMCreateMarketOperation.
func (op *PMCreateMarketOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMCreateMarketOperation to bytes.
func (op *PMCreateMarketOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMCreateMarket.Code()))
	enc.Encode(op.Creator)
	enc.Encode(op.Oracle)
	enc.Encode(op.MarketType)
	enc.EncodeArrString(op.Outcomes)
	enc.Encode(op.URL)
	enc.Encode(op.OracleFeePercent)
	enc.Encode(op.OracleFixedFee)
	enc.Encode(op.CreatorFeePercent)
	enc.Encode(op.LiquidityFeePercent)
	enc.Encode(op.Liquidity)
	enc.Encode(op.LMSRB)
	enc.Encode(op.BettingExpiration)
	enc.Encode(op.ResultExpiration)
	enc.Encode(op.TimePenaltyType)
	enc.Encode(op.TimePenaltyValue)
	enc.Encode(op.PenaltyCurveType)
	enc.Encode(op.AllowEarlyResolution)
	enc.Encode(op.AllowCancellation)
	enc.Encode(op.AllowBatch)
	enc.Encode(op.AllowInstantBet)
	enc.Encode(op.EndogeneityTier)
	enc.Encode(op.DisputeMode)
	enc.Encode(op.DisputeResolver)
	enc.Encode(op.DisputePenaltyPercent)
	enc.Encode(op.Metadata)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMOracleAcceptMarketOperation represents pm_oracle_accept_market operation data.
type PMOracleAcceptMarketOperation struct {
	Oracle           string       `json:"oracle"`
	MarketID         int64        `json:"market_id"`
	Accept           bool         `json:"accept"`
	OracleFeePercent uint16       `json:"oracle_fee_percent"`
	OracleFixedFee   *types.Asset `json:"oracle_fixed_fee"`
}

// Type function that defines the type of operation PMOracleAcceptMarketOperation.
func (op *PMOracleAcceptMarketOperation) Type() OpType {
	return TypePMOracleAcceptMarket
}

// Data returns the operation data PMOracleAcceptMarketOperation.
func (op *PMOracleAcceptMarketOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMOracleAcceptMarketOperation to bytes.
func (op *PMOracleAcceptMarketOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMOracleAcceptMarket.Code()))
	enc.Encode(op.Oracle)
	enc.Encode(op.MarketID)
	enc.Encode(op.Accept)
	enc.Encode(op.OracleFeePercent)
	enc.Encode(op.OracleFixedFee)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMPlaceBetOperation represents pm_place_bet operation data.
type PMPlaceBetOperation struct {
	Account      string       `json:"account"`
	MarketID     int64        `json:"market_id"`
	Side         int8         `json:"side"`
	OutcomeIndex int16        `json:"outcome_index"`
	Amount       *types.Asset `json:"amount"`
	MinTokens    int64        `json:"min_tokens"`
	Mode         uint8        `json:"mode"`
}

// Type function that defines the type of operation PMPlaceBetOperation.
func (op *PMPlaceBetOperation) Type() OpType {
	return TypePMPlaceBet
}

// Data returns the operation data PMPlaceBetOperation.
func (op *PMPlaceBetOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMPlaceBetOperation to bytes.
func (op *PMPlaceBetOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMPlaceBet.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.MarketID)
	enc.Encode(op.Side)
	enc.Encode(op.OutcomeIndex)
	enc.Encode(op.Amount)
	enc.Encode(op.MinTokens)
	enc.Encode(op.Mode)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMCommitBetOperation represents pm_commit_bet operation data.
type PMCommitBetOperation struct {
	Account            string       `json:"account"`
	MarketID           int64        `json:"market_id"`
	Commitment         string       `json:"commitment"`
	EscrowAmount       *types.Asset `json:"escrow_amount"`
	NoRevealFeePercent uint16       `json:"no_reveal_fee_percent"`
}

// Type function that defines the type of operation PMCommitBetOperation.
func (op *PMCommitBetOperation) Type() OpType {
	return TypePMCommitBet
}

// Data returns the operation data PMCommitBetOperation.
func (op *PMCommitBetOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMCommitBetOperation to bytes.
func (op *PMCommitBetOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMCommitBet.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.MarketID)
	enc.EncodeSHA256(op.Commitment)
	enc.Encode(op.EscrowAmount)
	enc.Encode(op.NoRevealFeePercent)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMRevealBetOperation represents pm_reveal_bet operation data.
type PMRevealBetOperation struct {
	Account      string       `json:"account"`
	CommitID     int64        `json:"commit_id"`
	Side         int8         `json:"side"`
	OutcomeIndex int16        `json:"outcome_index"`
	Amount       *types.Asset `json:"amount"`
	Salt         string       `json:"salt"`
	MinTokens    int64        `json:"min_tokens"`
}

// Type function that defines the type of operation PMRevealBetOperation.
func (op *PMRevealBetOperation) Type() OpType {
	return TypePMRevealBet
}

// Data returns the operation data PMRevealBetOperation.
func (op *PMRevealBetOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMRevealBetOperation to bytes.
func (op *PMRevealBetOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMRevealBet.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.CommitID)
	enc.Encode(op.Side)
	enc.Encode(op.OutcomeIndex)
	enc.Encode(op.Amount)
	enc.Encode(op.Salt)
	enc.Encode(op.MinTokens)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMCancelBetOperation represents pm_cancel_bet operation data.
type PMCancelBetOperation struct {
	Account   string `json:"account"`
	BetID     int64  `json:"bet_id"`
	MinReturn int64  `json:"min_return"`
}

// Type function that defines the type of operation PMCancelBetOperation.
func (op *PMCancelBetOperation) Type() OpType {
	return TypePMCancelBet
}

// Data returns the operation data PMCancelBetOperation.
func (op *PMCancelBetOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMCancelBetOperation to bytes.
func (op *PMCancelBetOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMCancelBet.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.BetID)
	enc.Encode(op.MinReturn)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMAddLiquidityOperation represents pm_add_liquidity operation data.
type PMAddLiquidityOperation struct {
	Provider string       `json:"provider"`
	MarketID int64        `json:"market_id"`
	Amount   *types.Asset `json:"amount"`
}

// Type function that defines the type of operation PMAddLiquidityOperation.
func (op *PMAddLiquidityOperation) Type() OpType {
	return TypePMAddLiquidity
}

// Data returns the operation data PMAddLiquidityOperation.
func (op *PMAddLiquidityOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMAddLiquidityOperation to bytes.
func (op *PMAddLiquidityOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMAddLiquidity.Code()))
	enc.Encode(op.Provider)
	enc.Encode(op.MarketID)
	enc.Encode(op.Amount)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMWithdrawLiquidityOperation represents pm_withdraw_liquidity operation data.
type PMWithdrawLiquidityOperation struct {
	Provider    string       `json:"provider"`
	LiquidityID int64        `json:"liquidity_id"`
	Amount      *types.Asset `json:"amount"`
}

// Type function that defines the type of operation PMWithdrawLiquidityOperation.
func (op *PMWithdrawLiquidityOperation) Type() OpType {
	return TypePMWithdrawLiquidity
}

// Data returns the operation data PMWithdrawLiquidityOperation.
func (op *PMWithdrawLiquidityOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMWithdrawLiquidityOperation to bytes.
func (op *PMWithdrawLiquidityOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMWithdrawLiquidity.Code()))
	enc.Encode(op.Provider)
	enc.Encode(op.LiquidityID)
	enc.Encode(op.Amount)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMResolveMarketOperation represents pm_resolve_market operation data.
type PMResolveMarketOperation struct {
	Oracle         string `json:"oracle"`
	MarketID       int64  `json:"market_id"`
	WinningOutcome int16  `json:"winning_outcome"`
	DecisionURL    string `json:"decision_url"`
	DecisionReason string `json:"decision_reason"`
}

// Type function that defines the type of operation PMResolveMarketOperation.
func (op *PMResolveMarketOperation) Type() OpType {
	return TypePMResolveMarket
}

// Data returns the operation data PMResolveMarketOperation.
func (op *PMResolveMarketOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMResolveMarketOperation to bytes.
func (op *PMResolveMarketOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMResolveMarket.Code()))
	enc.Encode(op.Oracle)
	enc.Encode(op.MarketID)
	enc.Encode(op.WinningOutcome)
	enc.Encode(op.DecisionURL)
	enc.Encode(op.DecisionReason)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMNoContestOperation represents pm_no_contest operation data.
type PMNoContestOperation struct {
	Oracle   string `json:"oracle"`
	MarketID int64  `json:"market_id"`
	Reason   string `json:"reason"`
}

// Type function that defines the type of operation PMNoContestOperation.
func (op *PMNoContestOperation) Type() OpType {
	return TypePMNoContest
}

// Data returns the operation data PMNoContestOperation.
func (op *PMNoContestOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMNoContestOperation to bytes.
func (op *PMNoContestOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMNoContest.Code()))
	enc.Encode(op.Oracle)
	enc.Encode(op.MarketID)
	enc.Encode(op.Reason)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMDisputeCreateOperation represents pm_dispute_create operation data.
type PMDisputeCreateOperation struct {
	Disputer        string `json:"disputer"`
	MarketID        int64  `json:"market_id"`
	ProposedOutcome int16  `json:"proposed_outcome"`
	Reason          string `json:"reason"`
}

// Type function that defines the type of operation PMDisputeCreateOperation.
func (op *PMDisputeCreateOperation) Type() OpType {
	return TypePMDisputeCreate
}

// Data returns the operation data PMDisputeCreateOperation.
func (op *PMDisputeCreateOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMDisputeCreateOperation to bytes.
func (op *PMDisputeCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMDisputeCreate.Code()))
	enc.Encode(op.Disputer)
	enc.Encode(op.MarketID)
	enc.Encode(op.ProposedOutcome)
	enc.Encode(op.Reason)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMDisputeVoteOperation represents pm_dispute_vote operation data.
type PMDisputeVoteOperation struct {
	Voter       string `json:"voter"`
	MarketID    int64  `json:"market_id"`
	VoteOutcome int16  `json:"vote_outcome"`
	VotePercent int16  `json:"vote_percent"`
}

// Type function that defines the type of operation PMDisputeVoteOperation.
func (op *PMDisputeVoteOperation) Type() OpType {
	return TypePMDisputeVote
}

// Data returns the operation data PMDisputeVoteOperation.
func (op *PMDisputeVoteOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMDisputeVoteOperation to bytes.
func (op *PMDisputeVoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMDisputeVote.Code()))
	enc.Encode(op.Voter)
	enc.Encode(op.MarketID)
	enc.Encode(op.VoteOutcome)
	enc.Encode(op.VotePercent)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMDisputeResolveOperation represents pm_dispute_resolve operation data.
type PMDisputeResolveOperation struct {
	Resolver        string       `json:"resolver"`
	MarketID        int64        `json:"market_id"`
	CorrectOutcome  int16        `json:"correct_outcome"`
	PenaltyAmount   *types.Asset `json:"penalty_amount"`
	BanOracle       bool         `json:"ban_oracle"`
	BanOracleUntil  *types.Time  `json:"ban_oracle_until"`
	BanCreator      bool         `json:"ban_creator"`
	BanCreatorUntil *types.Time  `json:"ban_creator_until"`
}

// Type function that defines the type of operation PMDisputeResolveOperation.
func (op *PMDisputeResolveOperation) Type() OpType {
	return TypePMDisputeResolve
}

// Data returns the operation data PMDisputeResolveOperation.
func (op *PMDisputeResolveOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMDisputeResolveOperation to bytes.
func (op *PMDisputeResolveOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMDisputeResolve.Code()))
	enc.Encode(op.Resolver)
	enc.Encode(op.MarketID)
	enc.Encode(op.CorrectOutcome)
	enc.Encode(op.PenaltyAmount)
	enc.Encode(op.BanOracle)
	enc.Encode(op.BanOracleUntil)
	enc.Encode(op.BanCreator)
	enc.Encode(op.BanCreatorUntil)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMTransferPositionOperation represents pm_transfer_position operation data.
type PMTransferPositionOperation struct {
	From   string `json:"from"`
	BetID  int64  `json:"bet_id"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
	Memo   string `json:"memo"`
}

// Type function that defines the type of operation PMTransferPositionOperation.
func (op *PMTransferPositionOperation) Type() OpType {
	return TypePMTransferPosition
}

// Data returns the operation data PMTransferPositionOperation.
func (op *PMTransferPositionOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMTransferPositionOperation to bytes.
func (op *PMTransferPositionOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMTransferPosition.Code()))
	enc.Encode(op.From)
	enc.Encode(op.BetID)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMLazyDepositOperation represents pm_lazy_deposit operation data.
type PMLazyDepositOperation struct {
	Account string       `json:"account"`
	Amount  *types.Asset `json:"amount"`
}

// Type function that defines the type of operation PMLazyDepositOperation.
func (op *PMLazyDepositOperation) Type() OpType {
	return TypePMLazyDeposit
}

// Data returns the operation data PMLazyDepositOperation.
func (op *PMLazyDepositOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMLazyDepositOperation to bytes.
func (op *PMLazyDepositOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMLazyDeposit.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Amount)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMLazyWithdrawOperation represents pm_lazy_withdraw operation data.
type PMLazyWithdrawOperation struct {
	Account   string `json:"account"`
	Shares    int64  `json:"shares"`
	Emergency bool   `json:"emergency"`
}

// Type function that defines the type of operation PMLazyWithdrawOperation.
func (op *PMLazyWithdrawOperation) Type() OpType {
	return TypePMLazyWithdraw
}

// Data returns the operation data PMLazyWithdrawOperation.
func (op *PMLazyWithdrawOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMLazyWithdrawOperation to bytes.
func (op *PMLazyWithdrawOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMLazyWithdraw.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Shares)
	enc.Encode(op.Emergency)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMLeverageOpenOperation represents pm_leverage_open operation data.
type PMLeverageOpenOperation struct {
	Account            string       `json:"account"`
	MarketID           int64        `json:"market_id"`
	OutcomeIndex       int16        `json:"outcome_index"`
	Collateral         *types.Asset `json:"collateral"`
	Loan               *types.Asset `json:"loan"`
	MinTokens          int64        `json:"min_tokens"`
	MaxSlippagePercent uint16       `json:"max_slippage_percent"`
}

// Type function that defines the type of operation PMLeverageOpenOperation.
func (op *PMLeverageOpenOperation) Type() OpType {
	return TypePMLeverageOpen
}

// Data returns the operation data PMLeverageOpenOperation.
func (op *PMLeverageOpenOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMLeverageOpenOperation to bytes.
func (op *PMLeverageOpenOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMLeverageOpen.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.MarketID)
	enc.Encode(op.OutcomeIndex)
	enc.Encode(op.Collateral)
	enc.Encode(op.Loan)
	enc.Encode(op.MinTokens)
	enc.Encode(op.MaxSlippagePercent)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMLeverageCloseOperation represents pm_leverage_close operation data.
type PMLeverageCloseOperation struct {
	Account    string `json:"account"`
	PositionID int64  `json:"position_id"`
	MinReturn  int64  `json:"min_return"`
}

// Type function that defines the type of operation PMLeverageCloseOperation.
func (op *PMLeverageCloseOperation) Type() OpType {
	return TypePMLeverageClose
}

// Data returns the operation data PMLeverageCloseOperation.
func (op *PMLeverageCloseOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMLeverageCloseOperation to bytes.
func (op *PMLeverageCloseOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMLeverageClose.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.PositionID)
	enc.Encode(op.MinReturn)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMLeverageConvertOperation represents pm_leverage_convert operation data.
type PMLeverageConvertOperation struct {
	Account              string `json:"account"`
	PositionID           int64  `json:"position_id"`
	ConversionProfitCost uint16 `json:"conversion_profit_cost"`
}

// Type function that defines the type of operation PMLeverageConvertOperation.
func (op *PMLeverageConvertOperation) Type() OpType {
	return TypePMLeverageConvert
}

// Data returns the operation data PMLeverageConvertOperation.
func (op *PMLeverageConvertOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMLeverageConvertOperation to bytes.
func (op *PMLeverageConvertOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMLeverageConvert.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.PositionID)
	enc.Encode(op.ConversionProfitCost)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMDisputeOracleRespondOperation represents pm_dispute_oracle_respond operation data.
type PMDisputeOracleRespondOperation struct {
	Oracle   string `json:"oracle"`
	MarketID int64  `json:"market_id"`
	Response string `json:"response"`
}

// Type function that defines the type of operation PMDisputeOracleRespondOperation.
func (op *PMDisputeOracleRespondOperation) Type() OpType {
	return TypePMDisputeOracleRespond
}

// Data returns the operation data PMDisputeOracleRespondOperation.
func (op *PMDisputeOracleRespondOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMDisputeOracleRespondOperation to bytes.
func (op *PMDisputeOracleRespondOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMDisputeOracleRespond.Code()))
	enc.Encode(op.Oracle)
	enc.Encode(op.MarketID)
	enc.Encode(op.Response)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// PMUnbanOperation represents pm_unban operation data.
type PMUnbanOperation struct {
	Resolver     string `json:"resolver"`
	Target       string `json:"target"`
	UnbanOracle  bool   `json:"unban_oracle"`
	UnbanCreator bool   `json:"unban_creator"`
}

// Type function that defines the type of operation PMUnbanOperation.
func (op *PMUnbanOperation) Type() OpType {
	return TypePMUnban
}

// Data returns the operation data PMUnbanOperation.
func (op *PMUnbanOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type PMUnbanOperation to bytes.
func (op *PMUnbanOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypePMUnban.Code()))
	enc.Encode(op.Resolver)
	enc.Encode(op.Target)
	enc.Encode(op.UnbanOracle)
	enc.Encode(op.UnbanCreator)
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}

// SetAgentPermissionOperation represents set_agent_permission operation data.
type SetAgentPermissionOperation struct {
	Account    string      `json:"account"`
	AgentName  string      `json:"agent_name"`
	AgentKey   string      `json:"agent_key"`
	Operations []string    `json:"operations"`
	Expiration *types.Time `json:"expiration"`
	Addons     []string    `json:"addons"`
}

// Type function that defines the type of operation SetAgentPermissionOperation.
func (op *SetAgentPermissionOperation) Type() OpType {
	return TypeSetAgentPermission
}

// Data returns the operation data SetAgentPermissionOperation.
func (op *SetAgentPermissionOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type SetAgentPermissionOperation to bytes.
func (op *SetAgentPermissionOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetAgentPermission.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.AgentName)
	enc.EncodePubKey(op.AgentKey)
	enc.EncodeArrString(sortedStrings(op.Operations))
	enc.Encode(op.Expiration)
	enc.EncodeArrString(sortedStrings(op.Addons))
	enc.EncodeUVarint(0) // extensions
	return enc.Err()
}
