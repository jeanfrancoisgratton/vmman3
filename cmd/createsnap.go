// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// createsnapCmd represents the createsnap command
var createsnapCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a snapshot for the VM",
	Long:  `This will create a snapshot for the named VM.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vmman3 snap ls called")
	},
}

func init() {
}
