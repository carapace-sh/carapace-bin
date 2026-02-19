package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_platform_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove platform(s) from the workspace file and updates the lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_removeCmd).Standalone()

	workspace_platform_removeCmd.Flags().StringP("feature", "f", "", "The feature for which the dependency should be modified")
	workspace_platform_removeCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_platform_removeCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	workspace_platformCmd.AddCommand(workspace_platform_removeCmd)

	carapace.Gen(workspace_platform_removeCmd).FlagCompletion(carapace.ActionMap{
		"feature":       pixi.ActionFeatures(),
		"manifest-path": carapace.ActionFiles(),
	})

	carapace.Gen(workspace_platform_removeCmd).PositionalAnyCompletion(
		pixi.ActionPlatforms(),
	)
}
