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
VERSION		DATE			COMMENT
-------		----			-------
0.500		2022.10.19		vmman up down reset completed
0.400		2022.10.16		ls completed+fixed, all commands now fully hypervisor-aware
0.300		2022.10.03		ls is completed, stop[All] near-completed
0.250		2022.09.20		db package extra work; reworked specfile (RPM)
0.200		2022.09.17		db package completed. Build dry run
0.150		2022.09.11		most db- packages are completed, except import
0.100		2022.08.24		db-bootstrap, hypervisor-aware
0.000		2022.06.18		initial version
`)
	fmt.Println()
}
