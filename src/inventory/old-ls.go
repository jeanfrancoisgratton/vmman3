// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// inventory/inventory-cmd-old-ls.go
// 2022-08-16 17:51:19

package inventory

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"vmman3/helpers"
)

func VM_List2() {
	vmspecs := collectInfo()

	helpers.SurroundText("All domains on hypervisor "+helpers.ConnectURI, false)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader((table.Row{"ID", "VM name", "State", "vMemory", "vCPUs", "Snapshots", "Curr snapshot", "iface name", "IP address", "Last status change", "Hypervisor"}))

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
