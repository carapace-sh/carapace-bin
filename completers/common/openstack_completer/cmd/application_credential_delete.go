package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var application_credential_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete application credentials(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(application_credential_deleteCmd).Standalone()

	application_credentialCmd.AddCommand(application_credential_deleteCmd)

	carapace.Gen(application_credential_deleteCmd).PositionalAnyCompletion(action.ActionApplicationCredentials(application_credential_deleteCmd))
}
