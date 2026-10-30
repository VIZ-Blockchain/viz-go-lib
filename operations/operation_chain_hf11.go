package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

// UseInviteBalanceOperation represents use_invite_balance operation data.
type UseInviteBalanceOperation struct {
	Initiator    string `json:"initiator"`
	Receiver     string `json:"receiver"`
	InviteSecret string `json:"invite_secret"`
}

// Type function that defines the type of operation UseInviteBalanceOperation.
func (op *UseInviteBalanceOperation) Type() OpType {
	return TypeUseInviteBalance
}

// Data returns the operation data UseInviteBalanceOperation.
func (op *UseInviteBalanceOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type UseInviteBalanceOperation to bytes.
func (op *UseInviteBalanceOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeUseInviteBalance.Code()))
	enc.Encode(op.Initiator)
	enc.Encode(op.Receiver)
	enc.Encode(op.InviteSecret)
	return enc.Err()
}

// FixedAwardOperation represents fixed_award operation data.
type FixedAwardOperation struct {
	Initiator      string              `json:"initiator"`
	Receiver       string              `json:"receiver"`
	RewardAmount   *types.Asset        `json:"reward_amount"`
	MaxEnergy      uint16              `json:"max_energy"`
	CustomSequence uint64              `json:"custom_sequence"`
	Memo           string              `json:"memo"`
	Beneficiaries  []types.Beneficiary `json:"beneficiaries"`
}

// Type function that defines the type of operation FixedAwardOperation.
func (op *FixedAwardOperation) Type() OpType {
	return TypeFixedAward
}

// Data returns the operation data FixedAwardOperation.
func (op *FixedAwardOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type FixedAwardOperation to bytes.
func (op *FixedAwardOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeFixedAward.Code()))
	enc.Encode(op.Initiator)
	enc.Encode(op.Receiver)
	enc.Encode(op.RewardAmount)
	enc.Encode(op.MaxEnergy)
	enc.Encode(op.CustomSequence)
	enc.Encode(op.Memo)
	enc.EncodeUVarint(uint64(len(op.Beneficiaries)))
	for _, val := range op.Beneficiaries {
		enc.Encode(val.Account)
		enc.Encode(val.Weight)
	}
	return enc.Err()
}

// TargetAccountSaleOperation represents target_account_sale operation data.
type TargetAccountSaleOperation struct {
	Account           string       `json:"account"`
	AccountSeller     string       `json:"account_seller"`
	TargetBuyer       string       `json:"target_buyer"`
	AccountOfferPrice *types.Asset `json:"account_offer_price"`
	AccountOnSale     bool         `json:"account_on_sale"`
}

// Type function that defines the type of operation TargetAccountSaleOperation.
func (op *TargetAccountSaleOperation) Type() OpType {
	return TypeTargetAccountSale
}

// Data returns the operation data TargetAccountSaleOperation.
func (op *TargetAccountSaleOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type TargetAccountSaleOperation to bytes.
func (op *TargetAccountSaleOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTargetAccountSale.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.AccountSeller)
	enc.Encode(op.TargetBuyer)
	enc.Encode(op.AccountOfferPrice)
	enc.Encode(op.AccountOnSale)
	return enc.Err()
}

// SetRewardSharingOperation represents set_reward_sharing operation data.
type SetRewardSharingOperation struct {
	Owner       string `json:"owner"`
	SharingRate uint16 `json:"sharing_rate"`
}

// Type function that defines the type of operation SetRewardSharingOperation.
func (op *SetRewardSharingOperation) Type() OpType {
	return TypeSetRewardSharing
}

// Data returns the operation data SetRewardSharingOperation.
func (op *SetRewardSharingOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type SetRewardSharingOperation to bytes.
func (op *SetRewardSharingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetRewardSharing.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.SharingRate)
	return enc.Err()
}
