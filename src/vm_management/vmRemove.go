// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/vm_management/vmRemove.go
// 2022-10-22 12:42:35

package vm_management

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"libvirt.org/go/libvirt"
	"log"
	"os"
	"strings"
	"vmman3/helpers"
)

// This will remove the VM, and optionally leave its storage there
func Remove(args []string) {
	var poolPaths []string
	var vmDisks []string

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
	defer conn.Close()

	for _, vmname := range args {
		//var host string
		domain, err := conn.LookupDomainByName(vmname)
		if err != nil {
			lverr, ok := err.(libvirt.Error)
			if ok {
				if strings.HasPrefix(lverr.Message, "Domain not found") {
					fmt.Println(lverr.Message)
					continue
				} else {
					fmt.Println(lverr.Message)
					os.Exit(-1)
				}
			}
		}
		// Shut the VM down, if active
		wait4Shutdown(domain, vmname)
		fmt.Println(vmname + " now shutdown. Proceeding to removal from inventory.")
		err = domain.UndefineFlags(libvirt.DOMAIN_UNDEFINE_SNAPSHOTS_METADATA)
		if err != nil {
			lverr, ok := err.(libvirt.Error)
			if ok {
				fmt.Println(lverr.Message)
				os.Exit(-1)
			}
		}
		if !helpers.BkeepStorage {
			poolPaths, vmDisks = getStorage4VM(vmname)
			removeStorage(poolPaths, vmDisks)
		}
		// Remove all VM information from the various tables
		removeFromDB(vmname, poolPaths, vmDisks)

		fmt.Println("VM %s has been removed.", vmname)
	}
}

// removeStorage(): remove the VM files from the disks
func removeStorage(paths []string, disks []string) {
	//fullpath := make([]string, len(paths))
	for i, _ := range paths {
		if !strings.HasSuffix(paths[i], "/") {
			paths[i] += "/"
		}
		if !strings.HasSuffix(disks[i], ".qcow2") {
			disks[i] += ".qcow2"
		}
		//fullpath[i] = path
		os.Remove(paths[i] + disks[i])
	}
}

// removeFromDB(): we wipe all VM info from the DB
func removeFromDB(vmname string, poolPaths []string, vmDisks []string) {
	var hypervisor string
	creds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", creds.DbUsr, creds.DbPasswd, creds.Hostname, creds.Port)
	ctx := context.Background()

	if strings.HasPrefix(helpers.ConnectURI, "qemu:///system") {
		hypervisor, _ = os.Hostname()
	} else {
		_, _, hypervisor = helpers.SplitConnectURI(helpers.ConnectURI)
	}
	dbconn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer dbconn.Close(ctx)

	// Table: vmstates
	sqlQuery := fmt.Sprintf("DELETE FROM vmstates WHERE vmname='%s' AND vmhypervisor='%';", vmname, hypervisor)
	_, err = dbconn.Exec(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	// Table: disks
	sqlQuery = fmt.Sprintf("DELETE FROM disks WHERE vmname='%s' AND vmhypervisor='%';", vmname, hypervisor)
	_, err = dbconn.Exec(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
}
