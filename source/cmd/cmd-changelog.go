// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"github.com/spf13/cobra"
	"vmman3/helpers"
)

// changelogCmd represents the changelog command
var changelogCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Application changelog",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		helpers.Changelog()
	},
}

func init() {
	rootCmd.AddCommand(changelogCmd)
}
