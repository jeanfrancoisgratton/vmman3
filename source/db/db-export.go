// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/db-export.go
// 2022-08-26 17:18:20

// SEE SCANNY ON https://stackoverflow.com/questions/61704842/how-to-scan-a-queryrow-into-a-struct-with-pgx

package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Export() : point d'entrée de l'exportation
func Export(filename string) {
	creds := json2creds()

	createDumpDir(filename)
	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	hypervisors := getHypervisorData(conn)
	storagePools := getSpData(conn)
	vmStates := getVmStateData(conn)

	if err := serialize(hypervisors, "hypervisors"); err != nil {
		log.Fatalln(err)
	}
	if err := serialize(storagePools, "storagepools"); err != nil {
		log.Fatalln(err)
	}
	if err := serialize(vmStates, "vmtates"); err != nil {
		log.Fatalln(err)
	}
}

// serialize() : Serialise toutes les tables, selon le(s) format(s) choisi(s) (pour le moment: json & yaml, pas encore sql)
func serialize(v interface{}, filename string) error {
	var data []byte
	var f *os.File
	var err error

	defer f.Close()
	if Bjson {
		f, err = os.Create(filename + ".json")
		if err != nil {
			return err
		}
		data, err = json.Marshal(v)
		if err != nil {
			return err
		}
		_, err = f.Write(data)
	}
	if Byaml {
		f, err = os.Create(filename + ".yaml")
		if err != nil {
			return err
		}
		data, err = yaml.Marshal(v)
		if err != nil {
			return err
		}
		_, err = f.Write(data)
	}

	if Bsql {
		//f, err = os.Create(filename + ".sql")
		//if err != nil {
		//	return err
		//}
		fmt.Println("SQL export not yet implemented.")
		return nil
		//if err != nil {
		//	return err
		//}
		//_, err = f.Write(data)
	}

	return err
}

// getHypervisorData(): importe le contenu de la table hypervisors
func getHypervisorData(conn *pgx.Conn) []dbHypervisors {
	var hyps []dbHypervisors

	rows, err := conn.Query(context.Background(), "SELECT hID, hName, hAddress from config.hypervisors")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dbh dbHypervisors
		err := rows.Scan(&dbh.HID, &dbh.Hname, &dbh.Haddress)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			hyps = append(hyps, dbh)
		}
	}
	return hyps
}

// getSPdata() : prend le contenu de la table storagePools
func getSpData(conn *pgx.Conn) []dbStoragePools {
	var sps []dbStoragePools

	rows, err := conn.Query(context.Background(), "SELECT spid, spname, sppath, spowner from config.storagepools")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sp dbStoragePools
		err := rows.Scan(&sp.SpID, &sp.SpName, &sp.SpPath, &sp.SpOwner)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			sps = append(sps, sp)
		}
	}
	return sps
}

// getVmStateData() : prend le contenu de la table vmstate
func getVmStateData(conn *pgx.Conn) []dbVmStates {
	var vmss []dbVmStates

	rows, err := conn.Query(context.Background(), "SELECT vmid, vmname, vmip, vmonline,vmlaststatechange from config.vmstate")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var vms dbVmStates
		err := rows.Scan(&vms.VmID, &vms.VmName, &vms.VmIP, &vms.VmOnline, &vms.VmLastStateChange)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			vmss = append(vmss, vms)
		}
	}
	return vmss
}
