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
// FIXME : transactions, anyone ? :p
func createTablesSchemas(hostname string, port int) {
	connString := fmt.Sprintf("postgresql://vmman:vmman@%s:%d/vmman", hostname, port)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Print("Drop/Create... ")
	conn.Exec(context.Background(), "DROP SCHEMA IF EXISTS config ;")
	conn.Exec(context.Background(), "CREATE SCHEMA IF NOT EXISTS config AUTHORIZATION vmman;")
	fmt.Print("Completed\n")
	fmt.Print("Sequences... ")
	createSeqs(conn)
	fmt.Print("Completed\n")
	fmt.Print("Tables... ")
	createTables(conn)
	fmt.Print("Completed\n")
	fmt.Print("Ownership... ")
	setTableOwnership(conn)
	fmt.Print("Completed\n\n")
}

// createSeqs() : crée les sequences dans la BD
func createSeqs(conn *pgx.Conn) {
	ctx := context.Background()
	conn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"storagepools_spid_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
	conn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"hypervisors_hid_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
	conn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"vmstate_vmid_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;")
	conn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"clusters_cid_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
}

// createTables() : crée les tables dans la BD
// TODO : transactions, anyone ? :p
func createTables(conn *pgx.Conn) {
	ctx := context.Background()
	_, err := conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS config.storagepools "+
		"(spid smallint NOT NULL DEFAULT nextval('config.\"storagepools_spid_seq\"'::regclass), "+
		"spname character varying(24) NOT NULL, sppath character varying(512) NOT NULL, "+
		"spowner character varying(24) NOT NULL DEFAULT 'localhost'::character varying, "+
		"CONSTRAINT storagepools_pkey PRIMARY KEY (spid));")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	_, err = conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS config.hypervisors "+
		"(hid smallint NOT NULL DEFAULT nextval('config.\"hypervisors_hid_seq\"'::regclass),"+
		"hname character varying(24) NOT NULL, haddress character varying(128) NOT NULL DEFAULT '127.0.0.1'::character varying,"+
		"hconnectinguser character varying(16) NOT NULL DEFAULT 'root',CONSTRAINT hypervisors_pkey PRIMARY KEY (hid));")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	_, err = conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS config.vmstates "+
		"(vmid integer NOT NULL DEFAULT nextval('config.\"vmstate_vmid_seq\"'::regclass), "+
		"vmname character varying(24) NOT NULL, vmip character varying(15), vmonline boolean NOT NULL DEFAULT false, "+
		"vmlaststatechange character varying(24) NOT NULL DEFAULT 'unseen', "+
		"vmoperatingsystem character varying(50) NOT NULL DEFAULT 'linux', "+
		"vmlasthypervisor character varying(24) NOT NULL DEFAULT 'unseen', "+
		"vmstoragepool character varying(24) NOT NULL DEFAULT 'vmpool', "+
		"CONSTRAINT vmState_pkey PRIMARY KEY (vmid));")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	_, err = conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS config.clusters "+
		"(cid smallint NOT NULL DEFAULT nextval('config.\"clusters_cid_seq\"'::regclass), "+
		"cname character varying(24) NOT NULL, CONSTRAINT clusters_pkey PRIMARY KEY (cid));")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
}

// setTableOwnership() : change la propriété des tables pour vmman
func setTableOwnership(conn *pgx.Conn) {
	ctx := context.Background()
	conn.Exec(ctx, "ALTER TABLE IF EXISTS config.storagePools OWNER to vmman;")
	conn.Exec(ctx, "ALTER TABLE IF EXISTS config.hypervisors OWNER to vmman;")
	conn.Exec(ctx, "ALTER TABLE IF EXISTS config.vmStates OWNER to vmman;")
	conn.Exec(ctx, "ALTER TABLE IF EXISTS config.clusters OWNER to vmman;")
}
