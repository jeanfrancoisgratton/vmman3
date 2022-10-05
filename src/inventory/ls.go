// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/ls.go
// 2022-09-17 20:07:29

package inventory

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"vmman3/db"
	"vmman3/helpers"
)

// LOTS of stuff to unpack in here.. FIXME

func VmInventory() {
	var hyps []db.DbHypervisors
	var allVMspecs []vmInfo

	// we need the hypervisors' list, except when BSingleHypervisor is false
	if !helpers.BSingleHypervisor {
		hyps = ListHypervisors()
	} else {
		// BSingleHypervisor is true
		host, _ := os.Hostname()
		hyps = []db.DbHypervisors{{HID: 0, Hname: host, Haddress: "127.0.0.1", Hconnectinguser: ""}}
	}

	// FIXME: looks unefficient....
	// if both -1 and -a are false, we need to fetch the full uri string from the db
	// we basically loop through all hypervisors in DB to fetch the one corresponding to the -c flag
	for _, name := range hyps {
		if name.Hname == helpers.ConnectURI {
			helpers.ConnectURI = fmt.Sprintf("qemu+ssh;//%s@%s/system/", name.Hconnectinguser, name.Haddress)
			break
		}
	}

	// First step: get the connection URI for a given hypervisor, and then iterate+connect on them
	for _, v := range hyps {
		if helpers.BSingleHypervisor {
			helpers.ConnectURI = "qemu:///system"
		} else {
			if helpers.ConnectURI == v.Hname && !helpers.BAllHypervisors {
				helpers.ConnectURI = fmt.Sprintf("qemu+ssh;//%s@%s/system/", v.Hconnectinguser, v.Haddress)
			}
			helpers.ConnectURI = fmt.Sprintf("qemu+ssh://%s@%s/system/", v.Hconnectinguser, v.Haddress)
		}

		// Second step: connect to hypervisor

		// Third step: collect the information
		vmspecs := collectInfo(v.Hname)
		allVMspecs = append(allVMspecs, vmspecs...)
	}

	// Fourth step: display information
	if helpers.BAllHypervisors {
		fmt.Println("Registered domains on all hypervisors")
	} else {
		fmt.Println("All domains on hypervisor ", helpers.ConnectURI)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "VM name", "State", "vMemory", "vCPUs", "Snapshots", "Curr snapshot", "iface name", "IP address", "Last status change", "Hypervisor"})

	for _, vmspec := range allVMspecs {
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
		t.AppendRow([]interface{}{sID, vmspec.viName, vmspec.viState, vmspec.viMem, vmspec.viCpu, vmspec.viSnapshot, vmspec.viCurrentSnapshot, vmspec.viInterfaceName, vmspec.viIPaddress, vmspec.viLastStatusChange, vmspec.viHypervisor, ""})

	}
	t.SortBy([]table.SortBy{
		{Name: "ID", Mode: table.Asc},
		{Name: "VM name", Mode: table.Asc},
		{Name: "Hypervisor", Mode: table.Asc},
	})
	t.SetStyle(table.StyleBold)
	//t.Style().Options.DrawBorder = false
	//t.Style().Options.SeparateColumns = false
	t.Style().Format.Header = text.FormatDefault
	t.SetRowPainter(func(row table.Row) text.Colors {
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
	})
	t.Render()
}
