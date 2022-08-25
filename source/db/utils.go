// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/utils.go
// 2022-08-25 13:32:28

package db

import (
	"encoding/json"
	"fmt"
	"os"
	"vmman3/helpers"
)

// La structure utilisée pour créer la bd originale
type dbCredsStruct struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

// creds2json() : sérialise la structure dbCredsStruct dans un fichier JSON
func creds2json(jsonFile string, creds dbCredsStruct) {
	jStream, err := json.Marshal(creds)
	if err != nil {
		fmt.Println("Error", err)
	}
	os.WriteFile(jsonFile, jStream, 0600)
}

func json2creds() dbCredsStruct {
	var payload dbCredsStruct
	rcFile := helpers.GetRCdir() + "database.json"
	jFile, _ := os.ReadFile(rcFile)
	err := json.Unmarshal(jFile, &payload)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return payload
}
