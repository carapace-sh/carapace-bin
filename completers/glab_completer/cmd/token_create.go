package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var token_createCmd = &cobra.Command{
	Use:     "create <name>",
	Short:   "Creates user, group, or project access tokens.",
	Aliases: []string{"create", "new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_createCmd).Standalone()

	token_createCmd.Flags().StringP("access-level", "A", "", "Access level of the token: one of 'guest', 'reporter', 'developer', 'maintainer', 'owner'.")
	token_createCmd.Flags().String("description", "", "Sets the token's description.")
	token_createCmd.Flags().StringP("duration", "D", "", "Sets the token duration, in hours. Maximum of 8760. Examples: 24h, 168h, 504h.")
	token_createCmd.Flags().StringP("expires-at", "E", "", "Sets the token's expiration date and time, in YYYY-MM-DD format. If not specified, --duration is used.")
	token_createCmd.Flags().StringP("group", "g", "", "Create a group access token. Ignored if a user or repository argument is set.")
	token_createCmd.Flags().StringP("output", "F", "", "Format output as 'text' for the token value, 'json' for the actual API token structure.")
	token_createCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	token_createCmd.Flags().StringSliceP("scope", "S", nil, "Scopes for the token. For a list, see https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#personal-access-token-scopes.")
	token_createCmd.Flags().StringP("user", "U", "", "Create a personal access token. For the current user, use @me.")
	tokenCmd.AddCommand(token_createCmd)

	carapace.Gen(token_createCmd).FlagCompletion(carapace.ActionMap{
		"access-level": carapace.ActionValues("guest", "reporter", "developer", "maintainer", "owner"),
		"group":        action.ActionGroups(token_createCmd),
		"output":       carapace.ActionValues("text", "json"),
		// TODO scope
		"user": action.ActionUsers(token_createCmd),
	})
}
