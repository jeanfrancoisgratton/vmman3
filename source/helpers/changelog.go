// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/helpers/changelog.go
// 2022-09-11 23:23:12

package helpers

import "fmt"

// Changelog() :
// Affiche simplement le changelog (le fichier _version)
func Changelog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Print(`
VERSION     DATE            COMMENT
-------     ----            -------
0.150		2022.09.11		most db- modules are completed, except import
0.100       2022.08.24      db-bootstrap, hypervisor-aware
0.000       2022.06.18      Initial version
`)
	fmt.Println()
}
