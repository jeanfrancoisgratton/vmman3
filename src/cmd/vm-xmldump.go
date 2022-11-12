// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/vm-xmldump.go
// 2022-11-05 13:42:54

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vmman3/vmmanagement"
)

// xmldumpCmd represents the xmldump command
var xmldumpCmd = &cobra.Command{
	Use:   "xmldump",
	Short: "Dumps the VM definition in an XML file",
	Long:  `This the equivalent of virsh xmldump, where you specify as arguments the VM name and the target XML file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("USAGE: vmman vm xmldump VMNAME XMLFILE")
			os.Exit(0)
		}
		vmmanagement.XmlDump(args[0], args[1])
	},
}

func init() {
	vmCmd.AddCommand(xmldumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xmldumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xmldumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
