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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a stub. You need a sub-command (add/remove)")
	},
}

func init() {
	rootCmd.AddCommand(hypervisorCmd)
}
