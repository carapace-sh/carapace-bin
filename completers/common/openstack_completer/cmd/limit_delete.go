package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var limit_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a limit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(limit_deleteCmd).Standalone()

	limitCmd.AddCommand(limit_deleteCmd)

	carapace.Gen(limit_deleteCmd).PositionalAnyCompletion(action.ActionLimits(limit_deleteCmd))
}
