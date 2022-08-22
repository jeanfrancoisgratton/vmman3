// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// helpers/connectionHelpers.go
// 2022-08-22 17:17:36

package helpers

import (
	"fmt"
	"libvirt.org/go/libvirt"
)

var ConnectURI string

// getConn() : ouvre la connexion à l'hyperviseur
func GetConn() libvirt.Connect {
	conn, err := libvirt.NewConnect(ConnectURI)

	if err != nil {
		fmt.Println("Error in inventory.getConn() : ", err)
	}

	return *conn
}

// getStateHelper() : transforme la variable DomainState (un int, en fait) en string
func GetStateHelper(state libvirt.DomainState) string {
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
	var conn = GetConn()

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)

	if err != nil {
		fmt.Println("Error in inventory.getVMlist() : ", err)
	}

	defer conn.Close()
	return doms
}
