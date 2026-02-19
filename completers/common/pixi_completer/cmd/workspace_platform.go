package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_platformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Commands to manage workspace platforms",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platformCmd).Standalone()

	workspace_platformCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspaceCmd.AddCommand(workspace_platformCmd)

	carapace.Gen(workspace_platformCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
