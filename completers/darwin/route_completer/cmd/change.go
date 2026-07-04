package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change aspects of a route",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changeCmd).Standalone()
	changeCmd.Flags().Bool("host", false, "Destination is a host")
	changeCmd.Flags().String("ifscope", "", "Bind to the specified interface")
	changeCmd.Flags().Bool("net", false, "Destination is a network")
	changeCmd.Flags().String("netmask", "", "Specify the network mask")
	rootCmd.AddCommand(changeCmd)
}
