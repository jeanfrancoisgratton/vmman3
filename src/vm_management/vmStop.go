// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// vmmanagement/vmStop.go
// 2022-08-22 13:13:14

package vm_management

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
	"vmman3/inventory"
)

func Stop(args []string) {
	var bIsActive bool
	conn, err := libvirt.NewConnect(helpers.ConnectURI)

	if err != nil {
		lverr, ok := err.(libvirt.Error)
		if ok && lverr.Message == "End of file while reading data: virt-ssh-helper: cannot connect to '/var/run/libvirt/libvirt-sock': Failed to connect socket to '/var/run/libvirt/libvirt-sock': Connection refused: Input/output error" {
			fmt.Printf("Hypervisor %s is offline\n", helpers.ConnectURI)
			return
		} else {
			panic(err)
		}
	}

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
	var vmlist []string
	domains := inventory.GetVMlist()

	for _, domain := range domains {
		var _, err = domain.GetID()
		if err == nil { // this means GetID() returned an ID, thus the VM is not shutdown (could be paused)
			vmname, _ := domain.GetName()
			vmlist = append(vmlist, vmname)
		}
	}
	Stop(vmlist)
}
