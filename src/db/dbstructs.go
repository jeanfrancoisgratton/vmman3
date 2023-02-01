// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/db-utils.go
// 2022-08-25 13:32:28

package db

// table: hypervisors
type DbHypervisors struct {
	HID             uint8  `json:"id" yaml:"id"`
	Hname           string `json:"hypervisor name" yaml:"hypervisor name"`
	Haddress        string `json:"address" yaml:"address"`
	Hconnectinguser string `json:"connecting user" yaml:"connecting user"`
}

// table: storagepools
type DbStoragePools struct {
	SpID    uint8  `json:"id" yaml:"id"`
	SpName  string `json:"storage pool name" yaml:"storage pool name"`
	SpPath  string `json:"path" yaml:"path"`
	SpOwner string `json:"owner" yaml:"owner"`
}

// table: vmstates
type dbVmStates struct {
	VmID              uint8  `json:"id" yaml:"id"`
	VmName            string `json:"vm name" yaml:"vm name"`
	VmIP              string `json:"ip" yaml:"ip"`
	VmOnline          bool   `json:"online" yaml:"online"`
	VmLastStateChange string `json:"last state change" yaml:"last state change"`
	VmOperatingSystem string `json:"os" yaml:"vmos"`
	VmHypervisor      string `json:"hypervisor" yaml:"hypervisor"`
	VmStoragePool     string `json:"storage pool" yaml:"storage pool"`
}

// table: clusters
type DbClusters struct {
	CID            uint8  `json:"id" yaml:"id"`
	Cname          string `json:"cluster name" yaml:"cluster name"`
	Cclustermember string `json:"cluster member" yaml:"cluster member"`
}

// table: templates
type dbTemplates struct {
	TID              uint8  `json:"id" yaml:"id"`
	Tname            string `json:"template name" yaml:"template name"`
	Towner           string `json:"owner" yaml:"owner"`
	TstoragePool     string `json:"storage pool" yaml:"storage pool"`
	ToperatingSystem string `json:"os" yaml:"os"`
}

type dbDisks struct {
	DID         uint   `json:"id" yaml:"id"`
	Dname       string `json:"disk name" yaml:"disk name"`
	Dpool       string `json:"storage pool" yaml:"storage pool"`
	Dvm         string `json:"vm" yaml:"vm"`
	Dhypervisor string `json:"hypervisor" yaml:"hypervisor"`
}

//// structure d'info sur les tables
//type tableInfo struct {
//	tablename     string
//	datastructure []interface{}
//}
