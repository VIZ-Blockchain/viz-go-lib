package types

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//SetSubaccountPriceOperation represents set_subaccount_price operation data.
type SetSubaccountPriceOperation struct {
	Account              string `json:"account"`
	SubaccountSeller     string `json:"subaccount_seller"`
	SubaccountOfferPrice *Asset `json:"subaccount_offer_price"`
	SubaccountOnSale     bool   `json:"subaccount_on_sale"`
}

//Type function that defines the type of operation SetSubaccountPriceOperation.
func (op *SetSubaccountPriceOperation) Type() OpType {
	return TypeSetSubaccountPrice
}

//Data returns the operation data SetSubaccountPriceOperation.
func (op *SetSubaccountPriceOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type SetSubaccountPriceOperation to bytes.
func (op *SetSubaccountPriceOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetSubaccountPrice.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.SubaccountSeller)
	enc.Encode(op.SubaccountOfferPrice)
	enc.EncodeBool(op.SubaccountOnSale)
	return enc.Err()
}
