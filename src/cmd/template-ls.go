/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"vmman3/template"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var templlsCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all templated VM",
	Long:  `This subcommand will list all templated VM with some extra information.`,
	Run: func(cmd *cobra.Command, args []string) {
		template.TemplateLS()
	},
}

func init() {
	templateCmd.AddCommand(templlsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
