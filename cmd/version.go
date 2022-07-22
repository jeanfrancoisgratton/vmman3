// Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows software version and changelog",
	Long:  `Shows the current version in the form of a CHAnGELOG.`,
	Run: func(cmd *cobra.Command, args []string) {
		showVer()
	},
}

func showVer() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	content, err := ioutil.ReadFile("_version")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
