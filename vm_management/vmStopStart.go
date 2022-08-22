// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// vm_management/vmStopStart.go
// 2022-08-22 13:13:14

package vm_management

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"vmman3/helpers"
)

func Stop(args []string) {
	conn := helpers.GetConn()
	var bIsActive bool

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
	domains := helpers.GetVMlist()

	for _, domain := range domains {
		var _, err = domain.GetID()
		if err == nil {
			vmname, _ := domain.GetName()
			vmlist = append(vmlist, vmname)
		}
	}
	Stop(vmlist)
}
