package types

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//ChainProperties is an additional structure used by other structures.
type ChainProperties struct {
	AccountCreationFee                    Asset  `json:"account_creation_fee,omitempty"`
	MaximumBlockSize                      uint32 `json:"maximum_block_size,omitempty"`
	CreateAccountDelegationRatio          uint32 `json:"create_account_delegation_ratio,omitempty"`
	CreateAccountDelegationTime           uint32 `json:"create_account_delegation_time,omitempty"`
	MinDelegation                         Asset  `json:"min_delegation,omitempty"`
	MinCurationPercent                    int16  `json:"min_curation_percent,omitempty"`
	MaxCurationPercent                    int16  `json:"max_curation_percent,omitempty"`
	BandwidthReservePercent               int16  `json:"bandwidth_reserve_percent,omitempty"`
	BandwidthReserveBelow                 Asset  `json:"bandwidth_reserve_below,omitempty"`
	FlagEnergyAdditionalCost              int16  `json:"flag_energy_additional_cost,omitempty"`
	VoteAccountingMinRshares              uint32 `json:"vote_accounting_min_rshares,omitempty"`
	CommitteeRequestApproveMinPercent     int16  `json:"committee_request_approve_min_percent,omitempty"`
	InflationWitnessPercent               int16  `json:"inflation_witness_percent,omitempty"`
	InflationRatioCommitteeVsRewardFund   int16  `json:"inflation_ratio_committee_vs_reward_fund,omitempty"`
	InflationRecalcPeriod                 uint32 `json:"inflation_recalc_period,omitempty"`
	DataOperationsCostAdditionalBandwidth uint32 `json:"data_operations_cost_additional_bandwidth,omitempty"`
	WitnessMissPenaltyPercent             int16  `json:"witness_miss_penalty_percent,omitempty"`
	WitnessMissPenaltyDuration            uint32 `json:"witness_miss_penalty_duration,omitempty"`
}

//MarshalTransaction is a function of converting type ChainProperties to bytes.
func (cp *ChainProperties) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.AccountCreationFee)
	enc.Encode(cp.MaximumBlockSize)
	enc.Encode(cp.CreateAccountDelegationRatio)
	enc.Encode(cp.CreateAccountDelegationTime)
	enc.Encode(cp.MinDelegation)
	enc.Encode(cp.MinCurationPercent)
	enc.Encode(cp.MaxCurationPercent)
	enc.Encode(cp.BandwidthReservePercent)
	enc.Encode(cp.BandwidthReserveBelow)
	enc.Encode(cp.FlagEnergyAdditionalCost)
	enc.Encode(cp.VoteAccountingMinRshares)
	enc.Encode(cp.CommitteeRequestApproveMinPercent)
	enc.Encode(cp.InflationWitnessPercent)
	enc.Encode(cp.InflationRatioCommitteeVsRewardFund)
	enc.Encode(cp.InflationRecalcPeriod)
	enc.Encode(cp.DataOperationsCostAdditionalBandwidth)
	enc.Encode(cp.WitnessMissPenaltyPercent)
	enc.Encode(cp.WitnessMissPenaltyDuration)
	return enc.Err()
}

// Очень старое возможно даже можно будет удалить.

//ChainProperties is an additional structure used by other structures.
type ChainPropertiesOLD struct {
	AccountCreationFee *Asset `json:"account_creation_fee"`
	MaximumBlockSize   uint32 `json:"maximum_block_size"`
	SBDInterestRate    uint16 `json:"sbd_interest_rate"`
}

//MarshalTransaction is a function of converting type ChainPropertiesOLD to bytes.
func (cp *ChainPropertiesOLD) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.AccountCreationFee)
	enc.Encode(cp.MaximumBlockSize)
	enc.Encode(cp.SBDInterestRate)
	return enc.Err()
}
