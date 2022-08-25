// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// cmd/dbBootstrap.go
// 2022-08-22 07:08:44

package cmd

import (
	"github.com/spf13/cobra"
	"vmman3/db"
)

// dbBootstrapCmd represents the dbBootstrap command
var dbBootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstraps (initializes) the PGSQL back-end for multi-hypervisors environments",
	Long: `This will use the user you will provide to create the software's built-in postgres user, db, schema and tables.

You use this command only if :
1- This is the very first time you use the software
2- You want to completely wipe the former DB and start over.

ALL PREVIOUS INFO WILL BE LOST`,
	Run: func(cmd *cobra.Command, args []string) {
		db.CreateDatabase()
	},
}

func init() {
	dbCmd.AddCommand(dbBootstrapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbBootstrapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbBootstrapCmd.Flags().BoolP("unstrap", "u", false, "Unstrap (wipe)")
}
