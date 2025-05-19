package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var milestones_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a milestone",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestones_deleteCmd).Standalone()

	milestones_deleteCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	milestones_deleteCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	milestones_deleteCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	milestones_deleteCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	milestonesCmd.AddCommand(milestones_deleteCmd)

	// TODO completion
	carapace.Gen(milestones_deleteCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
