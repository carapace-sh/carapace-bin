package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var pulls_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List pull requests of the repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_listCmd).Standalone()

	pulls_listCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	pulls_listCmd.Flags().String("limit", "", "specify limit of items per page")
	pulls_listCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_listCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pulls_listCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	pulls_listCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_listCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pulls_listCmd.Flags().String("state", "", "Filter by state (all|open|closed)")
	pullsCmd.AddCommand(pulls_listCmd)

	// TODO completion
	carapace.Gen(pulls_listCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionPullrequestFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"state":  carapace.ActionValues("all", "open", "closed").StyleF(style.ForKeyword),
	})
}
