package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_pool_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set subnet pool properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_pool_setCmd).Standalone()

	subnet_pool_setCmd.Flags().String("address-scope", "", "Set address scope associated with the subnet pool (name or ID), prefixes must be unique across address scopes")
	subnet_pool_setCmd.Flags().Bool("default", false, "Set this as a default subnet pool")
	subnet_pool_setCmd.Flags().String("default-prefix-length", "", "Set subnet pool default prefix length")
	subnet_pool_setCmd.Flags().String("default-quota", "", "Set default per-project quota for this subnet pool as the number of IP addresses that can be allocated from the subnet pool")
	subnet_pool_setCmd.Flags().String("description", "", "Set subnet pool description")
	subnet_pool_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	subnet_pool_setCmd.Flags().String("max-prefix-length", "", "Set subnet pool maximum prefix length")
	subnet_pool_setCmd.Flags().String("min-prefix-length", "", "Set subnet pool minimum prefix length")
	subnet_pool_setCmd.Flags().String("name", "", "Set subnet pool name")
	subnet_pool_setCmd.Flags().Bool("no-address-scope", false, "Remove address scope associated with the subnet pool")
	subnet_pool_setCmd.Flags().Bool("no-default", false, "Set this as a non-default subnet pool")
	subnet_pool_setCmd.Flags().Bool("no-tag", false, "Clear tags associated with the subnet pool.")
	subnet_pool_setCmd.Flags().String("pool-prefix", "", "Set subnet pool prefixes (in CIDR notation) (repeat option to set multiple prefixes)")
	subnet_pool_setCmd.Flags().String("tag", "", "Tag to be added to the subnet pool (repeat option to set multiple tags)")
	subnet_poolCmd.AddCommand(subnet_pool_setCmd)
}
