package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_removeCmd).Standalone()
	secret_removeCmd.Flags().StringP("app", "a", "", "Remove a secret for a specific application: {actions|codespaces|dependabot}")
	secret_removeCmd.Flags().StringP("env", "e", "", "Remove a secret for an environment")
	secret_removeCmd.Flags().StringP("org", "o", "", "Remove a secret for an organization")
	secret_removeCmd.Flags().BoolP("user", "u", false, "Remove a secret for your user")
	secretCmd.AddCommand(secret_removeCmd)

	carapace.Gen(secret_removeCmd).FlagCompletion(carapace.ActionMap{
		"app": carapace.ActionValues("actions", "codespaces", "dependabot"),
		"env": action.ActionEnvironments(secret_removeCmd),
		"org": action.ActionUsers(secret_removeCmd, action.UserOpts{Organizations: true}),
	})

	carapace.Gen(secret_removeCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSecrets(secret_removeCmd, action.SecretOpts{
				Org: secret_removeCmd.Flag("org").Value.String(),
				Env: secret_removeCmd.Flag("env").Value.String(),
			},
			)
		}),
	)
}
