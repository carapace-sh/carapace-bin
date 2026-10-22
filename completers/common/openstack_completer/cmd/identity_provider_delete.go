package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var identity_provider_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete identity provider(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identity_provider_deleteCmd).Standalone()

	identity_providerCmd.AddCommand(identity_provider_deleteCmd)

	carapace.Gen(identity_provider_deleteCmd).PositionalAnyCompletion(action.ActionIdentityProviders(identity_provider_deleteCmd))
}
