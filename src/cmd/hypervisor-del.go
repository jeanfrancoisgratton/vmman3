// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/hypervisor-del.go
// 2022-11-11 22:21:05

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"vmman3/hypervisor"
)

// hypervisorDelCmd represents the hypervisorDel command
var hypervisorDelCmd = &cobra.Command{
	Use:     "del",
	Aliases: []string{"rm", "delete"},
	Short:   "Removes an hypervisor from the DB",
	Long: `This is where you remove an hypervisor in the DB, optionally a user from an hypervisor.

USAGE: vmman hypervisor rm HYPERVISOR_NAME [HYPERVISOR_USER]`,
	Run: func(cmd *cobra.Command, args []string) {
		nArgs := len(args)
		switch {
		case nArgs == 1:
			hypervisor.DelHypervisor(args[0], "")
		case nArgs > 1:
			hypervisor.DelHypervisor(args[0], args[1])
		default:
			fmt.Println("USAGE: vmman hypervisor {rm|del} HYPERVISOR_NAME [HYPERVISOR_USER]")
		}
	},
}

func init() {
	hypervisorCmd.AddCommand(hypervisorDelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hypervisorDelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hypervisorDelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
