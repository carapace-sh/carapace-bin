package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var secret_deleteCmd = &cobra.Command{
	Use:     "delete <secret-name>",
	Short:   "Delete secrets",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_deleteCmd).Standalone()

	secret_deleteCmd.Flags().StringP("app", "a", "", "Delete a secret for a specific application: {actions|codespaces|dependabot}")
	secret_deleteCmd.Flags().StringP("env", "e", "", "Delete a secret for an environment")
	secret_deleteCmd.Flags().StringP("org", "o", "", "Delete a secret for an organization")
	secret_deleteCmd.Flags().BoolP("user", "u", false, "Delete a secret for your user")
	secretCmd.AddCommand(secret_deleteCmd)

	carapace.Gen(secret_deleteCmd).FlagCompletion(carapace.ActionMap{
		"app": carapace.ActionValues("actions", "codespaces", "dependabot"),
		"env": action.ActionEnvironments(secret_deleteCmd),
		"org": gh.ActionOrganizations(gh.HostOpts{}),
	})

	carapace.Gen(secret_deleteCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSecrets(secret_deleteCmd, action.SecretOpts{
				Org: secret_deleteCmd.Flag("org").Value.String(),
				Env: secret_deleteCmd.Flag("env").Value.String(),
			},
			)
		}),
	)
}
