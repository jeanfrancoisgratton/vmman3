// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// cmd/cmd-vm-start.go
// 2022-08-22 13:15:26

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vmstartCmd represents the vmstart command
var vmstartCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"up"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vmstart called")
	},
}

func init() {
	vmCmd.AddCommand(vmstartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmstartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmstartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
