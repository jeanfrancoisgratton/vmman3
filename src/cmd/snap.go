// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// snapCmd represents the snap command
var snapCmd = &cobra.Command{
	Use:   "snap",
	Short: "Snapshot management command",
	Long: `Snapshot management:

List, add or remove a snapshot.`,
	// TODO : replace this command below with a snap help text
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("A subcommand { create | ls | rm } must be passed to the snap command")
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
