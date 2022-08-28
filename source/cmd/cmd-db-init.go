// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/cmd/cmd-db-dbInit.go
// 2022-08-25 00:01:18

package cmd

import (
	"github.com/spf13/cobra"
	"vmman3/db"
)

// dbInitCmd represents the dbInit command
var dbInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the db with a data file",
	Long: `The database created with db bootstrap will now be populated with data that comes from a json, yaml or sql file.
Be aware that the software assumes that the file is syntaxically correct.

Also, this subcommand will ignore all subcommand argument except the last, which will be treated as a filename`,
	Run: func(cmd *cobra.Command, args []string) {
		checkFileType(args)
	},
}

func checkFileType(args []string) {
	filename := args[len(args)-1]
	nLen := len(filename)

	// input is YAML
	if filename[nLen-3:] == "yml" || filename[nLen-4:] == "yaml" {
		db.InitYaml()
	}
	// input is JSON
	if filename[nLen-4:] == "json" {
		db.InitJson()
	}
	// input is SQL
	if filename[nLen-3:] == "sql" {
		db.InitSql()
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
}
