// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// snapshot_management/snapshot.go
// 2022-08-16 17:51:19

package snapshot_management

import (
	"fmt"
	"libvirt.org/go/libvirt"
)

func GetCurrentSnapshot(conn libvirt.Connect, vmname string) string {
	var vm, _ = conn.LookupDomainByName(vmname)
	var snapshots, _ = vm.ListAllSnapshots(0)
	//var snapshot

	for snapshot := range snapshots {
		fmt.Printf("%T\n", snapshot)
		fmt.Println(snapshot)
	}

	return "" // not returning anything, yet
}
