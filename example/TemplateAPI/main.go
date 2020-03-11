package main

import (
	"fmt"

	"github.com/VIZ-Blockchain/viz-go-lib"
)

func main() {
	cls, _ := viz.NewClient("wss://solox.world/ws")
	defer cls.Close()

	ans, err := cls.API.GetOpsInBlock(35137873, true)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Answer :")
		for k, v := range ans {
			fmt.Printf("OP : %d -> %+v \n", k, v)
		}
	}
}
