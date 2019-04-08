package client

import (
	crand "crypto/rand"
	"crypto/sha256"
	"math/rand"
	"time"

	"github.com/VIZ-Blockchain/viz-go-lib/encoding/wif"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ripemd160"
)

var (
	//OpTypeKey include a description of the operation and the key needed to sign it
	OpTypeKey = make(map[types.OpType][]string)
	src       rand.Source
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

//Keys is used as a keystroke for a specific user.
//Only a few keys can be set.
type Keys struct {
	PKey []string
	AKey []string
	OKey []string
	MKey []string
}

func init() {
	OpTypeKey["vote"] = []string{"regular"}
	OpTypeKey["content"] = []string{"regular"}
	OpTypeKey["transfer"] = []string{"active"}
	OpTypeKey["transfer_to_vesting"] = []string{"active"}
	OpTypeKey["withdraw_vesting"] = []string{"active"}
	OpTypeKey["account_update"] = []string{"active"}
	OpTypeKey["witness_update"] = []string{"active"}
	OpTypeKey["account_witness_vote"] = []string{"regular"}
	OpTypeKey["account_witness_proxy"] = []string{"regular"}
	OpTypeKey["delete_content"] = []string{"regular"}
	OpTypeKey["custom"] = []string{"regular"}
	OpTypeKey["set_withdraw_vesting_route"] = []string{"active"}
	OpTypeKey["request_account_recovery"] = []string{"active"}
	OpTypeKey["recover_account"] = []string{"owner"}
	OpTypeKey["change_recovery_account"] = []string{"owner"}
	OpTypeKey["escrow_transfer"] = []string{"active"}
	OpTypeKey["escrow_dispute"] = []string{"active"}
	OpTypeKey["escrow_release"] = []string{"active"}
	OpTypeKey["escrow_approve"] = []string{"active"}
	OpTypeKey["delegate_vesting_shares"] = []string{"active"}
	OpTypeKey["account_create"] = []string{"active"}
	OpTypeKey["account_metadata"] = []string{"regular"}
	OpTypeKey["proposal_create"] = []string{"active"}
	OpTypeKey["proposal_update"] = []string{"active"}
	OpTypeKey["proposal_delete"] = []string{"active"}
	OpTypeKey["chain_properties_update"] = []string{"active"}
	OpTypeKey["committee_worker_create_request"] = []string{"regular"}
	OpTypeKey["committee_worker_cancel_request"] = []string{"regular"}
	OpTypeKey["committee_vote_request"] = []string{"regular"}
	OpTypeKey["create_invite"] = []string{"active"}
	OpTypeKey["claim_invite_balance"] = []string{"active"}
	OpTypeKey["invite_registration"] = []string{"active"}
	OpTypeKey["versioned_chain_properties_update"] = []string{"active"}
	OpTypeKey["award"] = []string{"regular"}
	OpTypeKey["set_paid_subscription"] = []string{"active"}
	OpTypeKey["paid_subscribe"] = []string{"active"}

	seed := time.Now().UnixNano()

	reader := crand.Reader
	i, err := crand.Prime(reader, 64)
	if err != nil {
		seed = seed ^ i.Int64()
	}

	src = rand.NewSource(seed)
}

//SetKeys you can specify keys for signing transactions.
func (client *Client) SetKeys(keys *Keys) {
	client.CurrentKeys = keys
}

//SigningKeys returns the key from the CurrentKeys
func (client *Client) SigningKeys(trx types.Operation) ([][]byte, error) {
	var keys [][]byte

	if client.CurrentKeys == nil {
		return nil, errors.New("Client Keys not initialized. Use SetKeys method")
	}

	opKeys := OpTypeKey[trx.Type()]
	for _, val := range opKeys {
		switch val {
		case "regular":
			for _, keyStr := range client.CurrentKeys.PKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Regular Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "active":
			for _, keyStr := range client.CurrentKeys.AKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Active Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "owner":
			for _, keyStr := range client.CurrentKeys.OKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Owner Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		case "memo":
			for _, keyStr := range client.CurrentKeys.MKey {
				privKey, err := wif.Decode(keyStr)
				if err != nil {
					return nil, errors.New("error decode Memo Key: " + err.Error())
				}
				keys = append(keys, privKey)
			}
		}
	}

	return keys, nil
}

//GenPassword allows to generate a 52-character password of the evil system.
func GenPassword() string {
	b := make([]byte, 51)
	for i, cache, remain := 51-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return "P" + string(b)
}

//GetPrivateKey generates a private key based on the specified parameters.
func GetPrivateKey(user, role, password string) string {
	hashSha256 := sha256.Sum256([]byte(user + role + password))
	pk := append([]byte{0x80}, hashSha256[:]...)
	chs := sha256.Sum256(pk)
	chs = sha256.Sum256(chs[:])
	b58 := append(pk, chs[:4]...)
	return base58.Encode(b58)
}

//GetPublicKey generates a public key based on the prefix and the private key.
func GetPublicKey(prefix, privatekey string) string {
	b58 := base58.Decode(privatekey)
	tpk := b58[:len(b58)-4]
	chs := b58[len(b58)-4:]
	nchs := sha256.Sum256(tpk)
	nchs = sha256.Sum256(nchs[:])
	if string(chs) != string(nchs[:4]) {
		return "Invalid WIF key (checksum miss-match)"
	}
	privKeyBytes := [32]byte{}
	copy(privKeyBytes[:], tpk[1:])
	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes[:])
	chHash := ripemd160.New()
	_, errHash := chHash.Write(priv.PubKey().SerializeCompressed())
	if errHash != nil {
		return errHash.Error()
	}
	nc := chHash.Sum(nil)
	pk := append(priv.PubKey().SerializeCompressed(), nc[:4]...)
	return prefix + base58.Encode(pk)
}
