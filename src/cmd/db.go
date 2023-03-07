// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/db"
	"vmman3/helpers"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Config and info DB management",
	Long:  `This is where you configure the backend.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("A subcommand { bootstrap | init | drop } must be passed to the db command")
	},
}

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

// dropCmd represents the drop command
var dropCmd = &cobra.Command{
	Use:     "drop",
	Aliases: []string{"wipe"},
	Short:   "Drops the database",
	Long: `The vmman database will be dropped (erased).

WARNING : This is irreversible !`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Drop()
	},
}

// exportCmd represents the export command
var dbExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Dump database in selected format",
	Long: `This will dump the DB tables in JSON format

You will need to specify the directory where to dump the files.`,
	Run: func(cmd *cobra.Command, args []string) {
		nLen := len(args)
		if nLen != 1 {
			fmt.Println("You need to specify a target directory")
			os.Exit(-1)
		} else {
			db.Export(args[nLen-1])
		}
	},
}

// dbInitCmd represents the dbInit command
var dbImportCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"init"},
	Short:   "Initializes the db with a data file",
	Long: `The database created with db bootstrap will now be populated with data that comes from a json, yaml or sql file.
Be aware that the software assumes that the file is syntaxically correct.

Also, this subcommand will ignore all subcommand argument except the last, which will be treated as a dirname`,
	Run: func(cmd *cobra.Command, args []string) {
		nLen := len(args)
		if nLen != 1 {
			fmt.Println("You need to specify a target directory")
			os.Exit(-1)
		} else {
			db.Import(args[nLen-1])
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbBootstrapCmd)
	dbCmd.AddCommand(dropCmd)
	dbCmd.AddCommand(dbExportCmd)
	dbCmd.AddCommand(dbImportCmd)

	dropCmd.PersistentFlags().BoolVarP(&helpers.DbDropAssumeYes, "yes", "y", false, "Assume yes to all questions")
}
