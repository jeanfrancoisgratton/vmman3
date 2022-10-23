// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/vm_management/vmRemove.go
// 2022-10-22 12:42:35

package vm_management

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
)

// This will remove the VM, and optionally leave its storage there
func Remove(args []string) {
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
	defer conn.Close()

	for _, vmname := range args {
		//var host string
		domain, _ := conn.LookupDomainByName(vmname)

		helpers.Wait4Shutdown(domain, vmname)
		fmt.Println(vmname + " now shutdown. Proceeding to removal from inventory.")
		err = domain.UndefineFlags(libvirt.DOMAIN_UNDEFINE_SNAPSHOTS_METADATA)
		if err != nil {
			lverr, ok := err.(libvirt.Error)
			if ok {
				fmt.Println(lverr.Message)
				os.Exit(-1)
			}
		}
		if !helpers.BkeepStorage {
			storage := helpers.GetStorage4VM(vmname)
			removeStorage(storage)
		}
	}
}

func removeStorage(storage []helpers.VmStorageDetails) {
	for _, storespecs := range storage {
		os.Remove()
	}
}
