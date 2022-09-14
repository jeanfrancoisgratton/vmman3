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

func interface2struct(hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates, vmc []dbClusters) ([]interface{}, []interface{}, []interface{}, []interface{}) {
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
	dbC := make([]interface{}, len(vmc))
	for i, v := range vmc {
		dbC[i] = v
	}

	return dbH, dbSP, dbVMs, dbC
}
