// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"vmman3/storagepool"
)

// poolCmd represents the pool command
var poolCmd = &cobra.Command{
	Use:   "pool",
	Short: "Storage pool management",
	Long:  `Commands to manage the storage pool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need a subcommand (add|del|ls) with vmman pool")
	},
}

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

// pooldelCmd represents the pooldel command
var pooldelCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"del"},
	Short:   "TODO: Deletes a pool storage",
	Long:    `TODO: This will remove a pool storage both in QEMU (libvirt) and the databse.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pool rm NOT IMPLEMENTED")
		//storagepool.PoolAdd(args[0])
	},
}

// TODO: add capability to list pools from QEMU
// poollistCmd represents the poollist command
var poollistCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all pools currently configured in the DB",
	Long:  `If a pool has been configured but has not been inserted in the database it will not show up.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LIMITED FUNCTIONALITY !! What you see here is what's in the DB, not in QEMU")
		storagepool.PoolList()
	},
}

func init() {
	rootCmd.AddCommand(poolCmd)
	poolCmd.AddCommand(pooladdCmd)
	poolCmd.AddCommand(pooldelCmd)
	poolCmd.AddCommand(poollistCmd)
}
