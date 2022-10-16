// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/collectInfo.go
// 2022-09-17 14:01:35

package inventory

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
	"vmman3/snapshot_management"
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
		numsnap, _ = d.SnapshotNum(snapshotflags)
		if numsnap > 0 {
			i.viCurrentSnapshot = snapshot_management.GetCurrentSnapshotName(*conn, i.viName)
		} else {
			i.viCurrentSnapshot = "n/a"
		}

		i.viLastStatusChange, i.viOperatingSystem, i.viStoragePool = getInfoFromDB(i.viName, i.viHypervisor)

		vmspec = append(vmspec, vmInfo{viId: i.viId, viName: i.viName, viState: getStateHelper(dState), viMem: specs.Memory / 1024, viCpu: specs.NrVirtCpu,
			viSnapshot: uint(numsnap), viCurrentSnapshot: i.viCurrentSnapshot, viInterfaceName: i.viInterfaceName, viIPaddress: i.viIPaddress, viLastStatusChange: i.viLastStatusChange,
			viOperatingSystem: i.viOperatingSystem, viStoragePool: i.viStoragePool, viHypervisor: i.viHypervisor})
	}
	return vmspec
}

func getInterfaceSpecs(dom libvirt.Domain, vmname string) (string, string) {
	var domainInterface []libvirt.DomainInterface
	var interfaceName, interfaceAddress string
	var err error

	domainInterface, err = dom.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT)
	if err != nil {
		fmt.Println("Oooops: ", err)
	}
	for _, di := range domainInterface {
		if len(di.Name) > 2 && (di.Name[:3] == "enp" || di.Name[:3] == "eth") {
			interfaceName = di.Name
			domainIPaddresses := di.Addrs
			for _, dipa := range domainIPaddresses {
				if dipa.Type == libvirt.IP_ADDR_TYPE_IPV4 {
					interfaceAddress = dipa.Addr
				}
			}

		}
	}
	return interfaceName, interfaceAddress
}