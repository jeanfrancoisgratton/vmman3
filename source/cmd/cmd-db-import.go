// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/cmd/cmd-db-dbInit.go
// 2022-08-25 00:01:18

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/db"
)

// dbInitCmd represents the dbInit command
var dbInitCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"init"},
	Short:   "Initializes the db with a data file",
	Long: `The database created with db bootstrap will now be populated with data that comes from a json, yaml or sql file.
Be aware that the software assumes that the file is syntaxically correct.

Also, this subcommand will ignore all subcommand argument except the last, which will be treated as a dirname`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Import(args[len(args)-1])
	},
}

func runImport(args []string) {
	nLen := len(args)
	if nLen != 1 {
		fmt.Println("You need to specify a target directory")
		os.Exit(-1)
	} else {
		db.Import(args[nLen-1])
	}
}

func init() {
	dbCmd.AddCommand(dbInitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbInitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbInitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	dbInitCmd.Flags().BoolVarP(&db.Bjson, "jsonfmt", "j", true, "Export in JSON format (default)")
	dbInitCmd.Flags().BoolVarP(&db.Byaml, "yamlfmt", "y", false, "Export in YAML format")
	dbInitCmd.Flags().BoolVarP(&db.Bsql, "sqlfmt", "s", false, "Export in SQL format")
}
