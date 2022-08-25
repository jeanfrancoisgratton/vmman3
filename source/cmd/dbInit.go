// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/cmd/dbInit.go
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
Be aware that the software assumes that the file is syntaxically correct.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Init()
	},
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
