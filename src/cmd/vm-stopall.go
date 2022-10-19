/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"vmman3/vmmanagement"

	"github.com/spf13/cobra"
)

// stopallCmd represents the stopall command
var stopallCmd = &cobra.Command{
	Use:   "stopall",
	Short: "Stop all VMs",
	Long:  `Stops all the VMs under the given hypervisor.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.StopAll()
	},
}

func init() {
	vmCmd.AddCommand(stopallCmd)
	rootCmd.AddCommand(stopallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
