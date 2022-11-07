// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/snap.go
// 2022-11-06 11:44:19

package cmd

import (
	"github.com/spf13/cobra"
)

// snapCmd represents the snap command
var snapCmd = &cobra.Command{
	Use:     "snap",
	Aliases: []string{"snapshot"},
	Short:   "Snapshot management subcommand",
	Long: `This is where you manage all snapshots on a given VM.

Subcommands include : list, add, remove`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(snapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
