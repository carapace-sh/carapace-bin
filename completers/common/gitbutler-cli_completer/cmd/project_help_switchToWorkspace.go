package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var project_help_switchToWorkspaceCmd = &cobra.Command{
	Use:   "switch-to-workspace",
	Short: "Switch back to the workspace branch for use of virtual branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_help_switchToWorkspaceCmd).Standalone()

	project_helpCmd.AddCommand(project_help_switchToWorkspaceCmd)
}
