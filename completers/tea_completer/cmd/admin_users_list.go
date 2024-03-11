package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var admin_users_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Users",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(admin_users_listCmd).Standalone()

	admin_users_listCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	admin_users_listCmd.Flags().String("limit", "", "specify limit of items per page")
	admin_users_listCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	admin_users_listCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	admin_users_listCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	admin_users_listCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	admin_users_listCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	admin_usersCmd.AddCommand(admin_users_listCmd)

	// TODO repo
	carapace.Gen(admin_users_listCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionUserFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
