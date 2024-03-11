package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var milestones_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an milestone on repository",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestones_createCmd).Standalone()

	milestones_createCmd.Flags().StringP("deadline", "x", "", "set milestone deadline (default is no due date)")
	milestones_createCmd.Flags().StringP("description", "d", "", "milestone description to create")
	milestones_createCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	milestones_createCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	milestones_createCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	milestones_createCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	milestones_createCmd.Flags().String("state", "", "set milestone state (default is open)")
	milestones_createCmd.Flags().StringP("title", "t", "", "milestone title to create")
	milestonesCmd.AddCommand(milestones_createCmd)

	// TODO completion
	carapace.Gen(milestones_createCmd).FlagCompletion(carapace.ActionMap{
		"deadline": time.ActionDate(),
		"login":    tea.ActionLogins(),
		"output":   tea.ActionOutputFormats(),
		"remote":   git.ActionRemotes(),
		"state":    carapace.ActionValues("open", "closed").StyleF(style.ForKeyword),
	})
}
