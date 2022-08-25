// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rmsnapCmd represents the rmsnap command
var rmsnapCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove snapshot",
	Long: `Removes the current snapshot from the VM.

NOTE: The snapshot cannot be removed if a child snapshot is present.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vmman3 snap rm called")
	},
}

func init() {
}
