package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var milestones_reopenCmd = &cobra.Command{
	Use:     "reopen",
	Short:   "Change state of one or more milestones to 'open'",
	Aliases: []string{"open"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestones_reopenCmd).Standalone()

	milestones_reopenCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	milestones_reopenCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	milestones_reopenCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	milestones_reopenCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	milestonesCmd.AddCommand(milestones_reopenCmd)

	// TODO completion
	carapace.Gen(milestones_reopenCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
