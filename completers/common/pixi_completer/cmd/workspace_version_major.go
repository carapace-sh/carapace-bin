package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_version_majorCmd = &cobra.Command{
	Use:   "major",
	Short: "Bump the workspace version to MAJOR",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_version_majorCmd).Standalone()

	workspace_version_majorCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_versionCmd.AddCommand(workspace_version_majorCmd)

	carapace.Gen(workspace_version_majorCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
