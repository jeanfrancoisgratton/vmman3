// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// db/db-import.go
// 2022-08-24 22:20:39

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Import() : injecte un JSON/YAML dans la BD. LA TABLE SE DOIT D'ÊTRE VIDE. Hard-requirement
func Import(directory string) {
	creds := json2creds()
	var hypervisors []dbHypervisors
	var storagePools []dbStoragePools
	var vmStates []dbVmStates

	//ctx := context.Background()

	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	if Byaml {
		hypervisors, storagePools, vmStates = getYamlTables(directory)
	} else {
		hypervisors, storagePools, vmStates = getJsonTables(directory)
	}

	structs2DB(conn, hypervisors, storagePools, vmStates)
}

// https://stackoverflow.com/questions/59406919/read-and-write-yaml-files-with-go
// https://zetcode.com/golang/yaml/
// --> https://kenanbek.medium.com/golang-how-to-parse-yaml-file-31b78141bda7  <--
// getYamlTables() : Collecte les données en format YAML
func getYamlTables(directory string) (hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates) {
	if !checkNOENT(directory, "hypervisors.yaml") && !checkNOENT(directory, "hypervisors.yml") {
		os.Exit(1)
	}
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return nil, nil, nil
}

// getJsonTables() : Collecte les données en format JSON
func getJsonTables(directory string) (hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates) {
	if !checkNOENT(directory, "hypervisors.json") {
		os.Exit(1)
	}

	return nil, nil, nil
}

// structs2DB() : Injecte les structures dans la BD
func structs2DB(conn *pgx.Conn, hyps []dbHypervisors, sps []dbStoragePools, vms []dbVmStates) {

}

// checkNOENT() : Vérifie si le fichier existe, les perms sont OK, ou autre
func checkNOENT(directory string, file string) bool {
	var fullpath string
	bExists := true

	if directory[:len(directory)-1] == "/" {
		fullpath = fmt.Sprintf("%s%s", directory, file)
	} else {
		fullpath = fmt.Sprintf("%s/%s", directory, file)
	}
	_, err := os.Stat(fullpath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s either does not exist or has permission issues. Aborting.\n", fullpath)
			bExists = false
		} else {
			fmt.Printf("Unhandled error with file %s :\n%s.\nAborting.\n", fullpath, err)
			bExists = false
		}
	}

	return bExists
}
