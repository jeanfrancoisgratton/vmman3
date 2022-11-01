// Copyright S 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"os"
	"strings"
	"vmman3/helpers"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands

var version = "0.550 (2022.11.01)"

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

	rootCmd.PersistentFlags().StringVarP(&helpers.ConnectURI, "connection", "c", "qemu:///system", "Target hypervisor")
	rootCmd.PersistentFlags().StringVarP(&helpers.EnvironmentFile, "environment", "e", "environment.json", "Environment file.")
	rootCmd.PersistentFlags().BoolVarP(&helpers.BAllHypervisors, "allHypervisors", "a", false, "Make vmman multi hypervisor-aware")
}

// -a will always override -c $HYPERVISOR_NAME
func initConfig() {
	helpers.ConnectURI, _ = rootCmd.Flags().GetString("connection")
	helpers.EnvironmentFile, _ = rootCmd.Flags().GetString("environment")
	helpers.BAllHypervisors, _ = rootCmd.Flags().GetBool("allHypervisors")
	//helpers.BSingleHypervisor, _ = rootCmd.Flags().GetBool("singleHypervisor")

	// Checks if environment file name ends with ".json"
	if !strings.HasSuffix(helpers.EnvironmentFile, ".json") {
		helpers.EnvironmentFile += ".json"
	}

	// Some logic to avoid parameter clash
	if helpers.BAllHypervisors {
		helpers.ConnectURI = ""
	} else {
		if helpers.ConnectURI != "qemu:///system" {
			helpers.ConnectURI = "qemu+ssh://" + helpers.BuildConnectURI(helpers.ConnectURI) + "@" + helpers.ConnectURI + "/system"
		}
	}
}
