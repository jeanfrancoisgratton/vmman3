// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// db/db-import.go
// 2022-08-24 22:20:39

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

// Import() : injecte un JSON/YAML dans la BD. LA TABLE SE DOIT D'ÊTRE VIDE. Hard-requirement
func Import(directory string) {
	creds := json2creds()
	var hypervisors []dbHypervisors
	var storagePools []dbStoragePools
	var vmStates []dbVmStates

	//ctx := context.Background()

	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	if Bjson {
		hypervisors, storagePools, vmStates = getJsonTables(directory)
	} else {
		hypervisors, storagePools, vmStates = getYamlTables(directory)
	}

	structs2DB(conn, hypervisors, storagePools, vmStates)
}

// getJsonTables() : Collecte les données en format JSON
func getJsonTables(directory string) (hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates) {
	if noEnt(directory)
	return nil, nil, nil
}

// getJsonTables() : Collecte les données en format JSON
func getYamlTables(directory string) (hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates) {
	return nil, nil, nil
}

// structs2DB() : Injecte les structures dans la BD
func structs2DB(conn *pgx.Conn, hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates) {

}