// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// db/db-import.go
// 2022-08-24 22:20:39

package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
	"vmman3/helpers"
)

// Import() : injecte un JSON/YAML dans la BD. LA TABLE SE DOIT D'ÊTRE VIDE. Hard-requirement
func Import(directory string) {
	creds := json2creds()
	var hypervisors []dbHypervisors
	var storagePools []dbStoragePools
	var vmStates []dbVmStates
	var vmClusters []dbClusters

	ctx := context.Background()

	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	hypervisors, storagePools, vmStates, vmClusters = getTables(directory)

	structs2DB(conn, hypervisors, storagePools, vmStates, vmClusters)
}

// Une fonction par table ? Ça aurait plus d'allure....
// getTables() : Collecte les données en format JSON
func getTables(directory string) (hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates, vmc []dbClusters) {

	fname := "hypervisors.json"
	if !checkNOENT(directory, fname) {
		os.Exit(1)
	}
	jsonFile, err := os.ReadFile(helpers.BuildPath(directory, fname))
	if err != nil {
		log.Printf("jsonFile.Get err   #%v ", err)
	}
	err = json.Unmarshal(jsonFile, &hyps)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return hyps, sps, vms, vmc
}

// structs2DB() : Injecte les structures dans la BD
func structs2DB(conn *pgx.Conn, hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates, vmClusters []dbClusters) {

}

// checkNOENT() : Vérifie si le fichier existe, les perms sont OK, ou autre
func checkNOENT(directory string, file string) bool {
	fullpath := helpers.BuildPath(directory, file)
	bExists := true

	_, err := os.Stat(fullpath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s either does not exist or has permission issues. Aborting.\n", fullpath)
			bExists = false
		} else {
			fmt.Printf("Unhandled error with file %s :\n%s.\nAborting.\n", fullpath, err)
			bExists = false
		}
	}

	return bExists
}
