/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// pooladdCmd represents the pooladd command
var pooladdCmd = &cobra.Command{
	Use:   "add",
	Short: "TODO: Add a new pool storage",
	Long:  `TODO: This will add a new pool storage both in QEMU (libvirt) and the databse.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pool add NOT IMPLEMENTED")
		//storagepool.PoolAdd(args[0], args[1])
	},
}

func init() {
	poolCmd.AddCommand(pooladdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pooladdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pooladdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
