package api

import (
	"encoding/json"

	"github.com/biter777/viz-go-lib/operations"
	"github.com/biter777/viz-go-lib/types"
)

// Account structure for the GetAccounts and LookupAccountNames function
type Account struct {
	ID                     types.Int       `json:"id"`
	Name                   string          `json:"name"`
	Master                 types.Authority `json:"master_authority"`
	Active                 types.Authority `json:"active_authority"`
	Regular                types.Authority `json:"regular_authority"`
	MemoKey                string          `json:"memo_key"`
	JSONMetadata           string          `json:"json_metadata"`
	Proxy                  string          `json:"proxy"`
	Referrer               string          `json:"referrer"`
	LastMasterUpdate       types.Time      `json:"last_master_update"`
	LastAccountUpdate      types.Time      `json:"last_account_update"`
	Created                types.Time      `json:"created"`
	RecoveryAccount        string          `json:"recovery_account"`
	LastAccountRecovery    types.Time      `json:"last_account_recovery"`
	SubcontentCount        uint32          `json:"subcontent_count"`
	VoteCount              uint32          `json:"vote_count"`
	ContentCount           uint32          `json:"content_count"`
	AwardedRshares         uint64          `json:"awarded_rshares"`
	CustomSequence         uint64          `json:"custom_sequence"`
	CustomSequenceBlockNum uint64          `json:"custom_sequence_block_num"`
	Energy                 int32           `json:"energy"`
	LastVoteTime           types.Time      `json:"last_vote_time"`
	Balance                types.Asset     `json:"balance"`
	CurationRewards        uint64          `json:"curation_rewards"`
	PostingRewards         uint64          `json:"posting_rewards"`
	ReceiverAwards         uint64          `json:"receiver_awards"`
	BenefactorAwards       uint64          `json:"benefactor_awards"`
	VestingShares          types.Asset     `json:"vesting_shares"`
	DelegatedVestingShares types.Asset     `json:"delegated_vesting_shares"`
	ReceivedVestingShares  types.Asset     `json:"received_vesting_shares"`
	VestingWithdrawRate    types.Asset     `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal  types.Time      `json:"next_vesting_withdrawal"`
	Withdrawn              json.Number     `json:"withdrawn"`
	ToWithdraw             json.Number     `json:"to_withdraw"`
	WithdrawRoutes         uint16          `json:"withdraw_routes"`
	ProxiedVsfVotes        []int64         `json:"proxied_vsf_votes"`
	WitnessesVotedFor      uint16          `json:"witnesses_voted_for"`
	WitnessesVoteWeight    json.Number     `json:"witnesses_vote_weight"`
	LastRootPost           types.Time      `json:"last_root_post"`
	LastPost               types.Time      `json:"last_post"`
	AverageBandwidth       json.Number     `json:"average_bandwidth"`
	LifetimeBandwidth      json.Number     `json:"lifetime_bandwidth"`
	LastBandwidthUpdate    types.Time      `json:"last_bandwidth_update"`
	WitnessVotes           []string        `json:"witness_votes"`
	Valid                  bool            `json:"valid"`
	AccountSeller          string          `json:"account_seller"`
	AccountOfferPrice      types.Asset     `json:"account_offer_price"`
	AccountOnSale          bool            `json:"account_on_sale"`
	SubaccountSeller       string          `json:"subaccount_seller"`
	SubaccountOfferPrice   types.Asset     `json:"subaccount_offer_price"`
	SubaccountOnSale       bool            `json:"subaccount_on_sale"`
}

// Block structure for the GetBlock function
type Block struct {
	BlockHeader
	WitnessSignature string                   `json:"witness_signature"`
	Transactions     []operations.Transaction `json:"transactions"`
}

// BlockHeader structure for the GetBlockHeader function
type BlockHeader struct {
	Number                uint32        `json:"-"`
	Timestamp             types.Time    `json:"timestamp"`
	Witness               string        `json:"witness"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Previous              string        `json:"previous"`
	Extensions            []interface{} `json:"extensions"`
}

