// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/hypervisor/hypervisor.go
// 2022-11-11 21:29:08

package hypervisor

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

// ListHypervisors() : lists all hypervisors
func ListHypervisors() {
	var hypervisor db.DbHypervisors
	var hypervisors []db.DbHypervisors

	envCreds := helpers.Json2creds()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", envCreds.DbUsr, envCreds.DbPasswd, envCreds.Hostname, envCreds.Port)

	dbconn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close(context.Background())

	rows, err := dbconn.Query(context.Background(), "SELECT * from hypervisors ORDER BY hid")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		retcode := rows.Scan(&hypervisor.HID, &hypervisor.Hname, &hypervisor.Haddress, &hypervisor.Hconnectinguser)
		if retcode != nil {
			fmt.Println("Error:", retcode)
			os.Exit(-9)
		}
		hypervisors = append(hypervisors, hypervisor)
	}
	helpers.SurroundText("Registered hypervisors", false)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Address", "Connecting user"})

	for _, hypervisor := range hypervisors {
		t.AppendRow([]interface{}{hypervisor.HID, hypervisor.Hname, hypervisor.Haddress, hypervisor.Hconnectinguser})
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

func AddtoCluster(member []string) {

}
