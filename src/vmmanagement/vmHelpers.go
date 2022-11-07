// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/helpers/vmHelpers.go
// 2022-10-16 10:08:00

package vmmanagement

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"libvirt.org/go/libvirt"
	"log"
	"os"
	"strings"
	"time"
	"vmman3/helpers"
)

// vmStateChange() : updates the database with current VM state
func vmStateChange(hypervisor string, vmname string) {
	creds := helpers.Json2creds()
	ctx := context.Background()
	newDate := time.Now().Format("2006.01.02 15:04:05")
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", creds.DbUsr, creds.DbPasswd, creds.Hostname, creds.Port)

	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	sqlUpdate := fmt.Sprintf("UPDATE vmstates SET vmlaststatechange='%s' WHERE vmhypervisor='%s' AND vmname='%s';", newDate, hypervisor, vmname)
	commandTag, err := dbconn.Exec(ctx, sqlUpdate)
	if err != nil {
		panic(err)
	}
	if commandTag.RowsAffected() != 1 {
		fmt.Println("--> No row found to delete")
	}
}

// wait4Shutdown() : Tries 15 seconds to gracefully shutdown the VM, if not it will shutdown forcefully
func wait4Shutdown(vm *libvirt.Domain, vmname string) {
	var bIsActive = false
	fmt.Println("Will await that the VM " + vmname + " gracefully shuts down on " + helpers.ConnectURI)
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

// getStorage4VM() : Lists all disks from VM
func getStorage4VM(vmname string) ([]string, []string) {
	var hypervisor string
	creds := helpers.Json2creds()
	ctx := context.Background()
	//var poolName string
	//var configuredDisks []string
	if helpers.ConnectURI == "qemu:///system" {
		hypervisor, _ = os.Hostname()
	} else {
		_, _, hypervisor = helpers.SplitConnectURI(helpers.ConnectURI)
	}
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", creds.DbUsr, creds.DbPasswd, creds.Hostname, creds.Port)

	dbconn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	// FIXME: how do we know the order of the returned rows in 1 SQL query matches the order of the 2nd ?
	// 1st: find the disks and storage pool for the given hypervisor+vm combination
	configuredPools, configuredDisks := getDisks(dbconn, vmname, hypervisor)

	// 2nd: find the storagepool paths
	poolPaths := getStoragePoolPaths(dbconn, configuredPools, hypervisor)
	return poolPaths, configuredDisks
}

// getDisks(): Fetches the disks for a given VM+hypervisor combination
func getDisks(dbconn *pgx.Conn, vm string, hypervisor string) ([]string, []string) {
	var configuredPools []string
	var configuredDisks []string

	sqlQuery := fmt.Sprintf("SELECT dname, dpool FROM disks WHERE dvm='%s' AND dhypervisor='%s';", vm, hypervisor)
	rows, err := dbconn.Query(context.Background(), sqlQuery)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var dn, np string
		retcode := rows.Scan(&dn, &np)
		if retcode != nil {
			fmt.Println("Error:", retcode)
		} else {
			// append extension to volume name if it's not already there
			if !strings.HasSuffix(dn, ".qcow2") {
				dn += ".qcow2"
			}
			configuredDisks = append(configuredDisks, dn)
			configuredPools = append(configuredPools, np)
		}
	}
	return configuredPools, configuredDisks
}

// getStoragePoolPaths(): Find the configured path for each defined storagepool used in the given hypervison
func getStoragePoolPaths(dbconn *pgx.Conn, pools []string, hypervisor string) []string {
	var paths []string
	var pathname string

	for _, sp := range pools {
		sqlQuery := fmt.Sprintf("SELECT sppath FROM storagepools WHERE spname='%s' AND (spowner='%s' or spowner='any');", sp, hypervisor)
		err := dbconn.QueryRow(context.Background(), sqlQuery).Scan(&pathname)
		if err != nil {
			panic(err)
		}
		paths = append(paths, pathname)
	}
	return paths
}
