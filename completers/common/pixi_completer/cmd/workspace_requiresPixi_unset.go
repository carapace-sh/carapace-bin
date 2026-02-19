package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_requiresPixi_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_requiresPixi_unsetCmd).Standalone()

	workspace_requiresPixi_unsetCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_requiresPixiCmd.AddCommand(workspace_requiresPixi_unsetCmd)

	carapace.Gen(workspace_requiresPixi_unsetCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
