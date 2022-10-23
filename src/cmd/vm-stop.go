// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// cmd/cmd-vm-stop.go
// 2022-08-22 13:16:31

package cmd

import (
	"github.com/spf13/cobra"
	"vmman3/vm_management"
)

// vmstopCmd represents the vmstop command
var vmstopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"down"},
	Short:   "Stop one or multiple VMs",
	Long: `This command is used to stop one or multiple virtual machines (VMs):

If more than a single VM needs to be stopped, you just add them to the commandline, space-separated..`,
	Run: func(cmd *cobra.Command, args []string) {
		vm_management.Stop(args)
	},
}

func init() {
	rootCmd.AddCommand(vmstopCmd)
	vmCmd.AddCommand(vmstopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmstopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmstopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
