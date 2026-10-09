package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_pool_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete subnet pool(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_pool_deleteCmd).Standalone()

	subnet_poolCmd.AddCommand(subnet_pool_deleteCmd)
}
