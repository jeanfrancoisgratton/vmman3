// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/vmmanagement/vmXmlDump.go
// 2022-11-05 13:45:39

package vmmanagement

import (
	"libvirt.org/go/libvirt"
	"os"
	"strings"
	"vmman3/helpers"
)

// XmlDump() : dumps the vm config in an xml file
// libvirt call: https://pkg.go.dev/libvirt.org/go/libvirt#Domain.GetXMLDesc
// TODO: more robust error handling here...
func XmlDump(vmname string, xmlfile string) {
	if !strings.HasSuffix(xmlfile, ".xml") {
		xmlfile += ".xml"
	}
	conn := helpers.Connect2HVM()
	defer conn.Close()

	domain, _ := conn.LookupDomainByName(vmname)
	defer domain.Free()

	helpers.Wait4Shutdown(domain, vmname)
	data, _ := domain.GetXMLDesc(libvirt.DOMAIN_XML_SECURE | libvirt.DOMAIN_XML_INACTIVE | libvirt.DOMAIN_XML_MIGRATABLE)

	file, _ := os.Create(xmlfile)
	defer file.Close()
	file.WriteString(data)
	file.Sync()
}
