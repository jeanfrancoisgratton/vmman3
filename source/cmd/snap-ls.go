// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// lssnapCmd represents the lssnap command
var lssnapCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all snapshots",
	Long: `This command will list all snapshots of a given VM:

The complete list will also inform you of the links (parent/child) between
the snapshots.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("snap ls called")
	},
}

func init() {
}
