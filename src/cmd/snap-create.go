// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/snap-create.go
// 2022-11-11 18:48:34

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/snapshotmanagement"
)

// createsnapCmd represents the createsnap command
var createsnapCmd = &cobra.Command{
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

func init() {
	snapCmd.AddCommand(createsnapCmd)
}
