// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// poolCmd represents the pool command
var poolCmd = &cobra.Command{
	Use:   "pool",
	Short: "Storage pool management",
	Long:  `Commands to manage the storage pool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pool called")
	},
}

func init() {
	rootCmd.AddCommand(poolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// poolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// poolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
