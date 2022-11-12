// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/hypervisor-del.go
// 2022-11-11 22:21:05

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hypervisorDelCmd represents the hypervisorDel command
var hypervisorDelCmd = &cobra.Command{
	Use:   "hypervisorDel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hypervisorDel called")
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
