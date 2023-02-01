// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/helpers/dbHelpers.go
// 2022-09-16 17:42:54

package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"strings"
)

// BuildConnectURI() : Builds a PGSQL connection string from the ConnectURI string
func BuildConnectURI(host string) string {
	var username string
	ctx := context.Background()
	creds := Json2creds()

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/vmman", creds.DbUsr, creds.DbPasswd, creds.Hostname, creds.Port)

	dbconn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer dbconn.Close(ctx)

	err = dbconn.QueryRow(ctx, "SELECT hconnectinguser FROM hypervisors WHERE hname='"+host+"';").Scan(&username)
	if err != nil {
		panic(err)
	}
	return username
}

// SplitConnectURI() : Extracts the username & host from the ConnectURI string
func SplitConnectURI(uri string) (string, string, string) {
	protoStr := strings.SplitAfter(uri, "://")
	atNdx := strings.Index(protoStr[1], "@")
	slashNdx := strings.Index(protoStr[1], "/")
	user := protoStr[1][0:atNdx]
	hypervisor := protoStr[1][atNdx+1 : slashNdx]

	return protoStr[0], user, hypervisor
}

// creds2json() : sérialise la structure EnvironmentStruct dans un fichier JSON
func Creds2json(jsonFile string, creds EnvironmentStruct) {
	jStream, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		fmt.Println("Error", err)
	}
	os.WriteFile(jsonFile, jStream, 0600)
}

func Json2creds() EnvironmentStruct {
	var payload EnvironmentStruct
	rcDir, _ := os.UserHomeDir()
	rcFile := rcDir + "/.config/vmman3/" + EnvironmentFile
	jFile, _ := os.ReadFile(rcFile)
	err := json.Unmarshal(jFile, &payload)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return payload
}
