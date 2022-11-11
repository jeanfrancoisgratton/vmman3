// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// snapshotmanagement/snapshot.go
// 2022-08-16 17:51:19

package snapshotmanagement

import (
	"encoding/xml"
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
)

// GetCurrentSnapshotName() : Gets the name of the current snapshot for a given VM
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

// ListSnapshots() : Lists all snapshots on current VM
func ListSnapshots(vmname string) {
	var snapXMLdata SnapshotXMLstruct
	var snaps []SnapshotXMLstruct
	conn := helpers.Connect2HVM()
	defer conn.Close()

	domain, _ := conn.LookupDomainByName(vmname)
	defer domain.Free()
	numsnap, _ := domain.SnapshotNum(0)
	if numsnap == 0 {
		fmt.Printf("Domain %s has no snapshot\n", vmname)
		os.Exit(0)
	}
	snapshots, _ := domain.ListAllSnapshots(0)

	for _, snap := range snapshots {
		data, _ := snap.GetXMLDesc(0)
		if err := xml.Unmarshal([]byte(data), &snapXMLdata); err != nil {
			fmt.Println("err:", err)
			os.Exit(-1)
		} else {

			snapXMLdata.CurrentSnapshot, _ = snap.IsCurrent(0)
			snaps = append(snaps, snapXMLdata)
		}
	}
	displaySnapshots(snaps, vmname)
}

// TODO: add capability to delete children snapshots (hint: VIR_DOMAIN_SNAPSHOT_DELETE_CHILDREN)
// RemoveSnapshots() : Removes snapshot of a given VM
func RemoveSnapshot(vmname string, snapshotname string) {
	var snap *libvirt.DomainSnapshot
	conn := helpers.Connect2HVM()
	defer conn.Close()

	domain, _ := conn.LookupDomainByName(vmname)
	defer domain.Free()
	numsnap, _ := domain.SnapshotNum(0)
	if numsnap == 0 {
		fmt.Printf("Domain %s has no snapshot\n", vmname)
		os.Exit(0)
	}

	helpers.Wait4Shutdown(domain, vmname)
	if snapshotname == "" {
		snapshotname = GetCurrentSnapshotName(conn, vmname)
	}
	err := snap.Delete(0)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
}

// CreateSnapshot() : create a snaspshot on the given VM
func CreateSnapshot(vmname string, snapname string) {}
