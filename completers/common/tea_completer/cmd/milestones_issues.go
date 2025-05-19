package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var milestones_issuesCmd = &cobra.Command{
	Use:     "issues",
	Short:   "manage issue/pull of an milestone",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestones_issuesCmd).Standalone()

	milestones_issuesCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	milestones_issuesCmd.Flags().String("kind", "", "Filter by kind (issue|pull)")
	milestones_issuesCmd.Flags().String("limit", "", "specify limit of items per page")
	milestones_issuesCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	milestones_issuesCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	milestones_issuesCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	milestones_issuesCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	milestones_issuesCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	milestones_issuesCmd.Flags().String("state", "", "Filter by issue state (all|open|closed)")
	milestonesCmd.AddCommand(milestones_issuesCmd)

	// TODO completion
	carapace.Gen(milestones_issuesCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionMilestoneFields().UniqueList(","),
		"kind":   carapace.ActionValues("issue", "pull"),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"state":  carapace.ActionValues("all", "open", "closed").StyleF(style.ForKeyword),
	})
}
