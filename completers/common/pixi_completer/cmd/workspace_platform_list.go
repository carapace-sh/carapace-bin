package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_platform_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the platforms in the workspace file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_listCmd).Standalone()

	workspace_platform_listCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_platformCmd.AddCommand(workspace_platform_listCmd)

	carapace.Gen(workspace_platform_listCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
