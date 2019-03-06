package types

//DeleteContentOperation represents delete_content operation data.
type DeleteContentOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//Type function that defines the type of operation DeleteContentOperation.
func (op *DeleteContentOperation) Type() OpType {
	return TypeDeleteContent
}

//Data returns the operation data DeleteContentOperation.
func (op *DeleteContentOperation) Data() interface{} {
	return op
}
