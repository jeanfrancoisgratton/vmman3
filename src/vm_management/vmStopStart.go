// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// vm_management/vmStopStart.go
// 2022-08-22 13:13:14

package vm_management

import (
	"fmt"
	"libvirt.org/go/libvirt"
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
		domain, _ := conn.LookupDomainByName(vmname)

		bIsActive, _ = domain.IsActive()
		if !bIsActive {
			fmt.Printf("Domain %s on %s is already shut down\n", vmname, helpers.ConnectURI)
		} else {
			err := domain.DestroyFlags(libvirt.DOMAIN_DESTROY_GRACEFUL)
			fmt.Printf("Domain %s is being shut down ...", vmname)
			if err != nil {
				fmt.Printf("\nERROR :\n")
				fmt.Println(err)
			} else {
				fmt.Printf("done\n")
			}
		}
	}
}

func StopAll() {
	var vmlist []string
	domains := inventory.GetVMlist()

	for _, domain := range domains {
		var _, err = domain.GetID()
		if err == nil {
			vmname, _ := domain.GetName()
			vmlist = append(vmlist, vmname)
		}
	}
	Stop(vmlist)
}
