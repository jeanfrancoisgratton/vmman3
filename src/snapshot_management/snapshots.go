// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// snapshot_management/snapshot.go
// 2022-08-16 17:51:19

package snapshot_management

import (
	"libvirt.org/go/libvirt"
)

func GetCurrentSnapshotName(conn libvirt.Connect, vmname string) string {
	var vm, _ = conn.LookupDomainByName(vmname)
	var currentSnapshot = "none"
	var snapshots, _ = vm.ListAllSnapshots(0)

	for _, snapshot := range snapshots {
		var isCurrent, _ = snapshot.IsCurrent(0)
		if isCurrent {
			currentSnapshot, _ = snapshot.GetName()
			break
		}
	}
	return currentSnapshot
}
