// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// cmd/dbBootstrap.go
// 2022-08-22 07:08:44

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// dbBootstrapCmd represents the dbBootstrap command
var dbBootstrapCmd = &cobra.Command{
	Use:   "db bootstrap",
	Short: "Bootstraps (initializes) the PGSQL back-end for multi-hypervisors environments",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dbBootstrap called")
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
	// dbBootstrapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
