package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_profile_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set network flavor profile properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_profile_setCmd).Standalone()

	network_flavor_profile_setCmd.Flags().String("description", "", "Description for the flavor profile")
	network_flavor_profile_setCmd.Flags().Bool("disable", false, "Disable the flavor profile")
	network_flavor_profile_setCmd.Flags().String("driver", "", "Python module path to driver.")
	network_flavor_profile_setCmd.Flags().Bool("enable", false, "Enable the flavor profile")
	network_flavor_profile_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_flavor_profile_setCmd.Flags().String("metainfo", "", "Metainfo for the flavor profile.")
	network_flavor_profileCmd.AddCommand(network_flavor_profile_setCmd)
}
