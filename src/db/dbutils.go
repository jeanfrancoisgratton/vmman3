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

// This might get converted to generics, at some point
func interface2struct(hyps []DbHypervisors, sps []dbStoragePools, vms []dbVmStates, vmc []dbClusters) ([]interface{}, []interface{}, []interface{}, []interface{}) {
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
