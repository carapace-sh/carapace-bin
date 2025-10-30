package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_project_switchToWorkspaceCmd = &cobra.Command{
	Use:   "switch-to-workspace",
	Short: "Switch back to the workspace branch for use of virtual branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_project_switchToWorkspaceCmd).Standalone()

	help_projectCmd.AddCommand(help_project_switchToWorkspaceCmd)
}
