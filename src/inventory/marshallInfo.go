// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/marshallInfo.go
// 2022-09-17 14:01:35

package inventory

import (
	"libvirt.org/go/libvirt"
	"vmman3/snapshot_management"
)

// FIXME: needs cleanup and/or readability fixes
// collectInfo() : Inventaire détaillé des VMs
func collectInfo() []vmInfo {
	var snapshotflags libvirt.DomainSnapshotListFlags
	var numsnap int
	vmspec := []vmInfo{}
	var i vmInfo
	var dState libvirt.DomainState
	doms := GetVMlist()
	var conn = GetConn()

	for _, dom := range doms {
		var specs, err = dom.GetInfo()
		i.viId, err = dom.GetID()
		if err != nil {
			i.viId = 0
		}
		// VM NAME
		i.viName, _ = dom.GetName()
		// VM STATE
		dState, _, _ = dom.GetState()
		i.viState = getStateHelper(dState)
		if i.viState == "Running" {
			// INTERFACE INFO
			i.viInterfaceName, i.viIPaddress = getInterfaceSpecs(dom, i.viName)
			// CURRENT HYPERVISOR
		} else {
			i.viInterfaceName = ""
			i.viIPaddress = ""
		}
		// SNAPSHOT INFO
		d, _ := conn.LookupDomainByName(i.viName)
		numsnap, _ = d.SnapshotNum(snapshotflags)
		if numsnap > 0 {
			i.viCurrentSnapshot = snapshot_management.GetCurrentSnapshotName(conn, i.viName)
		} else {
			i.viCurrentSnapshot = "n/a"
		}

		//lastStatusChange := getStatusChangeTS(i.viName)
		//get

		vmspec = append(vmspec, vmInfo{viId: i.viId, viName: i.viName, viState: getStateHelper(dState), viMem: specs.Memory / 1024, viCpu: specs.NrVirtCpu,
			viSnapshot: uint(numsnap), viCurrentSnapshot: i.viCurrentSnapshot, viInterfaceName: i.viInterfaceName, viIPaddress: i.viIPaddress})
	}

	return vmspec
}
