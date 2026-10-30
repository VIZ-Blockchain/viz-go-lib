package operations

// OpType represents a Golos operation type, i.e. vote, content and so on.
type OpType string

// Code returns the operation code associated with the given operation type.
func (kind OpType) Code() uint16 {
	return opCodes[kind]
}

const (
	TypeVote                           OpType = "vote"
	TypeContent                        OpType = "content"
	TypeTransfer                       OpType = "transfer"
	TypeTransferToVesting              OpType = "transfer_to_vesting"
	TypeWithdrawVesting                OpType = "withdraw_vesting"
	TypeAccountUpdate                  OpType = "account_update"
	TypeValidatorUpdate                OpType = "validator_update"
	TypeAccountValidatorVote           OpType = "account_validator_vote"
	TypeAccountValidatorProxy          OpType = "account_validator_proxy"
	TypeDeleteContent                  OpType = "delete_content"
	TypeCustom                         OpType = "custom"
	TypeSetWithdrawVestingRoute        OpType = "set_withdraw_vesting_route"
	TypeRequestAccountRecovery         OpType = "request_account_recovery"
	TypeRecoverAccount                 OpType = "recover_account"
	TypeChangeRecoveryAccount          OpType = "change_recovery_account"
	TypeEscrowTransfer                 OpType = "escrow_transfer"
	TypeEscrowDispute                  OpType = "escrow_dispute"
	TypeEscrowRelease                  OpType = "escrow_release"
	TypeEscrowApprove                  OpType = "escrow_approve"
	TypeDelegateVestingShares          OpType = "delegate_vesting_shares"
	TypeAccountCreate                  OpType = "account_create"
	TypeAccountMetadata                OpType = "account_metadata"
	TypeProposalCreate                 OpType = "proposal_create"
	TypeProposalUpdate                 OpType = "proposal_update"
	TypeProposalDelete                 OpType = "proposal_delete"
	TypeChainPropertiesUpdate          OpType = "chain_properties_update"
	TypeAuthorReward                   OpType = "author_reward"             //Virtual Operation
	TypeCurationReward                 OpType = "curation_reward"           //Virtual Operation
	TypeContentReward                  OpType = "content_reward"            //Virtual Operation
	TypeFillVestingWithdraw            OpType = "fill_vesting_withdraw"     //Virtual Operation
	TypeShutdownValidator              OpType = "shutdown_validator"        //Virtual Operation
	TypeHardfork                       OpType = "hardfork"                  //Virtual Operation
	TypeContentPayoutUpdate            OpType = "content_payout_update"     //Virtual Operation
	TypeContentBenefactorReward        OpType = "content_benefactor_reward" //Virtual Operation
	TypeReturnVestingDelegation        OpType = "return_vesting_delegation" //Virtual Operation
	TypeCommitteeWorkerCreateRequest   OpType = "committee_worker_create_request"
	TypeCommitteeWorkerCancelRequest   OpType = "committee_worker_cancel_request"
	TypeCommitteeVoteRequest           OpType = "committee_vote_request"
	TypeCommitteeCancelRequest         OpType = "committee_cancel_request"  //Virtual Operation
	TypeCommitteeApproveRequest        OpType = "committee_approve_request" //Virtual Operation
	TypeCommitteePayoutRequest         OpType = "committee_payout_request"  //Virtual Operation
	TypeCommitteePayRequest            OpType = "committee_pay_request"     //Virtual Operation
	TypeValidatorReward                OpType = "validator_reward"          //Virtual Operation
	TypeCreateInvite                   OpType = "create_invite"
	TypeClaimInviteBalance             OpType = "claim_invite_balance"
	TypeInviteRegistration             OpType = "invite_registration"
	TypeVersionedChainPropertiesUpdate OpType = "versioned_chain_properties_update"
	TypeAward                          OpType = "award"
	TypeReceiveAward                   OpType = "receive_award"    //Virtual Operation
	TypeBenefactorAward                OpType = "benefactor_award" //Virtual Operation
	TypeSetPaidSubscription            OpType = "set_paid_subscription"
	TypePaidSubscribe                  OpType = "paid_subscribe"
	TypePaidSubscriptionAction         OpType = "paid_subscription_action" //Virtual Operation
	TypeCancelPaidSubscription         OpType = "cancel_paid_subscription" //Virtual Operation
	TypeSetAccountPrice                OpType = "set_account_price"
	TypeSetSubaccountPrice             OpType = "set_subaccount_price"
	TypeBuyAccount                     OpType = "buy_account"
	TypeAccountSale                    OpType = "account_sale" //Virtual Operation

	TypeUseInviteBalance               OpType = "use_invite_balance"
	TypeExpireEscrowRatification       OpType = "expire_escrow_ratification" //Virtual Operation
	TypeFixedAward                     OpType = "fixed_award"
	TypeTargetAccountSale              OpType = "target_account_sale"
	TypeBid                            OpType = "bid"    //Virtual Operation
	TypeOutbid                         OpType = "outbid" //Virtual Operation
	TypeSetRewardSharing               OpType = "set_reward_sharing"
	TypeStakeholderReward              OpType = "stakeholder_reward" //Virtual Operation

	// HF14 predictions markets (wire names are the consensus op names, append only).
	TypePMOracleRegister        OpType = "pm_oracle_register"
	TypePMOracleUpdate          OpType = "pm_oracle_update"
	TypePMCreateMarket          OpType = "pm_create_market"
	TypePMOracleAcceptMarket    OpType = "pm_oracle_accept_market"
	TypePMPlaceBet              OpType = "pm_place_bet"
	TypePMCommitBet             OpType = "pm_commit_bet"
	TypePMRevealBet             OpType = "pm_reveal_bet"
	TypePMCancelBet             OpType = "pm_cancel_bet"
	TypePMAddLiquidity          OpType = "pm_add_liquidity"
	TypePMWithdrawLiquidity     OpType = "pm_withdraw_liquidity"
	TypePMResolveMarket         OpType = "pm_resolve_market"
	TypePMNoContest             OpType = "pm_no_contest"
	TypePMDisputeCreate         OpType = "pm_dispute_create"
	TypePMDisputeVote           OpType = "pm_dispute_vote"
	TypePMDisputeResolve        OpType = "pm_dispute_resolve"
	TypePMTransferPosition      OpType = "pm_transfer_position"
	TypePMLazyDeposit           OpType = "pm_lazy_deposit"
	TypePMLazyWithdraw          OpType = "pm_lazy_withdraw"
	TypePMBatchSettle           OpType = "pm_batch_settle"            //Virtual Operation
	TypePMCommitForfeit         OpType = "pm_commit_forfeit"          //Virtual Operation
	TypePMAutoPayout            OpType = "pm_auto_payout"             //Virtual Operation
	TypePMDisputeFinalize       OpType = "pm_dispute_finalize"        //Virtual Operation
	TypePMDisputeAutoClose      OpType = "pm_dispute_auto_close"      //Virtual Operation
	TypePMOracleMissedPenalty   OpType = "pm_oracle_missed_penalty"   //Virtual Operation
	TypePMLazyRecall            OpType = "pm_lazy_recall"             //Virtual Operation
	TypePMLeverageOpen          OpType = "pm_leverage_open"
	TypePMLeverageClose         OpType = "pm_leverage_close"
	TypePMLeverageConvert       OpType = "pm_leverage_convert"
	TypePMLeverageLiquidate     OpType = "pm_leverage_liquidate"      //Virtual Operation
	TypePMLeverageResolve       OpType = "pm_leverage_resolve"        //Virtual Operation
	TypePMMarketAccepted        OpType = "pm_market_accepted"         //Virtual Operation
	TypePMPayout                OpType = "pm_payout"                  //Virtual Operation
	TypePMDisputeOracleRespond  OpType = "pm_dispute_oracle_respond"
	TypePMUnban                 OpType = "pm_unban"
	TypePMBanExpired            OpType = "pm_ban_expired"             //Virtual Operation
	TypePMMarketExpired         OpType = "pm_market_expired"          //Virtual Operation
	TypePMDisputeOpened         OpType = "pm_dispute_opened"          //Virtual Operation
	TypePMEarlyExitClaimPaid    OpType = "pm_early_exit_claim_paid"   //Virtual Operation
	TypePMLPPayout              OpType = "pm_lp_payout"               //Virtual Operation

	// HF15 agent access.
	TypeSetAgentPermission OpType = "set_agent_permission"

	// HF16 QC.
	TypeValidatorEquivocation OpType = "validator_equivocation" //Virtual Operation

	// Deprecated: the chain renamed the validator operations in 2026; these aliases keep
	// existing source code compiling. They carry the NEW wire names.
	TypeWitnessUpdate        OpType = TypeValidatorUpdate
	TypeAccountWitnessVote   OpType = TypeAccountValidatorVote
	TypeAccountWitnessProxy  OpType = TypeAccountValidatorProxy
	TypeShutdownWitness      OpType = TypeShutdownValidator
	TypeWitnessReward        OpType = TypeValidatorReward
)

