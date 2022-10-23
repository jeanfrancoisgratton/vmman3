// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/helpers/vmHelpers.go
// 2022-10-16 10:08:00

package helpers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"libvirt.org/go/libvirt"
	"log"
	"os"
	"strings"
	"time"
)

type VmStorageDetails struct {
	Poolname string
	Diskname string
	PoolPath string
}

// VmStateChange() : updates the database with current VM state
func VmStateChange(hypervisor string, vmname string) {
	creds := Json2creds()
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
			}
		}
		bIsActive, _ = vm.IsActive()
		if bIsActive {
			vm.DestroyFlags(libvirt.DOMAIN_DESTROY_DEFAULT)
			fmt.Println("The VM " + vmname + " was slow to shutdown and was forcely shut down")
		}
	}
}

// GetStorage4VM() : Lists all disks from VM
func GetStorage4VM(vmname string) []VmStorageDetails {
	creds := Json2creds()
	ctx := context.Background()
	var sqlQuery string
	var storage []VmStorageDetails
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", creds.DbUsr, creds.DbPasswd, creds.Hostname, creds.Port)

	dbconn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())
	_, _, host := SplitConnectURI(ConnectURI)

	// TODO: THIS NEEDS A FINER SQL QUERY

	// 1st: find the pathname for StoragePool
	sqlQuery = fmt.Sprintf("SELECT spath FROM storagepools WHERE spiw='%s';",
	sqlQuery := fmt.Sprintf("SELECT dname, dpool FROM disks WHERE dvm='%s' AND dhypervisor='%s';", vmname, host)
	rows, err := dbconn.Query(ctx, sqlQuery)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sr VmStorageDetails
		retcode := rows.Scan(&sr.Diskname, &sr.Poolname)
		if retcode != nil {
			fmt.Println("Error:", retcode)
		} else {
			// append extension to volume name if it's not already there
			if !strings.HasSuffix(sr.Diskname, ".qcow2") {
				sr.Diskname += ".qcow2"
			}
			storage = append(storage, sr)
		}
	}

	return storage
}
