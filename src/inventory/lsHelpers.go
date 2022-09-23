// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/lsHelpers.go
// 2022-09-17 21:26:46

package inventory

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"vmman3/helpers"
)

// getConn() : ouvre la connexion à l'hyperviseur
func GetConn() libvirt.Connect {
	conn, err := libvirt.NewConnect(helpers.ConnectURI)

	if err != nil {
		fmt.Println("Error in inventory.getConn() : ", err)
	}

	return *conn
}

// getStateHelper() : transforme la variable DomainState (un int, en fait) en string
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
	conn := GetConn()

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)

	if err != nil {
		fmt.Println("Error in inventory.GetVMlist() : ", err)
	}

	defer conn.Close()
	return doms
}
