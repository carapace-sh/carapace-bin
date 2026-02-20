package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_requiresPixiCmd = &cobra.Command{
	Use:   "requires-pixi",
	Short: "Commands to manage the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_requiresPixiCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_requiresPixiCmd)
}
