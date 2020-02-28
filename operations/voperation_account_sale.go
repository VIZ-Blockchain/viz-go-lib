package operations

import (
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//AccountSaleOperation represents account_sale operation data.
type AccountSaleOperation struct {
	Account string      `json:"account"`
	Price   types.Asset `json:"price"`
	Buyer   string      `json:"buyer"`
	Seller  string      `json:"seller"`
}

//Type function that defines the type of operation AccountSaleOperation.
func (op *AccountSaleOperation) Type() OpType {
	return TypeAccountSale
}

//Data returns the operation data AccountSaleOperation.
func (op *AccountSaleOperation) Data() interface{} {
	return op
}
