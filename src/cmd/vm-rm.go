/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"vmman3/helpers"
	"vmman3/vm_management"

	"github.com/spf13/cobra"
)

// vmRmCmd represents the vmRm command
var vmRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		vm_management.Remove(args)
	},
}

func init() {
	rootCmd.AddCommand(vmRmCmd)
	vmCmd.AddCommand(vmRmCmd)

	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmRmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	vmRmCmd.Flags().BoolP("keepStorage", "k", false, "Keeps the storage (disks) of this VM after removal")

	helpers.BkeepStorage, _ = vmRmCmd.Flags().GetBool("keepStorage")
}
