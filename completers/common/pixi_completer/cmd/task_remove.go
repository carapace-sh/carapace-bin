package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var task_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a command from the workspace",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_removeCmd).Standalone()

	task_removeCmd.Flags().StringP("environment", "e", "", "The environment for which the task should be removed. The task is removed from the tasks defined inline on the environment")
	task_removeCmd.Flags().StringP("feature", "f", "", "The feature for which the task should be removed")
	task_removeCmd.Flags().StringP("platform", "p", "", "The platform for which the task should be removed")
	taskCmd.AddCommand(task_removeCmd)

	carapace.Gen(task_removeCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
		"feature":     pixi.ActionFeatures(),
		"platform":    pixi.ActionPlatforms(),
	})

	carapace.Gen(task_removeCmd).PositionalAnyCompletion(
		pixi.ActionTasks(),
	)
}
