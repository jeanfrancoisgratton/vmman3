// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// inventory/inventory-cmd-ls.go
// 2022-08-16 17:51:19

package inventory

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
	"vmman3/snapshot_management"
)

type vmInfo struct {
	viId                                            uint
	viName, viState                                 string
	viMem                                           uint64
	viCpu, viSnapshot                               uint
	viCurrentSnapshot, viInterfaceName, viIPaddress string
}

// getInterfaceSpecs() : va chercher le nom de l'interface réseau principale, et son adresse IP
func getInterfaceSpecs(dom libvirt.Domain, vmname string) (string, string) {
	var domainInterface []libvirt.DomainInterface
	var interfaceName, interfaceAddress string
	var err error

	domainInterface, err = dom.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT)
	if err != nil {
		fmt.Println("oooops.....")
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

// FIXME: needs cleanup and/or readability fixes
// collecteInfo() : Inventaire détaillé des VMs
func collectInfo() []vmInfo {
	var snapshotflags libvirt.DomainSnapshotListFlags
	var numsnap int
	vmspec := []vmInfo{}
	var i vmInfo
	var dState libvirt.DomainState
	doms := helpers.GetVMlist()
	var conn = helpers.GetConn()

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
		i.viState = helpers.GetStateHelper(dState)
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

		vmspec = append(vmspec, vmInfo{viId: i.viId, viName: i.viName, viState: helpers.GetStateHelper(dState), viMem: specs.Memory / 1024, viCpu: specs.NrVirtCpu,
			viSnapshot: uint(numsnap), viCurrentSnapshot: i.viCurrentSnapshot, viInterfaceName: i.viInterfaceName, viIPaddress: i.viIPaddress})
	}

	return vmspec
}

func VM_List() {
	var vmspecs = collectInfo()

	helpers.SurroundText("All domains on hypervisor "+helpers.ConnectURI, false)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if helpers.BsingleHypervisor {
		t.AppendHeader((table.Row{"ID", "VM name", "State", "vMemory", "vCPUs", "Snapshots", "Curr snapshot", "iface name", "IP address"}))
	} else {
		t.AppendHeader((table.Row{"ID", "VM name", "State", "vMemory", "vCPUs", "Snapshots", "Curr snapshot", "iface name", "IP address", "Last status change", "Hypervisor"}))
	}

	for _, vmspec := range vmspecs {
		sID := ""
		if vmspec.viId > 0 && vmspec.viId < 10 {
			sID = fmt.Sprintf("000%d", vmspec.viId)
		}
		if vmspec.viId > 9 && vmspec.viId < 100 {
			sID = fmt.Sprintf("00%d", vmspec.viId)
		}
		if vmspec.viId > 99 && vmspec.viId < 999 {
			sID = fmt.Sprintf("0%d", vmspec.viId)
		}
		if helpers.BsingleHypervisor {
			t.AppendRow([]interface{}{sID, vmspec.viName, vmspec.viState, vmspec.viMem, vmspec.viCpu, vmspec.viSnapshot, vmspec.viCurrentSnapshot, vmspec.viInterfaceName, vmspec.viIPaddress})
		} else {
			t.AppendRow([]interface{}{sID, vmspec.viName, vmspec.viState, vmspec.viMem, vmspec.viCpu, vmspec.viSnapshot, vmspec.viCurrentSnapshot, vmspec.viInterfaceName, vmspec.viIPaddress, "", ""})
		}

	}
	t.SortBy([]table.SortBy{
		{Name: "ID", Mode: table.Asc},
		{Name: "VM name", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)
	//t.Style().Options.DrawBorder = false
	//t.Style().Options.SeparateColumns = false
	t.Style().Format.Header = text.FormatDefault
	t.SetRowPainter(table.RowPainter(func(row table.Row) text.Colors {
		switch row[2] {
		case "Running":
			return text.Colors{text.BgBlack, text.FgHiGreen}
		case "Crashed":
			return text.Colors{text.BgBlack, text.FgHiRed}
		case "Blocked":
		case "Suspended":
		case "Paused":
			return text.Colors{text.BgHiBlack, text.FgHiYellow}
		}
		return nil
	}))
	t.Render()
}