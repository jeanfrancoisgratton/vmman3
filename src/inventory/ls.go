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

// 3 conditions :
// BAllHypervisors ? yes -> listhypervisors
func VM_Inventory() {
	var hyps []db.DbHypervisors
	var allVMspecs []vmInfo
	if helpers.BAllHypervisors {
		hyps = listHypervisors()
	} else {
		if helpers.BsingleHypervisor {
			hyps = []db.DbHypervisors{{HID: 0, Hname: "localhost", Haddress: "127.0.0.1"}}
		} else {
			hyps = []db.DbHypervisors{{HID: 0, Hname: helpers.ConnectURI, Haddress: helpers.ConnectURI}}
		}
	}

	// First step: get the connection URI for a given hypervisor, and then connect
	for _, v := range hyps {
		helpers.ConnectURI = getURI(v.Haddress, v.Hconnectinguser)
		// to be uncommented soon
		vmspecs := collectInfo()
		allVMspecs = append(allVMspecs, vmspecs...)
	}
	helpers.SurroundText("All domains on hypervisor "+helpers.ConnectURI, false)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader((table.Row{"ID", "VM name", "State", "vMemory", "vCPUs", "Snapshots", "Curr snapshot", "iface name", "IP address", "Last status change", "Hypervisor"}))

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
		t.AppendRow([]interface{}{sID, vmspec.viName, vmspec.viState, vmspec.viMem, vmspec.viCpu, vmspec.viSnapshot, vmspec.viCurrentSnapshot, vmspec.viInterfaceName, vmspec.viIPaddress, "", ""})

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
