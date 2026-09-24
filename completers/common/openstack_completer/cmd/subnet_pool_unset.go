package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_pool_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset subnet pool properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_pool_unsetCmd).Standalone()

	subnet_pool_unsetCmd.Flags().Bool("all-tag", false, "Clear all tags associated with the subnet pool")
	subnet_pool_unsetCmd.Flags().String("tag", "", "Tag to be removed from the subnet pool (repeat option to remove multiple tags)")
	subnet_poolCmd.AddCommand(subnet_pool_unsetCmd)
}
