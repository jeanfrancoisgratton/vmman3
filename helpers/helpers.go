// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// helpers/inventory-vm_man-helpers.go
// 2022-08-16 17:50:17

package helpers

import (
	"fmt"
	"log"
	"os"
)

var BsingleHypervisor bool

// Changelog() :
// Affiche simplement le changelog (le fichier _version)
func Changelog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	content, err := os.ReadFile("_version")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

// SurroundText()
// Fonction stupide pour afficher du texte "proprement" (avec un header-footer)
func SurroundText(text string, clearScr bool) {
	if clearScr == true {
		fmt.Println("\x1bc")
	}

	txLen := len(text)
	i := 0
	eq := ""

	for i < txLen {
		eq += "="
		i += 1
	}

	fmt.Println(eq)
	fmt.Println(text)
	fmt.Println(eq)
	fmt.Println()
}
