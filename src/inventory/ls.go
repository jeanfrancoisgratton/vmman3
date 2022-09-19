// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/ls.go
// 2022-09-17 20:07:29

package inventory

import (
	"fmt"
	"vmman3/db"
	"vmman3/helpers"
)

func VM_Inventory() {
	var hyps []db.DbHypervisors
	if helpers.BAllHypervisors {
		hyps = listHypervisors()
	} else {
		hyps = []db.DbHypervisors{{HID: 0, Hname: "localhost", Haddress: "127.0.0.1"}}
	}

	for _, v := range hyps {
		fmt.Println("allo", v.Hname)
	}
}
