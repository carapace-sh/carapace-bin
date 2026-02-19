package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_version_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the workspace version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_version_getCmd).Standalone()

	workspace_version_getCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_versionCmd.AddCommand(workspace_version_getCmd)

	carapace.Gen(workspace_version_getCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
