// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/dbCalls.go
// 2022-09-16 19:44:24

package inventory

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
	"vmman3/db"
)

func listHypervisors() []db.DbHypervisors {
	ctx := context.Background()
	creds := db.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	return db.GetHypervisorData(conn)
}

func getInfoFromDB(hostname string, hypervisor string) (string, string, string) {
	var statechange, operatingsystem, storagepool string
	ctx := context.Background()
	creds := db.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)
	dbconn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer dbconn.Close(ctx)
	querystring := fmt.Sprintf("SELECT vmlaststatechange,vmoperatingsystem,vmstoragepool FROM vmstates WHERE vmname = '%s' AND vmhypervisor = '%s';", hostname, hypervisor)
	err = dbconn.QueryRow(ctx, querystring).
		Scan(&statechange, &operatingsystem, &storagepool)
	if err != nil {
		panic(err)
	}

	return statechange, operatingsystem, storagepool
}
