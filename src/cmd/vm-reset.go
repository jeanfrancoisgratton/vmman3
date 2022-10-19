// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// cmd/cmd-vm-reset.go
// 2022-08-22 19:23:22

package cmd

import (
	"vmman3/vmmanagement"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:     "reset",
	Aliases: []string{"bounce", "reboot", "restart"},
	Short:   "Restart a single or multiple VM(s)",
	Long:    `The list of VMs needing to be restarted has to be space-separated`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Stop(args)
		vmmanagement.Start(args)
	},
}

func init() {
	vmCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
