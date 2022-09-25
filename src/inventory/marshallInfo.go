// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/marshallInfo.go
// 2022-09-17 14:01:35

package inventory

import (
	"libvirt.org/go/libvirt"
	"vmman3/snapshot_management"
)

// FIXME: needs cleanup and/or readability fixes
// collectInfo(hypervisorname string) : Inventaire détaillé des VMs
func collectInfo(hypervisorname string) []vmInfo {
	var snapshotflags libvirt.DomainSnapshotListFlags
	var numsnap int
	var vmspec []vmInfo
	var i vmInfo
	var dState libvirt.DomainState
	doms := GetVMlist()
	conn := GetConn()

	for _, dom := range doms {
		specs, err := dom.GetInfo()
		i.viId, err = dom.GetID()
		if err != nil {
			i.viId = 0
		}
		// NOTE: the following struct member is near-useless. Might be removed in future versions
		i.viHypervisor = hypervisorname
		// VM NAME
		i.viName, _ = dom.GetName()
		// VM STATE
		dState, _, _ = dom.GetState()
		i.viState = getStateHelper(dState)
		if i.viState == "Running" {
			// INTERFACE INFO
			i.viInterfaceName, i.viIPaddress = getInterfaceSpecs(dom, i.viName)
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

		i.viLastStatusChange, i.viOperatingSystem, i.viStoragePool = getInfoFromDB(i.viName, hypervisorname)

		vmspec = append(vmspec, vmInfo{viId: i.viId, viName: i.viName, viState: getStateHelper(dState), viMem: specs.Memory / 1024, viCpu: specs.NrVirtCpu,
			viSnapshot: uint(numsnap), viCurrentSnapshot: i.viCurrentSnapshot, viInterfaceName: i.viInterfaceName, viIPaddress: i.viIPaddress, viLastStatusChange: i.viLastStatusChange,
			viOperatingSystem: i.viOperatingSystem, viStoragePool: i.viStoragePool, viHypervisor: i.viHypervisor})
	}

	return vmspec
}
