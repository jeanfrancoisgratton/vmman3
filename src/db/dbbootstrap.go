// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// db/db-bootstrap.go
// 2022-08-22 20:02:37

// TODO
// FILE NEEDS CLEANUP AND GETTING RID OF PASSWORD IN JSON DOCUMENT

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"vmman3/helpers"
)

// CreateDatabase() : action du db bootstrap
func CreateDatabase() {
	var envCreds helpers.EnvironmentStruct

	rcFile, ok := helpers.CheckIfConfigExists()
	if ok {
		envCreds = helpers.Json2creds()
	} else {
		envCreds = getCreds()
		helpers.Creds2json(rcFile, envCreds)
	}

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/postgres", envCreds.RootUsr, envCreds.RootPasswd, envCreds.Hostname, envCreds.Port)

	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	if createUser(dbconn, envCreds.DbUsr, envCreds.DbPasswd) {
		createTablesSchemas(envCreds)
	}
}

// getCreds() : collecte les credentials nécessaires pour se connecter à la BD PGSQL, et créer la BD vmman
func getCreds() helpers.EnvironmentStruct {
	var envCreds helpers.EnvironmentStruct
	var err error

	fmt.Print("Please enter the database hostname: ")
	_, err = fmt.Scanln(&envCreds.Hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %s\n\n", err)
		os.Exit(-1)
	}
	fmt.Print("Please enter the database port: ")
	_, err = fmt.Scanln(&envCreds.Port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %s\n\n", err)
		os.Exit(-1)
	}
	fmt.Print("Please enter the administrative account username: ")
	_, err = fmt.Scanln(&envCreds.RootUsr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %s\n\n", err)
		os.Exit(-1)
	}
	envCreds.RootPasswd = helpers.GetPassword("Please enter that account's password: ")

	fmt.Print("Please enter the application's username: ")
	_, err = fmt.Scanln(&envCreds.DbUsr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %s\n\n", err)
		os.Exit(-1)
	}

	envCreds.DbPasswd = helpers.GetPassword("Please enter the application's user password: ")

	fmt.Print("Please enter the default hypervisor connecting username: ")
	_, err = fmt.Scanln(&envCreds.HypervisorDefaultUser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %s\n\n", err)
		os.Exit(-1)
	}

	fmt.Println()
	return envCreds
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
	dbconn.Exec(ctx, "ALTER DEFAULT PRIVILEGES FOR USER "+username+" IN SCHEMA vmman GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO "+username+";")
	return true
}