// Config structure for the GetConfig function.
type Config struct {
	Percent100                       uint16 `json:"CHAIN_100_PERCENT"`
	Percent1                         uint16 `json:"CHAIN_1_PERCENT"`
	AddressPrefix                    string `json:"CHAIN_ADDRESS_PREFIX"`
	BandwidthAverageWindowSeconds    uint32 `json:"CHAIN_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	BandwidthPrecision               uint32 `json:"CHAIN_BANDWIDTH_PRECISION"`
	ConsensusBandwidthReservePercent uint16 `json:"CONSENSUS_BANDWIDTH_RESERVE_PERCENT"`
	ConsensusBandwidthReserveBelow   uint32 `json:"CONSENSUS_BANDWIDTH_RESERVE_BELOW"`
	HardforkVersion                  string `json:"CHAIN_HARDFORK_VERSION"`
	Version                          string `json:"CHAIN_VERSION"`
	BlockInterval                    uint8  `json:"CHAIN_BLOCK_INTERVAL"`
	BlocksPerDay                     uint16 `json:"CHAIN_BLOCKS_PER_DAY"`
	BlocksPerYear                    uint32 `json:"CHAIN_BLOCKS_PER_YEAR"`
	CashoutWindowSeconds             uint32 `json:"CHAIN_CASHOUT_WINDOW_SECONDS"`
	ChainID                          string `json:"CHAIN_ID"`
	HardforkRequiredWitnesses        uint8  `json:"CHAIN_HARDFORK_REQUIRED_WITNESSES"`
	InitiatorName                    string `json:"CHAIN_INITIATOR_NAME"`
	InitiatorPublicKeyStr            string `json:"CHAIN_INITIATOR_PUBLIC_KEY_STR"`
	InitSupply                       string `json:"CHAIN_INIT_SUPPLY"`
	CommitteeAccount                 string `json:"CHAIN_COMMITTEE_ACCOUNT"`
	CommitteePublicKeyStr            string `json:"CHAIN_COMMITTEE_PUBLIC_KEY_STR"`
	IrreversibleThreshold            uint16 `json:"CHAIN_IRREVERSIBLE_THRESHOLD"`
	IrreversibleSupportMinRun        uint8  `json:"CHAIN_IRREVERSIBLE_SUPPORT_MIN_RUN"`
	MaxAccountNameLength             uint8  `json:"CHAIN_MAX_ACCOUNT_NAME_LENGTH"`
	MaxAccountWitnessVotes           uint8  `json:"CHAIN_MAX_ACCOUNT_WITNESS_VOTES"`
	BlockSize                        uint32 `json:"CHAIN_BLOCK_SIZE"`
	MaxCommentDepth                  uint32 `json:"CHAIN_MAX_COMMENT_DEPTH"`
	MaxMemoLength                    uint16 `json:"CHAIN_MAX_MEMO_LENGTH"`
	MaxWitnesses                     uint8  `json:"CHAIN_MAX_WITNESSES"`
	MaxProxyRecursionDepth           uint8  `json:"CHAIN_MAX_PROXY_RECURSION_DEPTH"`
	MaxReserveRatio                  uint16 `json:"CHAIN_MAX_RESERVE_RATIO"`
	MaxSupportWitnesses              uint8  `json:"CHAIN_MAX_SUPPORT_WITNESSES"`
	MaxShareSupply                   string `json:"CHAIN_MAX_SHARE_SUPPLY"`
	MaxSigCheckDepth                 uint8  `json:"CHAIN_MAX_SIG_CHECK_DEPTH"`
	MaxTimeUntilExpiration           uint16 `json:"CHAIN_MAX_TIME_UNTIL_EXPIRATION"`
	MaxTransactionSize               uint32 `json:"CHAIN_MAX_TRANSACTION_SIZE"`
	MaxUndoHistory                   uint16 `json:"CHAIN_MAX_UNDO_HISTORY"`
	MaxVoteChanges                   uint8  `json:"CHAIN_MAX_VOTE_CHANGES"`
	MaxTopWitnesses                  uint8  `json:"CHAIN_MAX_TOP_WITNESSES"`
	MaxWithdrawRoutes                uint8  `json:"CHAIN_MAX_WITHDRAW_ROUTES"`
	MaxWitnessURLLength              uint16 `json:"CHAIN_MAX_WITNESS_URL_LENGTH"`
	MinAccountCreationFee            uint16 `json:"CHAIN_MIN_ACCOUNT_CREATION_FEE"`
	MinAccountNameLength             uint8  `json:"CHAIN_MIN_ACCOUNT_NAME_LENGTH"`
	MinBlockSizeLimit                uint32 `json:"CHAIN_MIN_BLOCK_SIZE_LIMIT"`
	MaxBlockSizeLimit                uint32 `json:"CHAIN_MAX_BLOCK_SIZE_LIMIT"`
	NullAccount                      string `json:"CHAIN_NULL_ACCOUNT"`
	NumInitiators                    uint8  `json:"CHAIN_NUM_INITIATORS"`
	ProxyToSelfAccount               string `json:"CHAIN_PROXY_TO_SELF_ACCOUNT"`
	SecondsPerYear                   uint32 `json:"CHAIN_SECONDS_PER_YEAR"`
	VestingWithdrawIntervals         uint8  `json:"CHAIN_VESTING_WITHDRAW_INTERVALS"`
	VestingWithdrawIntervalSeconds   uint32 `json:"CHAIN_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	EnergyRegenerationSeconds        uint32 `json:"CHAIN_ENERGY_REGENERATION_SECONDS"`
	TokenSymbol                      uint32 `json:"TOKEN_SYMBOL"`
	SharesSymbol                     string `json:"SHARES_SYMBOL"`
	ChainName                        string `json:"CHAIN_NAME"`
	BlockGenerationPostponedTxLimit  uint8  `json:"CHAIN_BLOCK_GENERATION_POSTPONED_TX_LIMIT"`
	PendingTransactionExecutionLimit uint32 `json:"CHAIN_PENDING_TRANSACTION_EXECUTION_LIMIT"`
}

