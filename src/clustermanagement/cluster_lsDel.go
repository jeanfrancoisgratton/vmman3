package clustermanagement

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jwalton/gchalk"
	"log"
	"os"
	"vmman3/db"
	"vmman3/helpers"
)

// ListClusters() : simple cluster enumeration
func ListClusters() {
	var cluster db.DbClusters
	var clusters []db.DbClusters

	envCreds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", envCreds.DbUsr, envCreds.DbPasswd, envCreds.Hostname, envCreds.Port)

	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	rows, err := dbconn.Query(context.Background(), "SELECT DISTINCT (cname) from clusters ORDER BY cname")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		retcode := rows.Scan(&cluster.Cname)
		if retcode != nil {
			fmt.Println("Error:", retcode)
			os.Exit(-9)
		}
		clusters = append(clusters, cluster)
	}
	helpers.SurroundText("Registered clusters", false)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name"})

	for _, cluster := range clusters {
		t.AppendRow([]interface{}{cluster.Cname})
	}
	t.SortBy([]table.SortBy{
		{Name: "ID", Mode: table.Asc},
		{Name: "Name", Mode: table.Asc},
	})
	t.SetStyle(table.StyleDefault)
	//t.Style().Options.DrawBorder = false
	//t.Style().Options.SeparateColumns = false
	t.Style().Format.Header = text.FormatDefault
	t.Render()
}

// RemoveCluster() : remove the cluster or list of clusters
func RemoveCluster(args []string) {
	creds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", creds.DbUsr, creds.DbPasswd, creds.Hostname, creds.Port)
	ctx := context.Background()

	dbconn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer dbconn.Close(ctx)

	for _, arg := range args {
		sqlQuery := fmt.Sprintf("DELETE FROM clusters WHERE cname='%s';", arg)
		_, err = dbconn.Exec(ctx, sqlQuery)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("Cluster %s has been removed\n\n\n", arg)
		fmt.Printf("Cluster %s has been %s from the database.\n", gchalk.WithBrightWhite().Bold(arg), gchalk.BrightRed("removed"))
	}
}

// AddCluster() : add a cluster and members
func AddCluster(args []string) {
	creds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", creds.DbUsr, creds.DbPasswd, creds.Hostname, creds.Port)
	ctx := context.Background()

	dbconn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer dbconn.Close(ctx)

	for _, arg := range args {
		if memberIsValid(arg) {

		}
		sqlQuery := fmt.Sprintf("DELETE FROM clusters WHERE cname='%s';", arg)
		_, err = dbconn.Exec(ctx, sqlQuery)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("Cluster %s has been removed\n\n\n", arg)
		fmt.Printf("Cluster %s has been %s from the database.\n", gchalk.WithBrightWhite().Bold(arg), gchalk.BrightRed("removed"))
	}
}
