// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// cmd/ls.go
// 2022-08-16 17:47:06

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"vmman3/inventory"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List virtual machines",
	Long: `

Lists all virtual machines with some specs and their state.
Specs are : number of vCpu, amount of vMem, storage, ip address.
State is: status (running, stopped), the number of snapshots with
the name of the current one, if any, and so on.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\x1bc")
		inventory.VM_List()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
