// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/hypervisor-add.go
// 2022-11-11 22:20:54

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/hypervisor"
)

// hypervisorAddCmd represents the hypervisorAdd command
var hypervisorAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an hypervisor to the DB",
	Long: `This is where you add an hypervisor in the DB, with its up address (or resolvable hostname) and connecting username.

USAGE: vmman hypervisor add HYPERVISOR_NAME HYPERVISOR_ADDR HYPERVISOR_USER`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("USAGE: vmman hypervisor add HYPERVISOR_NAME HYPERVISOR_ADDR HYPERVISOR_USER")
			os.Exit(0)
		}
		hypervisor.AddHypervisor(args[0], args[1], args[2])
	},
}

func init() {
	hypervisorCmd.AddCommand(hypervisorAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hypervisorAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hypervisorAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
