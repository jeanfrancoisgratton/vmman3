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
	creds := Json2creds()
	hypervisors := getHypervisorTable(directory)
	storagePools := getStoragePoolTable(directory)
	vmStates := getVMStatesTable(directory)
	vmClusters := getClustersTable(directory)
	templates := getTemplatesTable(directory)

	ctx := context.Background()
	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	structs2DB(conn, hypervisors, storagePools, vmStates, vmClusters, templates)
	updateSequences(conn)
}

// structs2DB() : Injecte les structures dans la BD
// Ce n'est pas la méthode la plus efficace : on fait un INSERT par ligne, mais la quantité
// De données par table ne justifie pas l'emploi de transactions
func structs2DB(conn *pgx.Conn, hyps []DbHypervisors, sps []dbStoragePools, vms []dbVmStates, vmc []dbClusters, tpt []dbTemplates) {
	ctx := context.Background()
	// hyperviseurs
	for _, h := range hyps {
		sqlStr := fmt.Sprintf("INSERT INTO config.hypervisors (hid, hname, haddress, hconnectinguser) VALUES "+
			"(%d,'%s','%s','%s');", h.HID, h.Hname, h.Haddress, h.Hconnectinguser)
		_, err := conn.Exec(ctx, sqlStr)
		if err != nil {
			panic(err)
		}
	}
	// storagePools
	for _, s := range sps {
		sqlStr := fmt.Sprintf("INSERT INTO config.storagepools (spid, spname, sppath, spowner) VALUES "+
			"(%d,'%s','%s','%s');", s.SpID, s.SpName, s.SpPath, s.SpOwner)
		_, err := conn.Exec(ctx, sqlStr)
		if err != nil {
			panic(err)
		}
	}
	// vmstates
	for _, v := range vms {
		sqlStr := fmt.Sprintf("INSERT INTO config.vmstates "+
			"(vmid, vmname, vmip, vmonline,vmlaststatechange,vmoperatingsystem,vmlasthypervisor,vmstoragepool) VALUES "+
			"(%d,'%s','%s',%t,'%s','%s','%s','%s');", v.VmID, v.VmName, v.VmIP, v.VmOnline, v.VmLastStateChange, v.VmOperatingSystem, v.VmLastHypervisor, v.VmStoragePool)
		_, err := conn.Exec(ctx, sqlStr)
		if err != nil {
			panic(err)
		}
	}
	// clusters
	for _, c := range vmc {
		sqlStr := fmt.Sprintf("INSERT INTO config.clusters (cid, cname) VALUES (%d,'%s');", c.CID, c.Cname)
		_, err := conn.Exec(ctx, sqlStr)
		if err != nil {
			panic(err)
		}
	}
	// templates
	for _, t := range tpt {
		sqlStr := fmt.Sprintf("INSERT INTO config.templates (tid, tname, towner, tstoragepool) "+
			"VALUES (%d,'%s','%s','%s');", t.TID, t.Tname, t.Towner, t.TstoragePool)
		_, err := conn.Exec(ctx, sqlStr)
		if err != nil {
			panic(err)
		}
	}
}

// ALTER SEQUENCE config.xxx RESTART WITH yyy

// updateSequences() : Le nextvalue n'est pas mis à jour après un db import
func updateSequences(conn *pgx.Conn) {
	var vmid, hid, spid, cid, tid uint8
	ctx := context.Background()
	err := conn.QueryRow(ctx, "SELECT MAX(vmid) FROM config.vmstates;").Scan(&vmid)
	if err != nil {
		panic(err)
	}
	err = conn.QueryRow(ctx, "SELECT MAX(hid) FROM config.hypervisors;").Scan(&hid)
	if err != nil {
		panic(err)
	}
	err = conn.QueryRow(ctx, "SELECT MAX(spid) FROM config.storagepools;").Scan(&spid)
	if err != nil {
		panic(err)
	}
	err = conn.QueryRow(ctx, "SELECT MAX(cid) FROM config.clusters;").Scan(&cid)
	if err != nil {
		panic(err)
	}
	err = conn.QueryRow(ctx, "SELECT MAX(tid) FROM config.templates;").Scan(&tid)
	if err != nil {
		panic(err)
	}
	sqlStr := fmt.Sprintf("ALTER SEQUENCE IF EXISTS config.vmstate_vmid_seq RESTART WITH %d;", vmid+1)
	_, err = conn.Exec(ctx, sqlStr)
	if err != nil {
		panic(err)
	}
	sqlStr = fmt.Sprintf("ALTER SEQUENCE IF EXISTS config.hypervisors_hid_seq RESTART WITH %d;", hid+1)
	_, err = conn.Exec(ctx, sqlStr)
	if err != nil {
		panic(err)
	}
	sqlStr = fmt.Sprintf("ALTER SEQUENCE IF EXISTS config.storagepools_spid_seq RESTART WITH %d;", spid+1)
	_, err = conn.Exec(ctx, sqlStr)
	if err != nil {
		panic(err)
	}
	sqlStr = fmt.Sprintf("ALTER SEQUENCE IF EXISTS config.clusters_cid_seq RESTART WITH %d;", cid+1)
	_, err = conn.Exec(ctx, sqlStr)
	if err != nil {
		panic(err)
	}
	sqlStr = fmt.Sprintf("ALTER SEQUENCE IF EXISTS config.templates_tid_seq RESTART WITH %d;", tid+1)
	_, err = conn.Exec(ctx, sqlStr)
	if err != nil {
		panic(err)
	}
}

// getXXXTable() : une fonction par table, pour aller chercher le JSON des tables et l'intégre à la bonne struct
func getHypervisorTable(directory string) []DbHypervisors {
	var hyps []DbHypervisors
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

func getTemplatesTable(directory string) []dbTemplates {
	var dbt []dbTemplates
	fname := "templates.json"
	if !helpers.CheckNOENT(directory, fname) {
		os.Exit(1)
	}
	jsonFile, err := os.ReadFile(helpers.BuildPath(directory, fname))
	if err != nil {
		log.Printf("jsonFile.Get err   #%v ", err)
	}
	err = json.Unmarshal(jsonFile, &dbt)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return dbt
}
