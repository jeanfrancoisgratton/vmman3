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
	conn.Exec(context.Background(), "DROP SCHEMA IF EXISTS configs ;")
	conn.Exec(context.Background(), "CREATE SCHEMA IF NOT EXISTS configs AUTHORIZATION vmman;")
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
	conn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"storagePools_spID_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
	conn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"hypervisors_hID_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
	conn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"vmState_vmId_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;")
	conn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"clusters_cID_seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
	conn.Exec(ctx, "CREATE SEQUENCE IF NOT EXISTS config.\"servers_sID__seq\" "+
		"INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 32767 CACHE 1;")
}

// createTables() : crée les tables dans la BD
// FIXME : transactions, anyone ? :p
func createTables(conn *pgx.Conn) {
	ctx := context.Background()
	conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS config.storagePools "+
		"(spID smallint NOT NULL DEFAULT nextval('config.\"storagePools_spID_seq\"'::regclass),"+
		"spName character varying(24) NOT NULL, spPath character varying(512) NOT NULL,"+
		"spOwner character varying(24) CONSTRAINT storagePools_pkey PRIMARY KEY (spID)) TABLESPACE pg_default;")

	conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS config.hypervisors "+
		"(hID smallint NOT NULL DEFAULT nextval('config.\"hypervisors_hID_seq\"'::regclass),"+
		"hName character varying(24) NOT NULL, hAddress character varying(128) NOT NULL DEFAULT '127.0.0.1'::character varying,"+
		"CONSTRAINT hypervisors_pkey PRIMARY KEY (hID));")

	conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS config.vmState "+
		"(vmId integer NOT NULL DEFAULT nextval('config.\"vmState_vmId_seq\"'::regclass), "+
		"vmName character varying(24) NOT NULL, vmIP inet, vmOnline boolean NOT NULL DEFAULT false, "+
		"vmLastStateChange character varying(24) NOT NULL DEFAULT 'unseen', "+
		"CONSTRAINT vmState_pkey PRIMARY KEY (vmId));")

	conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS config.clusters "+
		"(cid smallint NOT NULL DEFAULT nextval('config.\"clusters_cID_seq\"'::regclass), "+
		"cname character(24) NOT NULL DEFAULT 'cluster'::bpchar, CONSTRAINT clusters_pkey PRIMARY KEY (cid)));")

	conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS config.servers "+
		"(sid smallint NOT NULL DEFAULT nextval('config.\"servers_sID_seq\"'::regclass), "+
		"sname character(16) NOT NULL DEFAULT 'server'::bpchar, soperatingsystem character(50) NOT NULL DEFAULT 'linux'::bpchar, "+
		"slasthypervisor character(24) NOT NULL DEFAULT 'localhost', CONSTRAINT servers_pkey PRIMARY KEY (sid));")

}

// setOwnership() : change la propriété des tables pour vmman
// FIXME : transactions, anyone ? :p
func setTableOwnership(conn *pgx.Conn) {
	ctx := context.Background()
	conn.Exec(ctx, "ALTER TABLE IF EXISTS config.storagePools OWNER to vmman;")
	conn.Exec(ctx, "ALTER TABLE IF EXISTS config.hypervisors OWNER to vmman;")
	conn.Exec(ctx, "ALTER TABLE IF EXISTS config.vmState OWNER to vmman;")
	conn.Exec(ctx, "ALTER TABLE IF EXISTS config.clusters OWNER to vmman;")
	conn.Exec(ctx, "ALTER TABLE IF EXISTS config.servers OWNER to vmman;")
}
