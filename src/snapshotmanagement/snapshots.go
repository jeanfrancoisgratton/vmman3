// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// snapshotmanagement/snapshot.go
// 2022-08-16 17:51:19

package snapshotmanagement

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
)

func GetCurrentSnapshotName(conn *libvirt.Connect, vmname string) string {
	domain, _ := conn.LookupDomainByName(vmname)
	defer domain.Free()
	var currentSnapshot = "none"
	var snapshots, _ = domain.ListAllSnapshots(0)

	for _, snapshot := range snapshots {
		var isCurrent, _ = snapshot.IsCurrent(0)
		if isCurrent {
			currentSnapshot, _ = snapshot.GetName()
			break
		}
	}
	return currentSnapshot
}

func ListSnapshots(vmname string) {
	conn := helpers.Connect2HVM()
	defer conn.Close()

	domain, _ := conn.LookupDomainByName(vmname)
	defer domain.Free()
	numsnap, _ := domain.SnapshotNum(0)
	if numsnap == 0 {
		fmt.Printf("Domain %s has no snapshot\n", vmname)
		os.Exit(0)
	}
	snapshots, _ := domain.ListAllSnapshots(0)

	for ndx, snap := range snapshots {
		data, _ := snap.GetXMLDesc(0)
		writeData(ndx, data)
		//name, parent, creationDate := parseXMLfile(ndx)
		os.Remove(fmt.Sprintf("/tmp/.snapshot-%d.xml", ndx))
	}
}

func writeData(index int, data string) {
	xmlfile := fmt.Sprintf("/tmp/.snapshot-%d.xml", index)

	file, _ := os.Create(xmlfile)
	defer file.Close()
	file.WriteString(data)
	file.Sync()
}

// This is f*ckin' ugly code to get around the fact that I do not know how to properly traverse long XML files in GO :(
func parseXMLfile(index int) (string, string, string) {
	// https://github.com/antchfx/xmlquery
	return "", "", ""
}

//	var snapname, snapparent, snapcreationdate string
//	var linesInFile []string
//
//	// STEP ONE: dump file in a []string var
//	rfile, err := os.Open(fmt.Sprintf("/tmp/.snapshot-%d.xml", index))
//	defer rfile.Close()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(-1)
//	}
//	fscanner := bufio.NewScanner(rfile)
//	fscanner.Split(bufio.ScanLines)
//
//	for fscanner.Scan() {
//		linesInFile = append(linesInFile, fscanner.Text())
//	}
//	// STEP TWO: iterate string slice to find the wanted values
//	for ndx, str := range linesInFile {
//		// snapshot name
//		if str == "<domainsnapshot>" {
//			snapparent = extractSname(linesInFile[ndx+1])
//		}
//		//a := strings.TrimSpace(str)
//		//if strings.HasPrefix(a, "name")
//	}
//	return "", "", ""
//}
//
//func extractSname(line string) string {
//	a := strings.TrimSpace(line)
//}