var opTypes = [...]OpType{
	TypeVote,
	TypeContent,
	TypeTransfer,
	TypeTransferToVesting,
	TypeWithdrawVesting,
	TypeAccountUpdate,
	TypeValidatorUpdate,
	TypeAccountValidatorVote,
	TypeAccountValidatorProxy,
	TypeDeleteContent,
	TypeCustom,
	TypeSetWithdrawVestingRoute,
	TypeRequestAccountRecovery,
	TypeRecoverAccount,
	TypeChangeRecoveryAccount,
	TypeEscrowTransfer,
	TypeEscrowDispute,
	TypeEscrowRelease,
	TypeEscrowApprove,
	TypeDelegateVestingShares,
	TypeAccountCreate,
	TypeAccountMetadata,
	TypeProposalCreate,
	TypeProposalUpdate,
	TypeProposalDelete,
	TypeChainPropertiesUpdate,
	TypeAuthorReward,            //Virtual Operation
	TypeCurationReward,          //Virtual Operation
	TypeContentReward,           //Virtual Operation
	TypeFillVestingWithdraw,     //Virtual Operation
	TypeShutdownValidator,       //Virtual Operation
	TypeHardfork,                //Virtual Operation
	TypeContentPayoutUpdate,     //Virtual Operation
	TypeContentBenefactorReward, //Virtual Operation
	TypeReturnVestingDelegation, //Virtual Operation
	TypeCommitteeWorkerCreateRequest,
	TypeCommitteeWorkerCancelRequest,
	TypeCommitteeVoteRequest,
	TypeCommitteeCancelRequest,  //Virtual Operation
	TypeCommitteeApproveRequest, //Virtual Operation
	TypeCommitteePayoutRequest,  //Virtual Operation
	TypeCommitteePayRequest,     //Virtual Operation
	TypeValidatorReward,         //Virtual Operation
	TypeCreateInvite,
	TypeClaimInviteBalance,
	TypeInviteRegistration,
	TypeVersionedChainPropertiesUpdate,
	TypeAward,
	TypeReceiveAward,    //Virtual Operation
	TypeBenefactorAward, //Virtual Operation
	TypeSetPaidSubscription,
	TypePaidSubscribe,
	TypePaidSubscriptionAction, //Virtual Operation
	TypeCancelPaidSubscription, //Virtual Operation
	TypeSetAccountPrice,
	TypeSetSubaccountPrice,
	TypeBuyAccount,
	TypeAccountSale, //Virtual Operation

	TypeUseInviteBalance,
	TypeExpireEscrowRatification, //Virtual Operation
	TypeFixedAward,
	TypeTargetAccountSale,
	TypeBid,    //Virtual Operation
	TypeOutbid, //Virtual Operation
	TypeSetRewardSharing,
	TypeStakeholderReward, //Virtual Operation

	TypePMOracleRegister,
	TypePMOracleUpdate,
	TypePMCreateMarket,
	TypePMOracleAcceptMarket,
	TypePMPlaceBet,
	TypePMCommitBet,
	TypePMRevealBet,
	TypePMCancelBet,
	TypePMAddLiquidity,
	TypePMWithdrawLiquidity,
	TypePMResolveMarket,
	TypePMNoContest,
	TypePMDisputeCreate,
	TypePMDisputeVote,
	TypePMDisputeResolve,
	TypePMTransferPosition,
	TypePMLazyDeposit,
	TypePMLazyWithdraw,
	TypePMBatchSettle,          //Virtual Operation
	TypePMCommitForfeit,        //Virtual Operation
	TypePMAutoPayout,           //Virtual Operation
	TypePMDisputeFinalize,      //Virtual Operation
	TypePMDisputeAutoClose,     //Virtual Operation
	TypePMOracleMissedPenalty,  //Virtual Operation
	TypePMLazyRecall,           //Virtual Operation
	TypePMLeverageOpen,
	TypePMLeverageClose,
	TypePMLeverageConvert,
	TypePMLeverageLiquidate,    //Virtual Operation
	TypePMLeverageResolve,      //Virtual Operation
	TypePMMarketAccepted,       //Virtual Operation
	TypePMPayout,               //Virtual Operation
	TypePMDisputeOracleRespond,
	TypePMUnban,
	TypePMBanExpired,           //Virtual Operation
	TypePMMarketExpired,        //Virtual Operation
	TypePMDisputeOpened,        //Virtual Operation
	TypePMEarlyExitClaimPaid,   //Virtual Operation
	TypePMLPPayout,             //Virtual Operation

	TypeSetAgentPermission,

	TypeValidatorEquivocation, //Virtual Operation
}

// opCodes keeps mapping operation type -> operation code.
var opCodes map[OpType]uint16

func init() {
	opCodes = make(map[OpType]uint16, len(opTypes))
	for i, opType := range opTypes {
		opCodes[opType] = uint16(i)
	}
}
