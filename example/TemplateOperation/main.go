package main

import (
	"fmt"

	"github.com/VIZ-Blockchain/viz-go-lib"
	"github.com/VIZ-Blockchain/viz-go-lib/operations"
)

func main() {
	cls, _ := viz.NewClient("wss://solox.world/ws")
	defer cls.Close()

	cls.SetKeys(&viz.Keys{PKey: []string{"<Privat Posting Key>"}})

	var trx []types.Operation

	tx := &types.VoteOperation{
		Voter:    "<UserName>",
		Author:   "<AuthorName>",
		Permlink: "<Permlink>",
		Weight:   types.Int16(Weight),
	}
	trx = append(trx, tx)

	resp, err := client.SendTrx("<UserName>", trx)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Answer : ", resp)
	}
}
