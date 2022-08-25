// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// db/initiator.go
// 2022-08-24 22:21:03

package db

import (
	"fmt"
	"vmman3/helpers"
)

func Init() {
	fmt.Println("Config file: ", helpers.GetRCdir()+"database.json")
	creds := json2creds()

	fmt.Printf("Hostname = %s\nPort = %d\n\n", creds.Hostname, creds.Port)
}
