// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// cmd/cmd-vm-start.go
// 2022-08-22 13:15:26

package cmd

import (
	"vmman3/vm_management"

	"github.com/spf13/cobra"
)

// vmstartCmd represents the vmstart command
var vmstartCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"up"},
	Short:   "Start one or multiple VMs",
	Long: `This command is used to start one or multiple virtual machines (VMs):

If more than a single VM needs to be started, you just add them to the commandline, space-separated..`,
	Run: func(cmd *cobra.Command, args []string) {
		vm_management.Start(args)
	},
}

func init() {
	rootCmd.AddCommand(vmstartCmd)
	vmCmd.AddCommand(vmstartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmstartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmstartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
