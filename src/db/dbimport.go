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
func structs2DB(conn *pgx.Conn, hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates, vmc []dbClusters) {
	ctx := context.Background()
	// hyperviseurs
	fmt.Print("Hyperviseurs.... ")
	for _, h := range hyps {
		sqlStr := fmt.Sprintf("INSERT INTO config.hypervisors (hid, hname, haddress) VALUES "+
			"(%d,'%s','%s');", h.HID, h.Hname, h.Haddress)
		_, err := conn.Exec(ctx, sqlStr)
		if err != nil {
			fmt.Println("\nErreur: ", err)
			os.Exit(-2)
		}
	}
	fmt.Println("Completé.")
	// storagePools
	fmt.Print("Storage pools.... ")
	for _, s := range sps {
		sqlStr := fmt.Sprintf("INSERT INTO config.storagepools (spid, spname, sppath, spowner) VALUES "+
			"(%d,'%s','%s','%s');", s.SpID, s.SpName, s.SpPath, s.SpOwner)
		_, err := conn.Exec(ctx, sqlStr)
		if err != nil {
			fmt.Println("\nErreur: ", err)
			os.Exit(-2)
		}
	}
	fmt.Println("Completé.")
	// vmstates
	fmt.Print("vm states.... ")
	for _, v := range vms {
		sqlStr := fmt.Sprintf("INSERT INTO config.vmstates "+
			"(vmid, vmname, vmip, vmonline,vmlaststatechange,vmoperatingsystem,vmlasthypervisor,vmstoragepool) VALUES "+
			"(%d,'%s','%s',%t,'%s','%s','%s','%s');", v.VmID, v.VmName, v.VmIP, v.VmOnline, v.VmLastStateChange, v.VmOperatingSystem, v.VmLastHypervisor, v.VmStoragePool)
		_, err := conn.Exec(ctx, sqlStr)
		if err != nil {
			fmt.Println("\nErreur: ", err)
			os.Exit(-2)
		}
	}
	fmt.Println("Completé.")
	// clusters
	fmt.Print("Clusters.... ")
	for _, c := range vmc {
		sqlStr := fmt.Sprintf("INSERT INTO config.clusters (cid, cname) VALUES (%d,'%s');", c.CID, c.Cname)
		_, err := conn.Exec(ctx, sqlStr)
		if err != nil {
			fmt.Println("\nErreur: ", err)
			os.Exit(-2)
		}
	}
	fmt.Println("Completé.")
}

// getXXXTable() : une fonction par table, pour aller chercher le JSON des tables et l'intégre à la bonne struct
func getHypervisorTable(directory string) []dbHypervisors {
	var hyps []dbHypervisors
	fname := "hypervisors.json"
	if !helpers.CheckNOENT(directory, fname) {
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
	if !helpers.CheckNOENT(directory, fname) {
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
	if !helpers.CheckNOENT(directory, fname) {
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
	if !helpers.CheckNOENT(directory, fname) {
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