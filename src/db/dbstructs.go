// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/db-utils.go
// 2022-08-25 13:32:28

package db

// La structure utilisée pour créer la bd originale
type DbCredsStruct struct {
	Hostname   string `json:"hostname" yaml:"hostname"`
	Port       int    `json:"port" yaml:"port"`
	RootUsr    string `json:"rootusr" yaml:"rootusr"`
	RootPasswd string `json:"rootpasswd" yaml:"rootpasswd"`
	DbUsr      string `json:"dbusr" yaml:"dbusr"`
	DbPasswd   string `json:"dbpasswd" yaml:"dbpasswd"`
}

// table: config.hypervisors
type DbHypervisors struct {
	HID             uint8  `json:"id" yaml:"id"`
	Hname           string `json:"name" yaml:"name"`
	Haddress        string `json:"address" yaml:"address"`
	Hconnectinguser string `json:"connectinguser" yaml:"connectinguser"`
}

// table: config.storagepools
type dbStoragePools struct {
	SpID    uint8  `json:"id" yaml:"id"`
	SpName  string `json:"name" yaml:"name"`
	SpPath  string `json:"path" yaml:"path"`
	SpOwner string `json:"owner" yaml:"owner"`
}

// table: config.vmstates
type dbVmStates struct {
	VmID              uint8  `json:"id" yaml:"id"`
	VmName            string `json:"name" yaml:"name"`
	VmIP              string `json:"ip" yaml:"ip"`
	VmOnline          bool   `json:"online" yaml:"online"`
	VmLastStateChange string `json:"laststatechange" yaml:"laststatechange"`
	VmOperatingSystem string `json:"os" yaml:"vmos"`
	VmLastHypervisor  string `json:"lasthypervisor" yaml:"lasthypervisor"`
	VmStoragePool     string `json:"storagepool" yaml:"storagepool"`
}

// table: config.clusters
type dbClusters struct {
	CID   uint8  `json:"id" yaml:"id"`
	Cname string `json:"name" yaml:"name"`
}

// table: config.templates
type dbTemplates struct {
	TID          uint8  `json:"id" yaml:"id"`
	Tname        string `json:"name" yaml:"name"`
	Towner       string `json:"owner" yaml:"owner"`
	TstoragePool string `json:"storagepool" yaml:"storagepool"`
}

// structure d'info sur les tables
type tableInfo struct {
	tablename     string
	datastructure []interface{}
}
