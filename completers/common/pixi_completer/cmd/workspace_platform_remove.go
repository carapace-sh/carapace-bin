package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_platform_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove platform(s) from the workspace file and updates the lockfile",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_removeCmd).Standalone()

	workspace_platform_removeCmd.Flags().StringP("feature", "f", "", "The name of the feature to remove the platform from")
	workspace_platform_removeCmd.Flags().Bool("no-install", false, "Don't update the environment, only remove the platform(s) from the lock-file")
	workspace_platformCmd.AddCommand(workspace_platform_removeCmd)

	carapace.Gen(workspace_platform_removeCmd).FlagCompletion(carapace.ActionMap{
		"feature": pixi.ActionFeatures(),
	})

	carapace.Gen(workspace_platform_removeCmd).PositionalAnyCompletion(
		pixi.ActionPlatforms(),
	)
}
