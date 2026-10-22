package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var trust_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete trust(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_deleteCmd).Standalone()

	trustCmd.AddCommand(trust_deleteCmd)

	carapace.Gen(trust_deleteCmd).PositionalAnyCompletion(action.ActionTrusts(trust_deleteCmd))
}
