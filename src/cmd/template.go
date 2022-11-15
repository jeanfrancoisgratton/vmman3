// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/templatemanagement.go

package cmd

import (
	"fmt"
	"vmman3/templatemanagement"

	"github.com/spf13/cobra"
)

// templateCmd represents the templatemanagement command
var templateCmd = &cobra.Command{
	Use:   "templatemanagement",
	Short: "Template-related commands",
	Long: `Here you will templatemanagement commands such as:

	- Editing templatemanagement specs
	- Removing templates
	- Adding new templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("templatemanagement called")
	},
}

// listCmd represents the list command
var templlsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "Lists all templated VM",
	Long:    `This subcommand will list all templated VM with some extra information.`,
	Run: func(cmd *cobra.Command, args []string) {
		templatemanagement.List()
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.AddCommand(templlsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
