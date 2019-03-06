package types

//Old

//VoteOperation represents vote operation data.
type VoteOperation struct {
	Voter    string `json:"voter"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Weight   Int16  `json:"weight"`
}

//Type function that defines the type of operation VoteOperation.
func (op *VoteOperation) Type() OpType {
	return TypeVote
}

//Data returns the operation data VoteOperation.
func (op *VoteOperation) Data() interface{} {
	return op
}
