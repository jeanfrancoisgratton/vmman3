// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/structures.go
// 2022-08-26 17:38:31

package db

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

//type dbVmStateSlice []dbVmStates
