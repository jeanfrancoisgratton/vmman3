// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/snapshotmanagement/snapHelpers.go
// 2022-11-09 19:15:22

package snapshotmanagement

import "encoding/xml"

// Snapshot XML definitions
type ParentElement struct {
	XMLName    xml.Name `xml:"parent"`
	ParentName string   `xml:"name"`
}
type SnapshotXMLstruct struct {
	SnapshotName    string        `xml:"name"`
	CreationTime    uint64        `xml:"creationTime"`
	Parent          ParentElement `xml:"parent"`
	CurrentSnapshot bool
}
