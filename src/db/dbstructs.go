// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/db-utils.go
// 2022-08-25 13:32:28

package db

// table: hypervisors
type DbHypervisors struct {
	HID             uint8  `json:"id" yaml:"id"`
	Hname           string `json:"name" yaml:"name"`
	Haddress        string `json:"address" yaml:"address"`
	Hconnectinguser string `json:"connectinguser" yaml:"connectinguser"`
}

// table: storagepools
type DbStoragePools struct {
	SpID    uint8  `json:"id" yaml:"id"`
	SpName  string `json:"name" yaml:"name"`
	SpPath  string `json:"path" yaml:"path"`
	SpOwner string `json:"owner" yaml:"owner"`
}

// table: vmstates
type dbVmStates struct {
	VmID              uint8  `json:"id" yaml:"id"`
	VmName            string `json:"name" yaml:"name"`
	VmIP              string `json:"ip" yaml:"ip"`
	VmOnline          bool   `json:"online" yaml:"online"`
	VmLastStateChange string `json:"laststatechange" yaml:"laststatechange"`
	VmOperatingSystem string `json:"os" yaml:"vmos"`
	VmHypervisor      string `json:"hypervisor" yaml:"hypervisor"`
	VmStoragePool     string `json:"storagepool" yaml:"storagepool"`
}

//// table: clusters
//type dbClusters struct {
//	CID   uint8  `json:"id" yaml:"id"`
//	Cname string `json:"name" yaml:"name"`
//}

// table: templates
type dbTemplates struct {
	TID              uint8  `json:"id" yaml:"id"`
	Tname            string `json:"name" yaml:"name"`
	Towner           string `json:"owner" yaml:"owner"`
	TstoragePool     string `json:"storagepool" yaml:"storagepool"`
	ToperatingSystem string `json:"os" yaml:"os"`
}

type dbDisks struct {
	DID         uint   `json:"id" yaml:"id"`
	Dname       string `json:"name" yaml:"name"`
	Dpool       string `json:"pool" yaml:"pool"`
	Dvm         string `json:"vm" yaml:"vm"`
	Dhypervisor string `json:"hypervisor" yaml:"hypervisor"`
}

//// structure d'info sur les tables
//type tableInfo struct {
//	tablename     string
//	datastructure []interface{}
//}
