// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/vm-rm.go
// 2022-10-28 18:57:39

package cmd

import (
	"vmman3/helpers"
	"vmman3/vmmanagement"

	"github.com/spf13/cobra"
)

// vmRmCmd represents the vmRm command
var vmRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a VM",
	Long: `This will shut the VM down if running, and optionally offer to keep its storage.
By default the storage is also removed.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Remove(args)
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
