package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var keypair_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete public or private key(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keypair_deleteCmd).Standalone()

	keypair_deleteCmd.Flags().String("user", "", "The owner of the keypair.")
	keypair_deleteCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	keypairCmd.AddCommand(keypair_deleteCmd)

	carapace.Gen(keypair_deleteCmd).FlagCompletion(carapace.ActionMap{
		"user":        action.ActionUsers(),
		"user-domain": action.ActionDomains(),
	})

	carapace.Gen(keypair_deleteCmd).PositionalAnyCompletion(action.ActionKeypairs())
}
