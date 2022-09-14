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

	hypervisors, storagePools, vmStates, vmClusters = getJsonTables(directory)

	structs2DB(conn, hypervisors, storagePools, vmStates, vmClusters)
}

// getJsonTables() : Collecte les données en format YAML
func getJsonTables(directory string) (hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates, vmc []dbClusters) {
	// mapping empty interfaces to data structures
	//dbH := make([]interface{}, len(hyps))
	//for i, v := range hyps {
	//	dbH[i] = v
	//}
	//dbSP := make([]interface{}, len(sps))
	//for i, v := range sps {
	//	dbSP[i] = v
	//}
	//dbVMs := make([]interface{}, len(vms))
	//for i, v := range vms {
	//	dbVMs[i] = v
	//}
	//dbC := make([]interface{}, len(vmc))
	//for i, v := range vmc {
	//	dbC[i] = v
	//}

	dbH, dbSP, dbVMs, dbC := interface2struct(hyps, sps, vms, vmc)
	// populating the tableInfo struct with the above mapping
	tables := []tableInfo{
		{tablename: "hypervisors", datastructure: dbH},
		{tablename: "storagepools", datastructure: dbSP},
		{tablename: "vmstates", datastructure: dbVMs},
		{tablename: "clusters", datastructure: dbC},
	}

	for table := range tables {
		fname := fmt.Sprintf("%s.json", table)
		if !checkNOENT(directory, fname) {
			os.Exit(1)
		}
		// TODO: The 2 next blocks need to be in a for .. range block
		jsonFile, err := os.ReadFile(helpers.BuildPath(directory, fname))
		if err != nil {
			log.Printf("jsonFile.Get err   #%v ", err)
		}
		err = json.Unmarshal(jsonFile, &hyps)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
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