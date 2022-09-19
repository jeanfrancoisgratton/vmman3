// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/inventory/lsStructs.go
// 2022-09-17 17:52:18

package inventory

type vmInfo struct {
	viId                                                uint
	viName, viState                                     string
	viMem                                               uint64
	viCpu, viSnapshot                                   uint
	viCurrentSnapshot, viInterfaceName, viIPaddress     string
	viLastStatusChange, viHypervisor, viOperatingSystem string
}
