package types

import (
	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

//Beneficiary is an additional structure used by other structures.
type Beneficiary struct {
	Account string `json:"account"`
	Weight  uint16 `json:"weight"`
}

//Beneficiaries is an additional structure used by other structures.
type Beneficiaries struct {
	Beneficiaries []Beneficiary `json:"beneficiaries"`
}

//MarshalTransaction is a function of converting type Beneficiaries to bytes.
func (op *Beneficiaries) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(len(op.Beneficiaries))
	for _, v := range op.Beneficiaries {
		enc.Encode(v.Account)
		enc.Encode(v.Weight)
	}
	return enc.Err()
}
