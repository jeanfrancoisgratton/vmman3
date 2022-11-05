// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/vm_management/vmXmlDump.go
// 2022-11-05 13:45:39

package vm_management

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"strings"
	"vmman3/helpers"
)

// XmlDump() : dumps the vm config in an xml file
// libvirt call: https://pkg.go.dev/libvirt.org/go/libvirt#Domain.GetXMLDesc
func XmlDump(vmname string, xmlfile string) {
	var data string
	var file *os.File

	if !strings.HasSuffix(xmlfile, ".xml") {
		xmlfile += ".xml"
	}
	conn, err := libvirt.NewConnect(helpers.ConnectURI)

	if err != nil {
		lverr, ok := err.(libvirt.Error)
		if ok && lverr.Message == "End of file while reading data: virt-ssh-helper: cannot connect to '/var/run/libvirt/libvirt-sock': Failed to connect socket to '/var/run/libvirt/libvirt-sock': Connection refused: Input/output error" {
			fmt.Printf("Hypervisor %s is offline\n", helpers.ConnectURI)
			return
		} else {
			panic(err)
		}
	}

	domain, _ := conn.LookupDomainByName(vmname)
	defer domain.Free()

	wait4Shutdown(domain, vmname)
	data, err = domain.GetXMLDesc(libvirt.DOMAIN_XML_SECURE | libvirt.DOMAIN_XML_INACTIVE | libvirt.DOMAIN_XML_MIGRATABLE)

	file, err = os.Create(xmlfile)
	defer file.Close()
	file.WriteString(data)
	file.Sync()
}
