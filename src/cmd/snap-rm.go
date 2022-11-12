// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/snapshotmanagement"
)

// rmsnapCmd represents the rmsnap command
var rmsnapCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove snapshot",
	Long: `Removes the current snapshot from the VM.

NOTE: The snapshot cannot be removed if a child snapshot is present.`,
	// usage: snap rm VMNAME SNAPNAME
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			snapshotmanagement.RemoveSnapshot(args[0], args[1])
			os.Exit(0)
		} else if len(args) == 1 {
			snapshotmanagement.RemoveSnapshot(args[0], "")
			os.Exit(0)
		} else {
			fmt.Println("USAGE: vmman {snap|snapshot} rm VMNAME [SNAPSHOTNAME]")
			os.Exit(0)
		}
	},
}

func init() {
	snapCmd.AddCommand(rmsnapCmd)
}
