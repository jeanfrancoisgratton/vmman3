// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/interfaceInfo.go
// 2022-09-17 14:00:19

package inventory

import (
	"fmt"
	"libvirt.org/go/libvirt"
)

// getInterfaceSpecs() : va chercher le nom de l'interface réseau principale, et son adresse IP
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
