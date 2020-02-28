package operations

import (
	"encoding/json"

	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

//ContentOperation represents content operation data.
type ContentOperation struct {
	ParentAuthor   string        `json:"parent_author"`
	ParentPermlink string        `json:"parent_permlink"`
	Author         string        `json:"author"`
	Permlink       string        `json:"permlink"`
	Title          string        `json:"title"`
	Body           string        `json:"body"`
	JSONMetadata   string        `json:"json_metadata"`
	Extensions     []interface{} `json:"extensions"`
}

//Type function that defines the type of operation ContentOperation.
func (op *ContentOperation) Type() OpType {
	return TypeContent
}

//Data returns the operation data ContentOperation.
func (op *ContentOperation) Data() interface{} {
	return op
}

//IsStory function specifies the type of publication.
func (op *ContentOperation) IsStory() bool {
	return op.ParentAuthor == ""
}

//MarshalTransaction is a function of converting type ContentOperation to bytes.
func (op *ContentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeContent.Code()))
	if !op.IsStory() {
		enc.Encode(op.ParentAuthor)
	} else {
		enc.Encode(byte(0))
	}
	enc.Encode(op.ParentPermlink)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Title)
	enc.Encode(op.Body)
	enc.Encode(op.JSONMetadata)
	if len(op.Extensions) > 0 {
		//Parse Beneficiaries
		z, _ := json.Marshal(op.Extensions[0])
		var l []interface{}
		_ = json.Unmarshal(z, &l)
		z1, _ := json.Marshal(l[1])
		var d types.CommentPayoutBeneficiaries
		_ = json.Unmarshal(z1, &d)

		enc.Encode(byte(1))
		enc.Encode(d)
	} else {
		enc.Encode(byte(0))
	}
	return enc.Err()
}
