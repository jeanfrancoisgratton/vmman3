// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/db-utils.go
// 2022-08-25 13:32:28

package db

import (
	"encoding/json"
	"fmt"
	"os"
	"vmman3/helpers"
)

// creds2json() : sérialise la structure dbCredsStruct dans un fichier JSON
func creds2json(jsonFile string, creds DbCredsStruct) {
	jStream, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		fmt.Println("Error", err)
	}
	os.WriteFile(jsonFile, jStream, 0600)
}

func Json2creds() DbCredsStruct {
	var payload DbCredsStruct
	rcFile := helpers.GetRCdir() + helpers.EnvironmentFile
	jFile, _ := os.ReadFile(rcFile)
	err := json.Unmarshal(jFile, &payload)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return payload
}

func createDumpDir(filename string) {
	_, err := os.Stat(filename)

	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(filename, 0755)
		} else {
			panic(err)
		}
	}
	os.Chdir(filename)
}
