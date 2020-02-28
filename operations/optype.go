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
	TypeWitnessUpdate                  OpType = "witness_update"
	TypeAccountWitnessVote             OpType = "account_witness_vote"
	TypeAccountWitnessProxy            OpType = "account_witness_proxy"
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
	TypeShutdownWitness                OpType = "shutdown_witness"          //Virtual Operation
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
	TypeWitnessReward                  OpType = "witness_reward"            //Virtual Operation
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
)

var opTypes = [...]OpType{
	TypeVote,
	TypeContent,
	TypeTransfer,
	TypeTransferToVesting,
	TypeWithdrawVesting,
	TypeAccountUpdate,
	TypeWitnessUpdate,
	TypeAccountWitnessVote,
	TypeAccountWitnessProxy,
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
	TypeShutdownWitness,         //Virtual Operation
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
	TypeWitnessReward,           //Virtual Operation
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
}

// opCodes keeps mapping operation type -> operation code.
var opCodes map[OpType]uint16

func init() {
	opCodes = make(map[OpType]uint16, len(opTypes))
	for i, opType := range opTypes {
		opCodes[opType] = uint16(i)
	}
}
