package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var projects_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_rmCmd).Standalone()

	projectsCmd.AddCommand(projects_rmCmd)

	carapace.Gen(projects_rmCmd).PositionalCompletion(
		action.ActionProjects(projects_rmCmd),
	)
}
