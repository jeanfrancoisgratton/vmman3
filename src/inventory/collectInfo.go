// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/collectInfo.go
// 2022-09-17 14:01:35

package inventory

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
	"vmman3/snapshotmanagement"
)

// TODO: needs cleanup and/or readability fixes
// collectInfo(hypervisorname string) : Inventaire détaillé des VMs
func collectInfo(hypervisorname string) []vmInfo {
	var snapshotflags libvirt.DomainSnapshotListFlags
	var numsnap int
	var vmspec []vmInfo
	var i vmInfo
	var dState libvirt.DomainState
	doms := GetVMlist()

	if doms == nil {
		return nil
	}

	conn, err := libvirt.NewConnect(helpers.ConnectURI)
	if err != nil {
		lverr, ok := err.(libvirt.Error)
		if ok && lverr.Message == "End of file while reading data: virt-ssh-helper: cannot connect to '/var/run/libvirt/libvirt-sock': Failed to connect socket to '/var/run/libvirt/libvirt-sock': Connection refused: Input/output error" {
			fmt.Printf("Hypervisor %s is offline\n", helpers.ConnectURI)
			return nil
		} else {
			panic(err)
		}
	}

	for _, dom := range doms {
		specs, err := dom.GetInfo()
		i.viId, err = dom.GetID()
		if err != nil {
			i.viId = 0
		}
		// NOTE: the following struct member is near-useless. Might be removed in future versions
		if hypervisorname == "qemu:///system" {
			i.viHypervisor, _ = os.Hostname()
		} else {
			i.viHypervisor = hypervisorname
		}

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
		defer d.Free()
		numsnap, _ = d.SnapshotNum(snapshotflags)
		if numsnap > 0 {
			i.viCurrentSnapshot = snapshotmanagement.GetCurrentSnapshotName(*conn, i.viName)
		} else {
			i.viCurrentSnapshot = "n/a"
		}

		i.viLastStatusChange, i.viOperatingSystem, i.viStoragePool = getInfoFromDB(i.viName, i.viHypervisor)

		// uptime is not yet working, so commenting out that block
		//if i.viId > 0 {
		//	// time.Unix(time.Now().Unix() - lastState, 0).Format("2006.01.02 15:04:05")
		//	i.viLastStatusChange = getUptime(i.viLastStatusChange)
		//}

		vmspec = append(vmspec, vmInfo{viId: i.viId, viName: i.viName, viState: getStateHelper(dState), viMem: specs.Memory / 1024, viCpu: specs.NrVirtCpu,
			viSnapshot: uint(numsnap), viCurrentSnapshot: i.viCurrentSnapshot, viInterfaceName: i.viInterfaceName, viIPaddress: i.viIPaddress, viLastStatusChange: i.viLastStatusChange,
			viOperatingSystem: i.viOperatingSystem, viStoragePool: i.viStoragePool, viHypervisor: i.viHypervisor})
	}
	return vmspec
}
