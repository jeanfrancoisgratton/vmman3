// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// vmmanagement/vmStop.go
// 2022-08-22 13:13:14

package vm_management

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/db"
	"vmman3/helpers"
	"vmman3/inventory"
)

func Stop(args []string) {
	var bIsActive bool
	conn := helpers.Connect2HVM()
	defer conn.Close()

	for _, vmname := range args {
		var host string
		domain, _ := conn.LookupDomainByName(vmname)
		defer domain.Free()

		bIsActive, _ = domain.IsActive()
		if !bIsActive {
			fmt.Printf("Domain %s on %s is already shut down\n", vmname, helpers.ConnectURI)
		} else {
			err := domain.ShutdownFlags(libvirt.DOMAIN_SHUTDOWN_DEFAULT)
			fmt.Printf("Domain %s is being shut down... ", vmname)
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

func StopAll() {
	var hyps []db.DbHypervisors
	var vmlist []string

	// we need the hypervisors' list, except when BSingleHypervisor is false
	if helpers.BAllHypervisors {
		hyps = inventory.ListHypervisors()
	} else {
		// This means we already have a valid ConnectURI, either qemu://system, or a qemu+ssh:// one
		hyps = []db.DbHypervisors{{HID: 0, Hname: helpers.ConnectURI, Haddress: helpers.ConnectURI}}
	}

	// First step: get the connection URI for a given hypervisor, and then iterate+connect on them
	for _, hyp := range hyps {
		if hyp.Hname != hyp.Haddress {
			// this here means that we have to build the URI from the DB because BAllHypervisors == true
			helpers.ConnectURI = fmt.Sprintf("qemu+ssh://%s@%s/system", hyp.Hconnectinguser, hyp.Hname)
		}
		domains := inventory.GetVMlist()
		for _, domain := range domains {
			var _, err = domain.GetID()
			if err == nil { // this means GetID() returned an ID, thus the VM is not shutdown (could be paused)
				vmname, _ := domain.GetName()
				vmlist = append(vmlist, vmname)
			}
		}
		Stop(vmlist)
		vmlist = nil
	}
}
