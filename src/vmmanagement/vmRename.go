// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/vmmanagement/vmRename.go
// 2022-10-29 19:04:58

package vmmanagement

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
)

func Rename(args []string) {
	var snapshotflags libvirt.DomainSnapshotListFlags

	if len(args) < 2 {
		fmt.Println("You need to provide a new name for the VM. Aborting")
		os.Exit(0)
	}
	oldName := args[0]
	newName := args[1]
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
	domain := helpers.GetDomain(conn, oldName)
	if domain == nil {
		os.Exit(0)
	}
	defer domain.Free()
	numsnap, _ := domain.SnapshotNum(snapshotflags)
	if numsnap > 0 {
		fmt.Println("You cannot rename " + oldName + " as this VM holds snapshots. The snapshots need to be removed, first.")
		os.Exit(1)
	}
	helpers.Wait4Shutdown(domain, oldName)
	err = domain.Rename(newName, 0)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("%s --> %s", oldName, newName))

	UpdateRenamedVMinDB(oldName, newName)
}

// Replace the old VM name with the new one in the vmstate and disks tables
func UpdateRenamedVMinDB(oldname string, newname string) {
	var hypervisor string
	if helpers.ConnectURI == "qemu:///system" {
		hypervisor, _ = os.Hostname()
	} else {
		_, _, hypervisor = helpers.SplitConnectURI(helpers.ConnectURI)
	}

	creds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", creds.DbUsr, creds.DbPasswd, creds.Hostname, creds.Port)
	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	//sq := fmt.Sprintf("UPDATE disks SET dvm='%s' WHERE dvm='%s' AND dhypervisor='%s';", newname, oldname, hypervisor)
	_, err = dbconn.Exec(context.Background(), fmt.Sprintf("UPDATE disks SET dvm='%s' WHERE dvm='%s' AND dhypervisor='%s';", newname, oldname, hypervisor))
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	//sq = fmt.Sprintf("UPDATE vmstate SET vmname='%s' WHERE vmname='%s' AND vmhypervisor='%s';", newname, oldname, hypervisor)
	_, err = dbconn.Exec(context.Background(), fmt.Sprintf("UPDATE vmstates SET vmname='%s' WHERE vmname='%s' AND vmhypervisor='%s';", newname, oldname, hypervisor))
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
}
