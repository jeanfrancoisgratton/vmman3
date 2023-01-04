package clustermanagement

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"vmman3/helpers"
)

// RemoveFromCluster() : Removes the vm from the database
func RemoveFromCluster(vmname string) {
	envCreds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", envCreds.DbUsr, envCreds.DbPasswd, envCreds.Hostname, envCreds.Port)

	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	hypervisor, _ := os.Hostname()
	sqlQuery := fmt.Sprintf("DELETE FROM clusters WHERE cclustermember='%s:%s';", hypervisor, vmname)
	_, err = dbconn.Exec(context.Background(), sqlQuery)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
}

// memberIsValid() : checks if the member argument is valid
func memberIsValid(member string) bool {
	return true
}
