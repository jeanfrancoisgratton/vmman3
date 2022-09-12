// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/db-utils.go
// 2022-08-25 13:32:28

package db

// La structure utilisée pour créer la bd originale
type dbCredsStruct struct {
	Hostname   string `json:"hostname",yaml:"hostname"`
	Port       int    `json:"port",yaml:"port"`
	RootUsr    string `json:"rootusr",yaml:"rootusr"`
	RootPasswd string `json:"rootpasswd",yaml:"rootpasswd"`
	DbUsr      string `json:"dbusr",yaml:"dbusr"`
	DbPasswd   string `json:"dbpasswd",yaml:"dbpasswd"`
}

// table: config.hypervisors
type dbHypervisors struct {
	HID      uint8  `json:"hid" yaml:"hid"`
	Hname    string `json:"hname" yaml:"hname"`
	Haddress string `json:"haddress" yaml:"address"`
}

// table: config.storagepools
type dbStoragePools struct {
	SpID    uint8  `json:"spid" yaml:"spid"`
	SpName  string `json:"spname" yaml:"spname"`
	SpPath  string `json:"sppath" yaml:"sppath"`
	SpOwner string `json:"spowner,omitempty" yaml:"spowner,omitempty"`
}

// table: config.vmstate
type dbVmStates struct {
	VmID              uint8  `json:"vmid" yaml:"vmid"`
	VmName            string `json:"vmname" yaml:"vmname"`
	VmIP              string `json:"vmip,omitempty" yaml:"vmip,omitempty"`
	VmOnline          bool   `json:"online" yaml:"online"`
	VmLastStateChange string `json:"laststatechange" yaml:"laststatechange"`
}

// table: config.clusters
type dbClusters struct {
	CID   uint8  `json:"cid",yaml:"cid"`
	Cname string `json:"cname",yaml:"cname"`
}

// table: config.servers
type dbServers struct {
	Sid              uint8  `json:"sid",yaml:"sid"`
	Sname            string `json:"sname",yaml:"sname"`
	SoperatingSystem string `json:"sos",yaml:"sos"`
	SlastHypervisor  string `json:"slasthypervisor",yaml:"slasthypervisor"`
}
