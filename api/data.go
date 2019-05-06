package api

import (
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//CommitteeObject structure for the GetCommitteeRequest function
type CommitteeObject struct {
	ID                     *types.Int            `json:"id"`
	RequestID              uint32                `json:"request_id"`
	Url                    string                `json:"url"`
	Creator                string                `json:"creator"`
	Worker                 string                `json:"worker"`
	RequiredAmount_min     *types.Asset          `json:"required_amount_min"`
	RequiredAmount_max     *types.Asset          `json:"required_amount_max"`
	StartTime              *types.Time           `json:"start_time"`
	Duration               uint32                `json:"duration"`
	EndTime                *types.Time           `json:"end_time"`
	Status                 uint16                `json:"status"`
	VotesCount             uint32                `json:"votes_count"`
	ConclusionTime         *types.Time           `json:"conclusion_time"`
	ConclusionPayoutAmount *types.Asset          `json:"conclusion_payout_amount"`
	PayoutAmount           *types.Asset          `json:"payout_amount"`
	RemainPayoutAmount     *types.Asset          `json:"remain_payout_amount"`
	LastPayoutTime         *types.Time           `json:"last_payout_time"`
	PayoutTime             *types.Time           `json:"payout_time"`
	Votes                  []*CommitteeVoteState `json:"votes"`
}

//CommitteeVoteState structure for the GetCommitteeRequestVotes function
type CommitteeVoteState struct {
	Voter       string      `json:"voter"`
	VotePercent int16       `json:"vote_percent"`
	LastUpdate  *types.Time `json:"last_update"`
}

//Account structure for the GetAccounts and LookupAccountNames functions
type Account struct {
	ID                     int64                  `json:"id"`
	Name                   string                 `json:"name"`
	Master                 *types.Authority       `json:"master_authority"`
	Active                 *types.Authority       `json:"active_authority"`
	Regular                *types.Authority       `json:"regular_authority"`
	Memo                   string                 `json:"memo_key"`
	JSONMetadata           *types.AccountMetadata `json:"json_metadata"`
	Proxy                  string                 `json:"proxy"`
	Referrer               string                 `json:"referrer"`
	LastMasterUpdate       *types.Time            `json:"last_master_update"`
	LastAccountUpdate      *types.Time            `json:"last_account_update"`
	Created                *types.Time            `json:"created"`
	RecoveryAccount        string                 `json:"recovery_account"`
	LastAccountRecovery    *types.Time            `json:"last_account_recovery"`
	SubcontentCount        uint32                 `json:"subcontent_count"`
	VoteCount              uint32                 `json:"vote_count"`
	ContentCount           uint32                 `json:"content_count"`
	AwardedRshares         uint64                 `json:"awarded_rshares"`
	CustomSequence         uint64                 `json:"custom_sequence"`
	CustomSequenceBlockNum uint64                 `json:"custom_sequence_block_num"`
	Energy                 int16                  `json:"energy"`
	LastVoteTime           *types.Time            `json:"last_vote_time"`
	Balance                *types.Asset           `json:"balance"`
	VestingShares          *types.Asset           `json:"vesting_shares"`
	DelegatedVestingShares *types.Asset           `json:"delegated_vesting_shares"`
	ReceivedVestingShares  *types.Asset           `json:"received_vesting_shares"`
	VestingWithdrawRate    *types.Asset           `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal  *types.Time            `json:"next_vesting_withdrawal"`
	Withdrawn              int64                  `json:"withdrawn"`
	ToWithdraw             int64                  `json:"to_withdraw"`
	WithdrawRoutes         uint16                 `json:"withdraw_routes"`
	CurationRewards        int64                  `json:"curation_rewards"`
	PostingRewards         int64                  `json:"posting_rewards"`
	ReceiverAwards         int64                  `json:"receiver_awards"`
	BenefactorAwards       int64                  `json:"benefactor_awards"`
	ProxiedVsfVotes        []int64                `json:"proxied_vsf_votes"`
	WitnessesVotedFor      uint16                 `json:"witnesses_voted_for"`
	WitnessesVoteWeight    string                 `json:"witnesses_vote_weight"`
	LastPost               *types.Time            `json:"last_post"`
	LastRootPost           *types.Time            `json:"last_root_post"`
	AverageBandwidth       int64                  `json:"average_bandwidth"`
	LifetimeBandwidth      string                 `json:"lifetime_bandwidth"`
	LastBandwidthUpdate    *types.Time            `json:"last_bandwidth_update"`
	WitnessVotes           []string               `json:"witness_votes"`
}

//Block structure for the GetBlock function
type Block struct {
	Number                uint32               `json:"-"`
	Timestamp             *types.Time          `json:"timestamp"`
	Witness               string               `json:"witness"`
	WitnessSignature      string               `json:"witness_signature"`
	TransactionMerkleRoot string               `json:"transaction_merkle_root"`
	Previous              string               `json:"previous"`
	Extensions            [][]interface{}      `json:"extensions"`
	Transactions          []*types.Transaction `json:"transactions"`
}

//BlockHeader structure for the GetBlockHeader function
type BlockHeader struct {
	Number                uint32        `json:"-"`
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	Witness               string        `json:"witness"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Extensions            []interface{} `json:"extensions"`
}

//Config structure for the GetConfig function.
type Config struct {
	Percent100                       int           `json:"CHAIN_100_PERCENT"`
	Percent1                         *types.Int    `json:"CHAIN_1_PERCENT"`
	AddressPrefix                    string        `json:"CHAIN_ADDRESS_PREFIX"`
	BandwidthAverageWindowSeconds    *types.Int    `json:"CHAIN_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	BandwidthPrecision               *types.Int    `json:"CHAIN_BANDWIDTH_PRECISION"`
	ConsensusBandwidthReservePercent *types.Int    `json:"CONSENSUS_BANDWIDTH_RESERVE_PERCENT"`
	ConsensusBandwidthReserveBelow   *types.Int    `json:"CONSENSUS_BANDWIDTH_RESERVE_BELOW"`
	HardforkVersion                  string        `json:"CHAIN_HARDFORK_VERSION"`
	Version                          string        `json:"CHAIN_VERSION"`
	BlockInterval                    uint          `json:"CHAIN_BLOCK_INTERVAL"`
	BlocksPerDay                     *types.Int    `json:"CHAIN_BLOCKS_PER_DAY"`
	BlocksPerYear                    *types.Int    `json:"CHAIN_BLOCKS_PER_YEAR"`
	CashoutWindowSeconds             *types.Int    `json:"CHAIN_CASHOUT_WINDOW_SECONDS"`
	ChainID                          string        `json:"CHAIN_ID"`
	HardforkRequiredWitnesses        *types.Int    `json:"CHAIN_HARDFORK_REQUIRED_WITNESSES"`
	InitiatorName                    string        `json:"CHAIN_INITIATOR_NAME"`
	InitiatorPublicKey               string        `json:"CHAIN_INITIATOR_PUBLIC_KEY_STR"`
	InitSupply                       *types.UInt32 `json:"CHAIN_INIT_SUPPLY"`
	CommitteeAccount                 string        `json:"CHAIN_COMMITTEE_ACCOUNT"`
	CommitteePublicKey               string        `json:"CHAIN_COMMITTEE_PUBLIC_KEY_STR"`
	IrreversibleThreshold            *types.Int    `json:"CHAIN_IRREVERSIBLE_THRESHOLD"`
	IrreversibleSupportMinRun        *types.Int    `json:"CHAIN_IRREVERSIBLE_SUPPORT_MIN_RUN"`
	MaxAccountNameLength             *types.Int    `json:"CHAIN_MAX_ACCOUNT_NAME_LENGTH"`
	MaxAccountWitnessVotes           *types.Int    `json:"CHAIN_MAX_ACCOUNT_WITNESS_VOTES"`
	BlockSize                        *types.Int    `json:"CHAIN_BLOCK_SIZE"`
	MaxCommentDepth                  *types.Int    `json:"CHAIN_MAX_COMMENT_DEPTH"`
	MaxMemoSize                      *types.Int    `json:"CHAIN_MAX_MEMO_SIZE"`
	MaxWitnesses                     *types.Int    `json:"CHAIN_MAX_WITNESSES"`
	MaxProxyRecursionDepth           *types.Int    `json:"CHAIN_MAX_PROXY_RECURSION_DEPTH"`
	MaxReserveRatio                  *types.Int    `json:"CHAIN_MAX_RESERVE_RATIO"`
	MaxSupportWitnesses              *types.Int    `json:"CHAIN_MAX_SUPPORT_WITNESSES"`
	MaxShareSupply                   string        `json:"CHAIN_MAX_SHARE_SUPPLY"`
	MaxSigCheckDepth                 *types.Int    `json:"CHAIN_MAX_SIG_CHECK_DEPTH"`
	MaxTimeUntilExpiration           *types.Int    `json:"CHAIN_MAX_TIME_UNTIL_EXPIRATION"`
	MaxTransactionSize               *types.Int    `json:"CHAIN_MAX_TRANSACTION_SIZE"`
	MaxUndoHistory                   *types.Int    `json:"CHAIN_MAX_UNDO_HISTORY"`
	MaxVoteChanges                   *types.Int    `json:"CHAIN_MAX_VOTE_CHANGES"`
	MaxTopWitnesses                  *types.Int    `json:"CHAIN_MAX_TOP_WITNESSES"`
	MaxWithdrawRoutes                *types.Int    `json:"CHAIN_MAX_WITHDRAW_ROUTES"`
	MaxWitnessURLLength              *types.Int    `json:"CHAIN_MAX_WITNESS_URL_LENGTH"`
	MinAccountCreationFee            *types.Int    `json:"CHAIN_MIN_ACCOUNT_CREATION_FEE"`
	MinAccountNameLength             *types.Int    `json:"CHAIN_MIN_ACCOUNT_NAME_LENGTH"`
	MinBlockSizeLimit                *types.Int    `json:"CHAIN_MIN_BLOCK_SIZE_LIMIT"`
	MaxBlockSizeLimit                *types.Int    `json:"CHAIN_MAX_BLOCK_SIZE_LIMIT"`
	NullAccount                      string        `json:"CHAIN_NULL_ACCOUNT"`
	NumInitiators                    *types.Int    `json:"CHAIN_NUM_INITIATORS"`
	ProxyToSelfAccount               string        `json:"CHAIN_PROXY_TO_SELF_ACCOUNT"`
	SecondsPerYear                   *types.Int    `json:"CHAIN_SECONDS_PER_YEAR"`
	VestingWithdrawIntervals         *types.Int    `json:"CHAIN_VESTING_WITHDRAW_INTERVALS"`
	VestingWithdrawIntervalSeconds   *types.Int    `json:"CHAIN_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	EnergyRegenerationSeconds        int           `json:"CHAIN_ENERGY_REGENERATION_SECONDS"`
	TokenSymbol                      *types.Int    `json:"TOKEN_SYMBOL"`
	SharesSymbol                     *types.Int    `json:"SHARES_SYMBOL"`
	ChainName                        string        `json:"CHAIN_NAME"`
	BlockGenerationPostponedTXLimit  *types.Int    `json:"CHAIN_BLOCK_GENERATION_POSTPONED_TX_LIMIT"`
	PendingTransactionExecutionLimit *types.Int    `json:"CHAIN_PENDING_TRANSACTION_EXECUTION_LIMIT"`
}

//DatabaseInfo structure for the GetDatabaseInfo function.
type DatabaseInfo struct {
	TotalSize    uint64              `json:"total_size"`
	FreeSize     uint64              `json:"free_size"`
	ReservedSize uint64              `json:"reserved_size"`
	UsedSize     uint64              `json:"used_size"`
	IndexList    []DatabaseInfoIndex `json:"index_list"`
}

//DatabaseInfoIndex additional structure for the function GetDatabaseInfo.
type DatabaseInfoIndex struct {
	Name        string `json:"name"`
	RecordCount uint64 `json:"record_count"`
}

//DynamicGlobalProperties structure for the GetDynamicGlobalProperties function.
type DynamicGlobalProperties struct {
	ID                         *types.Int   `json:"id"`
	HeadBlockNumber            uint32       `json:"head_block_number"`
	HeadBlockID                string       `json:"head_block_id"`
	GenesisTime                *types.Time  `json:"genesis_time"`
	Time                       *types.Time  `json:"time"`
	CurrentWitness             string       `json:"current_witness"`
	CommitteeFund              *types.Asset `json:"committee_fund"`
	CommitteeRequests          uint32       `json:"committee_requests"`
	CurrentSupply              *types.Asset `json:"current_supply"`
	TotalVersingFund           *types.Asset `json:"total_vesting_fund"`
	TotalVestingShares         *types.Asset `json:"total_vesting_shares"`
	TotalRewardFund            *types.Asset `json:"total_reward_fund"`
	TotalRewardShares          *types.Int64 `json:"total_reward_shares"`
	InflationCalcBlockNum      uint32       `json:"inflation_calc_block_num"`
	InflationWitnessPercent    int16        `json:"inflation_witness_percent"`
	InflationRatio             int16        `json:"inflation_ratio"`
	AverageBlockSize           uint32       `json:"average_block_size"`
	MaximumBlockSize           uint32       `json:"maximum_block_size"`
	CurrentAslot               uint64       `json:"current_aslot"`
	RecentSlotsFilled          string       `json:"recent_slots_filled"`
	ParticipationCount         uint8        `json:"participation_count"`
	LastIrreversibleBlockNum   uint32       `json:"last_irreversible_block_num"`
	MaxVirtualBandwidth        string       `json:"max_virtual_bandwidth"`
	CurrentReserveRatio        uint64       `json:"current_reserve_ratio"`
	VoteRegenerationPerDay     uint32       `json:"vote_regeneration_per_day"`
	BandwidthReserveCandidates uint32       `json:"bandwidth_reserve_candidates"`
}

//VestingDelegationExpiration structure for the GetExpiringVestingDelegations function.
type VestingDelegationExpiration struct {
	ID            *types.Int   `json:"id"`
	Delegator     string       `json:"delegator"`
	VestingShares *types.Asset `json:"vesting_shares"`
	Expiration    *types.Time  `json:"expiration"`
}

//NextScheduledHardfork structure for the GetNextScheduledHardfork function.
type NextScheduledHardfork struct {
	HfVersion string      `json:"hf_version"`
	LiveTime  *types.Time `json:"live_time"`
}

//MasterHistory structure for the GetMasterHistory function.
type MasterHistory struct {
	ID                      *types.Int       `json:"id"`
	Account                 string           `json:"account"`
	PreviousMasterAuthority *types.Authority `json:"previous_master_authority"`
	LastValidTime           *types.Time      `json:"last_valid_time"`
}

//ProposalObject structure for the GetProposedTransaction function.
type ProposalObject struct {
	Author                    string            `json:"author"`
	Title                     string            `json:"title"`
	Memo                      string            `json:"memo"`
	ExpirationTime            *types.Time       `json:"expiration_time"`
	ReviewPeriodTime          *types.Time       `json:"review_period_time,omitempty"`
	ProposedOperations        *types.Operations `json:"proposed_operations"`
	RequiredActiveApprovals   []string          `json:"required_active_approvals"`
	AvailableActiveApprovals  []string          `json:"available_active_approvals"`
	RequiredOwnerApprovals    []string          `json:"required_owner_approvals"`
	AvailableOwnerApprovals   []string          `json:"available_owner_approvals"`
	RequiredRegularApprovals  []string          `json:"required_regular_approvals"`
	AvailableRegularApprovals []string          `json:"available_regular_approvals"`
	AvailableKeyApprovals     []string          `json:"available_key_approvals"`
}

//VestingDelegation structure for the GetVestingDelegations function.
type VestingDelegation struct {
	ID                *types.Int   `json:"id"`
	Delegator         string       `json:"delegator"`
	Delegatee         string       `json:"delegatee"`
	VestingShares     *types.Asset `json:"vesting_shares"`
	MinDelegationTime *types.Time  `json:"min_delegation_time"`
}

//WithdrawVestingRoutes structure for the GetWithdrawRoutes function.
type WithdrawVestingRoutes struct {
	ID          *types.Int `json:"id"`
	FromAccount string     `json:"from_account"`
	ToAccount   string     `json:"to_account"`
	Percent     uint16     `json:"percent"`
	AutoVest    bool       `json:"auto_vest"`
}

//Escrow structure for the GetEscrow function.
type Escrow struct {
	ID                   *types.Int   `json:"id"`
	EscrowID             uint32       `json:"escrow_id"`
	From                 string       `json:"from"`
	To                   string       `json:"to"`
	Agent                string       `json:"agent"`
	RatificationDeadline *types.Time  `json:"ratification_deadline"`
	EscrowExpiration     *types.Time  `json:"escrow_expiration"`
	TokenBalance         *types.Asset `json:"token_balance"`
	Fee                  *types.Asset `json:"pending_fee"`
	ToApproved           bool         `json:"to_approved"`
	AgentApproved        bool         `json:"agent_approved"`
	Disputed             bool         `json:"disputed"`
}

//AccountRecoveryRequest structure for the GetRecoveryRequest function.
type AccountRecoveryRequest struct {
	ID                 *types.Int       `json:"id"`
	AccountToRecover   string           `json:"account_to_recover"`
	NewMasterAuthority *types.Authority `json:"new_master_authority"`
	Expires            *types.Time      `json:"expires"`
	Extensions         []interface{}    `json:"extensions"`
}

//Blogs structure for the GetBlog function
type Blogs struct {
	Content  *ContentData `json:"content"`
	Blog     string       `json:"blog"`
	ReblogOn *types.Time  `json:"reblog_on"`
	EntryID  uint32       `json:"entry_id"`
}

//ContentData additional structure for the function GetBlog.
type ContentData struct {
	ID                       *types.Int           `json:"id"`
	Title                    string               `json:"title"`
	Body                     string               `json:"body"`
	JSONMetadata             string               `json:"json_metadata"`
	ParentAuthor             string               `json:"parent_author"`
	ParentPermlink           string               `json:"parent_permlink"`
	Author                   string               `json:"author"`
	Permlink                 string               `json:"permlink"`
	LastUpdate               *types.Time          `json:"last_update"`
	Created                  *types.Time          `json:"created"`
	Active                   *types.Time          `json:"active"`
	LastPayout               *types.Time          `json:"last_payout"`
	Depth                    uint8                `json:"depth"`
	Children                 uint32               `json:"children"`
	ChildrenRshares          *types.UInt64        `json:"children_rshares"`
	NetRshares               int64                `json:"net_rshares"`
	AbsRshares               int64                `json:"abs_rshares"`
	VoteRshares              int64                `json:"vote_rshares"`
	CashoutTime              *types.Time          `json:"cashout_time"`
	TotalVoteWeight          uint64               `json:"total_vote_weight"`
	CurationPercent          int16                `json:"curation_percent"`
	ConsensusCurationPercent int16                `json:"consensus_curation_percent"`
	PayoutValue              *types.Asset         `json:"payout_value"`
	SharesPayoutValue        *types.Asset         `json:"shares_payout_value"`
	CuratorPayoutValue       *types.Asset         `json:"curator_payout_value"`
	BeneficiaryPayoutValue   *types.Asset         `json:"beneficiary_payout_value"`
	AuthorRewards            int64                `json:"author_rewards"`
	NetVotes                 int32                `json:"net_votes"`
	RootContent              *types.Int           `json:"root_content"`
	Beneficiaries            []*types.Beneficiary `json:"beneficiaries"`
}

//BlogAuthor structure for the GetBlogAuthors function
type BlogAuthor struct {
	Author string `json:"author"`
	Count  uint32 `json:"count"`
}

//BlogEntries structure for the GetBlogEntries function
type BlogEntries struct {
	Author   string      `json:"author"`
	Permlink string      `json:"permlink"`
	Blog     string      `json:"blog"`
	ReblogOn *types.Time `json:"reblog_on"`
	EntryID  uint32      `json:"entry_id"`
}

//Feeds structure for the GetFeed function
type Feeds struct {
	Content  *ContentData `json:"content"`
	ReblogBy []string     `json:"reblog_by"`
	ReblogOn *types.Time  `json:"reblog_on"`
	EntryID  uint32       `json:"entry_id"`
}

//FeedEntry structure for the GetFeedEntries function
type FeedEntry struct {
	Author   string      `json:"author"`
	Permlink string      `json:"permlink"`
	ReblogBy []string    `json:"reblog_by"`
	ReblogOn *types.Time `json:"reblog_on"`
	EntryID  uint32      `json:"entry_id"`
}

//FollowCount structure for the GetFollowCount function
type FollowCount struct {
	Account        string `json:"account"`
	FollowerCount  int    `json:"follower_count"`
	FollowingCount int    `json:"following_count"`
	Limit          int    `json:"limit"`
}

//FollowObject structure for the GetFollowers and GetFollowing functions
type FollowObject struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

//InviteObject structure for the GetInviteById and GetInviteByKey functions
type InviteObject struct {
	ID             *types.Int   `json:"id"`
	Creator        string       `json:"creator"`
	Receiver       string       `json:"receiver"`
	InviteKey      string       `json:"invite_key"`
	InviteSecret   string       `json:"invite_secret"`
	Balance        *types.Asset `json:"balance"`
	ClaimedBalance *types.Asset `json:"claimed_balance"`
	CreateTime     *types.Time  `json:"create_time"`
	ClaimTime      *types.Time  `json:"claim_time"`
	Status         uint16       `json:"status"`
}

//BroadcastResponse structure for the BroadcastTransactionSynchronous function
type BroadcastResponse struct {
	ID       string `json:"id"`
	BlockNum int32  `json:"block_num"`
	TrxNum   int32  `json:"trx_num"`
	Expired  bool   `json:"expired"`
}

//Votes structure for the GetAccountVotes function.
type Votes struct {
	Authorperm string      `json:"authorperm"`
	Weight     uint64      `json:"weight"`
	Rshares    int64       `json:"rshares"`
	Percent    int16       `json:"percent"`
	Time       *types.Time `json:"time"`
}

//VoteState structure for the GetActiveVotes, GetDiscussionsByActive, GetDiscussionsByAuthorBeforeDate, GetDiscussionsByBlog, GetDiscussionsByCashout, GetDiscussionsByChildren, GetDiscussionsByComments, GetDiscussionsByCreated, GetDiscussionsByFeed, GetDiscussionsByHot, GetDiscussionsByPayout, GetDiscussionsByPromoted, GetDiscussionsByTrending and GetDiscussionsByVotes functions.
type VoteState struct {
	Voter   string      `json:"voter"`
	Weight  uint64      `json:"weight"`
	Rshares int64       `json:"rshares"`
	Percent int16       `json:"percent"`
	Time    *types.Time `json:"time"`
}

//Content structure for the GetContent, GetContentReplies, GetRepliesByLastUpdate, GetDiscussionsByActive, GetDiscussionsByAuthorBeforeDate, GetDiscussionsByBlog, GetDiscussionsByCashout, GetDiscussionsByChildren, GetDiscussionsByComments, GetDiscussionsByCreated, GetDiscussionsByFeed, GetDiscussionsByHot, GetDiscussionsByPayout, GetDiscussionsByPromoted, GetDiscussionsByTrending and GetDiscussionsByVotes function.
type Content struct {
	ID                      *types.ID     `json:"id"`
	Author                  string        `json:"author"`
	Permlink                string        `json:"permlink"`
	Category                string        `json:"category"`
	ParentAuthor            string        `json:"parent_author"`
	ParentPermlink          string        `json:"parent_permlink"`
	Title                   string        `json:"title"`
	Body                    string        `json:"body"`
	JSONMetadata            string        `json:"json_metadata"`
	LastUpdate              *types.Time   `json:"last_update"`
	Created                 *types.Time   `json:"created"`
	Active                  *types.Time   `json:"active"`
	LastPayout              *types.Time   `json:"last_payout"`
	Depth                   *types.Int    `json:"depth"`
	Children                *types.Int    `json:"children"`
	ChildrenRshares2        *types.Int    `json:"children_rshares2"`
	NetRshares              *types.Int    `json:"net_rshares"`
	AbsRshares              *types.Int    `json:"abs_rshares"`
	VoteRshares             *types.Int    `json:"vote_rshares"`
	ChildrenAbsRshares      *types.Int    `json:"children_abs_rshares"`
	CashoutTime             *types.Time   `json:"cashout_time"`
	MaxCashoutTime          *types.Time   `json:"max_cashout_time"`
	TotalVoteWeight         *types.Int    `json:"total_vote_weight"`
	RewardWeight            *types.Int    `json:"reward_weight"`
	TotalPayoutValue        *types.Asset  `json:"total_payout_value"`
	CuratorPayoutValue      *types.Asset  `json:"curator_payout_value"`
	AuthorRewards           *types.Int    `json:"author_rewards"`
	NetVotes                *types.Int    `json:"net_votes"`
	RootComment             *types.Int    `json:"root_comment"`
	Mode                    string        `json:"mode"`
	MaxAcceptedPayout       *types.Asset  `json:"max_accepted_payout"`
	AllowReplies            bool          `json:"allow_replies"`
	AllowVotes              bool          `json:"allow_votes"`
	AllowCurationRewards    bool          `json:"allow_curation_rewards"`
	URL                     string        `json:"url"`
	RootTitle               string        `json:"root_title"`
	PendingPayoutValue      *types.Asset  `json:"pending_payout_value"`
	TotalPendingPayoutValue *types.Asset  `json:"total_pending_payout_value"`
	ActiveVotes             []*VoteState  `json:"active_votes"`
	ActiveVotesCount        uint32        `json:"active_votes_count"`
	Replies                 []*Content    `json:"replies"`
	AuthorReputation        *types.Int    `json:"author_reputation"`
	Promoted                *types.Asset  `json:"promoted"`
	BodyLength              *types.Int    `json:"body_length"`
	RebloggedBy             []interface{} `json:"reblogged_by"`
}

//IsStory operation that determines that Content is a publication.
func (content *Content) IsStory() bool {
	return content.ParentAuthor == ""
}

//DiscussionQuery structure used in queries.
type DiscussionQuery struct {
	Tag            string   `json:"tag"`
	Limit          uint32   `json:"limit"`
	VoteLimit      uint32   `json:"vote_limit"`
	FilterTags     []string `json:"filter_tags,omitempty"`
	StartAuthor    string   `json:"start_author,omitempty"`
	StartPermlink  string   `json:"start_permlink,omitempty"`
	ParentAuthor   string   `json:"parent_author,omitempty"`
	ParentPermlink string   `json:"parent_permlink,omitempty"`
}

//TrendingTags structure for the GetTrendingTags function.
type TrendingTags struct {
	Name                  string       `json:"name"`
	TotalChildrenRshares2 string       `json:"total_children_rshares2"`
	TotalPayouts          *types.Asset `json:"total_payouts"`
	NetVotes              *types.Int   `json:"net_votes"`
	TopPosts              *types.Int   `json:"top_posts"`
	Comments              *types.Int   `json:"comments"`
}

//WitnessSchedule structure for the GetWitnessSchedule function.
type WitnessSchedule struct {
	ID                       *types.Int             `json:"id"`
	CurrentVirtualTime       *types.UInt64          `json:"current_virtual_time"`
	NextShuffleBlockNum      uint32                 `json:"next_shuffle_block_num"`
	CurrentShuffledWitnesses string                 `json:"current_shuffled_witnesses"`
	NumScheduledWitnesses    uint8                  `json:"num_scheduled_witnesses"`
	MedianProps              *types.ChainProperties `json:"median_props"`
	MajorityVersion          string                 `json:"majority_version"`
}

//Witness structure for the GetWitnessByAccount, GetWitnesses and GetWitnessByVote function.
type Witness struct {
	ID                    *types.Int             `json:"id"`
	Owner                 string                 `json:"owner"`
	Created               *types.Time            `json:"created"`
	URL                   string                 `json:"url"`
	TotalMissed           uint32                 `json:"total_missed"`
	LastAslot             uint64                 `json:"last_aslot"`
	LastConfirmedBlockNum uint64                 `json:"last_confirmed_block_num"`
	SigningKey            string                 `json:"signing_key"`
	Props                 *types.ChainProperties `json:"props"`
	Votes                 int64                  `json:"votes"`
	VirtualLastUpdate     *types.UInt64          `json:"virtual_last_update"`
	VirtualPosition       *types.UInt64          `json:"virtual_position"`
	VirtualScheduledTime  *types.UInt64          `json:"virtual_scheduled_time"`
	LastWork              string                 `json:"last_work"`
	RunningVersion        string                 `json:"running_version"`
	HardforkVersionVote   string                 `json:"hardfork_version_vote"`
	HardforkTimeVote      *types.Time            `json:"hardfork_time_vote"`
}

//SubscriptionState structure for the GetPaidSubscriptionOptions function.
type SubscriptionState struct {
	Creator                                       string      `json:"creator"`
	Url                                           string      `json:"url"`
	Levels                                        uint16      `json:"levels"`
	Amount                                        int64       `json:"amount"`
	Period                                        uint16      `json:"period"`
	UpdateTime                                    *types.Time `json:"update_time"`
	ActiveSubscribers                             []string    `json:"active_subscribers"`
	ActiveSubscribersCount                        uint32      `json:"active_subscribers_count"`
	ActiveSubscribersSummaryAmount                int64       `json:"active_subscribers_summary_amount"`
	ActiveSubscribersWithAutoRenewal              []string    `json:"active_subscribers_with_auto_renewal"`
	ActiveSubscribersWithAutoRenewalCount         uint32      `json:"active_subscribers_with_auto_renewal_count"`
	ActiveSubscribersWithAutoRenewalSummaryAmount int64       `json:"active_subscribers_with_auto_renewal_summary_amount"`
}

//SubscribeState structure for the GetPaidSubscriptionStatus function.
type SubscribeState struct {
	Subscriber  string      `json:"subscriber"`
	Creator     string      `json:"creator"`
	Level       uint16      `json:"level"`
	Amount      int64       `json:"amount"`
	Period      uint16      `json:"period"`
	StartTime   *types.Time `json:"start_time"`
	NextTime    *types.Time `json:"next_time"`
	EndTime     *types.Time `json:"end_time"`
	Active      bool        `json:"active"`
	AutoRenewal bool        `json:"auto_renewal"`
}