// DatabaseInfo structure for the GetDatabaseInfo function.
type DatabaseInfo struct {
	TotalSize    string              `json:"total_size"`
	FreeSize     uint64              `json:"free_size"`
	ReservedSize uint64              `json:"reserved_size"`
	UsedSize     string              `json:"used_size"`
	IndexList    []DatabaseInfoIndex `json:"index_list"`
}

// DatabaseInfoIndex additional structure for the function GetDatabaseInfo.
type DatabaseInfoIndex struct {
	Name        string `json:"name"`
	RecordCount uint64 `json:"record_count"`
}

// DynamicGlobalProperties structure for the GetDynamicGlobalProperties function.
type DynamicGlobalProperties struct {
	ID                         types.Int   `json:"id"`
	HeadBlockNumber            uint32      `json:"head_block_number"`
	HeadBlockID                string      `json:"head_block_id"`
	GenesisTime                types.Time  `json:"genesis_time"`
	Time                       types.Time  `json:"time"`
	CurrentWitness             string      `json:"current_witness"`
	CommitteeFund              types.Asset `json:"committee_fund"`
	CommitteeRequests          uint32      `json:"committee_requests"`
	CurrentSupply              types.Asset `json:"current_supply"`
	TotalVersingFund           types.Asset `json:"total_vesting_fund"`
	TotalVestingShares         types.Asset `json:"total_vesting_shares"`
	TotalRewardFund            types.Asset `json:"total_reward_fund"`
	TotalRewardShares          string      `json:"total_reward_shares"`
	InflationCalcBlockNum      uint32      `json:"inflation_calc_block_num"`
	InflationWitnessPercent    int16       `json:"inflation_witness_percent"`
	InflationRatio             int16       `json:"inflation_ratio"`
	AverageBlockSize           uint32      `json:"average_block_size"`
	MaximumBlockSize           uint32      `json:"maximum_block_size"`
	CurrentAslot               uint64      `json:"current_aslot"`
	RecentSlotsFilled          string      `json:"recent_slots_filled"`
	ParticipationCount         uint8       `json:"participation_count"`
	LastIrreversibleBlockNum   uint32      `json:"last_irreversible_block_num"`
	MaxVirtualBandwidth        string      `json:"max_virtual_bandwidth"`
	CurrentReserveRatio        uint64      `json:"current_reserve_ratio"`
	VoteRegenerationPerDay     uint32      `json:"vote_regeneration_per_day"`
	BandwidthReserveCandidates uint32      `json:"bandwidth_reserve_candidates"`
}

// VestingDelegationExpiration structure for the GetExpiringVestingDelegations function.
type VestingDelegationExpiration struct {
	ID            types.Int   `json:"id"`
	Delegator     string      `json:"delegator"`
	VestingShares types.Asset `json:"vesting_shares"`
	Expiration    types.Time  `json:"expiration"`
}

// NextScheduledHardfork structure for the GetNextScheduledHardfork function.
type NextScheduledHardfork struct {
	HfVersion string     `json:"hf_version"`
	LiveTime  types.Time `json:"live_time"`
}

// MasterHistory structure for the GetMasterHistory function.
type MasterHistory struct {
	ID                      types.Int       `json:"id"`
	Account                 string          `json:"account"`
	PreviousMasterAuthority types.Authority `json:"previous_master_authority"`
	LastValidTime           string          `json:"last_valid_time"`
}

