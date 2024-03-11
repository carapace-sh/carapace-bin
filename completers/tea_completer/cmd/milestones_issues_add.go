package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var milestones_issues_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add an issue/pull to an milestone",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestones_issues_addCmd).Standalone()

	milestones_issues_addCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	milestones_issues_addCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	milestones_issues_addCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	milestones_issues_addCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	milestones_issuesCmd.AddCommand(milestones_issues_addCmd)

	// TODO completion
	carapace.Gen(milestones_issues_addCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
