// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// db/db-import.go
// 2022-08-24 22:20:39

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

// Import() : injecte un JSON/YAML dans la BD. LA TABLE SE DOIT D'ÊTRE VIDE. Hard-requirement
func Import(filename string) {
	creds := json2creds()

	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}
