// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/helpers/vmHelpers.go
// 2022-10-16 10:08:00

package helpers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"time"
)

func VMstateChange(hypervisor string, vmname string) {
	creds := Json2creds()
	ctx := context.Background()
	newDate := time.Now().Format("2006.01.02 15:04:05")
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", creds.DbUsr, creds.DbPasswd, creds.Hostname, creds.Port)

	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	sqlUpdate := fmt.Sprintf("UPDATE vmstates SET vmlaststatechange='%s' WHERE vmhypervisor='%s' AND vmname='%s';", newDate, hypervisor, vmname)
	commandTag, err := dbconn.Exec(ctx, sqlUpdate)
	if err != nil {
		panic(err)
	}
	if commandTag.RowsAffected() != 1 {
		fmt.Println("--> No row found to delete")
	}
}
