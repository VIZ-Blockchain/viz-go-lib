package types

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"

	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/pkg/errors"
)

var (
	TypeFollow         = "follow"
	TypeReblog         = "reblog"
	TypePrivateMessage = "private_message"
)

var customDataObjects = map[string]interface{}{
	TypeFollow:         &FollowOperation{},
	TypeReblog:         &ReblogOperation{},
	TypePrivateMessage: &PrivateMessageOperation{},
}

//CustomOperation represents custom operation data.
type CustomOperation struct {
	RequiredAuths        []string `json:"required_auths"`
	RequiredPostingAuths []string `json:"required_posting_auths"`
	ID                   string   `json:"id"`
	JSON                 string   `json:"json"`
}

//Type function that defines the type of operation.
func (op *CustomOperation) Type() OpType {
	return TypeCustom
}

//Data returns the operation data.
func (op *CustomOperation) Data() interface{} {
	return op
}

//UnmarshalData unpacking the JSON parameter in the CustomOperation type.
func (op *CustomOperation) UnmarshalData() (interface{}, error) {
	template, ok := customDataObjects[op.ID]
	if !ok {
		return nil, nil
	}

	opData := reflect.New(reflect.Indirect(reflect.ValueOf(template)).Type()).Interface()

	var bodyReader io.Reader
	if op.JSON[0] == '[' {
		rawTuple := make([]json.RawMessage, 2)
		if err := json.NewDecoder(strings.NewReader(op.JSON)).Decode(&rawTuple); err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal CustomOperation.JSON: \n%v", op.JSON)
		}
		if len(rawTuple) < 2 || rawTuple[1] == nil {
			return nil, errors.Errorf("invalid CustomOperation.JSON: \n%v", op.JSON)
		}
		bodyReader = bytes.NewReader([]byte(rawTuple[1]))
	} else {
		bodyReader = strings.NewReader(op.JSON)
	}

	// Unmarshal into the new object instance.
	if err := json.NewDecoder(bodyReader).Decode(opData); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal CustomOperation.JSON: \n%v", op.JSON)
	}

	return opData, nil
}

//MarshalTransaction is a function of converting type CustomOperation to bytes.
func (op *CustomOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCustom.Code()))
	enc.EncodeArrString(op.RequiredAuths)
	enc.EncodeArrString(op.RequiredPostingAuths)
	enc.Encode(op.ID)
	enc.Encode(op.JSON)
	return enc.Err()
}

//FollowOperation the structure for the operation CustomOperation.
type FollowOperation struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

//ReblogOperation the structure for the operation CustomOperation.
type ReblogOperation struct {
	Account  string `json:"account"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//PrivateMessageOperation the structure for the operation CustomOperation.
type PrivateMessageOperation struct {
	From             string `json:"from"`
	To               string `json:"to"`
	FromMemoKey      string `json:"from_memo_key"`
	ToMemoKey        string `json:"to_memo_key"`
	SentTime         uint64 `json:"sent_time"`
	Checksum         uint32 `json:"checksum"`
	EncryptedMessage string `json:"encrypted_message"`
}

//MarshalCustom generate a row from the structure fields.
func MarshalCustom(v interface{}) (string, error) {
	var tmp []interface{}

	typeInterface := reflect.TypeOf(v).Name()
	switch typeInterface {
	case "FollowOperation":
		tmp = append(tmp, TypeFollow)
	case "ReblogOperation":
		tmp = append(tmp, TypeReblog)
	case "PrivateMessageOperation":
		tmp = append(tmp, TypePrivateMessage)
	default:
		return "", errors.New("Unknown type")
	}

	tmp = append(tmp, v)

	b, err := json.Marshal(tmp)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
