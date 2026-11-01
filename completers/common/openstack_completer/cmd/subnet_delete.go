package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var subnet_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete subnet(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_deleteCmd).Standalone()

	subnetCmd.AddCommand(subnet_deleteCmd)

	carapace.Gen(subnet_deleteCmd).PositionalAnyCompletion(action.ActionSubnets(subnet_deleteCmd))
}
