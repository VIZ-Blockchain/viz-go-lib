package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//SetAccountPriceOperation represents set_account_price operation data.
type SetAccountPriceOperation struct {
	Account           string       `json:"account"`
	AccountSeller     string       `json:"account_seller"`
	AccountOfferPrice *types.Asset `json:"account_offer_price"`
	AccountOnSale     bool         `json:"account_on_sale"`
}

//Type function that defines the type of operation SetAccountPriceOperation.
func (op *SetAccountPriceOperation) Type() OpType {
	return TypeSetAccountPrice
}

//Data returns the operation data SetAccountPriceOperation.
func (op *SetAccountPriceOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type SetAccountPriceOperation to bytes.
func (op *SetAccountPriceOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetAccountPrice.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.AccountSeller)
	enc.Encode(op.AccountOfferPrice)
	enc.Encode(op.AccountOnSale)
	return enc.Err()
}
