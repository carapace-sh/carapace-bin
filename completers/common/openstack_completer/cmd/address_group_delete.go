package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var address_group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete address group(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_group_deleteCmd).Standalone()

	address_groupCmd.AddCommand(address_group_deleteCmd)

	carapace.Gen(address_group_deleteCmd).PositionalAnyCompletion(action.ActionAddressGroups(address_group_deleteCmd))
}
