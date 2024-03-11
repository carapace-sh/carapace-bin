package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var admin_usersCmd = &cobra.Command{
	Use:     "users",
	Short:   "Manage registered users",
	Aliases: []string{"u"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(admin_usersCmd).Standalone()

	admin_usersCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	admin_usersCmd.Flags().String("limit", "", "specify limit of items per page")
	admin_usersCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	admin_usersCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	admin_usersCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	admin_usersCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	admin_usersCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	adminCmd.AddCommand(admin_usersCmd)

	// TODO repo
	carapace.Gen(admin_users_listCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionUserFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
