package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_removeCmd = &cobra.Command{
	Use:   "remove <secret-name>",
	Short: "Remove an organization or repository secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	secret_removeCmd.Flags().StringP("org", "o", "", "List secrets for an organization")
	secretCmd.AddCommand(secret_removeCmd)

	carapace.Gen(secret_removeCmd).FlagCompletion(carapace.ActionMap{
		"org": action.ActionUsers(secret_removeCmd, action.UserOpts{Organizations: true}),
	})

	carapace.Gen(secret_removeCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSecrets(secret_removeCmd, secret_removeCmd.Flag("org").Value.String())
		}),
	)
}
