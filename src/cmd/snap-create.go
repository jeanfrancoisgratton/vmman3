// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"github.com/spf13/cobra"
	"vmman3/snapshotmanagement"
)

// createsnapCmd represents the createsnap command
var createsnapCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a snapshot for the VM",
	Long:  `This will create a snapshot for the named VM.`,
	Run: func(cmd *cobra.Command, args []string) {
		snapshotmanagement.CreateSnapshot(args[0], args[1])
	},
}

func init() {
}
