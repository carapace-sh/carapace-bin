package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var milestones_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List milestones of the repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestones_listCmd).Standalone()

	milestones_listCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	milestones_listCmd.Flags().String("limit", "", "specify limit of items per page")
	milestones_listCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	milestones_listCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	milestones_listCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	milestones_listCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	milestones_listCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	milestones_listCmd.Flags().String("state", "", "Filter by milestone state (all|open|closed)")
	milestonesCmd.AddCommand(milestones_listCmd)

	// TODO completion
	carapace.Gen(milestones_listCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionMilestoneFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"state":  carapace.ActionValues("all", "open", "closed").StyleF(style.ForKeyword),
	})
}
