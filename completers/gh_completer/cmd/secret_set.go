package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_setCmd = &cobra.Command{
	Use:   "set <secret-name>",
	Short: "Create or update secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	secret_setCmd.Flags().StringP("body", "b", "", "A value for the secret. Reads from STDIN if not specified.")
	secret_setCmd.Flags().StringP("env", "e", "", "Set a secret for an organization")
	secret_setCmd.Flags().StringP("org", "o", "", "List secrets for an organization")
	secret_setCmd.Flags().StringSliceP("repos", "r", []string{}, "List of repository names for `selected` visibility")
	secret_setCmd.Flags().StringP("visibility", "v", "private", "Set visibility for an organization secret: `all`, `private`, or `selected`")
	secretCmd.AddCommand(secret_setCmd)

	carapace.Gen(secret_setCmd).FlagCompletion(carapace.ActionMap{
		"env": action.ActionEnvironments(secret_setCmd),
		"org": action.ActionUsers(secret_setCmd, action.UserOpts{Organizations: true}),
		"repos": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionOwnerRepositories(secret_setCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"visibility": carapace.ActionValues("all", "private", "selected"),
	})

	carapace.Gen(secret_setCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSecrets(secret_setCmd, action.SecretOpts{
				Org: secret_setCmd.Flag("org").Value.String(),
				Env: secret_setCmd.Flag("env").Value.String(),
			},
			)
		}),
	)
}
