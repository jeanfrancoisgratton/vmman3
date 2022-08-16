// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// helpers/ls.go
// 2022-08-16 17:51:19

package helpers

import (
	"fmt"
	"libvirt.org/libvirt-go"
)

func VM_List() {

	conn, err := libvirt.NewConnect(ConnectURI)
	if err != nil {
		print("whatever")
	}

	defer conn.Close()

	//doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)

	if err != nil {
		print("whatever")
	}

	fmt.Printf("%d running domains:\n", len(doms))

	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			fmt.Printf("  %s\n", name)
		}
		dom.Free()
	}
}
