package types

//Old

//ContentOperation represents content operation data.
type ContentOperation struct {
	ParentAuthor    string        `json:"parent_author"`
	ParentPermlink  string        `json:"parent_permlink"`
	Author          string        `json:"author"`
	Permlink        string        `json:"permlink"`
	Title           string        `json:"title"`
	Body            string        `json:"body"`
	CurationPercent int16         `json:"curation_percent"`
	JSONMetadata    string        `json:"json_metadata"`
	Extensions      []interface{} `json:"extensions"`
}

//Type function that defines the type of operation ContentOperation.
func (op *ContentOperation) Type() OpType {
	return TypeContent
}

//Data returns the operation data ContentOperation.
func (op *ContentOperation) Data() interface{} {
	return op
}

//IsStoryOperation function specifies the type of publication.
func (op *ContentOperation) IsStoryOperation() bool {
	return op.ParentAuthor == ""
}
