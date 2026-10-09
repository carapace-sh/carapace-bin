package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set network flavor properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_setCmd).Standalone()

	network_flavor_setCmd.Flags().String("description", "", "Set network flavor description")
	network_flavor_setCmd.Flags().Bool("disable", false, "Disable network flavor")
	network_flavor_setCmd.Flags().Bool("enable", false, "Enable network flavor")
	network_flavor_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_flavor_setCmd.Flags().String("name", "", "Set flavor name")
	network_flavorCmd.AddCommand(network_flavor_setCmd)
}
