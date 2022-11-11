// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/vm-console.go
// 2022-11-05 17:55:36

package cmd

import (
	"github.com/spf13/cobra"
	"vmman3/vmmanagement"
)

// consoleCmd represents the console command
var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Connects the terminal to the VM console",
	Long: `Will connect you to the VM console.
Press CTRL+] to disconnect.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Console(args[0])
	},
}

func init() {
	rootCmd.AddCommand(consoleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consoleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consoleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
