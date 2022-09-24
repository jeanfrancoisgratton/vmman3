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
	"log"
	"os"
)

// Export() : point d'entrée de l'exportation
func Export(filename string) {
	creds := Json2creds()

	createDumpDir(filename)
	connString := fmt.Sprintf("postgresql://%s:vmman@%s:%d/vmman", creds.DbUsr, creds.Hostname, creds.Port)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	hypervisors := GetHypervisorData(conn)
	storagePools := getSpData(conn)
	vmStates := getVmStateData(conn)
	clusters := getClusterData(conn)
	templates := getTemplateData(conn)

	if err := serialize(hypervisors, "hypervisors.json"); err != nil {
		log.Fatalln(err)
	}
	if err := serialize(storagePools, "storagepools.json"); err != nil {
		log.Fatalln(err)
	}
	if err := serialize(vmStates, "vmstates.json"); err != nil {
		log.Fatalln(err)
	}
	if err := serialize(clusters, "clusters.json"); err != nil {
		log.Fatalln(err)
	}
	if err := serialize(templates, "templates.json"); err != nil {
		log.Fatalln(err)
	}
}

// serialize() : Serialise toutes les tables, selon le(s) format(s) choisi(s) (pour le moment: json & yaml, pas encore sql)
func serialize(v interface{}, filename string) error {
	var data []byte
	var f *os.File
	var err error

	defer f.Close()
	f, err = os.Create(filename)
	if err != nil {
		return err
	}
	data, err = json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}

// getHypervisorData(): importe le contenu de la table hypervisors
func GetHypervisorData(conn *pgx.Conn) []DbHypervisors {
	var hyps []DbHypervisors

	rows, err := conn.Query(context.Background(), "SELECT * from config.hypervisors ORDER BY hid")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dbh DbHypervisors
		retcode := rows.Scan(&dbh.HID, &dbh.Hname, &dbh.Haddress, &dbh.Hconnectinguser)
		if retcode != nil {
			fmt.Println("Error:", retcode)
		} else {
			hyps = append(hyps, dbh)
		}
	}
	return hyps
}

// getSPdata() : prend le contenu de la table storagePools
func getSpData(conn *pgx.Conn) []dbStoragePools {
	var sps []dbStoragePools

	rows, err := conn.Query(context.Background(), "SELECT * from config.storagepools ORDER BY spid")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sp dbStoragePools
		retcode := rows.Scan(&sp.SpID, &sp.SpName, &sp.SpPath, &sp.SpOwner)
		if retcode != nil {
			fmt.Println("Error:", retcode)
		} else {
			sps = append(sps, sp)
		}
	}
	return sps
}

// getVmStateData() : prend le contenu de la table vmstate
func getVmStateData(conn *pgx.Conn) []dbVmStates {
	var vmss []dbVmStates

	rows, retcode := conn.Query(context.Background(), "SELECT * from config.vmstates ORDER BY vmid;")
	if retcode != nil {
		fmt.Println("Error: ", retcode)
	}
	defer rows.Close()

	for rows.Next() {
		var vms dbVmStates
		retcode := rows.Scan(&vms.VmID, &vms.VmName, &vms.VmIP, &vms.VmOnline, &vms.VmLastStateChange, &vms.VmOperatingSystem, &vms.VmHypervisor, &vms.VmStoragePool)
		if retcode != nil {
			fmt.Println("Error:", retcode)
		} else {
			vmss = append(vmss, vms)
		}
	}
	return vmss
}

// getClusterData() : prend le contenu de la table servers
func getClusterData(conn *pgx.Conn) []dbClusters {
	var clusters []dbClusters

	rows, retcode := conn.Query(context.Background(), "SELECT * from config.clusters ORDER BY cid")
	if retcode != nil {
		fmt.Println("Error: ", retcode)
	}
	defer rows.Close()

	for rows.Next() {
		var cluster dbClusters
		err := rows.Scan(&cluster.CID, &cluster.Cname)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			clusters = append(clusters, cluster)
		}
	}
	return clusters
}

// getTemplateData(): importe le contenu de la table templates
func getTemplateData(conn *pgx.Conn) []dbTemplates {
	var temps []dbTemplates

	rows, err := conn.Query(context.Background(), "SELECT * from config.templates ORDER BY tid")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dbt dbTemplates
		retcode := rows.Scan(&dbt.TID, &dbt.Tname, &dbt.Towner, &dbt.TstoragePool, &dbt.ToperatingSystem)
		if retcode != nil {
			fmt.Println("Error:", retcode)
		} else {
			temps = append(temps, dbt)
		}
	}
	return temps
}
