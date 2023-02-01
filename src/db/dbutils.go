// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/db-utils.go
// 2022-08-25 13:32:28

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"strings"
	"vmman3/helpers"
)

// This might get converted to generics, at some point
// func interface2struct(hyps []DbHypervisors, sps []DbStoragePools, vms []dbVmStates, vmc []DbClusters) ([]interface{}, []interface{}, []interface{}, []interface{}) {
func interface2struct(hyps []DbHypervisors, sps []DbStoragePools, vms []dbVmStates) ([]interface{}, []interface{}, []interface{}) {
	dbH := make([]interface{}, len(hyps))
	for i, v := range hyps {
		dbH[i] = v
	}
	dbSP := make([]interface{}, len(sps))
	for i, v := range sps {
		dbSP[i] = v
	}
	dbVMs := make([]interface{}, len(vms))
	for i, v := range vms {
		dbVMs[i] = v
	}
	//dbC := make([]interface{}, len(vmc))
	//for i, v := range vmc {
	//	dbC[i] = v
	//}

	//return dbH, dbSP, dbVMs, dbC
	return dbH, dbSP, dbVMs
}

// Drop() : drop database
func Drop() {
	creds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/postgres", creds.RootUsr, creds.RootPasswd, creds.Hostname, creds.Port)
	ctx := context.Background()
	var confirmation string

	fmt.Print("WARNING !!! Host: " + creds.Hostname + ": This operation is irreversible. Are you sure you want to continue [Y/n] ? ")
	fmt.Scanln(&confirmation)

	if !strings.HasPrefix(strings.ToLower(confirmation), "y") {
		os.Exit(0)
	}

	dbconn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer dbconn.Close(ctx)

	_, err = dbconn.Exec(ctx, "DROP DATABASE IF EXISTS vmman;")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-2)
	}
}
