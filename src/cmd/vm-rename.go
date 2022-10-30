// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/vm-rename.go
// 2022-10-29 19:02:23

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a VM",
	Long:  `This command will rename a virtual machine. If the machine is running will be shut down before.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rename called")
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	vmCmd.AddCommand(vmRmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
