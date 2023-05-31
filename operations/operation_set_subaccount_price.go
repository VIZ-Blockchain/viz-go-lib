package operations

import (
	"github.com/biter777/viz-go-lib/encoding/transaction"
	"github.com/biter777/viz-go-lib/types"
)

// SetSubaccountPriceOperation represents set_subaccount_price operation data.
type SetSubaccountPriceOperation struct {
	Account              string       `json:"account"`
	SubAccountSeller     string       `json:"subaccount_seller"`
	SubAccountOfferPrice *types.Asset `json:"subaccount_offer_price"`
	SubAccountOnSale     bool         `json:"subaccount_on_sale"`
}

// Type function that defines the type of operation SetSubaccountPriceOperation.
func (op *SetSubaccountPriceOperation) Type() OpType {
	return TypeSetSubaccountPrice
}

// Data returns the operation data SetSubaccountPriceOperation.
func (op *SetSubaccountPriceOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type SetSubaccountPriceOperation to bytes.
func (op *SetSubaccountPriceOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetSubaccountPrice.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.SubAccountSeller)
	enc.Encode(op.SubAccountOfferPrice)
	enc.Encode(op.SubAccountOnSale)
	return enc.Err()
}
