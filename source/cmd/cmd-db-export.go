// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/cmd/cmd-db-export.go
// 2022-08-26 17:17:33

package cmd

import (
	"github.com/spf13/cobra"
	"vmman3/db"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Export()
	},
}

func init() {
	dbCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	exportCmd.Flags().BoolVarP(&db.Bjson, "jsonformat", "j", true, "Export in JSON format (default)")
	exportCmd.Flags().BoolVarP(&db.Byaml, "yamlformat", "y", false, "Export in YAML format")
	exportCmd.Flags().BoolVarP(&db.Bsql, "sqlformat", "s", false, "Export in SQL format")
}
