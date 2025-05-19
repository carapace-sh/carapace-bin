package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var token_rotateCmd = &cobra.Command{
	Use:     "rotate <token-name|token-id>",
	Short:   "Rotate user, group, or project access tokens",
	Aliases: []string{"rotate", "rot"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_rotateCmd).Standalone()

	token_rotateCmd.Flags().StringP("duration", "D", "", "Sets the token duration, in hours. Maximum of 8760. Examples: 24h, 168h, 504h.")
	token_rotateCmd.Flags().StringP("expires-at", "E", "", "Sets the token's expiration date and time, in YYYY-MM-DD format. If not specified, --duration is used.")
	token_rotateCmd.Flags().StringP("group", "g", "", "Rotate group access token. Ignored if a user or repository argument is set.")
	token_rotateCmd.Flags().StringP("output", "F", "", "Format output as: text, json. 'text' provides the new token value; 'json' outputs the token with metadata.")
	token_rotateCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	token_rotateCmd.Flags().StringP("user", "U", "", "Rotate personal access token. Use @me for the current user.")
	tokenCmd.AddCommand(token_rotateCmd)

	carapace.Gen(token_rotateCmd).FlagCompletion(carapace.ActionMap{
		"group":  action.ActionGroups(token_rotateCmd),
		"output": carapace.ActionValues("text", "json"),
		"repo":   action.ActionRepo(token_rotateCmd),
		"user":   action.ActionUsers(token_rotateCmd),
	})

	// TODO complete tokens
}
