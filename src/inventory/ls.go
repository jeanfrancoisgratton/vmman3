// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/ls.go
// 2022-09-17 20:07:29

package inventory

import (
	"vmman3/db"
	"vmman3/helpers"
)

// 3 conditions :
// BAllHypervisors ? yes -> listhypervisors
func VM_Inventory() {
	var hyps []db.DbHypervisors
	if helpers.BAllHypervisors {
		hyps = listHypervisors()
	} else {
		if helpers.BsingleHypervisor {
			hyps = []db.DbHypervisors{{HID: 0, Hname: "localhost", Haddress: "127.0.0.1"}}
		} else {
			hyps = []db.DbHypervisors{{HID: 0, Hname: helpers.ConnectURI, Haddress: helpers.ConnectURI}}
		}
	}

	// First step: get the connection URI for a given hypervisor, and then connect
	for _, v := range hyps {
		helpers.ConnectURI = getURI(v.Haddress, v.Hconnectinguser)
		// to be uncommented soon
		//domains := GetVMlist()
	}

}
