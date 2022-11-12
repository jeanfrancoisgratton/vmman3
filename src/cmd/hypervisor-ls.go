// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/hypervisor-ls.go
// 2022-11-12 11:29:22

package cmd

import (
	"vmman3/hypervisor"

	"github.com/spf13/cobra"
)

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
	hypervisorCmd.AddCommand(hlistCmd)
}
