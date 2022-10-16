// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// db/db-create-tables.go
// 2022-08-24 21:28:02

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

// createTablesSchemas() : crée la bd, schemas et tables
// TODO : transactions, anyone ? :p
func createTablesSchemas(hostname string, port int) {
	connString := fmt.Sprintf("postgresql://vmman:vmman@%s:%d/vmman", hostname, port)
	ctx := context.Background()

	dbconn, err := pgx.Connect(ctx, connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	//fmt.Print("Drop/Create... ")
	//dbconn.Exec(ctx, "DROP SCHEMA IF EXISTS config ;")
	//dbconn.Exec(ctx, "CREATE SCHEMA IF NOT EXISTS config AUTHORIZATION vmman;")
	//fmt.Print("Completed\n")
	fmt.Print("Sequences... ")
	createSeqs(dbconn)
	fmt.Print("Completed\n")
	fmt.Print("Tables... ")
	createTables(dbconn)
	fmt.Print("Completed\n")
	fmt.Print("Ownership... ")
	setTableOwnership(dbconn)
	fmt.Print("Completed\n\n")
}

// createSeqs() : crée les sequences dans la BD
func createSeqs(dbconn *pgx.Conn) {
	ctx := context.Background()
	//dbconn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"storagepools_spid_seq\" "+
	//	"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
	dbconn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS \"storagepools_spid_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
	dbconn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS \"hypervisors_hid_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
	dbconn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS \"vmstate_vmid_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;")
	dbconn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS \"clusters_cid_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
	dbconn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS \"templates_tid_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
}

// createTables() : crée les tables dans la BD
// TODO : transactions, anyone ? :p
func createTables(dbconn *pgx.Conn) {
	ctx := context.Background()
	_, err := dbconn.Exec(ctx, "CREATE TABLE IF NOT EXISTS storagepools "+
		"(spid smallint NOT NULL DEFAULT nextval('\"storagepools_spid_seq\"'::regclass), "+
		"spname character varying(24) NOT NULL, sppath character varying(512) NOT NULL, "+
		"spowner character varying(24) NOT NULL DEFAULT 'localhost'::character varying, "+
		"CONSTRAINT storagepools_pkey PRIMARY KEY (spid));")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	_, err = dbconn.Exec(ctx, "CREATE TABLE IF NOT EXISTS hypervisors "+
		"(hid smallint NOT NULL DEFAULT nextval('\"hypervisors_hid_seq\"'::regclass),"+
		"hname character varying(24) NOT NULL, haddress character varying(128) NOT NULL DEFAULT '127.0.0.1'::character varying,"+
		"hconnectinguser character varying(16) NOT NULL DEFAULT 'root',CONSTRAINT hypervisors_pkey PRIMARY KEY (hid));")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	_, err = dbconn.Exec(ctx, "CREATE TABLE IF NOT EXISTS vmstates "+
		"(vmid integer NOT NULL DEFAULT nextval('\"vmstate_vmid_seq\"'::regclass), "+
		"vmname character varying(24) NOT NULL, vmip character varying(15), vmonline boolean NOT NULL DEFAULT false, "+
		"vmlaststatechange character varying(24) NOT NULL DEFAULT 'unknown', "+
		"vmoperatingsystem character varying(50) NOT NULL DEFAULT 'linux', "+
		"vmhypervisor character varying(24) NOT NULL, "+
		"vmstoragepool character varying(24) NOT NULL DEFAULT 'vmpool', "+
		"CONSTRAINT vmState_pkey PRIMARY KEY (vmid));")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	_, err = dbconn.Exec(ctx, "CREATE TABLE IF NOT EXISTS clusters "+
		"(cid smallint NOT NULL DEFAULT nextval('\"clusters_cid_seq\"'::regclass), "+
		"cname character varying(24) NOT NULL, CONSTRAINT clusters_pkey PRIMARY KEY (cid));")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}

	_, err = dbconn.Exec(ctx, "CREATE TABLE IF NOT EXISTS templates "+
		"(tid smallint NOT NULL DEFAULT nextval('\"templates_tid_seq\"'::regclass), "+
		"tname character varying(24) NOT NULL, towner character varying(24) NOT NULL, "+
		"tstoragepool character varying(24), toperatingsystem character varying(50) NOT NULL DEFAULT 'linux', "+
		"CONSTRAINT templates_pkey PRIMARY KEY (tid));")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
}

// setTableOwnership() : change la propriété des tables pour vmman
func setTableOwnership(dbconn *pgx.Conn) {
	ctx := context.Background()
	dbconn.Exec(ctx, "ALTER TABLE IF EXISTS storagePools OWNER to vmman;")
	dbconn.Exec(ctx, "ALTER TABLE IF EXISTS hypervisors OWNER to vmman;")
	dbconn.Exec(ctx, "ALTER TABLE IF EXISTS vmStates OWNER to vmman;")
	dbconn.Exec(ctx, "ALTER TABLE IF EXISTS clusters OWNER to vmman;")
}
