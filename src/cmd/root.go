// Copyright S 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"vmman3/helpers"
)

// rootCmd represents the base command when called without any subcommands

var version = "0.320 (2022.10.03)"

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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&helpers.ConnectURI, "connection", "c", "", "Hypervisor URI.")
	rootCmd.PersistentFlags().StringVarP(&helpers.EnvironmentFile, "environment", "e", "environment.json", "Environment file.")
	rootCmd.PersistentFlags().BoolVarP(&helpers.BSingleHypervisor, "singleHypervisor", "1", false, "Connects to local hypervisor")
	rootCmd.PersistentFlags().BoolVarP(&helpers.BAllHypervisors, "allHypervisors", "a", true, "Make vmman multi hypervisor-aware")
}

// -a will always override -1 and -c $HYPERVISOR_NAME
// -1 will always override -c $HYPERVISOR_NAME : if -1 is set, it will act as if -c is set to qemu:///system
func initConfig() {
	helpers.ConnectURI, _ = rootCmd.Flags().GetString("connection")
	helpers.EnvironmentFile, _ = rootCmd.Flags().GetString("environment")
	helpers.BAllHypervisors, _ = rootCmd.Flags().GetBool("allHypervisors")
	helpers.BSingleHypervisor, _ = rootCmd.Flags().GetBool("singleHypervisor")

	// FIXME FIXME FIXME:
	// no flag screws ssh @ qemu URI
	if helpers.ConnectURI != "" {
		helpers.BAllHypervisors = false
		helpers.BSingleHypervisor = false
	}

	if helpers.BAllHypervisors {
		helpers.BSingleHypervisor = false
		helpers.ConnectURI = ""
	} else {
		if helpers.BSingleHypervisor {
			helpers.BAllHypervisors = false
			helpers.ConnectURI = ""
		}
	}
}
