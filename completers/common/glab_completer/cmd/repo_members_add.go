package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_members_addCmd = &cobra.Command{
	Use:   "add [flags]",
	Short: "Add a member to the project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_members_addCmd).Standalone()

	repo_members_addCmd.Flags().StringP("expires-at", "e", "", "Expiration date for the membership (YYYY-MM-DD)")
	repo_members_addCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	repo_members_addCmd.Flags().StringP("role", "r", "", "Role for the user (guest, reporter, developer, maintainer, owner)")
	repo_members_addCmd.Flags().String("role-id", "", "ID of a custom role defined in the project or group.")
	repo_members_addCmd.Flags().StringP("user-id", "u", "", "User ID instead of username.")
	repo_members_addCmd.Flags().String("username", "", "Username instead of user-id.")
	repo_membersCmd.AddCommand(repo_members_addCmd)

	carapace.Gen(repo_members_addCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(repo_members_addCmd),
	})
}
