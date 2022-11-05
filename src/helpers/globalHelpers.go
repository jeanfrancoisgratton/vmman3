// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/helpers/globalHelpers.go
// 2022-11-05 16:47:17

package helpers

import (
	"fmt"
	"libvirt.org/go/libvirt"
)

func Connect2HVM() *libvirt.Connect {
	conn, err := libvirt.NewConnect(ConnectURI)
	if err != nil {
		lverr, ok := err.(libvirt.Error)
		if ok && lverr.Message == "End of file while reading data: virt-ssh-helper: cannot connect to '/var/run/libvirt/libvirt-sock': Failed to connect socket to '/var/run/libvirt/libvirt-sock': Connection refused: Input/output error" {
			fmt.Printf("Hypervisor %s is offline\n", ConnectURI)
			return nil
		} else {
			panic(err)
		}
	}

	return conn
}
