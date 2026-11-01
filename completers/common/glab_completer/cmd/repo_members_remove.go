package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_members_removeCmd = &cobra.Command{
	Use:   "remove [flags]",
	Short: "Remove a member from the project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_members_removeCmd).Standalone()

	repo_members_removeCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	repo_members_removeCmd.Flags().StringP("user-id", "u", "", "User ID instead of username.")
	repo_members_removeCmd.Flags().String("username", "", "Username instead of user-id.")
	repo_membersCmd.AddCommand(repo_members_removeCmd)

	carapace.Gen(repo_members_removeCmd).FlagCompletion(carapace.ActionMap{
		"repo":     action.ActionRepo(repo_members_removeCmd),
		"username": action.ActionUsers(repo_members_removeCmd),
	})
}
