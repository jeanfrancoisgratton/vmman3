// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/helpers/structs.go
// 2022-11-11 22:33:11

package helpers

var BAllHypervisors = true
var EnvironmentFile string
var ConnectURI string
var BkeepStorage = false
var HypervisorUser = "root"

// La structure utilisée pour créer la bd originale
type EnvironmentStruct struct {
	Hostname              string `json:"hostname" yaml:"hostname"`
	Port                  int    `json:"port" yaml:"port"`
	RootUsr               string `json:"rootusr" yaml:"rootusr"`
	RootPasswd            string `json:"rootpasswd" yaml:"rootpasswd"`
	DbUsr                 string `json:"dbusr" yaml:"dbusr"`
	DbPasswd              string `json:"dbpasswd" yaml:"dbpasswd"`
	HypervisorDefaultUser string `json:"defaulthypervisoruser"`
}
