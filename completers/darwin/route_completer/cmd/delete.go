package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a route",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()
	deleteCmd.Flags().Bool("host", false, "Destination is a host")
	deleteCmd.Flags().String("ifscope", "", "Bind to the specified interface")
	deleteCmd.Flags().Bool("net", false, "Destination is a network")
	deleteCmd.Flags().String("netmask", "", "Specify the network mask")
	rootCmd.AddCommand(deleteCmd)
}
