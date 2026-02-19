package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_requiresPixi_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_requiresPixi_getCmd).Standalone()

	workspace_requiresPixi_getCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_requiresPixiCmd.AddCommand(workspace_requiresPixi_getCmd)

	carapace.Gen(workspace_requiresPixi_getCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
