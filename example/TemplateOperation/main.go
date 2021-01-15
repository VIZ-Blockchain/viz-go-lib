package main

import (
	"fmt"

	"github.com/VIZ-Blockchain/viz-go-lib"
	"github.com/VIZ-Blockchain/viz-go-lib/operations"
	"github.com/VIZ-Blockchain/viz-go-lib/types"
)

func main() {
	client, _ := viz.NewClient("wss://solox.world/ws")
	defer client.Close()

	client.SetKeys(&viz.Keys{PKey: []string{"<Private Posting Key>"}})

	var trx []operations.Operation

	beneficiaries := []types.Beneficiary{
		{Account: "<BeneficiaryName1>", Weight: 500},  // 5%
		{Account: "<BeneficiaryName2>", Weight: 1500}, // 15%
	}
	tx := &operations.AwardOperation{
		Initiator:      "<UserName>",
		Receiver:       "<ReceiverName>",
		Energy:         2000, // 20%
		CustomSequence: 0,
		Memo:           "test",
		Beneficiaries:  beneficiaries,
	}
	trx = append(trx, tx)

	resp, err := client.SendTrx("<UserName>", trx)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Answer : ", resp)
	}
}
