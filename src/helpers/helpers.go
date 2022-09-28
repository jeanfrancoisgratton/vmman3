// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// helpers/inventory-vm_man-lsHelpers.go
// 2022-08-16 17:50:17

package helpers

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/signal"
	"syscall"
)

// var BsingleHypervisor bool
var BSingleHypervisor = false
var BAllHypervisors = false
var EnvironmentFile string
var ConnectURI string // qemu:///system

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

// GetRCdir() : retourne le répertoire de configurations de l'usager
func GetRCdir() string {
	rcDir, _ := os.UserConfigDir()

	return rcDir + "/vmman3/"
}

// BuildPath() : une fonction pour construire le full pathname d'un fichier
func BuildPath(directory, file string) string {
	var fullpath string
	if directory[:len(directory)-1] == "/" {
		fullpath = fmt.Sprintf("%s%s", directory, file)
	} else {
		fullpath = fmt.Sprintf("%s/%s", directory, file)
	}
	return fullpath
}

// reference: https://gist.github.com/jlinoff/e8e26b4ffa38d379c7f1891fd174a6d0, the getPassword2.go
func GetPassword(prompt string) string {
	// Get the initial state of the terminal.
	initialTermState, e1 := terminal.GetState(syscall.Stdin)
	if e1 != nil {
		panic(e1)
	}

	// Restore it in the event of an interrupt.
	// CITATION: Konstantin Shaposhnikov - https://groups.google.com/forum/#!topic/golang-nuts/kTVAbtee9UA
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		_ = terminal.Restore(syscall.Stdin, initialTermState)
		os.Exit(1)
	}()

	// Now get the password.
	fmt.Print(prompt)
	p, err := terminal.ReadPassword(syscall.Stdin)
	fmt.Println("")
	if err != nil {
		panic(err)
	}

	// Stop looking for ^C on the channel.
	signal.Stop(c)

	// Return the password as a string.
	return string(p)
}
