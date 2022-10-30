// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/vm_management/vmRename.go
// 2022-10-29 19:04:58

package vm_management

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
)

func Rename(args []string) {
	var snapshotflags libvirt.DomainSnapshotListFlags

	if len(args) < 2 {
		fmt.Println("You need to provide a new name for the VM. Aborting")
		os.Exit(0)
	}
	oldName := args[0]
	newName := args[1]
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
	vm, _ := conn.LookupDomainByName(oldName)
	numsnap, _ := vm.SnapshotNum(snapshotflags)
	if numsnap > 0 {
		fmt.Println("You cannot rename " + oldName + " as this VM holds snapshots. The snapshots need to be removed, first.")
		os.Exit(1)
	}
	helpers.Wait4Shutdown(vm, oldName)
	vm.Rename(newName, 0)
	fmt.Println("%s --> %s", oldName, newName)
}
