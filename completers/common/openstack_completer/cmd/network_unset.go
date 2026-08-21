package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset network properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_unsetCmd).Standalone()

	network_unsetCmd.Flags().Bool("all-tag", false, "Clear all tags associated with the network")
	network_unsetCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_unsetCmd.Flags().String("tag", "", "Tag to be removed from the network (repeat option to remove multiple tags)")
	networkCmd.AddCommand(network_unsetCmd)
}
