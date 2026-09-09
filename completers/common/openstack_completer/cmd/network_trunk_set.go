package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_trunk_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set network trunk properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunk_setCmd).Standalone()

	network_trunk_setCmd.Flags().String("description", "", "A description of the trunk")
	network_trunk_setCmd.Flags().Bool("disable", false, "Disable trunk")
	network_trunk_setCmd.Flags().Bool("enable", false, "Enable trunk")
	network_trunk_setCmd.Flags().String("name", "", "Set trunk name")
	network_trunk_setCmd.Flags().String("subport", "", "Subport to add.")
	network_trunkCmd.AddCommand(network_trunk_setCmd)
}
