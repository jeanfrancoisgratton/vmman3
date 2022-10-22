/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"vmman3/vm_management"

	"github.com/spf13/cobra"
)

// startallCmd represents the startall command
var startallCmd = &cobra.Command{
	Use:   "startall",
	Short: "Starts all VMs",
	Long:  `Starts all the VMs under the given hypervisor.`,
	Run: func(cmd *cobra.Command, args []string) {
		vm_management.StartAll()
	},
}

func init() {
	rootCmd.AddCommand(startallCmd)
	vmCmd.AddCommand(startallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
