// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// inventory/ls.go
// 2022-08-16 17:51:19

package inventory

import (
	"fmt"
	_ "github.com/jedib0t/go-pretty/v6/table"
	"libvirt.org/go/libvirt"
	"vmman3/helpers"
	"vmman3/snapshot_management"
	_ "vmman3/snapshot_management"
)

type vmInfo struct {
	viId                           uint
	viName, viState                string
	viMem                          uint64
	viCpu, viSnapshot              uint
	viCurrentSnapshot, viIPaddress string
}

// getConn() : ouvre la connexion à l'hyperviseur
func getConn() libvirt.Connect {
	conn, err := libvirt.NewConnect(helpers.ConnectURI)

	if err != nil {
		fmt.Println("Error in inventory.getConn() : ", err)
	}

	return *conn
}

// getCurrentSnapshotName() : trouve le nom du snapshot actuel de la VM
func getCurrentSnapshotName(vmname string) string {
	return ""
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
func getVMlist() []libvirt.Domain {
	var conn = getConn()

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)

	if err != nil {
		fmt.Println("Error in inventory.getVMlist() : ", err)
	}

	defer conn.Close()
	return doms
}

// FIXME: needs cleanup
// VM_List() : Inventaire détaillé des VMs
//  1. On va chercher les noms des VMs (getVMlist())
//  2. Pour chaque VM de la liste, on cherche :
//     statut, connectionID (si online), #memore, #cpu, #snapshots, le nom du snapshot courrant, adresse IP
//  3. Pour chaque stat, on popule la structure vmInfo
//  4. On tabularise et on affice
func VM_List() {
	//var err error
	var snapshotflags libvirt.DomainSnapshotListFlags
	var numsnap int
	vmspec := []vmInfo{}
	var i vmInfo
	var dState libvirt.DomainState
	doms := getVMlist()
	var conn = getConn()

	for _, dom := range doms {
		var specs, err = dom.GetInfo()
		i.viId, err = dom.GetID()
		if err != nil {
			i.viId = 0
		}
		i.viName, _ = dom.GetName()
		dState, _, _ = dom.GetState()
		d, _ := conn.LookupDomainByName(i.viName)
		numsnap, _ = d.SnapshotNum(snapshotflags)
		if numsnap > 0 {
			//i.viCurrentSnapshot = snapshot_management.GetCurrentSnapshot(conn, i.viName)
			//[]libvirt.DomainSnapshot := snapshot_management.GetCurrentSnapshot(conn, i.viName)
			fmt.Println(snapshot_management.GetCurrentSnapshot(conn, i.viName))
		} else {
			i.viCurrentSnapshot = "n/a"
		}
		vmspec = append(vmspec, vmInfo{viId: i.viId, viName: i.viName, viState: getStateHelper(dState), viMem: specs.Memory / 1024, viCpu: specs.NrVirtCpu,
			viSnapshot: uint(numsnap)})
	}
	var y int
	{
		for y, i = range vmspec {
			fmt.Printf("%d: ID=%d Name=%s State=%s Mem=%d CPU=%d #snapshot=%d\n", y, i.viId, i.viName, i.viState, i.viMem, i.viCpu, i.viSnapshot)
		}
	}
}
