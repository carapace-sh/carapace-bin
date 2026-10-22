package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var credential_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete credential(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credential_deleteCmd).Standalone()

	credentialCmd.AddCommand(credential_deleteCmd)

	carapace.Gen(credential_deleteCmd).PositionalAnyCompletion(action.ActionCredentials())
}
