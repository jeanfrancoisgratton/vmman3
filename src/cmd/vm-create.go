// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/vm-create.go
// 2022-11-01 18:43:57

package cmd

import (
	"vmman3/vmmanagement"

	"github.com/spf13/cobra"
)

// vmcreateCmd represents the vmcreate command
var vmcreateCmd = &cobra.Command{
	Use:   "vmcreate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Create(args)
	},
}

func init() {
	vmCmd.AddCommand(vmcreateCmd)
	rootCmd.AddCommand(vmcreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmcreateCmd.Pe
}
