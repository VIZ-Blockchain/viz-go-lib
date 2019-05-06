package types

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//BuyAccountOperation represents buy_account operation data.
type BuyAccountOperation struct {
	Buyer                 string `json:"buyer"`
	Account               string `json:"account"`
	AccountOfferPrice     *Asset `json:"account_offer_price"`
	AccountAuthoritiesKey string `json:"account_authorities_key"`
	TokensToShares        *Asset `json:"tokens_to_shares"`
}

//Type function that defines the type of operation BuyAccountOperation.
func (op *BuyAccountOperation) Type() OpType {
	return TypeBuyAccount
}

//Data returns the operation data BuyAccountOperation.
func (op *BuyAccountOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type BuyAccountOperation to bytes.
func (op *BuyAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeBuyAccount.Code()))
	enc.Encode(op.Buyer)
	enc.Encode(op.Account)
	enc.Encode(op.AccountOfferPrice)
	enc.EncodePubKey(op.AccountAuthoritiesKey)
	enc.Encode(op.TokensToShares)
	return enc.Err()
}
