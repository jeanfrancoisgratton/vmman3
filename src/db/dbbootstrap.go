// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// db/db-bootstrap.go
// 2022-08-22 20:02:37

// FIXME FIXME FIXME
// FILE NEEDS CLEANUP AND GETTING RID OF PASSWORD IN JSON DOCUMENT
// FIXME FIXME FIXME

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"vmman3/helpers"
)

// CreateDatabase() : action du db bootstrap
func CreateDatabase() {
	var creds DbCredsStruct
	//connStr := "postgresql://<username>:<password>@<database_ip>:<port>/<dbname>?sslmode=disable

	// checkIfConfigExists() needs extra cleanup (subdivisions)
	rcFile := helpers.CheckIfConfigExists()
	if rcFile != "" {
		creds = getCreds()
		creds2json(rcFile, creds)
	}

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/postgres", creds.RootUsr, creds.RootPasswd, creds.Hostname, creds.Port)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	if createUser(conn, creds.DbUsr, creds.DbPasswd) {
		createTablesSchemas(creds.Hostname, creds.Port)
	}
}

// getCreds() : collecte les credentials nécessaires pour se connecter à la BD PGSQL, et créer la BD vmman
func getCreds() DbCredsStruct {
	var dbCreds DbCredsStruct
	var err error

	fmt.Print("Please enter the database hostname: ")
	_, err = fmt.Scanln(&dbCreds.Hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %s\n\n", err)
		os.Exit(-1)
	}
	fmt.Print("Please enter the database port: ")
	_, err = fmt.Scanln(&dbCreds.Port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %s\n\n", err)
		os.Exit(-1)
	}
	fmt.Print("Please enter the administrative account username: ")
	_, err = fmt.Scanln(&dbCreds.RootUsr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %s\n\n", err)
		os.Exit(-1)
	}
	dbCreds.RootPasswd = helpers.GetPassword("Please enter that account's password: ")

	fmt.Print("Please enter the application's username: ")
	_, err = fmt.Scanln(&dbCreds.DbUsr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %s\n\n", err)
		os.Exit(-1)
	}

	dbCreds.DbPasswd = helpers.GetPassword("Please enter the application's user password: ")

	fmt.Println()
	return dbCreds
}

// createUser() : crée le user vmman
// TODO: error checking
func createUser(dbconn *pgx.Conn, username string, passwd string) bool {
	ctx := context.Background()
	_, err := dbconn.Exec(ctx, "DROP DATABASE IF EXISTS vmman;")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	_, err = dbconn.Exec(context.Background(), "CREATE DATABASE vmman;")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	dbconn.Exec(ctx, "CREATE ROLE "+username+" CREATEDB INHERIT LOGIN PASSWORD '"+passwd+"';")
	dbconn.Exec(ctx, "GRANT CONNECT ON DATABASE vmman TO "+username+";")
	dbconn.Exec(ctx, "GRANT ALL PRIVILEGES ON DATABASE vmman TO "+username+";")
	dbconn.Exec(ctx, "ALTER USER "+username+" CREATEDB;")
	dbconn.Exec(ctx, "ALTER USER "+username+" WITH SUPERUSER;")
	dbconn.Exec(ctx, "ALTER DEFAULT PRIVILEGES FOR USER "+username+" IN SCHEMA vmman.public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO "+username+";")

	return true
}
