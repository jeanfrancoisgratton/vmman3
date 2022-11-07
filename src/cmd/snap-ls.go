// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/snap-ls.go
// 2022-11-06 11:46:53

package cmd

import (
	"vmman3/snapshotmanagement"

	"github.com/spf13/cobra"
)

// snapLSCmd represents the snapLS command
var snapLSCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all snapshots of a given VM",
	//	Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		snapshotmanagement.ListSnapshots(args[0])
	},
}

func init() {
	snapCmd.AddCommand(snapLSCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snapLSCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snapLSCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
