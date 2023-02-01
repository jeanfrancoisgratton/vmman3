// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// vmmanagement/vmStop.go
// 2022-08-22 13:13:14

package vmmanagement

import (
	"fmt"
	"os"
	"vmman3/helpers"
	"vmman3/inventory"
)

func Start(args []string) {
	var bIsActive bool

	conn := helpers.Connect2HVM()
	defer conn.Close()

	for _, vmname := range args {
		var host string
		domain := helpers.GetDomain(conn, vmname)
		if domain == nil {
			os.Exit(0)
		}
		defer domain.Free()

		bIsActive, _ = domain.IsActive()
		if bIsActive {
			fmt.Printf("Domain %s on %s is already up\n", vmname, helpers.ConnectURI)
		} else {
			err := domain.Create()
			fmt.Printf("Domain %s is starting... ", vmname)
			if err != nil {
				fmt.Printf("\nERROR :\n")
				fmt.Println(err)
			} else {
				fmt.Printf("done\n")
				// This is where we update the vmstates table
				if helpers.ConnectURI == "qemu:///system" {
					host, _ = os.Hostname()
				} else {
					_, _, host = helpers.SplitConnectURI(helpers.ConnectURI)
				}
				vmStateChange(host, vmname)
			}
		}
	}
}

func StartAll() {
	var vmlist []string
	domains := inventory.GetVMlist()

	for _, domain := range domains {
		var _, err = domain.GetID()
		if err != nil { // this means GetID() did not return an ID, thus candidate to be started
			vmname, _ := domain.GetName()
			vmlist = append(vmlist, vmname)
		}
	}
	Start(vmlist)
}
