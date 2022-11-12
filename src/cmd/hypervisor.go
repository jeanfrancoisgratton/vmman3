// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/hypervisor.go
// 2022-11-11 21:26:58

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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

func init() {
	rootCmd.AddCommand(hypervisorCmd)
}
