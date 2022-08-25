// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Cluster commands",
	Long: `This is where you will edit cluster (groups) of VMs to be running basic commands together.

Basic commands are: stop, start, reboot, revert snapshot, create snapshot.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cluster called")
	},
}

func init() {
	rootCmd.AddCommand(clusterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
