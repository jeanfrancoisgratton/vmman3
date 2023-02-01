// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// cmd/cmd-vm.go
// 2022-08-22 13:14:43

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
	"vmman3/vmmanagement"
)

// vmCmd represents the vm command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "VM management subcommands",
	Long:  `From here you manage all vm-related commands: create, delete, start, stop, etc.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// consoleCmd represents the console command
var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Connects the terminal to the VM console",
	Long: `Will connect you to the VM console.
Press CTRL+] to disconnect.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Console(args[0])
	},
}

// renameCmd represents the rename command
var vmrenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a VM",
	Long:  `This command will rename a virtual machine. If the machine is running will be shut down before.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Rename(args)
	},
}

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:     "reset",
	Aliases: []string{"bounce", "reboot", "restart"},
	Short:   "Restart a single or multiple VM(s)",
	Long:    `The list of VMs needing to be restarted has to be space-separated`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Stop(args)
		time.Sleep(5 * time.Second) // needed otherwise vmStart will think that the VM is already up
		vmmanagement.Start(args)
	},
}

// vmRmCmd represents the vmRm command
var vmRmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"del"},
	Short:   "Remove a VM",
	Long: `This will shut the VM down if running, and optionally offer to keep its storage.
By default the storage is also removed.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Remove(args)
	},
}

// vmstartCmd represents the vmstart command
var vmstartCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"up"},
	Short:   "Start one or multiple VMs",
	Long: `This command is used to start one or multiple virtual machines (VMs):

If more than a single VM needs to be started, you just add them to the commandline, space-separated..`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Start(args)
	},
}

// startallCmd represents the startall command
var startallCmd = &cobra.Command{
	Use:   "startall",
	Short: "Starts all VMs",
	Long:  `Starts all the VMs under the given hypervisor.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.StartAll()
	},
}

// xmldumpCmd represents the xmldump command
var xmldumpCmd = &cobra.Command{
	Use:     "xmldump",
	Short:   "Dumps the VM definition in an XML file",
	Aliases: []string{"dumpxml"},
	Long:    `This the equivalent of virsh xmldump, where you specify as arguments the VM name and the target XML file.`,
	Example: "vmman vm xmldump vmname xmlfile.",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("USAGE: vmman vm xmldump VMNAME XMLFILE")
			os.Exit(0)
		}
		vmmanagement.XmlDump(args[0], args[1])
	},
}

// stopallCmd represents the stopall command
var stopallCmd = &cobra.Command{
	Use:   "stopall",
	Short: "Stop all VMs",
	Long:  `Stops all the VMs under the given hypervisor.`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.StopAll()
	},
}

// vmstopCmd represents the vmstop command
var vmstopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"down"},
	Short:   "Stop one or multiple VMs",
	Long: `This command is used to stop one or multiple virtual machines (VMs):

If more than a single VM needs to be stopped, you just add them to the commandline, space-separated..`,
	Run: func(cmd *cobra.Command, args []string) {
		vmmanagement.Stop(args)
	},
}

func init() {
	rootCmd.AddCommand(vmCmd)
	rootCmd.AddCommand(consoleCmd)
	rootCmd.AddCommand(vmrenameCmd)
	vmCmd.AddCommand(vmrenameCmd)
	vmCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(vmRmCmd)
	vmCmd.AddCommand(vmRmCmd)
	rootCmd.AddCommand(vmstartCmd)
	vmCmd.AddCommand(vmstartCmd)
	rootCmd.AddCommand(startallCmd)
	vmCmd.AddCommand(startallCmd)
	rootCmd.AddCommand(vmstopCmd)
	vmCmd.AddCommand(vmstopCmd)
	vmCmd.AddCommand(stopallCmd)
	rootCmd.AddCommand(stopallCmd)
	vmCmd.AddCommand(xmldumpCmd)
}
