// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/snap.go
// 2022-11-06 11:44:19

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/snapshotmanagement"
)

// snapCmd represents the snap command
var snapCmd = &cobra.Command{
	Use:     "snap",
	Aliases: []string{"snapshot"},
	Short:   "Snapshot management subcommand",
	Long: `This is where you manage all snapshots on a given VM.

Subcommands include : list, add, remove, revert`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to provide a subcommand: {create, rm, revert, ls}")
	},
}

// createsnapCmd represents the createsnap command
var snapCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a snapshot for the VM",
	Long:  `This will create a snapshot for the named VM.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("USAGE: vmman {snap|snapshot} create VMNAME SNAPNAME")
			os.Exit(0)
		}
		snapshotmanagement.CreateSnapshot(args[0], args[1])
	},
}

// rmsnapCmd represents the rmsnap command
var snapRmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"del"},
	Short:   "Remove snapshot",
	Long: `Removes the current snapshot from the VM.

NOTE: The snapshot cannot be removed if a child snapshot is present....yet`,
	// usage: snap rm VMNAME SNAPNAME
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			snapshotmanagement.RemoveSnapshot(args[0], args[1])
			os.Exit(0)
		} else if len(args) == 1 {
			snapshotmanagement.RemoveSnapshot(args[0], "")
			os.Exit(0)
		} else {
			fmt.Println("USAGE: vmman {snap|snapshot} rm VMNAME [SNAPSHOTNAME]")
			os.Exit(0)
		}
	},
}

// revertCmd represents the revert command
var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Reverts the VM to a snapshot",
	Long: `This subcommand allows you to revert the VM to a specific snapshot.
If no snapshotname is specified, it reverts to the current one.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			snapshotmanagement.RevertSnapshot(args[0], args[1])
			os.Exit(0)
		} else if len(args) == 1 {
			snapshotmanagement.RevertSnapshot(args[0], "")
			os.Exit(0)
		} else {
			fmt.Println("USAGE: vmman {snap|snapshot} revert VMNAME [SNAPSHOTNAME]")
			os.Exit(0)
		}
	},
}

// snapLSCmd represents the snapLS command
var snapLSCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "Lists all snapshots of a given VM",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Usage: vmman {snap|snapshot} {ls|list} VirtualMachine")
			os.Exit(0)
		}
		snapshotmanagement.ListSnapshots(args[0])
	},
}

func init() {
	rootCmd.AddCommand(snapCmd)
	snapCmd.AddCommand(snapCreateCmd)
	snapCmd.AddCommand(snapLSCmd)
	snapCmd.AddCommand(revertCmd)
	snapCmd.AddCommand(snapRmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
