package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var task_aliasCmd = &cobra.Command{
	Use:     "alias",
	Short:   "Alias another specific command",
	Aliases: []string{"@"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_aliasCmd).Standalone()

	task_aliasCmd.Flags().String("description", "", "The description of the alias task")
	task_aliasCmd.Flags().StringP("environment", "e", "", "The environment for which the alias should be added. The alias is written to the tasks defined inline on the environment, creating the environment if it does not exist")
	task_aliasCmd.Flags().StringP("platform", "p", "", "The platform for which the alias should be added")
	taskCmd.AddCommand(task_aliasCmd)

	carapace.Gen(task_aliasCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
		"platform":    pixi.ActionPlatforms(),
	})

	carapace.Gen(task_aliasCmd).PositionalAnyCompletion(
		pixi.ActionTasks(),
	)
}
