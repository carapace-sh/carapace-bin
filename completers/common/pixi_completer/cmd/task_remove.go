package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var task_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a command from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_removeCmd).Standalone()

	task_removeCmd.Flags().StringP("feature", "f", "", "The feature for which the task should be removed")
	task_removeCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	task_removeCmd.Flags().StringP("platform", "p", "", "The platform for which the task should be removed")
	taskCmd.AddCommand(task_removeCmd)

	carapace.Gen(task_removeCmd).FlagCompletion(carapace.ActionMap{
		"feature":       pixi.ActionFeatures(),
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
	})

	carapace.Gen(task_removeCmd).PositionalAnyCompletion(
		pixi.ActionTasks(),
	)
}
