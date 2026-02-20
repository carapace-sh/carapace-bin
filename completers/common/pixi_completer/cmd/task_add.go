package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var task_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a command to the workspace",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_addCmd).Standalone()

	task_addCmd.Flags().StringSlice("arg", nil, "The arguments to pass to the task")
	task_addCmd.Flags().Bool("clean-env", false, "Isolate the task from the shell environment, and only use the pixi environment to run the task")
	task_addCmd.Flags().String("cwd", "", "The working directory relative to the root of the workspace")
	task_addCmd.Flags().String("default-environment", "", "Add a default environment for the task")
	task_addCmd.Flags().StringSlice("depends-on", nil, "Depends on these other commands")
	task_addCmd.Flags().String("description", "", "A description of the task to be added")
	task_addCmd.Flags().StringSlice("env", nil, "The environment variable to set, use --env key=value multiple times for more than one variable")
	task_addCmd.Flags().StringP("feature", "f", "", "The feature for which the task should be added")
	task_addCmd.Flags().StringP("platform", "p", "", "The platform for which the task should be added")
	taskCmd.AddCommand(task_addCmd)

	carapace.Gen(task_addCmd).FlagCompletion(carapace.ActionMap{
		"cwd":      carapace.ActionDirectories(),
		"feature":  pixi.ActionFeatures(),
		"platform": pixi.ActionPlatforms(),
	})
}