// ProposalObject structure for the GetProposedTransaction function.
type ProposalObject struct {
	Author                    string                `json:"author"`
	Title                     string                `json:"title"`
	Memo                      string                `json:"memo"`
	ExpirationTime            types.Time            `json:"expiration_time"`
	ReviewPeriodTime          types.Time            `json:"review_period_time,omitempty"`
	ProposedOperations        operations.Operations `json:"proposed_operations"`
	RequiredActiveApprovals   []string              `json:"required_active_approvals"`
	AvailableActiveApprovals  []string              `json:"available_active_approvals"`
	RequiredOwnerApprovals    []string              `json:"required_owner_approvals"`
	AvailableOwnerApprovals   []string              `json:"available_owner_approvals"`
	RequiredPostingApprovals  []string              `json:"required_posting_approvals"`
	AvailablePostingApprovals []string              `json:"available_posting_approvals"`
	AvailableKeyApprovals     []string              `json:"available_key_approvals"`
}

// VestingDelegation structure for the GetVestingDelegations function.
type VestingDelegation struct {
	ID                types.Int   `json:"id"`
	Delegator         string      `json:"delegator"`
	Delegatee         string      `json:"delegatee"`
	VestingShares     types.Asset `json:"vesting_shares"`
	MinDelegationTime types.Time  `json:"min_delegation_time"`
}

// WithdrawVestingRoutes structure for the GetWithdrawRoutes function.
type WithdrawVestingRoutes struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Percent     uint16 `json:"percent"`
	AutoVest    bool   `json:"auto_vest"`
}

// AccountOnSale structure for the GetAccountsOnSale function.
type AccountOnSale struct {
	Account           string      `json:"account"`
	AccountSeller     string      `json:"account_seller"`
	AccountOfferPrice types.Asset `json:"account_offer_price"`
}

// SubAccountOnSale structure for the GetSubAccountsOnSale function.
type SubAccountOnSale struct {
	Account           string      `json:"account"`
	AccountSeller     string      `json:"account_seller"`
	AccountOfferPrice types.Asset `json:"account_offer_price"`
}

// PaidSubscription structure for the GetPaidSubscriptions function.
type PaidSubscription struct {
	ID         types.Int  `json:"id"`
	Creator    string     `json:"creator"`
	URL        string     `json:"url"`
	Levels     uint16     `json:"levels"`
	Amount     int64      `json:"amount"`
	Period     uint16     `json:"period"`
	UpdateTime types.Time `json:"update_time"`
}

// PaidSubscriptionState structure for the GetPaidSubscriptionOptions function.
type PaidSubscriptionState struct {
	ID                                            types.Int  `json:"id"`
	Creator                                       string     `json:"creator"`
	URL                                           string     `json:"url"`
	Levels                                        uint16     `json:"levels"`
	Amount                                        int64      `json:"amount"`
	Period                                        uint16     `json:"period"`
	UpdateTime                                    types.Time `json:"update_time"`
	ActiveSubscribers                             []string   `json:"active_subscribers"`
	ActiveSubscribersCount                        uint32     `json:"active_subscribers_count"`
	ActiveSubscribersSummaryAmount                int64      `json:"active_subscribers_summary_amount"`
	ActiveSubscribersWithAutoRenewal              []string   `json:"active_subscribers_with_auto_renewal"`
	ActiveSubscribersWithAutoRenewalCount         uint32     `json:"active_subscribers_with_auto_renewal_count"`
	ActiveSubscribersWithAutoRenewalSummaryAmount int64      `json:"active_subscribers_with_auto_renewal_summary_amount"`
}

// PaidSubscribeState structure for the GetPaidSubscriptions function.
type PaidSubscribeState struct {
	ID          types.Int  `json:"id"`
	Subscriber  string     `json:"subscriber"`
	Creator     string     `json:"creator"`
	Level       uint16     `json:"level"`
	Amount      int64      `json:"amount"`
	Period      uint16     `json:"period"`
	StartTime   types.Time `json:"start_time"`
	NextTime    types.Time `json:"next_time"`
	EndTime     types.Time `json:"end_time"`
	Active      bool       `json:"active"`
	AutoRenewal bool       `json:"auto_renewal"`
}

// InviteObject structure for the GetInviteById and GetInviteByKey function.
type InviteObject struct {
	ID             types.Int   `json:"id"`
	Creator        string      `json:"creator"`
	Receiver       string      `json:"receiver"`
	InviteKey      string      `json:"invite_key"`
	InviteSecret   string      `json:"invite_secret"`
	Balance        types.Asset `json:"balance"`
	ClaimedBalance types.Asset `json:"claimed_balance"`
	CreateTime     types.Time  `json:"create_time"`
	ClaimTime      types.Time  `json:"claim_time"`
	Status         uint16      `json:"status"`
}

