package snapshot_management

import (
	"fmt"
	"libvirt.org/libvirt-go"
	"vmman3/helpers"
)

func GetCurrentSnapshot(vmname string) {
	conn, err := libvirt.NewConnect(helpers.ConnectURI)
	if err != nil {
		fmt.Println("Error while connecting: ", err)
	}
	defer conn.Close()
}
