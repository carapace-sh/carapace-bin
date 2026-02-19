package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_requiresPixiCmd = &cobra.Command{
	Use:   "requires-pixi",
	Short: "Commands to manage the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_requiresPixiCmd).Standalone()

	workspace_requiresPixiCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspaceCmd.AddCommand(workspace_requiresPixiCmd)

	carapace.Gen(workspace_requiresPixiCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
