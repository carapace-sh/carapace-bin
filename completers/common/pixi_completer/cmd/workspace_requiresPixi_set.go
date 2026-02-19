package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_requiresPixi_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the pixi minimum version requirement",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_requiresPixi_setCmd).Standalone()

	workspace_requiresPixi_setCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_requiresPixiCmd.AddCommand(workspace_requiresPixi_setCmd)

	carapace.Gen(workspace_requiresPixi_setCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
