// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/helpers/globalHelpers.go
// 2022-11-05 16:47:17

package helpers

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"strings"
	"time"
)

// Connect2HVM() : connects to the hypervisor, returning the Connect object
func Connect2HVM() *libvirt.Connect {
	conn, err := libvirt.NewConnect(ConnectURI)
	if err != nil {
		lverr, ok := err.(libvirt.Error)
		if ok && (lverr.Message == "End of file while reading data: virt-ssh-helper: cannot connect to '/var/run/libvirt/libvirt-sock': Failed to connect socket to '/var/run/libvirt/libvirt-sock': Connection refused: Input/output error" ||
			strings.HasPrefix(lverr.Message, "Cannot recv data: ssh: connect to host")) {
			fmt.Printf("Hypervisor %s is offline\n", ConnectURI)
			os.Exit(-1)
		} else {
			panic(err)
		}
	}

	return conn
}

// GetDomain() : Connects to the VM (domain), returning the Domainobject
func GetDomain(conn *libvirt.Connect, vmname string) *libvirt.Domain {
	domain, err := conn.LookupDomainByName(vmname)
	if err != nil {
		lverr, ok := err.(libvirt.Error)
		if ok {
			if strings.HasPrefix(lverr.Message, "Domain not found") {
				fmt.Println(lverr.Message)
				return nil
			} else {
				fmt.Println(lverr.Message)
				os.Exit(-1)
			}
		}
	}
	return domain
}

// Wait4Shutdown() : Tries 15 seconds to gracefully shutdown the VM, if not it will shutdown forcefully
func Wait4Shutdown(vm *libvirt.Domain, vmname string) {
	var bIsActive = false
	fmt.Println("Will await that the VM " + vmname + " gracefully shuts down on " + ConnectURI)
	bIsActive, _ = vm.IsActive()
	if bIsActive {
		n := 15
		vm.DestroyFlags(libvirt.DOMAIN_DESTROY_GRACEFUL)
		for n > 0 {
			bIsActive, _ = vm.IsActive()
			if bIsActive {
				n -= 1
				time.Sleep(1 * time.Second)
			} else {
				n = 0
			}
		}
		bIsActive, _ = vm.IsActive()
		if bIsActive {
			vm.DestroyFlags(libvirt.DOMAIN_DESTROY_DEFAULT)
			fmt.Println("The VM " + vmname + " was slow to shutdown and was forcely shut down")
		}
	}
}
