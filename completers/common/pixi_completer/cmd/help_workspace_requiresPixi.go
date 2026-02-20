package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_requiresPixiCmd = &cobra.Command{
	Use:   "requires-pixi",
	Short: "Commands to manage the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_requiresPixiCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_requiresPixiCmd)
}
