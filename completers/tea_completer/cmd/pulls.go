package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var pullsCmd = &cobra.Command{
	Use:     "pulls",
	Short:   "Manage and checkout pull requests",
	GroupID: "ENTITIES",
	Aliases: []string{"pull", "pr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullsCmd).Standalone()

	pullsCmd.Flags().Bool("comments", false, "Whether to display comments (will prompt if not provided & run interactively)")
	pullsCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	pullsCmd.Flags().String("limit", "", "specify limit of items per page")
	pullsCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pullsCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pullsCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	pullsCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pullsCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pullsCmd.Flags().String("state", "", "Filter by state (all|open|closed)")
	rootCmd.AddCommand(pullsCmd)

	// TODO completion
	carapace.Gen(pullsCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionPullrequestFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"state":  carapace.ActionValues("all", "open", "closed").StyleF(style.ForKeyword),
	})
}
