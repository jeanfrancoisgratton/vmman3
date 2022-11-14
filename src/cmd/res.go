// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/res.go
// 2022-11-13 22:57:36

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resCmd represents the res command
var resCmd = &cobra.Command{
	Use:   "res",
	Short: "Resources management",
	Long:  `This is where you add disks, vCPUS, vMEM, to an existing VM.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("res called")
	},
}

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cpu called")
	},
}

func init() {
	rootCmd.AddCommand(resCmd)
	resCmd.AddCommand(cpuCmd)
}
