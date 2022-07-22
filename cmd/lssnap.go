/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lssnapCmd represents the lssnap command
var lssnapCmd = &cobra.Command{
	Use:   "lssnap",
	Short: "List all snapshots of a given VM",
	Long:  `Lists all snapshots of a given MV, and also lists the parents and current snapshot.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lssnap called")
	},
}

func init() {
	rootCmd.AddCommand(lssnapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lssnapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lssnapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
