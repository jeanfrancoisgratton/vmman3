// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Config and info DB management",
	Long:  `This is where you configure the backend.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("A subcommand { bootstrap | init | drop } must be passed to the db command")
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
