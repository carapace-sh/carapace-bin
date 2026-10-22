package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var user_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_deleteCmd).Standalone()

	user_deleteCmd.Flags().String("domain", "", "Domain owning <user> (name or ID)")
	userCmd.AddCommand(user_deleteCmd)

	carapace.Gen(user_deleteCmd).FlagCompletion(carapace.ActionMap{
		"domain": action.ActionDomains(user_deleteCmd),
	})

	carapace.Gen(user_deleteCmd).PositionalAnyCompletion(action.ActionUsers(user_deleteCmd))
}
