// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/ls.go
// 2022-09-17 20:07:29

package inventory

import (
	"fmt"
	"os"
	"vmman3/db"
	"vmman3/helpers"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// LOTS of stuff to unpack in here.. FIXME

func VmInventory() {
	var hyps []db.DbHypervisors
	var allVMspecs []vmInfo

	// we need the hypervisors' list, except when BSingleHypervisor is false
	if helpers.BAllHypervisors {
		hyps = ListHypervisors()
	} else {
		// This means we already have a valid ConnectURI, either qemu://system, or a qemu+ssh:// one
		hyps = []db.DbHypervisors{{HID: 0, Hname: helpers.ConnectURI, Haddress: helpers.ConnectURI}}
	}

	// First step: get the connection URI for a given hypervisor, and then iterate+connect on them
	for _, hyp := range hyps {
		if hyp.Hname != hyp.Haddress {
			// this here means that we have to build the URI from the DB because BAllHypervisors == true
			helpers.ConnectURI = fmt.Sprintf("qemu+ssh://%s@%s/system", hyp.Hconnectinguser, hyp.Hname)
		}

		// Second step: collect the information
		if helpers.ConnectURI != "qemu:///system" {
			_, _, hyp.Hname = helpers.SplitConnectURI(helpers.ConnectURI)
		}
		vmspecs := collectInfo(hyp.Hname)
		allVMspecs = append(allVMspecs, vmspecs...)
	}

	// Third step: display information
	if helpers.BAllHypervisors {
		helpers.SurroundText("Registered domains on all hypervisors", false)
	} else {
		helpers.SurroundText(fmt.Sprintf("All domains on hypervisor %s", helpers.ConnectURI), false)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "VM name", "State", "vMem", "vCPUs", "Snaps", "Curr snap", "IP", "Last status change", "Hypervisor", "OS", "Storage"})

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
		t.AppendRow([]interface{}{sID, vmspec.viName, vmspec.viState, vmspec.viMem, vmspec.viCpu, vmspec.viSnapshot, vmspec.viCurrentSnapshot, vmspec.viIPaddress, vmspec.viLastStatusChange, vmspec.viHypervisor, vmspec.viOperatingSystem, vmspec.viStoragePool})

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
			//return text.Colors{text.BgBlack, text.FgHiGreen}
			return text.Colors{text.FgHiGreen}
		case "Crashed":
			//return text.Colors{text.BgBlack, text.FgHiRed}
			return text.Colors{text.FgHiRed}
		case "Blocked":
		case "Suspended":
		case "Paused":
			//return text.Colors{text.BgHiBlack, text.FgHiYellow}
			return text.Colors{text.FgHiYellow}
		}
		return nil
	})
	t.Render()
}
