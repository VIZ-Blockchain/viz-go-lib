package types

//BenefactorAwardOperation represents benefactor_award operation data.
type BenefactorAwardOperation struct {
	Benefactor     string `json:"benefactor"`
	Receiver       string `json:"receiver"`
	CustomSequence uint64 `json:"custom_sequence"`
	Memo           string `json:"memo"`
	Shares         *Asset `json:"shares"`
}

//Type function that defines the type of operation BenefactorAwardOperation.
func (op *BenefactorAwardOperation) Type() OpType {
	return TypeBenefactorAward
}

//Data returns the operation data BenefactorAwardOperation.
func (op *BenefactorAwardOperation) Data() interface{} {
	return op
}
