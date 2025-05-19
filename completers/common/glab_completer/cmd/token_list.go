package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var token_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List user, group, or project access tokens.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_listCmd).Standalone()

	token_listCmd.Flags().BoolP("active", "a", false, "List only the active tokens.")
	token_listCmd.Flags().StringP("group", "g", "", "List group access tokens. Ignored if a user or repository argument is set.")
	token_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json. text provides a readable table, json outputs the tokens with metadata.")
	token_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	token_listCmd.Flags().StringP("user", "U", "", "List personal access tokens. Use @me for the current user.")
	tokenCmd.AddCommand(token_listCmd)

	// TODO complete group
	carapace.Gen(token_listCmd).FlagCompletion(carapace.ActionMap{
		"group":  action.ActionGroups(token_listCmd),
		"output": carapace.ActionValues("text", "json"),
		"repo":   action.ActionRepo(token_listCmd),
		"user":   action.ActionUsers(token_listCmd),
	})
}
