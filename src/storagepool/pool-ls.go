package storagepool

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"vmman3/db"
	"vmman3/helpers"
)

// PoolList () : Lists all pools from the DB
func PoolList() {
	var storagepool db.DbStoragePools
	var storagepools []db.DbStoragePools
	envCreds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", envCreds.RootUsr, envCreds.RootPasswd, envCreds.Hostname, envCreds.Port)

	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	rows, err := dbconn.Query(context.Background(), "SELECT * from storagepools ORDER BY spid")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		retcode := rows.Scan(&storagepool.SpID, &storagepool.SpName, &storagepool.SpPath, &storagepool.SpOwner)
		if retcode != nil {
			fmt.Println("Error:", retcode)
			os.Exit(-9)
		}
		storagepools = append(storagepools, storagepool)
	}
	helpers.SurroundText("Registered storage pools", false)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	//t.AppendHeader(table.Row{"ID", "Name", "Path", "Owner", "Autostart ?", "Active ?"})
	t.AppendHeader(table.Row{"ID", "Name", "Path", "Owner"})

	for _, sp := range storagepools {
		//var currentHost, isAutostart, isActive string
		//if helpers.ConnectURI == "qemu:///system" {
		//	currentHost, _ = os.Hostname()
		//} else {
		//	_, _, currentHost = helpers.SplitConnectURI(helpers.ConnectURI)
		//}
		//if sp.SpOwner != "any" && currentHost != sp.SpOwner {
		//	isAutostart = "?"
		//	isActive = "?"
		//} else {
		//	isAutostart, isActive = getPoolStatus(sp.SpName)
		//}

		//t.AppendRow([]interface{}{sp.SpID, sp.SpName, sp.SpPath, sp.SpOwner, isAutostart, isActive})
		t.AppendRow([]interface{}{sp.SpID, sp.SpName, sp.SpPath, sp.SpOwner})
	}
	t.SortBy([]table.SortBy{
		{Name: "ID", Mode: table.Asc},
		{Name: "Name", Mode: table.Asc},
	})
	t.SetStyle(table.StyleLight)
	//t.Style().Options.DrawBorder = false
	//t.Style().Options.SeparateColumns = false
	t.Style().Format.Header = text.FormatDefault
	t.Render()
}
