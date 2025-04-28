package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var token_revokeCmd = &cobra.Command{
	Use:     "revoke <token-name|token-id>",
	Short:   "Revoke user, group or project access tokens",
	Aliases: []string{"revoke", "rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_revokeCmd).Standalone()

	token_revokeCmd.Flags().StringP("group", "g", "", "Revoke group access token. Ignored if a user or repository argument is set.")
	token_revokeCmd.Flags().StringP("output", "F", "", "Format output as: text, json. 'text' provides the name and ID of the revoked token; 'json' outputs the token with metadata.")
	token_revokeCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	token_revokeCmd.Flags().StringP("user", "U", "", "Revoke personal access token. Use @me for the current user.")
	tokenCmd.AddCommand(token_revokeCmd)

	carapace.Gen(token_revokeCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(token_revokeCmd),
		"repo":  action.ActionRepo(token_revokeCmd),
		"user":  action.ActionUsers(token_revokeCmd),
	})

	// TODO complete tokens
}