// CommitteeObject structure for the GetCommitteeRequest function.
type CommitteeObject struct {
	ID                     types.Int            `json:"id"`
	RequestID              uint32               `json:"request_id"`
	URL                    string               `json:"url"`
	Creator                string               `json:"creator"`
	Worker                 string               `json:"worker"`
	RequiredAmountMin      types.Asset          `json:"required_amount_min"`
	RequiredAmountMax      types.Asset          `json:"required_amount_max"`
	StartTime              types.Time           `json:"start_time"`
	Duration               uint32               `json:"duration"`
	EndTime                types.Time           `json:"end_time"`
	Status                 uint16               `json:"status"`
	VotesCount             uint32               `json:"votes_count"`
	ConclusionTime         types.Time           `json:"conclusion_time"`
	ConclusionPayoutAmount types.Asset          `json:"conclusion_payout_amount"`
	PayoutAmount           types.Asset          `json:"payout_amount"`
	RemainPayoutAmount     types.Asset          `json:"remain_payout_amount"`
	LastPayoutTime         types.Time           `json:"last_payout_time"`
	PayoutTime             types.Time           `json:"payout_time"`
	Votes                  []CommitteeVoteState `json:"votes"`
}

// CommitteeVoteState structure for the GetCommitteeRequest function.
type CommitteeVoteState struct {
	Voter       string     `json:"voter"`
	VotePercent int16      `json:"vote_percent"`
	LastUpdate  types.Time `json:"last_update"`
}

// Escrow structure for the GetEscrow function.
type Escrow struct {
	ID                   types.Int   `json:"id"`
	EscrowID             uint32      `json:"escrow_id"`
	From                 string      `json:"from"`
	To                   string      `json:"to"`
	Agent                string      `json:"agent"`
	RatificationDeadline types.Time  `json:"ratification_deadline"`
	EscrowExpiration     types.Time  `json:"escrow_expiration"`
	TokenBalance         types.Asset `json:"token_balance"`
	Fee                  types.Asset `json:"pending_fee"`
	ToApproved           bool        `json:"to_approved"`
	AgentApproved        bool        `json:"agent_approved"`
	Disputed             bool        `json:"disputed"`
	IsApproved           bool        `json:"is_approved"`
}

// AccountRecoveryRequest structure for the GetRecoveryRequest function.
type AccountRecoveryRequest struct {
	ID                 types.Int       `json:"id"`
	AccountToRecover   string          `json:"account_to_recover"`
	NewMasterAuthority types.Authority `json:"new_master_authority"`
	Expires            types.Time      `json:"expires"`
}

// BroadcastResponse structure for the BroadcastTransactionSynchronous function
type BroadcastResponse struct {
	ID       string `json:"id"`
	BlockNum int32  `json:"block_num"`
	TrxNum   int32  `json:"trx_num"`
	Expired  bool   `json:"expired"`
}

// WitnessSchedule structure for the GetWitnessSchedule function.
type WitnessSchedule struct {
	ID                       types.Int             `json:"id"`
	CurrentVirtualTime       string                `json:"current_virtual_time"`
	NextShuffleBlockNum      uint32                `json:"next_shuffle_block_num"`
	CurrentShuffledWitnesses string                `json:"current_shuffled_witnesses"`
	NumScheduledWitnesses    uint8                 `json:"num_scheduled_witnesses"`
	MedianProps              types.ChainProperties `json:"median_props"`
	MajorityVersion          string                `json:"majority_version"`
}

// Witness structure for the GetWitnessByAccount, GetWitnesses and GetWitnessByVote function.
type Witness struct {
	ID                    types.Int             `json:"id"`
	Owner                 string                `json:"owner"`
	Created               types.Time            `json:"created"`
	URL                   string                `json:"url"`
	TotalMissed           uint32                `json:"total_missed"`
	LastAslot             uint64                `json:"last_aslot"`
	LastConfirmedBlockNum uint64                `json:"last_confirmed_block_num"`
	SigningKey            string                `json:"signing_key"`
	Props                 types.ChainProperties `json:"props"`
	Votes                 string                `json:"votes"`
	PenaltyPercent        uint32                `json:"penalty_percent"`
	CountedVotes          string                `json:"counted_votes"`
	VirtualLastUpdate     string                `json:"virtual_last_update"`
	VirtualPosition       string                `json:"virtual_position"`
	VirtualScheduledTime  string                `json:"virtual_scheduled_time"`
	LastWork              string                `json:"last_work"`
	RunningVersion        string                `json:"running_version"`
	HardforkVersionVote   string                `json:"hardfork_version_vote"`
	HardforkTimeVote      types.Time            `json:"hardfork_time_vote"`
}

type CallbackBlockResponse struct {
	BlockNum   int                                  `json:"block_num"`
	Operations []operations.CallbackBlockOperations `json:"operations"`
}
