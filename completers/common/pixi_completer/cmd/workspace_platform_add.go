package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_platform_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds a platform(s) to the workspace file and updates the lockfile",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_addCmd).Standalone()

	workspace_platform_addCmd.Flags().StringP("feature", "f", "", "The name of the feature to add the platform to")
	workspace_platform_addCmd.Flags().Bool("no-install", false, "Don't update the environment, only add changed packages to the lock-file")
	workspace_platformCmd.AddCommand(workspace_platform_addCmd)

	carapace.Gen(workspace_platform_addCmd).FlagCompletion(carapace.ActionMap{
		"feature": pixi.ActionFeatures(),
	})

	carapace.Gen(workspace_platform_addCmd).PositionalAnyCompletion(
		pixi.ActionKnownPlatforms(),
	)
}
