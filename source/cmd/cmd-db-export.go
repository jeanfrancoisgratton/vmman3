// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/cmd/cmd-db-db-export.go
// 2022-08-26 17:17:33

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/db"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Dump database in selected format",
	Long: `This will dump the DB tables in the format you select (default is JSON).

All tables will be dumped in JSON unless you toggle it off with -j=false.
You will need to specify the full path and where to dump the files.

Please note: You HAVE to specify a target directory. If it does not exist, it will be created.`,
	Run: func(cmd *cobra.Command, args []string) {
		runExport(args)
	},
}

func runExport(args []string) {
	nLen := len(args)
	if nLen != 1 {
		fmt.Println("You need to specify a target directory")
		os.Exit(-1)
	} else {
		db.Export(args[nLen-1])
	}
}

func init() {
	dbCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	exportCmd.Flags().BoolVarP(&db.Bjson, "jsonfmt", "j", true, "Export in JSON format (default)")
	exportCmd.Flags().BoolVarP(&db.Byaml, "yamlfmt", "y", false, "Export in YAML format")
	exportCmd.Flags().BoolVarP(&db.Bsql, "sqlfmt", "s", false, "Export in SQL format")
}
