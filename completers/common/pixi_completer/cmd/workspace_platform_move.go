package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_platform_moveCmd = &cobra.Command{
	Use:     "move",
	Short:   "Reorder a workspace platform, changing its selection priority",
	Aliases: []string{"mv"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_moveCmd).Standalone()

	workspace_platform_moveCmd.Flags().String("after", "", "Move it directly after this platform")
	workspace_platform_moveCmd.Flags().String("before", "", "Move it directly before this platform")
	workspace_platform_moveCmd.Flags().Bool("no-install", false, "Don't update the environment, only refresh the lock-file")
	workspace_platform_moveCmd.Flags().Bool("to-bottom", false, "Move it to the bottom of the list (lowest selection priority)")
	workspace_platform_moveCmd.Flags().Bool("to-top", false, "Move it to the top of the list (highest selection priority)")
	workspace_platformCmd.AddCommand(workspace_platform_moveCmd)

	carapace.Gen(workspace_platform_moveCmd).FlagCompletion(carapace.ActionMap{
		"after":  pixi.ActionPlatforms(),
		"before": pixi.ActionPlatforms(),
	})

	carapace.Gen(workspace_platform_moveCmd).PositionalCompletion(
		pixi.ActionPlatforms(),
	)
}
