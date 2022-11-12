// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/snap-revert.go
// 2022-11-11 18:17:20

package cmd

import (
	"fmt"
	"os"
	"vmman3/snapshotmanagement"

	"github.com/spf13/cobra"
)

// revertCmd represents the revert command
var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Reverts the VM to a snapshot",
	Long: `This subcommand allows you to revert the VM to a specific snapshot.
If no snapshotname is specified, it reverts to the current one.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			snapshotmanagement.RevertSnapshot(args[0], args[1])
			os.Exit(0)
		} else if len(args) == 1 {
			snapshotmanagement.RevertSnapshot(args[0], "")
			os.Exit(0)
		} else {
			fmt.Println("USAGE: vmman {snap|snapshot} revert VMNAME [SNAPSHOTNAME]")
			os.Exit(0)
		}
	},
}

func init() {
	snapCmd.AddCommand(revertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// revertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// revertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
