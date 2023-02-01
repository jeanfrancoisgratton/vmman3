// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/snapshotmanagement/snapHelpers.go
// 2022-11-09 19:15:22

package snapshotmanagement

import (
	"encoding/xml"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"time"
	"vmman3/helpers"
)

// Snapshot XML definitions
type ParentElement struct {
	XMLName    xml.Name `xml:"parent"`
	ParentName string   `xml:"name"`
}
type SnapshotXMLstruct struct {
	SnapshotName    string        `xml:"name"`
	CreationTime    int64         `xml:"creationTime"`
	Parent          ParentElement `xml:"parent"`
	CurrentSnapshot bool
}

// displaySnapshots() : will display the actual snapshot info in a table
func displaySnapshots(snaps []SnapshotXMLstruct, vmname string) {

	helpers.SurroundText(fmt.Sprintf("All snapshots on %s/%s", helpers.ConnectURI, vmname), false)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Snapshot name", "Current", "Parent", "Creation time"})

	for _, snapshot := range snaps {
		tcreated := time.Unix(snapshot.CreationTime, 0).Format("2006.01.02 15:04:05")
		t.AppendRow([]interface{}{snapshot.SnapshotName, snapshot.CurrentSnapshot, snapshot.Parent.ParentName, tcreated})

	}
	t.SortBy([]table.SortBy{
		{Name: "Snapshot name", Mode: table.Asc},
	})
	t.SetStyle(table.StyleDefault)
	t.Style().Options.DrawBorder = false
	//t.Style().Options.SeparateColumns = false
	t.Style().Format.Header = text.FormatDefault
	t.SetRowPainter(func(row table.Row) text.Colors {
		switch row[1] {
		case true:
			return text.Colors{text.FgHiGreen}
		}
		return nil
	})
	t.Render()
}
