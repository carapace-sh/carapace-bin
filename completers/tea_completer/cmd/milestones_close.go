package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var milestones_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Change state of one or more milestones to 'closed'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestones_closeCmd).Standalone()

	milestones_closeCmd.Flags().BoolP("force", "f", false, "delete milestone")
	milestones_closeCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	milestones_closeCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	milestones_closeCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	milestones_closeCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	milestonesCmd.AddCommand(milestones_closeCmd)

	// TODO completion
	carapace.Gen(milestones_closeCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
