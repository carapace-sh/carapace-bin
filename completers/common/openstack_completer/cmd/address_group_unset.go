package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_group_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset address group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_group_unsetCmd).Standalone()

	address_group_unsetCmd.Flags().String("address", "", "IP address or CIDR (repeat option to unset multiple addresses)")
	address_groupCmd.AddCommand(address_group_unsetCmd)
}
