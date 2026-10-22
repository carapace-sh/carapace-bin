package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var flavor_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete flavor(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flavor_deleteCmd).Standalone()

	flavorCmd.AddCommand(flavor_deleteCmd)

	carapace.Gen(flavor_deleteCmd).PositionalAnyCompletion(action.ActionFlavors(flavor_deleteCmd))
}
