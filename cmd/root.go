// Copyright S 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/helpers"
)

// rootCmd represents the base command when called without any subcommands

var version = "1.000-0 (2022.08.16)"

//var connectURI string

var rootCmd = &cobra.Command{
	Use:     "vmman3",
	Version: version,
	Short:   "Go-based libvirtd client",
	Long: `This is a GoLang libvirtd client

This is a custom replacement for the stock virsh shell.
This program will allow you to manipulate vCPUs, vMEM, snapshots,
Stop/start/reboot a VM, list its specs, etc.

You can also manipulate VMs across hypervisors, handle clusters of VMs at once, etc.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vmman3.yaml)")
	rootCmd.PersistentFlags().StringVarP(&helpers.ConnectURI, "connection", "c", "qemu:///system", "Hypervisor URI.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	snapCmd.AddCommand(lssnapCmd)
	snapCmd.AddCommand(createsnapCmd)
	snapCmd.AddCommand(rmsnapCmd)
}

func initConfig() {
	helpers.ConnectURI, _ = rootCmd.Flags().GetString("connection")

	if helpers.ConnectURI != "qemu:///system" {
		connectURI := fmt.Sprintf("qemu+ssh://root@%s/system", helpers.ConnectURI)
		helpers.ConnectURI = connectURI
	}
}
