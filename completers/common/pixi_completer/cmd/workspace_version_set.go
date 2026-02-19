package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_version_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the workspace version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_version_setCmd).Standalone()

	workspace_version_setCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_versionCmd.AddCommand(workspace_version_setCmd)

	carapace.Gen(workspace_version_setCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
