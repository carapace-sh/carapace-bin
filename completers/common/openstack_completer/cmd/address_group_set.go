package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_group_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set address group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_group_setCmd).Standalone()

	address_group_setCmd.Flags().String("address", "", "IP address or CIDR (repeat option to set multiple addresses)")
	address_group_setCmd.Flags().String("description", "", "Set address group description")
	address_group_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	address_group_setCmd.Flags().String("name", "", "Set address group name")
	address_groupCmd.AddCommand(address_group_setCmd)
}
