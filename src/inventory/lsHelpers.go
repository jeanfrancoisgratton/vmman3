// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/lsHelpers.go
// 2022-09-17 21:26:46

package inventory

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"strconv"
	"time"
	"vmman3/helpers"
)

// getStateHelper() : gets DomainState value and transforms it in a string
func getStateHelper(state libvirt.DomainState) string {
	ds := ""
	switch state {
	case libvirt.DOMAIN_NOSTATE:
		ds = "no state"
	case libvirt.DOMAIN_RUNNING:
		ds = "Running"
	case libvirt.DOMAIN_BLOCKED:
		ds = "Blocked"
	case libvirt.DOMAIN_CRASHED:
		ds = "Crashed"
	case libvirt.DOMAIN_SHUTDOWN:
	case libvirt.DOMAIN_SHUTOFF:
		ds = "Shutdown"
	case libvirt.DOMAIN_PMSUSPENDED:
		ds = "Suspended"
	case libvirt.DOMAIN_PAUSED:
		ds = "Paused"
	default:
		ds = "n/a"
	}
	return ds
}

// getVMList() : Ammasse la liste des VMs sur cet hyperviseur
func GetVMlist() []libvirt.Domain {
	conn, err := libvirt.NewConnect(helpers.ConnectURI)
	if err != nil {
		lverr, ok := err.(libvirt.Error)
		if ok && (lverr.Message == "End of file while reading data: virt-ssh-helper: cannot connect to '/var/run/libvirt/libvirt-sock': Failed to connect socket to '/var/run/libvirt/libvirt-sock': Connection refused: Input/output error") ||
			lverr.Message == "internal error: unexpected qemu URI path '/system/', try qemu:///system" {
			fmt.Printf("Hypervisor %s is offline\n", helpers.ConnectURI)
			return nil
		} else {
			panic(err)
		}
	}

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)

	if err != nil {
		fmt.Println("Error in inventory.GetVMlist() : ", err)
	}

	defer conn.Close()
	return doms
}

// Some of it might not be needed anymore...
// getInterfaceSpecs(): this is where we get the interface name and its IP
func getInterfaceSpecs(dom libvirt.Domain, vmname string) (string, string) {
	var domainInterface []libvirt.DomainInterface
	var interfaceName, interfaceAddress string
	var err error

	domainInterface, err = dom.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT)
	if err != nil {
		fmt.Printf("\nOooops: %s\n\n", err)
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

// getUptime() : gets the active VM's uptime from the database
func getUptime(lastState string) string {
	//stateInDB := time.Unix(lastState, 0)
	a, _ := strconv.ParseInt(lastState, 10, 64)
	//deltaUnix := time.Now().Unix() - lastState
	return time.Unix(time.Now().Unix()-a, 0).Format("2006.01.02 15:04:05")
}
