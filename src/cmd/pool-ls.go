/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"vmman3/storagepool"

	"github.com/spf13/cobra"
)

// TODO: add capability to list pools from QEMU

// poollistCmd represents the poollist command
var poollistCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all pools currently configured in the DB",
	Long:  `If a pool has been configured but has not been inserted in the database it will not show up.`,
	Run: func(cmd *cobra.Command, args []string) {
		storagepool.PoolList()
	},
}

func init() {
	poolCmd.AddCommand(poollistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// poollistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// poollistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
