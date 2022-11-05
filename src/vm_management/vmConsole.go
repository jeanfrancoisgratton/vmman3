// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/vm_management/vmConsole.go
// 2022-11-05 15:48:18

package vm_management

import (
	"os"
	"os/exec"
	"vmman3/helpers"
)

// Console() : opens the VM's console
// https://pkg.go.dev/libvirt.org/go/libvirt#Domain
// https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainOpenConsole
// FIXME: this does not work, so I'll cheat....
//func Console(vmname string) {
//	var stream *libvirt.Stream
//	conn := helpers.Connect2HVM()
//	defer conn.Close()
//
//	domain, _ := conn.LookupDomainByName(vmname)
//	isup, _ := domain.IsActive()
//	if !isup {
//		fmt.Printf("%s needs to be up in order to access its console.")
//		os.Exit(0)
//	}
//	defer domain.Free()
//
//	domain.OpenConsole("", stream, 0)
//}

// Console() VERSION 2: ugly hack using the shell :p
func Console(vmname string) {
	//connectCmd := fmt.Sprintf("/usr/bin/virsh -c %s console %s", helpers.ConnectURI, vmname)

	cmd := exec.Command("virsh", "-c", helpers.ConnectURI, "console", vmname)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
