// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/hypervisor.go
// 2022-11-11 21:26:58

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/hypervisor"
)

// hypervisorCmd represents the hypervisor command
var hypervisorCmd = &cobra.Command{
	Use:   "hypervisor",
	Short: "Hypervisor handling",
	Long:  `This is where you insert your various hypervisors in the DB.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("USAGE: vmman hypervisor {add|del} ...")
		}
	},
}

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

// hypervisorDelCmd represents the hypervisorDel command
var hypervisorDelCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"rm", "del"},
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

// hlistCmd represents the hlist command
var hlistCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all hypervisors",
	Long:  `Lists all hypervisors and their connecting users.`,
	Run: func(cmd *cobra.Command, args []string) {
		hypervisor.ListHypervisors()
	},
}

func init() {
	rootCmd.AddCommand(hypervisorCmd)
	hypervisorCmd.AddCommand(hypervisorAddCmd)
	hypervisorCmd.AddCommand(hypervisorDelCmd)
	hypervisorCmd.AddCommand(hlistCmd)
}
