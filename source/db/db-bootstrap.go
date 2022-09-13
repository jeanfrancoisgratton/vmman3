// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// db/db-bootstrap.go
// 2022-08-22 20:02:37

// FIXME FIXME FIXME
// FILE NEEDS CLEANUP AND GETTING RID OF PASSWORD IN JSON/ANY DOCUMENT
// FIXME FIXME FIXME

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"vmman3/helpers"
)

// getCreds() : collecte les credentials nécessaires pour se connecter à la BD PGSQL, et créer la BD vmman
func getCreds() dbCredsStruct {
	var dbCreds dbCredsStruct
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
	fmt.Print("Please enter the root account username: ")
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

// checkIfConfigExists() : Vérifie si le répertoire existe; s'il existe, vérifie si le fichier de config existe
// s'il existe, on l'efface, il sera écrasé plus tard
func checkIfConfigExists() string {
	vmman3rcdir, _ := os.UserHomeDir()
	vmman3rcdir += "/.config/vmman3"

	_, err := os.Stat(vmman3rcdir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(vmman3rcdir, 0700)
		} else {
			panic(err)
		}
	}
	vmman3rcdir += "/database.json"

	_, err = os.Stat(vmman3rcdir)
	if err != nil {
		if os.IsNotExist(err) {
			return vmman3rcdir
		} else {
			panic(err)
		}
	} else {
		os.Remove(vmman3rcdir)
	}
	return vmman3rcdir
}

// CreateDatabase() : action du db bootstrap
func CreateDatabase() {
	var creds dbCredsStruct
	//connStr := "postgresql://<username>:<password>@<database_ip>:<port>/<dbname>?sslmode=disable

	// checkIfConfigExists() needs extra cleanup (subdivisions)
	rcFile := checkIfConfigExists()
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

// createUser() : crée le user vmman
// TODO: error checking
func createUser(conn *pgx.Conn, username string, passwd string) bool {
	conn.Exec(context.Background(), "DROP DATABASE IF EXISTS vmman;")
	conn.Exec(context.Background(), "CREATE DATABASE vmman;")
	conn.Exec(context.Background(), "CREATE ROLE "+username+" CREATEDB INHERIT LOGIN PASSWORD '"+passwd+"';")
	conn.Exec(context.Background(), "GRANT CONNECT ON DATABASE vmman TO "+username+";")
	conn.Exec(context.Background(), "GRANT ALL PRIVILEGES ON DATABASE vmman TO "+username+";")
	conn.Exec(context.Background(), "ALTER USER "+username+" CREATEDB;")
	conn.Exec(context.Background(), "ALTER USER "+username+" WITH SUPERUSER;")
	conn.Exec(context.Background(), "ALTER DEFAULT PRIVILEGES FOR USER "+username+" IN SCHEMA vmman.config GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO "+username+";")

	return true
}

// wipeDB() : Drop la base de données au complet
func wipeDB(conn *pgx.Conn) {

}
