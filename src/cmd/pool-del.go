/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pooldelCmd represents the pooldel command
var pooldelCmd = &cobra.Command{
	Use:     "r,",
	Aliases: []string{"del"},
	Short:   "TODO: Deletes a pool storage",
	Long:    `TODO: This will remove a pool storage both in QEMU (libvirt) and the databse.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pool rm NOT IMPLEMENTED")
		//storagepool.PoolAdd(args[0])
	},
}

func init() {
	poolCmd.AddCommand(pooldelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pooldelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pooldelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
