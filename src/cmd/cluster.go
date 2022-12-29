/*
Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Cluster management subcommands",
	Long: `A cluster is a group of VMs that you wish to manage all at once.

You can manage clusters the same way you manage single VMs: up, down, reboot, snapshots, etc.`,
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
