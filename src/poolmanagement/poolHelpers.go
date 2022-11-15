// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/poolmanagement/poolHelpers.go
// 2022-11-13 16:18:34

package poolmanagement

import (
	"fmt"
	"libvirt.org/go/libvirt"
	"os"
	"vmman3/helpers"
)

// getPoolStatus() : Get Pool info from QEMU, instead of LibVirtd
// WARNING: THIS IS A HUGE KLUDGE RIGHT NOW, FOR SOME REASON...
func getPoolStatus(poolname string) (autostart, active string) {
	conn := helpers.Connect2HVM()
	var aactive, astart bool
	pool, err := conn.LookupStoragePoolByName(poolname)
	if err != nil {
		fmt.Printf("Unable to fetch poolmanagement %s info: %s", poolname, err)
		os.Exit(-19)
	}
	aactive, err = pool.IsActive()
	if err != nil {
		lverror, ok := err.(libvirt.Error)
		if ok {
			fmt.Printf("Error: ", lverror.Message)
			active = ""
		}
	} else {
		active = fmt.Sprintf("%t", aactive)
	}
	astart, err = pool.GetAutostart()
	if err != nil {
		lverror, ok := err.(libvirt.Error)
		if ok {
			fmt.Println("Error: ", lverror.Message)
			autostart = ""
		}
	} else {
		autostart = fmt.Sprintf("%t", astart)
	}
	return active, autostart
}
