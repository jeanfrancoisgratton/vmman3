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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
