// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/helpers/dbHelpers.go
// 2022-09-16 17:42:54

package helpers

import (
	"fmt"
	"os"
)

// CheckNOENT() : Vérifie si le fichier existe, les perms sont OK, ou autre
func CheckNOENT(directory string, file string) bool {
	fullpath := BuildPath(directory, file)
	bExists := true

	_, err := os.Stat(fullpath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s either does not exist or has permission issues. Aborting.\n", fullpath)
			bExists = false
		} else {
			fmt.Printf("Unhandled error with file %s :\n%s.\nAborting.\n", fullpath, err)
			bExists = false
		}
	}
	return bExists
}

// checkIfConfigExists() : Vérifie si le répertoire existe; s'il existe, vérifie si le fichier de config existe
// s'il existe, on l'efface, il sera écrasé plus tard
func CheckIfConfigExists() string {
	vmman3rcdir, _ := os.UserHomeDir()
	vmman3rcdir += "/.config/vmman3"

	_, err := os.Stat(vmman3rcdir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(vmman3rcdir, 0700)
		} else {
			panic(err)
		}
	}
	vmman3rcdir += "/databaseCreds.json"

	_, err = os.Stat(vmman3rcdir)
	if err != nil {
		if os.IsNotExist(err) {
			return vmman3rcdir
		} else {
			panic(err)
		}
	} else {
		os.Remove(vmman3rcdir)
	}
	return vmman3rcdir
}
