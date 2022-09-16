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
	hypervisors := getHypervisorTable(directory)
	storagePools := getStoragePoolTable(directory)
	vmStates := getVMStatesTable(directory)
	vmClusters := getClustersTable(directory)

	ctx := context.Background()

	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	structs2DB(conn, hypervisors, storagePools, vmStates, vmClusters)
}

// structs2DB() : Injecte les structures dans la BD
// Ce n'est pas la méthode la plus efficace : on fait un INSERT par ligne, mais la quantité
// De données par table ne justifie pas l'emploi de transactions
func structs2DB(conn *pgx.Conn, hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates, vmClusters []dbClusters) {
	ctx := context.Background()
	// hyperviseurs
	for _, h := range hyps {
		sqlStr := fmt.Sprintf("INSERT INTO config.hypervisors (hid, hname, haddress) VALUES %s,%s,%s", h.HID, h.Hname, h.Haddress)
		conn.Exec(ctx, sqlStr)
	}
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

// getXXXTable() : une fonction par table, pour aller chercher le JSON des tables et l'intégre à la bonne struct
func getHypervisorTable(directory string) []dbHypervisors {
	var hyps []dbHypervisors
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
	return hyps
}

func getStoragePoolTable(directory string) []dbStoragePools {
	var sps []dbStoragePools
	fname := "storagepools.json"
	if !checkNOENT(directory, fname) {
		os.Exit(1)
	}
	jsonFile, err := os.ReadFile(helpers.BuildPath(directory, fname))
	if err != nil {
		log.Printf("jsonFile.Get err   #%v ", err)
	}
	err = json.Unmarshal(jsonFile, &sps)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return sps
}

func getVMStatesTable(directory string) []dbVmStates {
	var vms []dbVmStates
	fname := "vmstates.json"
	if !checkNOENT(directory, fname) {
		os.Exit(1)
	}
	jsonFile, err := os.ReadFile(helpers.BuildPath(directory, fname))
	if err != nil {
		log.Printf("jsonFile.Get err   #%v ", err)
	}
	err = json.Unmarshal(jsonFile, &vms)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return vms
}

func getClustersTable(directory string) []dbClusters {
	var dbc []dbClusters
	fname := "clusters.json"
	if !checkNOENT(directory, fname) {
		os.Exit(1)
	}
	jsonFile, err := os.ReadFile(helpers.BuildPath(directory, fname))
	if err != nil {
		log.Printf("jsonFile.Get err   #%v ", err)
	}
	err = json.Unmarshal(jsonFile, &dbc)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return dbc
}
