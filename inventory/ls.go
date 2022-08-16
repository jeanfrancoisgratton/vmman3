// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// helpers/ls.go
// 2022-08-16 17:51:19

package inventory

import (
	"fmt"
	_ "github.com/jedib0t/go-pretty/v6/table"
	"libvirt.org/libvirt-go"
	"vmman3/helpers"
)

type vmInfo struct {
	viId                           int
	viName, viState                string
	viMem, viCpu, viSnapshot       int
	viCurrentSnapshot, viIPaddress string
}

// getVMList() : retourne le nom de toutes les VMs sur l'hyperviseur
func getVMlist() []string {
	var domaines []string
	//	var vmSpec []vmInfo

	conn, err := libvirt.NewConnect(helpers.ConnectURI)
	if err != nil {
		print("whatever in inventory.getVMList()")
	}
	defer conn.Close()

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)

	if err != nil {
		print("whatever in inventory.getVMList()")
	}
	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			domaines = append(domaines, name)
		}
		dom.Free()
	}
	return domaines
}

// VM_List() : Inventaire détaillé des VMs
// 1. On va chercher les noms des VMs (getVMlist())
// 2. Pour chaque VM de la liste, on cherche :
// 		statut, connectionID (si online), #memore, #cpu, #snapshots, le nom du snapshot courrant, adresse IP
// 3. Pour chaque stat, on popule la structure vmInfo
// 4. On tabularise et on affice
func VM_List() {
	domaines := getVMlist()

	for _, dom := range domaines {
		fmt.Printf("%T %s\n", dom, dom)
	}

}
