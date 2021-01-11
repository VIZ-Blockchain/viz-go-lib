package wif

import (
	"bytes"

	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
)

// WifIsValid check that private key conform to public key
func WifIsValid(privateWif, publicWif string) (bool, error) {
	publicFromPrivate, err := GetPublicKey(privateWif)
	if err != nil {
		return false, err
	}
	var encodedPublicWif bytes.Buffer
	encoder := transaction.NewEncoder(&encodedPublicWif)
	if err := encoder.EncodePubKey(publicWif); err != nil {
		return false, err
	}
	return bytes.Compare(publicFromPrivate, encodedPublicWif.Bytes()) == 0, nil
}
