// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// snapshotmanagement/snapshot.go
// 2022-08-16 17:51:19

package snapshotmanagement

import (
	"libvirt.org/go/libvirt"
)

func GetCurrentSnapshotName(conn *libvirt.Connect, vmname string) string {
	domain, _ := conn.LookupDomainByName(vmname)
	defer domain.Free()
	var currentSnapshot = "none"
	var snapshots, _ = domain.ListAllSnapshots(0)

	for _, snapshot := range snapshots {
		var isCurrent, _ = snapshot.IsCurrent(0)
		if isCurrent {
			currentSnapshot, _ = snapshot.GetName()
			break
		}
	}
	return currentSnapshot
}
