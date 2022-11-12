// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/hypervisor/hypervisor.go
// 2022-11-11 21:29:08

package hypervisor

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"vmman3/helpers"
)

// AddHypervisor() : adds an hypervisor in the DB
// USAGE: vmman hypervisor add HYPERVISOR_NAME HYPERVISOR_ADDR HYPERVISOR_USER
func AddHypervisor(name, address, user string) {
	envCreds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", envCreds.RootUsr, envCreds.RootPasswd, envCreds.Hostname, envCreds.Port)

	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	sqlString := fmt.Sprintf("INSERT INTO hypervisors (hname,haddress,hconnectinguser) VALUES ('%s', '%s', '%s');", name, address, user)
	_, err = dbconn.Exec(context.Background(), sqlString)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	fmt.Printf("Completed:\n%s\n", sqlString)
}

// DelHypervisor() : deletes an hypervisor from DB
// USAGE: vmman hypervisor rm HYPERVISOR_NAME [HYPERVISOR_USER]
func DelHypervisor(name, user string) {
	whereClause := ";"

	if user != "" {
		whereClause = fmt.Sprintf(" AND hconnectinguser='%s';", user)
	}
	envCreds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", envCreds.RootUsr, envCreds.RootPasswd, envCreds.Hostname, envCreds.Port)

	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	sqlString := fmt.Sprintf("DELETE FROM hypervisors WHERE hname='%s'%s", name, whereClause)
	_, err = dbconn.Exec(context.Background(), sqlString)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
	fmt.Printf("Completed:\n%s\n", sqlString)
}
