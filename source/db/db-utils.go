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

var Bjson, Byaml, Bsql bool

// La structure utilisée pour créer la bd originale
type dbCredsStruct struct {
	Hostname   string `json:"hostname"`
	Port       int    `json:"port"`
	RootUsr    string `json:"rootusr"`
	RootPasswd string `json:"rootpasswd"`
	DbUsr      string `json:"dbusr"`
	DbPasswd   string `json:dbpasswd`
}

// table: config.hypervisors
type dbHypervisors struct {
	HID      uint8  `json:"hid" yaml:"hid"`
	Hname    string `json:"hname" yaml:"hname"`
	Haddress string `json:"haddress" yaml:"address"`
}

//type dbHypervisorSlice []dbHypervisors

// table: config.storagepools
type dbStoragePools struct {
	SpID    uint8  `json:"spid" yaml:"spid"`
	SpName  string `json:"spname" yaml:"spname"`
	SpPath  string `json:"sppath" yaml:"sppath"`
	SpOwner string `json:"spowner,omitempty" yaml:"spowner,omitempty"`
}

//type dbStoragePoolSlice []dbStoragePools

// table: config.vmstate
type dbVmStates struct {
	VmID              uint8  `json:"vmid" yaml:"vmid"`
	VmName            string `json:"vmname" yaml:"vmname"`
	VmIP              string `json:"vmip,omitempty" yaml:"vmip,omitempty"`
	VmOnline          bool   `json:"online" yaml:"online"`
	VmLastStateChange string `json:"laststatechange" yaml:"laststatechange"`
}

// type dbVmStateSlice []dbVmStates
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
